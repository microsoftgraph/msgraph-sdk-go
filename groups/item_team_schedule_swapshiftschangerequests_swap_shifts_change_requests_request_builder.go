package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
type ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilderGetQueryParameters the swap requests for shifts in the schedule.
type ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilderGetQueryParameters struct {
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
// ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilderGetQueryParameters
}
// ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySwapShiftsChangeRequestId provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
// returns a *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder when successful
func (m *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder) BySwapShiftsChangeRequestId(swapShiftsChangeRequestId string)(*ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if swapShiftsChangeRequestId != "" {
        urlTplParams["swapShiftsChangeRequest%2Did"] = swapShiftsChangeRequestId
    }
    return NewItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilderInternal instantiates a new ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder and sets the default values.
func NewItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder) {
    m := &ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/schedule/swapShiftsChangeRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder instantiates a new ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder and sets the default values.
func NewItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTeamScheduleSwapshiftschangerequestsCountRequestBuilder when successful
func (m *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder) Count()(*ItemTeamScheduleSwapshiftschangerequestsCountRequestBuilder) {
    return NewItemTeamScheduleSwapshiftschangerequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the swap requests for shifts in the schedule.
// returns a SwapShiftsChangeRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SwapShiftsChangeRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSwapShiftsChangeRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SwapShiftsChangeRequestCollectionResponseable), nil
}
// Post create new navigation property to swapShiftsChangeRequests for groups
// returns a SwapShiftsChangeRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SwapShiftsChangeRequestable, requestConfiguration *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SwapShiftsChangeRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSwapShiftsChangeRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SwapShiftsChangeRequestable), nil
}
// ToGetRequestInformation the swap requests for shifts in the schedule.
// returns a *RequestInformation when successful
func (m *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to swapShiftsChangeRequests for groups
// returns a *RequestInformation when successful
func (m *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SwapShiftsChangeRequestable, requestConfiguration *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder when successful
func (m *ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder) WithUrl(rawUrl string)(*ItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder) {
    return NewItemTeamScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
