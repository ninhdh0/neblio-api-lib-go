/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Object containing an array of transaction objects
type GetTxsResponse struct {

	// Number of pages of transactions
	PagesTotal float32 `json:"pagesTotal,omitempty"`

	// Array of transaction objects
	Txs []GetTxResponse `json:"txs,omitempty"`
}
