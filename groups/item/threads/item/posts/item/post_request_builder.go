package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1d5e0e0253764c43a7170c4dae6d6dada63a8577871ccfb6dbfa89ffcbb240cd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/multivalueextendedproperties"
    i5134eaa0bf73502180126762f8dba15d49dd36677713ba778154e0a60592a5bd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/forward"
    i68a4055bc841e0456a2e6639a547657de33eb9b48f925db253682599604a51e3 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/singlevalueextendedproperties"
    i75174cbcb1b898c11858f5f96b2d8847c330a38149bcd9be83e30ad77ed43c22 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/attachments"
    i87cc194ce5c8b6a9311a66fb5a22e3ba337812de9de5c80802d09bdb94b863aa "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/reply"
    ic5c310ab811a9efd80ec215346a10a39c86ebea5d6fadba8de5ac53f120fbdc1 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/extensions"
    ie143dd9268acc57eb633fcf1a20193d5695c0268106b5f598eb2c2f1e7a8a611 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/inreplyto"
    i27611e4379a9836b89c70ec5ef6aa8f3e16e6cc0deccd831272023f099a54b9a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/multivalueextendedproperties/item"
    i8ec098680e04f7df4b9ae597c1037f94b0a4442db8e1280e22d5710ee0602d81 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/attachments/item"
    ie25007fdca9d41586081045acb368264715f3e0d2b48bebac108a93b0d11781a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/extensions/item"
    if0f66d4e50c563fdb7e3c99373c63bc365617d543a82f7c19e2cfc84e3dcb3a8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/threads/item/posts/item/singlevalueextendedproperties/item"
)

// Builds and executes requests for operations under \groups\{group-id}\threads\{conversationThread-id}\posts\{post-id}
type PostRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type PostRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type PostRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PostRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Read-only. Nullable.
type PostRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type PostRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Post;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PostRequestBuilder) Attachments()(*i75174cbcb1b898c11858f5f96b2d8847c330a38149bcd9be83e30ad77ed43c22.AttachmentsRequestBuilder) {
    return i75174cbcb1b898c11858f5f96b2d8847c330a38149bcd9be83e30ad77ed43c22.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.threads.item.posts.item.attachments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PostRequestBuilder) AttachmentsById(id string)(*i8ec098680e04f7df4b9ae597c1037f94b0a4442db8e1280e22d5710ee0602d81.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i8ec098680e04f7df4b9ae597c1037f94b0a4442db8e1280e22d5710ee0602d81.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new PostRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPostRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PostRequestBuilder) {
    m := &PostRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/threads/{conversationThread_id}/posts/{post_id}{?select,expand}";
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
//  - options : Options for the request
func (m *PostRequestBuilder) CreateDeleteRequestInformation(options *PostRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *PostRequestBuilder) CreateGetRequestInformation(options *PostRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(options.Q)
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
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *PostRequestBuilder) CreatePatchRequestInformation(options *PostRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *PostRequestBuilder) Delete(options *PostRequestBuilderDeleteOptions)(error) {
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
func (m *PostRequestBuilder) Extensions()(*ic5c310ab811a9efd80ec215346a10a39c86ebea5d6fadba8de5ac53f120fbdc1.ExtensionsRequestBuilder) {
    return ic5c310ab811a9efd80ec215346a10a39c86ebea5d6fadba8de5ac53f120fbdc1.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.threads.item.posts.item.extensions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PostRequestBuilder) ExtensionsById(id string)(*ie25007fdca9d41586081045acb368264715f3e0d2b48bebac108a93b0d11781a.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ie25007fdca9d41586081045acb368264715f3e0d2b48bebac108a93b0d11781a.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PostRequestBuilder) Forward()(*i5134eaa0bf73502180126762f8dba15d49dd36677713ba778154e0a60592a5bd.ForwardRequestBuilder) {
    return i5134eaa0bf73502180126762f8dba15d49dd36677713ba778154e0a60592a5bd.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *PostRequestBuilder) Get(options *PostRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Post, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPost() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Post), nil
}
func (m *PostRequestBuilder) InReplyTo()(*ie143dd9268acc57eb633fcf1a20193d5695c0268106b5f598eb2c2f1e7a8a611.InReplyToRequestBuilder) {
    return ie143dd9268acc57eb633fcf1a20193d5695c0268106b5f598eb2c2f1e7a8a611.NewInReplyToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PostRequestBuilder) MultiValueExtendedProperties()(*i1d5e0e0253764c43a7170c4dae6d6dada63a8577871ccfb6dbfa89ffcbb240cd.MultiValueExtendedPropertiesRequestBuilder) {
    return i1d5e0e0253764c43a7170c4dae6d6dada63a8577871ccfb6dbfa89ffcbb240cd.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.threads.item.posts.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PostRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i27611e4379a9836b89c70ec5ef6aa8f3e16e6cc0deccd831272023f099a54b9a.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i27611e4379a9836b89c70ec5ef6aa8f3e16e6cc0deccd831272023f099a54b9a.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *PostRequestBuilder) Patch(options *PostRequestBuilderPatchOptions)(error) {
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
func (m *PostRequestBuilder) Reply()(*i87cc194ce5c8b6a9311a66fb5a22e3ba337812de9de5c80802d09bdb94b863aa.ReplyRequestBuilder) {
    return i87cc194ce5c8b6a9311a66fb5a22e3ba337812de9de5c80802d09bdb94b863aa.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PostRequestBuilder) SingleValueExtendedProperties()(*i68a4055bc841e0456a2e6639a547657de33eb9b48f925db253682599604a51e3.SingleValueExtendedPropertiesRequestBuilder) {
    return i68a4055bc841e0456a2e6639a547657de33eb9b48f925db253682599604a51e3.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.threads.item.posts.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PostRequestBuilder) SingleValueExtendedPropertiesById(id string)(*if0f66d4e50c563fdb7e3c99373c63bc365617d543a82f7c19e2cfc84e3dcb3a8.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return if0f66d4e50c563fdb7e3c99373c63bc365617d543a82f7c19e2cfc84e3dcb3a8.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
