package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i15a1cf85dcc7eeb107a61041c861c7e258520efb3c3af054aac161a329f2c15e "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/snoozereminder"
    i251db03f59cc75f48f84e7245a0dc6f3de00400380451646430f584012111760 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/tentativelyaccept"
    i29b0d509f093b40c14a151adbb10e0aab35cfe53c1d405745d4c2976d34fd389 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/decline"
    i29c1c4fc5baeb4284286c65b18c5c49cd5481668cdc91c23a63a6802353fae13 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/multivalueextendedproperties"
    i2d66a35c35d6c7f3f7ea7c5db7ef1dabea98b07af7035f0b05c692c32145e5ee "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/singlevalueextendedproperties"
    i352926d1ca610ee05029220da93f8400faee6b385271b5f15acb9c1b3d7d74f4 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/extensions"
    i3d78d5747cc294b4c322bc7aeeca2cc53380d9b4bad2965e21c621c160c79fa4 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/attachments"
    i72a6e89f7ed2091ef6b627975e8427d3601c1922304a5d1f43cc3dffdda9913e "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/forward"
    i90c9ce46a5ec66dcd1b80481bc789fbb3028677cbc3c32edfae1bc6d1d519a54 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/dismissreminder"
    i95bb8b9f1540cf7b97dc6ccf9364b1acc0b031644c22f4e03491a1700a790543 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/cancel"
    ib7a7049c46e443ff9e41b4c969d1d2ce89b58a39ddbce1e602cb74d15077df88 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/calendar"
    ie17a7797e2d679fd92ef43065db9290c8683a8f64c3279969a5eb21848d09d73 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/instances"
    if0ef0ca8fe12aeb76b29f3a545e38f8de508aed0b34c5517d3799e69f33f51f5 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/accept"
    i442111c87f659072c1c2f074352a85ebd625c1e18c68b3077096aabd48a95461 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/attachments/item"
    i51b055bd1bde3fd826011eb86802dc300fac30381c9947123d8a124cf68f9c84 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/extensions/item"
    ibf36b2390d0bb0a736c630f7f1cc55af3ed6227a90b2fbb973dc5bd7ac937b7b "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/singlevalueextendedproperties/item"
    ic6feaf2955661d48f6f8c3cf286ef82befbf22fa19c47374ebd2880710008c5d "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item"
    ica66966e5d0c3577314d4dcc18c27d0b23efe8e86f729b91df3bd4c3e6ecedb0 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item/multivalueextendedproperties/item"
)

// Builds and executes requests for operations under \me\calendarGroups\{calendarGroup-id}\calendars\{calendar-id}\events\{event-id}
type EventRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The events in the calendar. Navigation property. Read-only.
type EventRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Select properties to be returned
    Select_escaped []string;
}
func (m *EventRequestBuilder) Accept()(*if0ef0ca8fe12aeb76b29f3a545e38f8de508aed0b34c5517d3799e69f33f51f5.AcceptRequestBuilder) {
    return if0ef0ca8fe12aeb76b29f3a545e38f8de508aed0b34c5517d3799e69f33f51f5.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i3d78d5747cc294b4c322bc7aeeca2cc53380d9b4bad2965e21c621c160c79fa4.AttachmentsRequestBuilder) {
    return i3d78d5747cc294b4c322bc7aeeca2cc53380d9b4bad2965e21c621c160c79fa4.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.events.item.attachments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) AttachmentsById(id string)(*i442111c87f659072c1c2f074352a85ebd625c1e18c68b3077096aabd48a95461.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i442111c87f659072c1c2f074352a85ebd625c1e18c68b3077096aabd48a95461.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*ib7a7049c46e443ff9e41b4c969d1d2ce89b58a39ddbce1e602cb74d15077df88.CalendarRequestBuilder) {
    return ib7a7049c46e443ff9e41b4c969d1d2ce89b58a39ddbce1e602cb74d15077df88.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i95bb8b9f1540cf7b97dc6ccf9364b1acc0b031644c22f4e03491a1700a790543.CancelRequestBuilder) {
    return i95bb8b9f1540cf7b97dc6ccf9364b1acc0b031644c22f4e03491a1700a790543.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}/events/{event_id}{?select}";
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
// The events in the calendar. Navigation property. Read-only.
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
// The events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Decline()(*i29b0d509f093b40c14a151adbb10e0aab35cfe53c1d405745d4c2976d34fd389.DeclineRequestBuilder) {
    return i29b0d509f093b40c14a151adbb10e0aab35cfe53c1d405745d4c2976d34fd389.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) DismissReminder()(*i90c9ce46a5ec66dcd1b80481bc789fbb3028677cbc3c32edfae1bc6d1d519a54.DismissReminderRequestBuilder) {
    return i90c9ce46a5ec66dcd1b80481bc789fbb3028677cbc3c32edfae1bc6d1d519a54.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i352926d1ca610ee05029220da93f8400faee6b385271b5f15acb9c1b3d7d74f4.ExtensionsRequestBuilder) {
    return i352926d1ca610ee05029220da93f8400faee6b385271b5f15acb9c1b3d7d74f4.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.events.item.extensions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) ExtensionsById(id string)(*i51b055bd1bde3fd826011eb86802dc300fac30381c9947123d8a124cf68f9c84.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i51b055bd1bde3fd826011eb86802dc300fac30381c9947123d8a124cf68f9c84.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i72a6e89f7ed2091ef6b627975e8427d3601c1922304a5d1f43cc3dffdda9913e.ForwardRequestBuilder) {
    return i72a6e89f7ed2091ef6b627975e8427d3601c1922304a5d1f43cc3dffdda9913e.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Instances()(*ie17a7797e2d679fd92ef43065db9290c8683a8f64c3279969a5eb21848d09d73.InstancesRequestBuilder) {
    return ie17a7797e2d679fd92ef43065db9290c8683a8f64c3279969a5eb21848d09d73.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.events.item.instances.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) InstancesById(id string)(*ic6feaf2955661d48f6f8c3cf286ef82befbf22fa19c47374ebd2880710008c5d.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ic6feaf2955661d48f6f8c3cf286ef82befbf22fa19c47374ebd2880710008c5d.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i29c1c4fc5baeb4284286c65b18c5c49cd5481668cdc91c23a63a6802353fae13.MultiValueExtendedPropertiesRequestBuilder) {
    return i29c1c4fc5baeb4284286c65b18c5c49cd5481668cdc91c23a63a6802353fae13.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.events.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ica66966e5d0c3577314d4dcc18c27d0b23efe8e86f729b91df3bd4c3e6ecedb0.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ica66966e5d0c3577314d4dcc18c27d0b23efe8e86f729b91df3bd4c3e6ecedb0.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i2d66a35c35d6c7f3f7ea7c5db7ef1dabea98b07af7035f0b05c692c32145e5ee.SingleValueExtendedPropertiesRequestBuilder) {
    return i2d66a35c35d6c7f3f7ea7c5db7ef1dabea98b07af7035f0b05c692c32145e5ee.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item.calendars.item.events.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ibf36b2390d0bb0a736c630f7f1cc55af3ed6227a90b2fbb973dc5bd7ac937b7b.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ibf36b2390d0bb0a736c630f7f1cc55af3ed6227a90b2fbb973dc5bd7ac937b7b.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i15a1cf85dcc7eeb107a61041c861c7e258520efb3c3af054aac161a329f2c15e.SnoozeReminderRequestBuilder) {
    return i15a1cf85dcc7eeb107a61041c861c7e258520efb3c3af054aac161a329f2c15e.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i251db03f59cc75f48f84e7245a0dc6f3de00400380451646430f584012111760.TentativelyAcceptRequestBuilder) {
    return i251db03f59cc75f48f84e7245a0dc6f3de00400380451646430f584012111760.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
