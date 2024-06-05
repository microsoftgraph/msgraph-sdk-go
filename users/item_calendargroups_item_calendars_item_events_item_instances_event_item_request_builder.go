package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string `uriparametername:"endDateTime"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string `uriparametername:"startDateTime"`
}
// ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) Accept()(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) Attachments()(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAttachmentsRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemInstancesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemCalendarRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) Calendar()(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemCalendarRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemInstancesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemCancelRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) Cancel()(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemCancelRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemInstancesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilderInternal instantiates a new ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) {
    m := &ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/instances/{event%2Did1}?endDateTime={endDateTime}&startDateTime={startDateTime}{&%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder instantiates a new ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemDeclineRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) Decline()(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemDeclineRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemInstancesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) DismissReminder()(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemInstancesItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemExtensionsRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) Extensions()(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemExtensionsRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemInstancesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemForwardRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) Forward()(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemForwardRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemInstancesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable), nil
}
// SnoozeReminder provides operations to call the snoozeReminder method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemInstancesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
