package app

import (
	"bcos_lottery_demo/configs"
	utils "bcos_lottery_demo/pkg/http"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"strings"
)

// SaveTicket 存入交易数据
func SaveTicket(c *gin.Context) {
	//todo: add biz logic
	var (
		err            error
		ticket         Ticket
		lotterySession = configs.Conf.Client
	)
	if err = c.BindJSON(&ticket); err != nil {
		utils.BadResponse(c, 400, fmt.Sprintf("数据输入有误：%v", err))
		return
	}
	from := common.HexToAddress(ticket.PlayAddress)
	to := common.HexToAddress(ticket.PoolAddress)
	dataMap := map[string]interface{}{"amount": ticket.Amount, "timestamp": ticket.Timestamp, "lottery_num": ticket.LotteryNumber}
	dataBytes, _ := json.Marshal(dataMap)
	tx, _, err := lotterySession.SaveTransactionData(from, to, ticket.OrderHash, string(dataBytes)) // call Insert API
	if err != nil {
		utils.BadResponse(c, utils.Error, err.Error())
		return
	}
	utils.SimpleResponse(c, utils.OK, "", map[string]string{"hash": tx.Hash().Hex()})
}

// QueryTicket 查询数据
func QueryTicket(c *gin.Context) {
	var (
		resp           ContractQueryData
		lotterySession = configs.Conf.Client
	)
	txId := c.Query("txId")
	if txId == "" {
		utils.BadResponse(c, 400, "请指定交易hash")
		return
	}
	respX, err := lotterySession.GetTransactionData(txId) // call Insert API
	if err != nil {
		utils.BadResponse(c, utils.Error, err.Error())
		return
	} else {
		resp = ContractQueryData{
			From:  respX.Sender.String(),
			To:    respX.Receiver.String(),
			Value: respX.OrderHash,
			Data:  respX.Data,
		}
	}
	utils.SimpleResponse(c, utils.OK, "", resp)
}

func parseOutput(abiStr, name string, receipt *types.Receipt) (interface{}, error) {
	parsed, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		fmt.Printf("parse ABI failed, err: %v", err)
	}
	var ret interface{}
	b, err := hex.DecodeString(receipt.Output[2:])
	if err != nil {
		return nil, fmt.Errorf("decode receipt.Output[2:] failed, err: %v", err)
	}
	err = parsed.Unpack(&ret, name, b)
	if err != nil {
		return nil, fmt.Errorf("unpack %v failed, err: %v", name, err)
	}
	return ret, nil
}
