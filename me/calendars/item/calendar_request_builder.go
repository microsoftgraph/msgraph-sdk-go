package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i08263b120d4a64d516da1a9a009102f68f1ea3b72038e62049453064f3520daa "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/getschedule"
    i5b92fe7490f3adfa9b9eb8b7b27fcf837f66e0bf11a120b311cab89a7ab5ccb4 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/multivalueextendedproperties"
    i636df8302013a50dd4f709806f91e209c2a39540a8c8acb8ff6d284173bc9611 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events"
    i790abfff1579d62421ba888c419e99fcdb7f3342ca621795f1d4878bdf41107d "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarpermissions"
    i8d0a3a589be44f5a571e77f994200094b296ec5e5a51363c20dd536fdd042980 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/singlevalueextendedproperties"
    ib60e8318b810b2b05ebbda7d142451b6e905523a0d513153c156b9da64697fee "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview"
    if635722f2fe253209876c19e8ff9c2f5ed8ce31cb7fb3a6ef7c95fc4fa600390 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/allowedcalendarsharingroleswithuser"
    i0159be40cebef16f606f3683bed2c527ac757b7c0595a04fd5a7980fc524e986 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/multivalueextendedproperties/item"
    i5eeae220666e7414f299626922168ed676e16d02eb4326b487e9a593c4ea62b5 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarpermissions/item"
    i72d2962cc7d6783d327cd6ac936abb7f3a115f64393795d9f16cbed9c34fbc1a "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/events/item"
    i7441dd38f6817f39599227c0b3a93087e04dc084baa163401adef30ac3874560 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/singlevalueextendedproperties/item"
    if4d2b8a420506859db7a9f0c9c655e7bda330951aa669e4ab40546cf6bb9ce22 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item"
)

// Builds and executes requests for operations under \me\calendars\{calendar-id}
type CalendarRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The user's calendars. Read-only. Nullable.
type CalendarRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Select properties to be returned
    Select_escaped []string;
}
// Builds and executes requests for operations under \me\calendars\{calendar-id}\microsoft.graph.allowedCalendarSharingRoles(User='{User}')
// Parameters:
//  - User : Usage: User={User}
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*if635722f2fe253209876c19e8ff9c2f5ed8ce31cb7fb3a6ef7c95fc4fa600390.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return if635722f2fe253209876c19e8ff9c2f5ed8ce31cb7fb3a6ef7c95fc4fa600390.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*i790abfff1579d62421ba888c419e99fcdb7f3342ca621795f1d4878bdf41107d.CalendarPermissionsRequestBuilder) {
    return i790abfff1579d62421ba888c419e99fcdb7f3342ca621795f1d4878bdf41107d.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.calendarPermissions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i5eeae220666e7414f299626922168ed676e16d02eb4326b487e9a593c4ea62b5.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i5eeae220666e7414f299626922168ed676e16d02eb4326b487e9a593c4ea62b5.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*ib60e8318b810b2b05ebbda7d142451b6e905523a0d513153c156b9da64697fee.CalendarViewRequestBuilder) {
    return ib60e8318b810b2b05ebbda7d142451b6e905523a0d513153c156b9da64697fee.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.calendarView.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*if4d2b8a420506859db7a9f0c9c655e7bda330951aa669e4ab40546cf6bb9ce22.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return if4d2b8a420506859db7a9f0c9c655e7bda330951aa669e4ab40546cf6bb9ce22.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new CalendarRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/calendars/{calendar_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new CalendarRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCalendarRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarRequestBuilderInternal(urlParams, requestAdapter)
}
// The user's calendars. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *CalendarRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The user's calendars. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *CalendarRequestBuilder) CreateGetRequestInformation(q func (value *CalendarRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(CalendarRequestBuilderGetQueryParameters)
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
// The user's calendars. Read-only. Nullable.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *CalendarRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The user's calendars. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *CalendarRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *CalendarRequestBuilder) Events()(*i636df8302013a50dd4f709806f91e209c2a39540a8c8acb8ff6d284173bc9611.EventsRequestBuilder) {
    return i636df8302013a50dd4f709806f91e209c2a39540a8c8acb8ff6d284173bc9611.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.events.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) EventsById(id string)(*i72d2962cc7d6783d327cd6ac936abb7f3a115f64393795d9f16cbed9c34fbc1a.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i72d2962cc7d6783d327cd6ac936abb7f3a115f64393795d9f16cbed9c34fbc1a.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The user's calendars. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *CalendarRequestBuilder) Get(q func (value *CalendarRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewCalendar() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar), nil
}
func (m *CalendarRequestBuilder) GetSchedule()(*i08263b120d4a64d516da1a9a009102f68f1ea3b72038e62049453064f3520daa.GetScheduleRequestBuilder) {
    return i08263b120d4a64d516da1a9a009102f68f1ea3b72038e62049453064f3520daa.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i5b92fe7490f3adfa9b9eb8b7b27fcf837f66e0bf11a120b311cab89a7ab5ccb4.MultiValueExtendedPropertiesRequestBuilder) {
    return i5b92fe7490f3adfa9b9eb8b7b27fcf837f66e0bf11a120b311cab89a7ab5ccb4.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0159be40cebef16f606f3683bed2c527ac757b7c0595a04fd5a7980fc524e986.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i0159be40cebef16f606f3683bed2c527ac757b7c0595a04fd5a7980fc524e986.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The user's calendars. Read-only. Nullable.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *CalendarRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i8d0a3a589be44f5a571e77f994200094b296ec5e5a51363c20dd536fdd042980.SingleValueExtendedPropertiesRequestBuilder) {
    return i8d0a3a589be44f5a571e77f994200094b296ec5e5a51363c20dd536fdd042980.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7441dd38f6817f39599227c0b3a93087e04dc084baa163401adef30ac3874560.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i7441dd38f6817f39599227c0b3a93087e04dc084baa163401adef30ac3874560.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
