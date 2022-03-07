/*
 * Moralis API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Erc20TokenBalance struct {
	// The address of the token contract
	TokenAddress string `json:"token_address"`
	// The name of the token Contract
	Name string `json:"name"`
	// The symbol of the NFT contract
	Symbol string `json:"symbol"`
	// The logo of the token
	Logo string `json:"logo,omitempty"`
	// The thumbnail of the logo
	Thumbnail string `json:"thumbnail,omitempty"`
	// The number of decimals on of the token
	Decimals string `json:"decimals"`
	// Timestamp of when the contract was last synced with the node
	Balance string `json:"balance"`
}