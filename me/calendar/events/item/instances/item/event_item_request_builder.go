package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i115424fb947e5363ae5a2dbcd21aedc22f38dffe0cc73f1ef73b25306b7603d1 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item/snoozereminder"
    i4ecf9972e71ceed18714c2f2e8df53347f0e3eb946fa092de94977ea4fed1ceb "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item/decline"
    i5121a60846fa1d710d99f63cb11af175a28ec5079443325e118d875f5c00ff43 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item/forward"
    i6fcb984531f9e619175ab5500d10abe7f35a19a0222db8119e4558b14210ee7e "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item/accept"
    i7ed3cb920739c72fa84a300b2103c5b0d218576cca12359b662685935f0eaf1e "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item/calendar"
    i896eff89b7229688e940bcb6969ce2d2c0bdb24746c8c9588b2fd26db63028d7 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item/attachments"
    ib0e713445b5a213b2e1bc107fd4c5258be4b72d7d08cae4177f0f252ecc674e9 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item/singlevalueextendedproperties"
    ibf55e0cbfd799c432674550c326d27458a9d46ad44fb6847daaa94f3085266af "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item/cancel"
    ic320338b3b6969f4fefa9b2514352265bf35f14d06f238916bcfa74353b23abe "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item/multivalueextendedproperties"
    icedd85fbcd7041ab524d00f58e3fc8cd64b80a6040a18f66af1097654d14e25a "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item/extensions"
    if70024e0e615a818d126650626f99631a3f58b7c2c14bc1a5a482345208428b3 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item/dismissreminder"
    ife044224236e6edea2f74297331ca97a9b058bbe292a86fada70da05707d7460 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item/tentativelyaccept"
    i1f1103b1f10d19af70cdb5d44f1c4993d2eb251eb003bda9353d58a870ade6da "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item/extensions/item"
    i2879a2ee532e317841d70b52cee430d07751f97864ca66e3c10e502ec78ce73a "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item/multivalueextendedproperties/item"
    i2a39f49bad40b8b085e6b6ade7998f88875082e8dcb30589596db0523283fe95 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item/singlevalueextendedproperties/item"
    i6b2bff425bc13e13fd61c4d59ba9a0d4738b430aa1f1e13e97659caefc058e2d "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item/instances/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*i6fcb984531f9e619175ab5500d10abe7f35a19a0222db8119e4558b14210ee7e.AcceptRequestBuilder) {
    return i6fcb984531f9e619175ab5500d10abe7f35a19a0222db8119e4558b14210ee7e.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i896eff89b7229688e940bcb6969ce2d2c0bdb24746c8c9588b2fd26db63028d7.AttachmentsRequestBuilder) {
    return i896eff89b7229688e940bcb6969ce2d2c0bdb24746c8c9588b2fd26db63028d7.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i6b2bff425bc13e13fd61c4d59ba9a0d4738b430aa1f1e13e97659caefc058e2d.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i6b2bff425bc13e13fd61c4d59ba9a0d4738b430aa1f1e13e97659caefc058e2d.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i7ed3cb920739c72fa84a300b2103c5b0d218576cca12359b662685935f0eaf1e.CalendarRequestBuilder) {
    return i7ed3cb920739c72fa84a300b2103c5b0d218576cca12359b662685935f0eaf1e.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*ibf55e0cbfd799c432674550c326d27458a9d46ad44fb6847daaa94f3085266af.CancelRequestBuilder) {
    return ibf55e0cbfd799c432674550c326d27458a9d46ad44fb6847daaa94f3085266af.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/events/{event_id}/instances/{event_id1}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*i4ecf9972e71ceed18714c2f2e8df53347f0e3eb946fa092de94977ea4fed1ceb.DeclineRequestBuilder) {
    return i4ecf9972e71ceed18714c2f2e8df53347f0e3eb946fa092de94977ea4fed1ceb.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for me
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) DismissReminder()(*if70024e0e615a818d126650626f99631a3f58b7c2c14bc1a5a482345208428b3.DismissReminderRequestBuilder) {
    return if70024e0e615a818d126650626f99631a3f58b7c2c14bc1a5a482345208428b3.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*icedd85fbcd7041ab524d00f58e3fc8cd64b80a6040a18f66af1097654d14e25a.ExtensionsRequestBuilder) {
    return icedd85fbcd7041ab524d00f58e3fc8cd64b80a6040a18f66af1097654d14e25a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i1f1103b1f10d19af70cdb5d44f1c4993d2eb251eb003bda9353d58a870ade6da.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i1f1103b1f10d19af70cdb5d44f1c4993d2eb251eb003bda9353d58a870ade6da.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i5121a60846fa1d710d99f63cb11af175a28ec5079443325e118d875f5c00ff43.ForwardRequestBuilder) {
    return i5121a60846fa1d710d99f63cb11af175a28ec5079443325e118d875f5c00ff43.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Eventable), nil
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ic320338b3b6969f4fefa9b2514352265bf35f14d06f238916bcfa74353b23abe.MultiValueExtendedPropertiesRequestBuilder) {
    return ic320338b3b6969f4fefa9b2514352265bf35f14d06f238916bcfa74353b23abe.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i2879a2ee532e317841d70b52cee430d07751f97864ca66e3c10e502ec78ce73a.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i2879a2ee532e317841d70b52cee430d07751f97864ca66e3c10e502ec78ce73a.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in me
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ib0e713445b5a213b2e1bc107fd4c5258be4b72d7d08cae4177f0f252ecc674e9.SingleValueExtendedPropertiesRequestBuilder) {
    return ib0e713445b5a213b2e1bc107fd4c5258be4b72d7d08cae4177f0f252ecc674e9.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendar.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i2a39f49bad40b8b085e6b6ade7998f88875082e8dcb30589596db0523283fe95.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i2a39f49bad40b8b085e6b6ade7998f88875082e8dcb30589596db0523283fe95.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i115424fb947e5363ae5a2dbcd21aedc22f38dffe0cc73f1ef73b25306b7603d1.SnoozeReminderRequestBuilder) {
    return i115424fb947e5363ae5a2dbcd21aedc22f38dffe0cc73f1ef73b25306b7603d1.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*ife044224236e6edea2f74297331ca97a9b058bbe292a86fada70da05707d7460.TentativelyAcceptRequestBuilder) {
    return ife044224236e6edea2f74297331ca97a9b058bbe292a86fada70da05707d7460.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
