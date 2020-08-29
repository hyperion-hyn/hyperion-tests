package testScript

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"testing"
)

func TestGetValidatorInformation(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Printf("%v", err)
	}

	validatorAddress := common.HexToAddress("0x22082B79a5890A08d2a2a8Ee77994D7245090c3b")

	validator, err := client.GetValidatorInformation(context.Background(), validatorAddress, nil)

	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println(fmt.Sprintf("got validator address :%s", validator.Validator.ValidatorAddress.Hex()))

	for _, redelegation := range validator.Redelegations {
		fmt.Println(fmt.Sprintf("user: %s,amount: %s", redelegation.DelegatorAddress.Hex(), new(big.Int).Quo(redelegation.Amount, big.NewInt(params.Ether))))
	}

}