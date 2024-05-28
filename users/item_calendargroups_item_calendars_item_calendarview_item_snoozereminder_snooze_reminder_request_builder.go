package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder provides operations to call the snoozeReminder method.
type ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilderInternal instantiates a new ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder) {
    m := &ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/snoozeReminder", pathParameters),
    }
    return m
}
// NewItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder instantiates a new ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post postpone a reminder for an event in a user calendar until a new time.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-snoozereminder?view=graph-rest-1.0
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder) Post(ctx context.Context, body ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderPostRequestBodyable, requestConfiguration *ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderPostRequestBodyable, requestConfiguration *ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder) WithUrl(rawUrl string)(*ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
