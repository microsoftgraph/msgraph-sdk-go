package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i28a830c1e815101c0b85504c7b909848fd0dc929f2fa98ecbf66ae5908cc0e7e "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/copy"
    i57ff9c1da28e32d1438213a00ffd1d855134d619642894e7fd87e4532027a864 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/multivalueextendedproperties"
    i7f0870b87ccb71db8ecb5bfea28fb9e54778c808e443d1a99c4580eea1bd3f52 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/move"
    i8f7120f978f78e56c6efaa74e7732bcb02d100ea8df97bff99e5f8e11b74afc0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages"
    ic134c9b00254f702eada75fd8ad921541915d74fcca08820a255eae64692f139 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/singlevalueextendedproperties"
    id31d137ab3a6ab25d27f953a9a06af2a8a3fec04189a4cb157e4d5cb6cdaebf4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messagerules"
    if9a47801299e0be05f05c57b8b33e6c23c8bf48add445922ff4931f33ac0a26e "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders"
    i1285f2feb19f3fd1b1b5a3d0dc5440e58dafd55c4f69daede162ea2180f7adf5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item"
    i335a0ecd907517c064bd273cae9902fe9447a3a8ec23fe4de8b215b5356fc26c "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/multivalueextendedproperties/item"
    i9c44f0ee1a0a3f651db0ddfbaf98ec44de5c06f226dce2ae35c6c384f3a2bad4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/item"
    ic153e3fa3efccd7712fa90e3c5117dc286ddf11d28ccac6a7570f274e8f78032 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/singlevalueextendedproperties/item"
    if56e053996654e0ad839a540d45995ed16925363ccece02b277bf8d9fe87dd53 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messagerules/item"
)

// MailFolderItemRequestBuilder provides operations to manage the mailFolders property of the microsoft.graph.user entity.
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
// MailFolderItemRequestBuilderGetQueryParameters the user's mail folders. Read-only. Nullable.
type MailFolderItemRequestBuilderGetQueryParameters struct {
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
func (m *MailFolderItemRequestBuilder) ChildFolders()(*if9a47801299e0be05f05c57b8b33e6c23c8bf48add445922ff4931f33ac0a26e.ChildFoldersRequestBuilder) {
    return if9a47801299e0be05f05c57b8b33e6c23c8bf48add445922ff4931f33ac0a26e.NewChildFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildFoldersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.childFolders.item collection
func (m *MailFolderItemRequestBuilder) ChildFoldersById(id string)(*i1285f2feb19f3fd1b1b5a3d0dc5440e58dafd55c4f69daede162ea2180f7adf5.MailFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder_id1"] = id
    }
    return i1285f2feb19f3fd1b1b5a3d0dc5440e58dafd55c4f69daede162ea2180f7adf5.NewMailFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMailFolderItemRequestBuilderInternal instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewMailFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MailFolderItemRequestBuilder) {
    m := &MailFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}{?select}";
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
func (m *MailFolderItemRequestBuilder) Copy()(*i28a830c1e815101c0b85504c7b909848fd0dc929f2fa98ecbf66ae5908cc0e7e.CopyRequestBuilder) {
    return i28a830c1e815101c0b85504c7b909848fd0dc929f2fa98ecbf66ae5908cc0e7e.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property mailFolders for users
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
// CreateGetRequestInformation the user's mail folders. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property mailFolders in users
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
// Delete delete navigation property mailFolders for users
func (m *MailFolderItemRequestBuilder) Delete(options *MailFolderItemRequestBuilderDeleteOptions)(error) {
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
// Get the user's mail folders. Read-only. Nullable.
func (m *MailFolderItemRequestBuilder) Get(options *MailFolderItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailFolderable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateMailFolderFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailFolderable), nil
}
func (m *MailFolderItemRequestBuilder) MessageRules()(*id31d137ab3a6ab25d27f953a9a06af2a8a3fec04189a4cb157e4d5cb6cdaebf4.MessageRulesRequestBuilder) {
    return id31d137ab3a6ab25d27f953a9a06af2a8a3fec04189a4cb157e4d5cb6cdaebf4.NewMessageRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessageRulesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.messageRules.item collection
func (m *MailFolderItemRequestBuilder) MessageRulesById(id string)(*if56e053996654e0ad839a540d45995ed16925363ccece02b277bf8d9fe87dd53.MessageRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["messageRule_id"] = id
    }
    return if56e053996654e0ad839a540d45995ed16925363ccece02b277bf8d9fe87dd53.NewMessageRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderItemRequestBuilder) Messages()(*i8f7120f978f78e56c6efaa74e7732bcb02d100ea8df97bff99e5f8e11b74afc0.MessagesRequestBuilder) {
    return i8f7120f978f78e56c6efaa74e7732bcb02d100ea8df97bff99e5f8e11b74afc0.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.messages.item collection
func (m *MailFolderItemRequestBuilder) MessagesById(id string)(*i9c44f0ee1a0a3f651db0ddfbaf98ec44de5c06f226dce2ae35c6c384f3a2bad4.MessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message_id"] = id
    }
    return i9c44f0ee1a0a3f651db0ddfbaf98ec44de5c06f226dce2ae35c6c384f3a2bad4.NewMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderItemRequestBuilder) Move()(*i7f0870b87ccb71db8ecb5bfea28fb9e54778c808e443d1a99c4580eea1bd3f52.MoveRequestBuilder) {
    return i7f0870b87ccb71db8ecb5bfea28fb9e54778c808e443d1a99c4580eea1bd3f52.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderItemRequestBuilder) MultiValueExtendedProperties()(*i57ff9c1da28e32d1438213a00ffd1d855134d619642894e7fd87e4532027a864.MultiValueExtendedPropertiesRequestBuilder) {
    return i57ff9c1da28e32d1438213a00ffd1d855134d619642894e7fd87e4532027a864.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.multiValueExtendedProperties.item collection
func (m *MailFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i335a0ecd907517c064bd273cae9902fe9447a3a8ec23fe4de8b215b5356fc26c.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i335a0ecd907517c064bd273cae9902fe9447a3a8ec23fe4de8b215b5356fc26c.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property mailFolders in users
func (m *MailFolderItemRequestBuilder) Patch(options *MailFolderItemRequestBuilderPatchOptions)(error) {
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
func (m *MailFolderItemRequestBuilder) SingleValueExtendedProperties()(*ic134c9b00254f702eada75fd8ad921541915d74fcca08820a255eae64692f139.SingleValueExtendedPropertiesRequestBuilder) {
    return ic134c9b00254f702eada75fd8ad921541915d74fcca08820a255eae64692f139.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.singleValueExtendedProperties.item collection
func (m *MailFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ic153e3fa3efccd7712fa90e3c5117dc286ddf11d28ccac6a7570f274e8f78032.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ic153e3fa3efccd7712fa90e3c5117dc286ddf11d28ccac6a7570f274e8f78032.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
