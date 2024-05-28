package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
type ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilderGetQueryParameters retrieve the properties and relationships of all offerShiftRequest objects in a team.
type ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilderGetQueryParameters struct {
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
// ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilderGetQueryParameters
}
// ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByOfferShiftRequestId provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
// returns a *ItemScheduleOffershiftrequestsOfferShiftRequestItemRequestBuilder when successful
func (m *ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder) ByOfferShiftRequestId(offerShiftRequestId string)(*ItemScheduleOffershiftrequestsOfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if offerShiftRequestId != "" {
        urlTplParams["offerShiftRequest%2Did"] = offerShiftRequestId
    }
    return NewItemScheduleOffershiftrequestsOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilderInternal instantiates a new ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder and sets the default values.
func NewItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder) {
    m := &ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/schedule/offerShiftRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder instantiates a new ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder and sets the default values.
func NewItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemScheduleOffershiftrequestsCountRequestBuilder when successful
func (m *ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder) Count()(*ItemScheduleOffershiftrequestsCountRequestBuilder) {
    return NewItemScheduleOffershiftrequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the properties and relationships of all offerShiftRequest objects in a team.
// returns a OfferShiftRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/offershiftrequest-list?view=graph-rest-1.0
func (m *ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OfferShiftRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOfferShiftRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OfferShiftRequestCollectionResponseable), nil
}
// Post create an instance of an offerShiftRequest.
// returns a OfferShiftRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/offershiftrequest-post?view=graph-rest-1.0
func (m *ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OfferShiftRequestable, requestConfiguration *ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OfferShiftRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOfferShiftRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OfferShiftRequestable), nil
}
// ToGetRequestInformation retrieve the properties and relationships of all offerShiftRequest objects in a team.
// returns a *RequestInformation when successful
func (m *ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create an instance of an offerShiftRequest.
// returns a *RequestInformation when successful
func (m *ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OfferShiftRequestable, requestConfiguration *ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder when successful
func (m *ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder) WithUrl(rawUrl string)(*ItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder) {
    return NewItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
