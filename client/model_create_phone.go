/*
 * Customers API
 *
 * Customers ...
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CreatePhone struct {
	// phone number
	Number string `json:"number"`
	Type   string `json:"type"`
}