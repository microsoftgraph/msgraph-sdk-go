package children

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ie432fb01138850420e612dd342f68a9eaed12f35366219dc96926150776878b1 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/children/ref"
)

// ChildrenRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\children
type ChildrenRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ChildrenRequestBuilderGetOptions options for Get
type ChildrenRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ChildrenRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChildrenRequestBuilderGetQueryParameters collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
type ChildrenRequestBuilderGetQueryParameters struct {
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
// NewChildrenRequestBuilderInternal instantiates a new ChildrenRequestBuilder and sets the default values.
func NewChildrenRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChildrenRequestBuilder) {
    m := &ChildrenRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/children{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChildrenRequestBuilder instantiates a new ChildrenRequestBuilder and sets the default values.
func NewChildrenRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChildrenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChildrenRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
func (m *ChildrenRequestBuilder) CreateGetRequestInformation(options *ChildrenRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
func (m *ChildrenRequestBuilder) Get(options *ChildrenRequestBuilderGetOptions)(*ChildrenResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChildrenResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ChildrenResponse), nil
}
func (m *ChildrenRequestBuilder) Ref()(*ie432fb01138850420e612dd342f68a9eaed12f35366219dc96926150776878b1.RefRequestBuilder) {
    return ie432fb01138850420e612dd342f68a9eaed12f35366219dc96926150776878b1.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
