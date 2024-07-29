package tokenbizs

import (
	"context"
	"errors"
	"math/big"

	"backend-golang/blockchain"
	tokenstorages "backend-golang/modules/token/storage"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type TokenBiz interface {
	UploadData(ctx context.Context, docId string) error
	GetUploadSession(ctx context.Context, sessionId int64) (*blockchain.ControllerUploadSession, error)
	Confirm(ctx context.Context, docId, contentHash, proof string, sessionId, riskScore int64) error
	GetListSession(ctx context.Context) []blockchain.ControllerUploadSession
}

type TokenBizImpl struct {
	accountTxOpts *bind.TransactOpts
	controller    *blockchain.Controller
	tokenStorage  tokenstorages.TokenStorage
}

func NewTokenBiz(accountTxOpts *bind.TransactOpts, controller *blockchain.Controller, tokenStorage *tokenstorages.TokenStorageImpl) TokenBizImpl {
	return TokenBizImpl{
		accountTxOpts: accountTxOpts,
		controller:    controller,
		tokenStorage:  tokenStorage,
	}
}

func (t *TokenBizImpl) UploadData(ctx context.Context, docId string) error {
	_, err := t.controller.UploadData(t.accountTxOpts, docId)
	if err != nil {
		return err
	}

	return nil
}

func (t *TokenBizImpl) GetUploadSession(ctx context.Context, sessionId int64) (*blockchain.ControllerUploadSession, error) {
	currentSession, err := t.controller.GetSession(&bind.CallOpts{}, big.NewInt(sessionId))
	if err != nil {
		return nil, err
	}

	return &currentSession, nil
}

func (t *TokenBizImpl) Confirm(ctx context.Context, docId, contentHash, proof string, sessionId, riskScore int64) error {
	_, err := t.controller.Confirm(t.accountTxOpts, docId, contentHash, proof, big.NewInt(sessionId), big.NewInt(riskScore))
	if err != nil {
		return err
	}

	session, err := t.GetUploadSession(ctx, sessionId)
	if err != nil {
		return err
	}
	if session == nil {
		return errors.New("cannon get upload session")
	}

	err = t.tokenStorage.Confirm(ctx, session)
	if err != nil {
		return err
	}

	return nil
}

func (t *TokenBizImpl) GetListSession(ctx context.Context) []blockchain.ControllerUploadSession {
	return t.tokenStorage.GetListSession(ctx)
}
