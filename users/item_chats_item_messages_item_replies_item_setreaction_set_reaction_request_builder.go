package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilder provides operations to call the setReaction method.
type ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilderInternal instantiates a new ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilder and sets the default values.
func NewItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilder) {
    m := &ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/setReaction", pathParameters),
    }
    return m
}
// NewItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilder instantiates a new ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilder and sets the default values.
func NewItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action setReaction
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilder) Post(ctx context.Context, body ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionPostRequestBodyable, requestConfiguration *ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action setReaction
// returns a *RequestInformation when successful
func (m *ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionPostRequestBodyable, requestConfiguration *ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilder when successful
func (m *ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilder) WithUrl(rawUrl string)(*ItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilder) {
    return NewItemChatsItemMessagesItemRepliesItemSetreactionSetReactionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
