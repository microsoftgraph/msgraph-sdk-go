package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i083a9367fb14f0da75caea7bc857922fe10d91693dde0f11517ab6420e557018 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/attachments"
    i0f781326a67fbc9776be4d1ccbf43639b9bf8077faa925a7b0f90f0e03141c07 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/snoozereminder"
    i10a043ae8449d6b8061165bb67946e604bc2614783ade4a4557a2bb8d2a5a2b0 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/decline"
    i339d2d1a30983a01dd7de00efdd9e7d51cf0424dd9c4c96d335f9901fcff9763 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances"
    i49ad5e4dbfe1e924deb723580385a06e87f447b9ed73e648b64a934117f1c784 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/forward"
    i62d05692265cf61165d6bb2a1bb3a36d03ea2604586e3ebd550fae433333488e "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/singlevalueextendedproperties"
    i69bb0e8065fcc365cd51c4f9cb100c9289f8c378fe585dd0126d7eaee8953622 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/calendar"
    i87334fbab17098d26d733b7c6d4d979d97f28091d5d1b3c41c2e4adcd9d2f2a3 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/accept"
    id1d8fbd767f32fa77bb0a1e1f8b79c51f1a58b7dae1ccebcd3a3001af55d4f44 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/extensions"
    ie04ca0f3f9823daf845f8b3b165c34feaffc35ed4189263a607834108a606dbe "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/cancel"
    ie6641980e1225311bdf1c499d054bc5d07aa799f49902ef90e71b01800c2a555 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/dismissreminder"
    iee14c2abd2614ebab15ceaf1135456d72c4e1cd953dd5d7b3347f819c29c6002 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/tentativelyaccept"
    if78feb360a058ff87bf5a02baba6f7d4fb28a9e5f93fc081392b457ad6fae3eb "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/multivalueextendedproperties"
    i3f305339dd659581dddbb6081f2d1981394ba5f409b3fff8d5affc26648441c4 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/singlevalueextendedproperties/item"
    i658437a8657f51783d49c87367e2ae5b495c006ac48f6a7010d75098566455a0 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/multivalueextendedproperties/item"
    i868bcbc69b2c3a35a368f852b47fffbb78632a8f15944de468f0e130a0242300 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/attachments/item"
    i955a97236a9f337897b15ed49defa3aa5c5d7910ae15f5b5efe5fbae5e75cb05 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/extensions/item"
    ic8bb7a27f91fd567d4417f19bf0aacb6595974ef7d4a830d5f1842796aab82f7 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item"
    ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/delta"
)

// EventRequestBuilder builds and executes requests for operations under \me\calendar\events\{event-id}
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
// EventRequestBuilderGetQueryParameters the events in the calendar. Navigation property. Read-only.
type EventRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// EventRequestBuilderPatchOptions options for Patch
type EventRequestBuilderPatchOptions struct {
    // 
    Body *ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec.Event;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventRequestBuilder) Accept()(*i87334fbab17098d26d733b7c6d4d979d97f28091d5d1b3c41c2e4adcd9d2f2a3.AcceptRequestBuilder) {
    return i87334fbab17098d26d733b7c6d4d979d97f28091d5d1b3c41c2e4adcd9d2f2a3.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i083a9367fb14f0da75caea7bc857922fe10d91693dde0f11517ab6420e557018.AttachmentsRequestBuilder) {
    return i083a9367fb14f0da75caea7bc857922fe10d91693dde0f11517ab6420e557018.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.events.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*i868bcbc69b2c3a35a368f852b47fffbb78632a8f15944de468f0e130a0242300.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i868bcbc69b2c3a35a368f852b47fffbb78632a8f15944de468f0e130a0242300.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i69bb0e8065fcc365cd51c4f9cb100c9289f8c378fe585dd0126d7eaee8953622.CalendarRequestBuilder) {
    return i69bb0e8065fcc365cd51c4f9cb100c9289f8c378fe585dd0126d7eaee8953622.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*ie04ca0f3f9823daf845f8b3b165c34feaffc35ed4189263a607834108a606dbe.CancelRequestBuilder) {
    return ie04ca0f3f9823daf845f8b3b165c34feaffc35ed4189263a607834108a606dbe.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/events/{event_id}{?select}";
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
// CreateDeleteRequestInformation the events in the calendar. Navigation property. Read-only.
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
// CreateGetRequestInformation the events in the calendar. Navigation property. Read-only.
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
// CreatePatchRequestInformation the events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Decline()(*i10a043ae8449d6b8061165bb67946e604bc2614783ade4a4557a2bb8d2a5a2b0.DeclineRequestBuilder) {
    return i10a043ae8449d6b8061165bb67946e604bc2614783ade4a4557a2bb8d2a5a2b0.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the events in the calendar. Navigation property. Read-only.
func (m *EventRequestBuilder) Delete(options *EventRequestBuilderDeleteOptions)(error) {
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
func (m *EventRequestBuilder) DismissReminder()(*ie6641980e1225311bdf1c499d054bc5d07aa799f49902ef90e71b01800c2a555.DismissReminderRequestBuilder) {
    return ie6641980e1225311bdf1c499d054bc5d07aa799f49902ef90e71b01800c2a555.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*id1d8fbd767f32fa77bb0a1e1f8b79c51f1a58b7dae1ccebcd3a3001af55d4f44.ExtensionsRequestBuilder) {
    return id1d8fbd767f32fa77bb0a1e1f8b79c51f1a58b7dae1ccebcd3a3001af55d4f44.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.events.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*i955a97236a9f337897b15ed49defa3aa5c5d7910ae15f5b5efe5fbae5e75cb05.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i955a97236a9f337897b15ed49defa3aa5c5d7910ae15f5b5efe5fbae5e75cb05.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i49ad5e4dbfe1e924deb723580385a06e87f447b9ed73e648b64a934117f1c784.ForwardRequestBuilder) {
    return i49ad5e4dbfe1e924deb723580385a06e87f447b9ed73e648b64a934117f1c784.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the events in the calendar. Navigation property. Read-only.
func (m *EventRequestBuilder) Get(options *EventRequestBuilderGetOptions)(*ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec.NewEvent() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec.Event), nil
}
func (m *EventRequestBuilder) Instances()(*i339d2d1a30983a01dd7de00efdd9e7d51cf0424dd9c4c96d335f9901fcff9763.InstancesRequestBuilder) {
    return i339d2d1a30983a01dd7de00efdd9e7d51cf0424dd9c4c96d335f9901fcff9763.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.events.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*ic8bb7a27f91fd567d4417f19bf0aacb6595974ef7d4a830d5f1842796aab82f7.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ic8bb7a27f91fd567d4417f19bf0aacb6595974ef7d4a830d5f1842796aab82f7.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*if78feb360a058ff87bf5a02baba6f7d4fb28a9e5f93fc081392b457ad6fae3eb.MultiValueExtendedPropertiesRequestBuilder) {
    return if78feb360a058ff87bf5a02baba6f7d4fb28a9e5f93fc081392b457ad6fae3eb.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.events.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i658437a8657f51783d49c87367e2ae5b495c006ac48f6a7010d75098566455a0.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i658437a8657f51783d49c87367e2ae5b495c006ac48f6a7010d75098566455a0.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the events in the calendar. Navigation property. Read-only.
func (m *EventRequestBuilder) Patch(options *EventRequestBuilderPatchOptions)(error) {
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i62d05692265cf61165d6bb2a1bb3a36d03ea2604586e3ebd550fae433333488e.SingleValueExtendedPropertiesRequestBuilder) {
    return i62d05692265cf61165d6bb2a1bb3a36d03ea2604586e3ebd550fae433333488e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.events.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i3f305339dd659581dddbb6081f2d1981394ba5f409b3fff8d5affc26648441c4.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i3f305339dd659581dddbb6081f2d1981394ba5f409b3fff8d5affc26648441c4.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i0f781326a67fbc9776be4d1ccbf43639b9bf8077faa925a7b0f90f0e03141c07.SnoozeReminderRequestBuilder) {
    return i0f781326a67fbc9776be4d1ccbf43639b9bf8077faa925a7b0f90f0e03141c07.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*iee14c2abd2614ebab15ceaf1135456d72c4e1cd953dd5d7b3347f819c29c6002.TentativelyAcceptRequestBuilder) {
    return iee14c2abd2614ebab15ceaf1135456d72c4e1cd953dd5d7b3347f819c29c6002.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
