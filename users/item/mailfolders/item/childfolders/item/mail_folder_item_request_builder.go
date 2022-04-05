package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i01ac78197c17c23ecca173b1fb7acf4edf9b474f37a0ebcf52bb27074f4805ba "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages"
    i990fe89bba99adc4df3d6b8e35572e7504bd3f7757624d1b727ea074dd31f631 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/singlevalueextendedproperties"
    ia5a9078703c87169ca0a72f154ecda92ee298f176ca922ee15b15e73bc0d0c1f "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/copy"
    ib9239ea776e711cf86ad2218b84dc423b118eead0e9fdc5b44c2d4e817f7fa9d "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messagerules"
    id7b120c4d5fde20c14fd8826be84f56aab20ee7e9de6eae69a267bb3d2bdb86d "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/multivalueextendedproperties"
    if4e4166d7062e17a8da6fd62eb9b07b412a5a2d29ae9e66fa5e82953c4c30d19 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/move"
    i3df17c93c93c1f05f97aebafc6e566e725950ce62a85f12563acadc9925570d0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/singlevalueextendedproperties/item"
    i941aad2ca9f8e5631ff078dcdd2cf528c6330b158cad3254af5cf0ff24716a7e "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item"
    i943f365e8642a2dff088c70ccb525e63a545eb5681a535f016dc5bf089fb7fa8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/messagerules/item"
    iaae4278d9cf965b61eafe81ab302cfd39ddf559d2a9fd524dbad444a501056dd "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/item/multivalueextendedproperties/item"
)

// MailFolderItemRequestBuilder provides operations to manage the childFolders property of the microsoft.graph.mailFolder entity.
type MailFolderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MailFolderItemRequestBuilderDeleteOptions options for Delete
type MailFolderItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// MailFolderItemRequestBuilderGetOptions options for Get
type MailFolderItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *MailFolderItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
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
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailFolderable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewMailFolderItemRequestBuilderInternal instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewMailFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MailFolderItemRequestBuilder) {
    m := &MailFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}/childFolders/{mailFolder_id1}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMailFolderItemRequestBuilder instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewMailFolderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MailFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMailFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Copy the copy property
func (m *MailFolderItemRequestBuilder) Copy()(*ia5a9078703c87169ca0a72f154ecda92ee298f176ca922ee15b15e73bc0d0c1f.CopyRequestBuilder) {
    return ia5a9078703c87169ca0a72f154ecda92ee298f176ca922ee15b15e73bc0d0c1f.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property childFolders for users
func (m *MailFolderItemRequestBuilder) CreateDeleteRequestInformation(options *MailFolderItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection of child folders in the mailFolder.
func (m *MailFolderItemRequestBuilder) CreateGetRequestInformation(options *MailFolderItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property childFolders in users
func (m *MailFolderItemRequestBuilder) CreatePatchRequestInformation(options *MailFolderItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property childFolders for users
func (m *MailFolderItemRequestBuilder) Delete(options *MailFolderItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of child folders in the mailFolder.
func (m *MailFolderItemRequestBuilder) Get(options *MailFolderItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailFolderable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMailFolderFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailFolderable), nil
}
// MessageRules the messageRules property
func (m *MailFolderItemRequestBuilder) MessageRules()(*ib9239ea776e711cf86ad2218b84dc423b118eead0e9fdc5b44c2d4e817f7fa9d.MessageRulesRequestBuilder) {
    return ib9239ea776e711cf86ad2218b84dc423b118eead0e9fdc5b44c2d4e817f7fa9d.NewMessageRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessageRulesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.childFolders.item.messageRules.item collection
func (m *MailFolderItemRequestBuilder) MessageRulesById(id string)(*i943f365e8642a2dff088c70ccb525e63a545eb5681a535f016dc5bf089fb7fa8.MessageRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["messageRule_id"] = id
    }
    return i943f365e8642a2dff088c70ccb525e63a545eb5681a535f016dc5bf089fb7fa8.NewMessageRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *MailFolderItemRequestBuilder) Messages()(*i01ac78197c17c23ecca173b1fb7acf4edf9b474f37a0ebcf52bb27074f4805ba.MessagesRequestBuilder) {
    return i01ac78197c17c23ecca173b1fb7acf4edf9b474f37a0ebcf52bb27074f4805ba.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.childFolders.item.messages.item collection
func (m *MailFolderItemRequestBuilder) MessagesById(id string)(*i941aad2ca9f8e5631ff078dcdd2cf528c6330b158cad3254af5cf0ff24716a7e.MessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message_id"] = id
    }
    return i941aad2ca9f8e5631ff078dcdd2cf528c6330b158cad3254af5cf0ff24716a7e.NewMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Move the move property
func (m *MailFolderItemRequestBuilder) Move()(*if4e4166d7062e17a8da6fd62eb9b07b412a5a2d29ae9e66fa5e82953c4c30d19.MoveRequestBuilder) {
    return if4e4166d7062e17a8da6fd62eb9b07b412a5a2d29ae9e66fa5e82953c4c30d19.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *MailFolderItemRequestBuilder) MultiValueExtendedProperties()(*id7b120c4d5fde20c14fd8826be84f56aab20ee7e9de6eae69a267bb3d2bdb86d.MultiValueExtendedPropertiesRequestBuilder) {
    return id7b120c4d5fde20c14fd8826be84f56aab20ee7e9de6eae69a267bb3d2bdb86d.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.childFolders.item.multiValueExtendedProperties.item collection
func (m *MailFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iaae4278d9cf965b61eafe81ab302cfd39ddf559d2a9fd524dbad444a501056dd.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return iaae4278d9cf965b61eafe81ab302cfd39ddf559d2a9fd524dbad444a501056dd.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property childFolders in users
func (m *MailFolderItemRequestBuilder) Patch(options *MailFolderItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *MailFolderItemRequestBuilder) SingleValueExtendedProperties()(*i990fe89bba99adc4df3d6b8e35572e7504bd3f7757624d1b727ea074dd31f631.SingleValueExtendedPropertiesRequestBuilder) {
    return i990fe89bba99adc4df3d6b8e35572e7504bd3f7757624d1b727ea074dd31f631.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.mailFolders.item.childFolders.item.singleValueExtendedProperties.item collection
func (m *MailFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i3df17c93c93c1f05f97aebafc6e566e725950ce62a85f12563acadc9925570d0.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i3df17c93c93c1f05f97aebafc6e566e725950ce62a85f12563acadc9925570d0.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
