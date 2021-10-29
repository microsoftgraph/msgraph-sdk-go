package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1e31ad36dd6f6af890b8fbc16fdd724e1b79f825180e2146209442903e1536c6 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/send"
    i23d1871eaa3b34b9da4dc3013d78975116ba3c998a34b6b2f7deffdacf14481b "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/move"
    i4d76a5f9ebc4fd08800dbc2923dc874961ba25497bb7fa44bc7745d855311d6d "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/createreply"
    i525f8876c52a4815a231b5ca35b01865617a535594600f8665da18bdb784de49 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/value"
    i53c1e297b8c3f78ec3582370d703e4eb456f41af03120d845e2806968caea8a8 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/reply"
    i57c84c8fec460faeb8d03d0e3a5e76bff8ceceef2d1a73f98c319b311312185f "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/extensions"
    i7a919721ec6ea9b78407f2ee5aa295d1812c6812aeef0242dec74082f64d5fa2 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/attachments"
    i7aade3e497faeed2338d7b27d6d8b2216ae41fa1ae3717530eaf821301271040 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/replyall"
    i7e844353f50f778cff7d0fb3536d9e8900c8f5f4958bac6ec2331dabe2b1670c "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/createforward"
    i95c67b80ede95c8639d833166ec28903fde06719540ffe7742eec48fc08cd740 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/createreplyall"
    ia421f89e87068ec5411ac4214056c2997f61057486c146323e7ff0854e755cca "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/copy"
    iaf5c26a284e4a05c5a5a6cfa25d3c65b91541124829eb429f6ffc67343739406 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/singlevalueextendedproperties"
    ib4078a2858aac8ed3592dd9a4fd579ad952cb1df89a20fa52d2f42b6a1d69d91 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/calendarsharingmessage"
    idc279dd0f6da07bc29bf2a0f7b5d7528ee60dea48bfd2284a8f21c10e3e53e61 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/forward"
    ie70f42a3bdeda2a4c05e634ad649af47641c62c1ebdb49d12f0653327fce4e07 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/multivalueextendedproperties"
    i39886b4e20e3dea3e0ba21523a984da9b2ca189aff6561f3bc956ebf0d1c17ee "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/singlevalueextendedproperties/item"
    i93ef364c0f89b4997f731376d9fd8c3497ed63375a7efc0ee24e1a2122c20a64 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/attachments/item"
    id44a9e5325d96aaff187d75056e8b221eae0fe1190baf4c4c0318b6dfc483672 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/multivalueextendedproperties/item"
    idec392630c701cf349a9fadb8a05ca2f81715c4de9cfc77853598541c08f6920 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/messages/item/extensions/item"
)

// Builds and executes requests for operations under \me\mailFolders\{mailFolder-id}\messages\{message-id}
type MessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type MessageRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type MessageRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MessageRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The collection of messages in the mailFolder.
type MessageRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type MessageRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Message;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MessageRequestBuilder) Attachments()(*i7a919721ec6ea9b78407f2ee5aa295d1812c6812aeef0242dec74082f64d5fa2.AttachmentsRequestBuilder) {
    return i7a919721ec6ea9b78407f2ee5aa295d1812c6812aeef0242dec74082f64d5fa2.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.mailFolders.item.messages.item.attachments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *MessageRequestBuilder) AttachmentsById(id string)(*i93ef364c0f89b4997f731376d9fd8c3497ed63375a7efc0ee24e1a2122c20a64.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i93ef364c0f89b4997f731376d9fd8c3497ed63375a7efc0ee24e1a2122c20a64.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) CalendarSharingMessage()(*ib4078a2858aac8ed3592dd9a4fd579ad952cb1df89a20fa52d2f42b6a1d69d91.CalendarSharingMessageRequestBuilder) {
    return ib4078a2858aac8ed3592dd9a4fd579ad952cb1df89a20fa52d2f42b6a1d69d91.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new MessageRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageRequestBuilder) {
    m := &MessageRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/mailFolders/{mailFolder_id}/messages/{message_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new MessageRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewMessageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessageRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MessageRequestBuilder) Content()(*i525f8876c52a4815a231b5ca35b01865617a535594600f8665da18bdb784de49.ContentRequestBuilder) {
    return i525f8876c52a4815a231b5ca35b01865617a535594600f8665da18bdb784de49.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Copy()(*ia421f89e87068ec5411ac4214056c2997f61057486c146323e7ff0854e755cca.CopyRequestBuilder) {
    return ia421f89e87068ec5411ac4214056c2997f61057486c146323e7ff0854e755cca.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The collection of messages in the mailFolder.
// Parameters:
//  - options : Options for the request
func (m *MessageRequestBuilder) CreateDeleteRequestInformation(options *MessageRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MessageRequestBuilder) CreateForward()(*i7e844353f50f778cff7d0fb3536d9e8900c8f5f4958bac6ec2331dabe2b1670c.CreateForwardRequestBuilder) {
    return i7e844353f50f778cff7d0fb3536d9e8900c8f5f4958bac6ec2331dabe2b1670c.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The collection of messages in the mailFolder.
// Parameters:
//  - options : Options for the request
func (m *MessageRequestBuilder) CreateGetRequestInformation(options *MessageRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// The collection of messages in the mailFolder.
// Parameters:
//  - options : Options for the request
func (m *MessageRequestBuilder) CreatePatchRequestInformation(options *MessageRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MessageRequestBuilder) CreateReply()(*i4d76a5f9ebc4fd08800dbc2923dc874961ba25497bb7fa44bc7745d855311d6d.CreateReplyRequestBuilder) {
    return i4d76a5f9ebc4fd08800dbc2923dc874961ba25497bb7fa44bc7745d855311d6d.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) CreateReplyAll()(*i95c67b80ede95c8639d833166ec28903fde06719540ffe7742eec48fc08cd740.CreateReplyAllRequestBuilder) {
    return i95c67b80ede95c8639d833166ec28903fde06719540ffe7742eec48fc08cd740.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The collection of messages in the mailFolder.
// Parameters:
//  - options : Options for the request
func (m *MessageRequestBuilder) Delete(options *MessageRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *MessageRequestBuilder) Extensions()(*i57c84c8fec460faeb8d03d0e3a5e76bff8ceceef2d1a73f98c319b311312185f.ExtensionsRequestBuilder) {
    return i57c84c8fec460faeb8d03d0e3a5e76bff8ceceef2d1a73f98c319b311312185f.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.mailFolders.item.messages.item.extensions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *MessageRequestBuilder) ExtensionsById(id string)(*idec392630c701cf349a9fadb8a05ca2f81715c4de9cfc77853598541c08f6920.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return idec392630c701cf349a9fadb8a05ca2f81715c4de9cfc77853598541c08f6920.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) Forward()(*idc279dd0f6da07bc29bf2a0f7b5d7528ee60dea48bfd2284a8f21c10e3e53e61.ForwardRequestBuilder) {
    return idc279dd0f6da07bc29bf2a0f7b5d7528ee60dea48bfd2284a8f21c10e3e53e61.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The collection of messages in the mailFolder.
// Parameters:
//  - options : Options for the request
func (m *MessageRequestBuilder) Get(options *MessageRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Message, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewMessage() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Message), nil
}
func (m *MessageRequestBuilder) Move()(*i23d1871eaa3b34b9da4dc3013d78975116ba3c998a34b6b2f7deffdacf14481b.MoveRequestBuilder) {
    return i23d1871eaa3b34b9da4dc3013d78975116ba3c998a34b6b2f7deffdacf14481b.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) MultiValueExtendedProperties()(*ie70f42a3bdeda2a4c05e634ad649af47641c62c1ebdb49d12f0653327fce4e07.MultiValueExtendedPropertiesRequestBuilder) {
    return ie70f42a3bdeda2a4c05e634ad649af47641c62c1ebdb49d12f0653327fce4e07.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.mailFolders.item.messages.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *MessageRequestBuilder) MultiValueExtendedPropertiesById(id string)(*id44a9e5325d96aaff187d75056e8b221eae0fe1190baf4c4c0318b6dfc483672.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return id44a9e5325d96aaff187d75056e8b221eae0fe1190baf4c4c0318b6dfc483672.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The collection of messages in the mailFolder.
// Parameters:
//  - options : Options for the request
func (m *MessageRequestBuilder) Patch(options *MessageRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *MessageRequestBuilder) Reply()(*i53c1e297b8c3f78ec3582370d703e4eb456f41af03120d845e2806968caea8a8.ReplyRequestBuilder) {
    return i53c1e297b8c3f78ec3582370d703e4eb456f41af03120d845e2806968caea8a8.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) ReplyAll()(*i7aade3e497faeed2338d7b27d6d8b2216ae41fa1ae3717530eaf821301271040.ReplyAllRequestBuilder) {
    return i7aade3e497faeed2338d7b27d6d8b2216ae41fa1ae3717530eaf821301271040.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Send()(*i1e31ad36dd6f6af890b8fbc16fdd724e1b79f825180e2146209442903e1536c6.SendRequestBuilder) {
    return i1e31ad36dd6f6af890b8fbc16fdd724e1b79f825180e2146209442903e1536c6.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) SingleValueExtendedProperties()(*iaf5c26a284e4a05c5a5a6cfa25d3c65b91541124829eb429f6ffc67343739406.SingleValueExtendedPropertiesRequestBuilder) {
    return iaf5c26a284e4a05c5a5a6cfa25d3c65b91541124829eb429f6ffc67343739406.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.mailFolders.item.messages.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *MessageRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i39886b4e20e3dea3e0ba21523a984da9b2ca189aff6561f3bc956ebf0d1c17ee.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i39886b4e20e3dea3e0ba21523a984da9b2ca189aff6561f3bc956ebf0d1c17ee.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
