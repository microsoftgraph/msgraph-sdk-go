package insights

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0179bb6fa871e13743008821796dbc9552d7f97c10cb54f6706fbdf6be05df6e "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used"
    i3e4ce2336bbbc58648757b85760fb89c5881021a08e9098f363056f0c6a6195c "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared"
    iab3e868bd74f90122481d659f92b93937dc248e664cb771faabf5f029a1a664f "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending"
    ia144b1cd33e78603da2911ebd505e4e1bb01884db983335223d6dcd6008a9a5f "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item"
    if5106cc51c89d49dab777a3bad4d2f60f1791257ebdb9793f092fde43062bbe4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item"
)

// Builds and executes requests for operations under \users\{user-id}\insights
type InsightsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type InsightsRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type InsightsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *InsightsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Read-only. Nullable.
type InsightsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type InsightsRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OfficeGraphInsights;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new InsightsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewInsightsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InsightsRequestBuilder) {
    m := &InsightsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new InsightsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewInsightsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InsightsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInsightsRequestBuilderInternal(urlParams, requestAdapter)
}
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *InsightsRequestBuilder) CreateDeleteRequestInformation(options *InsightsRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *InsightsRequestBuilder) CreateGetRequestInformation(options *InsightsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *InsightsRequestBuilder) CreatePatchRequestInformation(options *InsightsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
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
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *InsightsRequestBuilder) Delete(options *InsightsRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *InsightsRequestBuilder) Get(options *InsightsRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OfficeGraphInsights, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOfficeGraphInsights() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OfficeGraphInsights), nil
}
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *InsightsRequestBuilder) Patch(options *InsightsRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *InsightsRequestBuilder) Shared()(*i3e4ce2336bbbc58648757b85760fb89c5881021a08e9098f363056f0c6a6195c.SharedRequestBuilder) {
    return i3e4ce2336bbbc58648757b85760fb89c5881021a08e9098f363056f0c6a6195c.NewSharedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.insights.shared.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *InsightsRequestBuilder) SharedById(id string)(*if5106cc51c89d49dab777a3bad4d2f60f1791257ebdb9793f092fde43062bbe4.SharedInsightRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedInsight_id"] = id
    }
    return if5106cc51c89d49dab777a3bad4d2f60f1791257ebdb9793f092fde43062bbe4.NewSharedInsightRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InsightsRequestBuilder) Trending()(*iab3e868bd74f90122481d659f92b93937dc248e664cb771faabf5f029a1a664f.TrendingRequestBuilder) {
    return iab3e868bd74f90122481d659f92b93937dc248e664cb771faabf5f029a1a664f.NewTrendingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.insights.trending.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *InsightsRequestBuilder) TrendingById(id string)(*iab3e868bd74f90122481d659f92b93937dc248e664cb771faabf5f029a1a664f.TrendingRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["trending_id"] = id
    }
    return iab3e868bd74f90122481d659f92b93937dc248e664cb771faabf5f029a1a664f.NewTrendingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InsightsRequestBuilder) Used()(*i0179bb6fa871e13743008821796dbc9552d7f97c10cb54f6706fbdf6be05df6e.UsedRequestBuilder) {
    return i0179bb6fa871e13743008821796dbc9552d7f97c10cb54f6706fbdf6be05df6e.NewUsedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.insights.used.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *InsightsRequestBuilder) UsedById(id string)(*ia144b1cd33e78603da2911ebd505e4e1bb01884db983335223d6dcd6008a9a5f.UsedInsightRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usedInsight_id"] = id
    }
    return ia144b1cd33e78603da2911ebd505e4e1bb01884db983335223d6dcd6008a9a5f.NewUsedInsightRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
