package neojson

type GetBlockResult struct {
	Hash              string `json:"hash"`
	Size              int    `json:"size"`
	Version           int    `json:"version"`
	Previousblockhash string `json:"previousblockhash"`
	Merkleroot        string `json:"merkleroot"`
	Time              int64    `json:"time"`
	Index             int64    `json:"index"`
	Nonce             string `json:"nonce"`
	Nextconsensus     string `json:"nextconsensus"`
	Script            struct {
		Invocation   string `json:"invocation"`
		Verification string `json:"verification"`
	} `json:"script"`
	Tx []struct {
		Txid       string        `json:"txid"`
		Size       int           `json:"size"`
		Type       string        `json:"type"`
		Version    int           `json:"version"`
		Attributes []interface{} `json:"attributes"`
		Vin        []struct {
			Txid string `json:"txid"`
			Vout int    `json:"vout"`
		} `json:"vin"`
		Vout []struct {
			N       int    `json:"n"`
			Asset   string `json:"asset"`
			Value   string `json:"value"`
			Address string `json:"address"`
		} `json:"vout"`
		SysFee  string `json:"sys_fee"`
		NetFee  string `json:"net_fee"`
		Scripts []struct {
			Invocation   string `json:"invocation"`
			Verification string `json:"verification"`
		} `json:"scripts"`
	} `json:"tx"`
	Confirmations int64    `json:"confirmations"`
	Nextblockhash string `json:"nextblockhash"`
}