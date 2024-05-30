package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder provides operations to call the dismissReminder method.
type ItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilderInternal instantiates a new ItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder and sets the default values.
func NewItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder) {
    m := &ItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/dismissReminder", pathParameters),
    }
    return m
}
// NewItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder instantiates a new ItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder and sets the default values.
func NewItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post dismiss a reminder that has been triggered for an event in a user calendar.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-dismissreminder?view=graph-rest-1.0
func (m *ItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation dismiss a reminder that has been triggered for an event in a user calendar.
// returns a *RequestInformation when successful
func (m *ItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
