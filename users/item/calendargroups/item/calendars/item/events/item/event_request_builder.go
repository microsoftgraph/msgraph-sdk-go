package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0f9d60ccac3e19508d94479f4590158d062306bb8e1db447f22799708e261e8f "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/accept"
    i26cba5917ffe762f3da7fb2f6465c97252639faf1ed6933904c03078326050b9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/calendar"
    i3d89795798d60a8674192206cb296fdc6c798b3a7e5fafad4ef7ad6fbc463f5c "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/forward"
    i4de7418aff4eb6faadca100368dec427222349438a55f7fe128d6dc9f8403113 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/extensions"
    i5691b3ac984532dca0d36718efc01d0da690a11db058e52f86dfca8e42e2e843 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/decline"
    i5d79f1cc3b5a62cb5d7f48597eaf612ff7bf8431c78bb65a7719c11356edcec6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/singlevalueextendedproperties"
    i5de5c408355445bfc2c1063b00fc6757f867b275c007f4ae093604ee53177a7a "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/cancel"
    i66cafcde733267d9d5f9dae37d9d80e46c78f6237b8690b1b8eba452272fd4fa "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/tentativelyaccept"
    i939f1e89656c7c04d15db716dcaed91b26add2c332547174bc549493d5a3cfe9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/snoozereminder"
    ib634327cbdcb39c4a84c74dea38795edde72c77311feaceb55ff7c45199ffb5e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances"
    id5b3d614d4ea44bb5c8853c3176e6a17bcc3750589d5a0ec78e9d14933175a0e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/attachments"
    id67ee023a9fcee00f4ca8cc1da5daa6089eaf483a96a4efbe32c64474f064a0c "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/dismissreminder"
    id7efed4f2e4818b45b9ced91c4bcece56e343696f115194afdb515269955d71f "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/multivalueextendedproperties"
    i81e2cf60be35c93438ebbb4baa8af217706df0cbc6f152f8baa1acdf2f6d21d4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/attachments/item"
    i9e26473fd65beefb7761db658a14b00167326427473ad98bc3de4f0521bb9be5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item"
    iaec9bf157c9c25abdf3e755897e9e055cf1a56654e1211f2ef0ceb2233b962c8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/multivalueextendedproperties/item"
    id6c5374cf8c61a96fbc34e5c44074670f36a69922f7f04e61e44cc2212db4c9c "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/extensions/item"
    ie7077f1b7f3633819b8377d4ff16e3b11b347d4adbaf777448a863505427f4b5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendargroups/item/calendars/item/events/item/singlevalueextendedproperties/item"
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
func (m *EventRequestBuilder) Accept()(*i0f9d60ccac3e19508d94479f4590158d062306bb8e1db447f22799708e261e8f.AcceptRequestBuilder) {
    return i0f9d60ccac3e19508d94479f4590158d062306bb8e1db447f22799708e261e8f.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*id5b3d614d4ea44bb5c8853c3176e6a17bcc3750589d5a0ec78e9d14933175a0e.AttachmentsRequestBuilder) {
    return id5b3d614d4ea44bb5c8853c3176e6a17bcc3750589d5a0ec78e9d14933175a0e.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) AttachmentsById(id string)(*i81e2cf60be35c93438ebbb4baa8af217706df0cbc6f152f8baa1acdf2f6d21d4.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i81e2cf60be35c93438ebbb4baa8af217706df0cbc6f152f8baa1acdf2f6d21d4.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i26cba5917ffe762f3da7fb2f6465c97252639faf1ed6933904c03078326050b9.CalendarRequestBuilder) {
    return i26cba5917ffe762f3da7fb2f6465c97252639faf1ed6933904c03078326050b9.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i5de5c408355445bfc2c1063b00fc6757f867b275c007f4ae093604ee53177a7a.CancelRequestBuilder) {
    return i5de5c408355445bfc2c1063b00fc6757f867b275c007f4ae093604ee53177a7a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}/events/{event_id}{?select,expand}";
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
func (m *EventRequestBuilder) Decline()(*i5691b3ac984532dca0d36718efc01d0da690a11db058e52f86dfca8e42e2e843.DeclineRequestBuilder) {
    return i5691b3ac984532dca0d36718efc01d0da690a11db058e52f86dfca8e42e2e843.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) DismissReminder()(*id67ee023a9fcee00f4ca8cc1da5daa6089eaf483a96a4efbe32c64474f064a0c.DismissReminderRequestBuilder) {
    return id67ee023a9fcee00f4ca8cc1da5daa6089eaf483a96a4efbe32c64474f064a0c.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i4de7418aff4eb6faadca100368dec427222349438a55f7fe128d6dc9f8403113.ExtensionsRequestBuilder) {
    return i4de7418aff4eb6faadca100368dec427222349438a55f7fe128d6dc9f8403113.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExtensionsById(id string)(*id6c5374cf8c61a96fbc34e5c44074670f36a69922f7f04e61e44cc2212db4c9c.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return id6c5374cf8c61a96fbc34e5c44074670f36a69922f7f04e61e44cc2212db4c9c.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i3d89795798d60a8674192206cb296fdc6c798b3a7e5fafad4ef7ad6fbc463f5c.ForwardRequestBuilder) {
    return i3d89795798d60a8674192206cb296fdc6c798b3a7e5fafad4ef7ad6fbc463f5c.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) Instances()(*ib634327cbdcb39c4a84c74dea38795edde72c77311feaceb55ff7c45199ffb5e.InstancesRequestBuilder) {
    return ib634327cbdcb39c4a84c74dea38795edde72c77311feaceb55ff7c45199ffb5e.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) InstancesById(id string)(*i9e26473fd65beefb7761db658a14b00167326427473ad98bc3de4f0521bb9be5.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i9e26473fd65beefb7761db658a14b00167326427473ad98bc3de4f0521bb9be5.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*id7efed4f2e4818b45b9ced91c4bcece56e343696f115194afdb515269955d71f.MultiValueExtendedPropertiesRequestBuilder) {
    return id7efed4f2e4818b45b9ced91c4bcece56e343696f115194afdb515269955d71f.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iaec9bf157c9c25abdf3e755897e9e055cf1a56654e1211f2ef0ceb2233b962c8.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return iaec9bf157c9c25abdf3e755897e9e055cf1a56654e1211f2ef0ceb2233b962c8.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i5d79f1cc3b5a62cb5d7f48597eaf612ff7bf8431c78bb65a7719c11356edcec6.SingleValueExtendedPropertiesRequestBuilder) {
    return i5d79f1cc3b5a62cb5d7f48597eaf612ff7bf8431c78bb65a7719c11356edcec6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ie7077f1b7f3633819b8377d4ff16e3b11b347d4adbaf777448a863505427f4b5.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ie7077f1b7f3633819b8377d4ff16e3b11b347d4adbaf777448a863505427f4b5.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i939f1e89656c7c04d15db716dcaed91b26add2c332547174bc549493d5a3cfe9.SnoozeReminderRequestBuilder) {
    return i939f1e89656c7c04d15db716dcaed91b26add2c332547174bc549493d5a3cfe9.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i66cafcde733267d9d5f9dae37d9d80e46c78f6237b8690b1b8eba452272fd4fa.TentativelyAcceptRequestBuilder) {
    return i66cafcde733267d9d5f9dae37d9d80e46c78f6237b8690b1b8eba452272fd4fa.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
