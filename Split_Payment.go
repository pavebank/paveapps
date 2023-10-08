package main

import (
	"time"

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

type createVirtualAccount struct {
	ParentAccountID string
	Name            string
	Description     string
	GenerateIBAN    bool
}

type virtualAccountResult struct {
	ID          string
	Name        string
	Description string
	VIBAN       string
	CreatedAt   time.Time
}

//export action_Log
func action_Log(x pdk.Args)

//export action_CreateVirtualAccount
func action_CreateVirtualAccount(x pdk.Args) pdk.Args

//export action_CreateP2P
func action_CreateP2P(x pdk.Args) pdk.Args

//export after_p2pTransactionReceived
func after_p2pTransactionReceived() {

	p2p := new(p2pResult)
	err := pdk.InputToStruct(p2p)
	if err != nil {
		action_Log(pdk.BytesToArgs([]byte(err.Error())))
	}

    // This is necessary to avoid infinite cycle if both sender and receiver are in the same legal entity.
    // after_p2pTransactionReceived requires p2p.DestinationIBAN != "YOUR_DESTINATION_IBAN"
    // after_p2pTransactionSent requires p2p.SourceIBAN != "YOUR_SOURCE_IBAN"
	if p2p.DestinationIBAN != "GE81PV6560277468153462" {
		return
	}
    
    // Store params in memory for action_CreateVirtualAccount and pass it returned offset.
	argOffset, err := pdk.StructToArgs(createVirtualAccount{
		ParentAccountID: "account_dfssy530vggcbd5c9712sl296a", // Click account in the UI and copy account_id from browser.
		Name:            "Save 50%",
		Description:     "My virtual account for saving",
		GenerateIBAN:    true,
	})
	if err != nil {
		action_Log(pdk.BytesToArgs([]byte(err.Error())))
	}
	virtualAccount := new(virtualAccountResult)
	resOffset := action_CreateVirtualAccount(argOffset)
	err = pdk.ArgsToStruct(resOffset, virtualAccount)
	if err != nil {
		action_Log(pdk.BytesToArgs([]byte(err.Error())))
	}

    // Store params in memory for action_CreateP2P and pass it returned offset.
	argOffset, err = pdk.StructToArgs(createP2P{
		Description: "Move 50%",
		Source:      p2p.DestinationIBAN,
		Destination: "GE40PV5244091074561136", // Your IBAN where you want to move 50% of received money.
		Asset:       "USD",
		AssetType:   "FIAT",
		Amount:      (p2p.Amount / 100) * 50,
	})
	if err != nil {
		action_Log(pdk.BytesToArgs([]byte(err.Error())))
	}
	moveP2P := new(p2pResult)
	resOffset = action_CreateP2P(argOffset)
	err = pdk.ArgsToStruct(resOffset, moveP2P)
	if err != nil {
		action_Log(pdk.BytesToArgs([]byte(err.Error())))
	}
}

func main() {}
