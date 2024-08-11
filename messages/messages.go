package messages

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Content     string
	IsEncrypted bool
	CreatedAt   time.Time
	ExpiresAt   time.Time
	Token       string
}

type MessagesMap map[string]*Message

type MessagesService struct {
	MessagesMap MessagesMap
	mMux        sync.Mutex
}

func NewMsgServices() *MessagesService {
	return &MessagesService{
		MessagesMap: make(MessagesMap),
	}
}

func (ms *MessagesService) AddMessage(content string) string {
	ms.mMux.Lock()
	defer ms.mMux.Unlock()

	token := uuid.New().String()
	now := time.Now()
	expiresAt := time.Now().Add(1 * time.Minute)
	msg := &Message{
		Content:   content,
		CreatedAt: now,
		ExpiresAt: expiresAt,
		Token:     token,
	}

	ms.MessagesMap[token] = msg
	log.Println(ms.MessagesMap)
	return token
}

func (ms *MessagesService) GetMessage(token string) (*Message, error) {
	ms.mMux.Lock()
	defer ms.mMux.Unlock()
	msg, exists := ms.MessagesMap[token]
	if !exists {
		return nil, errors.New("message doesn't exists")
	}
	return msg, nil
}

func (ms *MessagesService) DeleteMessage(token string) error {
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
