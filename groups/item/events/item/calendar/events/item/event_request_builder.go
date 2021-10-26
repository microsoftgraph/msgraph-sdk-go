package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i35b44b1913943c770c00ab44ed6a2e54f13394f04e4d26f17d29c13a754153d1 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/events/item/dismissreminder"
    i4fff99feb529bb46c0d0ca8a8abb833e2c0cbe2294c7d390cefd23db3957ca27 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/events/item/forward"
    i792255b9849f935465f04d3a10b6e200276b38f41cf2ff7ab2529ab5b360e172 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/events/item/decline"
    i9b8befd4d33ef78b646963d933777df94b88b1bda530296ec0ffff372f368583 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/events/item/snoozereminder"
    ib09926d6635d1d9e91ab40edfee1a8241f91f3ce30e24d67a541c89ffda0f4e2 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/events/item/accept"
    id8eb277815cfe71f400ab449b5653315092c441b52448d3129277f4a9e70fabc "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/events/item/tentativelyaccept"
    if8223036d67d813d86e64f7906c029d458e75d8e3277f4fd4d3ada7eae931b18 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/events/item/cancel"
)

// Builds and executes requests for operations under \groups\{group-id}\events\{event-id}\calendar\events\{event-id1}
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
func (m *EventRequestBuilder) Accept()(*ib09926d6635d1d9e91ab40edfee1a8241f91f3ce30e24d67a541c89ffda0f4e2.AcceptRequestBuilder) {
    return ib09926d6635d1d9e91ab40edfee1a8241f91f3ce30e24d67a541c89ffda0f4e2.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*if8223036d67d813d86e64f7906c029d458e75d8e3277f4fd4d3ada7eae931b18.CancelRequestBuilder) {
    return if8223036d67d813d86e64f7906c029d458e75d8e3277f4fd4d3ada7eae931b18.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/groups/{group_id}/events/{event_id}/calendar/events/{event_id1}{?select}";
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
func (m *EventRequestBuilder) Decline()(*i792255b9849f935465f04d3a10b6e200276b38f41cf2ff7ab2529ab5b360e172.DeclineRequestBuilder) {
    return i792255b9849f935465f04d3a10b6e200276b38f41cf2ff7ab2529ab5b360e172.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) DismissReminder()(*i35b44b1913943c770c00ab44ed6a2e54f13394f04e4d26f17d29c13a754153d1.DismissReminderRequestBuilder) {
    return i35b44b1913943c770c00ab44ed6a2e54f13394f04e4d26f17d29c13a754153d1.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i4fff99feb529bb46c0d0ca8a8abb833e2c0cbe2294c7d390cefd23db3957ca27.ForwardRequestBuilder) {
    return i4fff99feb529bb46c0d0ca8a8abb833e2c0cbe2294c7d390cefd23db3957ca27.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) SnoozeReminder()(*i9b8befd4d33ef78b646963d933777df94b88b1bda530296ec0ffff372f368583.SnoozeReminderRequestBuilder) {
    return i9b8befd4d33ef78b646963d933777df94b88b1bda530296ec0ffff372f368583.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*id8eb277815cfe71f400ab449b5653315092c441b52448d3129277f4a9e70fabc.TentativelyAcceptRequestBuilder) {
    return id8eb277815cfe71f400ab449b5653315092c441b52448d3129277f4a9e70fabc.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
