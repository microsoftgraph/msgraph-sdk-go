package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder provides operations to call the snoozeReminder method.
type ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderInternal instantiates a new ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder and sets the default values.
func NewItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) {
    m := &ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/snoozeReminder", pathParameters),
    }
    return m
}
// NewItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder instantiates a new ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder and sets the default values.
func NewItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post postpone a reminder for an event in a user calendar until a new time.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-snoozereminder?view=graph-rest-1.0
func (m *ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) Post(ctx context.Context, body ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderPostRequestBodyable, requestConfiguration *ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderPostRequestBodyable, requestConfiguration *ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
