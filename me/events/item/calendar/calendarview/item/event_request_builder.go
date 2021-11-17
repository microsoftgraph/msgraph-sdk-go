package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1a4e17d185b1a2b6abccb0defeb08b7409c1244f9402569e860654665a07bd94 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/calendarview/item/cancel"
    i34dd32683c7c9f334bc18a52be150aab75e13253aa28ff9077f5f15fd19d5b10 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/calendarview/item/dismissreminder"
    i4a633db77aff3e86902070cb6c219eadee91c3bd6d4326bed8bbd2ec3c59099d "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/calendarview/item/forward"
    i6a557c407f8eaaf786e2ede8b6448468e46adb4679560e14cbfa122b9dc0bbd3 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/calendarview/item/decline"
    ib2aede831c3065f6b2a5ad6d1b3ad7273a20ace9ef26514b94b5932134fdbdb0 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/calendarview/item/snoozereminder"
    ib917cc04e2b66df73351798ede67fb6975093e90f461e0f39f2ea779869ff82c "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/calendarview/item/accept"
    ica9f9bc668ad52cccdfe99a8ae86eb682912577700483c2302df2e06aa3584a9 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/calendarview/item/tentativelyaccept"
)

// Builds and executes requests for operations under \me\events\{event-id}\calendar\calendarView\{event-id1}
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
// The calendar view for the calendar. Navigation property. Read-only.
type EventRequestBuilderGetQueryParameters struct {
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
func (m *EventRequestBuilder) Accept()(*ib917cc04e2b66df73351798ede67fb6975093e90f461e0f39f2ea779869ff82c.AcceptRequestBuilder) {
    return ib917cc04e2b66df73351798ede67fb6975093e90f461e0f39f2ea779869ff82c.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i1a4e17d185b1a2b6abccb0defeb08b7409c1244f9402569e860654665a07bd94.CancelRequestBuilder) {
    return i1a4e17d185b1a2b6abccb0defeb08b7409c1244f9402569e860654665a07bd94.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/events/{event_id}/calendar/calendarView/{event_id1}{?select}";
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
// The calendar view for the calendar. Navigation property. Read-only.
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
// The calendar view for the calendar. Navigation property. Read-only.
// Parameters:
//  - options : Options for the request
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
// The calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Decline()(*i6a557c407f8eaaf786e2ede8b6448468e46adb4679560e14cbfa122b9dc0bbd3.DeclineRequestBuilder) {
    return i6a557c407f8eaaf786e2ede8b6448468e46adb4679560e14cbfa122b9dc0bbd3.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) DismissReminder()(*i34dd32683c7c9f334bc18a52be150aab75e13253aa28ff9077f5f15fd19d5b10.DismissReminderRequestBuilder) {
    return i34dd32683c7c9f334bc18a52be150aab75e13253aa28ff9077f5f15fd19d5b10.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i4a633db77aff3e86902070cb6c219eadee91c3bd6d4326bed8bbd2ec3c59099d.ForwardRequestBuilder) {
    return i4a633db77aff3e86902070cb6c219eadee91c3bd6d4326bed8bbd2ec3c59099d.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The calendar view for the calendar. Navigation property. Read-only.
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
// The calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) SnoozeReminder()(*ib2aede831c3065f6b2a5ad6d1b3ad7273a20ace9ef26514b94b5932134fdbdb0.SnoozeReminderRequestBuilder) {
    return ib2aede831c3065f6b2a5ad6d1b3ad7273a20ace9ef26514b94b5932134fdbdb0.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*ica9f9bc668ad52cccdfe99a8ae86eb682912577700483c2302df2e06aa3584a9.TentativelyAcceptRequestBuilder) {
    return ica9f9bc668ad52cccdfe99a8ae86eb682912577700483c2302df2e06aa3584a9.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
