package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i321470654b3ef34bacc414654086f423258f1193a26b214185a5cc558e2bb5e2 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i6af69cbc9d5ee4e56d0394b5dbb9ba47c157c0ee557e79646739cd31671483d3 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/multivalueextendedproperties"
    i8106511447607356fcb7172494684f7bb195617a2f3cfb8090371e4dca9494cf "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/singlevalueextendedproperties"
    ia93089358fcf18f5dfb97b53182f1524884b6a134ef09311582420c51b266f63 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/allowedcalendarsharingroleswithuser"
    ib746c7bd792e0607afacbee87215b032700dfb09c38134acd3e5367b09bd1eef "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/getschedule"
    ic3f5d8e3fd6959673a012943934769f057abadba324812644982eeef5f407def "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events"
    iefe8e12aae13619f806b3f0379ff65f77fd7287cd264dc16f120a83b59db3999 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarpermissions"
    i209c23501c843b5dd2a14fe5d4f0a449bc55aac8f20ad31a9d86e7cb75577e28 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/multivalueextendedproperties/item"
    i79296554e9c2c5b59da0a63741ace146ef92bbb0bc577e1866c1d097111c377c "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/singlevalueextendedproperties/item"
    i96d50ecda61b15c98effcc511b1eff207b365d1e3b05ef2f704013788d04ccae "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/events/item"
    i9e50aa45f8b8ff846d0226fbc640efa378ef39164e74364f6e3b11c0735dfc91 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarview/item"
    if33c45189b52ce3c1a36912a56beedbaacd4a8273d6116ef5f83bd582e5a8506 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar/calendarpermissions/item"
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
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*ia93089358fcf18f5dfb97b53182f1524884b6a134ef09311582420c51b266f63.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return ia93089358fcf18f5dfb97b53182f1524884b6a134ef09311582420c51b266f63.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*iefe8e12aae13619f806b3f0379ff65f77fd7287cd264dc16f120a83b59db3999.CalendarPermissionsRequestBuilder) {
    return iefe8e12aae13619f806b3f0379ff65f77fd7287cd264dc16f120a83b59db3999.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*if33c45189b52ce3c1a36912a56beedbaacd4a8273d6116ef5f83bd582e5a8506.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return if33c45189b52ce3c1a36912a56beedbaacd4a8273d6116ef5f83bd582e5a8506.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*i321470654b3ef34bacc414654086f423258f1193a26b214185a5cc558e2bb5e2.CalendarViewRequestBuilder) {
    return i321470654b3ef34bacc414654086f423258f1193a26b214185a5cc558e2bb5e2.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i9e50aa45f8b8ff846d0226fbc640efa378ef39164e74364f6e3b11c0735dfc91.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i9e50aa45f8b8ff846d0226fbc640efa378ef39164e74364f6e3b11c0735dfc91.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/calendar{?select,expand}";
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
func (m *CalendarRequestBuilder) Events()(*ic3f5d8e3fd6959673a012943934769f057abadba324812644982eeef5f407def.EventsRequestBuilder) {
    return ic3f5d8e3fd6959673a012943934769f057abadba324812644982eeef5f407def.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) EventsById(id string)(*i96d50ecda61b15c98effcc511b1eff207b365d1e3b05ef2f704013788d04ccae.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i96d50ecda61b15c98effcc511b1eff207b365d1e3b05ef2f704013788d04ccae.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) GetSchedule()(*ib746c7bd792e0607afacbee87215b032700dfb09c38134acd3e5367b09bd1eef.GetScheduleRequestBuilder) {
    return ib746c7bd792e0607afacbee87215b032700dfb09c38134acd3e5367b09bd1eef.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i6af69cbc9d5ee4e56d0394b5dbb9ba47c157c0ee557e79646739cd31671483d3.MultiValueExtendedPropertiesRequestBuilder) {
    return i6af69cbc9d5ee4e56d0394b5dbb9ba47c157c0ee557e79646739cd31671483d3.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i209c23501c843b5dd2a14fe5d4f0a449bc55aac8f20ad31a9d86e7cb75577e28.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i209c23501c843b5dd2a14fe5d4f0a449bc55aac8f20ad31a9d86e7cb75577e28.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i8106511447607356fcb7172494684f7bb195617a2f3cfb8090371e4dca9494cf.SingleValueExtendedPropertiesRequestBuilder) {
    return i8106511447607356fcb7172494684f7bb195617a2f3cfb8090371e4dca9494cf.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i79296554e9c2c5b59da0a63741ace146ef92bbb0bc577e1866c1d097111c377c.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i79296554e9c2c5b59da0a63741ace146ef92bbb0bc577e1866c1d097111c377c.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
