package chats

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSendactivitynotificationSendActivityNotificationRequestBuilder provides operations to call the sendActivityNotification method.
type ItemSendactivitynotificationSendActivityNotificationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSendactivitynotificationSendActivityNotificationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSendactivitynotificationSendActivityNotificationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSendactivitynotificationSendActivityNotificationRequestBuilderInternal instantiates a new ItemSendactivitynotificationSendActivityNotificationRequestBuilder and sets the default values.
func NewItemSendactivitynotificationSendActivityNotificationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSendactivitynotificationSendActivityNotificationRequestBuilder) {
    m := &ItemSendactivitynotificationSendActivityNotificationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/chats/{chat%2Did}/sendActivityNotification", pathParameters),
    }
    return m
}
// NewItemSendactivitynotificationSendActivityNotificationRequestBuilder instantiates a new ItemSendactivitynotificationSendActivityNotificationRequestBuilder and sets the default values.
func NewItemSendactivitynotificationSendActivityNotificationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSendactivitynotificationSendActivityNotificationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSendactivitynotificationSendActivityNotificationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send an activity feed notification in scope of a chat. For more information about sending notifications and the requirements for doing so, see sending Teams activity notifications.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chat-sendactivitynotification?view=graph-rest-1.0
func (m *ItemSendactivitynotificationSendActivityNotificationRequestBuilder) Post(ctx context.Context, body ItemSendactivitynotificationSendActivityNotificationPostRequestBodyable, requestConfiguration *ItemSendactivitynotificationSendActivityNotificationRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation send an activity feed notification in scope of a chat. For more information about sending notifications and the requirements for doing so, see sending Teams activity notifications.
// returns a *RequestInformation when successful
func (m *ItemSendactivitynotificationSendActivityNotificationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSendactivitynotificationSendActivityNotificationPostRequestBodyable, requestConfiguration *ItemSendactivitynotificationSendActivityNotificationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSendactivitynotificationSendActivityNotificationRequestBuilder when successful
func (m *ItemSendactivitynotificationSendActivityNotificationRequestBuilder) WithUrl(rawUrl string)(*ItemSendactivitynotificationSendActivityNotificationRequestBuilder) {
    return NewItemSendactivitynotificationSendActivityNotificationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
