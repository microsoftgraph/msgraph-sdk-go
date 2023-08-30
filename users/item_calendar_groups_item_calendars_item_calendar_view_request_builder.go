package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
type ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilderGetQueryParameters get the occurrences, exceptions and single instances of events in a calendar view defined by a time range,from a user's default calendar (../me/calendarView) or some other calendar of the user's.
type ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilderGetQueryParameters
}
// ByEventId provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilder) ByEventId(eventId string)(*ItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if eventId != "" {
        urlTplParams["event%2Did"] = eventId
    }
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilderInternal instantiates a new CalendarViewRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilder) {
    m := &ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilder instantiates a new CalendarViewRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilder) Count()(*ItemCalendarGroupsItemCalendarsItemCalendarViewCountRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delta provides operations to call the delta method.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilder) Delta()(*ItemCalendarGroupsItemCalendarsItemCalendarViewDeltaRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the occurrences, exceptions and single instances of events in a calendar view defined by a time range,from a user's default calendar (../me/calendarView) or some other calendar of the user's.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/calendar-list-calendarview?view=graph-rest-1.0
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EventCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEventCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EventCollectionResponseable), nil
}
// ToGetRequestInformation get the occurrences, exceptions and single instances of events in a calendar view defined by a time range,from a user's default calendar (../me/calendarView) or some other calendar of the user's.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
