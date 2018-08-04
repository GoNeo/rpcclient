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


type GetApplicationLogResult struct {
	GasConsumed   string `json:"gas_consumed"`
	Notifications []struct {
		Contract string `json:"contract"`
		State    struct {
			Type  string `json:"type"`
			Value []struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"value"`
		} `json:"state"`
	} `json:"notifications"`
	Stack []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"stack"`
	Txid    string `json:"txid"`
	Vmstate string `json:"vmstate"`
}


type InvokeFunctionResult struct {
	GasConsumed string `json:"gas_consumed"`
	Script      string `json:"script"`
	Stack       []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"stack"`
	State string `json:"state"`
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
