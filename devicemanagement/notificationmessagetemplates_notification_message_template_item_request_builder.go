package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder provides operations to manage the notificationMessageTemplates property of the microsoft.graph.deviceManagement entity.
type NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderGetQueryParameters read properties and relationships of the notificationMessageTemplate object.
type NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderGetQueryParameters
}
// NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewNotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderInternal instantiates a new NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder and sets the default values.
func NewNotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder) {
    m := &NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/notificationMessageTemplates/{notificationMessageTemplate%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewNotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder instantiates a new NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder and sets the default values.
func NewNotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a notificationMessageTemplate.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-notification-notificationmessagetemplate-delete?view=graph-rest-1.0
func (m *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get read properties and relationships of the notificationMessageTemplate object.
// returns a NotificationMessageTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-notification-notificationmessagetemplate-get?view=graph-rest-1.0
func (m *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NotificationMessageTemplateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateNotificationMessageTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NotificationMessageTemplateable), nil
}
// LocalizedNotificationMessages provides operations to manage the localizedNotificationMessages property of the microsoft.graph.notificationMessageTemplate entity.
// returns a *NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder when successful
func (m *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder) LocalizedNotificationMessages()(*NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder) {
    return NewNotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a notificationMessageTemplate object.
// returns a NotificationMessageTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-notification-notificationmessagetemplate-update?view=graph-rest-1.0
func (m *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NotificationMessageTemplateable, requestConfiguration *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NotificationMessageTemplateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateNotificationMessageTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NotificationMessageTemplateable), nil
}
// SendTestMessage provides operations to call the sendTestMessage method.
// returns a *NotificationmessagetemplatesItemSendtestmessageSendTestMessageRequestBuilder when successful
func (m *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder) SendTestMessage()(*NotificationmessagetemplatesItemSendtestmessageSendTestMessageRequestBuilder) {
    return NewNotificationmessagetemplatesItemSendtestmessageSendTestMessageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes a notificationMessageTemplate.
// returns a *RequestInformation when successful
func (m *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the notificationMessageTemplate object.
// returns a *RequestInformation when successful
func (m *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the properties of a notificationMessageTemplate object.
// returns a *RequestInformation when successful
func (m *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NotificationMessageTemplateable, requestConfiguration *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder when successful
func (m *NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder) WithUrl(rawUrl string)(*NotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder) {
    return NewNotificationmessagetemplatesNotificationMessageTemplateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
