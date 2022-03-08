package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i224a43a6f6c22839104e3af8e09c25d6ee5653d0bf039c8cc955d19054af9239 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item/calendar"
    i3b73bf1d367d8d58b43286e8815e65b652bdd9632a043961e96cdd4f21a4a456 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item/attachments"
    i62e84548f3d8b139eec410acddc1a9ec548d38f53c3adf0accda54573c280237 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item/accept"
    i6a76f7e1f0e5de6b644f66577ae554f00c63758d9839dbac747e2a975de2c062 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item/extensions"
    i708657df116b55afae5c575793f00bc62083de220954f2b30070d94329ad4c0b "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item/cancel"
    i78fa73479b6675db5fd8062a94af6edb28ce40ca11518c2cffa6b9e6892a94df "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item/tentativelyaccept"
    ia2d1e74d0c475d24b54809e71da0e1ba8c25369b3fc198918003686d5c13a3b8 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item/forward"
    iba5221e61e3b660e493309a9e0dc226fb7a9236ea0382bc762f0b58feb7753c5 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item/singlevalueextendedproperties"
    ic1bd788757d48993d62ddaea6350b29916d32a63078388e6faeb72dc45a05d7f "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item/multivalueextendedproperties"
    id303083c5a91f2be0053b04d5e3553e1dfe384bb4839cb4b08f8664a109d216f "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item/dismissreminder"
    id888541b36ea9db1d12cf8e62e960c429547d5e0db473794ec8fc9bcde6b3301 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item/decline"
    idf7d1d2c96845dbdba55aa4449dfe108112b1f0b7ff81ef7da936f0a0f979647 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item/snoozereminder"
    i122f3095d5e5a9401b8d3bf34749578d52cc4f098effb0dbeab87f3b7c74cb97 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item/singlevalueextendedproperties/item"
    i67c4f1ef75a41a5c6ae05d3b85e11b1c45c47f9e8bd5361a63b25adb073a82d3 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item/multivalueextendedproperties/item"
    ia34394981640465160cc339a2f60b56da7d55e5f7dd248039a8790e1081043f6 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item/attachments/item"
    ic06690588d7937c8b2811db52aac78b63b86909ef55badd8436b8736211d0c1f "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item/instances/item/extensions/item"
)

// EventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
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
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// EventItemRequestBuilderPatchOptions options for Patch
type EventItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Eventable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventItemRequestBuilder) Accept()(*i62e84548f3d8b139eec410acddc1a9ec548d38f53c3adf0accda54573c280237.AcceptRequestBuilder) {
    return i62e84548f3d8b139eec410acddc1a9ec548d38f53c3adf0accda54573c280237.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i3b73bf1d367d8d58b43286e8815e65b652bdd9632a043961e96cdd4f21a4a456.AttachmentsRequestBuilder) {
    return i3b73bf1d367d8d58b43286e8815e65b652bdd9632a043961e96cdd4f21a4a456.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ia34394981640465160cc339a2f60b56da7d55e5f7dd248039a8790e1081043f6.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ia34394981640465160cc339a2f60b56da7d55e5f7dd248039a8790e1081043f6.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i224a43a6f6c22839104e3af8e09c25d6ee5653d0bf039c8cc955d19054af9239.CalendarRequestBuilder) {
    return i224a43a6f6c22839104e3af8e09c25d6ee5653d0bf039c8cc955d19054af9239.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i708657df116b55afae5c575793f00bc62083de220954f2b30070d94329ad4c0b.CancelRequestBuilder) {
    return i708657df116b55afae5c575793f00bc62083de220954f2b30070d94329ad4c0b.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar_id}/events/{event_id}/instances/{event_id1}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property instances for me
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
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property instances in me
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
func (m *EventItemRequestBuilder) Decline()(*id888541b36ea9db1d12cf8e62e960c429547d5e0db473794ec8fc9bcde6b3301.DeclineRequestBuilder) {
    return id888541b36ea9db1d12cf8e62e960c429547d5e0db473794ec8fc9bcde6b3301.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for me
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) DismissReminder()(*id303083c5a91f2be0053b04d5e3553e1dfe384bb4839cb4b08f8664a109d216f.DismissReminderRequestBuilder) {
    return id303083c5a91f2be0053b04d5e3553e1dfe384bb4839cb4b08f8664a109d216f.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i6a76f7e1f0e5de6b644f66577ae554f00c63758d9839dbac747e2a975de2c062.ExtensionsRequestBuilder) {
    return i6a76f7e1f0e5de6b644f66577ae554f00c63758d9839dbac747e2a975de2c062.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ic06690588d7937c8b2811db52aac78b63b86909ef55badd8436b8736211d0c1f.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ic06690588d7937c8b2811db52aac78b63b86909ef55badd8436b8736211d0c1f.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*ia2d1e74d0c475d24b54809e71da0e1ba8c25369b3fc198918003686d5c13a3b8.ForwardRequestBuilder) {
    return ia2d1e74d0c475d24b54809e71da0e1ba8c25369b3fc198918003686d5c13a3b8.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Eventable), nil
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ic1bd788757d48993d62ddaea6350b29916d32a63078388e6faeb72dc45a05d7f.MultiValueExtendedPropertiesRequestBuilder) {
    return ic1bd788757d48993d62ddaea6350b29916d32a63078388e6faeb72dc45a05d7f.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i67c4f1ef75a41a5c6ae05d3b85e11b1c45c47f9e8bd5361a63b25adb073a82d3.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i67c4f1ef75a41a5c6ae05d3b85e11b1c45c47f9e8bd5361a63b25adb073a82d3.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in me
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*iba5221e61e3b660e493309a9e0dc226fb7a9236ea0382bc762f0b58feb7753c5.SingleValueExtendedPropertiesRequestBuilder) {
    return iba5221e61e3b660e493309a9e0dc226fb7a9236ea0382bc762f0b58feb7753c5.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i122f3095d5e5a9401b8d3bf34749578d52cc4f098effb0dbeab87f3b7c74cb97.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i122f3095d5e5a9401b8d3bf34749578d52cc4f098effb0dbeab87f3b7c74cb97.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*idf7d1d2c96845dbdba55aa4449dfe108112b1f0b7ff81ef7da936f0a0f979647.SnoozeReminderRequestBuilder) {
    return idf7d1d2c96845dbdba55aa4449dfe108112b1f0b7ff81ef7da936f0a0f979647.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i78fa73479b6675db5fd8062a94af6edb28ce40ca11518c2cffa6b9e6892a94df.TentativelyAcceptRequestBuilder) {
    return i78fa73479b6675db5fd8062a94af6edb28ce40ca11518c2cffa6b9e6892a94df.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
