package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemChatsItemSendActivityNotificationRequestBuilder provides operations to call the sendActivityNotification method.
type ItemChatsItemSendActivityNotificationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChatsItemSendActivityNotificationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemSendActivityNotificationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChatsItemSendActivityNotificationRequestBuilderInternal instantiates a new SendActivityNotificationRequestBuilder and sets the default values.
func NewItemChatsItemSendActivityNotificationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemSendActivityNotificationRequestBuilder) {
    m := &ItemChatsItemSendActivityNotificationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}/sendActivityNotification", pathParameters),
    }
    return m
}
// NewItemChatsItemSendActivityNotificationRequestBuilder instantiates a new SendActivityNotificationRequestBuilder and sets the default values.
func NewItemChatsItemSendActivityNotificationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemSendActivityNotificationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChatsItemSendActivityNotificationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send an activity feed notification in scope of a chat. For more details about sending notifications and the requirements for doing so, see sending Teams activity notifications.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chat-sendactivitynotification?view=graph-rest-1.0
func (m *ItemChatsItemSendActivityNotificationRequestBuilder) Post(ctx context.Context, body ItemChatsItemSendActivityNotificationPostRequestBodyable, requestConfiguration *ItemChatsItemSendActivityNotificationRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation send an activity feed notification in scope of a chat. For more details about sending notifications and the requirements for doing so, see sending Teams activity notifications.
func (m *ItemChatsItemSendActivityNotificationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemChatsItemSendActivityNotificationPostRequestBodyable, requestConfiguration *ItemChatsItemSendActivityNotificationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemChatsItemSendActivityNotificationRequestBuilder) WithUrl(rawUrl string)(*ItemChatsItemSendActivityNotificationRequestBuilder) {
    return NewItemChatsItemSendActivityNotificationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
