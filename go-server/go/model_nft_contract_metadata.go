/*
 * Moralis API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type NftContractMetadata struct {
	// The address of the token contract
	TokenAddress string `json:"token_address"`
	// The name of the token Contract
	Name string `json:"name"`
	// Timestamp of when the contract was last synced with the node
	SyncedAt string `json:"synced_at,omitempty"`
	// The symbol of the NFT contract
	Symbol string `json:"symbol"`
	// The type of NFT contract
	ContractType string `json:"contract_type"`
}
