zetrix-sdk-go
=======
A complete and simple library for Go developers to connect and use the Zetrix layer 1 blockchain.

## Docs & Useful Links

  * [SDK Documentation](https://docs.zetrix.com/en/developer-resources/sdk/go)
  * [Zetrix Explorer](https://explorer.zetrix.com)
  * [Zetrix Testnet Faucet](https://faucet.zetrix.com)
  * [Zetrix Smart Contract IDE](https://ide.zetrix.com/)
  * [Zetrix Wallet](https://www.zetrix.com/zetrix-wallet/)

## Installation & Prerequisite
Ensure your Go version is 1.10.1 or above.

To use the Zetrix Go SDK in a Go project, you can use the following commands to install the package:
#### Download packages

```sh
$ go get github.com/Zetrix-Chain/zetrix-sdk-go
```

#### Import the package in your project
```go
import (
	"github.com/Zetrix-Chain/zetrix-sdk-go/src/model"
	"github.com/Zetrix-Chain/zetrix-sdk-go/src/sdk"
)
```

## Quick Start & Basic Usages
Initialize the SDK with the Zetrix node URL:
#### Testnet
```go
var zetrixSDK sdk.Sdk
var reqData model.SDKInitRequest
reqData.SetUrl("https://test-node.zetrix.com")
resData := zetrixSDK.Init(reqData)
if resData.ErrorCode != 0 {
	fmt.Println(resData.ErrorDesc)
} else {
	fmt.Println("ZetrixSDK initialized.")
}
```

#### Mainnet
```go
var zetrixSDK sdk.Sdk
var reqData model.SDKInitRequest
reqData.SetUrl("https://node.zetrix.com")
resData := zetrixSDK.Init(reqData)
if resData.ErrorCode != 0 {
	fmt.Println(resData.ErrorDesc)
} else {
	fmt.Println("ZetrixSDK initialized.")
}
```

Getting the nonce for a particular address:
```go
var reqData model.AccountGetNonceRequest
reqData.SetAddress(address) // address in a string
resData := zetrixSDK.Account.GetNonce(reqData)
if resData.ErrorCode != 0 {
    return resData.ErrorDesc
} else {
    return resData.Result.Nonce
}
```

Submitting a transaction:
```go
var reqData model.TransactionSubmitRequest
reqData.SetBlob(txBlob) // txBlob in a string
reqData.SetSignatures(signedBlob) // signedBlob in a string
resDataSubmit := zetrixSDK.Transaction.Submit(reqData)
if resDataSubmit.ErrorCode != 0 {
    return resDataSubmit.ErrorDesc
} else {
    return resDataSubmit.Result.Hash
}
```

Getting the balance of a particular address:
```go
var reqData model.AccountGetBalanceRequest
reqData.SetAddress(address) // address in a string
resData := zetrixSDK.Account.GetBalance(reqData)
if resData.ErrorCode != 0 {
    return resData.ErrorDesc
} else {
    return resData.Result.Balance
}
```

Getting transaction info:
```go
var reqData model.TransactionGetInfoRequest
reqData.SetHash(hash) // hash in a string
resData := zetrixSDK.Transaction.GetInfo(reqData)
if resData.ErrorCode != 0 {
    return resData.ErrorDesc
} else {
    return resData.Result
}
```