package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i07aefebef2360a05f1f978a4dacfbfc958f6263953281cb4e4810c24e8fdbe51 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/forward"
    i130329fd346118a25e215e9ae08e6acafc2fb545bcc2544394704d98040e0fd5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/tentativelyaccept"
    i1b6d6c642ae53b6d75418c14f94432b149eb82d231e9ef4719cabf3dd36e4a84 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/decline"
    i47c119d74e0c4d1b7732433a7add8a684909bad78a2169bdadc752f8052b2cde "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/attachments"
    i8bc1b83864742db92b87f3ce691ba7194b9e4dfcdee44bc1f989db1fd24e0d9c "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/dismissreminder"
    i8daf6ca6792aac2f5e22afadd880b9416a73f6d66f3d3a4cb4ce49b702beabff "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances"
    iaafeb0437eaa3b26607893fda20fbdae77ac9d7e2333ed4d5b73883b77acff0a "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/snoozereminder"
    iabd200b2c68fcfd56b3a70da55accf4de8cff2e02b0264e62b3ec575b67993fd "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/cancel"
    ib1691f681b565035438b2134f4d2fff4a5be77dd06fbf2cefb004ae233672e01 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/calendar"
    ic2ca36017987ae7a5daa23b834522d6cb1b251d72fe6d72b471f8ca4ed2d2264 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/multivalueextendedproperties"
    ic9d89564024358edace4430f9c467c2085e6b36f70aa7b6fb0605f47881321b4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/extensions"
    icfb49c8c2637abda0a6e0b8db57c259abb3d4ec46f18de110de6cc049d986341 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/singlevalueextendedproperties"
    if35aa5926f23e100256e1aa6716e7d17471936c8947905b607e1035b11cc8f34 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/accept"
    i065657a3daf360b62ebf3c6e101d6bd29d67ca95848bd319d1ffeb09214c3cd5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/attachments/item"
    i49c7e84876761d3a9df15aa2f74f21992b0ba54e8eea9060140c4135b4b76920 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/singlevalueextendedproperties/item"
    i9facca5460444460406e305a714abf1b6c6c7e0be4c7bedda91eb9b394bd9ede "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/multivalueextendedproperties/item"
    ibe0f4b62b27e66ba9582cbb8b856bbb6bd0e3e8a30c85f69943b3fc01aac856e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/instances/item"
    ibf212a370d44f80b700c5b5b75239c5e02548b6abbd3484f7872aa4674bf95c4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/events/item/extensions/item"
)

// EventRequestBuilder builds and executes requests for operations under \users\{user-id}\calendar\events\{event-id}
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
func (m *EventRequestBuilder) Accept()(*if35aa5926f23e100256e1aa6716e7d17471936c8947905b607e1035b11cc8f34.AcceptRequestBuilder) {
    return if35aa5926f23e100256e1aa6716e7d17471936c8947905b607e1035b11cc8f34.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i47c119d74e0c4d1b7732433a7add8a684909bad78a2169bdadc752f8052b2cde.AttachmentsRequestBuilder) {
    return i47c119d74e0c4d1b7732433a7add8a684909bad78a2169bdadc752f8052b2cde.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.events.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*i065657a3daf360b62ebf3c6e101d6bd29d67ca95848bd319d1ffeb09214c3cd5.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i065657a3daf360b62ebf3c6e101d6bd29d67ca95848bd319d1ffeb09214c3cd5.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*ib1691f681b565035438b2134f4d2fff4a5be77dd06fbf2cefb004ae233672e01.CalendarRequestBuilder) {
    return ib1691f681b565035438b2134f4d2fff4a5be77dd06fbf2cefb004ae233672e01.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*iabd200b2c68fcfd56b3a70da55accf4de8cff2e02b0264e62b3ec575b67993fd.CancelRequestBuilder) {
    return iabd200b2c68fcfd56b3a70da55accf4de8cff2e02b0264e62b3ec575b67993fd.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendar/events/{event_id}{?select}";
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
func (m *EventRequestBuilder) Decline()(*i1b6d6c642ae53b6d75418c14f94432b149eb82d231e9ef4719cabf3dd36e4a84.DeclineRequestBuilder) {
    return i1b6d6c642ae53b6d75418c14f94432b149eb82d231e9ef4719cabf3dd36e4a84.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) DismissReminder()(*i8bc1b83864742db92b87f3ce691ba7194b9e4dfcdee44bc1f989db1fd24e0d9c.DismissReminderRequestBuilder) {
    return i8bc1b83864742db92b87f3ce691ba7194b9e4dfcdee44bc1f989db1fd24e0d9c.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*ic9d89564024358edace4430f9c467c2085e6b36f70aa7b6fb0605f47881321b4.ExtensionsRequestBuilder) {
    return ic9d89564024358edace4430f9c467c2085e6b36f70aa7b6fb0605f47881321b4.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.events.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*ibf212a370d44f80b700c5b5b75239c5e02548b6abbd3484f7872aa4674bf95c4.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ibf212a370d44f80b700c5b5b75239c5e02548b6abbd3484f7872aa4674bf95c4.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i07aefebef2360a05f1f978a4dacfbfc958f6263953281cb4e4810c24e8fdbe51.ForwardRequestBuilder) {
    return i07aefebef2360a05f1f978a4dacfbfc958f6263953281cb4e4810c24e8fdbe51.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) Instances()(*i8daf6ca6792aac2f5e22afadd880b9416a73f6d66f3d3a4cb4ce49b702beabff.InstancesRequestBuilder) {
    return i8daf6ca6792aac2f5e22afadd880b9416a73f6d66f3d3a4cb4ce49b702beabff.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.events.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*ibe0f4b62b27e66ba9582cbb8b856bbb6bd0e3e8a30c85f69943b3fc01aac856e.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ibe0f4b62b27e66ba9582cbb8b856bbb6bd0e3e8a30c85f69943b3fc01aac856e.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*ic2ca36017987ae7a5daa23b834522d6cb1b251d72fe6d72b471f8ca4ed2d2264.MultiValueExtendedPropertiesRequestBuilder) {
    return ic2ca36017987ae7a5daa23b834522d6cb1b251d72fe6d72b471f8ca4ed2d2264.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.events.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i9facca5460444460406e305a714abf1b6c6c7e0be4c7bedda91eb9b394bd9ede.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i9facca5460444460406e305a714abf1b6c6c7e0be4c7bedda91eb9b394bd9ede.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*icfb49c8c2637abda0a6e0b8db57c259abb3d4ec46f18de110de6cc049d986341.SingleValueExtendedPropertiesRequestBuilder) {
    return icfb49c8c2637abda0a6e0b8db57c259abb3d4ec46f18de110de6cc049d986341.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendar.events.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i49c7e84876761d3a9df15aa2f74f21992b0ba54e8eea9060140c4135b4b76920.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i49c7e84876761d3a9df15aa2f74f21992b0ba54e8eea9060140c4135b4b76920.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*iaafeb0437eaa3b26607893fda20fbdae77ac9d7e2333ed4d5b73883b77acff0a.SnoozeReminderRequestBuilder) {
    return iaafeb0437eaa3b26607893fda20fbdae77ac9d7e2333ed4d5b73883b77acff0a.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i130329fd346118a25e215e9ae08e6acafc2fb545bcc2544394704d98040e0fd5.TentativelyAcceptRequestBuilder) {
    return i130329fd346118a25e215e9ae08e6acafc2fb545bcc2544394704d98040e0fd5.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
