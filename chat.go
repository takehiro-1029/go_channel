package main

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID       uuid.UUID `json:"id"`
	CreateAt time.Time `json:"created_at"`
	Body     string    `json:"body"`
}

type ChatService struct {
	mutex    sync.RWMutex
	messages map[uuid.UUID]Message
}

func (service *ChatService) CreateMessage(body string) (Message, error) {
	message := Message{
		ID:       uuid.New(),
		CreateAt: time.Now().UTC(),
		Body:     body,
	}

	service.mutex.Lock()
	service.messages[message.ID] = message
	service.mutex.Unlock()

	return message, nil
}

func (service *ChatService) FindMessages(after *uuid.UUID) ([]Message, error) {
	service.mutex.RLock()
	messages := make([]Message, 0, len(service.messages))
	for _, message := range service.messages {
		if after != nil && *after == message.ID {
			messages = []Message{message}
			break
		}
		messages = append(messages, message)
	}
	service.mutex.RUnlock()

	return messages, nil
}
