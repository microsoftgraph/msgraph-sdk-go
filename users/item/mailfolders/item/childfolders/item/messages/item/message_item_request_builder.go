package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i078dc43cfacb492d7c14c400159d76effb3ddc82fd0975278340abf380ca4ac0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/createforward"
    i1e2e330cf63776672ee46cb29c0931a99d88b111c325d0eddcba469f590284ef "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/value"
    i283292ee0714f6d62978117a50ae89004640243f72c98fb77b3d3ad5d895724b "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/multivalueextendedproperties"
    i46afd8dc1f68253c5ec58fffe851f31d4606c9896eeca162270731a770163159 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/move"
    i4d24d04db91a658e0a0039257207563b182ae368b5564bc2651941a177fcec1f "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/send"
    i5a8c865ef21d4e26d59de283ac0bf75d8fa86b7ef5610dbf01669ac1a6a6426c "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/singlevalueextendedproperties"
    i601006ca5c8d4a2bd32a308480c958aa565af8b7eef89ca92de125fddcd8df37 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/createreply"
    i8bf596c0443d82bd098088cf313820c77518a4903e0996592bf7aa9bb0fc6204 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/forward"
    ia2b7d74de2d2ae40f811c9f5ca231ae8bcd6383bd423cbdadabf72fbdbfb6e62 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/replyall"
    ia3fdc47a4b68aebb3754bf14fd0396a90c0336bfcb6d210afc93388c25c3f2df "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/calendarsharingmessage"
    ia7dce7bbbad8834f61a47bf8d2d11a325ca6cb60f92e7e712275f7e8c25fc693 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/reply"
    ib2ca11a51f3340d7f385f534ea5827066380a80aa53f7fb73586d468aad2cf61 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/createreplyall"
    id3a837a6f0f6a0a991f73cd274bed4374b20b7ab621b89f0c8b54728f1438971 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/copy"
    id871d310ec2178d7aee89ca6f65864dd837e1389228be08e4040bb3486b367d4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/extensions"
    if72e21ea0481f36a049b51f38179e467041e8723702278110ed61ad1e65dde50 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/attachments"
    i5b5457ef2553e64079400434de96e0e9776d9caa96dc2683cdee41584b29dcf5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/multivalueextendedproperties/item"
    ic067f029b1404c524dee30760b61043b1957ec2a9b5b5f2858424190465ad31a "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/singlevalueextendedproperties/item"
    ie43dfbb21dca01399bcbaa4461736852466b5a2f62ecd895499c91d040c14a01 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/attachments/item"
    ief0529f669908aba2ad97bcc88a9d64281988927f10a7814f2b993f88439be1b "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/extensions/item"
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
func (m *MessageItemRequestBuilder) Attachments()(*if72e21ea0481f36a049b51f38179e467041e8723702278110ed61ad1e65dde50.AttachmentsRequestBuilder) {
    return if72e21ea0481f36a049b51f38179e467041e8723702278110ed61ad1e65dde50.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.childFolders.item.messages.item.attachments.item collection
func (m *MessageItemRequestBuilder) AttachmentsById(id string)(*ie43dfbb21dca01399bcbaa4461736852466b5a2f62ecd895499c91d040c14a01.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ie43dfbb21dca01399bcbaa4461736852466b5a2f62ecd895499c91d040c14a01.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) CalendarSharingMessage()(*ia3fdc47a4b68aebb3754bf14fd0396a90c0336bfcb6d210afc93388c25c3f2df.CalendarSharingMessageRequestBuilder) {
    return ia3fdc47a4b68aebb3754bf14fd0396a90c0336bfcb6d210afc93388c25c3f2df.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageItemRequestBuilder) {
    m := &MessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}/childFolders/{mailFolder_id1}/messages/{message_id}{?select,expand}";
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
func (m *MessageItemRequestBuilder) Content()(*i1e2e330cf63776672ee46cb29c0931a99d88b111c325d0eddcba469f590284ef.ContentRequestBuilder) {
    return i1e2e330cf63776672ee46cb29c0931a99d88b111c325d0eddcba469f590284ef.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Copy()(*id3a837a6f0f6a0a991f73cd274bed4374b20b7ab621b89f0c8b54728f1438971.CopyRequestBuilder) {
    return id3a837a6f0f6a0a991f73cd274bed4374b20b7ab621b89f0c8b54728f1438971.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property messages for users
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
func (m *MessageItemRequestBuilder) CreateForward()(*i078dc43cfacb492d7c14c400159d76effb3ddc82fd0975278340abf380ca4ac0.CreateForwardRequestBuilder) {
    return i078dc43cfacb492d7c14c400159d76effb3ddc82fd0975278340abf380ca4ac0.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// CreatePatchRequestInformation update the navigation property messages in users
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
func (m *MessageItemRequestBuilder) CreateReply()(*i601006ca5c8d4a2bd32a308480c958aa565af8b7eef89ca92de125fddcd8df37.CreateReplyRequestBuilder) {
    return i601006ca5c8d4a2bd32a308480c958aa565af8b7eef89ca92de125fddcd8df37.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) CreateReplyAll()(*ib2ca11a51f3340d7f385f534ea5827066380a80aa53f7fb73586d468aad2cf61.CreateReplyAllRequestBuilder) {
    return ib2ca11a51f3340d7f385f534ea5827066380a80aa53f7fb73586d468aad2cf61.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property messages for users
func (m *MessageItemRequestBuilder) Delete(options *MessageItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *MessageItemRequestBuilder) Extensions()(*id871d310ec2178d7aee89ca6f65864dd837e1389228be08e4040bb3486b367d4.ExtensionsRequestBuilder) {
    return id871d310ec2178d7aee89ca6f65864dd837e1389228be08e4040bb3486b367d4.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.childFolders.item.messages.item.extensions.item collection
func (m *MessageItemRequestBuilder) ExtensionsById(id string)(*ief0529f669908aba2ad97bcc88a9d64281988927f10a7814f2b993f88439be1b.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ief0529f669908aba2ad97bcc88a9d64281988927f10a7814f2b993f88439be1b.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Forward()(*i8bf596c0443d82bd098088cf313820c77518a4903e0996592bf7aa9bb0fc6204.ForwardRequestBuilder) {
    return i8bf596c0443d82bd098088cf313820c77518a4903e0996592bf7aa9bb0fc6204.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of messages in the mailFolder.
func (m *MessageItemRequestBuilder) Get(options *MessageItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Messageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateMessageFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Messageable), nil
}
func (m *MessageItemRequestBuilder) Move()(*i46afd8dc1f68253c5ec58fffe851f31d4606c9896eeca162270731a770163159.MoveRequestBuilder) {
    return i46afd8dc1f68253c5ec58fffe851f31d4606c9896eeca162270731a770163159.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) MultiValueExtendedProperties()(*i283292ee0714f6d62978117a50ae89004640243f72c98fb77b3d3ad5d895724b.MultiValueExtendedPropertiesRequestBuilder) {
    return i283292ee0714f6d62978117a50ae89004640243f72c98fb77b3d3ad5d895724b.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.childFolders.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5b5457ef2553e64079400434de96e0e9776d9caa96dc2683cdee41584b29dcf5.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i5b5457ef2553e64079400434de96e0e9776d9caa96dc2683cdee41584b29dcf5.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in users
func (m *MessageItemRequestBuilder) Patch(options *MessageItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *MessageItemRequestBuilder) Reply()(*ia7dce7bbbad8834f61a47bf8d2d11a325ca6cb60f92e7e712275f7e8c25fc693.ReplyRequestBuilder) {
    return ia7dce7bbbad8834f61a47bf8d2d11a325ca6cb60f92e7e712275f7e8c25fc693.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) ReplyAll()(*ia2b7d74de2d2ae40f811c9f5ca231ae8bcd6383bd423cbdadabf72fbdbfb6e62.ReplyAllRequestBuilder) {
    return ia2b7d74de2d2ae40f811c9f5ca231ae8bcd6383bd423cbdadabf72fbdbfb6e62.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Send()(*i4d24d04db91a658e0a0039257207563b182ae368b5564bc2651941a177fcec1f.SendRequestBuilder) {
    return i4d24d04db91a658e0a0039257207563b182ae368b5564bc2651941a177fcec1f.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) SingleValueExtendedProperties()(*i5a8c865ef21d4e26d59de283ac0bf75d8fa86b7ef5610dbf01669ac1a6a6426c.SingleValueExtendedPropertiesRequestBuilder) {
    return i5a8c865ef21d4e26d59de283ac0bf75d8fa86b7ef5610dbf01669ac1a6a6426c.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.childFolders.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ic067f029b1404c524dee30760b61043b1957ec2a9b5b5f2858424190465ad31a.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ic067f029b1404c524dee30760b61043b1957ec2a9b5b5f2858424190465ad31a.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
