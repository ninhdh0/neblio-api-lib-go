/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GetTransactionInfoResponseVin struct {

	// TXID of the input
	Txid string `json:"txid,omitempty"`

	// output index
	Vout float32 `json:"vout,omitempty"`

	ScriptSig *GetTransactionInfoResponseScriptSig `json:"scriptSig,omitempty"`

	Sequence float32 `json:"sequence,omitempty"`

	PreviousOutput *GetTransactionInfoResponsePreviousOutput `json:"previousOutput,omitempty"`

	Tokens []GetTransactionInfoResponseTokens `json:"tokens,omitempty"`

	// Value of input in NEBL satoshi
	Value float32 `json:"value,omitempty"`
}
