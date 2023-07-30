// send AR or wintson

package arhttp

import (
	"fmt"
	"math/big"
	"os"

	// origin goar import
	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	// my import
	// "github.com/ahhxfeng/goar"
	// "github.com/ahhxfeng/goar/types"
)

func SendAR() {
	wallet, err := goar.NewWalletFromPath("./key.json", "http://192.168.1.126:1984")
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

func SendData() {
	wallet, err := goar.NewWalletFromPath("./key.json", "http://192.168.1.126:1984")
	if err != nil {
		panic(err)
	}
	data, err := os.Open("./weave.jpg") // local file path
	if err != nil {
		panic(err)
	}
	defer data.Close()
	tags := []types.Tag{
		{Name: "Content-Type", Value: "img/jpeg"},
		{Name: "test", Value: "kevin-test"},
	}
	tx, err := wallet.SendDataStreamSpeedUp(data, tags, 10)
	fmt.Printf("Transcation id: %v, Error: %v", tx.ID, err)
	// fmt.Sprintf("Transcation id : %v, Error: %v", tx.ID, err)
}
