package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemConversationsItemThreadsItemPostsItemReplyRequestBuilder provides operations to call the reply method.
type ItemConversationsItemThreadsItemPostsItemReplyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemConversationsItemThreadsItemPostsItemReplyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConversationsItemThreadsItemPostsItemReplyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemConversationsItemThreadsItemPostsItemReplyRequestBuilderInternal instantiates a new ReplyRequestBuilder and sets the default values.
func NewItemConversationsItemThreadsItemPostsItemReplyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConversationsItemThreadsItemPostsItemReplyRequestBuilder) {
    m := &ItemConversationsItemThreadsItemPostsItemReplyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/conversations/{conversation%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/reply", pathParameters),
    }
    return m
}
// NewItemConversationsItemThreadsItemPostsItemReplyRequestBuilder instantiates a new ReplyRequestBuilder and sets the default values.
func NewItemConversationsItemThreadsItemPostsItemReplyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConversationsItemThreadsItemPostsItemReplyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConversationsItemThreadsItemPostsItemReplyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reply to a post and add a new post to the specified thread in a group conversation.  You can specify both the parent conversation and thread in the request, or, you can specify just the parent thread without the parent conversation. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/post-reply?view=graph-rest-1.0
func (m *ItemConversationsItemThreadsItemPostsItemReplyRequestBuilder) Post(ctx context.Context, body ItemConversationsItemThreadsItemPostsItemReplyPostRequestBodyable, requestConfiguration *ItemConversationsItemThreadsItemPostsItemReplyRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation reply to a post and add a new post to the specified thread in a group conversation.  You can specify both the parent conversation and thread in the request, or, you can specify just the parent thread without the parent conversation. This API is available in the following national cloud deployments.
func (m *ItemConversationsItemThreadsItemPostsItemReplyRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemConversationsItemThreadsItemPostsItemReplyPostRequestBodyable, requestConfiguration *ItemConversationsItemThreadsItemPostsItemReplyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemConversationsItemThreadsItemPostsItemReplyRequestBuilder) WithUrl(rawUrl string)(*ItemConversationsItemThreadsItemPostsItemReplyRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemReplyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
