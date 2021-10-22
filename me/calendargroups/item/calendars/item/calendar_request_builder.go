package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0853bc5e78a4ac665d8aa887531e364debb1ba91c3a1bc736be9331cc8db0b52 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/singlevalueextendedproperties"
    i25a4d487ad09f2eb45ba28744d3fe919b7b5dbbf1c645ddd9776a9b357d0ff96 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview"
    i3abccbd6b1092e6246d7bc4dd7dcac1b6938ee6bf024fb4eccaaeecfb29ff1b6 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events"
    i3e09978bd77ced5481016e4684368f72edfa92ef709a19504e9258cf8865bfc6 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/allowedcalendarsharingroleswithuser"
    i847a76068ecb832435c51c746fba99e17f01da45f86bc0ea0f29b5a1d42ae5b3 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/multivalueextendedproperties"
    ia0a09be4f5c1f5b3b99aa3064830cdc6c730ceea9bdb4dbc6934a7d5f0ffec6f "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/getschedule"
    iaac11e3c8bf2dd29e59be19f34577e748fb426a2bb9b94aae4be18a9b2f6474a "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarpermissions"
    i67198c8bd92c3d2fcf155bff41150fc55dc916147fb9b1eb9ca2ccf6b9d1ed6e "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/events/item"
    i857bdaf62ca4819c99b59159f0888909252fbffeba2fa7050dfc75ce2c6b598d "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/singlevalueextendedproperties/item"
    i8783c6ceaf93b1654906d93bb1e4267cb1060f0bcdbc2d302e4f64bfac674b76 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarview/item"
    i8b359363f8c83629e3e25072f613c3d3d807cc671fe5472a27f65565ce218c26 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/calendarpermissions/item"
    idefd7aa8b66a92e862195e44b4d12afbb8cb1f08136af3a806efad5c689470d7 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item/calendars/item/multivalueextendedproperties/item"
)

type CalendarRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type CalendarRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i3e09978bd77ced5481016e4684368f72edfa92ef709a19504e9258cf8865bfc6.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i3e09978bd77ced5481016e4684368f72edfa92ef709a19504e9258cf8865bfc6.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*iaac11e3c8bf2dd29e59be19f34577e748fb426a2bb9b94aae4be18a9b2f6474a.CalendarPermissionsRequestBuilder) {
    return iaac11e3c8bf2dd29e59be19f34577e748fb426a2bb9b94aae4be18a9b2f6474a.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i8b359363f8c83629e3e25072f613c3d3d807cc671fe5472a27f65565ce218c26.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i8b359363f8c83629e3e25072f613c3d3d807cc671fe5472a27f65565ce218c26.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*i25a4d487ad09f2eb45ba28744d3fe919b7b5dbbf1c645ddd9776a9b357d0ff96.CalendarViewRequestBuilder) {
    return i25a4d487ad09f2eb45ba28744d3fe919b7b5dbbf1c645ddd9776a9b357d0ff96.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i8783c6ceaf93b1654906d93bb1e4267cb1060f0bcdbc2d302e4f64bfac674b76.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i8783c6ceaf93b1654906d93bb1e4267cb1060f0bcdbc2d302e4f64bfac674b76.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewCalendarRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarRequestBuilderInternal(urlParams, requestAdapter)
}
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
func (m *CalendarRequestBuilder) Events()(*i3abccbd6b1092e6246d7bc4dd7dcac1b6938ee6bf024fb4eccaaeecfb29ff1b6.EventsRequestBuilder) {
    return i3abccbd6b1092e6246d7bc4dd7dcac1b6938ee6bf024fb4eccaaeecfb29ff1b6.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) EventsById(id string)(*i67198c8bd92c3d2fcf155bff41150fc55dc916147fb9b1eb9ca2ccf6b9d1ed6e.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i67198c8bd92c3d2fcf155bff41150fc55dc916147fb9b1eb9ca2ccf6b9d1ed6e.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
func (m *CalendarRequestBuilder) GetSchedule()(*ia0a09be4f5c1f5b3b99aa3064830cdc6c730ceea9bdb4dbc6934a7d5f0ffec6f.GetScheduleRequestBuilder) {
    return ia0a09be4f5c1f5b3b99aa3064830cdc6c730ceea9bdb4dbc6934a7d5f0ffec6f.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i847a76068ecb832435c51c746fba99e17f01da45f86bc0ea0f29b5a1d42ae5b3.MultiValueExtendedPropertiesRequestBuilder) {
    return i847a76068ecb832435c51c746fba99e17f01da45f86bc0ea0f29b5a1d42ae5b3.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*idefd7aa8b66a92e862195e44b4d12afbb8cb1f08136af3a806efad5c689470d7.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return idefd7aa8b66a92e862195e44b4d12afbb8cb1f08136af3a806efad5c689470d7.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i0853bc5e78a4ac665d8aa887531e364debb1ba91c3a1bc736be9331cc8db0b52.SingleValueExtendedPropertiesRequestBuilder) {
    return i0853bc5e78a4ac665d8aa887531e364debb1ba91c3a1bc736be9331cc8db0b52.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i857bdaf62ca4819c99b59159f0888909252fbffeba2fa7050dfc75ce2c6b598d.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i857bdaf62ca4819c99b59159f0888909252fbffeba2fa7050dfc75ce2c6b598d.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
