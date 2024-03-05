package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemSitesAddRequestBuilder provides operations to call the add method.
type ItemSitesItemSitesAddRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemSitesAddRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemSitesAddRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemSitesAddRequestBuilderInternal instantiates a new ItemSitesItemSitesAddRequestBuilder and sets the default values.
func NewItemSitesItemSitesAddRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemSitesAddRequestBuilder) {
    m := &ItemSitesItemSitesAddRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/sites/add", pathParameters),
    }
    return m
}
// NewItemSitesItemSitesAddRequestBuilder instantiates a new ItemSitesItemSitesAddRequestBuilder and sets the default values.
func NewItemSitesItemSitesAddRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemSitesAddRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemSitesAddRequestBuilderInternal(urlParams, requestAdapter)
}
// Post follow a user's site or multiple sites.
// Deprecated: This method is obsolete. Use PostAsAddPostResponse instead.
// returns a ItemSitesItemSitesAddResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/site-follow?view=graph-rest-1.0
func (m *ItemSitesItemSitesAddRequestBuilder) Post(ctx context.Context, body ItemSitesItemSitesAddPostRequestBodyable, requestConfiguration *ItemSitesItemSitesAddRequestBuilderPostRequestConfiguration)(ItemSitesItemSitesAddResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemSitesAddResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemSitesAddResponseable), nil
}
// PostAsAddPostResponse follow a user's site or multiple sites.
// returns a ItemSitesItemSitesAddPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/site-follow?view=graph-rest-1.0
func (m *ItemSitesItemSitesAddRequestBuilder) PostAsAddPostResponse(ctx context.Context, body ItemSitesItemSitesAddPostRequestBodyable, requestConfiguration *ItemSitesItemSitesAddRequestBuilderPostRequestConfiguration)(ItemSitesItemSitesAddPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemSitesAddPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemSitesAddPostResponseable), nil
}
// ToPostRequestInformation follow a user's site or multiple sites.
// returns a *RequestInformation when successful
func (m *ItemSitesItemSitesAddRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemSitesAddPostRequestBodyable, requestConfiguration *ItemSitesItemSitesAddRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemSitesAddRequestBuilder when successful
func (m *ItemSitesItemSitesAddRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemSitesAddRequestBuilder) {
    return NewItemSitesItemSitesAddRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
