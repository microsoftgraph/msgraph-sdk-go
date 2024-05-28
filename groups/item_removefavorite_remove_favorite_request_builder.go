package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemRemovefavoriteRemoveFavoriteRequestBuilder provides operations to call the removeFavorite method.
type ItemRemovefavoriteRemoveFavoriteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRemovefavoriteRemoveFavoriteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRemovefavoriteRemoveFavoriteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemRemovefavoriteRemoveFavoriteRequestBuilderInternal instantiates a new ItemRemovefavoriteRemoveFavoriteRequestBuilder and sets the default values.
func NewItemRemovefavoriteRemoveFavoriteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemovefavoriteRemoveFavoriteRequestBuilder) {
    m := &ItemRemovefavoriteRemoveFavoriteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/removeFavorite", pathParameters),
    }
    return m
}
// NewItemRemovefavoriteRemoveFavoriteRequestBuilder instantiates a new ItemRemovefavoriteRemoveFavoriteRequestBuilder and sets the default values.
func NewItemRemovefavoriteRemoveFavoriteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemovefavoriteRemoveFavoriteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRemovefavoriteRemoveFavoriteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove the group from the list of the current user's favorite groups. Supported for Microsoft 365 groups only.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/group-removefavorite?view=graph-rest-1.0
func (m *ItemRemovefavoriteRemoveFavoriteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemRemovefavoriteRemoveFavoriteRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation remove the group from the list of the current user's favorite groups. Supported for Microsoft 365 groups only.
// returns a *RequestInformation when successful
func (m *ItemRemovefavoriteRemoveFavoriteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemRemovefavoriteRemoveFavoriteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRemovefavoriteRemoveFavoriteRequestBuilder when successful
func (m *ItemRemovefavoriteRemoveFavoriteRequestBuilder) WithUrl(rawUrl string)(*ItemRemovefavoriteRemoveFavoriteRequestBuilder) {
    return NewItemRemovefavoriteRemoveFavoriteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
