package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0f87642757745921f13e8e7c175214c33a54594efbdd2d330281f077b7de312c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/events/item/accept"
    i3277a4355e3576f9710f96152dc2d5f4897caae3316b7a8a1fd800fc650cca42 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/events/item/forward"
    i359d9014ba05f4ce00b0dddff9408f8a5eed5bd038a83eaa2d2ec7710629b33d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/events/item/decline"
    i3c9bd42eccac418e1a5f66b7296a146c480c8d4006c973b42c7138f4d4ccc590 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/events/item/dismissreminder"
    icacfd755f37f31d187c0d0e2a3fd3c8754366bd254c5bad9284f1a6576a70a07 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/events/item/snoozereminder"
    icf15ef86ccd7cf9f7cbf8d11edfb6139a96b4d9ab142256177402d9c687221f8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/events/item/tentativelyaccept"
    id1eb21aeb9d6082148818bf0bdc0db3654298e4a9eeb0f847cdeef841aa7fc2e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/events/item/cancel"
)

// Builds and executes requests for operations under \groups\{group-id}\calendarView\{event-id}\calendar\events\{event-id1}
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
func (m *EventRequestBuilder) Accept()(*i0f87642757745921f13e8e7c175214c33a54594efbdd2d330281f077b7de312c.AcceptRequestBuilder) {
    return i0f87642757745921f13e8e7c175214c33a54594efbdd2d330281f077b7de312c.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*id1eb21aeb9d6082148818bf0bdc0db3654298e4a9eeb0f847cdeef841aa7fc2e.CancelRequestBuilder) {
    return id1eb21aeb9d6082148818bf0bdc0db3654298e4a9eeb0f847cdeef841aa7fc2e.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/groups/{group_id}/calendarView/{event_id}/calendar/events/{event_id1}{?select}";
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
func (m *EventRequestBuilder) Decline()(*i359d9014ba05f4ce00b0dddff9408f8a5eed5bd038a83eaa2d2ec7710629b33d.DeclineRequestBuilder) {
    return i359d9014ba05f4ce00b0dddff9408f8a5eed5bd038a83eaa2d2ec7710629b33d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) DismissReminder()(*i3c9bd42eccac418e1a5f66b7296a146c480c8d4006c973b42c7138f4d4ccc590.DismissReminderRequestBuilder) {
    return i3c9bd42eccac418e1a5f66b7296a146c480c8d4006c973b42c7138f4d4ccc590.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i3277a4355e3576f9710f96152dc2d5f4897caae3316b7a8a1fd800fc650cca42.ForwardRequestBuilder) {
    return i3277a4355e3576f9710f96152dc2d5f4897caae3316b7a8a1fd800fc650cca42.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) SnoozeReminder()(*icacfd755f37f31d187c0d0e2a3fd3c8754366bd254c5bad9284f1a6576a70a07.SnoozeReminderRequestBuilder) {
    return icacfd755f37f31d187c0d0e2a3fd3c8754366bd254c5bad9284f1a6576a70a07.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*icf15ef86ccd7cf9f7cbf8d11edfb6139a96b4d9ab142256177402d9c687221f8.TentativelyAcceptRequestBuilder) {
    return icf15ef86ccd7cf9f7cbf8d11edfb6139a96b4d9ab142256177402d9c687221f8.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
