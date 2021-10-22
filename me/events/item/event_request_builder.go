package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1426802efbd0f5039b4ed42db7c3a88b7863b74e5a46d0d4bc134d1e85fbd90c "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/accept"
    i26b1c7cc63b87d2dd7618405741954826378e5687e20fdad5c9aa35ed448f8ff "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/decline"
    i2badd6557828e3ec37a7444fdf50ab5f7b8fdb7ae4b2a7ee749da38e2e461018 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/dismissreminder"
    i4209b20d5ab4261e9fa52d2682ea4586281037bc7cc1dfad2dcf1191b7d6dbc2 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar"
    i44038414241e1d1ac04a807191ed07589543d0a83a4e70fb548ebd312bf47055 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/attachments"
    i54078d1892815baae8a8fd0cece8461711b63ba916021bd5ced1c30195a1a698 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/multivalueextendedproperties"
    i5726cc676e11fe9fa90959a9a58ddf2f16c34f2a86c819a81f6f2ce4937b710a "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/snoozereminder"
    i774cf8a56cdf6984d443e1380e0d0dcd2ca8d118dde0425d0b60e04c5f1c436c "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/tentativelyaccept"
    i95058b33c206d566c03da284448a36f4ab33704d456ac0a7ebb2cd7b6cd2fb0f "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/extensions"
    ib59086c56b738f1f0717fb1ceaf00dffeac922ef570d9103a1ff449f48fca45d "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances"
    id5711ae3f0bfb22fb24dc70cb4be78fa08cb4a509118586258c0f33adb016cf2 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/singlevalueextendedproperties"
    idcc656e3c5d595c7e17926a2796ec2029a16b9df71a54b6dd2ced6343b191c8d "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/forward"
    ie36ea007a91599e24d84e7f363e1abf4b6e65cf46bfa7af40b5e30187cfa1602 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/cancel"
    i2a46836ccae42574e0c03c77e8963f70e5d011a92ca307f636d258be89ac1916 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/multivalueextendedproperties/item"
    i4cdf65b58617609797304bf46ce9451a705dd1606550f29e20aa67daf8236b42 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/instances/item"
    i8414fee0ad5be4f22aa3ea04ac1422e9124533bafff06b8fcddc21e0c7ca538a "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/extensions/item"
    ia9d78ab81a6804e22ae5ba33ac5f181e55a589f99d2e4208a5aa8aa0856893e1 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/attachments/item"
    ic43a8f89b11e7c86745f7bafd213b5d40252e1081389e9965812ca0da060a016 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/singlevalueextendedproperties/item"
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
func (m *EventRequestBuilder) Accept()(*i1426802efbd0f5039b4ed42db7c3a88b7863b74e5a46d0d4bc134d1e85fbd90c.AcceptRequestBuilder) {
    return i1426802efbd0f5039b4ed42db7c3a88b7863b74e5a46d0d4bc134d1e85fbd90c.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i44038414241e1d1ac04a807191ed07589543d0a83a4e70fb548ebd312bf47055.AttachmentsRequestBuilder) {
    return i44038414241e1d1ac04a807191ed07589543d0a83a4e70fb548ebd312bf47055.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) AttachmentsById(id string)(*ia9d78ab81a6804e22ae5ba33ac5f181e55a589f99d2e4208a5aa8aa0856893e1.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ia9d78ab81a6804e22ae5ba33ac5f181e55a589f99d2e4208a5aa8aa0856893e1.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i4209b20d5ab4261e9fa52d2682ea4586281037bc7cc1dfad2dcf1191b7d6dbc2.CalendarRequestBuilder) {
    return i4209b20d5ab4261e9fa52d2682ea4586281037bc7cc1dfad2dcf1191b7d6dbc2.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*ie36ea007a91599e24d84e7f363e1abf4b6e65cf46bfa7af40b5e30187cfa1602.CancelRequestBuilder) {
    return ie36ea007a91599e24d84e7f363e1abf4b6e65cf46bfa7af40b5e30187cfa1602.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/events/{event_id}{?select,expand}";
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
func (m *EventRequestBuilder) Decline()(*i26b1c7cc63b87d2dd7618405741954826378e5687e20fdad5c9aa35ed448f8ff.DeclineRequestBuilder) {
    return i26b1c7cc63b87d2dd7618405741954826378e5687e20fdad5c9aa35ed448f8ff.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) DismissReminder()(*i2badd6557828e3ec37a7444fdf50ab5f7b8fdb7ae4b2a7ee749da38e2e461018.DismissReminderRequestBuilder) {
    return i2badd6557828e3ec37a7444fdf50ab5f7b8fdb7ae4b2a7ee749da38e2e461018.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i95058b33c206d566c03da284448a36f4ab33704d456ac0a7ebb2cd7b6cd2fb0f.ExtensionsRequestBuilder) {
    return i95058b33c206d566c03da284448a36f4ab33704d456ac0a7ebb2cd7b6cd2fb0f.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExtensionsById(id string)(*i8414fee0ad5be4f22aa3ea04ac1422e9124533bafff06b8fcddc21e0c7ca538a.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i8414fee0ad5be4f22aa3ea04ac1422e9124533bafff06b8fcddc21e0c7ca538a.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*idcc656e3c5d595c7e17926a2796ec2029a16b9df71a54b6dd2ced6343b191c8d.ForwardRequestBuilder) {
    return idcc656e3c5d595c7e17926a2796ec2029a16b9df71a54b6dd2ced6343b191c8d.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) Instances()(*ib59086c56b738f1f0717fb1ceaf00dffeac922ef570d9103a1ff449f48fca45d.InstancesRequestBuilder) {
    return ib59086c56b738f1f0717fb1ceaf00dffeac922ef570d9103a1ff449f48fca45d.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) InstancesById(id string)(*i4cdf65b58617609797304bf46ce9451a705dd1606550f29e20aa67daf8236b42.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i4cdf65b58617609797304bf46ce9451a705dd1606550f29e20aa67daf8236b42.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i54078d1892815baae8a8fd0cece8461711b63ba916021bd5ced1c30195a1a698.MultiValueExtendedPropertiesRequestBuilder) {
    return i54078d1892815baae8a8fd0cece8461711b63ba916021bd5ced1c30195a1a698.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i2a46836ccae42574e0c03c77e8963f70e5d011a92ca307f636d258be89ac1916.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i2a46836ccae42574e0c03c77e8963f70e5d011a92ca307f636d258be89ac1916.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*id5711ae3f0bfb22fb24dc70cb4be78fa08cb4a509118586258c0f33adb016cf2.SingleValueExtendedPropertiesRequestBuilder) {
    return id5711ae3f0bfb22fb24dc70cb4be78fa08cb4a509118586258c0f33adb016cf2.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ic43a8f89b11e7c86745f7bafd213b5d40252e1081389e9965812ca0da060a016.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ic43a8f89b11e7c86745f7bafd213b5d40252e1081389e9965812ca0da060a016.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i5726cc676e11fe9fa90959a9a58ddf2f16c34f2a86c819a81f6f2ce4937b710a.SnoozeReminderRequestBuilder) {
    return i5726cc676e11fe9fa90959a9a58ddf2f16c34f2a86c819a81f6f2ce4937b710a.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i774cf8a56cdf6984d443e1380e0d0dcd2ca8d118dde0425d0b60e04c5f1c436c.TentativelyAcceptRequestBuilder) {
    return i774cf8a56cdf6984d443e1380e0d0dcd2ca8d118dde0425d0b60e04c5f1c436c.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
