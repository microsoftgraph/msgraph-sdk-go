package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder provides operations to call the setReaction method.
type ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilderInternal instantiates a new ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder and sets the default values.
func NewItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder) {
    m := &ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/primaryChannel/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/setReaction", pathParameters),
    }
    return m
}
// NewItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder instantiates a new ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder and sets the default values.
func NewItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action setReaction
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder) Post(ctx context.Context, body ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionPostRequestBodyable, requestConfiguration *ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionPostRequestBodyable, requestConfiguration *ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder when successful
func (m *ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder) WithUrl(rawUrl string)(*ItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder) {
    return NewItemTeamPrimarychannelMessagesItemRepliesItemSetreactionSetReactionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
