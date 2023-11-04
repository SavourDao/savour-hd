package near

import (
	"fmt"
	"testing"

	"github.com/savour-labs/wallet-hd-chain/config"
)

func newTestClient() *NearClient {

	client, _ := newNearClients(&config.Config{Fullnode: config.Fullnode{
		Near: config.Node{
			RPCs: []*config.RPC{
				{
					RPCURL: "https://rpc.mainnet.near.org",
				},
			},
		},
	},
	})
	return client[0]
}

func TestGetLatestBlockHeight(t *testing.T) {
	client := newTestClient()
	height, _ := client.GetLatestBlockHeight()
	fmt.Println(height)
}

func TestGetTx(t *testing.T) {
	client := newTestClient()
	height, _ := client.GetTx("9872cefe5ba026ba4806bd0ab9e4ebe77790febd31f09813eb9c798a19634e45", 1, 5)
	fmt.Println(height)
}

func TestGetTxByHash(t *testing.T) {
	client := newTestClient()
	tx, _ := client.GetTxByHash("ZpsS5Ahkvq2AybQUDHe4YnZmmXKtd5BCV48CxZXGvaK")
	fmt.Println(tx)
}

func TestGetActionByHash(t *testing.T) {
	client := newTestClient()
	tx, _ := client.GetActionByHash([]string{"3QrSMQN7VYV6HtM8czwKJ4sA2AENN8xSMGwqTAWvrpWN"})
	fmt.Println(tx)
}

func TestGetBalance(t *testing.T) {
	client := newTestClient()
	balance, _ := client.GetBalance("c7aada6fdd947d56a0bf58c57c579a601f31e5610d8e3c2883cd2ce26109bba7")
	fmt.Println(balance)
}

func TestGetAccount(t *testing.T) {
	client := newTestClient()
	pri, address, _ := client.GetAccount()
	fmt.Println(pri)
	fmt.Println(address)
}

func TestSendTx(t *testing.T) {
	client := newTestClient()
	pri := ""
	from := "9872cefe5ba026ba4806bd0ab9e4ebe77790febd31f09813eb9c798a19634e45"
	to := "2d4e63523b106f42b08dc16b59e0914273d8136af37f4474258a5c90ed08c453"
	hash, _ := client.SignAndSendTx(pri, from, to, "1000000000000000000000")
	fmt.Println(hash)
}

func TestGetBlockListByHash(t *testing.T) {
	client := newTestClient()
	blocks, _ := client.GetBlockListByHash([]string{"2FPDwfy7ib42xtLwFRAzivdLx2D7ce8H4unHNvsYmgPp"})
	fmt.Println(len(blocks))
}
