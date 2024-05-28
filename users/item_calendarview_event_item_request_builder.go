package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendarviewEventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.user entity.
type ItemCalendarviewEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarviewEventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Read-only. Nullable.
type ItemCalendarviewEventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string `uriparametername:"endDateTime"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string `uriparametername:"startDateTime"`
}
// ItemCalendarviewEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarviewEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarviewEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarviewItemAcceptRequestBuilder when successful
func (m *ItemCalendarviewEventItemRequestBuilder) Accept()(*ItemCalendarviewItemAcceptRequestBuilder) {
    return NewItemCalendarviewItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarviewItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarviewEventItemRequestBuilder) Attachments()(*ItemCalendarviewItemAttachmentsRequestBuilder) {
    return NewItemCalendarviewItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarviewItemCalendarRequestBuilder when successful
func (m *ItemCalendarviewEventItemRequestBuilder) Calendar()(*ItemCalendarviewItemCalendarRequestBuilder) {
    return NewItemCalendarviewItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarviewItemCancelRequestBuilder when successful
func (m *ItemCalendarviewEventItemRequestBuilder) Cancel()(*ItemCalendarviewItemCancelRequestBuilder) {
    return NewItemCalendarviewItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarviewEventItemRequestBuilderInternal instantiates a new ItemCalendarviewEventItemRequestBuilder and sets the default values.
func NewItemCalendarviewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarviewEventItemRequestBuilder) {
    m := &ItemCalendarviewEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarView/{event%2Did}?endDateTime={endDateTime}&startDateTime={startDateTime}{&%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarviewEventItemRequestBuilder instantiates a new ItemCalendarviewEventItemRequestBuilder and sets the default values.
func NewItemCalendarviewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarviewEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarviewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarviewItemDeclineRequestBuilder when successful
func (m *ItemCalendarviewEventItemRequestBuilder) Decline()(*ItemCalendarviewItemDeclineRequestBuilder) {
    return NewItemCalendarviewItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarviewItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemCalendarviewEventItemRequestBuilder) DismissReminder()(*ItemCalendarviewItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemCalendarviewItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarviewItemExtensionsRequestBuilder when successful
func (m *ItemCalendarviewEventItemRequestBuilder) Extensions()(*ItemCalendarviewItemExtensionsRequestBuilder) {
    return NewItemCalendarviewItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarviewItemForwardRequestBuilder when successful
func (m *ItemCalendarviewEventItemRequestBuilder) Forward()(*ItemCalendarviewItemForwardRequestBuilder) {
    return NewItemCalendarviewItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the calendar view for the calendar. Read-only. Nullable.
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarviewEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarviewEventItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
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
// returns a *ItemCalendarviewItemInstancesRequestBuilder when successful
func (m *ItemCalendarviewEventItemRequestBuilder) Instances()(*ItemCalendarviewItemInstancesRequestBuilder) {
    return NewItemCalendarviewItemInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SnoozeReminder provides operations to call the snoozeReminder method.
// returns a *ItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarviewEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendarviewItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarviewEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemCalendarviewItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the calendar view for the calendar. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemCalendarviewEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarviewEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarviewEventItemRequestBuilder when successful
func (m *ItemCalendarviewEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarviewEventItemRequestBuilder) {
    return NewItemCalendarviewEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
