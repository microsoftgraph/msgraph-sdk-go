package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1398a2b6b7b6fa6075615f0b696c5d5d393449822e80a67ff6e1f1b006570589 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/cancel"
    i23879bc3fa3602cbaf93cee98173f588cf737baf3fa364fa5d532b4a7ccffb6f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/instances"
    i4b518553c867ab557e4a9014d64bff52b4a5614abc228a534b622097747f58c4 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/extensions"
    i855851d197827c1ee6475b539166033ff1c226b94affc6b7876227a822cefe56 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/singlevalueextendedproperties"
    i91e19aea3690b9482c589e1b4fffc198beb93c8882e7ad5dc30a2c61b975241d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/forward"
    i967611b2c88418ecd27c9ba55d8e1a0fb980af96f161a7941b0122d0af98e20b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/multivalueextendedproperties"
    i973bdbcda3dd1f6f25a1482be6ac8c9929a425f352242a5f72ad56ed2695a133 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/dismissreminder"
    i9d4cd76f12610ad565059847d3a79b13115ec8cc1967d272b7f9e33b62908d70 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/decline"
    ia9eabac6ecd410545dc4495208341bd846651d6673db813cabee1323b7ccebb1 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/attachments"
    iaf6ca7f88694008763f7c64dc5f9e9a4696d364070e649a47f06c9cbb9720f79 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/accept"
    ic09438710245b2671bf04abe5db2656009ca7799d1d80105aa23a063cf2f4e40 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/tentativelyaccept"
    idaa6a05522f747ce1a1a1c6a68f0551d5b3d05627b4ce5ddef9c85b80d4684b2 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar"
    ie274417b453065e3b9c1c84afef158b9149fe29aa54c90fc8fe2a8361a6a6880 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/snoozereminder"
    i3924f06647673effb94eed5dd5d2d7fe2b5b2f58ac82d0f02eab48c1219f1c07 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/multivalueextendedproperties/item"
    i5042ede149a551ff0f8fb6fd788882bbcfd795931aefebb153289527246e5e4d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/instances/item"
    i68e56f812a05f1c50a7fc8625d20c43951316b2122068afdbe3b9ec586bd6fe2 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/extensions/item"
    i9fcaa4fe7aa3132d8c368ba6abc3e2992d0ff4b8a67614921b91e92aab8cab92 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/singlevalueextendedproperties/item"
    ia3fe1bf9b023096d20271ede0e7f809b9ea0d5dc59475200a395629f1ecb917c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/attachments/item"
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
func (m *EventRequestBuilder) Accept()(*iaf6ca7f88694008763f7c64dc5f9e9a4696d364070e649a47f06c9cbb9720f79.AcceptRequestBuilder) {
    return iaf6ca7f88694008763f7c64dc5f9e9a4696d364070e649a47f06c9cbb9720f79.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*ia9eabac6ecd410545dc4495208341bd846651d6673db813cabee1323b7ccebb1.AttachmentsRequestBuilder) {
    return ia9eabac6ecd410545dc4495208341bd846651d6673db813cabee1323b7ccebb1.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) AttachmentsById(id string)(*ia3fe1bf9b023096d20271ede0e7f809b9ea0d5dc59475200a395629f1ecb917c.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ia3fe1bf9b023096d20271ede0e7f809b9ea0d5dc59475200a395629f1ecb917c.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*idaa6a05522f747ce1a1a1c6a68f0551d5b3d05627b4ce5ddef9c85b80d4684b2.CalendarRequestBuilder) {
    return idaa6a05522f747ce1a1a1c6a68f0551d5b3d05627b4ce5ddef9c85b80d4684b2.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i1398a2b6b7b6fa6075615f0b696c5d5d393449822e80a67ff6e1f1b006570589.CancelRequestBuilder) {
    return i1398a2b6b7b6fa6075615f0b696c5d5d393449822e80a67ff6e1f1b006570589.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/groups/{group_id}/events/{event_id}{?select,expand}";
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
func (m *EventRequestBuilder) Decline()(*i9d4cd76f12610ad565059847d3a79b13115ec8cc1967d272b7f9e33b62908d70.DeclineRequestBuilder) {
    return i9d4cd76f12610ad565059847d3a79b13115ec8cc1967d272b7f9e33b62908d70.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) DismissReminder()(*i973bdbcda3dd1f6f25a1482be6ac8c9929a425f352242a5f72ad56ed2695a133.DismissReminderRequestBuilder) {
    return i973bdbcda3dd1f6f25a1482be6ac8c9929a425f352242a5f72ad56ed2695a133.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i4b518553c867ab557e4a9014d64bff52b4a5614abc228a534b622097747f58c4.ExtensionsRequestBuilder) {
    return i4b518553c867ab557e4a9014d64bff52b4a5614abc228a534b622097747f58c4.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExtensionsById(id string)(*i68e56f812a05f1c50a7fc8625d20c43951316b2122068afdbe3b9ec586bd6fe2.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i68e56f812a05f1c50a7fc8625d20c43951316b2122068afdbe3b9ec586bd6fe2.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i91e19aea3690b9482c589e1b4fffc198beb93c8882e7ad5dc30a2c61b975241d.ForwardRequestBuilder) {
    return i91e19aea3690b9482c589e1b4fffc198beb93c8882e7ad5dc30a2c61b975241d.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) Instances()(*i23879bc3fa3602cbaf93cee98173f588cf737baf3fa364fa5d532b4a7ccffb6f.InstancesRequestBuilder) {
    return i23879bc3fa3602cbaf93cee98173f588cf737baf3fa364fa5d532b4a7ccffb6f.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) InstancesById(id string)(*i5042ede149a551ff0f8fb6fd788882bbcfd795931aefebb153289527246e5e4d.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i5042ede149a551ff0f8fb6fd788882bbcfd795931aefebb153289527246e5e4d.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i967611b2c88418ecd27c9ba55d8e1a0fb980af96f161a7941b0122d0af98e20b.MultiValueExtendedPropertiesRequestBuilder) {
    return i967611b2c88418ecd27c9ba55d8e1a0fb980af96f161a7941b0122d0af98e20b.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i3924f06647673effb94eed5dd5d2d7fe2b5b2f58ac82d0f02eab48c1219f1c07.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i3924f06647673effb94eed5dd5d2d7fe2b5b2f58ac82d0f02eab48c1219f1c07.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i855851d197827c1ee6475b539166033ff1c226b94affc6b7876227a822cefe56.SingleValueExtendedPropertiesRequestBuilder) {
    return i855851d197827c1ee6475b539166033ff1c226b94affc6b7876227a822cefe56.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i9fcaa4fe7aa3132d8c368ba6abc3e2992d0ff4b8a67614921b91e92aab8cab92.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i9fcaa4fe7aa3132d8c368ba6abc3e2992d0ff4b8a67614921b91e92aab8cab92.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*ie274417b453065e3b9c1c84afef158b9149fe29aa54c90fc8fe2a8361a6a6880.SnoozeReminderRequestBuilder) {
    return ie274417b453065e3b9c1c84afef158b9149fe29aa54c90fc8fe2a8361a6a6880.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*ic09438710245b2671bf04abe5db2656009ca7799d1d80105aa23a063cf2f4e40.TentativelyAcceptRequestBuilder) {
    return ic09438710245b2671bf04abe5db2656009ca7799d1d80105aa23a063cf2f4e40.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
