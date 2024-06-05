package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendarsItemCalendarviewEventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
type ItemCalendarsItemCalendarviewEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarsItemCalendarviewEventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type ItemCalendarsItemCalendarviewEventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string `uriparametername:"endDateTime"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string `uriparametername:"startDateTime"`
}
// ItemCalendarsItemCalendarviewEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemCalendarviewEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarsItemCalendarviewEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarsItemCalendarviewItemAcceptRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewEventItemRequestBuilder) Accept()(*ItemCalendarsItemCalendarviewItemAcceptRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemCalendarviewItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewEventItemRequestBuilder) Attachments()(*ItemCalendarsItemCalendarviewItemAttachmentsRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemCalendarviewItemCalendarRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewEventItemRequestBuilder) Calendar()(*ItemCalendarsItemCalendarviewItemCalendarRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarsItemCalendarviewItemCancelRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewEventItemRequestBuilder) Cancel()(*ItemCalendarsItemCalendarviewItemCancelRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarsItemCalendarviewEventItemRequestBuilderInternal instantiates a new ItemCalendarsItemCalendarviewEventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarviewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemCalendarviewEventItemRequestBuilder) {
    m := &ItemCalendarsItemCalendarviewEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}?endDateTime={endDateTime}&startDateTime={startDateTime}{&%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarsItemCalendarviewEventItemRequestBuilder instantiates a new ItemCalendarsItemCalendarviewEventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarviewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemCalendarviewEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarsItemCalendarviewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarsItemCalendarviewItemDeclineRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewEventItemRequestBuilder) Decline()(*ItemCalendarsItemCalendarviewItemDeclineRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarsItemCalendarviewItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewEventItemRequestBuilder) DismissReminder()(*ItemCalendarsItemCalendarviewItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemCalendarviewItemExtensionsRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewEventItemRequestBuilder) Extensions()(*ItemCalendarsItemCalendarviewItemExtensionsRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarsItemCalendarviewItemForwardRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewEventItemRequestBuilder) Forward()(*ItemCalendarsItemCalendarviewItemForwardRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the calendar view for the calendar. Navigation property. Read-only.
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarsItemCalendarviewEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarsItemCalendarviewEventItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
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
// returns a *ItemCalendarsItemCalendarviewItemInstancesRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewEventItemRequestBuilder) Instances()(*ItemCalendarsItemCalendarviewItemInstancesRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SnoozeReminder provides operations to call the snoozeReminder method.
// returns a *ItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarsItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarsItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
// returns a *RequestInformation when successful
func (m *ItemCalendarsItemCalendarviewEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarsItemCalendarviewEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarsItemCalendarviewEventItemRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarsItemCalendarviewEventItemRequestBuilder) {
    return NewItemCalendarsItemCalendarviewEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
