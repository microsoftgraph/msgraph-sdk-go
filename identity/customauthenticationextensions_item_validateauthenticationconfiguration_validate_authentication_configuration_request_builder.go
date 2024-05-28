package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder provides operations to call the validateAuthenticationConfiguration method.
type CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilderInternal instantiates a new CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder and sets the default values.
func NewCustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder) {
    m := &CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/customAuthenticationExtensions/{customAuthenticationExtension%2Did}/validateAuthenticationConfiguration", pathParameters),
    }
    return m
}
// NewCustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder instantiates a new CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder and sets the default values.
func NewCustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post an API to check validity of the endpoint and and authentication configuration for a customAuthenticationExtension.
// returns a AuthenticationConfigurationValidationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/customauthenticationextension-validateauthenticationconfiguration?view=graph-rest-1.0
func (m *CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder) Post(ctx context.Context, requestConfiguration *CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConfigurationValidationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationConfigurationValidationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConfigurationValidationable), nil
}
// ToPostRequestInformation an API to check validity of the endpoint and and authentication configuration for a customAuthenticationExtension.
// returns a *RequestInformation when successful
func (m *CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder when successful
func (m *CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder) WithUrl(rawUrl string)(*CustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder) {
    return NewCustomauthenticationextensionsItemValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
