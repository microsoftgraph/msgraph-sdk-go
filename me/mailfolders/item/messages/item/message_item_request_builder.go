package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
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

// MessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
type MessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MessageItemRequestBuilderDeleteOptions options for Delete
type MessageItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MessageItemRequestBuilderGetOptions options for Get
type MessageItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MessageItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MessageItemRequestBuilderGetQueryParameters the collection of messages in the mailFolder.
type MessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// MessageItemRequestBuilderPatchOptions options for Patch
type MessageItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Messageable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MessageItemRequestBuilder) Attachments()(*i7a919721ec6ea9b78407f2ee5aa295d1812c6812aeef0242dec74082f64d5fa2.AttachmentsRequestBuilder) {
    return i7a919721ec6ea9b78407f2ee5aa295d1812c6812aeef0242dec74082f64d5fa2.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.mailFolders.item.messages.item.attachments.item collection
func (m *MessageItemRequestBuilder) AttachmentsById(id string)(*i93ef364c0f89b4997f731376d9fd8c3497ed63375a7efc0ee24e1a2122c20a64.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i93ef364c0f89b4997f731376d9fd8c3497ed63375a7efc0ee24e1a2122c20a64.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) CalendarSharingMessage()(*ib4078a2858aac8ed3592dd9a4fd579ad952cb1df89a20fa52d2f42b6a1d69d91.CalendarSharingMessageRequestBuilder) {
    return ib4078a2858aac8ed3592dd9a4fd579ad952cb1df89a20fa52d2f42b6a1d69d91.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageItemRequestBuilder) {
    m := &MessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mailFolders/{mailFolder_id}/messages/{message_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMessageItemRequestBuilder instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MessageItemRequestBuilder) Content()(*i525f8876c52a4815a231b5ca35b01865617a535594600f8665da18bdb784de49.ContentRequestBuilder) {
    return i525f8876c52a4815a231b5ca35b01865617a535594600f8665da18bdb784de49.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Copy()(*ia421f89e87068ec5411ac4214056c2997f61057486c146323e7ff0854e755cca.CopyRequestBuilder) {
    return ia421f89e87068ec5411ac4214056c2997f61057486c146323e7ff0854e755cca.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property messages for me
func (m *MessageItemRequestBuilder) CreateDeleteRequestInformation(options *MessageItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MessageItemRequestBuilder) CreateForward()(*i7e844353f50f778cff7d0fb3536d9e8900c8f5f4958bac6ec2331dabe2b1670c.CreateForwardRequestBuilder) {
    return i7e844353f50f778cff7d0fb3536d9e8900c8f5f4958bac6ec2331dabe2b1670c.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the collection of messages in the mailFolder.
func (m *MessageItemRequestBuilder) CreateGetRequestInformation(options *MessageItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property messages in me
func (m *MessageItemRequestBuilder) CreatePatchRequestInformation(options *MessageItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MessageItemRequestBuilder) CreateReply()(*i4d76a5f9ebc4fd08800dbc2923dc874961ba25497bb7fa44bc7745d855311d6d.CreateReplyRequestBuilder) {
    return i4d76a5f9ebc4fd08800dbc2923dc874961ba25497bb7fa44bc7745d855311d6d.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) CreateReplyAll()(*i95c67b80ede95c8639d833166ec28903fde06719540ffe7742eec48fc08cd740.CreateReplyAllRequestBuilder) {
    return i95c67b80ede95c8639d833166ec28903fde06719540ffe7742eec48fc08cd740.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property messages for me
func (m *MessageItemRequestBuilder) Delete(options *MessageItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *MessageItemRequestBuilder) Extensions()(*i57c84c8fec460faeb8d03d0e3a5e76bff8ceceef2d1a73f98c319b311312185f.ExtensionsRequestBuilder) {
    return i57c84c8fec460faeb8d03d0e3a5e76bff8ceceef2d1a73f98c319b311312185f.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.mailFolders.item.messages.item.extensions.item collection
func (m *MessageItemRequestBuilder) ExtensionsById(id string)(*idec392630c701cf349a9fadb8a05ca2f81715c4de9cfc77853598541c08f6920.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return idec392630c701cf349a9fadb8a05ca2f81715c4de9cfc77853598541c08f6920.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Forward()(*idc279dd0f6da07bc29bf2a0f7b5d7528ee60dea48bfd2284a8f21c10e3e53e61.ForwardRequestBuilder) {
    return idc279dd0f6da07bc29bf2a0f7b5d7528ee60dea48bfd2284a8f21c10e3e53e61.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of messages in the mailFolder.
func (m *MessageItemRequestBuilder) Get(options *MessageItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Messageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateMessageFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Messageable), nil
}
func (m *MessageItemRequestBuilder) Move()(*i23d1871eaa3b34b9da4dc3013d78975116ba3c998a34b6b2f7deffdacf14481b.MoveRequestBuilder) {
    return i23d1871eaa3b34b9da4dc3013d78975116ba3c998a34b6b2f7deffdacf14481b.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) MultiValueExtendedProperties()(*ie70f42a3bdeda2a4c05e634ad649af47641c62c1ebdb49d12f0653327fce4e07.MultiValueExtendedPropertiesRequestBuilder) {
    return ie70f42a3bdeda2a4c05e634ad649af47641c62c1ebdb49d12f0653327fce4e07.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.mailFolders.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*id44a9e5325d96aaff187d75056e8b221eae0fe1190baf4c4c0318b6dfc483672.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return id44a9e5325d96aaff187d75056e8b221eae0fe1190baf4c4c0318b6dfc483672.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in me
func (m *MessageItemRequestBuilder) Patch(options *MessageItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *MessageItemRequestBuilder) Reply()(*i53c1e297b8c3f78ec3582370d703e4eb456f41af03120d845e2806968caea8a8.ReplyRequestBuilder) {
    return i53c1e297b8c3f78ec3582370d703e4eb456f41af03120d845e2806968caea8a8.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) ReplyAll()(*i7aade3e497faeed2338d7b27d6d8b2216ae41fa1ae3717530eaf821301271040.ReplyAllRequestBuilder) {
    return i7aade3e497faeed2338d7b27d6d8b2216ae41fa1ae3717530eaf821301271040.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Send()(*i1e31ad36dd6f6af890b8fbc16fdd724e1b79f825180e2146209442903e1536c6.SendRequestBuilder) {
    return i1e31ad36dd6f6af890b8fbc16fdd724e1b79f825180e2146209442903e1536c6.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) SingleValueExtendedProperties()(*iaf5c26a284e4a05c5a5a6cfa25d3c65b91541124829eb429f6ffc67343739406.SingleValueExtendedPropertiesRequestBuilder) {
    return iaf5c26a284e4a05c5a5a6cfa25d3c65b91541124829eb429f6ffc67343739406.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.mailFolders.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i39886b4e20e3dea3e0ba21523a984da9b2ca189aff6561f3bc956ebf0d1c17ee.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i39886b4e20e3dea3e0ba21523a984da9b2ca189aff6561f3bc956ebf0d1c17ee.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
