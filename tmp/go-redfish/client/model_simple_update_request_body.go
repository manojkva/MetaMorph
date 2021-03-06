/*
 * Redfish OAPI specification
 *
 * Partial Redfish OAPI specification for a limited client
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// SimpleUpdateRequestBody struct for SimpleUpdateRequestBody
type SimpleUpdateRequestBody struct {
	ImageURI string `json:"ImageURI"`
	Targets []string `json:"Targets,omitempty"`
	TransferProtocolType TransferProtocolType `json:"TransferProtocolType,omitempty"`
}
