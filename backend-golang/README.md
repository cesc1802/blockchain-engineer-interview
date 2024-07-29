##### Table of contents
- [Start the Golang Backend](#start-the-golang-backend)
  - [Requirements](#requirements)
  - [In the genomicdao repository](#in-the-genomicdao-repository)
    - [The Controller smart contract address after deployed](#the-controller-smart-contract-address-after-deployed)
    - [Hardhat Node User Accounts and Private Keys for testing purpose](#hardhat-node-user-accounts-and-private-keys-for-testing-purpose)
  - [In the golang-backend repository](#in-the-golang-backend-repository)

# Start the Golang Backend
## Requirements
- The **genomicdao** repository, ensure dependencies version:
  - **@nomicfoundation/hardhat-toolbox** - ^5.0.0
  - **hardhat** - ^2.22.6
  - **@openzeppelin/contracts** - ^4.9.3
  - **solcjs** - ^0.8.26
  - **abigen** - latest
- The **golang-backend** repository:
  - **Go Version** - 1.22.5
## In the genomicdao repository
- Install dependencies by this command:
  ```bash
  npm install
  ```
- Run Hardhat Node by this command:
  ```bash
  npx hardhat node
  ```
  Or
  ```bash
  make node
  ```
- Deploy the Controller smart contract by this command:
  ```bash
  npx hardhat run ./scripts/deploy.js --network localhost
  ```
  Or
  ```bash
  make deploy
  ```
- If you want to generate abi and go files:
  - Install abigen by running this command, make sure you have installed go and export ~/go/bin in the $PATH:
    ```bash
    go install github.com/ethereum/go-ethereum/cmd/abigen@latest
    ```
  - Compile to abi and generate go files by running this commmand:
    ```bash
    make compile
    ```
### The Controller smart contract address after deployed
```
0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0
```
- In case the **Controller smart contract address** does not match, look into this log(**Contract address**) after running the Hardhat node:
  ```
  eth_sendTransaction
    Contract deployment: Controller
    Contract address:    0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0
    Transaction:         0xb9008b5d1bf3d6afa33fb3ef512d8c36d92f72003ce5b156b4db34cd8d5c7a88
    From:                0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266
    Value:               0 ETH
    Gas used:            1356994 of 30000000
    Block #3:            0x7f48c650543f5f6c949d4cece80f0a9f2e51a6941c90e4dc0a0b44ad0e2d294a
  ```

### Hardhat Node User Accounts and Private Keys for testing purpose
```
Account #0: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 (10000 ETH)
Private Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

Account #1: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 (10000 ETH)
Private Key: 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
...
```

## In the golang-backend repository
- Start the server by this command:
  ```bash
  RPC_URL="ws://127.0.0.1:8545/ws" \
  SMART_CONTRACT_CONTROLLER_ADDRESS="0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0" \
  USER_ACCOUNT="0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266" \
  USER_HEX_PRIVATE_KEY="ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" \
  go run .
  ```
In this:
  - **RPC_URL** is used to starts a JSON-RPC server on top of Hardhat Network. In this case, we use the server on localhost.
  - **SMART_CONTRACT_CONTROLLER_ADDRESS** is the address shown in the log of the Controller smart contract after we deployed.
  - **USER_ACCOUNT** is the account we used to do the transaction on the Controller smart contract.
  - **USER_HEX_PRIVATE_KEY** is the account's private key we used to sign the transaction on the Controller smart contract.