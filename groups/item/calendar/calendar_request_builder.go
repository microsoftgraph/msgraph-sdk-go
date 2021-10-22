package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i325e8c5a186a2ac5ba63357864499a2daeca2d6b1fe397a8e242931a81e90fd6 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/singlevalueextendedproperties"
    i564e44237d1fd29264e29294fe61a277a2b68966ba3da977d4dc4188473efc6f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/allowedcalendarsharingroleswithuser"
    i59567e44c84457909a7acd2691855472309ae72d46905ad5b927e32ac1bc6775 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarpermissions"
    i8a7b3ddf32454c04e11c42fafddbe8a65a1b8c29699e948e42fb5fb489efff7a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/getschedule"
    i91f19a09d718f883e167c3c1fea6e1dc0bafbd34f8ca677f95c72d7745bcd847 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/multivalueextendedproperties"
    ia4056a24d617675f13dc550d9eb8d60439d535c22994554b767157caae834324 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview"
    iae525d555fa999780edca2ea54ea681baf7b616326144106b66e173f4930b985 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events"
    i066e4b225a90f837574cba0d06582cb5234a89a58269aad5ccca1e0d592385ed "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item"
    i32f1c0e0a1af3233116ddc476f2459674db4e63222c4a5762ca7049d50e06a83 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/singlevalueextendedproperties/item"
    i84f844d52d5e497fca0d718f40dd524efb38a3a493f27e5b266d0f0b0d06bd58 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarpermissions/item"
    ieb71583a78a26be4af7a9008915150796bd995e96567f7d812206ea5c47e34ee "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item"
    iff61068eadc42d2652152ea70da0c96f72f2b3e78201eeb2e58a48230426fa5a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/multivalueextendedproperties/item"
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
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i564e44237d1fd29264e29294fe61a277a2b68966ba3da977d4dc4188473efc6f.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i564e44237d1fd29264e29294fe61a277a2b68966ba3da977d4dc4188473efc6f.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*i59567e44c84457909a7acd2691855472309ae72d46905ad5b927e32ac1bc6775.CalendarPermissionsRequestBuilder) {
    return i59567e44c84457909a7acd2691855472309ae72d46905ad5b927e32ac1bc6775.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i84f844d52d5e497fca0d718f40dd524efb38a3a493f27e5b266d0f0b0d06bd58.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i84f844d52d5e497fca0d718f40dd524efb38a3a493f27e5b266d0f0b0d06bd58.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*ia4056a24d617675f13dc550d9eb8d60439d535c22994554b767157caae834324.CalendarViewRequestBuilder) {
    return ia4056a24d617675f13dc550d9eb8d60439d535c22994554b767157caae834324.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i066e4b225a90f837574cba0d06582cb5234a89a58269aad5ccca1e0d592385ed.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i066e4b225a90f837574cba0d06582cb5234a89a58269aad5ccca1e0d592385ed.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/groups/{group_id}/calendar{?select,expand}";
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
func (m *CalendarRequestBuilder) Events()(*iae525d555fa999780edca2ea54ea681baf7b616326144106b66e173f4930b985.EventsRequestBuilder) {
    return iae525d555fa999780edca2ea54ea681baf7b616326144106b66e173f4930b985.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) EventsById(id string)(*ieb71583a78a26be4af7a9008915150796bd995e96567f7d812206ea5c47e34ee.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return ieb71583a78a26be4af7a9008915150796bd995e96567f7d812206ea5c47e34ee.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) GetSchedule()(*i8a7b3ddf32454c04e11c42fafddbe8a65a1b8c29699e948e42fb5fb489efff7a.GetScheduleRequestBuilder) {
    return i8a7b3ddf32454c04e11c42fafddbe8a65a1b8c29699e948e42fb5fb489efff7a.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i91f19a09d718f883e167c3c1fea6e1dc0bafbd34f8ca677f95c72d7745bcd847.MultiValueExtendedPropertiesRequestBuilder) {
    return i91f19a09d718f883e167c3c1fea6e1dc0bafbd34f8ca677f95c72d7745bcd847.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iff61068eadc42d2652152ea70da0c96f72f2b3e78201eeb2e58a48230426fa5a.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return iff61068eadc42d2652152ea70da0c96f72f2b3e78201eeb2e58a48230426fa5a.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i325e8c5a186a2ac5ba63357864499a2daeca2d6b1fe397a8e242931a81e90fd6.SingleValueExtendedPropertiesRequestBuilder) {
    return i325e8c5a186a2ac5ba63357864499a2daeca2d6b1fe397a8e242931a81e90fd6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i32f1c0e0a1af3233116ddc476f2459674db4e63222c4a5762ca7049d50e06a83.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i32f1c0e0a1af3233116ddc476f2459674db4e63222c4a5762ca7049d50e06a83.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
