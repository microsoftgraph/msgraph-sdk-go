package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19 "github.com/microsoftgraph/msgraph-sdk-go/models/callrecords"
)

// CallrecordsItemSessionsItemSegmentsRequestBuilder provides operations to manage the segments property of the microsoft.graph.callRecords.session entity.
type CallrecordsItemSessionsItemSegmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallrecordsItemSessionsItemSegmentsRequestBuilderGetQueryParameters the list of segments involved in the session. Read-only. Nullable.
type CallrecordsItemSessionsItemSegmentsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// CallrecordsItemSessionsItemSegmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallrecordsItemSessionsItemSegmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallrecordsItemSessionsItemSegmentsRequestBuilderGetQueryParameters
}
// CallrecordsItemSessionsItemSegmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallrecordsItemSessionsItemSegmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySegmentId provides operations to manage the segments property of the microsoft.graph.callRecords.session entity.
// returns a *CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder when successful
func (m *CallrecordsItemSessionsItemSegmentsRequestBuilder) BySegmentId(segmentId string)(*CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if segmentId != "" {
        urlTplParams["segment%2Did"] = segmentId
    }
    return NewCallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCallrecordsItemSessionsItemSegmentsRequestBuilderInternal instantiates a new CallrecordsItemSessionsItemSegmentsRequestBuilder and sets the default values.
func NewCallrecordsItemSessionsItemSegmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallrecordsItemSessionsItemSegmentsRequestBuilder) {
    m := &CallrecordsItemSessionsItemSegmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/callRecords/{callRecord%2Did}/sessions/{session%2Did}/segments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCallrecordsItemSessionsItemSegmentsRequestBuilder instantiates a new CallrecordsItemSessionsItemSegmentsRequestBuilder and sets the default values.
func NewCallrecordsItemSessionsItemSegmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallrecordsItemSessionsItemSegmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallrecordsItemSessionsItemSegmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CallrecordsItemSessionsItemSegmentsCountRequestBuilder when successful
func (m *CallrecordsItemSessionsItemSegmentsRequestBuilder) Count()(*CallrecordsItemSessionsItemSegmentsCountRequestBuilder) {
    return NewCallrecordsItemSessionsItemSegmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of segments involved in the session. Read-only. Nullable.
// returns a SegmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsItemSessionsItemSegmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *CallrecordsItemSessionsItemSegmentsRequestBuilderGetRequestConfiguration)(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.SegmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.CreateSegmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.SegmentCollectionResponseable), nil
}
// Post create new navigation property to segments for communications
// returns a Segmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsItemSessionsItemSegmentsRequestBuilder) Post(ctx context.Context, body iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Segmentable, requestConfiguration *CallrecordsItemSessionsItemSegmentsRequestBuilderPostRequestConfiguration)(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Segmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.CreateSegmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Segmentable), nil
}
// ToGetRequestInformation the list of segments involved in the session. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *CallrecordsItemSessionsItemSegmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallrecordsItemSessionsItemSegmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to segments for communications
// returns a *RequestInformation when successful
func (m *CallrecordsItemSessionsItemSegmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Segmentable, requestConfiguration *CallrecordsItemSessionsItemSegmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CallrecordsItemSessionsItemSegmentsRequestBuilder when successful
func (m *CallrecordsItemSessionsItemSegmentsRequestBuilder) WithUrl(rawUrl string)(*CallrecordsItemSessionsItemSegmentsRequestBuilder) {
    return NewCallrecordsItemSessionsItemSegmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
