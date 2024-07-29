package tokenstorages

import (
	"context"
	"errors"
	"sync"

	"backend-golang/blockchain"
)

type SessionId int64
type DocId string
type NTFId int64

type TokenStorage interface {
	SaveSession(context.Context, *blockchain.ControllerUploadSession)
	Confirm(context.Context, *blockchain.ControllerUploadSession) error
	GetListSession(ctx context.Context) []blockchain.ControllerUploadSession
}

type TokenStorageImpl struct {
	mu         sync.Mutex
	Sessions   map[SessionId]blockchain.ControllerUploadSession
	Docs       map[DocId]blockchain.ControllerDataDoc
	NTFDocs    map[NTFId]string
	DocSubmits map[DocId]bool
}

func NewTokenStore() TokenStorageImpl {
	return TokenStorageImpl{
		Sessions:   map[SessionId]blockchain.ControllerUploadSession{},
		Docs:       map[DocId]blockchain.ControllerDataDoc{},
		NTFDocs:    map[NTFId]string{},
		DocSubmits: map[DocId]bool{},
	}
}

func (t *TokenStorageImpl) SaveSession(ctx context.Context, uploadSession *blockchain.ControllerUploadSession) {
	t.mu.Lock()
	t.Sessions[SessionId(uploadSession.Id.Int64())] = *uploadSession
	t.mu.Unlock()
}

func (t *TokenStorageImpl) Confirm(ctx context.Context, session *blockchain.ControllerUploadSession) error {
	t.mu.Lock()
	if _, ok := t.Sessions[SessionId(session.Id.Int64())]; !ok {
		return errors.New("session does not exist")
	}
	t.Sessions[SessionId(session.Id.Int64())] = *session
	t.mu.Unlock()
	return nil
}

func (t *TokenStorageImpl) GetListSession(ctx context.Context) []blockchain.ControllerUploadSession {
	t.mu.Lock()
	var uploadSessions []blockchain.ControllerUploadSession
	for _, s := range t.Sessions {
		uploadSessions = append(uploadSessions, s)
	}
	t.mu.Unlock()
	return uploadSessions
}
