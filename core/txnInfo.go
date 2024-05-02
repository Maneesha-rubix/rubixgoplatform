package core

import (
	"github.com/rubixchain/rubixgoplatform/core/wallet"
)

func (c *Core) GetTransInfo(txnID string, didstr string) (wallet.TransactionDetails, error) {
	var txnInf wallet.TransactionDetails
	res, err := c.w.GetTxnInfo(txnID, didstr)
	if err != nil {
		return txnInf, err
	}
	return res, nil
}
