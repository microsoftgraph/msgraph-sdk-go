package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilder provides operations to call the validateCredentials method.
type ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilderInternal instantiates a new ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilder) {
    m := &ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/synchronization/jobs/{synchronizationJob%2Did}/validateCredentials", pathParameters),
    }
    return m
}
// NewItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilder instantiates a new ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post validate that the credentials are valid in the tenant.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/synchronization-synchronizationjob-validatecredentials?view=graph-rest-1.0
func (m *ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilder) Post(ctx context.Context, body ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsPostRequestBodyable, requestConfiguration *ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation validate that the credentials are valid in the tenant.
// returns a *RequestInformation when successful
func (m *ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsPostRequestBodyable, requestConfiguration *ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilder when successful
func (m *ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilder) {
    return NewItemSynchronizationJobsItemValidatecredentialsValidateCredentialsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
