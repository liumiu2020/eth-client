/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-11-20 10:49:55
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-11-21 12:30:26
 * @FilePath: \go-client\route\bid.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package route

import (
	"go-client/common"
	"go-client/contracts"
	"go-client/route/param"
	"math/big"
	"net/http"

	ethCommn "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func Bid(c *gin.Context) {
	var param = new(param.BidParam)
	err := c.BindJSON(&param)

	if err != nil {
		c.JSON(http.StatusExpectationFailed, "bid params analysis wrong")
	}

	client := common.EthClient{}

	bid, err := common.NewContract(contracts.NewBid)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, "bid init error")
	}
	// bookAddr:=common.HexToAddress("0xAb3B467dB70e99342fbF47de6AD74A1371518e20");
	bookAddr := ethCommn.HexToAddress("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0")

	var incr = new(big.Int)
	incr.SetInt64(param.BidAmount)
	nowAmount, err := bid.Bid(client.CallOption().CallOpt, 0, bookAddr, incr)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, "bid error")
	}
	c.JSON(http.StatusOK, nowAmount.String())
}
