compile:
	solc --abi --include-path node_modules/ --base-path . ./contracts/NFT.sol -o ./blockchain --overwrite
	solc --abi --include-path node_modules/ --base-path . ./contracts/Token.sol -o ./blockchain --overwrite
	solc --abi --include-path node_modules/ --base-path . ./contracts/Controller.sol -o ./blockchain --overwrite
# go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	abigen --abi ./blockchain/GeneNFT.abi --pkg blockchain --type GeneNFT --out ./blockchain/GeneNFT.go
	abigen --abi ./blockchain/PostCovidStrokePrevention.abi --pkg blockchain --type PostCovidStrokePrevention --out ./blockchain/PostCovidStrokePrevention.go
	abigen --abi ./blockchain/Controller.abi --pkg blockchain --type Controller --out ./blockchain/Controller.go
deploy:
	npx hardhat run ./scripts/deploy.js --network localhost
node:
	npx hardhat node
