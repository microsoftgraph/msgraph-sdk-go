package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
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

// MessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
type MessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MessageItemRequestBuilderGetQueryParameters the collection of messages in the mailFolder.
type MessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MessageItemRequestBuilderGetQueryParameters
}
// MessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments the attachments property
func (m *MessageItemRequestBuilder) Attachments()(*i325c2d93beaffbeb094ac9cd1175734c42683a7c63d9419aadcd8a0a9a992575.AttachmentsRequestBuilder) {
    return i325c2d93beaffbeb094ac9cd1175734c42683a7c63d9419aadcd8a0a9a992575.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.messages.item.attachments.item collection
func (m *MessageItemRequestBuilder) AttachmentsById(id string)(*iaaf4ea6378bf11c338d473c01fe02eb0bcbab7e0785a2c2b63a8d33693b642b7.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return iaaf4ea6378bf11c338d473c01fe02eb0bcbab7e0785a2c2b63a8d33693b642b7.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarSharingMessage the calendarSharingMessage property
func (m *MessageItemRequestBuilder) CalendarSharingMessage()(*i2dfc35da4c65425a2b10fff308a7173c784eb400aa7a77530c2b84b758ca6d16.CalendarSharingMessageRequestBuilder) {
    return i2dfc35da4c65425a2b10fff308a7173c784eb400aa7a77530c2b84b758ca6d16.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessageItemRequestBuilder) {
    m := &MessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMessageItemRequestBuilder instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the Content property
func (m *MessageItemRequestBuilder) Content()(*ib23139dd32ff7cd0654948f9728fdf3f7b9c4d5c115350a61811f192243497be.ContentRequestBuilder) {
    return ib23139dd32ff7cd0654948f9728fdf3f7b9c4d5c115350a61811f192243497be.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy the copy property
func (m *MessageItemRequestBuilder) Copy()(*icc5059e80c114c09025d473f2afa217eb771396c5f74d5be284ca48ae46d0cab.CopyRequestBuilder) {
    return icc5059e80c114c09025d473f2afa217eb771396c5f74d5be284ca48ae46d0cab.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property messages for users
func (m *MessageItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property messages for users
func (m *MessageItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *MessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateForward the createForward property
func (m *MessageItemRequestBuilder) CreateForward()(*i6ad07a59d46ab03d44f649d85e219feb09f1b9de5800cb7b9590ff92347ff356.CreateForwardRequestBuilder) {
    return i6ad07a59d46ab03d44f649d85e219feb09f1b9de5800cb7b9590ff92347ff356.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the collection of messages in the mailFolder.
func (m *MessageItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the collection of messages in the mailFolder.
func (m *MessageItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property messages in users
func (m *MessageItemRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property messages in users
func (m *MessageItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, requestConfiguration *MessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateReply the createReply property
func (m *MessageItemRequestBuilder) CreateReply()(*i5de64d7ba319e004e4fad5389782964371cfeaa2956093290434d9564c85bcab.CreateReplyRequestBuilder) {
    return i5de64d7ba319e004e4fad5389782964371cfeaa2956093290434d9564c85bcab.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateReplyAll the createReplyAll property
func (m *MessageItemRequestBuilder) CreateReplyAll()(*ib514357c2438d9d84bb6a90ddd348310bf097fc75d97905b8420f5cc3de23bba.CreateReplyAllRequestBuilder) {
    return ib514357c2438d9d84bb6a90ddd348310bf097fc75d97905b8420f5cc3de23bba.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property messages for users
func (m *MessageItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property messages for users
func (m *MessageItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *MessageItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Extensions the extensions property
func (m *MessageItemRequestBuilder) Extensions()(*id0b54dcecaa8542b3fa53a87c14da851c997cae296f64c5235715ba05fca4def.ExtensionsRequestBuilder) {
    return id0b54dcecaa8542b3fa53a87c14da851c997cae296f64c5235715ba05fca4def.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.messages.item.extensions.item collection
func (m *MessageItemRequestBuilder) ExtensionsById(id string)(*iccad5f03bb49efc1650d1fe0e050ffa514f5fa92e483e5f27b73aa5dfbf0c755.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return iccad5f03bb49efc1650d1fe0e050ffa514f5fa92e483e5f27b73aa5dfbf0c755.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *MessageItemRequestBuilder) Forward()(*ib4ec94e919a89eabb07aab3948adafcf0ad3ab5c11f2849ce02c0bedfac417db.ForwardRequestBuilder) {
    return ib4ec94e919a89eabb07aab3948adafcf0ad3ab5c11f2849ce02c0bedfac417db.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of messages in the mailFolder.
func (m *MessageItemRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the collection of messages in the mailFolder.
func (m *MessageItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *MessageItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMessageFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable), nil
}
// Move the move property
func (m *MessageItemRequestBuilder) Move()(*i062a3ff4d1d6529c88f923d81d7a2d73cc91a7aee2ce785d6ea82b627d35e210.MoveRequestBuilder) {
    return i062a3ff4d1d6529c88f923d81d7a2d73cc91a7aee2ce785d6ea82b627d35e210.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *MessageItemRequestBuilder) MultiValueExtendedProperties()(*iee1de9a57b157c37bef3dd47f8c4371431b421428ba049687d93529e0f49bbe3.MultiValueExtendedPropertiesRequestBuilder) {
    return iee1de9a57b157c37bef3dd47f8c4371431b421428ba049687d93529e0f49bbe3.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0d29fac26fc9096d9957be27133c2cb630d1d9aff452d0071baddca02c01fcf5.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i0d29fac26fc9096d9957be27133c2cb630d1d9aff452d0071baddca02c01fcf5.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in users
func (m *MessageItemRequestBuilder) Patch(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property messages in users
func (m *MessageItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, requestConfiguration *MessageItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Reply the reply property
func (m *MessageItemRequestBuilder) Reply()(*ia30921aeaaa3c5a09c4d881ec3a58b3c88454249f15e78743583738180fd9b67.ReplyRequestBuilder) {
    return ia30921aeaaa3c5a09c4d881ec3a58b3c88454249f15e78743583738180fd9b67.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReplyAll the replyAll property
func (m *MessageItemRequestBuilder) ReplyAll()(*i19224a4e273490c542daba438c5ae16a7a3f0e0c07160c68599fe9e49e6d6db5.ReplyAllRequestBuilder) {
    return i19224a4e273490c542daba438c5ae16a7a3f0e0c07160c68599fe9e49e6d6db5.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Send the send property
func (m *MessageItemRequestBuilder) Send()(*ib4f36fecd48150daa2f2c8cf15e035efd83748164c0c3d487b3ed0f3b758f30c.SendRequestBuilder) {
    return ib4f36fecd48150daa2f2c8cf15e035efd83748164c0c3d487b3ed0f3b758f30c.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *MessageItemRequestBuilder) SingleValueExtendedProperties()(*i989a6eef61b5f2d97064f6b03429c220fba6a058beb4c94807f0f085b89b103a.SingleValueExtendedPropertiesRequestBuilder) {
    return i989a6eef61b5f2d97064f6b03429c220fba6a058beb4c94807f0f085b89b103a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7406385c3a11e571b5a4a129b5889d6096ac6a18d6213a54ebcf5e3200bdec80.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i7406385c3a11e571b5a4a129b5889d6096ac6a18d6213a54ebcf5e3200bdec80.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
