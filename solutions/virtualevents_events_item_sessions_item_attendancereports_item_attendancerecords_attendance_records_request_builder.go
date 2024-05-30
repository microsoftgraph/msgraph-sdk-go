package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder provides operations to manage the attendanceRecords property of the microsoft.graph.meetingAttendanceReport entity.
type VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderGetQueryParameters list of attendance records of an attendance report. Read-only.
type VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderGetQueryParameters struct {
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
// VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderGetQueryParameters
}
// VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAttendanceRecordId provides operations to manage the attendanceRecords property of the microsoft.graph.meetingAttendanceReport entity.
// returns a *VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordItemRequestBuilder when successful
func (m *VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) ByAttendanceRecordId(attendanceRecordId string)(*VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if attendanceRecordId != "" {
        urlTplParams["attendanceRecord%2Did"] = attendanceRecordId
    }
    return NewVirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderInternal instantiates a new VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder and sets the default values.
func NewVirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) {
    m := &VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/events/{virtualEvent%2Did}/sessions/{virtualEventSession%2Did}/attendanceReports/{meetingAttendanceReport%2Did}/attendanceRecords{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder instantiates a new VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder and sets the default values.
func NewVirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsCountRequestBuilder when successful
func (m *VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) Count()(*VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsCountRequestBuilder) {
    return NewVirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of attendance records of an attendance report. Read-only.
// returns a AttendanceRecordCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendanceRecordCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAttendanceRecordCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendanceRecordCollectionResponseable), nil
}
// Post create new navigation property to attendanceRecords for solutions
// returns a AttendanceRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendanceRecordable, requestConfiguration *VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendanceRecordable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAttendanceRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendanceRecordable), nil
}
// ToGetRequestInformation list of attendance records of an attendance report. Read-only.
// returns a *RequestInformation when successful
func (m *VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to attendanceRecords for solutions
// returns a *RequestInformation when successful
func (m *VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendanceRecordable, requestConfiguration *VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder when successful
func (m *VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) WithUrl(rawUrl string)(*VirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder) {
    return NewVirtualeventsEventsItemSessionsItemAttendancereportsItemAttendancerecordsAttendanceRecordsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
