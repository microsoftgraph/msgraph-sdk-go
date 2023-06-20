package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SendActivityNotificationToRecipientsRequestBuilder provides operations to call the sendActivityNotificationToRecipients method.
type SendActivityNotificationToRecipientsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SendActivityNotificationToRecipientsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SendActivityNotificationToRecipientsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSendActivityNotificationToRecipientsRequestBuilderInternal instantiates a new SendActivityNotificationToRecipientsRequestBuilder and sets the default values.
func NewSendActivityNotificationToRecipientsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SendActivityNotificationToRecipientsRequestBuilder) {
    m := &SendActivityNotificationToRecipientsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/sendActivityNotificationToRecipients", pathParameters),
    }
    return m
}
// NewSendActivityNotificationToRecipientsRequestBuilder instantiates a new SendActivityNotificationToRecipientsRequestBuilder and sets the default values.
func NewSendActivityNotificationToRecipientsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SendActivityNotificationToRecipientsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSendActivityNotificationToRecipientsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send activity feed notifications to multiple users, in bulk.  For more details about sending notifications and the requirements for doing so, seesending Teams activity notifications.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/teamwork-sendactivitynotificationtorecipients?view=graph-rest-1.0
func (m *SendActivityNotificationToRecipientsRequestBuilder) Post(ctx context.Context, body SendActivityNotificationToRecipientsPostRequestBodyable, requestConfiguration *SendActivityNotificationToRecipientsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation send activity feed notifications to multiple users, in bulk.  For more details about sending notifications and the requirements for doing so, seesending Teams activity notifications.
func (m *SendActivityNotificationToRecipientsRequestBuilder) ToPostRequestInformation(ctx context.Context, body SendActivityNotificationToRecipientsPostRequestBodyable, requestConfiguration *SendActivityNotificationToRecipientsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
