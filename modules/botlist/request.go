package botlist

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://api.twitchinsights.net/v1/bots/online"

func fetchBots() ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("invalid HTTP response")
	}
	return io.ReadAll(res.Body)
}
