package externalcolumns

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i7561827c7bf6a1d5e7586808487675d9f7bc93ac11d07b8d8cdba84b59ddb8c5 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/externalcolumns/ref"
)

// ExternalColumnsRequestBuilder builds and executes requests for operations under \sites\{site-id}\externalColumns
type ExternalColumnsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ExternalColumnsRequestBuilderGetOptions options for Get
type ExternalColumnsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ExternalColumnsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ExternalColumnsRequestBuilderGetQueryParameters the collection of column definitions available in the site that are referenced from the sites in the parent hierarchy of the current site.
type ExternalColumnsRequestBuilderGetQueryParameters struct {
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
// NewExternalColumnsRequestBuilderInternal instantiates a new ExternalColumnsRequestBuilder and sets the default values.
func NewExternalColumnsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExternalColumnsRequestBuilder) {
    m := &ExternalColumnsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/externalColumns{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExternalColumnsRequestBuilder instantiates a new ExternalColumnsRequestBuilder and sets the default values.
func NewExternalColumnsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExternalColumnsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExternalColumnsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the collection of column definitions available in the site that are referenced from the sites in the parent hierarchy of the current site.
func (m *ExternalColumnsRequestBuilder) CreateGetRequestInformation(options *ExternalColumnsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the collection of column definitions available in the site that are referenced from the sites in the parent hierarchy of the current site.
func (m *ExternalColumnsRequestBuilder) Get(options *ExternalColumnsRequestBuilderGetOptions)(*ExternalColumnsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExternalColumnsResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ExternalColumnsResponse), nil
}
func (m *ExternalColumnsRequestBuilder) Ref()(*i7561827c7bf6a1d5e7586808487675d9f7bc93ac11d07b8d8cdba84b59ddb8c5.RefRequestBuilder) {
    return i7561827c7bf6a1d5e7586808487675d9f7bc93ac11d07b8d8cdba84b59ddb8c5.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
