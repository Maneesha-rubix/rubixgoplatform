package model

import "github.com/rubixchain/rubixgoplatform/core/wallet"

type TxnInfoResponse struct {
	Status     bool   `json:"status"`
	Msg        string `json:"message"`
	TxnDetails []wallet.TransactionDetails
}

type TxnInfoRequest struct {
	UserAddr string `json:"user"`
	TxnID    string `json:"txnId"`
}
