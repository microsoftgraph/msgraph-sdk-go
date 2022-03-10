package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i01ec3c483e02043ffcf9719047c0715f653bcff2cc50ed31df82e09e8920121d "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/childfolders/item/move"
    i0b868ecc056c79c6636dd51ef482a36c5c894daf4f588668d67f7b22da8ce77d "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/childfolders/item/messages"
    i1743205569c2a7c4d5bb6d44ff0107a2e7e60d549e2b33fc00b68911f1572aaa "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/childfolders/item/copy"
    i70113464c82bb92ef63f33259096dbfe2f13fe756f26b3075549b1294a1d8d82 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/childfolders/item/multivalueextendedproperties"
    i72dc3d7edde3c8160fe7cadf4cff9ad965fa3d951f399fadd4bd4e45b47f2e90 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/childfolders/item/messagerules"
    iabfcbaada8ea53d26b0a397e0f48fb8df66733cda6d760b402ebe678e743750c "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/childfolders/item/singlevalueextendedproperties"
    i076172bf8e9645aa1ed174294d545f0bcb13978581051cd8ef7e93371f949076 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/childfolders/item/singlevalueextendedproperties/item"
    i4bad71fe4da4717a6d793b680796b60364f5758c65ea16ba378a07935aaed713 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/childfolders/item/multivalueextendedproperties/item"
    i8d9aad822d253c42b11a52bfea737460988d9b2b3512881497cb4b0bdf408e29 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/childfolders/item/messages/item"
    ieefd001b236c83104f4cb9fe64f0f159586d25f0328105a728c316df9621cc75 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item/childfolders/item/messagerules/item"
)

// MailFolderItemRequestBuilder provides operations to manage the childFolders property of the microsoft.graph.mailFolder entity.
type MailFolderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MailFolderItemRequestBuilderDeleteOptions options for Delete
type MailFolderItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MailFolderItemRequestBuilderGetOptions options for Get
type MailFolderItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MailFolderItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MailFolderItemRequestBuilderGetQueryParameters the collection of child folders in the mailFolder.
type MailFolderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// MailFolderItemRequestBuilderPatchOptions options for Patch
type MailFolderItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailFolderable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewMailFolderItemRequestBuilderInternal instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewMailFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MailFolderItemRequestBuilder) {
    m := &MailFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mailFolders/{mailFolder_id}/childFolders/{mailFolder_id1}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMailFolderItemRequestBuilder instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewMailFolderItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MailFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMailFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MailFolderItemRequestBuilder) Copy()(*i1743205569c2a7c4d5bb6d44ff0107a2e7e60d549e2b33fc00b68911f1572aaa.CopyRequestBuilder) {
    return i1743205569c2a7c4d5bb6d44ff0107a2e7e60d549e2b33fc00b68911f1572aaa.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property childFolders for me
func (m *MailFolderItemRequestBuilder) CreateDeleteRequestInformation(options *MailFolderItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of child folders in the mailFolder.
func (m *MailFolderItemRequestBuilder) CreateGetRequestInformation(options *MailFolderItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property childFolders in me
func (m *MailFolderItemRequestBuilder) CreatePatchRequestInformation(options *MailFolderItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property childFolders for me
func (m *MailFolderItemRequestBuilder) Delete(options *MailFolderItemRequestBuilderDeleteOptions)(error) {
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
// Get the collection of child folders in the mailFolder.
func (m *MailFolderItemRequestBuilder) Get(options *MailFolderItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailFolderable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateMailFolderFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailFolderable), nil
}
func (m *MailFolderItemRequestBuilder) MessageRules()(*i72dc3d7edde3c8160fe7cadf4cff9ad965fa3d951f399fadd4bd4e45b47f2e90.MessageRulesRequestBuilder) {
    return i72dc3d7edde3c8160fe7cadf4cff9ad965fa3d951f399fadd4bd4e45b47f2e90.NewMessageRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessageRulesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.mailFolders.item.childFolders.item.messageRules.item collection
func (m *MailFolderItemRequestBuilder) MessageRulesById(id string)(*ieefd001b236c83104f4cb9fe64f0f159586d25f0328105a728c316df9621cc75.MessageRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["messageRule_id"] = id
    }
    return ieefd001b236c83104f4cb9fe64f0f159586d25f0328105a728c316df9621cc75.NewMessageRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderItemRequestBuilder) Messages()(*i0b868ecc056c79c6636dd51ef482a36c5c894daf4f588668d67f7b22da8ce77d.MessagesRequestBuilder) {
    return i0b868ecc056c79c6636dd51ef482a36c5c894daf4f588668d67f7b22da8ce77d.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.mailFolders.item.childFolders.item.messages.item collection
func (m *MailFolderItemRequestBuilder) MessagesById(id string)(*i8d9aad822d253c42b11a52bfea737460988d9b2b3512881497cb4b0bdf408e29.MessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message_id"] = id
    }
    return i8d9aad822d253c42b11a52bfea737460988d9b2b3512881497cb4b0bdf408e29.NewMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderItemRequestBuilder) Move()(*i01ec3c483e02043ffcf9719047c0715f653bcff2cc50ed31df82e09e8920121d.MoveRequestBuilder) {
    return i01ec3c483e02043ffcf9719047c0715f653bcff2cc50ed31df82e09e8920121d.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderItemRequestBuilder) MultiValueExtendedProperties()(*i70113464c82bb92ef63f33259096dbfe2f13fe756f26b3075549b1294a1d8d82.MultiValueExtendedPropertiesRequestBuilder) {
    return i70113464c82bb92ef63f33259096dbfe2f13fe756f26b3075549b1294a1d8d82.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.mailFolders.item.childFolders.item.multiValueExtendedProperties.item collection
func (m *MailFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i4bad71fe4da4717a6d793b680796b60364f5758c65ea16ba378a07935aaed713.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i4bad71fe4da4717a6d793b680796b60364f5758c65ea16ba378a07935aaed713.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property childFolders in me
func (m *MailFolderItemRequestBuilder) Patch(options *MailFolderItemRequestBuilderPatchOptions)(error) {
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
func (m *MailFolderItemRequestBuilder) SingleValueExtendedProperties()(*iabfcbaada8ea53d26b0a397e0f48fb8df66733cda6d760b402ebe678e743750c.SingleValueExtendedPropertiesRequestBuilder) {
    return iabfcbaada8ea53d26b0a397e0f48fb8df66733cda6d760b402ebe678e743750c.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.mailFolders.item.childFolders.item.singleValueExtendedProperties.item collection
func (m *MailFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i076172bf8e9645aa1ed174294d545f0bcb13978581051cd8ef7e93371f949076.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i076172bf8e9645aa1ed174294d545f0bcb13978581051cd8ef7e93371f949076.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
