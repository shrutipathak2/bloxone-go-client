/*
BloxOne FW API

BloxOne Threat Defense Cloud is an extension of the BloxOne Suite that provides visibility into infected and compromised off-premises devices, roaming users, remote sites, and branch offices. You can subscribe to BloxOne Cloud and use its functionality to mitigate and control malware as well as provide unprecedented insight into your network security posture and enable timely action. BloxOne Cloud also offers unified policy management, reporting, and threat analytics across the entire spectrum. Using automated and high-quality threat intelligence feeds and unique behavioral analytics, it automatically stops device communications with C&Cs/botnets and prevents DNS based data exfiltration.  The mission-critical DNS infrastructure can become a vulnerable component in your network when it is inadequately protected by traditional security solutions and consequently used as an attack surface. Compromised DNS services can result in catastrophic network and system failures. To fully protect your network in today’s cyber security threat environment, Infoblox sets a new DNS security standard by offering scalable, enterprise-grade, and integrated protection for your DNS infrastructure.  Through the Infoblox Cloud Services Portal, you can view the status of your subscription and threat intelligence feeds, manage your network scope and roaming end users, and learn more about threats on your networks through the Infoblox Threat Lookup tool and predefined reports.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fw

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type PopRegionsAPI interface {
	/*
		PopRegionsListPoPRegions List PoP Regions.

		Use this method to retrieve information on all Point of Presence (PoP) regions availablein the database.



		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPopRegionsListPoPRegionsRequest
	*/
	PopRegionsListPoPRegions(ctx context.Context) ApiPopRegionsListPoPRegionsRequest

	// PopRegionsListPoPRegionsExecute executes the request
	//  @return AtcfwListPoPRegionsResponse
	PopRegionsListPoPRegionsExecute(r ApiPopRegionsListPoPRegionsRequest) (*AtcfwListPoPRegionsResponse, *http.Response, error)
	/*
		PopRegionsReadPoPRegion Read PoP Region.

		Use this method to retrieve information on the specified PoP region object.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id The PoP region object identifier
		@return ApiPopRegionsReadPoPRegionRequest
	*/
	PopRegionsReadPoPRegion(ctx context.Context, id int32) ApiPopRegionsReadPoPRegionRequest

	// PopRegionsReadPoPRegionExecute executes the request
	//  @return AtcfwReadPoPRegionResponse
	PopRegionsReadPoPRegionExecute(r ApiPopRegionsReadPoPRegionRequest) (*AtcfwReadPoPRegionResponse, *http.Response, error)
}

// PopRegionsAPIService PopRegionsAPI service
type PopRegionsAPIService internal.Service

type ApiPopRegionsListPoPRegionsRequest struct {
	ctx        context.Context
	ApiService PopRegionsAPI
	filter     *string
	fields     *string
	offset     *int32
	limit      *int32
	pageToken  *string
	tfilter    *string
	torderBy   *string
}

// A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name               | type   | Supported Ops    | | ------------------ | ------ | ---------------- | | region             | string | &#x3D;&#x3D;, !&#x3D;           | | location           | string | ~, !~            |  Grouping operators (and, or, not, ()) are not supported between different fields.
func (r ApiPopRegionsListPoPRegionsRequest) Filter(filter string) ApiPopRegionsListPoPRegionsRequest {
	r.filter = &filter
	return r
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiPopRegionsListPoPRegionsRequest) Fields(fields string) ApiPopRegionsListPoPRegionsRequest {
	r.fields = &fields
	return r
}

// The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.
func (r ApiPopRegionsListPoPRegionsRequest) Offset(offset int32) ApiPopRegionsListPoPRegionsRequest {
	r.offset = &offset
	return r
}

// The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.
func (r ApiPopRegionsListPoPRegionsRequest) Limit(limit int32) ApiPopRegionsListPoPRegionsRequest {
	r.limit = &limit
	return r
}

// The service-defined string used to identify a page of resources. A null value indicates the first page.
func (r ApiPopRegionsListPoPRegionsRequest) PageToken(pageToken string) ApiPopRegionsListPoPRegionsRequest {
	r.pageToken = &pageToken
	return r
}

// Filtering by tags.
func (r ApiPopRegionsListPoPRegionsRequest) Tfilter(tfilter string) ApiPopRegionsListPoPRegionsRequest {
	r.tfilter = &tfilter
	return r
}

// Sorting by tags.
func (r ApiPopRegionsListPoPRegionsRequest) TorderBy(torderBy string) ApiPopRegionsListPoPRegionsRequest {
	r.torderBy = &torderBy
	return r
}

func (r ApiPopRegionsListPoPRegionsRequest) Execute() (*AtcfwListPoPRegionsResponse, *http.Response, error) {
	return r.ApiService.PopRegionsListPoPRegionsExecute(r)
}

/*
PopRegionsListPoPRegions List PoP Regions.

Use this method to retrieve information on all Point of Presence (PoP) regions availablein the database.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPopRegionsListPoPRegionsRequest
*/
func (a *PopRegionsAPIService) PopRegionsListPoPRegions(ctx context.Context) ApiPopRegionsListPoPRegionsRequest {
	return ApiPopRegionsListPoPRegionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AtcfwListPoPRegionsResponse
func (a *PopRegionsAPIService) PopRegionsListPoPRegionsExecute(r ApiPopRegionsListPoPRegionsRequest) (*AtcfwListPoPRegionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *AtcfwListPoPRegionsResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "PopRegionsAPIService.PopRegionsListPoPRegions")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/pop_regions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_filter", r.filter, "")
	}
	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.offset != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_offset", r.offset, "")
	}
	if r.limit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_limit", r.limit, "")
	}
	if r.pageToken != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_page_token", r.pageToken, "")
	}
	if r.tfilter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_tfilter", r.tfilter, "")
	}
	if r.torderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_torder_by", r.torderBy, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		if localVarHTTPResponse.StatusCode == 500 {
			var v AccessCodesListAccessCodes500Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPopRegionsReadPoPRegionRequest struct {
	ctx        context.Context
	ApiService PopRegionsAPI
	id         int32
}

func (r ApiPopRegionsReadPoPRegionRequest) Execute() (*AtcfwReadPoPRegionResponse, *http.Response, error) {
	return r.ApiService.PopRegionsReadPoPRegionExecute(r)
}

/*
PopRegionsReadPoPRegion Read PoP Region.

Use this method to retrieve information on the specified PoP region object.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The PoP region object identifier
	@return ApiPopRegionsReadPoPRegionRequest
*/
func (a *PopRegionsAPIService) PopRegionsReadPoPRegion(ctx context.Context, id int32) ApiPopRegionsReadPoPRegionRequest {
	return ApiPopRegionsReadPoPRegionRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AtcfwReadPoPRegionResponse
func (a *PopRegionsAPIService) PopRegionsReadPoPRegionExecute(r ApiPopRegionsReadPoPRegionRequest) (*AtcfwReadPoPRegionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *AtcfwReadPoPRegionResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "PopRegionsAPIService.PopRegionsReadPoPRegion")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/pop_regions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		if localVarHTTPResponse.StatusCode == 404 {
			var v PopRegionsReadPoPRegion404Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v AccessCodesListAccessCodes500Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}