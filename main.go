package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/DrDelphi/ElrondPrestaker/prestake"
	myToken "github.com/DrDelphi/ElrondPrestaker/token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

const (
	PrivateKey = ""

	GasLimit = 200000
	GasPrice = 30000000000      // in wei (30 gwei)
	MinETH   = 5000000000000000 // 0.005

	TokenAddress    = "0xF9986D445ceD31882377b5D6a5F58EaEa72288c3"
	PrestakeAddress = "0xE9F97FCa1C8C1dDC6649104A6dBCD400D073A445"
)

var (
	InfuraClient     *ethclient.Client
	TokenInstance    *myToken.Token
	PrestakeInstance *prestake.Token
)

func main() {
	pk := PrivateKey
	if len(os.Args) > 1 {
		pk = os.Args[1]
	}
	if err := prestakeNow(pk); err != nil {
		fmt.Println(err)
	}
}

func prestakeNow(pk string) error {
	var err error
	// Connect to Infura
	InfuraClient, err = ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		return errors.New("Unable to connect to Infura")
	}
	// Prepare contracts
	prestakeAddress := common.HexToAddress(PrestakeAddress)
	PrestakeInstance, _ = prestake.NewToken(prestakeAddress, InfuraClient)
	tokenAddress := common.HexToAddress(TokenAddress)
	TokenInstance, _ = myToken.NewToken(tokenAddress, InfuraClient)

	// Init privKey and generate public address
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return errors.New("Invalid private key")
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Verify balances
	eth, _ := InfuraClient.BalanceAt(context.Background(), fromAddress, nil)
	erd, _ := TokenInstance.BalanceOf(&bind.CallOpts{}, fromAddress)
	if erd.Int64() == 0 {
		return errors.New("No ERD in your account")
	}
	minEth := big.NewInt(MinETH)
	if eth.Cmp(minEth) == -1 {
		return errors.New("Not enough ETH for gas")
	}

	// Initialize approve transaction
	nonce, _ := InfuraClient.PendingNonceAt(context.Background(), fromAddress)
	value := big.NewInt(0)
	approveFnSignature := []byte("approve(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(approveFnSignature)
	approveMethodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(prestakeAddress.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(erd.Bytes(), 32)
	var data1 []byte
	data1 = append(data1, approveMethodID...)
	data1 = append(data1, paddedAddress...)
	data1 = append(data1, paddedAmount...)
	gasLimit := uint64(GasLimit)
	gasPrice := big.NewInt(GasPrice)
	tx1 := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data1)
	chainID, _ := InfuraClient.NetworkID(context.Background())
	signedTx1, _ := types.SignTx(tx1, types.NewEIP155Signer(chainID), privateKey)
	fmt.Println("Approve: " + signedTx1.Hash().Hex())
	err = InfuraClient.SendTransaction(context.Background(), signedTx1)
	if err != nil {
		return err
	}
	// Verify approve transaction status
	for {
		receipt, err := InfuraClient.TransactionReceipt(context.Background(), signedTx1.Hash())
		if err == nil {
			if receipt.Status != 1 {
				return errors.New(fmt.Sprintf("Failed: %v", receipt.Status))
			}
			break
		}
		time.Sleep(time.Second)
	}

	// Initialize deposit transaction
	nonce++
	depositFnSignature := []byte("deposit(uint256)")
	hash = sha3.NewLegacyKeccak256()
	hash.Write(depositFnSignature)
	depositMethodID := hash.Sum(nil)[:4]
	var data2 []byte
	data2 = append(data2, depositMethodID...)
	data2 = append(data2, paddedAmount...)
	tx2 := types.NewTransaction(nonce, prestakeAddress, value, gasLimit, gasPrice, data2)
	signedTx2, _ := types.SignTx(tx2, types.NewEIP155Signer(chainID), privateKey)
	fmt.Println("Deposit: " + signedTx2.Hash().Hex())
	err = InfuraClient.SendTransaction(context.Background(), signedTx2)
	if err != nil {
		return err
	}
	// Verify deposit transaction status
	for {
		receipt, err := InfuraClient.TransactionReceipt(context.Background(), signedTx2.Hash())
		if err == nil {
			if receipt.Status != 1 {
				return errors.New(fmt.Sprintf("Failed: %v", receipt.Status))
			}
			break
		}
		time.Sleep(time.Second)
	}
	fmt.Println("Done.")
	return nil
}
