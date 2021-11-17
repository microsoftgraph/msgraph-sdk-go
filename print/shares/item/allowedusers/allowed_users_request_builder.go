package allowedusers

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i636de18a266b417f396e31d0dcf1e4274ac84a186cd1e420aa1b71c3c3f188c6 "github.com/microsoftgraph/msgraph-sdk-go/print/shares/item/allowedusers/ref"
)

// Builds and executes requests for operations under \print\shares\{printerShare-id}\allowedUsers
type AllowedUsersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type AllowedUsersRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AllowedUsersRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The users who have access to print using the printer.
type AllowedUsersRequestBuilderGetQueryParameters struct {
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
// Instantiates a new AllowedUsersRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAllowedUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AllowedUsersRequestBuilder) {
    m := &AllowedUsersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/shares/{printerShare_id}/allowedUsers{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new AllowedUsersRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAllowedUsersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AllowedUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAllowedUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// The users who have access to print using the printer.
// Parameters:
//  - options : Options for the request
func (m *AllowedUsersRequestBuilder) CreateGetRequestInformation(options *AllowedUsersRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The users who have access to print using the printer.
// Parameters:
//  - options : Options for the request
func (m *AllowedUsersRequestBuilder) Get(options *AllowedUsersRequestBuilderGetOptions)(*AllowedUsersResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAllowedUsersResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*AllowedUsersResponse), nil
}
func (m *AllowedUsersRequestBuilder) Ref()(*i636de18a266b417f396e31d0dcf1e4274ac84a186cd1e420aa1b71c3c3f188c6.RefRequestBuilder) {
    return i636de18a266b417f396e31d0dcf1e4274ac84a186cd1e420aa1b71c3c3f188c6.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
