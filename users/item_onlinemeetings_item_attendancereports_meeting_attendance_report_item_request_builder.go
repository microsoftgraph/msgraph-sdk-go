package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeetingBase entity.
type ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderGetQueryParameters the attendance reports of an online meeting. Read-only.
type ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderGetQueryParameters
}
// ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AttendanceRecords provides operations to manage the attendanceRecords property of the microsoft.graph.meetingAttendanceReport entity.
// returns a *ItemOnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder when successful
func (m *ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) AttendanceRecords()(*ItemOnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) {
    return NewItemOnlinemeetingsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderInternal instantiates a new ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder and sets the default values.
func NewItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) {
    m := &ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onlineMeetings/{onlineMeeting%2Did}/attendanceReports/{meetingAttendanceReport%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder instantiates a new ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder and sets the default values.
func NewItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property attendanceReports for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable, error) {
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
// Patch update the navigation property attendanceReports in users
// returns a MeetingAttendanceReportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable, requestConfiguration *ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable, error) {
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
// ToDeleteRequestInformation delete navigation property attendanceReports for users
// returns a *RequestInformation when successful
func (m *ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property attendanceReports in users
// returns a *RequestInformation when successful
func (m *ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingAttendanceReportable, requestConfiguration *ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder when successful
func (m *ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) WithUrl(rawUrl string)(*ItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder) {
    return NewItemOnlinemeetingsItemAttendancereportsMeetingAttendanceReportItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
