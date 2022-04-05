package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i133599b4fac6c134a2f752fe88a6312033f8cd2a7e45b8b3dd53b4ca839d76d9 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders/item/contacts"
    i1fc7b2e92f96aa6ce6ba87d3f159178016332ebd3880278adcf1b917efe74b31 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders/item/multivalueextendedproperties"
    iba26fe604a0cbb0b4c8a3b3f3d86114f1f5d3288b9f7c045822a86109b8ecbcb "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders/item/singlevalueextendedproperties"
    ibc8762fdc34cde51d2f827a741dacde78c1fec7ad1688a27333add31120387a0 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders/item/multivalueextendedproperties/item"
    id81a099a3cc874726c1e03f9f676d5725b21fee6fdea29251d8b3e0ff26e7d77 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders/item/contacts/item"
    if98cc278ef17097b18473ff4e44f397a59f675e168984c4ffa285da08b892273 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders/item/singlevalueextendedproperties/item"
)

// ContactFolderItemRequestBuilder provides operations to manage the childFolders property of the microsoft.graph.contactFolder entity.
type ContactFolderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContactFolderItemRequestBuilderDeleteOptions options for Delete
type ContactFolderItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ContactFolderItemRequestBuilderGetOptions options for Get
type ContactFolderItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *ContactFolderItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ContactFolderItemRequestBuilderGetQueryParameters the collection of child folders in the folder. Navigation property. Read-only. Nullable.
type ContactFolderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ContactFolderItemRequestBuilderPatchOptions options for Patch
type ContactFolderItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContactFolderable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewContactFolderItemRequestBuilderInternal instantiates a new ContactFolderItemRequestBuilder and sets the default values.
func NewContactFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContactFolderItemRequestBuilder) {
    m := &ContactFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/contactFolders/{contactFolder_id}/childFolders/{contactFolder_id1}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContactFolderItemRequestBuilder instantiates a new ContactFolderItemRequestBuilder and sets the default values.
func NewContactFolderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContactFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Contacts the contacts property
func (m *ContactFolderItemRequestBuilder) Contacts()(*i133599b4fac6c134a2f752fe88a6312033f8cd2a7e45b8b3dd53b4ca839d76d9.ContactsRequestBuilder) {
    return i133599b4fac6c134a2f752fe88a6312033f8cd2a7e45b8b3dd53b4ca839d76d9.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contactFolders.item.childFolders.item.contacts.item collection
func (m *ContactFolderItemRequestBuilder) ContactsById(id string)(*id81a099a3cc874726c1e03f9f676d5725b21fee6fdea29251d8b3e0ff26e7d77.ContactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact_id"] = id
    }
    return id81a099a3cc874726c1e03f9f676d5725b21fee6fdea29251d8b3e0ff26e7d77.NewContactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property childFolders for me
func (m *ContactFolderItemRequestBuilder) CreateDeleteRequestInformation(options *ContactFolderItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of child folders in the folder. Navigation property. Read-only. Nullable.
func (m *ContactFolderItemRequestBuilder) CreateGetRequestInformation(options *ContactFolderItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property childFolders in me
func (m *ContactFolderItemRequestBuilder) CreatePatchRequestInformation(options *ContactFolderItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property childFolders for me
func (m *ContactFolderItemRequestBuilder) Delete(options *ContactFolderItemRequestBuilderDeleteOptions)(error) {
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
// Get the collection of child folders in the folder. Navigation property. Read-only. Nullable.
func (m *ContactFolderItemRequestBuilder) Get(options *ContactFolderItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContactFolderable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContactFolderFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContactFolderable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *ContactFolderItemRequestBuilder) MultiValueExtendedProperties()(*i1fc7b2e92f96aa6ce6ba87d3f159178016332ebd3880278adcf1b917efe74b31.MultiValueExtendedPropertiesRequestBuilder) {
    return i1fc7b2e92f96aa6ce6ba87d3f159178016332ebd3880278adcf1b917efe74b31.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contactFolders.item.childFolders.item.multiValueExtendedProperties.item collection
func (m *ContactFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ibc8762fdc34cde51d2f827a741dacde78c1fec7ad1688a27333add31120387a0.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ibc8762fdc34cde51d2f827a741dacde78c1fec7ad1688a27333add31120387a0.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property childFolders in me
func (m *ContactFolderItemRequestBuilder) Patch(options *ContactFolderItemRequestBuilderPatchOptions)(error) {
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
func (m *ContactFolderItemRequestBuilder) SingleValueExtendedProperties()(*iba26fe604a0cbb0b4c8a3b3f3d86114f1f5d3288b9f7c045822a86109b8ecbcb.SingleValueExtendedPropertiesRequestBuilder) {
    return iba26fe604a0cbb0b4c8a3b3f3d86114f1f5d3288b9f7c045822a86109b8ecbcb.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contactFolders.item.childFolders.item.singleValueExtendedProperties.item collection
func (m *ContactFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*if98cc278ef17097b18473ff4e44f397a59f675e168984c4ffa285da08b892273.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return if98cc278ef17097b18473ff4e44f397a59f675e168984c4ffa285da08b892273.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
