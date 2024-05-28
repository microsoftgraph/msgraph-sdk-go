package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19 "github.com/microsoftgraph/msgraph-sdk-go/models/callrecords"
)

// CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder provides operations to manage the segments property of the microsoft.graph.callRecords.session entity.
type CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderGetQueryParameters the list of segments involved in the session. Read-only. Nullable.
type CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderGetQueryParameters
}
// CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderInternal instantiates a new CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder and sets the default values.
func NewCallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) {
    m := &CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/callRecords/{callRecord%2Did}/sessions/{session%2Did}/segments/{segment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder instantiates a new CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder and sets the default values.
func NewCallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property segments for communications
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the list of segments involved in the session. Read-only. Nullable.
// returns a Segmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderGetRequestConfiguration)(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Segmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property segments in communications
// returns a Segmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) Patch(ctx context.Context, body iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Segmentable, requestConfiguration *CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderPatchRequestConfiguration)(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Segmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property segments for communications
// returns a *RequestInformation when successful
func (m *CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of segments involved in the session. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property segments in communications
// returns a *RequestInformation when successful
func (m *CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Segmentable, requestConfiguration *CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder when successful
func (m *CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) WithUrl(rawUrl string)(*CallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder) {
    return NewCallrecordsItemSessionsItemSegmentsSegmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
