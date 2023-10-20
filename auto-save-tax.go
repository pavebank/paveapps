package main

import (
	"github.com/pavebank/pdk-go"
)

type createP2P struct {
	Source      string
	Destination string
	Amount      int64
	Description string
	Asset       string
	AssetType   string
}

type p2pResult struct {
	TransactionID   string
	Source          string
	SourceIBAN      string
	Destination     string
	DestinationIBAN string
	Amount          int64
	Asset           string
	AssetType       string
	Description     string
	Status          string
}

//export action_Log
func action_Log(x pdk.Args)

//export action_CreateP2P
func action_CreateP2P(x pdk.Args) pdk.Args

//export after_p2pTransactionReceived
func after_p2pTransactionReceived() {

	p2p := new(p2pResult)
	err := pdk.InputToStruct(p2p)
	if err != nil {
		action_Log(pdk.BytesToArgs([]byte(err.Error())))
	}

	if p2p.DestinationIBAN != "CHANGEME" { // the iban of the tax pot
		return // do nothing if this payment is already going into the tax pot
	}

    if p2p.Amount < 200000 { // $2000 
		return // we're only going to save the tax if the inbound amount is greater than 2000 for this example
	}
  
    // Store params in memory for action_CreateP2P and pass it returned offset.
	argOffset, err := pdk.StructToArgs(createP2P{
		Description: "Moving 20% because the payment is greater than $2,000",
		Source:      p2p.DestinationIBAN,
		Destination: "CHANGEME", //the tax pot iban
		Asset:       "USD",
		AssetType:   "FIAT",
		Amount:      (p2p.Amount / 100) * 20,
	})
	if err != nil {
		action_Log(pdk.BytesToArgs([]byte(err.Error())))
	}
	moveP2P := new(p2pResult)
	resOffset := action_CreateP2P(argOffset)
	err = pdk.ArgsToStruct(resOffset, moveP2P)
	if err != nil {
		action_Log(pdk.BytesToArgs([]byte(err.Error())))
	}
}

func main() {}
