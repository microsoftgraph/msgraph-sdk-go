package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i2d2964f562f1bca8626a91b6ab1e35ad2385f27a304d332970defc2383259b60 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/events/item/accept"
    i5a4a0d4caaab4c3d4e8eeef1ea661c6bccc8fbde0a17a7b113951d8a163f4f51 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/events/item/tentativelyaccept"
    i6583eabdd322186fdeb9534b553d0872ec4dada37eb8ad6813c2e5eec5b579de "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/events/item/cancel"
    i6a74e4de2207e6d56f8066e0e3b5ecde154635d31257b545aa5b1034bad3a092 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/events/item/snoozereminder"
    i91b8c0987ce08a4915f236ac350c8e7882cbc4098b5911f4595b6a2e872dbb0d "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/events/item/forward"
    ic135a15594dc68274729ae7303521a26d41f7a1bb8398edde9cabae5dcadc67a "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/events/item/dismissreminder"
    if52eda4b547abfb1f0ff6f697b2b11d5de90429a754fe2e7ab45440ab679b01e "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/events/item/decline"
)

// Builds and executes requests for operations under \me\calendarView\{event-id}\calendar\events\{event-id1}
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
func (m *EventRequestBuilder) Accept()(*i2d2964f562f1bca8626a91b6ab1e35ad2385f27a304d332970defc2383259b60.AcceptRequestBuilder) {
    return i2d2964f562f1bca8626a91b6ab1e35ad2385f27a304d332970defc2383259b60.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i6583eabdd322186fdeb9534b553d0872ec4dada37eb8ad6813c2e5eec5b579de.CancelRequestBuilder) {
    return i6583eabdd322186fdeb9534b553d0872ec4dada37eb8ad6813c2e5eec5b579de.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/calendarView/{event_id}/calendar/events/{event_id1}{?select}";
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
func (m *EventRequestBuilder) Decline()(*if52eda4b547abfb1f0ff6f697b2b11d5de90429a754fe2e7ab45440ab679b01e.DeclineRequestBuilder) {
    return if52eda4b547abfb1f0ff6f697b2b11d5de90429a754fe2e7ab45440ab679b01e.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) DismissReminder()(*ic135a15594dc68274729ae7303521a26d41f7a1bb8398edde9cabae5dcadc67a.DismissReminderRequestBuilder) {
    return ic135a15594dc68274729ae7303521a26d41f7a1bb8398edde9cabae5dcadc67a.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i91b8c0987ce08a4915f236ac350c8e7882cbc4098b5911f4595b6a2e872dbb0d.ForwardRequestBuilder) {
    return i91b8c0987ce08a4915f236ac350c8e7882cbc4098b5911f4595b6a2e872dbb0d.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) SnoozeReminder()(*i6a74e4de2207e6d56f8066e0e3b5ecde154635d31257b545aa5b1034bad3a092.SnoozeReminderRequestBuilder) {
    return i6a74e4de2207e6d56f8066e0e3b5ecde154635d31257b545aa5b1034bad3a092.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i5a4a0d4caaab4c3d4e8eeef1ea661c6bccc8fbde0a17a7b113951d8a163f4f51.TentativelyAcceptRequestBuilder) {
    return i5a4a0d4caaab4c3d4e8eeef1ea661c6bccc8fbde0a17a7b113951d8a163f4f51.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
