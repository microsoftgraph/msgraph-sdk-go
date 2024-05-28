package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnlinemeetingsOnlineMeetingItemRequestBuilder provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
type ItemOnlinemeetingsOnlineMeetingItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnlinemeetingsOnlineMeetingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlinemeetingsOnlineMeetingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOnlinemeetingsOnlineMeetingItemRequestBuilderGetQueryParameters information about a meeting, including the URL used to join a meeting, the attendees list, and the description.
type ItemOnlinemeetingsOnlineMeetingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOnlinemeetingsOnlineMeetingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlinemeetingsOnlineMeetingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnlinemeetingsOnlineMeetingItemRequestBuilderGetQueryParameters
}
// ItemOnlinemeetingsOnlineMeetingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlinemeetingsOnlineMeetingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AttendanceReports provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeetingBase entity.
// returns a *ItemOnlinemeetingsItemAttendancereportsAttendanceReportsRequestBuilder when successful
func (m *ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) AttendanceReports()(*ItemOnlinemeetingsItemAttendancereportsAttendanceReportsRequestBuilder) {
    return NewItemOnlinemeetingsItemAttendancereportsAttendanceReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AttendeeReport provides operations to manage the media for the user entity.
// returns a *ItemOnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder when successful
func (m *ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) AttendeeReport()(*ItemOnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder) {
    return NewItemOnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOnlinemeetingsOnlineMeetingItemRequestBuilderInternal instantiates a new ItemOnlinemeetingsOnlineMeetingItemRequestBuilder and sets the default values.
func NewItemOnlinemeetingsOnlineMeetingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) {
    m := &ItemOnlinemeetingsOnlineMeetingItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onlineMeetings/{onlineMeeting%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemOnlinemeetingsOnlineMeetingItemRequestBuilder instantiates a new ItemOnlinemeetingsOnlineMeetingItemRequestBuilder and sets the default values.
func NewItemOnlinemeetingsOnlineMeetingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlinemeetingsOnlineMeetingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property onlineMeetings for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemOnlinemeetingsOnlineMeetingItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get information about a meeting, including the URL used to join a meeting, the attendees list, and the description.
// returns a OnlineMeetingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnlinemeetingsOnlineMeetingItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnlineMeetingable, error) {
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
// returns a *ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder when successful
func (m *ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) GetVirtualAppointmentJoinWebUrl()(*ItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    return NewItemOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property onlineMeetings in users
// returns a OnlineMeetingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnlineMeetingable, requestConfiguration *ItemOnlinemeetingsOnlineMeetingItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnlineMeetingable, error) {
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
// returns a *ItemOnlinemeetingsItemRecordingsRequestBuilder when successful
func (m *ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) Recordings()(*ItemOnlinemeetingsItemRecordingsRequestBuilder) {
    return NewItemOnlinemeetingsItemRecordingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendVirtualAppointmentReminderSms provides operations to call the sendVirtualAppointmentReminderSms method.
// returns a *ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder when successful
func (m *ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) SendVirtualAppointmentReminderSms()(*ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder) {
    return NewItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendVirtualAppointmentSms provides operations to call the sendVirtualAppointmentSms method.
// returns a *ItemOnlinemeetingsItemSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder when successful
func (m *ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) SendVirtualAppointmentSms()(*ItemOnlinemeetingsItemSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder) {
    return NewItemOnlinemeetingsItemSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property onlineMeetings for users
// returns a *RequestInformation when successful
func (m *ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemOnlinemeetingsOnlineMeetingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation information about a meeting, including the URL used to join a meeting, the attendees list, and the description.
// returns a *RequestInformation when successful
func (m *ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnlinemeetingsOnlineMeetingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property onlineMeetings in users
// returns a *RequestInformation when successful
func (m *ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnlineMeetingable, requestConfiguration *ItemOnlinemeetingsOnlineMeetingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOnlinemeetingsItemTranscriptsRequestBuilder when successful
func (m *ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) Transcripts()(*ItemOnlinemeetingsItemTranscriptsRequestBuilder) {
    return NewItemOnlinemeetingsItemTranscriptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOnlinemeetingsOnlineMeetingItemRequestBuilder when successful
func (m *ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) WithUrl(rawUrl string)(*ItemOnlinemeetingsOnlineMeetingItemRequestBuilder) {
    return NewItemOnlinemeetingsOnlineMeetingItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
