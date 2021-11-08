package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i12bbca421752c1a216af39c1b248202a000c0523508478cb12f65bebbdf390b1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/singlevalueextendedproperties"
    i60ab60f82d14e2dac827c4a94548731a12742fd43ee5e3c8ad4e981c63b3430e "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/childfolders"
    i65553c18ffeaa784c3939c2a02713d4581a841081044d39399ed69407816f934 "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/multivalueextendedproperties"
    ida83c3cad119a461d4351a90d275d85f7ab07a13493267b770834de6b0df1e46 "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/contacts"
    i313408c9cc7b38ecc6251cca569a0b7c6640fe8f5556623d6b010c4ad94526c7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/childfolders/item"
    i316f60e76d66ed07f62c18eb8b9dbe052faea000d861699ea2e1e9ee3745269a "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/contacts/item"
    ibb87b307fcec740337cd562826187e93e9cf34543efe25a06a3c17bea866ffc8 "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/singlevalueextendedproperties/item"
    ibd6ce6f8bcc14aa05c644fdf6a85e41e90bf174f08ca7bfa4ed65ee287544500 "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/multivalueextendedproperties/item"
)

// Builds and executes requests for operations under \users\{user-id}\contactFolders\{contactFolder-id}
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
func (m *ContactFolderRequestBuilder) ChildFolders()(*i60ab60f82d14e2dac827c4a94548731a12742fd43ee5e3c8ad4e981c63b3430e.ChildFoldersRequestBuilder) {
    return i60ab60f82d14e2dac827c4a94548731a12742fd43ee5e3c8ad4e981c63b3430e.NewChildFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.users.item.contactFolders.item.childFolders.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContactFolderRequestBuilder) ChildFoldersById(id string)(*i313408c9cc7b38ecc6251cca569a0b7c6640fe8f5556623d6b010c4ad94526c7.ContactFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contactFolder_id1"] = id
    }
    return i313408c9cc7b38ecc6251cca569a0b7c6640fe8f5556623d6b010c4ad94526c7.NewContactFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new ContactFolderRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewContactFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactFolderRequestBuilder) {
    m := &ContactFolderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/contactFolders/{contactFolder_id}{?select}";
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
func (m *ContactFolderRequestBuilder) Contacts()(*ida83c3cad119a461d4351a90d275d85f7ab07a13493267b770834de6b0df1e46.ContactsRequestBuilder) {
    return ida83c3cad119a461d4351a90d275d85f7ab07a13493267b770834de6b0df1e46.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.users.item.contactFolders.item.contacts.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContactFolderRequestBuilder) ContactsById(id string)(*i316f60e76d66ed07f62c18eb8b9dbe052faea000d861699ea2e1e9ee3745269a.ContactRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact_id"] = id
    }
    return i316f60e76d66ed07f62c18eb8b9dbe052faea000d861699ea2e1e9ee3745269a.NewContactRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *ContactFolderRequestBuilder) MultiValueExtendedProperties()(*i65553c18ffeaa784c3939c2a02713d4581a841081044d39399ed69407816f934.MultiValueExtendedPropertiesRequestBuilder) {
    return i65553c18ffeaa784c3939c2a02713d4581a841081044d39399ed69407816f934.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.users.item.contactFolders.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContactFolderRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ibd6ce6f8bcc14aa05c644fdf6a85e41e90bf174f08ca7bfa4ed65ee287544500.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ibd6ce6f8bcc14aa05c644fdf6a85e41e90bf174f08ca7bfa4ed65ee287544500.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *ContactFolderRequestBuilder) SingleValueExtendedProperties()(*i12bbca421752c1a216af39c1b248202a000c0523508478cb12f65bebbdf390b1.SingleValueExtendedPropertiesRequestBuilder) {
    return i12bbca421752c1a216af39c1b248202a000c0523508478cb12f65bebbdf390b1.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.users.item.contactFolders.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContactFolderRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ibb87b307fcec740337cd562826187e93e9cf34543efe25a06a3c17bea866ffc8.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ibb87b307fcec740337cd562826187e93e9cf34543efe25a06a3c17bea866ffc8.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
