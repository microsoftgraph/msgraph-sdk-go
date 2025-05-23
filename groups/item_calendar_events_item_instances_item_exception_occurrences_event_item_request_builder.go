// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters contains the id property values of the event instances that are exceptions in a recurring series.Exceptions can differ from other occurrences in a recurring series, such as the subject, start or end times, or attendees. Exceptions don't include canceled occurrences.Returned only on $select and $expand in a GET operation that specifies the ID (seriesMasterId property value) of a series master event.
type ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder when successful
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Accept()(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder) {
    return NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder when successful
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemCancelRequestBuilder when successful
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Cancel()(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemCancelRequestBuilder) {
    return NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/calendar/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder instantiates a new ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilder when successful
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Decline()(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilder) {
    return NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilder when successful
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) DismissReminder()(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilder) {
    return NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilder when successful
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemForwardRequestBuilder when successful
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Forward()(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemForwardRequestBuilder) {
    return NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get contains the id property values of the event instances that are exceptions in a recurring series.Exceptions can differ from other occurrences in a recurring series, such as the subject, start or end times, or attendees. Exceptions don't include canceled occurrences.Returned only on $select and $expand in a GET operation that specifies the ID (seriesMasterId property value) of a series master event.
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
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
// PermanentDelete provides operations to call the permanentDelete method.
// returns a *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemPermanentDeleteRequestBuilder when successful
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) PermanentDelete()(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemPermanentDeleteRequestBuilder) {
    return NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemPermanentDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SnoozeReminder provides operations to call the snoozeReminder method.
// returns a *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    return NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder) {
    return NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation contains the id property values of the event instances that are exceptions in a recurring series.Exceptions can differ from other occurrences in a recurring series, such as the subject, start or end times, or attendees. Exceptions don't include canceled occurrences.Returned only on $select and $expand in a GET operation that specifies the ID (seriesMasterId property value) of a series master event.
// returns a *RequestInformation when successful
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder when successful
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    return NewItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
