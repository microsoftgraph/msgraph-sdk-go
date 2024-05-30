package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder provides operations to call the sendVirtualAppointmentReminderSms method.
type OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilderInternal instantiates a new OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder and sets the default values.
func NewOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder) {
    m := &OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/onlineMeetings/{onlineMeeting%2Did}/sendVirtualAppointmentReminderSms", pathParameters),
    }
    return m
}
// NewOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder instantiates a new OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder and sets the default values.
func NewOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send an SMS reminder to external attendees for a Teams virtual appointment. This feature requires Teams premium and attendees must have a valid United States phone number to receive SMS notifications.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualappointment-sendvirtualappointmentremindersms?view=graph-rest-1.0
func (m *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder) Post(ctx context.Context, body OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBodyable, requestConfiguration *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation send an SMS reminder to external attendees for a Teams virtual appointment. This feature requires Teams premium and attendees must have a valid United States phone number to receive SMS notifications.
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder) ToPostRequestInformation(ctx context.Context, body OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBodyable, requestConfiguration *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder when successful
func (m *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder) WithUrl(rawUrl string)(*OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder) {
    return NewOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
