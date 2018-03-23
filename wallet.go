package rpcclient

import (
	"github.com/GoNeo/rpcclient/neojson"
	"encoding/json"
)

// FutureGetNewAddressResult is a future promise to deliver the result of a
// GetNewAddressAsync RPC invocation (or an applicable error).
type FutureGetNewAddressResult chan *response

// Receive waits for the response promised by the future and returns a new
// address.
func (r FutureGetNewAddressResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	// Unmarshal result as a string.
	var addr string
	err = json.Unmarshal(res, &addr)
	if err != nil {
		return "", err
	}

	return addr, nil
}

// GetNewAddressAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetNewAddress for the blocking version and more details.
func (c *Client) GetNewAddressAsync() FutureGetNewAddressResult {
	cmd := neojson.NewGetNewAddressCmd()
	return c.sendCmd(cmd)
}

// GetNewAddress returns a new address.
func (c *Client) GetNewAddress() (string, error) {
	return c.GetNewAddressAsync().Receive()
}

// FutureGetTransactionResult is a future promise to deliver the result
// of a GetTransactionAsync RPC invocation (or an applicable error).
type FutureGetTransactionResult chan *response

// Receive waits for the response promised by the future and returns detailed
// information about a wallet transaction.
func (r FutureGetTransactionResult) Receive() (*neojson.GetTransactionResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a gettransaction result object
	var getTx neojson.GetTransactionResult
	err = json.Unmarshal(res, &getTx)
	if err != nil {
		return nil, err
	}

	return &getTx, nil
}

// GetTransactionAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See GetTransaction for the blocking version and more details.
func (c *Client) GetTransactionAsync(txID string, verbose *int) FutureGetTransactionResult {
	cmd := neojson.NewGetTransactionCmd(txID, verbose)
	return c.sendCmd(cmd)
}

// GetTransaction returns detailed information about a wallet transaction.
//
// See GetRawTransaction to return the raw transaction instead.
func (c *Client) GetRawTransaction(txID string) (*neojson.GetTransactionResult, error) {
	verbose := 1
	return c.GetTransactionAsync(txID, &verbose).Receive()
}

// FutureSendToAddressResult is a future promise to deliver the result of a
// SendToAddressAsync RPC invocation (or an applicable error).
type FutureSendToAddressResult chan *response

// Receive waits for the response promised by the future and returns the hash
// of the transaction sending the passed amount to the given address.
func (r FutureSendToAddressResult) Receive() (*neojson.SendToAddressResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var sendToaddr neojson.SendToAddressResult
	err = json.Unmarshal(res, &sendToaddr)
	if err != nil {
		return nil, err
	}

	return &sendToaddr, nil
}

// SendToAddressAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See SendToAddress for the blocking version and more details.
func (c *Client) SendToAddressAsync(asset_id, address string, amount int64) FutureSendToAddressResult {
	cmd := neojson.NewSendToAddressCmd(asset_id, address, amount)
	return c.sendCmd(cmd)
}

// SendToAddress sends the passed amount to the given address.
//
// See SendToAddressComment to associate comments with the transaction in the
// wallet.  The comments are not part of the transaction and are only internal
// to the wallet.
//
// NOTE: This function requires to the wallet to be unlocked.  See the
// WalletPassphrase function for more details.
func (c *Client) SendToAddress(asset_id, address string, amount int64) (*neojson.SendToAddressResult, error) {
	return c.SendToAddressAsync(asset_id, address, amount).Receive()
}

type FutureSendManyResult chan *response

func (r FutureSendManyResult) Receive() (*neojson.SendManyResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var sendToaddr neojson.SendManyResult
	err = json.Unmarshal(res, &sendToaddr)
	if err != nil {
		return nil, err
	}

	return &sendToaddr, nil
}

func (c *Client) SendManyAsync(asset_address_values *[]neojson.AssetAddressValue) FutureSendManyResult {
	cmd := neojson.NewSendManyCmd(asset_address_values)
	return c.sendCmd(cmd)
}

func (c *Client) SendMany(asset_address_values *[]neojson.AssetAddressValue) (*neojson.SendManyResult, error) {
	return c.SendManyAsync(asset_address_values).Receive()
}

// FutureGetBalanceResult is a future promise to deliver the result of a
// GetBalanceAsync or GetBalanceMinConfAsync RPC invocation (or an applicable
// error).
type FutureGetBalanceResult chan *response

// Receive waits for the response promised by the future and returns the
// available balance from the server for the specified account.
func (r FutureGetBalanceResult) Receive() (*neojson.GetBalanceResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string
	var balanceResult neojson.GetBalanceResult
	err = json.Unmarshal(res, &balanceResult)
	if err != nil {
		return nil, err
	}

	return &balanceResult, nil
}

// GetBalanceAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetBalance for the blocking version and more details.
func (c *Client) GetBalanceAsync(asset_id string) FutureGetBalanceResult {
	cmd := neojson.NewGetBalanceCmd(asset_id)
	return c.sendCmd(cmd)
}

// GetBalance returns the available balance from the server for the specified
// account using the default number of minimum confirmations.  The account may
// be "*" for all accounts.
//
// See GetBalanceMinConf to override the minimum number of confirmations.
func (c *Client) GetBalance(asset_id string) (*neojson.GetBalanceResult, error) {
	return c.GetBalanceAsync(asset_id).Receive()
}
