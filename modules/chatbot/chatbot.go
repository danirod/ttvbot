package chatbot

import (
	"fmt"
	"sync/atomic"

	"github.com/gempir/go-twitch-irc/v4"
)

func newTwitchClient(botUsername, botToken string) *twitch.Client {
	client := twitch.NewClient(botUsername, botToken)
	client.Capabilities = append(
		client.Capabilities,
		twitch.CommandsCapability,
		twitch.MembershipCapability,
		twitch.TagsCapability,
	)
	return client
}

func newChatbot(botUsername, botToken, targetChannel string) *Chatbot {
	client := newTwitchClient(botUsername, botToken)
	client.Join(targetChannel)
	return &Chatbot{
		client:        client,
		botUsername:   botUsername,
		targetChannel: targetChannel,
		onJoin:        []JoinFunc{},
		onLeave:       []LeaveFunc{},
		onMessage:     []MessageFunc{},
	}
}

type MessageFlags struct {
	First bool
}

type JoinFunc func(channel, username string)
type LeaveFunc func(channel, username string)
type MessageFunc func(channel, username, content string, flags *MessageFlags)

type Chatbot struct {
	client        *twitch.Client
	connected     atomic.Bool
	botUsername   string
	targetChannel string

	onJoin    []JoinFunc
	onLeave   []LeaveFunc
	onMessage []MessageFunc
}

func (bot *Chatbot) SendMessage(content string) {
	bot.client.Say(bot.targetChannel, content)
}

func (bot *Chatbot) OnUserJoin(joinFunc JoinFunc) {
	bot.onJoin = append(bot.onJoin, joinFunc)
}

func (bot *Chatbot) OnUserLeave(leaveFunc LeaveFunc) {
	bot.onLeave = append(bot.onLeave, leaveFunc)
}

func (bot *Chatbot) OnMessage(msgFunc MessageFunc) {
	bot.onMessage = append(bot.onMessage, msgFunc)
}

func (bot *Chatbot) connect() error {
	bot.client.OnUserJoinMessage(func(message twitch.UserJoinMessage) {
		for _, function := range bot.onJoin {
			function(message.Channel, message.User)
		}
	})
	bot.client.OnUserPartMessage(func(message twitch.UserPartMessage) {
		for _, function := range bot.onLeave {
			function(message.Channel, message.User)
		}
	})
	bot.client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		var flags MessageFlags
		flags.First = message.FirstMessage
		for _, function := range bot.onMessage {
			function(message.Channel, message.User.Name, message.Message, &flags)
		}
	})
	bot.client.OnSelfJoinMessage(func(message twitch.UserJoinMessage) {
		fmt.Println("The bot has joined", message.Channel)
	})
	bot.client.OnConnect(func() {
		bot.connected.Store(true)
		fmt.Println("The bot is connected to Twitch")
	})

	// TODO: Connect() actually blocks. Move this to a local goroutine.
	return bot.client.Connect()
}

func (bot *Chatbot) disconnect() error {
	defer bot.connected.Store(false)
	return bot.client.Disconnect()
}
