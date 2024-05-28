package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder provides operations to call the resetDecisions method.
type AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilderInternal instantiates a new AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder) {
    m := &AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/resetDecisions", pathParameters),
    }
    return m
}
// NewAccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder instantiates a new AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post resets all accessReviewInstanceDecisionItem objects on an accessReviewInstance to notReviewed.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-resetdecisions?view=graph-rest-1.0
func (m *AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder) Post(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation resets all accessReviewInstanceDecisionItem objects on an accessReviewInstance to notReviewed.
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemResetdecisionsResetDecisionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
