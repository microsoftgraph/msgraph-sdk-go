package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilder provides operations to call the sendActivityNotificationToRecipients method.
type SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilderInternal instantiates a new SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilder and sets the default values.
func NewSendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilder) {
    m := &SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/sendActivityNotificationToRecipients", pathParameters),
    }
    return m
}
// NewSendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilder instantiates a new SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilder and sets the default values.
func NewSendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send activity feed notifications to multiple users, in bulk.  For more information, see sending Teams activity notifications.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/teamwork-sendactivitynotificationtorecipients?view=graph-rest-1.0
func (m *SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilder) Post(ctx context.Context, body SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsPostRequestBodyable, requestConfiguration *SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation send activity feed notifications to multiple users, in bulk.  For more information, see sending Teams activity notifications.
// returns a *RequestInformation when successful
func (m *SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilder) ToPostRequestInformation(ctx context.Context, body SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsPostRequestBodyable, requestConfiguration *SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilder when successful
func (m *SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilder) WithUrl(rawUrl string)(*SendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilder) {
    return NewSendactivitynotificationtorecipientsSendActivityNotificationToRecipientsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
