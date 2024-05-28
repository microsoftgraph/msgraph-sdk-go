package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder provides operations to call the sendReminder method.
type AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilderInternal instantiates a new AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder) {
    m := &AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/sendReminder", pathParameters),
    }
    return m
}
// NewAccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder instantiates a new AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send a reminder to the reviewers of an active accessReviewInstance.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-sendreminder?view=graph-rest-1.0
func (m *AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder) Post(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation send a reminder to the reviewers of an active accessReviewInstance.
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemSendreminderSendReminderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
