package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0acef849aee4ed2dae2d03e2761b4d2d11499bb4e3f7e240209ebe5cda1170d6 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/singlevalueextendedproperties"
    i43e4646345ea106ee726f101004135b0094a2df4214ee9f35c4c71c39fc3f689 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/multivalueextendedproperties"
    i4e6b455f1eeb8b9ccd9dbd76bf9f10f94f3623f746a732e58a36e76d7ca6fa06 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/calendarpermissions"
    i5e27aa8fc6aeb6944c5bbb8a3d4b81cb9f0f0cd3d18dce0e1418177a10949111 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/allowedcalendarsharingroleswithuser"
    i6a14c47b472ca4c040950e43d768346e1a2aa53eb080fe9eb1ad0af6e3b9f12c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/getschedule"
    i75125b08852f749d38bfbc9ccec9ad54ebf64c4791bafd2cadc1fbc2ad1d7b91 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/events"
    ie508f85483b663b20b1ceb98a59f616c70c0e776a328f86dd5a7f73628830b74 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/calendarview"
    i01338c46341d1b67f3679d661f4920e14ee9231fae6ccf8f6299304f522efae8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/multivalueextendedproperties/item"
    i036cfa7473e6def0fbbef47c4ca974f0c3f18fe7fbaf63ec8c70a77939d74d60 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/calendarpermissions/item"
    i0f36dec42b0613c8a0385fbbe11758edffc9d08c45be537bf26edd9a8c5dd040 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/events/item"
    i3d9c70b9bf8cbb06b5864c9f850a40465016663728792465f5320e9d74ed8d8d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/calendarview/item"
    idaefa7f993ae51db9765dd1fd885f6a493455efbf2e1654ae13bfd6374d1efe1 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendarview/item/calendar/singlevalueextendedproperties/item"
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
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i5e27aa8fc6aeb6944c5bbb8a3d4b81cb9f0f0cd3d18dce0e1418177a10949111.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i5e27aa8fc6aeb6944c5bbb8a3d4b81cb9f0f0cd3d18dce0e1418177a10949111.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*i4e6b455f1eeb8b9ccd9dbd76bf9f10f94f3623f746a732e58a36e76d7ca6fa06.CalendarPermissionsRequestBuilder) {
    return i4e6b455f1eeb8b9ccd9dbd76bf9f10f94f3623f746a732e58a36e76d7ca6fa06.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i036cfa7473e6def0fbbef47c4ca974f0c3f18fe7fbaf63ec8c70a77939d74d60.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i036cfa7473e6def0fbbef47c4ca974f0c3f18fe7fbaf63ec8c70a77939d74d60.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*ie508f85483b663b20b1ceb98a59f616c70c0e776a328f86dd5a7f73628830b74.CalendarViewRequestBuilder) {
    return ie508f85483b663b20b1ceb98a59f616c70c0e776a328f86dd5a7f73628830b74.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i3d9c70b9bf8cbb06b5864c9f850a40465016663728792465f5320e9d74ed8d8d.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i3d9c70b9bf8cbb06b5864c9f850a40465016663728792465f5320e9d74ed8d8d.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/groups/{group_id}/calendarView/{event_id}/calendar{?select,expand}";
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
func (m *CalendarRequestBuilder) Events()(*i75125b08852f749d38bfbc9ccec9ad54ebf64c4791bafd2cadc1fbc2ad1d7b91.EventsRequestBuilder) {
    return i75125b08852f749d38bfbc9ccec9ad54ebf64c4791bafd2cadc1fbc2ad1d7b91.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) EventsById(id string)(*i0f36dec42b0613c8a0385fbbe11758edffc9d08c45be537bf26edd9a8c5dd040.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i0f36dec42b0613c8a0385fbbe11758edffc9d08c45be537bf26edd9a8c5dd040.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) GetSchedule()(*i6a14c47b472ca4c040950e43d768346e1a2aa53eb080fe9eb1ad0af6e3b9f12c.GetScheduleRequestBuilder) {
    return i6a14c47b472ca4c040950e43d768346e1a2aa53eb080fe9eb1ad0af6e3b9f12c.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i43e4646345ea106ee726f101004135b0094a2df4214ee9f35c4c71c39fc3f689.MultiValueExtendedPropertiesRequestBuilder) {
    return i43e4646345ea106ee726f101004135b0094a2df4214ee9f35c4c71c39fc3f689.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i01338c46341d1b67f3679d661f4920e14ee9231fae6ccf8f6299304f522efae8.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i01338c46341d1b67f3679d661f4920e14ee9231fae6ccf8f6299304f522efae8.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i0acef849aee4ed2dae2d03e2761b4d2d11499bb4e3f7e240209ebe5cda1170d6.SingleValueExtendedPropertiesRequestBuilder) {
    return i0acef849aee4ed2dae2d03e2761b4d2d11499bb4e3f7e240209ebe5cda1170d6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*idaefa7f993ae51db9765dd1fd885f6a493455efbf2e1654ae13bfd6374d1efe1.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return idaefa7f993ae51db9765dd1fd885f6a493455efbf2e1654ae13bfd6374d1efe1.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
