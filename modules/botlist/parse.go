package botlist

import (
	"bytes"
	"encoding/json"
)

type bot struct {
	username string
	channels int64
	lastSeen int64
}

func (b *bot) UnmarshalJSON(bs []byte) error {
	var (
		reader  = bytes.NewReader(bs)
		decoder = json.NewDecoder(reader)
		arr     []interface{}
	)
	decoder.UseNumber()
	if err := decoder.Decode(&arr); err != nil {
		return err
	}
	b.username = arr[0].(string)
	if num, err := arr[1].(json.Number).Int64(); err != nil {
		return err
	} else {
		b.channels = num
	}
	if num, err := arr[2].(json.Number).Int64(); err != nil {
		return err
	} else {
		b.lastSeen = num
	}
	return nil
}

type botsResponse struct {
	Bots  []bot `json:"bots"`
	Total uint  `json:"_total"`
}

func unmarshalResponse(data []byte) ([]string, error) {
	response := botsResponse{}
	if err := json.Unmarshal([]byte(data), &response); err != nil {
		return nil, err
	}

	botNames := make([]string, response.Total)
	for i, bot := range response.Bots {
		botNames[i] = bot.username
	}
	return botNames, nil
}
