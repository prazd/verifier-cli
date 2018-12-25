package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"crypto/ecdsa"
	"context"
	"math/big"
	"strconv"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"os"
	"math"
	"github.com/onrik/ethrpc"
)

func TransferETH(amountStr, recipientStr string) {

	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		fmt.Println(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PKEY"))
	if err != nil {
		fmt.Println(err)
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("error casting public key to ECDSA")
	}

	ctx := context.Background()

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		fmt.Println(err)
	}

	amountFloat, err := strconv.ParseFloat(amountStr,64)
	if err != nil{
		fmt.Println(err)
	}

	amountWei := amountFloat * math.Pow(10,18)

	amount := big.NewInt(int64(amountWei)) // in wei

	gasLimit := uint64(21000) // in units

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		fmt.Println(err)
	}

	recipient := common.HexToAddress(recipientStr)

	var data []byte

	tx := types.NewTransaction(nonce, recipient, amount, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		fmt.Println(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex() + "\n")

}

func AccountBalance(accStr string){

	ethClient := ethrpc.New("https://rinkeby.infura.io")
	balance, err := ethClient.EthGetBalance(accStr, "latest")
	if err != nil {
		fmt.Println(err)
	}

	floatBalance, _ := strconv.ParseFloat(balance.String(), 64)

	ethBalance := floatBalance / math.Pow(10, 18)

	fmt.Println(ethBalance)
}