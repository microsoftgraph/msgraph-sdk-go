package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1d0689ad330d2744c1fb18da8874f5ab0d3d9f9e089b981c1847e0732b28d4b6 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/tentativelyaccept"
    i465e55e32341d082edd97b4cbbe30b74404cd0753d5a94c441f6e442632d344d "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/cancel"
    i53f43a3531f80c6e0571cd6533e13dfb0a9bec5d78a4451ab57d06d69569e313 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/decline"
    i68fc9e57c0f6013ec42006755700065554fc5bf5e29fcd5025a538b57570720b "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/singlevalueextendedproperties"
    i6a090f392ca1e1ce7d5306cbfcc6ecedfebdc0920501b06deec7800a7c0a2a4d "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/extensions"
    i7ebbe5973f1cfdd153c7211a7d729e6acf754214d73692548caaa452fc437931 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/dismissreminder"
    i890bc6825225252197ba0329e5325181337dcc5018722ea3f24133acc97d663d "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/attachments"
    ib37f35045d733a190b9416b6893f854438f6a18b0b560d7fd6225d7b528fc019 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/instances"
    ib4cb01573cfaac49f1987ed83b0866022e257847a6e6a595526d0a91e42cd048 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/snoozereminder"
    ibd0b52fb44db2ad950860bec1e2b4ed68ce5d6349861ec6e194d282bde7bcc51 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/forward"
    ibe852a23e0398849c98cf64aa13ee2a52ee98610fea6b7e377c6a7dd1c1628c0 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/multivalueextendedproperties"
    ie12260a766bb83c9b88880512999d5f33b231a5677bf79c4d07097242d01255a "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar"
    ie711528247ed53f1948a3ebd7b38785aa216df02cdea65cc2a13d9fb60e90e54 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/accept"
    i624e81c2db0b3bfef96d02d14ccc5aafca1e94da30a6ef51f34764348903e10e "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/attachments/item"
    i6d746bffa422f6cab9bce504d82f353fa1e96f30d9135162a2867c0b6ee56409 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/extensions/item"
    i981aed02b0a2dbdb188fef5063ef9909c9da485cac50836f2c3d65d618577a23 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/singlevalueextendedproperties/item"
    ic2fee68681a1462f085ba98b6ba6395da8df913e922d9bff315377a4a254da2e "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/instances/item"
    icbdb2876a9fc3cde1981cefae88c00544c63c09c8176a43b5750b723fee1cfa3 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/multivalueextendedproperties/item"
)

// EventRequestBuilder builds and executes requests for operations under \me\calendarView\{event-id}
type EventRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventRequestBuilderDeleteOptions options for Delete
type EventRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventRequestBuilderGetOptions options for Get
type EventRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EventRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventRequestBuilderGetQueryParameters the calendar view for the calendar. Read-only. Nullable.
type EventRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string;
    // Select properties to be returned
    Select []string;
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string;
}
// EventRequestBuilderPatchOptions options for Patch
type EventRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventRequestBuilder) Accept()(*ie711528247ed53f1948a3ebd7b38785aa216df02cdea65cc2a13d9fb60e90e54.AcceptRequestBuilder) {
    return ie711528247ed53f1948a3ebd7b38785aa216df02cdea65cc2a13d9fb60e90e54.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i890bc6825225252197ba0329e5325181337dcc5018722ea3f24133acc97d663d.AttachmentsRequestBuilder) {
    return i890bc6825225252197ba0329e5325181337dcc5018722ea3f24133acc97d663d.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarView.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*i624e81c2db0b3bfef96d02d14ccc5aafca1e94da30a6ef51f34764348903e10e.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i624e81c2db0b3bfef96d02d14ccc5aafca1e94da30a6ef51f34764348903e10e.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*ie12260a766bb83c9b88880512999d5f33b231a5677bf79c4d07097242d01255a.CalendarRequestBuilder) {
    return ie12260a766bb83c9b88880512999d5f33b231a5677bf79c4d07097242d01255a.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i465e55e32341d082edd97b4cbbe30b74404cd0753d5a94c441f6e442632d344d.CancelRequestBuilder) {
    return i465e55e32341d082edd97b4cbbe30b74404cd0753d5a94c441f6e442632d344d.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarView/{event_id}{?startDateTime,endDateTime,select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventRequestBuilder instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the calendar view for the calendar. Read-only. Nullable.
func (m *EventRequestBuilder) CreateDeleteRequestInformation(options *EventRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the calendar view for the calendar. Read-only. Nullable.
func (m *EventRequestBuilder) CreateGetRequestInformation(options *EventRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the calendar view for the calendar. Read-only. Nullable.
func (m *EventRequestBuilder) CreatePatchRequestInformation(options *EventRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventRequestBuilder) Decline()(*i53f43a3531f80c6e0571cd6533e13dfb0a9bec5d78a4451ab57d06d69569e313.DeclineRequestBuilder) {
    return i53f43a3531f80c6e0571cd6533e13dfb0a9bec5d78a4451ab57d06d69569e313.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the calendar view for the calendar. Read-only. Nullable.
func (m *EventRequestBuilder) Delete(options *EventRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) DismissReminder()(*i7ebbe5973f1cfdd153c7211a7d729e6acf754214d73692548caaa452fc437931.DismissReminderRequestBuilder) {
    return i7ebbe5973f1cfdd153c7211a7d729e6acf754214d73692548caaa452fc437931.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i6a090f392ca1e1ce7d5306cbfcc6ecedfebdc0920501b06deec7800a7c0a2a4d.ExtensionsRequestBuilder) {
    return i6a090f392ca1e1ce7d5306cbfcc6ecedfebdc0920501b06deec7800a7c0a2a4d.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarView.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*i6d746bffa422f6cab9bce504d82f353fa1e96f30d9135162a2867c0b6ee56409.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i6d746bffa422f6cab9bce504d82f353fa1e96f30d9135162a2867c0b6ee56409.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*ibd0b52fb44db2ad950860bec1e2b4ed68ce5d6349861ec6e194d282bde7bcc51.ForwardRequestBuilder) {
    return ibd0b52fb44db2ad950860bec1e2b4ed68ce5d6349861ec6e194d282bde7bcc51.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Read-only. Nullable.
func (m *EventRequestBuilder) Get(options *EventRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEvent() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event), nil
}
func (m *EventRequestBuilder) Instances()(*ib37f35045d733a190b9416b6893f854438f6a18b0b560d7fd6225d7b528fc019.InstancesRequestBuilder) {
    return ib37f35045d733a190b9416b6893f854438f6a18b0b560d7fd6225d7b528fc019.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarView.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*ic2fee68681a1462f085ba98b6ba6395da8df913e922d9bff315377a4a254da2e.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ic2fee68681a1462f085ba98b6ba6395da8df913e922d9bff315377a4a254da2e.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*ibe852a23e0398849c98cf64aa13ee2a52ee98610fea6b7e377c6a7dd1c1628c0.MultiValueExtendedPropertiesRequestBuilder) {
    return ibe852a23e0398849c98cf64aa13ee2a52ee98610fea6b7e377c6a7dd1c1628c0.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*icbdb2876a9fc3cde1981cefae88c00544c63c09c8176a43b5750b723fee1cfa3.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return icbdb2876a9fc3cde1981cefae88c00544c63c09c8176a43b5750b723fee1cfa3.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the calendar view for the calendar. Read-only. Nullable.
func (m *EventRequestBuilder) Patch(options *EventRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i68fc9e57c0f6013ec42006755700065554fc5bf5e29fcd5025a538b57570720b.SingleValueExtendedPropertiesRequestBuilder) {
    return i68fc9e57c0f6013ec42006755700065554fc5bf5e29fcd5025a538b57570720b.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i981aed02b0a2dbdb188fef5063ef9909c9da485cac50836f2c3d65d618577a23.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i981aed02b0a2dbdb188fef5063ef9909c9da485cac50836f2c3d65d618577a23.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*ib4cb01573cfaac49f1987ed83b0866022e257847a6e6a595526d0a91e42cd048.SnoozeReminderRequestBuilder) {
    return ib4cb01573cfaac49f1987ed83b0866022e257847a6e6a595526d0a91e42cd048.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i1d0689ad330d2744c1fb18da8874f5ab0d3d9f9e089b981c1847e0732b28d4b6.TentativelyAcceptRequestBuilder) {
    return i1d0689ad330d2744c1fb18da8874f5ab0d3d9f9e089b981c1847e0732b28d4b6.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
