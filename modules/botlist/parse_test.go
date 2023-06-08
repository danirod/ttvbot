package botlist

import (
	"testing"
)

const list = `
{
    "bots": [
        [
            "kattah",
            53433,
            1685949394
        ],
        [
            "drapsnatt",
            53089,
            1685949394
        ]
    ],
    "_total": 2
}
`

func TestUnmarshal(t *testing.T) {
	data, err := unmarshalResponse([]byte(list))
	if err != nil {
		t.Fail()
	}
	if len(data) != 2 {
		t.Fail()
	}
	if (data[0] != "kattah") || data[1] != "drapsnatt" {
		t.Fail()
	}
}
