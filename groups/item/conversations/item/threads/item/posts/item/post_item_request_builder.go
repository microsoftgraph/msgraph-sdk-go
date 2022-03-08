package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i524bf6311b24b6f141d885728d606289b8ee24cbd67c3599ccb2ec6084cb17bd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/reply"
    i78e7a8cdddabb1f91bbd12d096fb8e9c132c2ec07d32fd988ce7552cfd89352b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/forward"
    i9884a55ef28596c5da82f51922aa27afafadf6bc4977af30511de7bda6744049 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/singlevalueextendedproperties"
    ia669f29093557a86be3b54973c4aba4e78b09891dc40e9d967d02b3938cbb08d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto"
    ie5a656c04186903bcf4da8c7768bc7acdefb45b9b92af375c844ab6dde924c3b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/multivalueextendedproperties"
    if00c8974308ec134f8b54a4b85c9b080451e578b97ae45f3a96f2983e98b235b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/attachments"
    if33d4aff747bbefe22c9d4c0581a32241f79f8c00b65004202a691770673b2d8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/extensions"
    i81cf5d0e294523ab35bef33098c9a475ca60e40c2f945ab2e768942447194086 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/multivalueextendedproperties/item"
    i85387a3ef1c4b59c2db207f407864b3a4fca529975d344ab83cd21fd615de2e6 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/singlevalueextendedproperties/item"
    i98e0b63002f7d7bb02acdac8815a30c2058502950f307d6652a0b2d766b7f02b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/attachments/item"
    i9deed3605dbdfd0f6b5413628ad1be773f48a2ee7f6266b85266a4248af81477 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/extensions/item"
)

// PostItemRequestBuilder provides operations to manage the posts property of the microsoft.graph.conversationThread entity.
type PostItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PostItemRequestBuilderDeleteOptions options for Delete
type PostItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PostItemRequestBuilderGetOptions options for Get
type PostItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PostItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PostItemRequestBuilderGetQueryParameters read-only. Nullable.
type PostItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PostItemRequestBuilderPatchOptions options for Patch
type PostItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Postable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PostItemRequestBuilder) Attachments()(*if00c8974308ec134f8b54a4b85c9b080451e578b97ae45f3a96f2983e98b235b.AttachmentsRequestBuilder) {
    return if00c8974308ec134f8b54a4b85c9b080451e578b97ae45f3a96f2983e98b235b.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.conversations.item.threads.item.posts.item.attachments.item collection
func (m *PostItemRequestBuilder) AttachmentsById(id string)(*i98e0b63002f7d7bb02acdac8815a30c2058502950f307d6652a0b2d766b7f02b.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i98e0b63002f7d7bb02acdac8815a30c2058502950f307d6652a0b2d766b7f02b.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPostItemRequestBuilderInternal instantiates a new PostItemRequestBuilder and sets the default values.
func NewPostItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PostItemRequestBuilder) {
    m := &PostItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/conversations/{conversation_id}/threads/{conversationThread_id}/posts/{post_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPostItemRequestBuilder instantiates a new PostItemRequestBuilder and sets the default values.
func NewPostItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PostItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPostItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property posts for groups
func (m *PostItemRequestBuilder) CreateDeleteRequestInformation(options *PostItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation read-only. Nullable.
func (m *PostItemRequestBuilder) CreateGetRequestInformation(options *PostItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property posts in groups
func (m *PostItemRequestBuilder) CreatePatchRequestInformation(options *PostItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property posts for groups
func (m *PostItemRequestBuilder) Delete(options *PostItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *PostItemRequestBuilder) Extensions()(*if33d4aff747bbefe22c9d4c0581a32241f79f8c00b65004202a691770673b2d8.ExtensionsRequestBuilder) {
    return if33d4aff747bbefe22c9d4c0581a32241f79f8c00b65004202a691770673b2d8.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.conversations.item.threads.item.posts.item.extensions.item collection
func (m *PostItemRequestBuilder) ExtensionsById(id string)(*i9deed3605dbdfd0f6b5413628ad1be773f48a2ee7f6266b85266a4248af81477.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i9deed3605dbdfd0f6b5413628ad1be773f48a2ee7f6266b85266a4248af81477.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PostItemRequestBuilder) Forward()(*i78e7a8cdddabb1f91bbd12d096fb8e9c132c2ec07d32fd988ce7552cfd89352b.ForwardRequestBuilder) {
    return i78e7a8cdddabb1f91bbd12d096fb8e9c132c2ec07d32fd988ce7552cfd89352b.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get read-only. Nullable.
func (m *PostItemRequestBuilder) Get(options *PostItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Postable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreatePostFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Postable), nil
}
func (m *PostItemRequestBuilder) InReplyTo()(*ia669f29093557a86be3b54973c4aba4e78b09891dc40e9d967d02b3938cbb08d.InReplyToRequestBuilder) {
    return ia669f29093557a86be3b54973c4aba4e78b09891dc40e9d967d02b3938cbb08d.NewInReplyToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PostItemRequestBuilder) MultiValueExtendedProperties()(*ie5a656c04186903bcf4da8c7768bc7acdefb45b9b92af375c844ab6dde924c3b.MultiValueExtendedPropertiesRequestBuilder) {
    return ie5a656c04186903bcf4da8c7768bc7acdefb45b9b92af375c844ab6dde924c3b.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.conversations.item.threads.item.posts.item.multiValueExtendedProperties.item collection
func (m *PostItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i81cf5d0e294523ab35bef33098c9a475ca60e40c2f945ab2e768942447194086.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i81cf5d0e294523ab35bef33098c9a475ca60e40c2f945ab2e768942447194086.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property posts in groups
func (m *PostItemRequestBuilder) Patch(options *PostItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *PostItemRequestBuilder) Reply()(*i524bf6311b24b6f141d885728d606289b8ee24cbd67c3599ccb2ec6084cb17bd.ReplyRequestBuilder) {
    return i524bf6311b24b6f141d885728d606289b8ee24cbd67c3599ccb2ec6084cb17bd.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PostItemRequestBuilder) SingleValueExtendedProperties()(*i9884a55ef28596c5da82f51922aa27afafadf6bc4977af30511de7bda6744049.SingleValueExtendedPropertiesRequestBuilder) {
    return i9884a55ef28596c5da82f51922aa27afafadf6bc4977af30511de7bda6744049.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.conversations.item.threads.item.posts.item.singleValueExtendedProperties.item collection
func (m *PostItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i85387a3ef1c4b59c2db207f407864b3a4fca529975d344ab83cd21fd615de2e6.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i85387a3ef1c4b59c2db207f407864b3a4fca529975d344ab83cd21fd615de2e6.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
