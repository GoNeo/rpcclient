package neojson

// GetTransactionResult models the data from the gettransaction command.
type GetTransactionResult struct {
	Txid       string        `json:"txid"`
	Size       int           `json:"size"`
	Type       string        `json:"type"`
	Version    int           `json:"version"`
	Attributes []interface{} `json:"attributes"`
	Vin []struct {
		Txid string `json:"txid"`
		Vout int    `json:"vout"`
	} `json:"vin"`
	Vout []struct {
		N       int    `json:"n"`
		Asset   string `json:"asset"`
		Value   string `json:"value"`
		Address string `json:"address"`
	} `json:"vout"`
	SysFee string `json:"sys_fee"`
	NetFee string `json:"net_fee"`
	Scripts []struct {
		Invocation   string `json:"invocation"`
		Verification string `json:"verification"`
	} `json:"scripts"`
	Blockhash     string `json:"blockhash"`
	Confirmations int64    `json:"confirmations"`
	Blocktime     int64    `json:"blocktime"`
}

type SendToAddressResult struct {
	Txid       string        `json:"txid"`
	Size       int           `json:"size"`
	Type       string        `json:"type"`
	Version    int           `json:"version"`
	Attributes []interface{} `json:"attributes"`
	Vin []struct {
		Txid string `json:"txid"`
		Vout int    `json:"vout"`
	} `json:"vin"`
	Vout []struct {
		N       int    `json:"n"`
		Asset   string `json:"asset"`
		Value   string `json:"value"`
		Address string `json:"address"`
	} `json:"vout"`
	SysFee string `json:"sys_fee"`
	NetFee string `json:"net_fee"`
	Scripts []struct {
		Invocation   string `json:"invocation"`
		Verification string `json:"verification"`
	} `json:"scripts"`
}

type GetBalanceResult struct {
	Balance   string `json:"balance"`
	Confirmed string `json:"confirmed"`
}

type SendManyResult struct {
	Txid       string        `json:"txid"`
	Size       int           `json:"size"`
	Type       string        `json:"type"`
	Version    int           `json:"version"`
	Attributes []interface{} `json:"attributes"`
	Vin []struct {
		Txid string `json:"txid"`
		Vout int    `json:"vout"`
	} `json:"vin"`
	Vout []struct {
		N       int    `json:"n"`
		Asset   string `json:"asset"`
		Value   string `json:"value"`
		Address string `json:"address"`
	} `json:"vout"`
	SysFee string `json:"sys_fee"`
	NetFee string `json:"net_fee"`
	Scripts []struct {
		Invocation   string `json:"invocation"`
		Verification string `json:"verification"`
	} `json:"scripts"`
}
