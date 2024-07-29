package events

import (
	"context"
	"math/big"
	"os"
	"time"

	"backend-golang/blockchain"
	"backend-golang/component/asyncjob"
	tokenbizs "backend-golang/modules/token/biz"
	tokenstorages "backend-golang/modules/token/storage"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type UploadSessionEvent struct {
	rpcURL       string
	tokenStorage tokenstorages.TokenStorage
	tokenBiz     tokenbizs.TokenBiz
}

func NewUploadSessionEvent(rpcURL string, tokenStorage *tokenstorages.TokenStorageImpl, tokenBiz *tokenbizs.TokenBizImpl) UploadSessionEvent {
	return UploadSessionEvent{
		rpcURL:       rpcURL,
		tokenStorage: tokenStorage,
		tokenBiz:     tokenBiz,
	}
}

func (e *UploadSessionEvent) Run(ctx context.Context) {
	client, err := ethclient.DialContext(ctx, e.rpcURL)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(os.Getenv("SMART_CONTRACT_CONTROLLER_ADDRESS"))

	userHEXPrivateKey := os.Getenv("USER_HEX_PRIVATE_KEY")
	userPrivateKey, err := crypto.HexToECDSA(userHEXPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(userPrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	auth.GasLimit = uint64(0)
	auth.GasPrice = big.NewInt(2000000000)

	controller, err := blockchain.NewController(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	controllerEventChan := make(chan *blockchain.ControllerUploadData)
	sub, err := controller.WatchUploadData(&bind.WatchOpts{Context: ctx, Start: nil}, controllerEventChan)
	if err != nil {
		log.Println(err)
	}
	defer sub.Unsubscribe()

	for event := range controllerEventChan {
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return e.handleUploadSession(ctx, event)
		})

		job.SetRetryDurations(time.Second, time.Second, time.Second, time.Second) // 4 times (1s each)

		if err := job.Execute(ctx); err != nil {
			log.Errorln(err)
		}
	}
}

func (e *UploadSessionEvent) handleUploadSession(ctx context.Context, event *blockchain.ControllerUploadData) error {
	uploadSession, err := e.tokenBiz.GetUploadSession(ctx, event.SessionId.Int64())
	if err != nil {
		return err
	}

	e.tokenStorage.SaveSession(ctx, uploadSession)
	return nil
}
