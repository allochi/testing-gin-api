package main

import (
	"log"
	"testing-gin-api/mocks"

	"github.com/gin-gonic/gin"
)

type QueryAccountsRequest struct {
	Sort   string `form:"sort"`
	Desc   bool   `form:"desc"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type QueryAccountsResponse struct {
	Accounts []mocks.Account `json:"accounts"`
	Sort     string          `form:"sort"`
	Desc     bool            `form:"desc"`
	Limit    int             `form:"limit"`
	Offset   int             `form:"offset"`
	Error    string          `json:"error"`
}

func GetAccounts(c *gin.Context) {
	query := QueryAccountsRequest{}
	if err := c.BindQuery(&query); err != nil {
		log.Println("error binding query params", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var accounts []mocks.Account
	if query.Sort == "address" {
		accounts = mocks.GetSortedAccounts()
	} else {
		accounts = mocks.GetAccounts()
	}

	response := QueryAccountsResponse{
		Accounts: accounts,
		Sort:     query.Sort,
		Desc:     query.Desc,
		Limit:    query.Limit,
		Offset:   query.Offset,
	}

	c.JSON(200, response)
}
