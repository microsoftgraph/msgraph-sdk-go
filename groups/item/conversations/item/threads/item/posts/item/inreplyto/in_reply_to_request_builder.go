package inreplyto

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i0b7197b35866560bb75d5d0f4efa896921a6b20436b029349143d8acf9ad3267 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/extensions"
    i5a96cf81c3b745fcb623e502cb75bbd77276c4b7ec8789db5e8e4f2a6f262330 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/attachments"
    i92495245446099b92ab7c7f0065f500a4d14975a7d0b9c96573f137956ca320b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/multivalueextendedproperties"
    i9d94bd11d6d6bd3a2e1f8fa9a4b9aadfb3092f358c65db04ee394627108d0ce5 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/forward"
    ib88622d9430ce7882a6ecfd828e52de9a25082dfb3670f150a4e6d6e20463b04 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/singlevalueextendedproperties"
    if19c96c1013ba36d2ec558a175696439aec92252a53d9eabf93ad72118c4a530 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/reply"
    i0c96d2db214e9fe53246ea7d2b3be413b98c5662576cab468ddd391583915a33 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/singlevalueextendedproperties/item"
    i3e8cb7e4bf2b4a512a1cbb6155635c676f8c2f106acb67a77867e8fe307c6e2c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/attachments/item"
    ic430108688a3dd1740084b1d6f5636f80554b3bca84a2ff93f74d8950f0ff2fd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/multivalueextendedproperties/item"
    ifbfba05084a93efb147200ec38527fedfa451327d8d0168fc33f1244aa12772f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/extensions/item"
)

// InReplyToRequestBuilder provides operations to manage the inReplyTo property of the microsoft.graph.post entity.
type InReplyToRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// InReplyToRequestBuilderDeleteOptions options for Delete
type InReplyToRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// InReplyToRequestBuilderGetOptions options for Get
type InReplyToRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *InReplyToRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// InReplyToRequestBuilderGetQueryParameters read-only. Supports $expand.
type InReplyToRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// InReplyToRequestBuilderPatchOptions options for Patch
type InReplyToRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Postable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *InReplyToRequestBuilder) Attachments()(*i5a96cf81c3b745fcb623e502cb75bbd77276c4b7ec8789db5e8e4f2a6f262330.AttachmentsRequestBuilder) {
    return i5a96cf81c3b745fcb623e502cb75bbd77276c4b7ec8789db5e8e4f2a6f262330.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.conversations.item.threads.item.posts.item.inReplyTo.attachments.item collection
func (m *InReplyToRequestBuilder) AttachmentsById(id string)(*i3e8cb7e4bf2b4a512a1cbb6155635c676f8c2f106acb67a77867e8fe307c6e2c.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i3e8cb7e4bf2b4a512a1cbb6155635c676f8c2f106acb67a77867e8fe307c6e2c.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewInReplyToRequestBuilderInternal instantiates a new InReplyToRequestBuilder and sets the default values.
func NewInReplyToRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InReplyToRequestBuilder) {
    m := &InReplyToRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/conversations/{conversation_id}/threads/{conversationThread_id}/posts/{post_id}/inReplyTo{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInReplyToRequestBuilder instantiates a new InReplyToRequestBuilder and sets the default values.
func NewInReplyToRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InReplyToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInReplyToRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property inReplyTo for groups
func (m *InReplyToRequestBuilder) CreateDeleteRequestInformation(options *InReplyToRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation read-only. Supports $expand.
func (m *InReplyToRequestBuilder) CreateGetRequestInformation(options *InReplyToRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property inReplyTo in groups
func (m *InReplyToRequestBuilder) CreatePatchRequestInformation(options *InReplyToRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property inReplyTo for groups
func (m *InReplyToRequestBuilder) Delete(options *InReplyToRequestBuilderDeleteOptions)(error) {
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
func (m *InReplyToRequestBuilder) Extensions()(*i0b7197b35866560bb75d5d0f4efa896921a6b20436b029349143d8acf9ad3267.ExtensionsRequestBuilder) {
    return i0b7197b35866560bb75d5d0f4efa896921a6b20436b029349143d8acf9ad3267.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.conversations.item.threads.item.posts.item.inReplyTo.extensions.item collection
func (m *InReplyToRequestBuilder) ExtensionsById(id string)(*ifbfba05084a93efb147200ec38527fedfa451327d8d0168fc33f1244aa12772f.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ifbfba05084a93efb147200ec38527fedfa451327d8d0168fc33f1244aa12772f.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InReplyToRequestBuilder) Forward()(*i9d94bd11d6d6bd3a2e1f8fa9a4b9aadfb3092f358c65db04ee394627108d0ce5.ForwardRequestBuilder) {
    return i9d94bd11d6d6bd3a2e1f8fa9a4b9aadfb3092f358c65db04ee394627108d0ce5.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get read-only. Supports $expand.
func (m *InReplyToRequestBuilder) Get(options *InReplyToRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Postable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreatePostFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Postable), nil
}
func (m *InReplyToRequestBuilder) MultiValueExtendedProperties()(*i92495245446099b92ab7c7f0065f500a4d14975a7d0b9c96573f137956ca320b.MultiValueExtendedPropertiesRequestBuilder) {
    return i92495245446099b92ab7c7f0065f500a4d14975a7d0b9c96573f137956ca320b.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.conversations.item.threads.item.posts.item.inReplyTo.multiValueExtendedProperties.item collection
func (m *InReplyToRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ic430108688a3dd1740084b1d6f5636f80554b3bca84a2ff93f74d8950f0ff2fd.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ic430108688a3dd1740084b1d6f5636f80554b3bca84a2ff93f74d8950f0ff2fd.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property inReplyTo in groups
func (m *InReplyToRequestBuilder) Patch(options *InReplyToRequestBuilderPatchOptions)(error) {
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
func (m *InReplyToRequestBuilder) Reply()(*if19c96c1013ba36d2ec558a175696439aec92252a53d9eabf93ad72118c4a530.ReplyRequestBuilder) {
    return if19c96c1013ba36d2ec558a175696439aec92252a53d9eabf93ad72118c4a530.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InReplyToRequestBuilder) SingleValueExtendedProperties()(*ib88622d9430ce7882a6ecfd828e52de9a25082dfb3670f150a4e6d6e20463b04.SingleValueExtendedPropertiesRequestBuilder) {
    return ib88622d9430ce7882a6ecfd828e52de9a25082dfb3670f150a4e6d6e20463b04.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.conversations.item.threads.item.posts.item.inReplyTo.singleValueExtendedProperties.item collection
func (m *InReplyToRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i0c96d2db214e9fe53246ea7d2b3be413b98c5662576cab468ddd391583915a33.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i0c96d2db214e9fe53246ea7d2b3be413b98c5662576cab468ddd391583915a33.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
