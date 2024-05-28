package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemFollowedsitesAddRequestBuilder provides operations to call the add method.
type ItemFollowedsitesAddRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemFollowedsitesAddRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFollowedsitesAddRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemFollowedsitesAddRequestBuilderInternal instantiates a new ItemFollowedsitesAddRequestBuilder and sets the default values.
func NewItemFollowedsitesAddRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFollowedsitesAddRequestBuilder) {
    m := &ItemFollowedsitesAddRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/followedSites/add", pathParameters),
    }
    return m
}
// NewItemFollowedsitesAddRequestBuilder instantiates a new ItemFollowedsitesAddRequestBuilder and sets the default values.
func NewItemFollowedsitesAddRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFollowedsitesAddRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFollowedsitesAddRequestBuilderInternal(urlParams, requestAdapter)
}
// Post follow a user's site or multiple sites.
// Deprecated: This method is obsolete. Use PostAsAddPostResponse instead.
// returns a ItemFollowedsitesAddResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/site-follow?view=graph-rest-1.0
func (m *ItemFollowedsitesAddRequestBuilder) Post(ctx context.Context, body ItemFollowedsitesAddPostRequestBodyable, requestConfiguration *ItemFollowedsitesAddRequestBuilderPostRequestConfiguration)(ItemFollowedsitesAddResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemFollowedsitesAddResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemFollowedsitesAddResponseable), nil
}
// PostAsAddPostResponse follow a user's site or multiple sites.
// returns a ItemFollowedsitesAddPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/site-follow?view=graph-rest-1.0
func (m *ItemFollowedsitesAddRequestBuilder) PostAsAddPostResponse(ctx context.Context, body ItemFollowedsitesAddPostRequestBodyable, requestConfiguration *ItemFollowedsitesAddRequestBuilderPostRequestConfiguration)(ItemFollowedsitesAddPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemFollowedsitesAddPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemFollowedsitesAddPostResponseable), nil
}
// ToPostRequestInformation follow a user's site or multiple sites.
// returns a *RequestInformation when successful
func (m *ItemFollowedsitesAddRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemFollowedsitesAddPostRequestBodyable, requestConfiguration *ItemFollowedsitesAddRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemFollowedsitesAddRequestBuilder when successful
func (m *ItemFollowedsitesAddRequestBuilder) WithUrl(rawUrl string)(*ItemFollowedsitesAddRequestBuilder) {
    return NewItemFollowedsitesAddRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
