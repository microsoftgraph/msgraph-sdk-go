package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i12c00cf6d17f76414eb8bf4af116ae0f94a375b16a3fdcc2a520db9d8b51bb69 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/calendar"
    i382bc5ab005174ca0a157f3ddbbe59b93465e03639aea98bec85f98a02704a20 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/extensions"
    i52fe9bd763a706a85cde883eea661538b98807c8ed4896d571ad9f373de4a24b "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/dismissreminder"
    i8cb640b8e824bfdf805685b8399a0ddb642f5f041f6f7a6d1a2adbbd8b04242c "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/attachments"
    i91a25c006065bdbda1ec6dbad739849f30b2ed0ea0a864cf75a845796dfb463a "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/cancel"
    ia41b7c24cc3f3baeaa7815e53e95f6fede450c797c61686c04ae5509cff1fee9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/accept"
    iac5010b0952b1db93baa23920b116d827c3b2ec1b7e9168259f8668f64c64732 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/tentativelyaccept"
    ib2891d93c6ded684ce7abfd2909cc1a20fc2090c50db20bbe6da0375eb2f3fe7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances"
    ib2f603875392f9c2b340caee4eab840215f32e0520cb72749ad970c7f46e8ced "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/decline"
    id148b2204d6e284392369bbf295c7ad603ca498161a3166593a744b1705d35d7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/multivalueextendedproperties"
    ieabc03565c45dc93b6dd73031a0f62fcc81744682614225dee84662b38c937e3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/forward"
    if1d75ca20ba55d1dcc9dfcad50836d4235106da1e2a01e8cbd7d0a07f344df20 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/singlevalueextendedproperties"
    ife09ca7c64720802742bbf52d953fd6747f5910742fbabcc2a740709adb9c2cc "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/snoozereminder"
    i8fcec19e7c3ce8aa06ca5657386121d5140cfc8138e4f515339a4475a80f52be "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/instances/item"
    iabd4185ab5ac976d8418c3afea4e9d5b739795c1459c53296cbe08e4cdb2a44a "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/extensions/item"
    id1c9a5e9fa77008270b7af81a4a4663bc2e9e1151e2512f7f7783693e4a62854 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/singlevalueextendedproperties/item"
    ie0f0bdcc9602484baf79d027f6483fa9fbafd07fe65976cf1a96cc9ced359802 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/attachments/item"
    ie456cc01b9ad3635053091de1d8edb69f6dfa30b76d79790c6668bf7307efe8e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item/multivalueextendedproperties/item"
)

// EventRequestBuilder builds and executes requests for operations under \users\{user-id}\calendars\{calendar-id}\events\{event-id}
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
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventRequestBuilder) Accept()(*ia41b7c24cc3f3baeaa7815e53e95f6fede450c797c61686c04ae5509cff1fee9.AcceptRequestBuilder) {
    return ia41b7c24cc3f3baeaa7815e53e95f6fede450c797c61686c04ae5509cff1fee9.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i8cb640b8e824bfdf805685b8399a0ddb642f5f041f6f7a6d1a2adbbd8b04242c.AttachmentsRequestBuilder) {
    return i8cb640b8e824bfdf805685b8399a0ddb642f5f041f6f7a6d1a2adbbd8b04242c.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.events.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*ie0f0bdcc9602484baf79d027f6483fa9fbafd07fe65976cf1a96cc9ced359802.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ie0f0bdcc9602484baf79d027f6483fa9fbafd07fe65976cf1a96cc9ced359802.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i12c00cf6d17f76414eb8bf4af116ae0f94a375b16a3fdcc2a520db9d8b51bb69.CalendarRequestBuilder) {
    return i12c00cf6d17f76414eb8bf4af116ae0f94a375b16a3fdcc2a520db9d8b51bb69.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i91a25c006065bdbda1ec6dbad739849f30b2ed0ea0a864cf75a845796dfb463a.CancelRequestBuilder) {
    return i91a25c006065bdbda1ec6dbad739849f30b2ed0ea0a864cf75a845796dfb463a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendars/{calendar_id}/events/{event_id}{?select}";
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
func (m *EventRequestBuilder) Decline()(*ib2f603875392f9c2b340caee4eab840215f32e0520cb72749ad970c7f46e8ced.DeclineRequestBuilder) {
    return ib2f603875392f9c2b340caee4eab840215f32e0520cb72749ad970c7f46e8ced.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) DismissReminder()(*i52fe9bd763a706a85cde883eea661538b98807c8ed4896d571ad9f373de4a24b.DismissReminderRequestBuilder) {
    return i52fe9bd763a706a85cde883eea661538b98807c8ed4896d571ad9f373de4a24b.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i382bc5ab005174ca0a157f3ddbbe59b93465e03639aea98bec85f98a02704a20.ExtensionsRequestBuilder) {
    return i382bc5ab005174ca0a157f3ddbbe59b93465e03639aea98bec85f98a02704a20.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.events.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*iabd4185ab5ac976d8418c3afea4e9d5b739795c1459c53296cbe08e4cdb2a44a.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return iabd4185ab5ac976d8418c3afea4e9d5b739795c1459c53296cbe08e4cdb2a44a.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*ieabc03565c45dc93b6dd73031a0f62fcc81744682614225dee84662b38c937e3.ForwardRequestBuilder) {
    return ieabc03565c45dc93b6dd73031a0f62fcc81744682614225dee84662b38c937e3.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Instances()(*ib2891d93c6ded684ce7abfd2909cc1a20fc2090c50db20bbe6da0375eb2f3fe7.InstancesRequestBuilder) {
    return ib2891d93c6ded684ce7abfd2909cc1a20fc2090c50db20bbe6da0375eb2f3fe7.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.events.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*i8fcec19e7c3ce8aa06ca5657386121d5140cfc8138e4f515339a4475a80f52be.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i8fcec19e7c3ce8aa06ca5657386121d5140cfc8138e4f515339a4475a80f52be.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*id148b2204d6e284392369bbf295c7ad603ca498161a3166593a744b1705d35d7.MultiValueExtendedPropertiesRequestBuilder) {
    return id148b2204d6e284392369bbf295c7ad603ca498161a3166593a744b1705d35d7.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.events.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ie456cc01b9ad3635053091de1d8edb69f6dfa30b76d79790c6668bf7307efe8e.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ie456cc01b9ad3635053091de1d8edb69f6dfa30b76d79790c6668bf7307efe8e.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*if1d75ca20ba55d1dcc9dfcad50836d4235106da1e2a01e8cbd7d0a07f344df20.SingleValueExtendedPropertiesRequestBuilder) {
    return if1d75ca20ba55d1dcc9dfcad50836d4235106da1e2a01e8cbd7d0a07f344df20.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.events.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*id1c9a5e9fa77008270b7af81a4a4663bc2e9e1151e2512f7f7783693e4a62854.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return id1c9a5e9fa77008270b7af81a4a4663bc2e9e1151e2512f7f7783693e4a62854.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*ife09ca7c64720802742bbf52d953fd6747f5910742fbabcc2a740709adb9c2cc.SnoozeReminderRequestBuilder) {
    return ife09ca7c64720802742bbf52d953fd6747f5910742fbabcc2a740709adb9c2cc.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*iac5010b0952b1db93baa23920b116d827c3b2ec1b7e9168259f8668f64c64732.TentativelyAcceptRequestBuilder) {
    return iac5010b0952b1db93baa23920b116d827c3b2ec1b7e9168259f8668f64c64732.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
