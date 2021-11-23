package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i09c7e9344c4531879e9a4292bc9a67eb53529e726c1878c6c2e9e8e53916ba47 "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/send"
    i2632294208677b2133bf271ae434a941bb72968036431b6303130fd27a80ddb4 "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/createforward"
    i263d542cf24ae9e9ef78bc9a276ebed9e7e6590c4a4ae274bd71b3012abf44d2 "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/reply"
    i2bb36c6899b162cafc6e656fe32fc77e6a563cef2a267d7fb5f954e663b9e808 "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/value"
    i3933a4d952f201657cf6f03d324ef4e33d0780a4ea49f70704cf76a2c191f976 "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/copy"
    i4ccf4f74cd764907f4b93d428388f1e772673c2f57b25c776ed0040e37810d8e "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/replyall"
    i5d0b9e0e9542b209896282893fed24a1ce4587679e91e205e15df0f1aeb037ff "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/attachments"
    i638c3159dd70df11db383b956f2700e3cdf9726dede7a3126289ca8deebbba33 "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/multivalueextendedproperties"
    i83f515fce5f469eba14ae8086d23694bda995238a9cf6aeda9e4a77f60d46e51 "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/createreplyall"
    i9a10feb98b69df8b4be73cf64661437fb8eba6c417528ea15e3fa7f97690cf45 "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/createreply"
    ia12ddc34b64b94dfc125c48cb2ce271ab4c81b4e234623928b468956f0ac7218 "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/singlevalueextendedproperties"
    iac1ed303e6357b3776696be9bf22be9bc234c3414b61d5a3e53b788b022c3b55 "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/calendarsharingmessage"
    id541c444118e749467445e6ddae5eccd7ce1bcbde5376a875eb7a6ac14639c51 "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/extensions"
    iee30ea951ebbeb52a3eefc3847404284a2701aa3eee3ab7b930e1466e6871b06 "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/forward"
    ifb4c655ccae06417a7b44976ed7bbc4f45bb8f773bda157b976061d66736f4c7 "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/move"
    i66622497d9843fd844c973e3655196d0fcde41543efedcdb9bc3c54e70023033 "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/extensions/item"
    i6e6c8882e6f07057df4c3ec98b0922cb303131fe0d4d661ef64ab6392a2c2221 "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/attachments/item"
    ib81b123557158ed5006240abf09bd6ea17d88c3dbd6ea5625f18100a9f024c8a "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/multivalueextendedproperties/item"
    ic3747645bd85467205b599ea11e617092b39c2dcfaef23e7852445e46b0d20fa "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item/singlevalueextendedproperties/item"
)

// MessageRequestBuilder builds and executes requests for operations under \me\messages\{message-id}
type MessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MessageRequestBuilderDeleteOptions options for Delete
type MessageRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MessageRequestBuilderGetOptions options for Get
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
// MessageRequestBuilderGetQueryParameters the messages in a mailbox or folder. Read-only. Nullable.
type MessageRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select_escaped []string;
}
// MessageRequestBuilderPatchOptions options for Patch
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
func (m *MessageRequestBuilder) Attachments()(*i5d0b9e0e9542b209896282893fed24a1ce4587679e91e205e15df0f1aeb037ff.AttachmentsRequestBuilder) {
    return i5d0b9e0e9542b209896282893fed24a1ce4587679e91e205e15df0f1aeb037ff.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.messages.item.attachments.item collection
func (m *MessageRequestBuilder) AttachmentsById(id string)(*i6e6c8882e6f07057df4c3ec98b0922cb303131fe0d4d661ef64ab6392a2c2221.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i6e6c8882e6f07057df4c3ec98b0922cb303131fe0d4d661ef64ab6392a2c2221.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) CalendarSharingMessage()(*iac1ed303e6357b3776696be9bf22be9bc234c3414b61d5a3e53b788b022c3b55.CalendarSharingMessageRequestBuilder) {
    return iac1ed303e6357b3776696be9bf22be9bc234c3414b61d5a3e53b788b022c3b55.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessageRequestBuilderInternal instantiates a new MessageRequestBuilder and sets the default values.
func NewMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageRequestBuilder) {
    m := &MessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/messages/{message_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMessageRequestBuilder instantiates a new MessageRequestBuilder and sets the default values.
func NewMessageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessageRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MessageRequestBuilder) Content()(*i2bb36c6899b162cafc6e656fe32fc77e6a563cef2a267d7fb5f954e663b9e808.ContentRequestBuilder) {
    return i2bb36c6899b162cafc6e656fe32fc77e6a563cef2a267d7fb5f954e663b9e808.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Copy()(*i3933a4d952f201657cf6f03d324ef4e33d0780a4ea49f70704cf76a2c191f976.CopyRequestBuilder) {
    return i3933a4d952f201657cf6f03d324ef4e33d0780a4ea49f70704cf76a2c191f976.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the messages in a mailbox or folder. Read-only. Nullable.
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
func (m *MessageRequestBuilder) CreateForward()(*i2632294208677b2133bf271ae434a941bb72968036431b6303130fd27a80ddb4.CreateForwardRequestBuilder) {
    return i2632294208677b2133bf271ae434a941bb72968036431b6303130fd27a80ddb4.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageRequestBuilder) CreateGetRequestInformation(options *MessageRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the messages in a mailbox or folder. Read-only. Nullable.
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
func (m *MessageRequestBuilder) CreateReply()(*i9a10feb98b69df8b4be73cf64661437fb8eba6c417528ea15e3fa7f97690cf45.CreateReplyRequestBuilder) {
    return i9a10feb98b69df8b4be73cf64661437fb8eba6c417528ea15e3fa7f97690cf45.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) CreateReplyAll()(*i83f515fce5f469eba14ae8086d23694bda995238a9cf6aeda9e4a77f60d46e51.CreateReplyAllRequestBuilder) {
    return i83f515fce5f469eba14ae8086d23694bda995238a9cf6aeda9e4a77f60d46e51.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the messages in a mailbox or folder. Read-only. Nullable.
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
func (m *MessageRequestBuilder) Extensions()(*id541c444118e749467445e6ddae5eccd7ce1bcbde5376a875eb7a6ac14639c51.ExtensionsRequestBuilder) {
    return id541c444118e749467445e6ddae5eccd7ce1bcbde5376a875eb7a6ac14639c51.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.messages.item.extensions.item collection
func (m *MessageRequestBuilder) ExtensionsById(id string)(*i66622497d9843fd844c973e3655196d0fcde41543efedcdb9bc3c54e70023033.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i66622497d9843fd844c973e3655196d0fcde41543efedcdb9bc3c54e70023033.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) Forward()(*iee30ea951ebbeb52a3eefc3847404284a2701aa3eee3ab7b930e1466e6871b06.ForwardRequestBuilder) {
    return iee30ea951ebbeb52a3eefc3847404284a2701aa3eee3ab7b930e1466e6871b06.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the messages in a mailbox or folder. Read-only. Nullable.
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
func (m *MessageRequestBuilder) Move()(*ifb4c655ccae06417a7b44976ed7bbc4f45bb8f773bda157b976061d66736f4c7.MoveRequestBuilder) {
    return ifb4c655ccae06417a7b44976ed7bbc4f45bb8f773bda157b976061d66736f4c7.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) MultiValueExtendedProperties()(*i638c3159dd70df11db383b956f2700e3cdf9726dede7a3126289ca8deebbba33.MultiValueExtendedPropertiesRequestBuilder) {
    return i638c3159dd70df11db383b956f2700e3cdf9726dede7a3126289ca8deebbba33.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.messages.item.multiValueExtendedProperties.item collection
func (m *MessageRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ib81b123557158ed5006240abf09bd6ea17d88c3dbd6ea5625f18100a9f024c8a.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ib81b123557158ed5006240abf09bd6ea17d88c3dbd6ea5625f18100a9f024c8a.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the messages in a mailbox or folder. Read-only. Nullable.
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
func (m *MessageRequestBuilder) Reply()(*i263d542cf24ae9e9ef78bc9a276ebed9e7e6590c4a4ae274bd71b3012abf44d2.ReplyRequestBuilder) {
    return i263d542cf24ae9e9ef78bc9a276ebed9e7e6590c4a4ae274bd71b3012abf44d2.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) ReplyAll()(*i4ccf4f74cd764907f4b93d428388f1e772673c2f57b25c776ed0040e37810d8e.ReplyAllRequestBuilder) {
    return i4ccf4f74cd764907f4b93d428388f1e772673c2f57b25c776ed0040e37810d8e.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Send()(*i09c7e9344c4531879e9a4292bc9a67eb53529e726c1878c6c2e9e8e53916ba47.SendRequestBuilder) {
    return i09c7e9344c4531879e9a4292bc9a67eb53529e726c1878c6c2e9e8e53916ba47.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) SingleValueExtendedProperties()(*ia12ddc34b64b94dfc125c48cb2ce271ab4c81b4e234623928b468956f0ac7218.SingleValueExtendedPropertiesRequestBuilder) {
    return ia12ddc34b64b94dfc125c48cb2ce271ab4c81b4e234623928b468956f0ac7218.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.messages.item.singleValueExtendedProperties.item collection
func (m *MessageRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ic3747645bd85467205b599ea11e617092b39c2dcfaef23e7852445e46b0d20fa.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ic3747645bd85467205b599ea11e617092b39c2dcfaef23e7852445e46b0d20fa.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
