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

// EventItemRequestBuilder builds and executes requests for operations under \me\events\{event-id}
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventItemRequestBuilderDeleteOptions options for Delete
type EventItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventItemRequestBuilderGetOptions options for Get
type EventItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EventItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventItemRequestBuilderGetQueryParameters the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// EventItemRequestBuilderPatchOptions options for Patch
type EventItemRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventItemRequestBuilder) Accept()(*i1426802efbd0f5039b4ed42db7c3a88b7863b74e5a46d0d4bc134d1e85fbd90c.AcceptRequestBuilder) {
    return i1426802efbd0f5039b4ed42db7c3a88b7863b74e5a46d0d4bc134d1e85fbd90c.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i44038414241e1d1ac04a807191ed07589543d0a83a4e70fb548ebd312bf47055.AttachmentsRequestBuilder) {
    return i44038414241e1d1ac04a807191ed07589543d0a83a4e70fb548ebd312bf47055.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ia9d78ab81a6804e22ae5ba33ac5f181e55a589f99d2e4208a5aa8aa0856893e1.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ia9d78ab81a6804e22ae5ba33ac5f181e55a589f99d2e4208a5aa8aa0856893e1.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i4209b20d5ab4261e9fa52d2682ea4586281037bc7cc1dfad2dcf1191b7d6dbc2.CalendarRequestBuilder) {
    return i4209b20d5ab4261e9fa52d2682ea4586281037bc7cc1dfad2dcf1191b7d6dbc2.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*ie36ea007a91599e24d84e7f363e1abf4b6e65cf46bfa7af40b5e30187cfa1602.CancelRequestBuilder) {
    return ie36ea007a91599e24d84e7f363e1abf4b6e65cf46bfa7af40b5e30187cfa1602.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/events/{event_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation(options *EventItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformation(options *EventItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(options *EventItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *EventItemRequestBuilder) Decline()(*i26b1c7cc63b87d2dd7618405741954826378e5687e20fdad5c9aa35ed448f8ff.DeclineRequestBuilder) {
    return i26b1c7cc63b87d2dd7618405741954826378e5687e20fdad5c9aa35ed448f8ff.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) DismissReminder()(*i2badd6557828e3ec37a7444fdf50ab5f7b8fdb7ae4b2a7ee749da38e2e461018.DismissReminderRequestBuilder) {
    return i2badd6557828e3ec37a7444fdf50ab5f7b8fdb7ae4b2a7ee749da38e2e461018.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i95058b33c206d566c03da284448a36f4ab33704d456ac0a7ebb2cd7b6cd2fb0f.ExtensionsRequestBuilder) {
    return i95058b33c206d566c03da284448a36f4ab33704d456ac0a7ebb2cd7b6cd2fb0f.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i8414fee0ad5be4f22aa3ea04ac1422e9124533bafff06b8fcddc21e0c7ca538a.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i8414fee0ad5be4f22aa3ea04ac1422e9124533bafff06b8fcddc21e0c7ca538a.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*idcc656e3c5d595c7e17926a2796ec2029a16b9df71a54b6dd2ced6343b191c8d.ForwardRequestBuilder) {
    return idcc656e3c5d595c7e17926a2796ec2029a16b9df71a54b6dd2ced6343b191c8d.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEvent() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event), nil
}
func (m *EventItemRequestBuilder) Instances()(*ib59086c56b738f1f0717fb1ceaf00dffeac922ef570d9103a1ff449f48fca45d.InstancesRequestBuilder) {
    return ib59086c56b738f1f0717fb1ceaf00dffeac922ef570d9103a1ff449f48fca45d.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i4cdf65b58617609797304bf46ce9451a705dd1606550f29e20aa67daf8236b42.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i4cdf65b58617609797304bf46ce9451a705dd1606550f29e20aa67daf8236b42.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i54078d1892815baae8a8fd0cece8461711b63ba916021bd5ced1c30195a1a698.MultiValueExtendedPropertiesRequestBuilder) {
    return i54078d1892815baae8a8fd0cece8461711b63ba916021bd5ced1c30195a1a698.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i2a46836ccae42574e0c03c77e8963f70e5d011a92ca307f636d258be89ac1916.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i2a46836ccae42574e0c03c77e8963f70e5d011a92ca307f636d258be89ac1916.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*id5711ae3f0bfb22fb24dc70cb4be78fa08cb4a509118586258c0f33adb016cf2.SingleValueExtendedPropertiesRequestBuilder) {
    return id5711ae3f0bfb22fb24dc70cb4be78fa08cb4a509118586258c0f33adb016cf2.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ic43a8f89b11e7c86745f7bafd213b5d40252e1081389e9965812ca0da060a016.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ic43a8f89b11e7c86745f7bafd213b5d40252e1081389e9965812ca0da060a016.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i5726cc676e11fe9fa90959a9a58ddf2f16c34f2a86c819a81f6f2ce4937b710a.SnoozeReminderRequestBuilder) {
    return i5726cc676e11fe9fa90959a9a58ddf2f16c34f2a86c819a81f6f2ce4937b710a.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i774cf8a56cdf6984d443e1380e0d0dcd2ca8d118dde0425d0b60e04c5f1c436c.TentativelyAcceptRequestBuilder) {
    return i774cf8a56cdf6984d443e1380e0d0dcd2ca8d118dde0425d0b60e04c5f1c436c.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
