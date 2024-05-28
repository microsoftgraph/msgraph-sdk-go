package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAddfavoriteAddFavoriteRequestBuilder provides operations to call the addFavorite method.
type ItemAddfavoriteAddFavoriteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAddfavoriteAddFavoriteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAddfavoriteAddFavoriteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAddfavoriteAddFavoriteRequestBuilderInternal instantiates a new ItemAddfavoriteAddFavoriteRequestBuilder and sets the default values.
func NewItemAddfavoriteAddFavoriteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddfavoriteAddFavoriteRequestBuilder) {
    m := &ItemAddfavoriteAddFavoriteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/addFavorite", pathParameters),
    }
    return m
}
// NewItemAddfavoriteAddFavoriteRequestBuilder instantiates a new ItemAddfavoriteAddFavoriteRequestBuilder and sets the default values.
func NewItemAddfavoriteAddFavoriteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAddfavoriteAddFavoriteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAddfavoriteAddFavoriteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add the group to the list of the current user's favorite groups. Supported for Microsoft 365 groups only.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/group-addfavorite?view=graph-rest-1.0
func (m *ItemAddfavoriteAddFavoriteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemAddfavoriteAddFavoriteRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation add the group to the list of the current user's favorite groups. Supported for Microsoft 365 groups only.
// returns a *RequestInformation when successful
func (m *ItemAddfavoriteAddFavoriteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemAddfavoriteAddFavoriteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAddfavoriteAddFavoriteRequestBuilder when successful
func (m *ItemAddfavoriteAddFavoriteRequestBuilder) WithUrl(rawUrl string)(*ItemAddfavoriteAddFavoriteRequestBuilder) {
    return NewItemAddfavoriteAddFavoriteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
