package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CommunicationsRequestBuilder provides operations to manage the cloudCommunications singleton.
type CommunicationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CommunicationsRequestBuilderGetQueryParameters get communications
type CommunicationsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CommunicationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommunicationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CommunicationsRequestBuilderGetQueryParameters
}
// CommunicationsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommunicationsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CallRecords provides operations to manage the callRecords property of the microsoft.graph.cloudCommunications entity.
func (m *CommunicationsRequestBuilder) CallRecords()(*CommunicationsCallRecordsRequestBuilder) {
    return NewCommunicationsCallRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CallRecordsById provides operations to manage the callRecords property of the microsoft.graph.cloudCommunications entity.
func (m *CommunicationsRequestBuilder) CallRecordsById(id string)(*CommunicationsCallRecordsCallRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["callRecord%2Did"] = id
    }
    return NewCommunicationsCallRecordsCallRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calls provides operations to manage the calls property of the microsoft.graph.cloudCommunications entity.
func (m *CommunicationsRequestBuilder) Calls()(*CommunicationsCallsRequestBuilder) {
    return NewCommunicationsCallsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CallsById provides operations to manage the calls property of the microsoft.graph.cloudCommunications entity.
func (m *CommunicationsRequestBuilder) CallsById(id string)(*CommunicationsCallsCallItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["call%2Did"] = id
    }
    return NewCommunicationsCallsCallItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCommunicationsRequestBuilderInternal instantiates a new CommunicationsRequestBuilder and sets the default values.
func NewCommunicationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommunicationsRequestBuilder) {
    m := &CommunicationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCommunicationsRequestBuilder instantiates a new CommunicationsRequestBuilder and sets the default values.
func NewCommunicationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommunicationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCommunicationsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get communications
func (m *CommunicationsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *CommunicationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update communications
func (m *CommunicationsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudCommunicationsable, requestConfiguration *CommunicationsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get communications
func (m *CommunicationsRequestBuilder) Get(ctx context.Context, requestConfiguration *CommunicationsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudCommunicationsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCloudCommunicationsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudCommunicationsable), nil
}
// GetPresencesByUserId provides operations to call the getPresencesByUserId method.
func (m *CommunicationsRequestBuilder) GetPresencesByUserId()(*CommunicationsGetPresencesByUserIdRequestBuilder) {
    return NewCommunicationsGetPresencesByUserIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetings provides operations to manage the onlineMeetings property of the microsoft.graph.cloudCommunications entity.
func (m *CommunicationsRequestBuilder) OnlineMeetings()(*CommunicationsOnlineMeetingsRequestBuilder) {
    return NewCommunicationsOnlineMeetingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetingsById provides operations to manage the onlineMeetings property of the microsoft.graph.cloudCommunications entity.
func (m *CommunicationsRequestBuilder) OnlineMeetingsById(id string)(*CommunicationsOnlineMeetingsOnlineMeetingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onlineMeeting%2Did"] = id
    }
    return NewCommunicationsOnlineMeetingsOnlineMeetingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update communications
func (m *CommunicationsRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudCommunicationsable, requestConfiguration *CommunicationsRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudCommunicationsable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCloudCommunicationsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudCommunicationsable), nil
}
// Presences provides operations to manage the presences property of the microsoft.graph.cloudCommunications entity.
func (m *CommunicationsRequestBuilder) Presences()(*CommunicationsPresencesRequestBuilder) {
    return NewCommunicationsPresencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PresencesById provides operations to manage the presences property of the microsoft.graph.cloudCommunications entity.
func (m *CommunicationsRequestBuilder) PresencesById(id string)(*CommunicationsPresencesPresenceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["presence%2Did"] = id
    }
    return NewCommunicationsPresencesPresenceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
