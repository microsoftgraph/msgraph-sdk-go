package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i5b3c5a1d1e7fbc3d2cab902cce2aa5e8077c4d1ea108f26690d141eea491bc11 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts/item/photo"
    i99160c73b1259c3dbc6c400fc70439ce9e3cc2e871ea6f0a4a0c008b7a8ef536 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts/item/singlevalueextendedproperties"
    ib8cc74adfce6dfc0789ee0acb82354ad0bcc5e628492e567a0c358afb1fe741f "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts/item/extensions"
    ife1d7f7e3b6d757a7201f2349a51e5e4d89c0ec8aa32c691a7bc983d58f03031 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts/item/multivalueextendedproperties"
    i707bf214b57d0fb124dfd1e57eb1ad7127873d5d018f3aa458aa74c2697719cf "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts/item/extensions/item"
    ib15d886582231557e1b5d33561dc0a80a183185d3af5dfb3a67e56a0f47cfe85 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts/item/singlevalueextendedproperties/item"
    ifb7b100c3862aac4d11c9eec0cd8a2dd8abb0681790b9f59c49e5519657c8849 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts/item/multivalueextendedproperties/item"
)

// ContactItemRequestBuilder provides operations to manage the contacts property of the microsoft.graph.contactFolder entity.
type ContactItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContactItemRequestBuilderDeleteOptions options for Delete
type ContactItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ContactItemRequestBuilderGetOptions options for Get
type ContactItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *ContactItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ContactItemRequestBuilderGetQueryParameters the contacts in the folder. Navigation property. Read-only. Nullable.
type ContactItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ContactItemRequestBuilderPatchOptions options for Patch
type ContactItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Contactable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewContactItemRequestBuilderInternal instantiates a new ContactItemRequestBuilder and sets the default values.
func NewContactItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContactItemRequestBuilder) {
    m := &ContactItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/contactFolders/{contactFolder_id}/contacts/{contact_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContactItemRequestBuilder instantiates a new ContactItemRequestBuilder and sets the default values.
func NewContactItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContactItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property contacts for me
func (m *ContactItemRequestBuilder) CreateDeleteRequestInformation(options *ContactItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the contacts in the folder. Navigation property. Read-only. Nullable.
func (m *ContactItemRequestBuilder) CreateGetRequestInformation(options *ContactItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property contacts in me
func (m *ContactItemRequestBuilder) CreatePatchRequestInformation(options *ContactItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property contacts for me
func (m *ContactItemRequestBuilder) Delete(options *ContactItemRequestBuilderDeleteOptions)(error) {
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
// Extensions the extensions property
func (m *ContactItemRequestBuilder) Extensions()(*ib8cc74adfce6dfc0789ee0acb82354ad0bcc5e628492e567a0c358afb1fe741f.ExtensionsRequestBuilder) {
    return ib8cc74adfce6dfc0789ee0acb82354ad0bcc5e628492e567a0c358afb1fe741f.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contactFolders.item.contacts.item.extensions.item collection
func (m *ContactItemRequestBuilder) ExtensionsById(id string)(*i707bf214b57d0fb124dfd1e57eb1ad7127873d5d018f3aa458aa74c2697719cf.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i707bf214b57d0fb124dfd1e57eb1ad7127873d5d018f3aa458aa74c2697719cf.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the contacts in the folder. Navigation property. Read-only. Nullable.
func (m *ContactItemRequestBuilder) Get(options *ContactItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Contactable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContactFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Contactable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *ContactItemRequestBuilder) MultiValueExtendedProperties()(*ife1d7f7e3b6d757a7201f2349a51e5e4d89c0ec8aa32c691a7bc983d58f03031.MultiValueExtendedPropertiesRequestBuilder) {
    return ife1d7f7e3b6d757a7201f2349a51e5e4d89c0ec8aa32c691a7bc983d58f03031.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contactFolders.item.contacts.item.multiValueExtendedProperties.item collection
func (m *ContactItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ifb7b100c3862aac4d11c9eec0cd8a2dd8abb0681790b9f59c49e5519657c8849.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ifb7b100c3862aac4d11c9eec0cd8a2dd8abb0681790b9f59c49e5519657c8849.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property contacts in me
func (m *ContactItemRequestBuilder) Patch(options *ContactItemRequestBuilderPatchOptions)(error) {
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
// Photo the photo property
func (m *ContactItemRequestBuilder) Photo()(*i5b3c5a1d1e7fbc3d2cab902cce2aa5e8077c4d1ea108f26690d141eea491bc11.PhotoRequestBuilder) {
    return i5b3c5a1d1e7fbc3d2cab902cce2aa5e8077c4d1ea108f26690d141eea491bc11.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *ContactItemRequestBuilder) SingleValueExtendedProperties()(*i99160c73b1259c3dbc6c400fc70439ce9e3cc2e871ea6f0a4a0c008b7a8ef536.SingleValueExtendedPropertiesRequestBuilder) {
    return i99160c73b1259c3dbc6c400fc70439ce9e3cc2e871ea6f0a4a0c008b7a8ef536.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contactFolders.item.contacts.item.singleValueExtendedProperties.item collection
func (m *ContactItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ib15d886582231557e1b5d33561dc0a80a183185d3af5dfb3a67e56a0f47cfe85.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ib15d886582231557e1b5d33561dc0a80a183185d3af5dfb3a67e56a0f47cfe85.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
