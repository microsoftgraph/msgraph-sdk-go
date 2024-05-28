package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder provides operations to manage the contentSharingSessions property of the microsoft.graph.call entity.
type CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderGetQueryParameters retrieve a list of contentSharingSession objects in a call.
type CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderGetQueryParameters struct {
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
// CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderGetQueryParameters
}
// CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByContentSharingSessionId provides operations to manage the contentSharingSessions property of the microsoft.graph.call entity.
// returns a *CallsItemContentsharingsessionsContentSharingSessionItemRequestBuilder when successful
func (m *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) ByContentSharingSessionId(contentSharingSessionId string)(*CallsItemContentsharingsessionsContentSharingSessionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if contentSharingSessionId != "" {
        urlTplParams["contentSharingSession%2Did"] = contentSharingSessionId
    }
    return NewCallsItemContentsharingsessionsContentSharingSessionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCallsItemContentsharingsessionsContentSharingSessionsRequestBuilderInternal instantiates a new CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder and sets the default values.
func NewCallsItemContentsharingsessionsContentSharingSessionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) {
    m := &CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/calls/{call%2Did}/contentSharingSessions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCallsItemContentsharingsessionsContentSharingSessionsRequestBuilder instantiates a new CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder and sets the default values.
func NewCallsItemContentsharingsessionsContentSharingSessionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemContentsharingsessionsContentSharingSessionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CallsItemContentsharingsessionsCountRequestBuilder when successful
func (m *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) Count()(*CallsItemContentsharingsessionsCountRequestBuilder) {
    return NewCallsItemContentsharingsessionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of contentSharingSession objects in a call.
// returns a ContentSharingSessionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/call-list-contentsharingsessions?view=graph-rest-1.0
func (m *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) Get(ctx context.Context, requestConfiguration *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentSharingSessionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContentSharingSessionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentSharingSessionCollectionResponseable), nil
}
// Post create new navigation property to contentSharingSessions for communications
// returns a ContentSharingSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentSharingSessionable, requestConfiguration *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentSharingSessionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContentSharingSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentSharingSessionable), nil
}
// ToGetRequestInformation retrieve a list of contentSharingSession objects in a call.
// returns a *RequestInformation when successful
func (m *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to contentSharingSessions for communications
// returns a *RequestInformation when successful
func (m *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentSharingSessionable, requestConfiguration *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder when successful
func (m *CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) WithUrl(rawUrl string)(*CallsItemContentsharingsessionsContentSharingSessionsRequestBuilder) {
    return NewCallsItemContentsharingsessionsContentSharingSessionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
