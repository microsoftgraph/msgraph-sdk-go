package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1ba18e3209541b1af1f9b854baf8bb74f243072ce23325f9f3a4b2ed2b083afd "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders/item/contacts/item/photo"
    i2e96a860e5ba0d38f5d7d5e8e9dbad3a0783d888435154e1ba38bc0438733511 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders/item/contacts/item/extensions"
    i3585d830b5ea1b53168178d1dbc60876711e9d6b5fc5880f82daae458e857c93 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders/item/contacts/item/singlevalueextendedproperties"
    ibd4921c05ad2da86cb447c5801b5ba4878870939a085c34dc020ef310dd85fe8 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders/item/contacts/item/multivalueextendedproperties"
    i4fdb5e70dee965993aa52f85502010d479a4611820f5b3a5214b7e3ff7c7247f "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders/item/contacts/item/multivalueextendedproperties/item"
    i82f3dfbd3929dec4ed9d7a9c791d969e8c06e84c0fe48c8aa642555b61648b1f "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders/item/contacts/item/singlevalueextendedproperties/item"
    if973172fef7e6df1f7ece5b38f56f6c29021660408ae8e3ddb92713cb0c8b498 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders/item/contacts/item/extensions/item"
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
    m.urlTemplate = "{+baseurl}/me/contactFolders/{contactFolder_id}/childFolders/{contactFolder_id1}/contacts/{contact_id}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property contacts for me
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
// CreatePatchRequestInformation update the navigation property contacts in me
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
// Delete delete navigation property contacts for me
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
func (m *ContactItemRequestBuilder) Extensions()(*i2e96a860e5ba0d38f5d7d5e8e9dbad3a0783d888435154e1ba38bc0438733511.ExtensionsRequestBuilder) {
    return i2e96a860e5ba0d38f5d7d5e8e9dbad3a0783d888435154e1ba38bc0438733511.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contactFolders.item.childFolders.item.contacts.item.extensions.item collection
func (m *ContactItemRequestBuilder) ExtensionsById(id string)(*if973172fef7e6df1f7ece5b38f56f6c29021660408ae8e3ddb92713cb0c8b498.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return if973172fef7e6df1f7ece5b38f56f6c29021660408ae8e3ddb92713cb0c8b498.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *ContactItemRequestBuilder) MultiValueExtendedProperties()(*ibd4921c05ad2da86cb447c5801b5ba4878870939a085c34dc020ef310dd85fe8.MultiValueExtendedPropertiesRequestBuilder) {
    return ibd4921c05ad2da86cb447c5801b5ba4878870939a085c34dc020ef310dd85fe8.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contactFolders.item.childFolders.item.contacts.item.multiValueExtendedProperties.item collection
func (m *ContactItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i4fdb5e70dee965993aa52f85502010d479a4611820f5b3a5214b7e3ff7c7247f.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i4fdb5e70dee965993aa52f85502010d479a4611820f5b3a5214b7e3ff7c7247f.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property contacts in me
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
func (m *ContactItemRequestBuilder) Photo()(*i1ba18e3209541b1af1f9b854baf8bb74f243072ce23325f9f3a4b2ed2b083afd.PhotoRequestBuilder) {
    return i1ba18e3209541b1af1f9b854baf8bb74f243072ce23325f9f3a4b2ed2b083afd.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactItemRequestBuilder) SingleValueExtendedProperties()(*i3585d830b5ea1b53168178d1dbc60876711e9d6b5fc5880f82daae458e857c93.SingleValueExtendedPropertiesRequestBuilder) {
    return i3585d830b5ea1b53168178d1dbc60876711e9d6b5fc5880f82daae458e857c93.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contactFolders.item.childFolders.item.contacts.item.singleValueExtendedProperties.item collection
func (m *ContactItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i82f3dfbd3929dec4ed9d7a9c791d969e8c06e84c0fe48c8aa642555b61648b1f.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i82f3dfbd3929dec4ed9d7a9c791d969e8c06e84c0fe48c8aa642555b61648b1f.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
