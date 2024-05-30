package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder provides operations to call the snoozeReminder method.
type ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderInternal instantiates a new ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) {
    m := &ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/snoozeReminder", pathParameters),
    }
    return m
}
// NewItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder instantiates a new ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post postpone a reminder for an event in a user calendar until a new time.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-snoozereminder?view=graph-rest-1.0
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) Post(ctx context.Context, body ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderPostRequestBodyable, requestConfiguration *ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation postpone a reminder for an event in a user calendar until a new time.
// returns a *RequestInformation when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderPostRequestBodyable, requestConfiguration *ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) WithUrl(rawUrl string)(*ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
