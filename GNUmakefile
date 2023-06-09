GO_FILES=$(shell find . -name '*.go' -type f)

.PHONY: ttvbot

ttvbot: $(GO_FILES)
	go build ./cmd/ttvbot

README.txt: docs/ttvbot.8
	groff -man docs/ttvbot.8 -Tutf8 -P-c -P-b -P-u > README.txt

clean:
	rm -f ttvbot
