package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemChatsItemMarkChatReadForUserRequestBuilder provides operations to call the markChatReadForUser method.
type ItemChatsItemMarkChatReadForUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChatsItemMarkChatReadForUserRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemMarkChatReadForUserRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChatsItemMarkChatReadForUserRequestBuilderInternal instantiates a new MarkChatReadForUserRequestBuilder and sets the default values.
func NewItemChatsItemMarkChatReadForUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemMarkChatReadForUserRequestBuilder) {
    m := &ItemChatsItemMarkChatReadForUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}/markChatReadForUser", pathParameters),
    }
    return m
}
// NewItemChatsItemMarkChatReadForUserRequestBuilder instantiates a new MarkChatReadForUserRequestBuilder and sets the default values.
func NewItemChatsItemMarkChatReadForUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemMarkChatReadForUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChatsItemMarkChatReadForUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Post mark a chat as read for a user.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chat-markchatreadforuser?view=graph-rest-1.0
func (m *ItemChatsItemMarkChatReadForUserRequestBuilder) Post(ctx context.Context, body ItemChatsItemMarkChatReadForUserPostRequestBodyable, requestConfiguration *ItemChatsItemMarkChatReadForUserRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation mark a chat as read for a user.
func (m *ItemChatsItemMarkChatReadForUserRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemChatsItemMarkChatReadForUserPostRequestBodyable, requestConfiguration *ItemChatsItemMarkChatReadForUserRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemChatsItemMarkChatReadForUserRequestBuilder) WithUrl(rawUrl string)(*ItemChatsItemMarkChatReadForUserRequestBuilder) {
    return NewItemChatsItemMarkChatReadForUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
