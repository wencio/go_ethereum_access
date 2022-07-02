package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// We can switch from Infura to Ganache for demo purposes8
var infuraURL = "https://rinkeby.infura.io/v3/550dd3ed604f4342aaf4aa938937a274"
var ganacheURL = "http://localhost:8545"

func main() {

	// Get de ethereum client

	client, err := ethclient.DialContext(context.Background(), infuraURL)

	// Check for errors
	if err != nil {
		log.Fatalf("Error creating a ether client:%v", err)
	}

	// Close after finish
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)

	// Check for errors
	if err != nil {
		log.Fatalf("Error getting a block:%v", err)
	}

	// We took this address for demo purpose

	addr := "0xD6aE8250b8348C94847280928c79fb3b63cA453e"
	// Cast to Hexadecimal
	address := common.HexToAddress(addr)
	fmt.Println(address)
	// Check for Balance of the demo address
	balance, err := client.BalanceAt(context.Background(), address, nil)
	//Check for errors
	if err != nil {
		log.Fatalf("Error getting the balance %v", err)
	}

	// Convert from Wei to Ether
	fBlance := new(big.Float)
	fBlance.SetString(balance.String())
	balanceEther := new(big.Float).Quo(fBlance, big.NewFloat(math.Pow10(18)))
	fmt.Println("Balance in Ether:", balanceEther)
	//fmt.Println("The balance is:", balance)

	fmt.Println("The last block mined:", block.Number())
	// Another info from the  Block
	//fmt.Println(block.Difficulty())
	//fmt.Println(block.GasLimit())
	//fmt.Println(block.Hash())
	//fmt.Println(block.Header())

}
