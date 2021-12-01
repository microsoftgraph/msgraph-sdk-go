package createdobjects

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i6f42bc516a096b55e74a963f59b1f3896b161e60c5cf202a3b22cd998284d1e5 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/item/createdobjects/ref"
)

// CreatedObjectsRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\createdObjects
type CreatedObjectsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CreatedObjectsRequestBuilderGetOptions options for Get
type CreatedObjectsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CreatedObjectsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CreatedObjectsRequestBuilderGetQueryParameters directory objects created by this service principal. Read-only. Nullable.
type CreatedObjectsRequestBuilderGetQueryParameters struct {
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
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// NewCreatedObjectsRequestBuilderInternal instantiates a new CreatedObjectsRequestBuilder and sets the default values.
func NewCreatedObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreatedObjectsRequestBuilder) {
    m := &CreatedObjectsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal_id}/createdObjects{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCreatedObjectsRequestBuilder instantiates a new CreatedObjectsRequestBuilder and sets the default values.
func NewCreatedObjectsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreatedObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCreatedObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation directory objects created by this service principal. Read-only. Nullable.
func (m *CreatedObjectsRequestBuilder) CreateGetRequestInformation(options *CreatedObjectsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get directory objects created by this service principal. Read-only. Nullable.
func (m *CreatedObjectsRequestBuilder) Get(options *CreatedObjectsRequestBuilderGetOptions)(*CreatedObjectsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCreatedObjectsResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*CreatedObjectsResponse), nil
}
func (m *CreatedObjectsRequestBuilder) Ref()(*i6f42bc516a096b55e74a963f59b1f3896b161e60c5cf202a3b22cd998284d1e5.RefRequestBuilder) {
    return i6f42bc516a096b55e74a963f59b1f3896b161e60c5cf202a3b22cd998284d1e5.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
