package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i57af6d665855a0506938ef0e172bba0a14925f6b6ca627e2a7f4b6b2b38fd3bc "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/forward"
    i5b955e22c663573118720c3bc93ddf49ea1d54fc25dbd270017f41a676a78121 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/calendar"
    i6886d0c1332bc9d55516ea947d572efa5b319910b95fbb0820df71ec1f978f5a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/instances"
    i860a97d3a0a455053fe8451f67950c18e9a2e873fe28c966f761d1441a1b3f8f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/cancel"
    i8c7ff56f4d95c55b53a19a6b026409669658ef2e87e0166571395d96a4f5dc9b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/accept"
    i970025de78959a55575efb35fa3a413a32ebb5ab4dd5a708532d02f291175272 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/extensions"
    ia5e7630df574c5850c03f09e0d5676dde3236e1bca7c9041678b2fa25ee07e57 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/decline"
    ia6e2cb4fce4bf7b4978f9f272441f1b3a147cf7e6cb2a0de353a804070b27112 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/tentativelyaccept"
    ib354b8e7fb23c71741722b1fefce34b3b747a56eef37f83a062018a757f1aa72 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/attachments"
    iba41db6ae0a52d498c587df3e5bfcd8faf2437bcf7d4e46fd758bb57a9b0793e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/singlevalueextendedproperties"
    ic243b07f463bb3f30c6cb6f2fd2d449e4e4cec617ebca670effd9fc299701f4c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/multivalueextendedproperties"
    id98aa3c369bca6039a3982c6e8eaad54a60e63bab85c3ca6d5ddcb7fef06e0ab "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/dismissreminder"
    if141e32134f129fe69c623ef5329c6c06edda15c3911a704e36f0879d7277b76 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/snoozereminder"
    i4d40e556c4fcba97401013dc12446fe5534169b8814c68f0f34bc69689136da0 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/singlevalueextendedproperties/item"
    i931f24b3eb23f7e9c765dc76b0676613a564dedcec2ade920230227f4d4da606 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/multivalueextendedproperties/item"
    ibfd185eccf652dada583d3ca36c3d4bfc599e1bd77b6d8088608b2683fa14880 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/extensions/item"
    ie19dffbaf9991ce0123621be61d8c60393863cee5c24c78b5a5311caf84c056f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/attachments/item"
    ie90f5d53aab3ae5d6e53998a425e8dac6aec14b0af9e64448368e99bd0616cb5 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item/instances/item"
)

type EventRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EventRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *EventRequestBuilder) Accept()(*i8c7ff56f4d95c55b53a19a6b026409669658ef2e87e0166571395d96a4f5dc9b.AcceptRequestBuilder) {
    return i8c7ff56f4d95c55b53a19a6b026409669658ef2e87e0166571395d96a4f5dc9b.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*ib354b8e7fb23c71741722b1fefce34b3b747a56eef37f83a062018a757f1aa72.AttachmentsRequestBuilder) {
    return ib354b8e7fb23c71741722b1fefce34b3b747a56eef37f83a062018a757f1aa72.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) AttachmentsById(id string)(*ie19dffbaf9991ce0123621be61d8c60393863cee5c24c78b5a5311caf84c056f.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ie19dffbaf9991ce0123621be61d8c60393863cee5c24c78b5a5311caf84c056f.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i5b955e22c663573118720c3bc93ddf49ea1d54fc25dbd270017f41a676a78121.CalendarRequestBuilder) {
    return i5b955e22c663573118720c3bc93ddf49ea1d54fc25dbd270017f41a676a78121.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i860a97d3a0a455053fe8451f67950c18e9a2e873fe28c966f761d1441a1b3f8f.CancelRequestBuilder) {
    return i860a97d3a0a455053fe8451f67950c18e9a2e873fe28c966f761d1441a1b3f8f.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/groups/{group_id}/calendar/events/{event_id}{?select,expand}";
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
func (m *EventRequestBuilder) Decline()(*ia5e7630df574c5850c03f09e0d5676dde3236e1bca7c9041678b2fa25ee07e57.DeclineRequestBuilder) {
    return ia5e7630df574c5850c03f09e0d5676dde3236e1bca7c9041678b2fa25ee07e57.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) DismissReminder()(*id98aa3c369bca6039a3982c6e8eaad54a60e63bab85c3ca6d5ddcb7fef06e0ab.DismissReminderRequestBuilder) {
    return id98aa3c369bca6039a3982c6e8eaad54a60e63bab85c3ca6d5ddcb7fef06e0ab.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i970025de78959a55575efb35fa3a413a32ebb5ab4dd5a708532d02f291175272.ExtensionsRequestBuilder) {
    return i970025de78959a55575efb35fa3a413a32ebb5ab4dd5a708532d02f291175272.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExtensionsById(id string)(*ibfd185eccf652dada583d3ca36c3d4bfc599e1bd77b6d8088608b2683fa14880.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ibfd185eccf652dada583d3ca36c3d4bfc599e1bd77b6d8088608b2683fa14880.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i57af6d665855a0506938ef0e172bba0a14925f6b6ca627e2a7f4b6b2b38fd3bc.ForwardRequestBuilder) {
    return i57af6d665855a0506938ef0e172bba0a14925f6b6ca627e2a7f4b6b2b38fd3bc.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) Instances()(*i6886d0c1332bc9d55516ea947d572efa5b319910b95fbb0820df71ec1f978f5a.InstancesRequestBuilder) {
    return i6886d0c1332bc9d55516ea947d572efa5b319910b95fbb0820df71ec1f978f5a.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) InstancesById(id string)(*ie90f5d53aab3ae5d6e53998a425e8dac6aec14b0af9e64448368e99bd0616cb5.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ie90f5d53aab3ae5d6e53998a425e8dac6aec14b0af9e64448368e99bd0616cb5.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*ic243b07f463bb3f30c6cb6f2fd2d449e4e4cec617ebca670effd9fc299701f4c.MultiValueExtendedPropertiesRequestBuilder) {
    return ic243b07f463bb3f30c6cb6f2fd2d449e4e4cec617ebca670effd9fc299701f4c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i931f24b3eb23f7e9c765dc76b0676613a564dedcec2ade920230227f4d4da606.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i931f24b3eb23f7e9c765dc76b0676613a564dedcec2ade920230227f4d4da606.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*iba41db6ae0a52d498c587df3e5bfcd8faf2437bcf7d4e46fd758bb57a9b0793e.SingleValueExtendedPropertiesRequestBuilder) {
    return iba41db6ae0a52d498c587df3e5bfcd8faf2437bcf7d4e46fd758bb57a9b0793e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i4d40e556c4fcba97401013dc12446fe5534169b8814c68f0f34bc69689136da0.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i4d40e556c4fcba97401013dc12446fe5534169b8814c68f0f34bc69689136da0.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*if141e32134f129fe69c623ef5329c6c06edda15c3911a704e36f0879d7277b76.SnoozeReminderRequestBuilder) {
    return if141e32134f129fe69c623ef5329c6c06edda15c3911a704e36f0879d7277b76.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*ia6e2cb4fce4bf7b4978f9f272441f1b3a147cf7e6cb2a0de353a804070b27112.TentativelyAcceptRequestBuilder) {
    return ia6e2cb4fce4bf7b4978f9f272441f1b3a147cf7e6cb2a0de353a804070b27112.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
