package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilder provides operations to call the onenotePatchContent method.
type ItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilderInternal instantiates a new OnenotePatchContentRequestBuilder and sets the default values.
func NewItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilder) {
    m := &ItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onenote/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}/onenotePatchContent", pathParameters),
    }
    return m
}
// NewItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilder instantiates a new OnenotePatchContentRequestBuilder and sets the default values.
func NewItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action onenotePatchContent
func (m *ItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilder) Post(ctx context.Context, body ItemOnenoteSectionsItemPagesItemOnenotePatchContentPostRequestBodyable, requestConfiguration *ItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action onenotePatchContent
func (m *ItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemOnenoteSectionsItemPagesItemOnenotePatchContentPostRequestBodyable, requestConfiguration *ItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilder) WithUrl(rawUrl string)(*ItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilder) {
    return NewItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
