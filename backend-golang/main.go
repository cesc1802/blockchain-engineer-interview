package main

import (
	"context"
	"math/big"
	"os"

	contracts "backend-golang/blockchain"
	"backend-golang/component/appctx"
	"backend-golang/events"
	"backend-golang/middleware"
	tokenbizs "backend-golang/modules/token/biz"
	tokenstorages "backend-golang/modules/token/storage"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})

	assetDomain := os.Getenv("ASSET_DOMAIN")
	secretKey := os.Getenv("SECRET_KEY")

	appCtx := appctx.NewAppContext(nil, assetDomain, secretKey)

	rpcURL := os.Getenv("RPC_URL")
	clientCtx := context.Background()
	client, err := ethclient.DialContext(clientCtx, rpcURL)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(os.Getenv("SMART_CONTRACT_CONTROLLER_ADDRESS"))

	userHEXPrivateKey := os.Getenv("USER_HEX_PRIVATE_KEY")
	userPrivateKey, err := crypto.HexToECDSA(userHEXPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(clientCtx)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(userPrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	auth.GasLimit = uint64(0)
	auth.GasPrice = big.NewInt(2000000000)

	controller, err := contracts.NewController(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tokenStoreImpl := tokenstorages.NewTokenStore()
	tokenBiz := tokenbizs.NewTokenBiz(auth, controller, &tokenStoreImpl)

	event := events.NewUploadSessionEvent(rpcURL, &tokenStoreImpl, &tokenBiz)
	go event.Run(clientCtx)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.Use(middleware.AllowCORS())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")

	Route(v1, appCtx, &tokenBiz)

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
