package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1233ad0ba006ab96918402fc9449bdbea95cc2732eac86365bde41c84b357d21 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/singlevalueextendedproperties"
    i133098166421eafcbb7d41e4a68f3d29757bcf6688e847c4a7808bf543993aef "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/tentativelyaccept"
    i2ef9251ae81121975f3c9a073ec23b2469d36ed1973d4d03af6e70a2f6fc3c4d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/multivalueextendedproperties"
    i32caceec19579e592d7312d6a514ad4bfa9807198b6c2414606a3703116e102f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/snoozereminder"
    i42df08a14e97e22bb26eddf9cbaa35aa7bf02405a8bbc84fe4fe5b4203cc2ff8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances"
    i53ba2fe10fcac93a89e1ab826cf6d7b8e1e1069ce7a61b20bc3ce310e26333ab "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/accept"
    i6a59555a6a577b7a15b03f87f857fa23cfe6af2eace90222104dbcdc30a0d24d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/attachments"
    i6bf2dbc96674e29a2505ab70f10b8f1735eb7fe1b7640bf2d65e16312c0a071d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/calendar"
    i8c3a0045c4c0ac4da728d4aa37dbba4e6df99e52c194414aedb749b76e669fca "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/forward"
    i921d4dffa29707e7b15bb069fe546aa8d25d138f815d39cdd30ed48b40c27465 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/extensions"
    iac711011c7c99048797a7e80afa2fb7cb707667348b6c9c773d351f73e0a403f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/cancel"
    ie6ebec3d6348227c40f20857d0105bccb0382669c86a08c61932f7f04269ae7a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/dismissreminder"
    ie8cb6c49b7e9e2ac0ad0271442195855cb47449992c78270fe676de14c335081 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/decline"
    i00430d12f816c0c6997639410f009cc5f413c588c399506a408cafb1d4a794ec "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/multivalueextendedproperties/item"
    i1ec50c12270bd7bd6ab657694752ebec839c5bc0a607be433717875d65bff30e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/extensions/item"
    i9be5b8dcebdfca6e7ace3db98d602d80c97a8f974e4125c75bacc246c066e21e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/singlevalueextendedproperties/item"
    i9fe44f8859760d86bd88baff7a59fc2a46f6264caa52b2adf30c8d7f4ae24672 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/attachments/item"
    id60fc1ecd13479f120a7edbab2577fd0b8fff9ebc834f6c7d280345dd5f73d2e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item"
)

// Builds and executes requests for operations under \groups\{group-id}\calendar\calendarView\{event-id}
type EventRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type EventRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
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
// The calendar view for the calendar. Navigation property. Read-only.
type EventRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string;
    // Select properties to be returned
    Select_escaped []string;
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string;
}
// Options for Patch
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
func (m *EventRequestBuilder) Accept()(*i53ba2fe10fcac93a89e1ab826cf6d7b8e1e1069ce7a61b20bc3ce310e26333ab.AcceptRequestBuilder) {
    return i53ba2fe10fcac93a89e1ab826cf6d7b8e1e1069ce7a61b20bc3ce310e26333ab.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i6a59555a6a577b7a15b03f87f857fa23cfe6af2eace90222104dbcdc30a0d24d.AttachmentsRequestBuilder) {
    return i6a59555a6a577b7a15b03f87f857fa23cfe6af2eace90222104dbcdc30a0d24d.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.calendarView.item.attachments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) AttachmentsById(id string)(*i9fe44f8859760d86bd88baff7a59fc2a46f6264caa52b2adf30c8d7f4ae24672.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i9fe44f8859760d86bd88baff7a59fc2a46f6264caa52b2adf30c8d7f4ae24672.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i6bf2dbc96674e29a2505ab70f10b8f1735eb7fe1b7640bf2d65e16312c0a071d.CalendarRequestBuilder) {
    return i6bf2dbc96674e29a2505ab70f10b8f1735eb7fe1b7640bf2d65e16312c0a071d.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*iac711011c7c99048797a7e80afa2fb7cb707667348b6c9c773d351f73e0a403f.CancelRequestBuilder) {
    return iac711011c7c99048797a7e80afa2fb7cb707667348b6c9c773d351f73e0a403f.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/calendar/calendarView/{event_id}{?startDateTime,endDateTime,select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventRequestBuilderInternal(urlParams, requestAdapter)
}
// The calendar view for the calendar. Navigation property. Read-only.
// Parameters:
//  - options : Options for the request
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
// The calendar view for the calendar. Navigation property. Read-only.
// Parameters:
//  - options : Options for the request
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
// The calendar view for the calendar. Navigation property. Read-only.
// Parameters:
//  - options : Options for the request
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
func (m *EventRequestBuilder) Decline()(*ie8cb6c49b7e9e2ac0ad0271442195855cb47449992c78270fe676de14c335081.DeclineRequestBuilder) {
    return ie8cb6c49b7e9e2ac0ad0271442195855cb47449992c78270fe676de14c335081.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The calendar view for the calendar. Navigation property. Read-only.
// Parameters:
//  - options : Options for the request
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
func (m *EventRequestBuilder) DismissReminder()(*ie6ebec3d6348227c40f20857d0105bccb0382669c86a08c61932f7f04269ae7a.DismissReminderRequestBuilder) {
    return ie6ebec3d6348227c40f20857d0105bccb0382669c86a08c61932f7f04269ae7a.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i921d4dffa29707e7b15bb069fe546aa8d25d138f815d39cdd30ed48b40c27465.ExtensionsRequestBuilder) {
    return i921d4dffa29707e7b15bb069fe546aa8d25d138f815d39cdd30ed48b40c27465.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.calendarView.item.extensions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) ExtensionsById(id string)(*i1ec50c12270bd7bd6ab657694752ebec839c5bc0a607be433717875d65bff30e.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i1ec50c12270bd7bd6ab657694752ebec839c5bc0a607be433717875d65bff30e.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i8c3a0045c4c0ac4da728d4aa37dbba4e6df99e52c194414aedb749b76e669fca.ForwardRequestBuilder) {
    return i8c3a0045c4c0ac4da728d4aa37dbba4e6df99e52c194414aedb749b76e669fca.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The calendar view for the calendar. Navigation property. Read-only.
// Parameters:
//  - options : Options for the request
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
func (m *EventRequestBuilder) Instances()(*i42df08a14e97e22bb26eddf9cbaa35aa7bf02405a8bbc84fe4fe5b4203cc2ff8.InstancesRequestBuilder) {
    return i42df08a14e97e22bb26eddf9cbaa35aa7bf02405a8bbc84fe4fe5b4203cc2ff8.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.calendarView.item.instances.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) InstancesById(id string)(*id60fc1ecd13479f120a7edbab2577fd0b8fff9ebc834f6c7d280345dd5f73d2e.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return id60fc1ecd13479f120a7edbab2577fd0b8fff9ebc834f6c7d280345dd5f73d2e.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i2ef9251ae81121975f3c9a073ec23b2469d36ed1973d4d03af6e70a2f6fc3c4d.MultiValueExtendedPropertiesRequestBuilder) {
    return i2ef9251ae81121975f3c9a073ec23b2469d36ed1973d4d03af6e70a2f6fc3c4d.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.calendarView.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i00430d12f816c0c6997639410f009cc5f413c588c399506a408cafb1d4a794ec.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i00430d12f816c0c6997639410f009cc5f413c588c399506a408cafb1d4a794ec.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The calendar view for the calendar. Navigation property. Read-only.
// Parameters:
//  - options : Options for the request
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i1233ad0ba006ab96918402fc9449bdbea95cc2732eac86365bde41c84b357d21.SingleValueExtendedPropertiesRequestBuilder) {
    return i1233ad0ba006ab96918402fc9449bdbea95cc2732eac86365bde41c84b357d21.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.calendarView.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i9be5b8dcebdfca6e7ace3db98d602d80c97a8f974e4125c75bacc246c066e21e.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i9be5b8dcebdfca6e7ace3db98d602d80c97a8f974e4125c75bacc246c066e21e.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i32caceec19579e592d7312d6a514ad4bfa9807198b6c2414606a3703116e102f.SnoozeReminderRequestBuilder) {
    return i32caceec19579e592d7312d6a514ad4bfa9807198b6c2414606a3703116e102f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i133098166421eafcbb7d41e4a68f3d29757bcf6688e847c4a7808bf543993aef.TentativelyAcceptRequestBuilder) {
    return i133098166421eafcbb7d41e4a68f3d29757bcf6688e847c4a7808bf543993aef.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
