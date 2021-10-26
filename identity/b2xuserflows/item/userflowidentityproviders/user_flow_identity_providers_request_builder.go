package userflowidentityproviders

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i7458e2a8c1750e8336a16e93580de785722f2083e4ede36c629a2a3d477240aa "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/userflowidentityproviders/availableprovidertypes"
    id1418e897ecd09b3d256be063da8c4b869278ce8ee0614463dc6dedaeca5afc0 "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/userflowidentityproviders/ref"
)

// Builds and executes requests for operations under \identity\b2xUserFlows\{b2xIdentityUserFlow-id}\userFlowIdentityProviders
type UserFlowIdentityProvidersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Get userFlowIdentityProviders from identity
type UserFlowIdentityProvidersRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// Builds and executes requests for operations under \identity\b2xUserFlows\{b2xIdentityUserFlow-id}\userFlowIdentityProviders\microsoft.graph.availableProviderTypes()
func (m *UserFlowIdentityProvidersRequestBuilder) AvailableProviderTypes()(*i7458e2a8c1750e8336a16e93580de785722f2083e4ede36c629a2a3d477240aa.AvailableProviderTypesRequestBuilder) {
    return i7458e2a8c1750e8336a16e93580de785722f2083e4ede36c629a2a3d477240aa.NewAvailableProviderTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new UserFlowIdentityProvidersRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewUserFlowIdentityProvidersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserFlowIdentityProvidersRequestBuilder) {
    m := &UserFlowIdentityProvidersRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/identity/b2xUserFlows/{b2xIdentityUserFlow_id}/userFlowIdentityProviders{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new UserFlowIdentityProvidersRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewUserFlowIdentityProvidersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserFlowIdentityProvidersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserFlowIdentityProvidersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get userFlowIdentityProviders from identity
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *UserFlowIdentityProvidersRequestBuilder) CreateGetRequestInformation(q func (value *UserFlowIdentityProvidersRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(UserFlowIdentityProvidersRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get userFlowIdentityProviders from identity
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *UserFlowIdentityProvidersRequestBuilder) Get(q func (value *UserFlowIdentityProvidersRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*UserFlowIdentityProvidersResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserFlowIdentityProvidersResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*UserFlowIdentityProvidersResponse), nil
}
func (m *UserFlowIdentityProvidersRequestBuilder) Ref()(*id1418e897ecd09b3d256be063da8c4b869278ce8ee0614463dc6dedaeca5afc0.RefRequestBuilder) {
    return id1418e897ecd09b3d256be063da8c4b869278ce8ee0614463dc6dedaeca5afc0.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
