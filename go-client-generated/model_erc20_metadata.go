/*
 * Moralis API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Erc20Metadata struct {
	// The address of the token contract
	Address string `json:"address"`
	// The name of the token Contract
	Name string `json:"name"`
	// The symbol of the NFT contract
	Symbol string `json:"symbol"`
	// The number of decimals on of the token
	Decimals string `json:"decimals"`
	// The logo of the token
	Logo string `json:"logo,omitempty"`
	// The logo hash
	LogoHash string `json:"logo_hash,omitempty"`
	// The thumbnail of the logo
	Thumbnail string `json:"thumbnail,omitempty"`
	BlockNumber string `json:"block_number,omitempty"`
	Validated string `json:"validated,omitempty"`
}
