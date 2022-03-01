package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i4ff5990c26216bf96fa13edbc712ebf52c93877c789c04d00f4bbc45a8610b7e "github.com/microsoftgraph/msgraph-sdk-go/me/onlinemeetings/item/attendancereports/item/attendancerecords"
    i175613690773dafc36090c5a5b9d10ac791e0323bee00615b1128c23db26645b "github.com/microsoftgraph/msgraph-sdk-go/me/onlinemeetings/item/attendancereports/item/attendancerecords/item"
)

// MeetingAttendanceReportItemRequestBuilder builds and executes requests for operations under \me\onlineMeetings\{onlineMeeting-id}\attendanceReports\{meetingAttendanceReport-id}
type MeetingAttendanceReportItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MeetingAttendanceReportItemRequestBuilderDeleteOptions options for Delete
type MeetingAttendanceReportItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MeetingAttendanceReportItemRequestBuilderGetOptions options for Get
type MeetingAttendanceReportItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MeetingAttendanceReportItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MeetingAttendanceReportItemRequestBuilderGetQueryParameters the attendance reports of an online meeting. Read-only.
type MeetingAttendanceReportItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// MeetingAttendanceReportItemRequestBuilderPatchOptions options for Patch
type MeetingAttendanceReportItemRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MeetingAttendanceReport;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MeetingAttendanceReportItemRequestBuilder) AttendanceRecords()(*i4ff5990c26216bf96fa13edbc712ebf52c93877c789c04d00f4bbc45a8610b7e.AttendanceRecordsRequestBuilder) {
    return i4ff5990c26216bf96fa13edbc712ebf52c93877c789c04d00f4bbc45a8610b7e.NewAttendanceRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceRecordsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.onlineMeetings.item.attendanceReports.item.attendanceRecords.item collection
func (m *MeetingAttendanceReportItemRequestBuilder) AttendanceRecordsById(id string)(*i175613690773dafc36090c5a5b9d10ac791e0323bee00615b1128c23db26645b.AttendanceRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attendanceRecord_id"] = id
    }
    return i175613690773dafc36090c5a5b9d10ac791e0323bee00615b1128c23db26645b.NewAttendanceRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMeetingAttendanceReportItemRequestBuilderInternal instantiates a new MeetingAttendanceReportItemRequestBuilder and sets the default values.
func NewMeetingAttendanceReportItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MeetingAttendanceReportItemRequestBuilder) {
    m := &MeetingAttendanceReportItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onlineMeetings/{onlineMeeting_id}/attendanceReports/{meetingAttendanceReport_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeetingAttendanceReportItemRequestBuilder instantiates a new MeetingAttendanceReportItemRequestBuilder and sets the default values.
func NewMeetingAttendanceReportItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MeetingAttendanceReportItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeetingAttendanceReportItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the attendance reports of an online meeting. Read-only.
func (m *MeetingAttendanceReportItemRequestBuilder) CreateDeleteRequestInformation(options *MeetingAttendanceReportItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the attendance reports of an online meeting. Read-only.
func (m *MeetingAttendanceReportItemRequestBuilder) CreateGetRequestInformation(options *MeetingAttendanceReportItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation the attendance reports of an online meeting. Read-only.
func (m *MeetingAttendanceReportItemRequestBuilder) CreatePatchRequestInformation(options *MeetingAttendanceReportItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete the attendance reports of an online meeting. Read-only.
func (m *MeetingAttendanceReportItemRequestBuilder) Delete(options *MeetingAttendanceReportItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the attendance reports of an online meeting. Read-only.
func (m *MeetingAttendanceReportItemRequestBuilder) Get(options *MeetingAttendanceReportItemRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MeetingAttendanceReport, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewMeetingAttendanceReport() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MeetingAttendanceReport), nil
}
// Patch the attendance reports of an online meeting. Read-only.
func (m *MeetingAttendanceReportItemRequestBuilder) Patch(options *MeetingAttendanceReportItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
