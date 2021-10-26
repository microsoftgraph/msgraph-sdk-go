package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
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

// Builds and executes requests for operations under \groups\{group-id}\conversations\{conversation-id}\threads\{conversationThread-id}\posts\{post-id}
type PostRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Read-only. Nullable.
type PostRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *PostRequestBuilder) Attachments()(*if00c8974308ec134f8b54a4b85c9b080451e578b97ae45f3a96f2983e98b235b.AttachmentsRequestBuilder) {
    return if00c8974308ec134f8b54a4b85c9b080451e578b97ae45f3a96f2983e98b235b.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.conversations.item.threads.item.posts.item.attachments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PostRequestBuilder) AttachmentsById(id string)(*i98e0b63002f7d7bb02acdac8815a30c2058502950f307d6652a0b2d766b7f02b.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i98e0b63002f7d7bb02acdac8815a30c2058502950f307d6652a0b2d766b7f02b.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new PostRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPostRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PostRequestBuilder) {
    m := &PostRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/groups/{group_id}/conversations/{conversation_id}/threads/{conversationThread_id}/posts/{post_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new PostRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPostRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PostRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPostRequestBuilderInternal(urlParams, requestAdapter)
}
// Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *PostRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *PostRequestBuilder) CreateGetRequestInformation(q func (value *PostRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(PostRequestBuilderGetQueryParameters)
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
// Read-only. Nullable.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *PostRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Post, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *PostRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PostRequestBuilder) Extensions()(*if33d4aff747bbefe22c9d4c0581a32241f79f8c00b65004202a691770673b2d8.ExtensionsRequestBuilder) {
    return if33d4aff747bbefe22c9d4c0581a32241f79f8c00b65004202a691770673b2d8.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.conversations.item.threads.item.posts.item.extensions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PostRequestBuilder) ExtensionsById(id string)(*i9deed3605dbdfd0f6b5413628ad1be773f48a2ee7f6266b85266a4248af81477.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i9deed3605dbdfd0f6b5413628ad1be773f48a2ee7f6266b85266a4248af81477.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PostRequestBuilder) Forward()(*i78e7a8cdddabb1f91bbd12d096fb8e9c132c2ec07d32fd988ce7552cfd89352b.ForwardRequestBuilder) {
    return i78e7a8cdddabb1f91bbd12d096fb8e9c132c2ec07d32fd988ce7552cfd89352b.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *PostRequestBuilder) Get(q func (value *PostRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Post, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPost() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Post), nil
}
func (m *PostRequestBuilder) InReplyTo()(*ia669f29093557a86be3b54973c4aba4e78b09891dc40e9d967d02b3938cbb08d.InReplyToRequestBuilder) {
    return ia669f29093557a86be3b54973c4aba4e78b09891dc40e9d967d02b3938cbb08d.NewInReplyToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PostRequestBuilder) MultiValueExtendedProperties()(*ie5a656c04186903bcf4da8c7768bc7acdefb45b9b92af375c844ab6dde924c3b.MultiValueExtendedPropertiesRequestBuilder) {
    return ie5a656c04186903bcf4da8c7768bc7acdefb45b9b92af375c844ab6dde924c3b.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.conversations.item.threads.item.posts.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PostRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i81cf5d0e294523ab35bef33098c9a475ca60e40c2f945ab2e768942447194086.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i81cf5d0e294523ab35bef33098c9a475ca60e40c2f945ab2e768942447194086.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Read-only. Nullable.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *PostRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Post, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PostRequestBuilder) Reply()(*i524bf6311b24b6f141d885728d606289b8ee24cbd67c3599ccb2ec6084cb17bd.ReplyRequestBuilder) {
    return i524bf6311b24b6f141d885728d606289b8ee24cbd67c3599ccb2ec6084cb17bd.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PostRequestBuilder) SingleValueExtendedProperties()(*i9884a55ef28596c5da82f51922aa27afafadf6bc4977af30511de7bda6744049.SingleValueExtendedPropertiesRequestBuilder) {
    return i9884a55ef28596c5da82f51922aa27afafadf6bc4977af30511de7bda6744049.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.conversations.item.threads.item.posts.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PostRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i85387a3ef1c4b59c2db207f407864b3a4fca529975d344ab83cd21fd615de2e6.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i85387a3ef1c4b59c2db207f407864b3a4fca529975d344ab83cd21fd615de2e6.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
