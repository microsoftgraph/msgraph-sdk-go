package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// OnlinemeetingsOnlineMeetingItemRequestBuilder provides operations to manage the onlineMeetings property of the microsoft.graph.cloudCommunications entity.
type OnlinemeetingsOnlineMeetingItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlinemeetingsOnlineMeetingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsOnlineMeetingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnlinemeetingsOnlineMeetingItemRequestBuilderGetQueryParameters get onlineMeetings from communications
type OnlinemeetingsOnlineMeetingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnlinemeetingsOnlineMeetingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsOnlineMeetingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnlinemeetingsOnlineMeetingItemRequestBuilderGetQueryParameters
}
// OnlinemeetingsOnlineMeetingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsOnlineMeetingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AttendanceReports provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeetingBase entity.
// returns a *OnlinemeetingsItemAttendancereportsAttendanceReportsRequestBuilder when successful
func (m *OnlinemeetingsOnlineMeetingItemRequestBuilder) AttendanceReports()(*OnlinemeetingsItemAttendancereportsAttendanceReportsRequestBuilder) {
    return NewOnlinemeetingsItemAttendancereportsAttendanceReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AttendeeReport provides operations to manage the media for the cloudCommunications entity.
// returns a *OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder when successful
func (m *OnlinemeetingsOnlineMeetingItemRequestBuilder) AttendeeReport()(*OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder) {
    return NewOnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewOnlinemeetingsOnlineMeetingItemRequestBuilderInternal instantiates a new OnlinemeetingsOnlineMeetingItemRequestBuilder and sets the default values.
func NewOnlinemeetingsOnlineMeetingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsOnlineMeetingItemRequestBuilder) {
    m := &OnlinemeetingsOnlineMeetingItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/onlineMeetings/{onlineMeeting%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewOnlinemeetingsOnlineMeetingItemRequestBuilder instantiates a new OnlinemeetingsOnlineMeetingItemRequestBuilder and sets the default values.
func NewOnlinemeetingsOnlineMeetingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsOnlineMeetingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlinemeetingsOnlineMeetingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property onlineMeetings for communications
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsOnlineMeetingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnlinemeetingsOnlineMeetingItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get onlineMeetings from communications
// returns a OnlineMeetingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsOnlineMeetingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlinemeetingsOnlineMeetingItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnlineMeetingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnlineMeetingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnlineMeetingable), nil
}
// GetVirtualAppointmentJoinWebUrl provides operations to call the getVirtualAppointmentJoinWebUrl method.
// returns a *OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder when successful
func (m *OnlinemeetingsOnlineMeetingItemRequestBuilder) GetVirtualAppointmentJoinWebUrl()(*OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    return NewOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property onlineMeetings in communications
// returns a OnlineMeetingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsOnlineMeetingItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnlineMeetingable, requestConfiguration *OnlinemeetingsOnlineMeetingItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnlineMeetingable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnlineMeetingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnlineMeetingable), nil
}
// Recordings provides operations to manage the recordings property of the microsoft.graph.onlineMeeting entity.
// returns a *OnlinemeetingsItemRecordingsRequestBuilder when successful
func (m *OnlinemeetingsOnlineMeetingItemRequestBuilder) Recordings()(*OnlinemeetingsItemRecordingsRequestBuilder) {
    return NewOnlinemeetingsItemRecordingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendVirtualAppointmentReminderSms provides operations to call the sendVirtualAppointmentReminderSms method.
// returns a *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder when successful
func (m *OnlinemeetingsOnlineMeetingItemRequestBuilder) SendVirtualAppointmentReminderSms()(*OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder) {
    return NewOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendVirtualAppointmentSms provides operations to call the sendVirtualAppointmentSms method.
// returns a *OnlinemeetingsItemSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder when successful
func (m *OnlinemeetingsOnlineMeetingItemRequestBuilder) SendVirtualAppointmentSms()(*OnlinemeetingsItemSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder) {
    return NewOnlinemeetingsItemSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property onlineMeetings for communications
// returns a *RequestInformation when successful
func (m *OnlinemeetingsOnlineMeetingItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OnlinemeetingsOnlineMeetingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get onlineMeetings from communications
// returns a *RequestInformation when successful
func (m *OnlinemeetingsOnlineMeetingItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlinemeetingsOnlineMeetingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property onlineMeetings in communications
// returns a *RequestInformation when successful
func (m *OnlinemeetingsOnlineMeetingItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnlineMeetingable, requestConfiguration *OnlinemeetingsOnlineMeetingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Transcripts provides operations to manage the transcripts property of the microsoft.graph.onlineMeeting entity.
// returns a *OnlinemeetingsItemTranscriptsRequestBuilder when successful
func (m *OnlinemeetingsOnlineMeetingItemRequestBuilder) Transcripts()(*OnlinemeetingsItemTranscriptsRequestBuilder) {
    return NewOnlinemeetingsItemTranscriptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OnlinemeetingsOnlineMeetingItemRequestBuilder when successful
func (m *OnlinemeetingsOnlineMeetingItemRequestBuilder) WithUrl(rawUrl string)(*OnlinemeetingsOnlineMeetingItemRequestBuilder) {
    return NewOnlinemeetingsOnlineMeetingItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
