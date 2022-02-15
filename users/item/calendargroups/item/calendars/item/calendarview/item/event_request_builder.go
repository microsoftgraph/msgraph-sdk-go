package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ica42cd38493faf35d7bf41e56bd3d77d5273039e1066c2dae628f21ac8349eec "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/delta"
    i13d9d853f6f596efbe21ea18f40d0592475a8ac990e8a695d81a6fbba14f4d26 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/cancel"
    i1a10f47c94aaf40994f428a074851ae6a632b047055a7f1cf3bdd2b7c7b6c0a5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/singlevalueextendedproperties"
    i1b64d030401b44fa9d8e6c1050f609bba7b33042a59456f5e5f4bf592c0d9b93 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/snoozereminder"
    i2e1f9fb6aaa8abf5f8df5e11ba377fe0143a5e43f512624a8d9cd6fb27421a78 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/dismissreminder"
    i91354a288a328a774d7e175c52b2d23e26aa9ea2c90072fa527d0a581ca4bc8e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/tentativelyaccept"
    ia088d5c6b1f3a9f9a2aeaad3b76325dbf4d601ee30542e1676570dbf971be3f8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/decline"
    iaf599fb4f2ae1ccff34614f5a03bcc714f19cbff2c171838f87122ed74aecb31 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances"
    ica7ba27254da644cf113708d63a508c5069debc7ee258b273bee580a0147682d "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/multivalueextendedproperties"
    idd8083a89a62c4a90c13d7095b22cbdaa8ecd14901ed17edda30c0d6194c4c4f "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/calendar"
    ie21405766b71afd03cdca486d941e721d602f890de39f16cfb24032bae39294b "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/extensions"
    ie8336fa6b5a4ea27b8de9a04a40237f8ebccb1dfe6f97770ce15b89d6c3e2f95 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/attachments"
    ie8d2a5a2eab16f2219d42c95df3d68b757fb3773211e5b6ae612f6646a67a1b1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/forward"
    if3456a6c7e36e0e413ac6ae411ab193961324f5aa72c986631a16de60564c13e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/accept"
    i089a2accd06c2c34e85f4442bcfe40065651bb4f57e23851aa07d1ea2becdc21 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/extensions/item"
    i0be4955da7d4a2edcd0b70a9d0b170efd62f033d7459723d8439132a58f846c9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item"
    iadb079fbc3d0d06ba36c2c0c66d5f3073b14e52cd15b1db7a9bd6befec932fa3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/multivalueextendedproperties/item"
    ib451c2fc927d457e02148536325b3839f57ef91bec0997dd64fb4888db642574 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/singlevalueextendedproperties/item"
    iee3b48332651e2bc68d677bf62b811a958d2bac5d0f1786cad3f5eece9d248f8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/attachments/item"
)

// EventRequestBuilder builds and executes requests for operations under \users\{user-id}\calendarGroups\{calendarGroup-id}\calendars\{calendar-id}\calendarView\{event-id}
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
func (m *EventRequestBuilder) Accept()(*if3456a6c7e36e0e413ac6ae411ab193961324f5aa72c986631a16de60564c13e.AcceptRequestBuilder) {
    return if3456a6c7e36e0e413ac6ae411ab193961324f5aa72c986631a16de60564c13e.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*ie8336fa6b5a4ea27b8de9a04a40237f8ebccb1dfe6f97770ce15b89d6c3e2f95.AttachmentsRequestBuilder) {
    return ie8336fa6b5a4ea27b8de9a04a40237f8ebccb1dfe6f97770ce15b89d6c3e2f95.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*iee3b48332651e2bc68d677bf62b811a958d2bac5d0f1786cad3f5eece9d248f8.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return iee3b48332651e2bc68d677bf62b811a958d2bac5d0f1786cad3f5eece9d248f8.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*idd8083a89a62c4a90c13d7095b22cbdaa8ecd14901ed17edda30c0d6194c4c4f.CalendarRequestBuilder) {
    return idd8083a89a62c4a90c13d7095b22cbdaa8ecd14901ed17edda30c0d6194c4c4f.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i13d9d853f6f596efbe21ea18f40d0592475a8ac990e8a695d81a6fbba14f4d26.CancelRequestBuilder) {
    return i13d9d853f6f596efbe21ea18f40d0592475a8ac990e8a695d81a6fbba14f4d26.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}/calendarView/{event_id}{?select}";
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
func (m *EventRequestBuilder) Decline()(*ia088d5c6b1f3a9f9a2aeaad3b76325dbf4d601ee30542e1676570dbf971be3f8.DeclineRequestBuilder) {
    return ia088d5c6b1f3a9f9a2aeaad3b76325dbf4d601ee30542e1676570dbf971be3f8.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) DismissReminder()(*i2e1f9fb6aaa8abf5f8df5e11ba377fe0143a5e43f512624a8d9cd6fb27421a78.DismissReminderRequestBuilder) {
    return i2e1f9fb6aaa8abf5f8df5e11ba377fe0143a5e43f512624a8d9cd6fb27421a78.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*ie21405766b71afd03cdca486d941e721d602f890de39f16cfb24032bae39294b.ExtensionsRequestBuilder) {
    return ie21405766b71afd03cdca486d941e721d602f890de39f16cfb24032bae39294b.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*i089a2accd06c2c34e85f4442bcfe40065651bb4f57e23851aa07d1ea2becdc21.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i089a2accd06c2c34e85f4442bcfe40065651bb4f57e23851aa07d1ea2becdc21.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*ie8d2a5a2eab16f2219d42c95df3d68b757fb3773211e5b6ae612f6646a67a1b1.ForwardRequestBuilder) {
    return ie8d2a5a2eab16f2219d42c95df3d68b757fb3773211e5b6ae612f6646a67a1b1.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) Instances()(*iaf599fb4f2ae1ccff34614f5a03bcc714f19cbff2c171838f87122ed74aecb31.InstancesRequestBuilder) {
    return iaf599fb4f2ae1ccff34614f5a03bcc714f19cbff2c171838f87122ed74aecb31.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*i0be4955da7d4a2edcd0b70a9d0b170efd62f033d7459723d8439132a58f846c9.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i0be4955da7d4a2edcd0b70a9d0b170efd62f033d7459723d8439132a58f846c9.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*ica7ba27254da644cf113708d63a508c5069debc7ee258b273bee580a0147682d.MultiValueExtendedPropertiesRequestBuilder) {
    return ica7ba27254da644cf113708d63a508c5069debc7ee258b273bee580a0147682d.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iadb079fbc3d0d06ba36c2c0c66d5f3073b14e52cd15b1db7a9bd6befec932fa3.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return iadb079fbc3d0d06ba36c2c0c66d5f3073b14e52cd15b1db7a9bd6befec932fa3.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i1a10f47c94aaf40994f428a074851ae6a632b047055a7f1cf3bdd2b7c7b6c0a5.SingleValueExtendedPropertiesRequestBuilder) {
    return i1a10f47c94aaf40994f428a074851ae6a632b047055a7f1cf3bdd2b7c7b6c0a5.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ib451c2fc927d457e02148536325b3839f57ef91bec0997dd64fb4888db642574.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ib451c2fc927d457e02148536325b3839f57ef91bec0997dd64fb4888db642574.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i1b64d030401b44fa9d8e6c1050f609bba7b33042a59456f5e5f4bf592c0d9b93.SnoozeReminderRequestBuilder) {
    return i1b64d030401b44fa9d8e6c1050f609bba7b33042a59456f5e5f4bf592c0d9b93.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i91354a288a328a774d7e175c52b2d23e26aa9ea2c90072fa527d0a581ca4bc8e.TentativelyAcceptRequestBuilder) {
    return i91354a288a328a774d7e175c52b2d23e26aa9ea2c90072fa527d0a581ca4bc8e.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
