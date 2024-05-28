package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder provides operations to call the onenotePatchContent method.
type ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilderInternal instantiates a new ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder and sets the default values.
func NewItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder) {
    m := &ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/pages/{onenotePage%2Did}/onenotePatchContent", pathParameters),
    }
    return m
}
// NewItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder instantiates a new ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder and sets the default values.
func NewItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action onenotePatchContent
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder) Post(ctx context.Context, body ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentPostRequestBodyable, requestConfiguration *ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action onenotePatchContent
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentPostRequestBodyable, requestConfiguration *ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder when successful
func (m *ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder) {
    return NewItemSitesItemOnenotePagesItemOnenotepatchcontentOnenotePatchContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
