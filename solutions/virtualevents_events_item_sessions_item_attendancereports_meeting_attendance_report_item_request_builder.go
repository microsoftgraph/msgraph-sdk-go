package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeetingBase entity.
type VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderGetQueryParameters the attendance reports of an online meeting. Read-only.
type VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderGetQueryParameters
}
// VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AttendanceRecords provides operations to manage the attendanceRecords property of the microsoft.graph.meetingAttendanceReport entity.
// returns a *VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder when successful
func (m *VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) AttendanceRecords()(*VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) {
    return NewVirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderInternal instantiates a new VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder and sets the default values.
func NewVirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) {
    m := &VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/events/{virtualEvent%2Did}/sessions/{virtualEventSession%2Did}/attendanceReports/{meetingAttendanceReport%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder instantiates a new VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder and sets the default values.
func NewVirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property attendanceReports for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the attendance reports of an online meeting. Read-only.
// returns a MeetingAttendanceReportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property attendanceReports in solutions
// returns a MeetingAttendanceReportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable, requestConfiguration *VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property attendanceReports for solutions
// returns a *RequestInformation when successful
func (m *VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the attendance reports of an online meeting. Read-only.
// returns a *RequestInformation when successful
func (m *VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property attendanceReports in solutions
// returns a *RequestInformation when successful
func (m *VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable, requestConfiguration *VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder when successful
func (m *VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) WithUrl(rawUrl string)(*VirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) {
    return NewVirtualeventsEventsItemSessionsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
