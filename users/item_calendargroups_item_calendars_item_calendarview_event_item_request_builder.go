package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
type ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string `uriparametername:"endDateTime"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string `uriparametername:"startDateTime"`
}
// ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemAcceptRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) Accept()(*ItemCalendargroupsItemCalendarsItemCalendarviewItemAcceptRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) Attachments()(*ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemCalendarRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) Calendar()(*ItemCalendargroupsItemCalendarsItemCalendarviewItemCalendarRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemCancelRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) Cancel()(*ItemCalendargroupsItemCalendarsItemCalendarviewItemCancelRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilderInternal instantiates a new ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) {
    m := &ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}?endDateTime={endDateTime}&startDateTime={startDateTime}{&%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder instantiates a new ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemDeclineRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) Decline()(*ItemCalendargroupsItemCalendarsItemCalendarviewItemDeclineRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) DismissReminder()(*ItemCalendargroupsItemCalendarsItemCalendarviewItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemExtensionsRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) Extensions()(*ItemCalendargroupsItemCalendarsItemCalendarviewItemExtensionsRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemForwardRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) Forward()(*ItemCalendargroupsItemCalendarsItemCalendarviewItemForwardRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the calendar view for the calendar. Navigation property. Read-only.
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
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
// Instances provides operations to manage the instances property of the microsoft.graph.event entity.
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) Instances()(*ItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SnoozeReminder provides operations to call the snoozeReminder method.
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) SnoozeReminder()(*ItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) TentativelyAccept()(*ItemCalendargroupsItemCalendarsItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
// returns a *RequestInformation when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
