package insights

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i29d2c155f681081063d0b2c0ab92796f0e2b78c469f71af45a16d8240b348ebf "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i8893ecb7e9d38ec8961e185f29391303f74779021473d1c37caeafd1abebef9a "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending"
    i99fad623e5b2b4a7ae85326fa692745befba2f6b4fb179d4572c3e7219fbf3ac "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used"
    i7eb746d3d91c620af79236bba2af4b708c6b65d5f4ec8f9cc42e5d4305f1c396 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item"
    ibf98f8f6ecbe2d059f97c4d3970656db9710f9bafa095bf25ac791007a80e997 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item"
    ic47fb69b7ecc13ce7f81973ffdde08efe7b70445311d1e83964064ef79497098 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item"
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
    m.urlTemplate = "{+baseurl}/me/insights{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property insights for me
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
// CreatePatchRequestInformation update the navigation property insights in me
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
// Delete delete navigation property insights for me
func (m *InsightsRequestBuilder) Delete(options *InsightsRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
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
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateOfficeGraphInsightsFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OfficeGraphInsightsable), nil
}
// Patch update the navigation property insights in me
func (m *InsightsRequestBuilder) Patch(options *InsightsRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *InsightsRequestBuilder) Shared()(*i29d2c155f681081063d0b2c0ab92796f0e2b78c469f71af45a16d8240b348ebf.SharedRequestBuilder) {
    return i29d2c155f681081063d0b2c0ab92796f0e2b78c469f71af45a16d8240b348ebf.NewSharedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.insights.shared.item collection
func (m *InsightsRequestBuilder) SharedById(id string)(*i7eb746d3d91c620af79236bba2af4b708c6b65d5f4ec8f9cc42e5d4305f1c396.SharedInsightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedInsight_id"] = id
    }
    return i7eb746d3d91c620af79236bba2af4b708c6b65d5f4ec8f9cc42e5d4305f1c396.NewSharedInsightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InsightsRequestBuilder) Trending()(*i8893ecb7e9d38ec8961e185f29391303f74779021473d1c37caeafd1abebef9a.TrendingRequestBuilder) {
    return i8893ecb7e9d38ec8961e185f29391303f74779021473d1c37caeafd1abebef9a.NewTrendingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TrendingById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.insights.trending.item collection
func (m *InsightsRequestBuilder) TrendingById(id string)(*ibf98f8f6ecbe2d059f97c4d3970656db9710f9bafa095bf25ac791007a80e997.TrendingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["trending_id"] = id
    }
    return ibf98f8f6ecbe2d059f97c4d3970656db9710f9bafa095bf25ac791007a80e997.NewTrendingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InsightsRequestBuilder) Used()(*i99fad623e5b2b4a7ae85326fa692745befba2f6b4fb179d4572c3e7219fbf3ac.UsedRequestBuilder) {
    return i99fad623e5b2b4a7ae85326fa692745befba2f6b4fb179d4572c3e7219fbf3ac.NewUsedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsedById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.insights.used.item collection
func (m *InsightsRequestBuilder) UsedById(id string)(*ic47fb69b7ecc13ce7f81973ffdde08efe7b70445311d1e83964064ef79497098.UsedInsightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usedInsight_id"] = id
    }
    return ic47fb69b7ecc13ce7f81973ffdde08efe7b70445311d1e83964064ef79497098.NewUsedInsightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
