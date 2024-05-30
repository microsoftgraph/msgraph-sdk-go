package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder provides operations to call the validateAuthenticationConfiguration method.
type CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilderInternal instantiates a new CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder and sets the default values.
func NewCustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder) {
    m := &CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/customAuthenticationExtensions/validateAuthenticationConfiguration", pathParameters),
    }
    return m
}
// NewCustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder instantiates a new CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder and sets the default values.
func NewCustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action validateAuthenticationConfiguration
// returns a AuthenticationConfigurationValidationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder) Post(ctx context.Context, body CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationPostRequestBodyable, requestConfiguration *CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConfigurationValidationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action validateAuthenticationConfiguration
// returns a *RequestInformation when successful
func (m *CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder) ToPostRequestInformation(ctx context.Context, body CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationPostRequestBodyable, requestConfiguration *CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder when successful
func (m *CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder) WithUrl(rawUrl string)(*CustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder) {
    return NewCustomauthenticationextensionsValidateauthenticationconfigurationValidateAuthenticationConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
