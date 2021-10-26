package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i26e3d572d3ce4fd2f7a9942597d4ec417f0055c7257ff7e19ff1942e96efdada "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/calendarview/item/accept"
    i2f55abcaa4f7436ff7c12805370566b718ca489240475d112654c5101dfa5e14 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/calendarview/item/forward"
    i66c573bbbeba4cbafbf042d4c1af3942eaf35758524eaf7cc2fe8a10f9dc5a9a "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/calendarview/item/dismissreminder"
    i83f9dc3b0ed0c3ef6fe3132386877869e8a1245ca9224e14807a3f941ff992b3 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/calendarview/item/cancel"
    ib8540d03308b08bb61e85c064c8e292638aac60c9564a9e79225b0e5e7825e0e "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/calendarview/item/snoozereminder"
    if31f3f341041d6b9fcd7f18d98d9f390eb4808151bba1b6e87cd9b0fa5751698 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/calendarview/item/decline"
    if979a6c55d8b39c8a9aa26433f2b154ed2a99e41a6c8d58c48cf69c31ef06365 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/calendarview/item/tentativelyaccept"
)

// Builds and executes requests for operations under \me\calendarView\{event-id}\calendar\calendarView\{event-id1}
type EventRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The calendar view for the calendar. Navigation property. Read-only.
type EventRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Select properties to be returned
    Select_escaped []string;
}
func (m *EventRequestBuilder) Accept()(*i26e3d572d3ce4fd2f7a9942597d4ec417f0055c7257ff7e19ff1942e96efdada.AcceptRequestBuilder) {
    return i26e3d572d3ce4fd2f7a9942597d4ec417f0055c7257ff7e19ff1942e96efdada.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i83f9dc3b0ed0c3ef6fe3132386877869e8a1245ca9224e14807a3f941ff992b3.CancelRequestBuilder) {
    return i83f9dc3b0ed0c3ef6fe3132386877869e8a1245ca9224e14807a3f941ff992b3.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/calendarView/{event_id}/calendar/calendarView/{event_id1}{?select}";
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
// The calendar view for the calendar. Navigation property. Read-only.
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
// The calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Decline()(*if31f3f341041d6b9fcd7f18d98d9f390eb4808151bba1b6e87cd9b0fa5751698.DeclineRequestBuilder) {
    return if31f3f341041d6b9fcd7f18d98d9f390eb4808151bba1b6e87cd9b0fa5751698.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) DismissReminder()(*i66c573bbbeba4cbafbf042d4c1af3942eaf35758524eaf7cc2fe8a10f9dc5a9a.DismissReminderRequestBuilder) {
    return i66c573bbbeba4cbafbf042d4c1af3942eaf35758524eaf7cc2fe8a10f9dc5a9a.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i2f55abcaa4f7436ff7c12805370566b718ca489240475d112654c5101dfa5e14.ForwardRequestBuilder) {
    return i2f55abcaa4f7436ff7c12805370566b718ca489240475d112654c5101dfa5e14.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The calendar view for the calendar. Navigation property. Read-only.
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
// The calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) SnoozeReminder()(*ib8540d03308b08bb61e85c064c8e292638aac60c9564a9e79225b0e5e7825e0e.SnoozeReminderRequestBuilder) {
    return ib8540d03308b08bb61e85c064c8e292638aac60c9564a9e79225b0e5e7825e0e.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*if979a6c55d8b39c8a9aa26433f2b154ed2a99e41a6c8d58c48cf69c31ef06365.TentativelyAcceptRequestBuilder) {
    return if979a6c55d8b39c8a9aa26433f2b154ed2a99e41a6c8d58c48cf69c31ef06365.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
