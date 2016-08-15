package dao

import (
	"github.com/gocraft/dbr"
	"errors"
)

// Tx is a transaction for the given Session
type Tx interface {
	Begin() error
	Commit() error
	RollbackUnlessCommitted()
}

type TxImpl struct {
	sess *dbr.Session
	tx *dbr.Tx
}

// NewTx returns an transaction instance for the given Session.
func NewTx(sess *dbr.Session) *TxImpl {
	return &TxImpl{sess: sess}
}

// Begin a transaction. If ever transaction has not finished, it fails.
func (t TxImpl) Begin() error {
	if t.tx != nil {
		return errors.New("transaction is unfinished")
	}

	if tx, err := t.sess.Begin(); err != nil {
		return err
	}else{
		t.tx = tx
		return nil
	}
}

// Commit finishes the transaction.
func (t TxImpl) Commit() error {
	err := t.tx.Commit()
	t.tx = nil
	return err
}

// Rollsback the transaction unless it has already been committed. Useful to defer tx.RollbackUnlessCommitted()
func (t TxImpl) RollbackUnlessCommitted() {
	t.tx.RollbackUnlessCommitted()
	t.tx = nil
}