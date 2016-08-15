package service

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/sambaiz/go-api-with-test/dao"
	"github.com/sambaiz/go-api-with-test/model"
	"github.com/stretchr/testify/assert"
)

func TestMessageGetMessages(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	txMoc := dao.NewMockTx(ctl)

	messages := []model.Message{{ID: 1, Content: "メッセージ"}}
	messageDaoMoc := dao.NewMockMessage(ctl)
	messageDaoMoc.EXPECT().FindWithLimitOffset(uint64(1), uint64(1)).Return(messages, nil)

	messageService := NewMessage(txMoc, messageDaoMoc)

	actualMessages, err := messageService.GetMessages(1, 1)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, messages, actualMessages, "")
}

func TestMessageCreateMessage(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	txMoc := dao.NewMockTx(ctl)
	txMoc.EXPECT().Begin().Return(nil)
	txMoc.EXPECT().Commit().Return(nil)
	txMoc.EXPECT().RollbackUnlessCommitted()

	message := model.Message{ID: 1, Content: "メッセージ"}
	messageDaoMoc := dao.NewMockMessage(ctl)
	messageDaoMoc.EXPECT().Create(message.Content).Return(int64(1), nil)
	messageDaoMoc.EXPECT().FindById(message.ID).Return(&message, nil)

	messageService := NewMessage(txMoc, messageDaoMoc)

	actualMessage, err := messageService.CreateMessage(message.Content)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, message, *actualMessage, "")
}