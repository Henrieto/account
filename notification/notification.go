package notification

import (
	"errors"
	"sync"
)

var (
	Manager *NotificationManager = NewManager()
)

type Notifier interface {
	Send(any) error
}

type NotificationManager struct {
	notifiers map[string]Notifier
	mutex     sync.Mutex
}

func NewManager() *NotificationManager {
	return &NotificationManager{
		notifiers: map[string]Notifier{},
	}
}

func (ntm *NotificationManager) Notify(channel string, payload any) error {
	notifier, ok := ntm.notifiers[channel]
	if !ok {
		return errors.New("notifier for the channel is not available")
	}
	err := notifier.Send(payload)
	if err != nil {
		return err
	}
	return nil
}

func (ntm *NotificationManager) Register(channel string, notifier Notifier) {
	ntm.mutex.Lock()
	defer ntm.mutex.Unlock()
	ntm.notifiers[channel] = notifier
}
