package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i3f79f65f324cc8f63e5311e3d4da88846b6af43dd68bbb77c63594059ec3a659 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/calendarview"
    i40a4d74c974865cc786a5f83edf3f4c30db936ba322a974f3d031fe907ee3bd0 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/singlevalueextendedproperties"
    i760d5c18cb49d9b46d954f9d71dc07087b06bee27cf7b46c867cb40885f05a1d "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/events"
    i87522b379cb368c7d67407c880eda95994d1966e04c354e1c7254c81d94c3df2 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/allowedcalendarsharingroleswithuser"
    ib3b3c31045d0f81371a2811f71d9b7541bac5974b70d21a29ba76dfe2c42005b "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/multivalueextendedproperties"
    iebeb9c025a78843e3645bc28b4e7f104796dbf8e64fa8a0e1c84e95c2ca352df "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/getschedule"
    if12cff0b93f69c1c889a8521b612ebd50353ed7ac09f8f4e295ffa2bab6a046b "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/calendarpermissions"
    i1437799c8e26980aef12efe3c53b8bfa00297a1469de145fadc289296e837d25 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/events/item"
    ia0abf3598f2b36c1131205830a4246fce4e4bfd333e7b0be84174ba49bb40162 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/calendarpermissions/item"
    ic3d178a18fab594dd1843b4b4d3ce5d1e59640f10c7d2bd570b3a38de6f3319e "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/multivalueextendedproperties/item"
    idbbbf63238b6489392f341b3a7b34642c4614b48e494061ec91f20bd70d131a3 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/calendarview/item"
    ideca945de53e399384bbc41ab854b4d1cdb613880c5e7a4e63c7ef95ee5a8490 "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item/calendar/singlevalueextendedproperties/item"
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
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i87522b379cb368c7d67407c880eda95994d1966e04c354e1c7254c81d94c3df2.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i87522b379cb368c7d67407c880eda95994d1966e04c354e1c7254c81d94c3df2.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*if12cff0b93f69c1c889a8521b612ebd50353ed7ac09f8f4e295ffa2bab6a046b.CalendarPermissionsRequestBuilder) {
    return if12cff0b93f69c1c889a8521b612ebd50353ed7ac09f8f4e295ffa2bab6a046b.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*ia0abf3598f2b36c1131205830a4246fce4e4bfd333e7b0be84174ba49bb40162.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return ia0abf3598f2b36c1131205830a4246fce4e4bfd333e7b0be84174ba49bb40162.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*i3f79f65f324cc8f63e5311e3d4da88846b6af43dd68bbb77c63594059ec3a659.CalendarViewRequestBuilder) {
    return i3f79f65f324cc8f63e5311e3d4da88846b6af43dd68bbb77c63594059ec3a659.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*idbbbf63238b6489392f341b3a7b34642c4614b48e494061ec91f20bd70d131a3.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return idbbbf63238b6489392f341b3a7b34642c4614b48e494061ec91f20bd70d131a3.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/calendarView/{event_id}/calendar{?select,expand}";
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
func (m *CalendarRequestBuilder) Events()(*i760d5c18cb49d9b46d954f9d71dc07087b06bee27cf7b46c867cb40885f05a1d.EventsRequestBuilder) {
    return i760d5c18cb49d9b46d954f9d71dc07087b06bee27cf7b46c867cb40885f05a1d.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) EventsById(id string)(*i1437799c8e26980aef12efe3c53b8bfa00297a1469de145fadc289296e837d25.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i1437799c8e26980aef12efe3c53b8bfa00297a1469de145fadc289296e837d25.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) GetSchedule()(*iebeb9c025a78843e3645bc28b4e7f104796dbf8e64fa8a0e1c84e95c2ca352df.GetScheduleRequestBuilder) {
    return iebeb9c025a78843e3645bc28b4e7f104796dbf8e64fa8a0e1c84e95c2ca352df.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*ib3b3c31045d0f81371a2811f71d9b7541bac5974b70d21a29ba76dfe2c42005b.MultiValueExtendedPropertiesRequestBuilder) {
    return ib3b3c31045d0f81371a2811f71d9b7541bac5974b70d21a29ba76dfe2c42005b.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ic3d178a18fab594dd1843b4b4d3ce5d1e59640f10c7d2bd570b3a38de6f3319e.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ic3d178a18fab594dd1843b4b4d3ce5d1e59640f10c7d2bd570b3a38de6f3319e.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i40a4d74c974865cc786a5f83edf3f4c30db936ba322a974f3d031fe907ee3bd0.SingleValueExtendedPropertiesRequestBuilder) {
    return i40a4d74c974865cc786a5f83edf3f4c30db936ba322a974f3d031fe907ee3bd0.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ideca945de53e399384bbc41ab854b4d1cdb613880c5e7a4e63c7ef95ee5a8490.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ideca945de53e399384bbc41ab854b4d1cdb613880c5e7a4e63c7ef95ee5a8490.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
