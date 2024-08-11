package messages

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Content   string
	CreatedAt time.Time
	ExpiresAt time.Time
	Token     uuid.UUID
}

type MessagesMap map[uuid.UUID]*Message

type MessagesService struct {
	MessagesMap MessagesMap
	mMux        sync.Mutex
}

func NewMsgServices() *MessagesService {
	return &MessagesService{
		MessagesMap: make(MessagesMap),
	}
}

func (ms *MessagesService) AddMessage(content string) uuid.UUID {
	ms.mMux.Lock()
	defer ms.mMux.Unlock()

	token := uuid.New()
	now := time.Now()
	expiresAt := time.Now().Add(1 * time.Minute)
	msg := &Message{
		Content:   content,
		CreatedAt: now,
		ExpiresAt: expiresAt,
		Token:     token,
	}

	ms.MessagesMap[token] = msg

	return token
}

func (ms *MessagesService) DeleteMessage(token uuid.UUID) error {
	ms.mMux.Lock()
	defer ms.mMux.Unlock()

	if _, exists := ms.MessagesMap[token]; !exists {
		return errors.New("message not found")
	}

	delete(ms.MessagesMap, token)
	return nil
}

func (ms *MessagesService) CleanExpiredMessages() {
	ms.mMux.Lock()
	defer ms.mMux.Unlock()

	now := time.Now()
	for token, msg := range ms.MessagesMap {
		if now.After(msg.ExpiresAt) {
			delete(ms.MessagesMap, token)
		}
	}
}

func (ms *MessagesService) CleanExpiredMessagesLoop() {
	for {
		ms.CleanExpiredMessages()
		log.Println("<Msg> Cleaned Expired Messages")
		time.Sleep(4 * (time.Hour))
	}
}
