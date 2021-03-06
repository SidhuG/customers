/*
 * Customers API
 *
 * Customers focuses on solving authentic identification of humans who are legally able to hold and transfer currency within the US. Primarily this project solves [Know Your Customer](https://en.wikipedia.org/wiki/Know_your_customer) (KYC), [Customer Identification Program](https://en.wikipedia.org/wiki/Customer_Identification_Program) (CIP), [Office of Foreign Asset Control](https://www.treasury.gov/about/organizational-structure/offices/Pages/Office-of-Foreign-Assets-Control.aspx) (OFAC) checks and verification workflows to comply with United States federal law and ensure authentic transfers. Customers has an objective to be a service for detailed due diligence on individuals and companies for Financial Institutions and services in a modernized and extensible way.  Customer phone numbers and addresses are stored and partially used in KYC/OFAC validation. Arbitrary key/value pairs can be stored for a Customer. Documents and Disclaimers, and their acknowledgment are also stored under a Customer as they're accepted. Bank Accounts, which can be validated with micro-deposits currently, are stored under each Customer.  ![](https://raw.githubusercontent.com/adamdecaf/customers/create-accounts/docs/images/customer.png)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	_context "context"
	"github.com/antihax/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// AccountValidationApiService AccountValidationApi service
type AccountValidationApiService service

// CompleteAccountValidationOpts Optional parameters for the method 'CompleteAccountValidation'
type CompleteAccountValidationOpts struct {
	XRequestID    optional.String
	XOrganization optional.String
}

/*
CompleteAccountValidation Complete Account Validation
Complete account validation with specified strategy and vendor.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customerID customerID of the Customer the accountID belongs to
 * @param accountID accountID of the Account to validate
 * @param completeAccountValidationRequest
 * @param optional nil or *CompleteAccountValidationOpts - Optional Parameters:
 * @param "XRequestID" (optional.String) -  Optional requestID allows application developer to trace requests through the systems logs
 * @param "XOrganization" (optional.String) -  Value used to separate and identify models
@return CompleteAccountValidationResponse
*/
func (a *AccountValidationApiService) CompleteAccountValidation(ctx _context.Context, customerID string, accountID string, completeAccountValidationRequest CompleteAccountValidationRequest, localVarOptionals *CompleteAccountValidationOpts) (CompleteAccountValidationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CompleteAccountValidationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/customers/{customerID}/accounts/{accountID}/validations"
	localVarPath = strings.Replace(localVarPath, "{"+"customerID"+"}", _neturl.QueryEscape(parameterToString(customerID, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", _neturl.QueryEscape(parameterToString(accountID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XOrganization.IsSet() {
		localVarHeaderParams["X-Organization"] = parameterToString(localVarOptionals.XOrganization.Value(), "")
	}
	// body params
	localVarPostBody = &completeAccountValidationRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetAccountValidationOpts Optional parameters for the method 'GetAccountValidation'
type GetAccountValidationOpts struct {
	XRequestID    optional.String
	XOrganization optional.String
}

/*
GetAccountValidation Get Account Validation
Get information about account validation strategy, status, etc.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customerID customerID of the Customer the accountID belongs to
 * @param accountID accountID of the Account to validate
 * @param validationID ID of the Validation
 * @param optional nil or *GetAccountValidationOpts - Optional Parameters:
 * @param "XRequestID" (optional.String) -  Optional requestID allows application developer to trace requests through the systems logs
 * @param "XOrganization" (optional.String) -  Value used to separate and identify models
@return AccountValidationResponse
*/
func (a *AccountValidationApiService) GetAccountValidation(ctx _context.Context, customerID string, accountID string, validationID string, localVarOptionals *GetAccountValidationOpts) (AccountValidationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AccountValidationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/customers/{customerID}/accounts/{accountID}/validations/{validationID}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerID"+"}", _neturl.QueryEscape(parameterToString(customerID, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", _neturl.QueryEscape(parameterToString(accountID, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"validationID"+"}", _neturl.QueryEscape(parameterToString(validationID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XOrganization.IsSet() {
		localVarHeaderParams["X-Organization"] = parameterToString(localVarOptionals.XOrganization.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// InitAccountValidationOpts Optional parameters for the method 'InitAccountValidation'
type InitAccountValidationOpts struct {
	XRequestID    optional.String
	XOrganization optional.String
}

/*
InitAccountValidation Initiate Account Validation
Initiate account validation with specified strategy and vendor.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customerID customerID of the Customer the accountID belongs to
 * @param accountID accountID of the Account to validate
 * @param initAccountValidationRequest
 * @param optional nil or *InitAccountValidationOpts - Optional Parameters:
 * @param "XRequestID" (optional.String) -  Optional requestID allows application developer to trace requests through the systems logs
 * @param "XOrganization" (optional.String) -  Value used to separate and identify models
@return InitAccountValidationResponse
*/
func (a *AccountValidationApiService) InitAccountValidation(ctx _context.Context, customerID string, accountID string, initAccountValidationRequest InitAccountValidationRequest, localVarOptionals *InitAccountValidationOpts) (InitAccountValidationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InitAccountValidationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/customers/{customerID}/accounts/{accountID}/validations"
	localVarPath = strings.Replace(localVarPath, "{"+"customerID"+"}", _neturl.QueryEscape(parameterToString(customerID, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", _neturl.QueryEscape(parameterToString(accountID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XOrganization.IsSet() {
		localVarHeaderParams["X-Organization"] = parameterToString(localVarOptionals.XOrganization.Value(), "")
	}
	// body params
	localVarPostBody = &initAccountValidationRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
