// send AR or wintson

package arhttp

import (
	"fmt"
	"math/big"
	"github.com/everFinance/goar/types"
	"github.com/everFinance/goar"
)

func SendAR() bool {
	wallet, err := goar.NewWalletFromPath("./key.json", "http://120.79.132.16:1984")
	if err != nil {
		panic(err)
		return false
	}

	tx, err := wallet.SendAR(
		big.NewFloat(1.0),
		{{target}},
		[]types.Tag{},
	)

	fmt.Println(tx.ID, err)
	return true
	

}

func SendData() bool {
	wallet, err := goar.NewWalletFromPath("./key.json", "http://120.79.132.16:1984")
	tx, err := wallet.SendData(
		[]byte("123"), // Data bytes
		[]types.Tag{
		  types.Tag{
			Name:  "testSendData",
			Value: "123",
		  },
		},
	  )
	  
	  fmt.Println(id, err) // {{id}}, nil
}

