package main

import (
	"log"
	"github.com/GoNeo/rpcclient"
	"github.com/GoNeo/rpcclient/neojson"
)

func main() {
	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:20332",
		HTTPPostMode: true, //  supports HTTP POST mode
		DisableTLS:   true, //  does not provide TLS by default
	}

	client, err := rpcclient.New(connCfg)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	//addr, err := client.GetNewAddress()
	//log.Printf("address %s %v", addr, err)

	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)

	txID := "0xd4c538554320c461c7c4511d44ae9d20b2bf2d820b6f24b0e749b140e854a5e1"
	tx, err := client.GetRawTransaction(txID)
	log.Println("transaction ", tx, err)

	blockHeight := int64(1261895)
	blockhash, err := client.GetBlockHash(blockHeight)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block hash: %s ,by index %d", blockhash, blockHeight)

	block, err := client.GetBlock(blockhash)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block : %v ,by block hash %d", block, blockhash)

	bestblockhash, err := client.GetBestBlockHash()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("best Block hash: %s ", bestblockhash)

	asset_id := "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b"
	//to := "AVtPp6Yc5iuJU8G6ht65i4nSyzdxG9NBiw"
	//amount := int64(1)
	//sendToAddrResult, err := client.SendToAddress(asset_id, to, amount)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("sendToAddrResult : %v ", sendToAddrResult)

	balance, err := client.GetBalance(asset_id)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("best Block hash: %v ", balance)

	// send many
	asset_address_values := []neojson.AssetAddressValue{}

	asset_address_value1 := neojson.AssetAddressValue{
		Asset:   asset_id,
		Address: "AVtPp6Yc5iuJU8G6ht65i4nSyzdxG9NBiw",
		Value:   2,
	}
	asset_address_values = append(asset_address_values, asset_address_value1)
	asset_address_value2 := neojson.AssetAddressValue{
		Asset:   asset_id,
		Address: "AVtPp6Yc5iuJU8G6ht65i4nSyzdxG9NBiw",
		Value:   3,
	}
	asset_address_values = append(asset_address_values, asset_address_value2)

	sendManyResult, err := client.SendMany(&asset_address_values)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("send Many result: %v ", sendManyResult)

}
