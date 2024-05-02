package client

import (
	"github.com/rubixchain/rubixgoplatform/core/model"
	"github.com/rubixchain/rubixgoplatform/setup"
)

func (c *Client) GetTransInfo(userAddr string, txnId string) (*model.TxnInfoResponse, error) {

	mp := make(map[string]string)
	mp["userAddr"] = userAddr
	mp["txnID"] = txnId

	var resp model.TxnInfoResponse

	err := c.sendJSONRequest("GET", setup.APIGetTransInfo, mp, nil, &resp)
	if err != nil {
		resp.Status = false
		return nil, err
	}
	return &resp, nil

}
