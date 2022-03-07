/*
 * Moralis API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Trade struct {
	// The transaction hash
	TransactionHash string `json:"transaction_hash"`
	// The transaction index
	TransactionIndex string `json:"transaction_index"`
	// The token id(s) traded
	TokenIds *Array `json:"token_ids"`
	// The address that sold the NFT
	SellerAddress string `json:"seller_address"`
	// The address that bought the NFT
	BuyerAddress string `json:"buyer_address"`
	// The address of the contract that traded the NFT
	MarketplaceAddress string `json:"marketplace_address"`
	// The value that was sent in the transaction (ETH/BNB/etc..)
	Price string `json:"price"`
	// The block timestamp
	BlockTimestamp string `json:"block_timestamp"`
	// The blocknumber of the transaction
	BlockNumber string `json:"block_number"`
	// The block hash
	BlockHash string `json:"block_hash"`
}