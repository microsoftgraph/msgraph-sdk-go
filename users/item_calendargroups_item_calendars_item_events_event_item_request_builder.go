package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.calendar entity.
type ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderGetQueryParameters the events in the calendar. Navigation property. Read-only.
type ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderGetQueryParameters
}
// ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemAcceptRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) Accept()(*ItemCalendargroupsItemCalendarsItemEventsItemAcceptRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemAttachmentsRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) Attachments()(*ItemCalendargroupsItemCalendarsItemEventsItemAttachmentsRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemCalendarRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) Calendar()(*ItemCalendargroupsItemCalendarsItemEventsItemCalendarRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemCancelRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) Cancel()(*ItemCalendargroupsItemCalendarsItemEventsItemCancelRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderInternal instantiates a new ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) {
    m := &ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder instantiates a new ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemDeclineRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) Decline()(*ItemCalendargroupsItemCalendarsItemEventsItemDeclineRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property events for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) DismissReminder()(*ItemCalendargroupsItemCalendarsItemEventsItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemExtensionsRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) Extensions()(*ItemCalendargroupsItemCalendarsItemEventsItemExtensionsRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemForwardRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) Forward()(*ItemCalendargroupsItemCalendarsItemEventsItemForwardRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the events in the calendar. Navigation property. Read-only.
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
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
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemInstancesRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) Instances()(*ItemCalendargroupsItemCalendarsItemEventsItemInstancesRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property events in users
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, requestConfiguration *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) SnoozeReminder()(*ItemCalendargroupsItemCalendarsItemEventsItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) TentativelyAccept()(*ItemCalendargroupsItemCalendarsItemEventsItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property events for users
// returns a *RequestInformation when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the events in the calendar. Navigation property. Read-only.
// returns a *RequestInformation when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property events in users
// returns a *RequestInformation when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, requestConfiguration *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
