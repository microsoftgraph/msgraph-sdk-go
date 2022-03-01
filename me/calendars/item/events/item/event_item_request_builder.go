package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i09aae1c2c441b5461fcc90aa35dcd1c4c2ca0552f429aa68c3db1c56bf0d2153 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/cancel"
    i27c7643dce2bb6c7a69be2cdf3257e7c4a990cf16f691aac9ce3596d45aad48f "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/forward"
    i383bbe46697a339b06b5bafed4a3590534909db04af55c794a8faec30e41ca1a "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/snoozereminder"
    i613d576972f3d8f4d501ac839ed6758ffca96792d87fe3cdc16bfbbcd999e339 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/dismissreminder"
    i72cf340f081c108008d2cadd6aaa88898cb9a7ec52a019ac18c447c5fc9e3cc4 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/singlevalueextendedproperties"
    i882ced2bf5312be86e2c2ae6ea20a683cd710fdb806ed0768072570524785b86 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/decline"
    i8e5f4a41163bc4194e03167ef61b98c0e69797a47d44c62adb54062836a1bbe6 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/multivalueextendedproperties"
    i9b6b878aac98fae97c8910a9408f847174ad681c9aa60585e78ec39998ab2778 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/attachments"
    iac14f0c5ccb2a5f3b9bafd14cee7eed24b0f2d2d6879f08fbd6269784857e360 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances"
    ib838f0df0ca92ac8dcade427c8a94d2f558237dc9cb00aa645b3cec1b9614bee "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/accept"
    ibd576f1728c4dc9fbe5962a0ee16ce5f68c7cfab3b2852ddf9c17f07ed387504 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/calendar"
    idb076a8a69d2a57d5dcdc849f6cbb856c0d28a0d737152e3dd6285793373a285 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/tentativelyaccept"
    idb61d479748362a6e089afd5c0b858f09487d50b611d7917271603f399cddcc3 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/extensions"
    i072f4a99f49d2b8f6761bf43cfd5b2b95f06274c07ba44d9a737e87df17d7886 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/attachments/item"
    i157450aaee9c55248f9ba5ffbbf66bf16cc3d8d3d2bd8c9cfd4a765ad8bae272 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/multivalueextendedproperties/item"
    i4b79cd0561a1ab0685360c4682018160dedd8ecad4d5fd7dd32992b4017c80fd "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item"
    i5b791a9e8a95f008bb13124c648ebdef495fb58e993ff30cf6006693ece7f80a "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/singlevalueextendedproperties/item"
    ifd72b733fd36497510451fbda93ed7a025d727b889c80b685dc6fcdf606233c6 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/extensions/item"
)

// EventItemRequestBuilder builds and executes requests for operations under \me\calendars\{calendar-id}\events\{event-id}
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventItemRequestBuilderDeleteOptions options for Delete
type EventItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventItemRequestBuilderGetOptions options for Get
type EventItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EventItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventItemRequestBuilderGetQueryParameters the events in the calendar. Navigation property. Read-only.
type EventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// EventItemRequestBuilderPatchOptions options for Patch
type EventItemRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventItemRequestBuilder) Accept()(*ib838f0df0ca92ac8dcade427c8a94d2f558237dc9cb00aa645b3cec1b9614bee.AcceptRequestBuilder) {
    return ib838f0df0ca92ac8dcade427c8a94d2f558237dc9cb00aa645b3cec1b9614bee.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i9b6b878aac98fae97c8910a9408f847174ad681c9aa60585e78ec39998ab2778.AttachmentsRequestBuilder) {
    return i9b6b878aac98fae97c8910a9408f847174ad681c9aa60585e78ec39998ab2778.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i072f4a99f49d2b8f6761bf43cfd5b2b95f06274c07ba44d9a737e87df17d7886.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i072f4a99f49d2b8f6761bf43cfd5b2b95f06274c07ba44d9a737e87df17d7886.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*ibd576f1728c4dc9fbe5962a0ee16ce5f68c7cfab3b2852ddf9c17f07ed387504.CalendarRequestBuilder) {
    return ibd576f1728c4dc9fbe5962a0ee16ce5f68c7cfab3b2852ddf9c17f07ed387504.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i09aae1c2c441b5461fcc90aa35dcd1c4c2ca0552f429aa68c3db1c56bf0d2153.CancelRequestBuilder) {
    return i09aae1c2c441b5461fcc90aa35dcd1c4c2ca0552f429aa68c3db1c56bf0d2153.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar_id}/events/{event_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the events in the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation(options *EventItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventItemRequestBuilder) CreateGetRequestInformation(options *EventItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(options *EventItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventItemRequestBuilder) Decline()(*i882ced2bf5312be86e2c2ae6ea20a683cd710fdb806ed0768072570524785b86.DeclineRequestBuilder) {
    return i882ced2bf5312be86e2c2ae6ea20a683cd710fdb806ed0768072570524785b86.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the events in the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
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
func (m *EventItemRequestBuilder) DismissReminder()(*i613d576972f3d8f4d501ac839ed6758ffca96792d87fe3cdc16bfbbcd999e339.DismissReminderRequestBuilder) {
    return i613d576972f3d8f4d501ac839ed6758ffca96792d87fe3cdc16bfbbcd999e339.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*idb61d479748362a6e089afd5c0b858f09487d50b611d7917271603f399cddcc3.ExtensionsRequestBuilder) {
    return idb61d479748362a6e089afd5c0b858f09487d50b611d7917271603f399cddcc3.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ifd72b733fd36497510451fbda93ed7a025d727b889c80b685dc6fcdf606233c6.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ifd72b733fd36497510451fbda93ed7a025d727b889c80b685dc6fcdf606233c6.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i27c7643dce2bb6c7a69be2cdf3257e7c4a990cf16f691aac9ce3596d45aad48f.ForwardRequestBuilder) {
    return i27c7643dce2bb6c7a69be2cdf3257e7c4a990cf16f691aac9ce3596d45aad48f.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the events in the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEvent() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event), nil
}
func (m *EventItemRequestBuilder) Instances()(*iac14f0c5ccb2a5f3b9bafd14cee7eed24b0f2d2d6879f08fbd6269784857e360.InstancesRequestBuilder) {
    return iac14f0c5ccb2a5f3b9bafd14cee7eed24b0f2d2d6879f08fbd6269784857e360.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i4b79cd0561a1ab0685360c4682018160dedd8ecad4d5fd7dd32992b4017c80fd.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i4b79cd0561a1ab0685360c4682018160dedd8ecad4d5fd7dd32992b4017c80fd.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i8e5f4a41163bc4194e03167ef61b98c0e69797a47d44c62adb54062836a1bbe6.MultiValueExtendedPropertiesRequestBuilder) {
    return i8e5f4a41163bc4194e03167ef61b98c0e69797a47d44c62adb54062836a1bbe6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i157450aaee9c55248f9ba5ffbbf66bf16cc3d8d3d2bd8c9cfd4a765ad8bae272.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i157450aaee9c55248f9ba5ffbbf66bf16cc3d8d3d2bd8c9cfd4a765ad8bae272.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the events in the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i72cf340f081c108008d2cadd6aaa88898cb9a7ec52a019ac18c447c5fc9e3cc4.SingleValueExtendedPropertiesRequestBuilder) {
    return i72cf340f081c108008d2cadd6aaa88898cb9a7ec52a019ac18c447c5fc9e3cc4.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i5b791a9e8a95f008bb13124c648ebdef495fb58e993ff30cf6006693ece7f80a.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i5b791a9e8a95f008bb13124c648ebdef495fb58e993ff30cf6006693ece7f80a.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i383bbe46697a339b06b5bafed4a3590534909db04af55c794a8faec30e41ca1a.SnoozeReminderRequestBuilder) {
    return i383bbe46697a339b06b5bafed4a3590534909db04af55c794a8faec30e41ca1a.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*idb076a8a69d2a57d5dcdc849f6cbb856c0d28a0d737152e3dd6285793373a285.TentativelyAcceptRequestBuilder) {
    return idb076a8a69d2a57d5dcdc849f6cbb856c0d28a0d737152e3dd6285793373a285.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
