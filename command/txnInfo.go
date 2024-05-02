package command

import "fmt"

func (cmd *Command) getTransInfo() {
	msg, err := cmd.c.GetTransInfo(cmd.userAddr, cmd.txnID)
	if err != nil {
		cmd.log.Error("failed to get transaction info", "err", err)
		cmd.log.Info(msg.Msg)
		return
	}
	if !msg.Status {
		cmd.log.Error("failed to get transaction info of txn ID", cmd.txnID)
		cmd.log.Info(msg.Msg)
	} else {
		cmd.log.Info(msg.Msg)
		fmt.Printf("TransactionID : %s\n, TransactionType : %s\n, BlockID : %s\n, Mode : %d\n, SenderDID : %s\n, ReceiverDID %s\n, Amount : %10.5f\n, TotalTime : %10.5f\n, Comment : %s\n, DateTime : %s\n, Status : %t\n, DeployerDID : %s\n", msg.TxnDetails[0].TransactionID, msg.TxnDetails[0].TransactionType, msg.TxnDetails[0].BlockID, msg.TxnDetails[0].Mode, msg.TxnDetails[0].SenderDID, msg.TxnDetails[0].ReceiverDID, msg.TxnDetails[0].Amount, msg.TxnDetails[0].TotalTime, msg.TxnDetails[0].Comment, msg.TxnDetails[0].DateTime.String(), msg.TxnDetails[0].Status, msg.TxnDetails[0].DeployerDID)
	}
}
