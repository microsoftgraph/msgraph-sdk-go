package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeetingBase entity.
type VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderGetQueryParameters get a list of meetingAttendanceReport objects for an onlineMeeting or a virtualEvent. Each time an online meeting or a virtual event ends, an attendance report is generated for that session.
type VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderGetQueryParameters
}
// VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMeetingAttendanceReportId provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeetingBase entity.
// returns a *VirtualeventsWebinarsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder when successful
func (m *VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) ByMeetingAttendanceReportId(meetingAttendanceReportId string)(*VirtualeventsWebinarsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if meetingAttendanceReportId != "" {
        urlTplParams["meetingAttendanceReport%2Did"] = meetingAttendanceReportId
    }
    return NewVirtualeventsWebinarsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderInternal instantiates a new VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder and sets the default values.
func NewVirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) {
    m := &VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/{virtualEventWebinar%2Did}/sessions/{virtualEventSession%2Did}/attendanceReports{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder instantiates a new VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder and sets the default values.
func NewVirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualeventsWebinarsItemSessionsItemAttendancereportsCountRequestBuilder when successful
func (m *VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) Count()(*VirtualeventsWebinarsItemSessionsItemAttendancereportsCountRequestBuilder) {
    return NewVirtualeventsWebinarsItemSessionsItemAttendancereportsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of meetingAttendanceReport objects for an onlineMeeting or a virtualEvent. Each time an online meeting or a virtual event ends, an attendance report is generated for that session.
// returns a MeetingAttendanceReportCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/meetingattendancereport-list?view=graph-rest-1.0
func (m *VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMeetingAttendanceReportCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportCollectionResponseable), nil
}
// Post create new navigation property to attendanceReports for solutions
// returns a MeetingAttendanceReportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable, requestConfiguration *VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMeetingAttendanceReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable), nil
}
// ToGetRequestInformation get a list of meetingAttendanceReport objects for an onlineMeeting or a virtualEvent. Each time an online meeting or a virtual event ends, an attendance report is generated for that session.
// returns a *RequestInformation when successful
func (m *VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to attendanceReports for solutions
// returns a *RequestInformation when successful
func (m *VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable, requestConfiguration *VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder when successful
func (m *VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) WithUrl(rawUrl string)(*VirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder) {
    return NewVirtualeventsWebinarsItemSessionsItemAttendancereportsAttendanceReportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
