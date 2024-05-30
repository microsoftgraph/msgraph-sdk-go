package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilder provides operations to call the validateCredentials method.
type ItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilderInternal instantiates a new ItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilder and sets the default values.
func NewItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilder) {
    m := &ItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/synchronization/jobs/validateCredentials", pathParameters),
    }
    return m
}
// NewItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilder instantiates a new ItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilder and sets the default values.
func NewItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action validateCredentials
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilder) Post(ctx context.Context, body ItemSynchronizationJobsValidatecredentialsValidateCredentialsPostRequestBodyable, requestConfiguration *ItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action validateCredentials
// returns a *RequestInformation when successful
func (m *ItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSynchronizationJobsValidatecredentialsValidateCredentialsPostRequestBodyable, requestConfiguration *ItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilder when successful
func (m *ItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilder) {
    return NewItemSynchronizationJobsValidatecredentialsValidateCredentialsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
