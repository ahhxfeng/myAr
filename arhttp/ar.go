// send AR or wintson

package arhttp

import (
	"fmt"
	"math/big"

	// origin goar import
	// "github.com/everFinance/goar"
	// "github.com/everFinance/goar/types"

	"github.com/ahhxfeng/goar"
	"github.com/ahhxfeng/goar/types"
)

func SendAR() {
	wallet, err := goar.NewWalletFromPath("./key.json", "http://192.168.1.102:1984")
	if err != nil {
		panic(err)

	}

	tx, err := wallet.SendAR(
		big.NewFloat(1.0),
		"6qyjHgv-IPje-zZQj1S_9ncPEwGWGleVeSbTn0dZFJ0",
		[]types.Tag{},
	)

	fmt.Println(tx.ID, err)

}

// func SendData() {
// 	wallet, err := goar.NewWalletFromPath("./key.json", "http://192.168.1.102:1984")
// 	tx, err := wallet.SendData(
// 		[]byte("123"), // Data bytes
// 		[]types.Tag{
// 			types.Tag{
// 				Name:  "testSendData",
// 				Value: "123",
// 			},
// 		},
// 	)

// 	fmt.Println(id, err) // {{id}}, nil
// }
