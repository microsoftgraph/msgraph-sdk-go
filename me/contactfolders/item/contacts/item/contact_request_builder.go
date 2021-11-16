package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i5b3c5a1d1e7fbc3d2cab902cce2aa5e8077c4d1ea108f26690d141eea491bc11 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts/item/photo"
    i99160c73b1259c3dbc6c400fc70439ce9e3cc2e871ea6f0a4a0c008b7a8ef536 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts/item/singlevalueextendedproperties"
    ib8cc74adfce6dfc0789ee0acb82354ad0bcc5e628492e567a0c358afb1fe741f "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts/item/extensions"
    ife1d7f7e3b6d757a7201f2349a51e5e4d89c0ec8aa32c691a7bc983d58f03031 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts/item/multivalueextendedproperties"
    i707bf214b57d0fb124dfd1e57eb1ad7127873d5d018f3aa458aa74c2697719cf "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts/item/extensions/item"
    ib15d886582231557e1b5d33561dc0a80a183185d3af5dfb3a67e56a0f47cfe85 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts/item/singlevalueextendedproperties/item"
    ifb7b100c3862aac4d11c9eec0cd8a2dd8abb0681790b9f59c49e5519657c8849 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts/item/multivalueextendedproperties/item"
)

// Builds and executes requests for operations under \me\contactFolders\{contactFolder-id}\contacts\{contact-id}
type ContactRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ContactRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ContactRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContactRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The contacts in the folder. Navigation property. Read-only. Nullable.
type ContactRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ContactRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Contact;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new ContactRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewContactRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactRequestBuilder) {
    m := &ContactRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/contactFolders/{contactFolder_id}/contacts/{contact_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ContactRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewContactRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactRequestBuilderInternal(urlParams, requestAdapter)
}
// The contacts in the folder. Navigation property. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ContactRequestBuilder) CreateDeleteRequestInformation(options *ContactRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The contacts in the folder. Navigation property. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ContactRequestBuilder) CreateGetRequestInformation(options *ContactRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The contacts in the folder. Navigation property. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ContactRequestBuilder) CreatePatchRequestInformation(options *ContactRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The contacts in the folder. Navigation property. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ContactRequestBuilder) Delete(options *ContactRequestBuilderDeleteOptions)(error) {
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
func (m *ContactRequestBuilder) Extensions()(*ib8cc74adfce6dfc0789ee0acb82354ad0bcc5e628492e567a0c358afb1fe741f.ExtensionsRequestBuilder) {
    return ib8cc74adfce6dfc0789ee0acb82354ad0bcc5e628492e567a0c358afb1fe741f.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contactFolders.item.contacts.item.extensions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContactRequestBuilder) ExtensionsById(id string)(*i707bf214b57d0fb124dfd1e57eb1ad7127873d5d018f3aa458aa74c2697719cf.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i707bf214b57d0fb124dfd1e57eb1ad7127873d5d018f3aa458aa74c2697719cf.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The contacts in the folder. Navigation property. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ContactRequestBuilder) Get(options *ContactRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Contact, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewContact() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Contact), nil
}
func (m *ContactRequestBuilder) MultiValueExtendedProperties()(*ife1d7f7e3b6d757a7201f2349a51e5e4d89c0ec8aa32c691a7bc983d58f03031.MultiValueExtendedPropertiesRequestBuilder) {
    return ife1d7f7e3b6d757a7201f2349a51e5e4d89c0ec8aa32c691a7bc983d58f03031.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contactFolders.item.contacts.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContactRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ifb7b100c3862aac4d11c9eec0cd8a2dd8abb0681790b9f59c49e5519657c8849.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ifb7b100c3862aac4d11c9eec0cd8a2dd8abb0681790b9f59c49e5519657c8849.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The contacts in the folder. Navigation property. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ContactRequestBuilder) Patch(options *ContactRequestBuilderPatchOptions)(error) {
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
func (m *ContactRequestBuilder) Photo()(*i5b3c5a1d1e7fbc3d2cab902cce2aa5e8077c4d1ea108f26690d141eea491bc11.PhotoRequestBuilder) {
    return i5b3c5a1d1e7fbc3d2cab902cce2aa5e8077c4d1ea108f26690d141eea491bc11.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactRequestBuilder) SingleValueExtendedProperties()(*i99160c73b1259c3dbc6c400fc70439ce9e3cc2e871ea6f0a4a0c008b7a8ef536.SingleValueExtendedPropertiesRequestBuilder) {
    return i99160c73b1259c3dbc6c400fc70439ce9e3cc2e871ea6f0a4a0c008b7a8ef536.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contactFolders.item.contacts.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContactRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ib15d886582231557e1b5d33561dc0a80a183185d3af5dfb3a67e56a0f47cfe85.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ib15d886582231557e1b5d33561dc0a80a183185d3af5dfb3a67e56a0f47cfe85.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
