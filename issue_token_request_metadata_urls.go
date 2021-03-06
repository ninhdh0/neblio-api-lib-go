/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type IssueTokenRequestMetadataUrls struct {

	// Name of the URL
	Name string `json:"name,omitempty"`

	// Actual URL
	Url string `json:"url,omitempty"`

	// mimeType of URL content
	MimeType string `json:"mimeType,omitempty"`

	// Hash of data at the URL, used for verification
	DataHash string `json:"dataHash,omitempty"`
}
