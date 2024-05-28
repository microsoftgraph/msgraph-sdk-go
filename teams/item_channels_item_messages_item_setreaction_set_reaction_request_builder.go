package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilder provides operations to call the setReaction method.
type ItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilderInternal instantiates a new ItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilder and sets the default values.
func NewItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilder) {
    m := &ItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/channels/{channel%2Did}/messages/{chatMessage%2Did}/setReaction", pathParameters),
    }
    return m
}
// NewItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilder instantiates a new ItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilder and sets the default values.
func NewItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action setReaction
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilder) Post(ctx context.Context, body ItemChannelsItemMessagesItemSetreactionSetReactionPostRequestBodyable, requestConfiguration *ItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemChannelsItemMessagesItemSetreactionSetReactionPostRequestBodyable, requestConfiguration *ItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilder when successful
func (m *ItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilder) WithUrl(rawUrl string)(*ItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilder) {
    return NewItemChannelsItemMessagesItemSetreactionSetReactionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
