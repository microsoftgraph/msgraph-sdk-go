package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i08cbf44ccd2006085019ad7f1db8b25b151f76f2226572e379b633b6bb8bfd0c "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/calendar/getschedule"
    i1db1756c9dc1485dd2b0664968ced55c60fb6cedf15a9ad4465685c39a873c3c "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/calendar/multivalueextendedproperties"
    i219badc27beeb66203e54e1c9e7275884fabf7d0a341e0fed598c056b42bc949 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/calendar/calendarpermissions"
    i27ab0705e58d5dd61a01392bcb5538edd5c5700f3c9bd903d8683b8275043898 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/calendar/allowedcalendarsharingroleswithuser"
    i4737ccf6c4de9b43ff9f15bf803182421cb705085720cbf9f8a6a7af0a794d4e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/calendar/singlevalueextendedproperties"
    i8389f88bdb4ee4f62ee7a0604d9dcab46b3d93897658f8841cda6dee48bcfac4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/calendar/events"
    ie8992c33461c5dd5f1492fc7f025c205fc80302921cd2a20f5486e4f347c9297 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/calendar/calendarview"
    i04f33d6865a19224c6c9fa20c6baa0238e921bdc3a0d4bd6cc9a7e088cc9f193 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/calendar/calendarview/item"
    i0c04a96c97affe987eaae7b174f477b5ba285b8f2dd6d566e40ab0f8064b7018 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/calendar/singlevalueextendedproperties/item"
    i2a394921f13c5d351ad74acea56ad02a0071d76166d31f229272c3c015465589 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/calendar/events/item"
    id4c9062372ada07ecb88af216f6d03b793ae55cff5fe79d4af913328e16639ea "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/calendar/calendarpermissions/item"
    if7609ad3ebcae15de5b2163c9709b7d734c9ff6a29512410f6c64114f4220070 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendarview/item/calendar/multivalueextendedproperties/item"
)

// Builds and executes requests for operations under \users\{user-id}\calendarView\{event-id}\calendar
type CalendarRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The calendar that contains the event. Navigation property. Read-only.
type CalendarRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Select properties to be returned
    Select_escaped []string;
}
// Builds and executes requests for operations under \users\{user-id}\calendarView\{event-id}\calendar\microsoft.graph.allowedCalendarSharingRoles(User='{User}')
// Parameters:
//  - User : Usage: User={User}
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i27ab0705e58d5dd61a01392bcb5538edd5c5700f3c9bd903d8683b8275043898.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i27ab0705e58d5dd61a01392bcb5538edd5c5700f3c9bd903d8683b8275043898.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*i219badc27beeb66203e54e1c9e7275884fabf7d0a341e0fed598c056b42bc949.CalendarPermissionsRequestBuilder) {
    return i219badc27beeb66203e54e1c9e7275884fabf7d0a341e0fed598c056b42bc949.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarView.item.calendar.calendarPermissions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*id4c9062372ada07ecb88af216f6d03b793ae55cff5fe79d4af913328e16639ea.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return id4c9062372ada07ecb88af216f6d03b793ae55cff5fe79d4af913328e16639ea.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*ie8992c33461c5dd5f1492fc7f025c205fc80302921cd2a20f5486e4f347c9297.CalendarViewRequestBuilder) {
    return ie8992c33461c5dd5f1492fc7f025c205fc80302921cd2a20f5486e4f347c9297.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarView.item.calendar.calendarView.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i04f33d6865a19224c6c9fa20c6baa0238e921bdc3a0d4bd6cc9a7e088cc9f193.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i04f33d6865a19224c6c9fa20c6baa0238e921bdc3a0d4bd6cc9a7e088cc9f193.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new CalendarRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/calendarView/{event_id}/calendar{?select}";
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
// The calendar that contains the event. Navigation property. Read-only.
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
// The calendar that contains the event. Navigation property. Read-only.
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
// The calendar that contains the event. Navigation property. Read-only.
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
// The calendar that contains the event. Navigation property. Read-only.
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
func (m *CalendarRequestBuilder) Events()(*i8389f88bdb4ee4f62ee7a0604d9dcab46b3d93897658f8841cda6dee48bcfac4.EventsRequestBuilder) {
    return i8389f88bdb4ee4f62ee7a0604d9dcab46b3d93897658f8841cda6dee48bcfac4.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarView.item.calendar.events.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) EventsById(id string)(*i2a394921f13c5d351ad74acea56ad02a0071d76166d31f229272c3c015465589.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i2a394921f13c5d351ad74acea56ad02a0071d76166d31f229272c3c015465589.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The calendar that contains the event. Navigation property. Read-only.
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
func (m *CalendarRequestBuilder) GetSchedule()(*i08cbf44ccd2006085019ad7f1db8b25b151f76f2226572e379b633b6bb8bfd0c.GetScheduleRequestBuilder) {
    return i08cbf44ccd2006085019ad7f1db8b25b151f76f2226572e379b633b6bb8bfd0c.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i1db1756c9dc1485dd2b0664968ced55c60fb6cedf15a9ad4465685c39a873c3c.MultiValueExtendedPropertiesRequestBuilder) {
    return i1db1756c9dc1485dd2b0664968ced55c60fb6cedf15a9ad4465685c39a873c3c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarView.item.calendar.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*if7609ad3ebcae15de5b2163c9709b7d734c9ff6a29512410f6c64114f4220070.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return if7609ad3ebcae15de5b2163c9709b7d734c9ff6a29512410f6c64114f4220070.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The calendar that contains the event. Navigation property. Read-only.
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i4737ccf6c4de9b43ff9f15bf803182421cb705085720cbf9f8a6a7af0a794d4e.SingleValueExtendedPropertiesRequestBuilder) {
    return i4737ccf6c4de9b43ff9f15bf803182421cb705085720cbf9f8a6a7af0a794d4e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendarView.item.calendar.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i0c04a96c97affe987eaae7b174f477b5ba285b8f2dd6d566e40ab0f8064b7018.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i0c04a96c97affe987eaae7b174f477b5ba285b8f2dd6d566e40ab0f8064b7018.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
