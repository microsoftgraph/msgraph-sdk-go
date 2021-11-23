package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i062a3ff4d1d6529c88f923d81d7a2d73cc91a7aee2ce785d6ea82b627d35e210 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/move"
    i19224a4e273490c542daba438c5ae16a7a3f0e0c07160c68599fe9e49e6d6db5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/replyall"
    i2dfc35da4c65425a2b10fff308a7173c784eb400aa7a77530c2b84b758ca6d16 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/calendarsharingmessage"
    i325c2d93beaffbeb094ac9cd1175734c42683a7c63d9419aadcd8a0a9a992575 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/attachments"
    i5de64d7ba319e004e4fad5389782964371cfeaa2956093290434d9564c85bcab "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/createreply"
    i6ad07a59d46ab03d44f649d85e219feb09f1b9de5800cb7b9590ff92347ff356 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/createforward"
    i989a6eef61b5f2d97064f6b03429c220fba6a058beb4c94807f0f085b89b103a "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/singlevalueextendedproperties"
    ia30921aeaaa3c5a09c4d881ec3a58b3c88454249f15e78743583738180fd9b67 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/reply"
    ib23139dd32ff7cd0654948f9728fdf3f7b9c4d5c115350a61811f192243497be "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/value"
    ib4ec94e919a89eabb07aab3948adafcf0ad3ab5c11f2849ce02c0bedfac417db "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/forward"
    ib4f36fecd48150daa2f2c8cf15e035efd83748164c0c3d487b3ed0f3b758f30c "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/send"
    ib514357c2438d9d84bb6a90ddd348310bf097fc75d97905b8420f5cc3de23bba "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/createreplyall"
    icc5059e80c114c09025d473f2afa217eb771396c5f74d5be284ca48ae46d0cab "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/copy"
    id0b54dcecaa8542b3fa53a87c14da851c997cae296f64c5235715ba05fca4def "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/extensions"
    iee1de9a57b157c37bef3dd47f8c4371431b421428ba049687d93529e0f49bbe3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/multivalueextendedproperties"
    i0d29fac26fc9096d9957be27133c2cb630d1d9aff452d0071baddca02c01fcf5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/multivalueextendedproperties/item"
    i7406385c3a11e571b5a4a129b5889d6096ac6a18d6213a54ebcf5e3200bdec80 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/singlevalueextendedproperties/item"
    iaaf4ea6378bf11c338d473c01fe02eb0bcbab7e0785a2c2b63a8d33693b642b7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/attachments/item"
    iccad5f03bb49efc1650d1fe0e050ffa514f5fa92e483e5f27b73aa5dfbf0c755 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item/extensions/item"
)

// messageRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}
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
// messageRequestBuilderGetQueryParameters the collection of messages in the mailFolder.
type MessageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
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
func (m *MessageRequestBuilder) Attachments()(*i325c2d93beaffbeb094ac9cd1175734c42683a7c63d9419aadcd8a0a9a992575.AttachmentsRequestBuilder) {
    return i325c2d93beaffbeb094ac9cd1175734c42683a7c63d9419aadcd8a0a9a992575.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.messages.item.attachments.item collection
func (m *MessageRequestBuilder) AttachmentsById(id string)(*iaaf4ea6378bf11c338d473c01fe02eb0bcbab7e0785a2c2b63a8d33693b642b7.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return iaaf4ea6378bf11c338d473c01fe02eb0bcbab7e0785a2c2b63a8d33693b642b7.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) CalendarSharingMessage()(*i2dfc35da4c65425a2b10fff308a7173c784eb400aa7a77530c2b84b758ca6d16.CalendarSharingMessageRequestBuilder) {
    return i2dfc35da4c65425a2b10fff308a7173c784eb400aa7a77530c2b84b758ca6d16.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessageRequestBuilderInternal instantiates a new MessageRequestBuilder and sets the default values.
func NewMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageRequestBuilder) {
    m := &MessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}/messages/{message_id}{?select,expand}";
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
func (m *MessageRequestBuilder) Content()(*ib23139dd32ff7cd0654948f9728fdf3f7b9c4d5c115350a61811f192243497be.ContentRequestBuilder) {
    return ib23139dd32ff7cd0654948f9728fdf3f7b9c4d5c115350a61811f192243497be.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Copy()(*icc5059e80c114c09025d473f2afa217eb771396c5f74d5be284ca48ae46d0cab.CopyRequestBuilder) {
    return icc5059e80c114c09025d473f2afa217eb771396c5f74d5be284ca48ae46d0cab.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the collection of messages in the mailFolder.
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
func (m *MessageRequestBuilder) CreateForward()(*i6ad07a59d46ab03d44f649d85e219feb09f1b9de5800cb7b9590ff92347ff356.CreateForwardRequestBuilder) {
    return i6ad07a59d46ab03d44f649d85e219feb09f1b9de5800cb7b9590ff92347ff356.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the collection of messages in the mailFolder.
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
// CreatePatchRequestInformation the collection of messages in the mailFolder.
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
func (m *MessageRequestBuilder) CreateReply()(*i5de64d7ba319e004e4fad5389782964371cfeaa2956093290434d9564c85bcab.CreateReplyRequestBuilder) {
    return i5de64d7ba319e004e4fad5389782964371cfeaa2956093290434d9564c85bcab.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) CreateReplyAll()(*ib514357c2438d9d84bb6a90ddd348310bf097fc75d97905b8420f5cc3de23bba.CreateReplyAllRequestBuilder) {
    return ib514357c2438d9d84bb6a90ddd348310bf097fc75d97905b8420f5cc3de23bba.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the collection of messages in the mailFolder.
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
func (m *MessageRequestBuilder) Extensions()(*id0b54dcecaa8542b3fa53a87c14da851c997cae296f64c5235715ba05fca4def.ExtensionsRequestBuilder) {
    return id0b54dcecaa8542b3fa53a87c14da851c997cae296f64c5235715ba05fca4def.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.messages.item.extensions.item collection
func (m *MessageRequestBuilder) ExtensionsById(id string)(*iccad5f03bb49efc1650d1fe0e050ffa514f5fa92e483e5f27b73aa5dfbf0c755.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return iccad5f03bb49efc1650d1fe0e050ffa514f5fa92e483e5f27b73aa5dfbf0c755.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) Forward()(*ib4ec94e919a89eabb07aab3948adafcf0ad3ab5c11f2849ce02c0bedfac417db.ForwardRequestBuilder) {
    return ib4ec94e919a89eabb07aab3948adafcf0ad3ab5c11f2849ce02c0bedfac417db.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of messages in the mailFolder.
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
func (m *MessageRequestBuilder) Move()(*i062a3ff4d1d6529c88f923d81d7a2d73cc91a7aee2ce785d6ea82b627d35e210.MoveRequestBuilder) {
    return i062a3ff4d1d6529c88f923d81d7a2d73cc91a7aee2ce785d6ea82b627d35e210.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) MultiValueExtendedProperties()(*iee1de9a57b157c37bef3dd47f8c4371431b421428ba049687d93529e0f49bbe3.MultiValueExtendedPropertiesRequestBuilder) {
    return iee1de9a57b157c37bef3dd47f8c4371431b421428ba049687d93529e0f49bbe3.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0d29fac26fc9096d9957be27133c2cb630d1d9aff452d0071baddca02c01fcf5.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i0d29fac26fc9096d9957be27133c2cb630d1d9aff452d0071baddca02c01fcf5.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the collection of messages in the mailFolder.
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
func (m *MessageRequestBuilder) Reply()(*ia30921aeaaa3c5a09c4d881ec3a58b3c88454249f15e78743583738180fd9b67.ReplyRequestBuilder) {
    return ia30921aeaaa3c5a09c4d881ec3a58b3c88454249f15e78743583738180fd9b67.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) ReplyAll()(*i19224a4e273490c542daba438c5ae16a7a3f0e0c07160c68599fe9e49e6d6db5.ReplyAllRequestBuilder) {
    return i19224a4e273490c542daba438c5ae16a7a3f0e0c07160c68599fe9e49e6d6db5.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Send()(*ib4f36fecd48150daa2f2c8cf15e035efd83748164c0c3d487b3ed0f3b758f30c.SendRequestBuilder) {
    return ib4f36fecd48150daa2f2c8cf15e035efd83748164c0c3d487b3ed0f3b758f30c.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) SingleValueExtendedProperties()(*i989a6eef61b5f2d97064f6b03429c220fba6a058beb4c94807f0f085b89b103a.SingleValueExtendedPropertiesRequestBuilder) {
    return i989a6eef61b5f2d97064f6b03429c220fba6a058beb4c94807f0f085b89b103a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7406385c3a11e571b5a4a129b5889d6096ac6a18d6213a54ebcf5e3200bdec80.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i7406385c3a11e571b5a4a129b5889d6096ac6a18d6213a54ebcf5e3200bdec80.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
