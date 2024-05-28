package chats

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemUnhideforuserUnhideForUserRequestBuilder provides operations to call the unhideForUser method.
type ItemUnhideforuserUnhideForUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemUnhideforuserUnhideForUserRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemUnhideforuserUnhideForUserRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemUnhideforuserUnhideForUserRequestBuilderInternal instantiates a new ItemUnhideforuserUnhideForUserRequestBuilder and sets the default values.
func NewItemUnhideforuserUnhideForUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUnhideforuserUnhideForUserRequestBuilder) {
    m := &ItemUnhideforuserUnhideForUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/chats/{chat%2Did}/unhideForUser", pathParameters),
    }
    return m
}
// NewItemUnhideforuserUnhideForUserRequestBuilder instantiates a new ItemUnhideforuserUnhideForUserRequestBuilder and sets the default values.
func NewItemUnhideforuserUnhideForUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUnhideforuserUnhideForUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemUnhideforuserUnhideForUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unhide a chat for a user.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chat-unhideforuser?view=graph-rest-1.0
func (m *ItemUnhideforuserUnhideForUserRequestBuilder) Post(ctx context.Context, body ItemUnhideforuserUnhideForUserPostRequestBodyable, requestConfiguration *ItemUnhideforuserUnhideForUserRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation unhide a chat for a user.
// returns a *RequestInformation when successful
func (m *ItemUnhideforuserUnhideForUserRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemUnhideforuserUnhideForUserPostRequestBodyable, requestConfiguration *ItemUnhideforuserUnhideForUserRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemUnhideforuserUnhideForUserRequestBuilder when successful
func (m *ItemUnhideforuserUnhideForUserRequestBuilder) WithUrl(rawUrl string)(*ItemUnhideforuserUnhideForUserRequestBuilder) {
    return NewItemUnhideforuserUnhideForUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
