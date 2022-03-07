/*
 * Moralis API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Erc20Price struct {

	NativePrice *NativeErc20Price `json:"nativePrice,omitempty"`
	// The price in USD for the token
	UsdPrice float64 `json:"usdPrice"`
	// The address of the exchange used to calculate the price
	ExchangeAddress string `json:"exchangeAddress,omitempty"`
	// The name of the exchange used for calculating the price
	ExchangeName string `json:"exchangeName,omitempty"`
}
