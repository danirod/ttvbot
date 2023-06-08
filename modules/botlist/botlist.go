package botlist

import (
	"fmt"
	"sync"
	"time"
)

func newBotList() *Botlist {
	return &Botlist{
		usernames:    make(map[string]bool),
		lastFetch:    time.UnixMilli(0),
		allowListBot: make(map[string]bool),
	}
}

type Botlist struct {
	mutex        sync.Mutex
	usernames    map[string]bool
	lastFetch    time.Time
	allowListBot map[string]bool
}

func (bots *Botlist) AllowBot(names ...string) {
	bots.mutex.Lock()
	defer bots.mutex.Unlock()

	for _, bot := range names {
		bots.allowListBot[bot] = true
	}
}

func (bots *Botlist) CountBots() int {
	bots.mutex.Lock()
	defer bots.mutex.Unlock()

	count := len(bots.usernames)
	return count
}

func (bots *Botlist) IsBot(username string) bool {
	bots.mutex.Lock()
	defer bots.mutex.Unlock()

	// First check if the username is in the allowlist.
	_, safe := bots.allowListBot[username]
	if safe {
		return false
	}

	// Otherwise, check the list of bots
	_, contains := bots.usernames[username]
	return contains
}

func (bots *Botlist) needsRefresh() bool {
	bots.mutex.Lock()
	defer bots.mutex.Unlock()

	diff := time.Since(bots.lastFetch)
	return diff.Minutes() > 10
}

func (bots *Botlist) fetchBots() error {
	bots.mutex.Lock()
	defer bots.mutex.Unlock()

	body, err := fetchBots()
	if err != nil {
		return err
	}

	usernames, err := unmarshalResponse(body)
	if err != nil {
		return err
	}

	nextUsernames := make(map[string]bool)
	for _, bot := range usernames {
		nextUsernames[bot] = true
	}
	bots.usernames = nextUsernames
	bots.lastFetch = time.Now()
	return nil
}

func (bots *Botlist) tryRefresh() error {
	var err error = nil
	if bots.needsRefresh() {
		fmt.Println("Bot list needs a refresh. Fetching...")
		err = bots.fetchBots()
		fmt.Printf("Fetched %d bots.\n", bots.CountBots())
	}
	return err
}
