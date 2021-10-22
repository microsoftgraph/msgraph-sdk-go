package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0f0ce6d736b6f504b471613c5ac15739a55461fedfe57875a5750f9b53c0f226 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/decline"
    i1a18a6b6a39e454f6c8d96179b980f6da06555b91b37580c9ec8ba0305e6de5f "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/extensions"
    i1c6a239e026b8cb9867b574711b33427f6c5744c2f14261d944366070fc353e5 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/calendar"
    i1ea6a33ab7759770447ac7a4d978f10017eacb31c6f0871731eeb148649fb5bc "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/tentativelyaccept"
    i26977c51a41815fe0d7ff9cbafcc8af4a348722bd293191d771b43d03047b305 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/singlevalueextendedproperties"
    i566f9c16a6fe85b63ce8dbb3f2fea3693448cd4a91cd02b2e3a7e71a63267860 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances"
    i77eb73a898b6f8d95bb82e9c2efcf03f741c8f2f7e7176dde95d01eed01893c0 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/attachments"
    i78a836c35965e9975e5e331baffe1a63b8525eef66bb1b8645b6e1e39525bef3 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/accept"
    i829fbd366d6ff12b081486076f81163508113569bf85669ed18cd726a41affc8 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/forward"
    ib0d55d703ff4d551e433c43443d5dd4e5ee73676034a7c9cb395c174c3e3c434 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/cancel"
    ie6bafb4b17c2a9ac6f86bde54240348db1f716dae601c3aaa3ba5246a106e2c7 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/multivalueextendedproperties"
    ie7a0d6568144c2179c8661e11538ef71f2732b97fd1f89e98eaa928afb3f952e "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/dismissreminder"
    ifc94e234d0cd341ab21ed1133e622aa778dde1f8670fa77373fecf95bb710d09 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/snoozereminder"
    i435b79a3073f9e3b67954c2f479ceba4b84b060beffbf4eddfdcd11b9c0afdd6 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/multivalueextendedproperties/item"
    i5c4cc58062b30213bec3fe8c70fe2c7eba1480b941dd3c248011fe0fb3c4dcf9 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/extensions/item"
    i9d4b2dde6d1a46763d115eb03242b93910eb5891036290fb5e8a49322ead62da "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/instances/item"
    i9e8c599582a73f84f65086c644266f1ecd869813360b75664d307bffb48d75bc "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/attachments/item"
    if591e3f43c598d182e086c7492b69164db7e06e4a23e37dd47fab1c1dfa49207 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item/calendarview/item/singlevalueextendedproperties/item"
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
func (m *EventRequestBuilder) Accept()(*i78a836c35965e9975e5e331baffe1a63b8525eef66bb1b8645b6e1e39525bef3.AcceptRequestBuilder) {
    return i78a836c35965e9975e5e331baffe1a63b8525eef66bb1b8645b6e1e39525bef3.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i77eb73a898b6f8d95bb82e9c2efcf03f741c8f2f7e7176dde95d01eed01893c0.AttachmentsRequestBuilder) {
    return i77eb73a898b6f8d95bb82e9c2efcf03f741c8f2f7e7176dde95d01eed01893c0.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) AttachmentsById(id string)(*i9e8c599582a73f84f65086c644266f1ecd869813360b75664d307bffb48d75bc.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i9e8c599582a73f84f65086c644266f1ecd869813360b75664d307bffb48d75bc.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i1c6a239e026b8cb9867b574711b33427f6c5744c2f14261d944366070fc353e5.CalendarRequestBuilder) {
    return i1c6a239e026b8cb9867b574711b33427f6c5744c2f14261d944366070fc353e5.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*ib0d55d703ff4d551e433c43443d5dd4e5ee73676034a7c9cb395c174c3e3c434.CancelRequestBuilder) {
    return ib0d55d703ff4d551e433c43443d5dd4e5ee73676034a7c9cb395c174c3e3c434.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/calendars/{calendar_id}/calendarView/{event_id}{?startDateTime,endDateTime,select,expand}";
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
func (m *EventRequestBuilder) Decline()(*i0f0ce6d736b6f504b471613c5ac15739a55461fedfe57875a5750f9b53c0f226.DeclineRequestBuilder) {
    return i0f0ce6d736b6f504b471613c5ac15739a55461fedfe57875a5750f9b53c0f226.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) DismissReminder()(*ie7a0d6568144c2179c8661e11538ef71f2732b97fd1f89e98eaa928afb3f952e.DismissReminderRequestBuilder) {
    return ie7a0d6568144c2179c8661e11538ef71f2732b97fd1f89e98eaa928afb3f952e.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i1a18a6b6a39e454f6c8d96179b980f6da06555b91b37580c9ec8ba0305e6de5f.ExtensionsRequestBuilder) {
    return i1a18a6b6a39e454f6c8d96179b980f6da06555b91b37580c9ec8ba0305e6de5f.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExtensionsById(id string)(*i5c4cc58062b30213bec3fe8c70fe2c7eba1480b941dd3c248011fe0fb3c4dcf9.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i5c4cc58062b30213bec3fe8c70fe2c7eba1480b941dd3c248011fe0fb3c4dcf9.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i829fbd366d6ff12b081486076f81163508113569bf85669ed18cd726a41affc8.ForwardRequestBuilder) {
    return i829fbd366d6ff12b081486076f81163508113569bf85669ed18cd726a41affc8.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) Instances()(*i566f9c16a6fe85b63ce8dbb3f2fea3693448cd4a91cd02b2e3a7e71a63267860.InstancesRequestBuilder) {
    return i566f9c16a6fe85b63ce8dbb3f2fea3693448cd4a91cd02b2e3a7e71a63267860.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) InstancesById(id string)(*i9d4b2dde6d1a46763d115eb03242b93910eb5891036290fb5e8a49322ead62da.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i9d4b2dde6d1a46763d115eb03242b93910eb5891036290fb5e8a49322ead62da.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*ie6bafb4b17c2a9ac6f86bde54240348db1f716dae601c3aaa3ba5246a106e2c7.MultiValueExtendedPropertiesRequestBuilder) {
    return ie6bafb4b17c2a9ac6f86bde54240348db1f716dae601c3aaa3ba5246a106e2c7.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i435b79a3073f9e3b67954c2f479ceba4b84b060beffbf4eddfdcd11b9c0afdd6.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i435b79a3073f9e3b67954c2f479ceba4b84b060beffbf4eddfdcd11b9c0afdd6.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i26977c51a41815fe0d7ff9cbafcc8af4a348722bd293191d771b43d03047b305.SingleValueExtendedPropertiesRequestBuilder) {
    return i26977c51a41815fe0d7ff9cbafcc8af4a348722bd293191d771b43d03047b305.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*if591e3f43c598d182e086c7492b69164db7e06e4a23e37dd47fab1c1dfa49207.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return if591e3f43c598d182e086c7492b69164db7e06e4a23e37dd47fab1c1dfa49207.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*ifc94e234d0cd341ab21ed1133e622aa778dde1f8670fa77373fecf95bb710d09.SnoozeReminderRequestBuilder) {
    return ifc94e234d0cd341ab21ed1133e622aa778dde1f8670fa77373fecf95bb710d09.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i1ea6a33ab7759770447ac7a4d978f10017eacb31c6f0871731eeb148649fb5bc.TentativelyAcceptRequestBuilder) {
    return i1ea6a33ab7759770447ac7a4d978f10017eacb31c6f0871731eeb148649fb5bc.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
