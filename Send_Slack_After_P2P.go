package main

import (
    "fmt"
    "time"

    pdk "github.com/pavebank/pdk-go"
)

type Transaction struct {
    ID              string    `json:"id"`
    TransactionID   string    `json:"transaction_id"`
    Source          string    `json:"source"`
    SourceIban      string    `json:"source_iban"`
    Destination     string    `json:"destination"`
    DestinationIban string    `json:"destination_iban"`
    Amount          int       `json:"amount"`
    AssetType       string    `json:"asset_type"`
    Description     string    `json:"description"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}

type SlackMessage struct {
    WebHookUrl string `json:"webhook_url"`
    Text       string `json:"text"`
}

//export action_Log
func action_Log(x pdk.Args)

//export action_SlackSendMessage
func action_SlackSendMessage(x pdk.Args)

//export after_p2pTransaction
func after_p2pTransaction() {

    var transaction Transaction
    err := pdk.InputToStruct(&transaction)
    if err != nil {
        action_Log(pdk.BytesToArgs([]byte("Error")))
    }

    in := SlackMessage{
        WebHookUrl: "https://hooks.slack.com/services/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",
        Text: fmt.Sprintf("---------\n"+
            "Transaction `%s` has been successfully sent. \n"+
            "Source `%s` (IBAN: *%s*) \n"+
            "Destination: `%s` (IBAN: *%s*) \n"+
            "Amount: *%s %d* \n"+
            "Description: *%s*",
            transaction.ID,
            transaction.Source,
            transaction.SourceIban,
            transaction.Destination,
            transaction.DestinationIban,
            transaction.AssetType,
            transaction.Amount,
            transaction.Description,
        ),
    }

    offsetInput, err := pdk.StructToArgs(in)
    if err != nil {
        action_Log(pdk.BytesToArgs([]byte("Error")))
    }

    action_SlackSendMessage(offsetInput)

}

func main() {}
