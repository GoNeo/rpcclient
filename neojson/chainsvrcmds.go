package neojson

// GetBlockCountCmd defines the getblockcount JSON-RPC command.
type GetBlockCountCmd struct{}

// NewGetBlockCountCmd returns a new instance which can be used to issue a
// getblockcount JSON-RPC command.
func NewGetBlockCountCmd() *GetBlockCountCmd {
	return &GetBlockCountCmd{}
}


// GetBlockCmd defines the getblock JSON-RPC command.
type GetBlockCmd struct {
	Hash      string
	Verbose   *int `jsonrpcdefault:"1"`
}

// NewGetBlockCmd returns a new instance which can be used to issue a getblock
// JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewGetBlockCmd(hash string, verbose *int) *GetBlockCmd {
	return &GetBlockCmd{
		Hash:      hash,
		Verbose:   verbose,
	}
}

// GetBlockHashCmd defines the getblockhash JSON-RPC command.
type GetBlockHashCmd struct {
	Index int64
}

// NewGetBlockHashCmd returns a new instance which can be used to issue a
// getblockhash JSON-RPC command.
func NewGetBlockHashCmd(index int64) *GetBlockHashCmd {
	return &GetBlockHashCmd{
		Index: index,
	}
}


// GetBestBlockHashCmd defines the getbestblockhash JSON-RPC command.
type GetBestBlockHashCmd struct{}

// NewGetBestBlockHashCmd returns a new instance which can be used to issue a
// getbestblockhash JSON-RPC command.
func NewGetBestBlockHashCmd() *GetBestBlockHashCmd {
	return &GetBestBlockHashCmd{}
}


func init() {
	// No special flags for commands in this file.
	flags := UsageFlag(0)

	MustRegisterCmd("getblockcount", (*GetBlockCountCmd)(nil), flags)
	MustRegisterCmd("getblockhash", (*GetBlockHashCmd)(nil), flags)
	MustRegisterCmd("getblock", (*GetBlockCmd)(nil), flags)
	MustRegisterCmd("getbestblockhash", (*GetBestBlockHashCmd)(nil), flags)



}
