package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder provides operations to call the unsetReaction method.
type ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderInternal instantiates a new ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder and sets the default values.
func NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) {
    m := &ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/primaryChannel/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/unsetReaction", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder instantiates a new ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder and sets the default values.
func NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action unsetReaction
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) Post(ctx context.Context, body ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionPostRequestBodyable, requestConfiguration *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action unsetReaction
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionPostRequestBodyable, requestConfiguration *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelMessagesItemRepliesItemUnsetreactionUnsetReactionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
