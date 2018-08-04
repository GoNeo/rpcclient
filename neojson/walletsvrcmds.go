package neojson

// GetNewAddressCmd defines the getnewaddress JSON-RPC command.
type GetNewAddressCmd struct {
}

// NewGetNewAddressCmd returns a new instance which can be used to issue a
// getnewaddress JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewGetNewAddressCmd() *GetNewAddressCmd {
	return &GetNewAddressCmd{}
}

// GetTransactionCmd defines the gettransaction JSON-RPC command.
type GetTransactionCmd struct {
	Txid    string
	Verbose *int `jsonrpcdefault:"1"`
}

// NewGetTransactionCmd returns a new instance which can be used to issue a
// gettransaction JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewGetTransactionCmd(txHash string, verbose *int) *GetTransactionCmd {
	return &GetTransactionCmd{
		Txid:    txHash,
		Verbose: verbose,
	}
}

// GetApplicationLogCmd defines the getapplicationlog JSON-RPC command.
type GetApplicationLogCmd struct {
	Txid    string
	Verbose *int `jsonrpcdefault:"1"`
}

func NewGetApplicationLogCmd(txHash string, verbose *int) *GetApplicationLogCmd {
	return &GetApplicationLogCmd{
		Txid:    txHash,
		Verbose: verbose,
	}
}

type InvokeFunctoinCmd struct {
	ScriptHash string
	MethodName string
	Args       []InvokeFunctionArgs
}

type InvokeFunctionArgs struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func NewInvokeFunctoinCmd(scriptHash string, methodName string, args []InvokeFunctionArgs) *InvokeFunctoinCmd {
	return &InvokeFunctoinCmd{
		ScriptHash: scriptHash,
		MethodName: methodName,
		Args:       args,
	}
}

// SendToAddressCmd defines the sendtoaddress JSON-RPC command.
type SendToAddressCmd struct {
	AssetId string
	Address string
	Amount  int64
}

// NewSendToAddressCmd returns a new instance which can be used to issue a
// sendtoaddress JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewSendToAddressCmd(asset_id, address string, amount int64) *SendToAddressCmd {
	return &SendToAddressCmd{
		AssetId: asset_id,
		Address: address,
		Amount:  amount,
	}
}

type AssetAddressValue struct {
	Asset   string  `json:"asset"`
	Value   float64 `json:"value"`
	Address string  `json:"address"`
}

type SendManyCmd struct {
	AddressAmounts *[]AssetAddressValue
}

func NewSendManyCmd(asset_address_values *[]AssetAddressValue) *SendManyCmd {
	return &SendManyCmd{
		asset_address_values,
	}
}

// GetBalanceCmd defines the getbalance JSON-RPC command.
type GetBalanceCmd struct {
	AssetId string
}

// NewGetBalanceCmd returns a new instance which can be used to issue a
// getbalance JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewGetBalanceCmd(asset_id string) *GetBalanceCmd {
	return &GetBalanceCmd{
		AssetId: asset_id,
	}
}

func init() {
	// The commands in this file are only usable with a wallet server.
	flags := UFWalletOnly

	MustRegisterCmd("getnewaddress", (*GetNewAddressCmd)(nil), flags)
	MustRegisterCmd("getrawtransaction", (*GetTransactionCmd)(nil), flags)
	MustRegisterCmd("getapplicationlog", (*GetApplicationLogCmd)(nil), flags)
	MustRegisterCmd("invokefunction", (*InvokeFunctoinCmd)(nil), flags)
	MustRegisterCmd("sendtoaddress", (*SendToAddressCmd)(nil), flags)
	MustRegisterCmd("getbalance", (*GetBalanceCmd)(nil), flags)
	MustRegisterCmd("sendmany", (*SendManyCmd)(nil), flags)

}
