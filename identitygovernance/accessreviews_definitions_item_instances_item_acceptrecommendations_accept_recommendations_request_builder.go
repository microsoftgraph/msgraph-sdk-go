package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder provides operations to call the acceptRecommendations method.
type AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilderInternal instantiates a new AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder) {
    m := &AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/acceptRecommendations", pathParameters),
    }
    return m
}
// NewAccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder instantiates a new AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post allows the acceptance of recommendations on all accessReviewInstanceDecisionItem objects that haven't been reviewed on an accessReviewInstance object for which the calling user is a reviewer.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-acceptrecommendations?view=graph-rest-1.0
func (m *AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder) Post(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation allows the acceptance of recommendations on all accessReviewInstanceDecisionItem objects that haven't been reviewed on an accessReviewInstance object for which the calling user is a reviewer.
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemAcceptrecommendationsAcceptRecommendationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
