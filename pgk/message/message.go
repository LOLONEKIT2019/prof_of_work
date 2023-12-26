package message

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	messageSplitter = "|"
)

type (
	Type int

	Message struct {
		Type Type
		Body string
	}
)

func ParseMessage(data string) (*Message, error) {

	parts := strings.Split(data, messageSplitter)
	if len(parts) < 1 || len(parts) > 2 {
		return nil, fmt.Errorf("invalid message format")
	}

	messageType, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid message type")
	}

	msg := Message{Type: Type(messageType)}

	if len(parts) == 2 {
		msg.Body = parts[1]
	}

	return &msg, nil
}

func (m *Message) String() string {
	return fmt.Sprintf("%d|%s", m.Type, m.Body)
}
