package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i5022b7717b28336f47553a616806d1a28694d39d2336ba16f991ea8b6024102a "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/contacts/item/singlevalueextendedproperties"
    id8c33b60c0d6c294af0eadd9706875d14b368aa0f2a21148a729038ba9a807af "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/contacts/item/multivalueextendedproperties"
    ie51a5139f2522c8a5795eb41dece31f19381bef2fb940355e6b5718ff531a859 "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/contacts/item/extensions"
    if2a9be7d4a1466c82a09f8c40d96131d299667c9a5dc451bc9b8fc5b68ae4f02 "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/contacts/item/photo"
    i40a06616769f69552f5b46470e0c262e3856ffefcc8ad14204283fb68cb55a8a "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/contacts/item/singlevalueextendedproperties/item"
    i99baeae1c43f49fcdb0f1e1d4ecd41a3a52dc95e17cdb56141c586bbcff09078 "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/contacts/item/extensions/item"
    idc046c969eac1b40f7a04ad4a98854a7726a894f0c28ac98ade4d8af2b9cce7e "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/contacts/item/multivalueextendedproperties/item"
)

// ContactItemRequestBuilder provides operations to manage the contacts property of the microsoft.graph.contactFolder entity.
type ContactItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContactItemRequestBuilderDeleteOptions options for Delete
type ContactItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContactItemRequestBuilderGetOptions options for Get
type ContactItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContactItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
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
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Contactable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewContactItemRequestBuilderInternal instantiates a new ContactItemRequestBuilder and sets the default values.
func NewContactItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactItemRequestBuilder) {
    m := &ContactItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/contactFolders/{contactFolder_id}/contacts/{contact_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContactItemRequestBuilder instantiates a new ContactItemRequestBuilder and sets the default values.
func NewContactItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property contacts for users
func (m *ContactItemRequestBuilder) CreateDeleteRequestInformation(options *ContactItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the contacts in the folder. Navigation property. Read-only. Nullable.
func (m *ContactItemRequestBuilder) CreateGetRequestInformation(options *ContactItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property contacts in users
func (m *ContactItemRequestBuilder) CreatePatchRequestInformation(options *ContactItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property contacts for users
func (m *ContactItemRequestBuilder) Delete(options *ContactItemRequestBuilderDeleteOptions)(error) {
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
func (m *ContactItemRequestBuilder) Extensions()(*ie51a5139f2522c8a5795eb41dece31f19381bef2fb940355e6b5718ff531a859.ExtensionsRequestBuilder) {
    return ie51a5139f2522c8a5795eb41dece31f19381bef2fb940355e6b5718ff531a859.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.contactFolders.item.contacts.item.extensions.item collection
func (m *ContactItemRequestBuilder) ExtensionsById(id string)(*i99baeae1c43f49fcdb0f1e1d4ecd41a3a52dc95e17cdb56141c586bbcff09078.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i99baeae1c43f49fcdb0f1e1d4ecd41a3a52dc95e17cdb56141c586bbcff09078.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the contacts in the folder. Navigation property. Read-only. Nullable.
func (m *ContactItemRequestBuilder) Get(options *ContactItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Contactable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateContactFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Contactable), nil
}
func (m *ContactItemRequestBuilder) MultiValueExtendedProperties()(*id8c33b60c0d6c294af0eadd9706875d14b368aa0f2a21148a729038ba9a807af.MultiValueExtendedPropertiesRequestBuilder) {
    return id8c33b60c0d6c294af0eadd9706875d14b368aa0f2a21148a729038ba9a807af.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.contactFolders.item.contacts.item.multiValueExtendedProperties.item collection
func (m *ContactItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*idc046c969eac1b40f7a04ad4a98854a7726a894f0c28ac98ade4d8af2b9cce7e.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return idc046c969eac1b40f7a04ad4a98854a7726a894f0c28ac98ade4d8af2b9cce7e.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property contacts in users
func (m *ContactItemRequestBuilder) Patch(options *ContactItemRequestBuilderPatchOptions)(error) {
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
func (m *ContactItemRequestBuilder) Photo()(*if2a9be7d4a1466c82a09f8c40d96131d299667c9a5dc451bc9b8fc5b68ae4f02.PhotoRequestBuilder) {
    return if2a9be7d4a1466c82a09f8c40d96131d299667c9a5dc451bc9b8fc5b68ae4f02.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactItemRequestBuilder) SingleValueExtendedProperties()(*i5022b7717b28336f47553a616806d1a28694d39d2336ba16f991ea8b6024102a.SingleValueExtendedPropertiesRequestBuilder) {
    return i5022b7717b28336f47553a616806d1a28694d39d2336ba16f991ea8b6024102a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.contactFolders.item.contacts.item.singleValueExtendedProperties.item collection
func (m *ContactItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i40a06616769f69552f5b46470e0c262e3856ffefcc8ad14204283fb68cb55a8a.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i40a06616769f69552f5b46470e0c262e3856ffefcc8ad14204283fb68cb55a8a.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
