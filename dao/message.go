package dao

import (
	"github.com/gocraft/dbr"
	"github.com/sambaiz/go-api-with-test/model"
	"errors"
	"unicode/utf8"
)

// Message table DAO
type Message interface {
	Create(content string) (int64, error)
	FindWithLimitOffset(limit uint64, offset uint64) ([]model.Message, error)
	FindById(id int64) (*model.Message, error)
}

type MessageImpl struct {
	sess *dbr.Session
}

// NewMessage returns Message instance
func NewMessage(sess *dbr.Session) *MessageImpl {
	return &MessageImpl{sess: sess}
}

// Creates a message record.
func (t MessageImpl) Create(content string) (int64, error) {
	if(utf8.RuneCountInString(content) > 2000){
		return 0, errors.New("content is too long")
	}
	res, err := t.sess.InsertInto("message").Columns("content").Record(model.Message{Content: content}).Exec()
	if err != nil{
		return 0, err
	}
	return res.LastInsertId()
}

// Find messsage records with limit and offset order by ID.
func (t MessageImpl) FindWithLimitOffset(limit uint64, offset uint64) ([]model.Message, error) {
	var messages []model.Message
	_, err := t.sess.Select("*").From("message").Limit(limit).Offset(offset).OrderBy("id").Load(&messages)
	if err != nil{
		return nil, err
	}
	return messages, nil
}

// Find a messsage record by ID.
func (t MessageImpl) FindById(id int64) (*model.Message, error) {
	var messages []model.Message
	_, err := t.sess.Select("*").From("message").Where("id = ?", id).Load(&messages)
	if err != nil{
		return nil, err
	}
	if len(messages) == 0 {
		return nil, errors.New("message is not found")
	}
	return &messages[0], nil
}