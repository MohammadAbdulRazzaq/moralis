/*
 * Moralis API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type LogEvent struct {
	// The transaction hash
	TransactionHash string `json:"transaction_hash"`
	// The address of the contract
	Address string `json:"address"`
	// The block timestamp
	BlockTimestamp string `json:"block_timestamp"`
	// The block number
	BlockNumber string `json:"block_number"`
	// The block hash
	BlockHash string `json:"block_hash"`
	// The content of the event
	Data *interface{} `json:"data"`
}