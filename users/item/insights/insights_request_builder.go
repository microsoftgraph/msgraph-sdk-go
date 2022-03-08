package insights

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0179bb6fa871e13743008821796dbc9552d7f97c10cb54f6706fbdf6be05df6e "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used"
    i3e4ce2336bbbc58648757b85760fb89c5881021a08e9098f363056f0c6a6195c "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared"
    iab3e868bd74f90122481d659f92b93937dc248e664cb771faabf5f029a1a664f "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending"
    ia144b1cd33e78603da2911ebd505e4e1bb01884db983335223d6dcd6008a9a5f "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item"
    ie5cd453259522930b20937b7a3b030629926bd41d23d0488babbe3bc6a06c20c "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item"
    if5106cc51c89d49dab777a3bad4d2f60f1791257ebdb9793f092fde43062bbe4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item"
)

// InsightsRequestBuilder provides operations to manage the insights property of the microsoft.graph.user entity.
type InsightsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// InsightsRequestBuilderDeleteOptions options for Delete
type InsightsRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// InsightsRequestBuilderGetOptions options for Get
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
// InsightsRequestBuilderGetQueryParameters read-only. Nullable.
type InsightsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// InsightsRequestBuilderPatchOptions options for Patch
type InsightsRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OfficeGraphInsightsable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewInsightsRequestBuilderInternal instantiates a new InsightsRequestBuilder and sets the default values.
func NewInsightsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InsightsRequestBuilder) {
    m := &InsightsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInsightsRequestBuilder instantiates a new InsightsRequestBuilder and sets the default values.
func NewInsightsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InsightsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInsightsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property insights for users
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
// CreateGetRequestInformation read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property insights in users
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
// Delete delete navigation property insights for users
func (m *InsightsRequestBuilder) Delete(options *InsightsRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read-only. Nullable.
func (m *InsightsRequestBuilder) Get(options *InsightsRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OfficeGraphInsightsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateOfficeGraphInsightsFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OfficeGraphInsightsable), nil
}
// Patch update the navigation property insights in users
func (m *InsightsRequestBuilder) Patch(options *InsightsRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *InsightsRequestBuilder) Shared()(*i3e4ce2336bbbc58648757b85760fb89c5881021a08e9098f363056f0c6a6195c.SharedRequestBuilder) {
    return i3e4ce2336bbbc58648757b85760fb89c5881021a08e9098f363056f0c6a6195c.NewSharedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.insights.shared.item collection
func (m *InsightsRequestBuilder) SharedById(id string)(*if5106cc51c89d49dab777a3bad4d2f60f1791257ebdb9793f092fde43062bbe4.SharedInsightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedInsight_id"] = id
    }
    return if5106cc51c89d49dab777a3bad4d2f60f1791257ebdb9793f092fde43062bbe4.NewSharedInsightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InsightsRequestBuilder) Trending()(*iab3e868bd74f90122481d659f92b93937dc248e664cb771faabf5f029a1a664f.TrendingRequestBuilder) {
    return iab3e868bd74f90122481d659f92b93937dc248e664cb771faabf5f029a1a664f.NewTrendingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TrendingById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.insights.trending.item collection
func (m *InsightsRequestBuilder) TrendingById(id string)(*ie5cd453259522930b20937b7a3b030629926bd41d23d0488babbe3bc6a06c20c.TrendingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["trending_id"] = id
    }
    return ie5cd453259522930b20937b7a3b030629926bd41d23d0488babbe3bc6a06c20c.NewTrendingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InsightsRequestBuilder) Used()(*i0179bb6fa871e13743008821796dbc9552d7f97c10cb54f6706fbdf6be05df6e.UsedRequestBuilder) {
    return i0179bb6fa871e13743008821796dbc9552d7f97c10cb54f6706fbdf6be05df6e.NewUsedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsedById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.insights.used.item collection
func (m *InsightsRequestBuilder) UsedById(id string)(*ia144b1cd33e78603da2911ebd505e4e1bb01884db983335223d6dcd6008a9a5f.UsedInsightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usedInsight_id"] = id
    }
    return ia144b1cd33e78603da2911ebd505e4e1bb01884db983335223d6dcd6008a9a5f.NewUsedInsightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
