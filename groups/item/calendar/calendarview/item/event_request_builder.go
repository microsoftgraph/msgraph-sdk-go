package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1233ad0ba006ab96918402fc9449bdbea95cc2732eac86365bde41c84b357d21 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/singlevalueextendedproperties"
    i133098166421eafcbb7d41e4a68f3d29757bcf6688e847c4a7808bf543993aef "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/tentativelyaccept"
    i2ef9251ae81121975f3c9a073ec23b2469d36ed1973d4d03af6e70a2f6fc3c4d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/multivalueextendedproperties"
    i32caceec19579e592d7312d6a514ad4bfa9807198b6c2414606a3703116e102f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/snoozereminder"
    i42df08a14e97e22bb26eddf9cbaa35aa7bf02405a8bbc84fe4fe5b4203cc2ff8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances"
    i53ba2fe10fcac93a89e1ab826cf6d7b8e1e1069ce7a61b20bc3ce310e26333ab "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/accept"
    i6a59555a6a577b7a15b03f87f857fa23cfe6af2eace90222104dbcdc30a0d24d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/attachments"
    i6bf2dbc96674e29a2505ab70f10b8f1735eb7fe1b7640bf2d65e16312c0a071d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/calendar"
    i8c3a0045c4c0ac4da728d4aa37dbba4e6df99e52c194414aedb749b76e669fca "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/forward"
    i921d4dffa29707e7b15bb069fe546aa8d25d138f815d39cdd30ed48b40c27465 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/extensions"
    iac711011c7c99048797a7e80afa2fb7cb707667348b6c9c773d351f73e0a403f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/cancel"
    ie6ebec3d6348227c40f20857d0105bccb0382669c86a08c61932f7f04269ae7a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/dismissreminder"
    ie8cb6c49b7e9e2ac0ad0271442195855cb47449992c78270fe676de14c335081 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/decline"
    i00430d12f816c0c6997639410f009cc5f413c588c399506a408cafb1d4a794ec "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/multivalueextendedproperties/item"
    i1ec50c12270bd7bd6ab657694752ebec839c5bc0a607be433717875d65bff30e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/extensions/item"
    i9be5b8dcebdfca6e7ace3db98d602d80c97a8f974e4125c75bacc246c066e21e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/singlevalueextendedproperties/item"
    i9fe44f8859760d86bd88baff7a59fc2a46f6264caa52b2adf30c8d7f4ae24672 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/attachments/item"
    id60fc1ecd13479f120a7edbab2577fd0b8fff9ebc834f6c7d280345dd5f73d2e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item"
)

type EventRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EventRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    EndDateTime *string;
    Expand []string;
    Select_escpaped []string;
    StartDateTime *string;
}
func (m *EventRequestBuilder) Accept()(*i53ba2fe10fcac93a89e1ab826cf6d7b8e1e1069ce7a61b20bc3ce310e26333ab.AcceptRequestBuilder) {
    return i53ba2fe10fcac93a89e1ab826cf6d7b8e1e1069ce7a61b20bc3ce310e26333ab.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i6a59555a6a577b7a15b03f87f857fa23cfe6af2eace90222104dbcdc30a0d24d.AttachmentsRequestBuilder) {
    return i6a59555a6a577b7a15b03f87f857fa23cfe6af2eace90222104dbcdc30a0d24d.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) AttachmentsById(id string)(*i9fe44f8859760d86bd88baff7a59fc2a46f6264caa52b2adf30c8d7f4ae24672.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i9fe44f8859760d86bd88baff7a59fc2a46f6264caa52b2adf30c8d7f4ae24672.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i6bf2dbc96674e29a2505ab70f10b8f1735eb7fe1b7640bf2d65e16312c0a071d.CalendarRequestBuilder) {
    return i6bf2dbc96674e29a2505ab70f10b8f1735eb7fe1b7640bf2d65e16312c0a071d.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*iac711011c7c99048797a7e80afa2fb7cb707667348b6c9c773d351f73e0a403f.CancelRequestBuilder) {
    return iac711011c7c99048797a7e80afa2fb7cb707667348b6c9c773d351f73e0a403f.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/groups/{group_id}/calendar/calendarView/{event_id}{?startDateTime,endDateTime,select,expand}";
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
func NewEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventRequestBuilderInternal(urlParams, requestAdapter)
}
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
func (m *EventRequestBuilder) Decline()(*ie8cb6c49b7e9e2ac0ad0271442195855cb47449992c78270fe676de14c335081.DeclineRequestBuilder) {
    return ie8cb6c49b7e9e2ac0ad0271442195855cb47449992c78270fe676de14c335081.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
func (m *EventRequestBuilder) DismissReminder()(*ie6ebec3d6348227c40f20857d0105bccb0382669c86a08c61932f7f04269ae7a.DismissReminderRequestBuilder) {
    return ie6ebec3d6348227c40f20857d0105bccb0382669c86a08c61932f7f04269ae7a.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i921d4dffa29707e7b15bb069fe546aa8d25d138f815d39cdd30ed48b40c27465.ExtensionsRequestBuilder) {
    return i921d4dffa29707e7b15bb069fe546aa8d25d138f815d39cdd30ed48b40c27465.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExtensionsById(id string)(*i1ec50c12270bd7bd6ab657694752ebec839c5bc0a607be433717875d65bff30e.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i1ec50c12270bd7bd6ab657694752ebec839c5bc0a607be433717875d65bff30e.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i8c3a0045c4c0ac4da728d4aa37dbba4e6df99e52c194414aedb749b76e669fca.ForwardRequestBuilder) {
    return i8c3a0045c4c0ac4da728d4aa37dbba4e6df99e52c194414aedb749b76e669fca.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
func (m *EventRequestBuilder) Instances()(*i42df08a14e97e22bb26eddf9cbaa35aa7bf02405a8bbc84fe4fe5b4203cc2ff8.InstancesRequestBuilder) {
    return i42df08a14e97e22bb26eddf9cbaa35aa7bf02405a8bbc84fe4fe5b4203cc2ff8.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) InstancesById(id string)(*id60fc1ecd13479f120a7edbab2577fd0b8fff9ebc834f6c7d280345dd5f73d2e.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return id60fc1ecd13479f120a7edbab2577fd0b8fff9ebc834f6c7d280345dd5f73d2e.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i2ef9251ae81121975f3c9a073ec23b2469d36ed1973d4d03af6e70a2f6fc3c4d.MultiValueExtendedPropertiesRequestBuilder) {
    return i2ef9251ae81121975f3c9a073ec23b2469d36ed1973d4d03af6e70a2f6fc3c4d.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i00430d12f816c0c6997639410f009cc5f413c588c399506a408cafb1d4a794ec.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i00430d12f816c0c6997639410f009cc5f413c588c399506a408cafb1d4a794ec.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i1233ad0ba006ab96918402fc9449bdbea95cc2732eac86365bde41c84b357d21.SingleValueExtendedPropertiesRequestBuilder) {
    return i1233ad0ba006ab96918402fc9449bdbea95cc2732eac86365bde41c84b357d21.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i9be5b8dcebdfca6e7ace3db98d602d80c97a8f974e4125c75bacc246c066e21e.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i9be5b8dcebdfca6e7ace3db98d602d80c97a8f974e4125c75bacc246c066e21e.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i32caceec19579e592d7312d6a514ad4bfa9807198b6c2414606a3703116e102f.SnoozeReminderRequestBuilder) {
    return i32caceec19579e592d7312d6a514ad4bfa9807198b6c2414606a3703116e102f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i133098166421eafcbb7d41e4a68f3d29757bcf6688e847c4a7808bf543993aef.TentativelyAcceptRequestBuilder) {
    return i133098166421eafcbb7d41e4a68f3d29757bcf6688e847c4a7808bf543993aef.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
