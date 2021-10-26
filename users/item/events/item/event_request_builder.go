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

// Builds and executes requests for operations under \users\{user-id}\events\{event-id}
type EventRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
type EventRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Select properties to be returned
    Select_escaped []string;
}
func (m *EventRequestBuilder) Accept()(*i7bd7ba2ce48d243a859e20ef25790047a8675315ee1da0126957422798414ec0.AcceptRequestBuilder) {
    return i7bd7ba2ce48d243a859e20ef25790047a8675315ee1da0126957422798414ec0.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*ib5ca28c6c72c1db8b6f37b269c2876fab903b374b0d30cc8404a167730803940.AttachmentsRequestBuilder) {
    return ib5ca28c6c72c1db8b6f37b269c2876fab903b374b0d30cc8404a167730803940.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.attachments.item collection
// Parameters:
//  - id : Unique identifier of the item
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
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/events/{event_id}{?select}";
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
// The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *EventRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *EventRequestBuilder) CreateGetRequestInformation(q func (value *EventRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EventRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *EventRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *EventRequestBuilder) Decline()(*i735b02580c067f7e9c04d97f0df8b9805febcc2badf5811dc6135a7bcbd36880.DeclineRequestBuilder) {
    return i735b02580c067f7e9c04d97f0df8b9805febcc2badf5811dc6135a7bcbd36880.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *EventRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
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
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.extensions.item collection
// Parameters:
//  - id : Unique identifier of the item
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
// The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *EventRequestBuilder) Get(q func (value *EventRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEvent() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event), nil
}
func (m *EventRequestBuilder) Instances()(*i36d3a71fa12eb21299e38ecfe8e5c1c9858e945064a266a1dc3886953114eece.InstancesRequestBuilder) {
    return i36d3a71fa12eb21299e38ecfe8e5c1c9858e945064a266a1dc3886953114eece.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.instances.item collection
// Parameters:
//  - id : Unique identifier of the item
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
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
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
// The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *EventRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*ibc4cc90649e04bb4574982b4a76356636e2914d984eb4131c9ea6e287f13ebdb.SingleValueExtendedPropertiesRequestBuilder) {
    return ibc4cc90649e04bb4574982b4a76356636e2914d984eb4131c9ea6e287f13ebdb.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
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
