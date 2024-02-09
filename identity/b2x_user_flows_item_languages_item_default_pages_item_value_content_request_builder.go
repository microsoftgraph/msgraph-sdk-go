package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder provides operations to manage the media for the identityContainer entity.
type B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderGetQueryParameters get media content for the navigation property defaultPages from identity
type B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderGetQueryParameters struct {
    // Format of the content
    Format *string `uriparametername:"%24format"`
}
// B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderGetQueryParameters
}
// B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderInternal instantiates a new B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder and sets the default values.
func NewB2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder) {
    m := &B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/languages/{userFlowLanguageConfiguration%2Did}/defaultPages/{userFlowLanguagePage%2Did}/$value{?%24format*}", pathParameters),
    }
    return m
}
// NewB2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder instantiates a new B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder and sets the default values.
func NewB2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get media content for the navigation property defaultPages from identity
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put update media content for the navigation property defaultPages in identity
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get media content for the navigation property defaultPages from identity
// returns a *RequestInformation when successful
func (m *B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation update media content for the navigation property defaultPages in identity
// returns a *RequestInformation when successful
func (m *B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/languages/{userFlowLanguageConfiguration%2Did}/defaultPages/{userFlowLanguagePage%2Did}/$value", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetStreamContentAndContentType(body, "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder when successful
func (m *B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder) WithUrl(rawUrl string)(*B2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder) {
    return NewB2xUserFlowsItemLanguagesItemDefaultPagesItemValueContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
