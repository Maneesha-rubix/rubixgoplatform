package server

import (
	"fmt"
	"net/http"

	"github.com/rubixchain/rubixgoplatform/core/model"
	"github.com/rubixchain/rubixgoplatform/core/wallet"
	"github.com/rubixchain/rubixgoplatform/util"
	"github.com/rubixchain/rubixgoplatform/wrapper/ensweb"
)

// @Summary Get transaction details by Transcation ID
// @Description Retrieves the details of a transaction based on its ID.
// @ID get-txn-details-by-id
// @Tags         Account
// @Accept json
// @Produce json
// @Param txnID query string true "The ID of the transaction to retrieve"
// @Success 200 {object} model.BasicResponse
// @Router /api/get-txn-info [get]
func (s *Server) APIGetTransInfo(req *ensweb.Request) *ensweb.Result {
	txnId := s.GetQuerry(req, "txnID")
	userAddr := s.GetQuerry(req, "userAddr")

	_, didstr, ok := util.ParseAddress(userAddr)
	if !ok {
		fmt.Println("couldn't parse userAddr")
		return s.BasicResponse(req, false, "Invalid user address", nil)
	}
	// if !s.validateDIDAccess(req, did) {
	// 	fmt.Println("couldn't validate did access")
	// 	return s.BasicResponse(req, false, "DID does not have an access", nil)
	// }

	res, err := s.c.GetTransInfo(txnId, didstr)
	if err != nil {
		if err.Error() == "no records found" {
			s.log.Info("User doesn't have any information on Txn ID " + txnId)
			transResp := &model.TxnInfoResponse{
				Status:     false,
				Msg:        "User doesn't have any information on Txn ID " + txnId,
				TxnDetails: make([]wallet.TransactionDetails, 0),
			}
			return s.RenderJSON(req, &transResp, http.StatusOK)
		}

		s.log.Info("error", err)
		transResp := &model.TxnInfoResponse{
			Status:     false,
			Msg:        "User doesn't have any information on Txn ID " + txnId,
			TxnDetails: make([]wallet.TransactionDetails, 0),
		}
		return s.RenderJSON(req, &transResp, http.StatusOK)
	}

	transResp := &model.TxnInfoResponse{
		Status:     true,
		Msg:        "Transaction info fetched successfully",
		TxnDetails: make([]wallet.TransactionDetails, 0),
	}
	transResp.TxnDetails = append(transResp.TxnDetails, res)

	return s.RenderJSON(req, &transResp, http.StatusOK)

}
