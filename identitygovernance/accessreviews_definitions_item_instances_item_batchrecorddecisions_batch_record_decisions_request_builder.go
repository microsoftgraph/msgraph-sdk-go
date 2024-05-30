package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder provides operations to call the batchRecordDecisions method.
type AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilderInternal instantiates a new AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) {
    m := &AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/batchRecordDecisions", pathParameters),
    }
    return m
}
// NewAccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder instantiates a new AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post enables reviewers to review all accessReviewInstanceDecisionItem objects in batches by using principalId, resourceId, or neither.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-batchrecorddecisions?view=graph-rest-1.0
func (m *AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) Post(ctx context.Context, body AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsPostRequestBodyable, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation enables reviewers to review all accessReviewInstanceDecisionItem objects in batches by using principalId, resourceId, or neither.
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsPostRequestBodyable, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
