package main

import (
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CRUDContract contract with simple create, read, update delete
type CRUDContract struct {
	contractapi.Contract
}

// Create adds to the world state
func (c *CRUDContract) Create(ctx contractapi.TransactionContextInterface, key string, value string) error {
	return ctx.GetStub().PutState(key, []byte(value))
}

// Read adds to the world state
func (c *CRUDContract) Read(ctx contractapi.TransactionContextInterface, key string, value string) (string, error) {
	bytes, err := ctx.GetStub().GetState(key)

	if err != nil {
		return "", err
	}

	if bytes == nil {
		return "", errors.New("Key not found")
	}

	return string(bytes), nil
}

// Update adds to the world state
func (c *CRUDContract) Update(ctx contractapi.TransactionContextInterface, key string, value string) error {
	return ctx.GetStub().PutState(key, []byte(value))
}

// Delete removes from the world state
func (c *CRUDContract) Delete(ctx contractapi.TransactionContextInterface, key string, value string) error {
	return ctx.GetStub().DelState(key)
}

func main() {
	contract := new(CRUDContract)

	cc, err := contractapi.NewChaincode(contract)

	if err != nil {
		panic(err.Error())
	}

	err = cc.Start()

	if err != nil {
		panic(err.Error())
	}
}
