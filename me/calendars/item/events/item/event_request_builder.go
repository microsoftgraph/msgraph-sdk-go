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

// Builds and executes requests for operations under \me\calendars\{calendar-id}\events\{event-id}
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
// The events in the calendar. Navigation property. Read-only.
type EventRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Select properties to be returned
    Select_escaped []string;
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
func (m *EventRequestBuilder) Accept()(*ib838f0df0ca92ac8dcade427c8a94d2f558237dc9cb00aa645b3cec1b9614bee.AcceptRequestBuilder) {
    return ib838f0df0ca92ac8dcade427c8a94d2f558237dc9cb00aa645b3cec1b9614bee.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i9b6b878aac98fae97c8910a9408f847174ad681c9aa60585e78ec39998ab2778.AttachmentsRequestBuilder) {
    return i9b6b878aac98fae97c8910a9408f847174ad681c9aa60585e78ec39998ab2778.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.calendars.item.events.item.attachments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) AttachmentsById(id string)(*i072f4a99f49d2b8f6761bf43cfd5b2b95f06274c07ba44d9a737e87df17d7886.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i072f4a99f49d2b8f6761bf43cfd5b2b95f06274c07ba44d9a737e87df17d7886.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*ibd576f1728c4dc9fbe5962a0ee16ce5f68c7cfab3b2852ddf9c17f07ed387504.CalendarRequestBuilder) {
    return ibd576f1728c4dc9fbe5962a0ee16ce5f68c7cfab3b2852ddf9c17f07ed387504.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i09aae1c2c441b5461fcc90aa35dcd1c4c2ca0552f429aa68c3db1c56bf0d2153.CancelRequestBuilder) {
    return i09aae1c2c441b5461fcc90aa35dcd1c4c2ca0552f429aa68c3db1c56bf0d2153.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
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
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventRequestBuilderInternal(urlParams, requestAdapter)
}
// The events in the calendar. Navigation property. Read-only.
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
// The events in the calendar. Navigation property. Read-only.
// Parameters:
//  - options : Options for the request
func (m *EventRequestBuilder) CreateGetRequestInformation(options *EventRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// The events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Decline()(*i882ced2bf5312be86e2c2ae6ea20a683cd710fdb806ed0768072570524785b86.DeclineRequestBuilder) {
    return i882ced2bf5312be86e2c2ae6ea20a683cd710fdb806ed0768072570524785b86.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) DismissReminder()(*i613d576972f3d8f4d501ac839ed6758ffca96792d87fe3cdc16bfbbcd999e339.DismissReminderRequestBuilder) {
    return i613d576972f3d8f4d501ac839ed6758ffca96792d87fe3cdc16bfbbcd999e339.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*idb61d479748362a6e089afd5c0b858f09487d50b611d7917271603f399cddcc3.ExtensionsRequestBuilder) {
    return idb61d479748362a6e089afd5c0b858f09487d50b611d7917271603f399cddcc3.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.calendars.item.events.item.extensions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) ExtensionsById(id string)(*ifd72b733fd36497510451fbda93ed7a025d727b889c80b685dc6fcdf606233c6.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ifd72b733fd36497510451fbda93ed7a025d727b889c80b685dc6fcdf606233c6.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i27c7643dce2bb6c7a69be2cdf3257e7c4a990cf16f691aac9ce3596d45aad48f.ForwardRequestBuilder) {
    return i27c7643dce2bb6c7a69be2cdf3257e7c4a990cf16f691aac9ce3596d45aad48f.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Instances()(*iac14f0c5ccb2a5f3b9bafd14cee7eed24b0f2d2d6879f08fbd6269784857e360.InstancesRequestBuilder) {
    return iac14f0c5ccb2a5f3b9bafd14cee7eed24b0f2d2d6879f08fbd6269784857e360.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.calendars.item.events.item.instances.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) InstancesById(id string)(*i4b79cd0561a1ab0685360c4682018160dedd8ecad4d5fd7dd32992b4017c80fd.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i4b79cd0561a1ab0685360c4682018160dedd8ecad4d5fd7dd32992b4017c80fd.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i8e5f4a41163bc4194e03167ef61b98c0e69797a47d44c62adb54062836a1bbe6.MultiValueExtendedPropertiesRequestBuilder) {
    return i8e5f4a41163bc4194e03167ef61b98c0e69797a47d44c62adb54062836a1bbe6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.calendars.item.events.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i157450aaee9c55248f9ba5ffbbf66bf16cc3d8d3d2bd8c9cfd4a765ad8bae272.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i157450aaee9c55248f9ba5ffbbf66bf16cc3d8d3d2bd8c9cfd4a765ad8bae272.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i72cf340f081c108008d2cadd6aaa88898cb9a7ec52a019ac18c447c5fc9e3cc4.SingleValueExtendedPropertiesRequestBuilder) {
    return i72cf340f081c108008d2cadd6aaa88898cb9a7ec52a019ac18c447c5fc9e3cc4.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.calendars.item.events.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i5b791a9e8a95f008bb13124c648ebdef495fb58e993ff30cf6006693ece7f80a.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i5b791a9e8a95f008bb13124c648ebdef495fb58e993ff30cf6006693ece7f80a.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i383bbe46697a339b06b5bafed4a3590534909db04af55c794a8faec30e41ca1a.SnoozeReminderRequestBuilder) {
    return i383bbe46697a339b06b5bafed4a3590534909db04af55c794a8faec30e41ca1a.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*idb076a8a69d2a57d5dcdc849f6cbb856c0d28a0d737152e3dd6285793373a285.TentativelyAcceptRequestBuilder) {
    return idb076a8a69d2a57d5dcdc849f6cbb856c0d28a0d737152e3dd6285793373a285.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
