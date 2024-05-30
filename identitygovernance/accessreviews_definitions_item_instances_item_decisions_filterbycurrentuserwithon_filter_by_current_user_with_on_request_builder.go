package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder provides operations to call the filterByCurrentUser method.
type AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetQueryParameters retrieve all decision items for an instance of an access review or a stage of an instance of a multi-stage access review, for which the calling user is the reviewer. The decision items are represented by accessReviewInstanceDecisionItem objects on a given accessReviewInstance or accessReviewStage for which the calling user is the reviewer.
type AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetQueryParameters struct {
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
// AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetQueryParameters
}
// NewAccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal instantiates a new AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, on *string)(*AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    m := &AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/decisions/filterByCurrentUser(on='{on}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if on != nil {
        m.BaseRequestBuilder.PathParameters["on"] = *on
    }
    return m
}
// NewAccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder instantiates a new AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get retrieve all decision items for an instance of an access review or a stage of an instance of a multi-stage access review, for which the calling user is the reviewer. The decision items are represented by accessReviewInstanceDecisionItem objects on a given accessReviewInstance or accessReviewStage for which the calling user is the reviewer.
// Deprecated: This method is obsolete. Use GetAsFilterByCurrentUserWithOnGetResponse instead.
// returns a AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstancedecisionitem-filterbycurrentuser?view=graph-rest-1.0
func (m *AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable), nil
}
// GetAsFilterByCurrentUserWithOnGetResponse retrieve all decision items for an instance of an access review or a stage of an instance of a multi-stage access review, for which the calling user is the reviewer. The decision items are represented by accessReviewInstanceDecisionItem objects on a given accessReviewInstance or accessReviewStage for which the calling user is the reviewer.
// returns a AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstancedecisionitem-filterbycurrentuser?view=graph-rest-1.0
func (m *AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) GetAsFilterByCurrentUserWithOnGetResponse(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable), nil
}
// ToGetRequestInformation retrieve all decision items for an instance of an access review or a stage of an instance of a multi-stage access review, for which the calling user is the reviewer. The decision items are represented by accessReviewInstanceDecisionItem objects on a given accessReviewInstance or accessReviewStage for which the calling user is the reviewer.
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemDecisionsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
