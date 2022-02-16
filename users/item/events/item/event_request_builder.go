package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1a3b6643861b0a0caa0ff1eb68e5c488a43aff0aa1f3a06093f6e226fa2b917d "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/dismissreminder"
    i32c37b1c04b5be6ff9366461bd21b5fde755e861ca804c535263e878e67f6c27 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/forward"
    i36d3a71fa12eb21299e38ecfe8e5c1c9858e945064a266a1dc3886953114eece "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances"
    i40517a193f3b25c5466a4fb68f9a11b390590bed4b80557218b6fa94bba2bb03 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/multivalueextendedproperties"
    i5823ac8df2710b8a2e2cffa605921bb14b14de21e070c27b7581b8e8a5ba871e "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/cancel"
    i70149e4bcbdc49f55d93a41203bb62f881cf96a4136c6be6de44d5750c55cbec "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/tentativelyaccept"
    i735b02580c067f7e9c04d97f0df8b9805febcc2badf5811dc6135a7bcbd36880 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/decline"
    i73fdb8c241b8badf777c9058cb4d1f11e3c5329111efcaf2510b52537daa9874 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/snoozereminder"
    i7bd7ba2ce48d243a859e20ef25790047a8675315ee1da0126957422798414ec0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/accept"
    i91c0a715e42a5cbf02de3870dd658b8c3716027385e23055057583d5a1b3597e "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/extensions"
    ib5ca28c6c72c1db8b6f37b269c2876fab903b374b0d30cc8404a167730803940 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/attachments"
    ibc4cc90649e04bb4574982b4a76356636e2914d984eb4131c9ea6e287f13ebdb "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/singlevalueextendedproperties"
    idd1ffd3d2adc7fab21de995ec4b6949ca9bb47b13c5e7cc18129c1449dc2651a "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/calendar"
    i58869d35545fe6338f6dd5b999816eceba5ff2ae808d4a6636dd981dc0025fa3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/multivalueextendedproperties/item"
    i95164ba391501ae53794b155baaed0825128639bb413adda940b803c1b5ae14a "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/instances/item"
    ief46ce2c42a9faa7186bd9c6c9a601dc9f0722341a69677d6743c991650b60fd "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/attachments/item"
    if6d2f1f347111fdd7f5b334372ce95e9ea9e20681fbc3bd7a3a39fdb53b19ca4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/extensions/item"
    ifccd98b65865537ee9ddb082f187a02b159c058c0cdbeb4ffa28207f93bca168 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/singlevalueextendedproperties/item"
)

// EventRequestBuilder builds and executes requests for operations under \users\{user-id}\events\{event-id}
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
// EventRequestBuilderGetQueryParameters the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
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
func (m *EventRequestBuilder) Accept()(*i7bd7ba2ce48d243a859e20ef25790047a8675315ee1da0126957422798414ec0.AcceptRequestBuilder) {
    return i7bd7ba2ce48d243a859e20ef25790047a8675315ee1da0126957422798414ec0.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*ib5ca28c6c72c1db8b6f37b269c2876fab903b374b0d30cc8404a167730803940.AttachmentsRequestBuilder) {
    return ib5ca28c6c72c1db8b6f37b269c2876fab903b374b0d30cc8404a167730803940.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*ief46ce2c42a9faa7186bd9c6c9a601dc9f0722341a69677d6743c991650b60fd.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ief46ce2c42a9faa7186bd9c6c9a601dc9f0722341a69677d6743c991650b60fd.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*idd1ffd3d2adc7fab21de995ec4b6949ca9bb47b13c5e7cc18129c1449dc2651a.CalendarRequestBuilder) {
    return idd1ffd3d2adc7fab21de995ec4b6949ca9bb47b13c5e7cc18129c1449dc2651a.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i5823ac8df2710b8a2e2cffa605921bb14b14de21e070c27b7581b8e8a5ba871e.CancelRequestBuilder) {
    return i5823ac8df2710b8a2e2cffa605921bb14b14de21e070c27b7581b8e8a5ba871e.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/events/{event_id}{?select}";
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
// CreateDeleteRequestInformation the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
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
// CreateGetRequestInformation the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
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
// CreatePatchRequestInformation the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
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
func (m *EventRequestBuilder) Decline()(*i735b02580c067f7e9c04d97f0df8b9805febcc2badf5811dc6135a7bcbd36880.DeclineRequestBuilder) {
    return i735b02580c067f7e9c04d97f0df8b9805febcc2badf5811dc6135a7bcbd36880.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
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
func (m *EventRequestBuilder) DismissReminder()(*i1a3b6643861b0a0caa0ff1eb68e5c488a43aff0aa1f3a06093f6e226fa2b917d.DismissReminderRequestBuilder) {
    return i1a3b6643861b0a0caa0ff1eb68e5c488a43aff0aa1f3a06093f6e226fa2b917d.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i91c0a715e42a5cbf02de3870dd658b8c3716027385e23055057583d5a1b3597e.ExtensionsRequestBuilder) {
    return i91c0a715e42a5cbf02de3870dd658b8c3716027385e23055057583d5a1b3597e.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*if6d2f1f347111fdd7f5b334372ce95e9ea9e20681fbc3bd7a3a39fdb53b19ca4.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return if6d2f1f347111fdd7f5b334372ce95e9ea9e20681fbc3bd7a3a39fdb53b19ca4.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i32c37b1c04b5be6ff9366461bd21b5fde755e861ca804c535263e878e67f6c27.ForwardRequestBuilder) {
    return i32c37b1c04b5be6ff9366461bd21b5fde755e861ca804c535263e878e67f6c27.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventRequestBuilder) Get(options *EventRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, error) {
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
func (m *EventRequestBuilder) Instances()(*i36d3a71fa12eb21299e38ecfe8e5c1c9858e945064a266a1dc3886953114eece.InstancesRequestBuilder) {
    return i36d3a71fa12eb21299e38ecfe8e5c1c9858e945064a266a1dc3886953114eece.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*i95164ba391501ae53794b155baaed0825128639bb413adda940b803c1b5ae14a.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i95164ba391501ae53794b155baaed0825128639bb413adda940b803c1b5ae14a.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i40517a193f3b25c5466a4fb68f9a11b390590bed4b80557218b6fa94bba2bb03.MultiValueExtendedPropertiesRequestBuilder) {
    return i40517a193f3b25c5466a4fb68f9a11b390590bed4b80557218b6fa94bba2bb03.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i58869d35545fe6338f6dd5b999816eceba5ff2ae808d4a6636dd981dc0025fa3.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i58869d35545fe6338f6dd5b999816eceba5ff2ae808d4a6636dd981dc0025fa3.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*ibc4cc90649e04bb4574982b4a76356636e2914d984eb4131c9ea6e287f13ebdb.SingleValueExtendedPropertiesRequestBuilder) {
    return ibc4cc90649e04bb4574982b4a76356636e2914d984eb4131c9ea6e287f13ebdb.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ifccd98b65865537ee9ddb082f187a02b159c058c0cdbeb4ffa28207f93bca168.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ifccd98b65865537ee9ddb082f187a02b159c058c0cdbeb4ffa28207f93bca168.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i73fdb8c241b8badf777c9058cb4d1f11e3c5329111efcaf2510b52537daa9874.SnoozeReminderRequestBuilder) {
    return i73fdb8c241b8badf777c9058cb4d1f11e3c5329111efcaf2510b52537daa9874.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i70149e4bcbdc49f55d93a41203bb62f881cf96a4136c6be6de44d5750c55cbec.TentativelyAcceptRequestBuilder) {
    return i70149e4bcbdc49f55d93a41203bb62f881cf96a4136c6be6de44d5750c55cbec.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
