// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListFastConnectProviderServicesRequest wrapper for the ListFastConnectProviderServices operation
type ListFastConnectProviderServicesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`
}

func (request ListFastConnectProviderServicesRequest) String() string {
	return common.PointerString(request)
}

// ListFastConnectProviderServicesResponse wrapper for the ListFastConnectProviderServices operation
type ListFastConnectProviderServicesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []FastConnectProviderService instance
	Items []FastConnectProviderService `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListFastConnectProviderServicesResponse) String() string {
	return common.PointerString(response)
}
