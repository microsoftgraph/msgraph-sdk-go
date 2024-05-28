package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder provides operations to manage the defaultPages property of the microsoft.graph.userFlowLanguageConfiguration entity.
type B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderGetQueryParameters collection of pages with the default content to display in a user flow for a specified language. This collection doesn't allow any kind of modification.
type B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderGetQueryParameters
}
// B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderInternal instantiates a new B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder and sets the default values.
func NewB2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder) {
    m := &B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/languages/{userFlowLanguageConfiguration%2Did}/defaultPages/{userFlowLanguagePage%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewB2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder instantiates a new B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder and sets the default values.
func NewB2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the identityContainer entity.
// returns a *B2xuserflowsItemLanguagesItemDefaultpagesItemValueContentRequestBuilder when successful
func (m *B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder) Content()(*B2xuserflowsItemLanguagesItemDefaultpagesItemValueContentRequestBuilder) {
    return NewB2xuserflowsItemLanguagesItemDefaultpagesItemValueContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property defaultPages for identity
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get collection of pages with the default content to display in a user flow for a specified language. This collection doesn't allow any kind of modification.
// returns a UserFlowLanguagePageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserFlowLanguagePageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFlowLanguagePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserFlowLanguagePageable), nil
}
// Patch update the navigation property defaultPages in identity
// returns a UserFlowLanguagePageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserFlowLanguagePageable, requestConfiguration *B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserFlowLanguagePageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFlowLanguagePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserFlowLanguagePageable), nil
}
// ToDeleteRequestInformation delete navigation property defaultPages for identity
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of pages with the default content to display in a user flow for a specified language. This collection doesn't allow any kind of modification.
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property defaultPages in identity
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserFlowLanguagePageable, requestConfiguration *B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder when successful
func (m *B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder) WithUrl(rawUrl string)(*B2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder) {
    return NewB2xuserflowsItemLanguagesItemDefaultpagesUserFlowLanguagePageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
