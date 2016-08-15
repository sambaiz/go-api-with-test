package service

import (
	"github.com/sambaiz/go-api-with-test/dao"
	"github.com/sambaiz/go-api-with-test/model"
)

type Message interface {
	GetMessages(limit uint64, offset uint64) ([]model.Message, error)
	CreateMessage(content string) (*model.Message, error)
}

type MessageImpl struct {
	tx dao.Tx
	messageDao dao.Message
}

// NewMessage returns Message instance.
func NewMessage(tx dao.Tx, messageDao dao.Message) *MessageImpl {
	return &MessageImpl{tx: tx, messageDao: messageDao}
}

// Get messages.
func (m MessageImpl) GetMessages(limit uint64, offset uint64) ([]model.Message, error) {
	return m.messageDao.FindWithLimitOffset(limit, offset)
}

// Creates a message and returns it.
func (m MessageImpl) CreateMessage(content string) (*model.Message, error) {
	m.tx.Begin()

	defer m.tx.RollbackUnlessCommitted()

	id, err := m.messageDao.Create(content);
	if err != nil {
		return nil, err
	}

	message, err := m.messageDao.FindById(id)
	if err != nil {
		return nil, err
	}

	if err := m.tx.Commit(); err != nil {
		return nil, err
	}

	return message, nil
}
