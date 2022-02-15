package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i145fb7990a2e3fdab872136aac6bb8be775248efc831fc0ceac357c425bae97f "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/decline"
    i263c92c411c89f2be6a8e54de5fc9a4dde982d0e6d8de040c71d4afe63a7748c "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/cancel"
    i580c1b6081b469d14dea10fa0ff4e3fcf1e92e486940fe417a29f9f5671bbb53 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances"
    i896d4bd262b2cdabe891cd22c1b08cbcde12e3b522ff979daf20e82ad1ea711b "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/snoozereminder"
    i9e33412fe1bdfa6c018b792dd1429c9276d49c06fe8c88117fede514ab522720 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/forward"
    ia02addaa9051893cf138a732042bf251456286968b685c328064520251cb0cf3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/extensions"
    ia3dd18af0ea5a7f297b2a2edda6570fc313fc5708c7e2634ba8259e0ae32b2a5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/tentativelyaccept"
    ia62de877959a200aca4887792a575daf1a8f28bf9cfcb7a1fff2251ff773c997 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/accept"
    ia93978b34743023ce443a9997d7630ed381895aed8d859d3f4ac245be62fb849 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/multivalueextendedproperties"
    iaae9964b69928703e801f04d74fc2a4f7d6498b08c64baf44a7eb2b78040bcfc "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/calendar"
    iae0f8cfe0219a22eb2de5c896f00609e7fb3fb5d12dfd6203fee1bdcc65508d7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/dismissreminder"
    ib55ea188c0bda46cccf674fa54b66612f2fa73c5e8d49e10718ad5878343eb9f "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/attachments"
    id8bc39608b45df9d2acb9059f8be81c090091fe783571dbbd283d47b33d831d5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/singlevalueextendedproperties"
    i01f8c58bdadd020940a9f0c219ace3903f9dd0225f07ec8fe1fdfc9eb6c2c0f8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/multivalueextendedproperties/item"
    i061c7c8e6a1b9d55c8cc8301b7c4370b5310fd0b6c3a2c8c5217710a38ec869b "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/extensions/item"
    i57a0e5d0a9a8aa62edc5ca53b30b8ad23c80c88a7af4e91c43ef9504c35fd4a4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/singlevalueextendedproperties/item"
    i66e4305766d125651f3fd13dcddaba80315e4ca5673cb2b5eda12fdabe6174fb "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/item"
    ia402890f08866a24e37023bade2f4fd90979ed0d599b4b006b83fd80bee80d01 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/attachments/item"
    ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/delta"
)

// EventRequestBuilder builds and executes requests for operations under \users\{user-id}\calendar\calendarView\{event-id}
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
// EventRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
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
    Body *ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec.Event;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventRequestBuilder) Accept()(*ia62de877959a200aca4887792a575daf1a8f28bf9cfcb7a1fff2251ff773c997.AcceptRequestBuilder) {
    return ia62de877959a200aca4887792a575daf1a8f28bf9cfcb7a1fff2251ff773c997.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*ib55ea188c0bda46cccf674fa54b66612f2fa73c5e8d49e10718ad5878343eb9f.AttachmentsRequestBuilder) {
    return ib55ea188c0bda46cccf674fa54b66612f2fa73c5e8d49e10718ad5878343eb9f.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.calendarView.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*ia402890f08866a24e37023bade2f4fd90979ed0d599b4b006b83fd80bee80d01.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ia402890f08866a24e37023bade2f4fd90979ed0d599b4b006b83fd80bee80d01.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*iaae9964b69928703e801f04d74fc2a4f7d6498b08c64baf44a7eb2b78040bcfc.CalendarRequestBuilder) {
    return iaae9964b69928703e801f04d74fc2a4f7d6498b08c64baf44a7eb2b78040bcfc.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i263c92c411c89f2be6a8e54de5fc9a4dde982d0e6d8de040c71d4afe63a7748c.CancelRequestBuilder) {
    return i263c92c411c89f2be6a8e54de5fc9a4dde982d0e6d8de040c71d4afe63a7748c.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendar/calendarView/{event_id}{?startDateTime,endDateTime,select}";
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
// CreateDeleteRequestInformation the calendar view for the calendar. Navigation property. Read-only.
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
// CreateGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
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
// CreatePatchRequestInformation the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Decline()(*i145fb7990a2e3fdab872136aac6bb8be775248efc831fc0ceac357c425bae97f.DeclineRequestBuilder) {
    return i145fb7990a2e3fdab872136aac6bb8be775248efc831fc0ceac357c425bae97f.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) DismissReminder()(*iae0f8cfe0219a22eb2de5c896f00609e7fb3fb5d12dfd6203fee1bdcc65508d7.DismissReminderRequestBuilder) {
    return iae0f8cfe0219a22eb2de5c896f00609e7fb3fb5d12dfd6203fee1bdcc65508d7.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*ia02addaa9051893cf138a732042bf251456286968b685c328064520251cb0cf3.ExtensionsRequestBuilder) {
    return ia02addaa9051893cf138a732042bf251456286968b685c328064520251cb0cf3.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.calendarView.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*i061c7c8e6a1b9d55c8cc8301b7c4370b5310fd0b6c3a2c8c5217710a38ec869b.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i061c7c8e6a1b9d55c8cc8301b7c4370b5310fd0b6c3a2c8c5217710a38ec869b.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i9e33412fe1bdfa6c018b792dd1429c9276d49c06fe8c88117fede514ab522720.ForwardRequestBuilder) {
    return i9e33412fe1bdfa6c018b792dd1429c9276d49c06fe8c88117fede514ab522720.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Instances()(*i580c1b6081b469d14dea10fa0ff4e3fcf1e92e486940fe417a29f9f5671bbb53.InstancesRequestBuilder) {
    return i580c1b6081b469d14dea10fa0ff4e3fcf1e92e486940fe417a29f9f5671bbb53.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.calendarView.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*i66e4305766d125651f3fd13dcddaba80315e4ca5673cb2b5eda12fdabe6174fb.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i66e4305766d125651f3fd13dcddaba80315e4ca5673cb2b5eda12fdabe6174fb.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*ia93978b34743023ce443a9997d7630ed381895aed8d859d3f4ac245be62fb849.MultiValueExtendedPropertiesRequestBuilder) {
    return ia93978b34743023ce443a9997d7630ed381895aed8d859d3f4ac245be62fb849.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i01f8c58bdadd020940a9f0c219ace3903f9dd0225f07ec8fe1fdfc9eb6c2c0f8.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i01f8c58bdadd020940a9f0c219ace3903f9dd0225f07ec8fe1fdfc9eb6c2c0f8.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*id8bc39608b45df9d2acb9059f8be81c090091fe783571dbbd283d47b33d831d5.SingleValueExtendedPropertiesRequestBuilder) {
    return id8bc39608b45df9d2acb9059f8be81c090091fe783571dbbd283d47b33d831d5.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i57a0e5d0a9a8aa62edc5ca53b30b8ad23c80c88a7af4e91c43ef9504c35fd4a4.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i57a0e5d0a9a8aa62edc5ca53b30b8ad23c80c88a7af4e91c43ef9504c35fd4a4.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i896d4bd262b2cdabe891cd22c1b08cbcde12e3b522ff979daf20e82ad1ea711b.SnoozeReminderRequestBuilder) {
    return i896d4bd262b2cdabe891cd22c1b08cbcde12e3b522ff979daf20e82ad1ea711b.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*ia3dd18af0ea5a7f297b2a2edda6570fc313fc5708c7e2634ba8259e0ae32b2a5.TentativelyAcceptRequestBuilder) {
    return ia3dd18af0ea5a7f297b2a2edda6570fc313fc5708c7e2634ba8259e0ae32b2a5.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
