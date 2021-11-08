package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0035901ffc213fc4b59f4993863d1daec79b347f97e62b284d79c7f4e51e541a "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders"
    ib939870f831b5d1e6f646bd6624ff59df0efd3d63e40ad5d6874229710fb5a0f "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts"
    ibe30f52eff3d7782e030050a6a70ab017912ac09c1eaaafeb1f5bc12d8b3a5c0 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/singlevalueextendedproperties"
    ic1c22c979ab4e96bd99861c36b59000262714ad01fb89b523d57edd0e2710294 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/multivalueextendedproperties"
    i6249342797ad051a9fde0138fc7a4de2b582b010423ef26c264a2a10b971def7 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/multivalueextendedproperties/item"
    i73617fc20b6e8870c7f350b9a031e29042cc9e456837fc4b6df34bbf3918f396 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/contacts/item"
    i7b320e3aa4ff049100b60cd68585037930958ce6c760db10289e1a7baef49df1 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/childfolders/item"
    i82ebac0048f2c4d4678fe54780699f0f1230c24f422393c29fcc9bccf43bf443 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item/singlevalueextendedproperties/item"
)

// Builds and executes requests for operations under \me\contactFolders\{contactFolder-id}
type ContactFolderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ContactFolderRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ContactFolderRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContactFolderRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The user's contacts folders. Read-only. Nullable.
type ContactFolderRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ContactFolderRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContactFolder;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ContactFolderRequestBuilder) ChildFolders()(*i0035901ffc213fc4b59f4993863d1daec79b347f97e62b284d79c7f4e51e541a.ChildFoldersRequestBuilder) {
    return i0035901ffc213fc4b59f4993863d1daec79b347f97e62b284d79c7f4e51e541a.NewChildFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.contactFolders.item.childFolders.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContactFolderRequestBuilder) ChildFoldersById(id string)(*i7b320e3aa4ff049100b60cd68585037930958ce6c760db10289e1a7baef49df1.ContactFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contactFolder_id1"] = id
    }
    return i7b320e3aa4ff049100b60cd68585037930958ce6c760db10289e1a7baef49df1.NewContactFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new ContactFolderRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewContactFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactFolderRequestBuilder) {
    m := &ContactFolderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/contactFolders/{contactFolder_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ContactFolderRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewContactFolderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactFolderRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContactFolderRequestBuilder) Contacts()(*ib939870f831b5d1e6f646bd6624ff59df0efd3d63e40ad5d6874229710fb5a0f.ContactsRequestBuilder) {
    return ib939870f831b5d1e6f646bd6624ff59df0efd3d63e40ad5d6874229710fb5a0f.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.contactFolders.item.contacts.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContactFolderRequestBuilder) ContactsById(id string)(*i73617fc20b6e8870c7f350b9a031e29042cc9e456837fc4b6df34bbf3918f396.ContactRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact_id"] = id
    }
    return i73617fc20b6e8870c7f350b9a031e29042cc9e456837fc4b6df34bbf3918f396.NewContactRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The user's contacts folders. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ContactFolderRequestBuilder) CreateDeleteRequestInformation(options *ContactFolderRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The user's contacts folders. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ContactFolderRequestBuilder) CreateGetRequestInformation(options *ContactFolderRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// The user's contacts folders. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ContactFolderRequestBuilder) CreatePatchRequestInformation(options *ContactFolderRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The user's contacts folders. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ContactFolderRequestBuilder) Delete(options *ContactFolderRequestBuilderDeleteOptions)(error) {
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
// The user's contacts folders. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ContactFolderRequestBuilder) Get(options *ContactFolderRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContactFolder, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewContactFolder() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContactFolder), nil
}
func (m *ContactFolderRequestBuilder) MultiValueExtendedProperties()(*ic1c22c979ab4e96bd99861c36b59000262714ad01fb89b523d57edd0e2710294.MultiValueExtendedPropertiesRequestBuilder) {
    return ic1c22c979ab4e96bd99861c36b59000262714ad01fb89b523d57edd0e2710294.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.contactFolders.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContactFolderRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i6249342797ad051a9fde0138fc7a4de2b582b010423ef26c264a2a10b971def7.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i6249342797ad051a9fde0138fc7a4de2b582b010423ef26c264a2a10b971def7.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The user's contacts folders. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ContactFolderRequestBuilder) Patch(options *ContactFolderRequestBuilderPatchOptions)(error) {
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
func (m *ContactFolderRequestBuilder) SingleValueExtendedProperties()(*ibe30f52eff3d7782e030050a6a70ab017912ac09c1eaaafeb1f5bc12d8b3a5c0.SingleValueExtendedPropertiesRequestBuilder) {
    return ibe30f52eff3d7782e030050a6a70ab017912ac09c1eaaafeb1f5bc12d8b3a5c0.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.contactFolders.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContactFolderRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i82ebac0048f2c4d4678fe54780699f0f1230c24f422393c29fcc9bccf43bf443.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i82ebac0048f2c4d4678fe54780699f0f1230c24f422393c29fcc9bccf43bf443.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
