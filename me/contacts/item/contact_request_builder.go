package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7900dfd0bd161ba16cfab166ed3b4b73f28d6b6bdf7693846d0060cbc83fe7c6 "github.com/microsoftgraph/msgraph-sdk-go/me/contacts/item/multivalueextendedproperties"
    i89438ec9ae223d0bedcb8242f200f32139b896bc60ecdc15e450b1129f445d59 "github.com/microsoftgraph/msgraph-sdk-go/me/contacts/item/extensions"
    ib7bc02a205ca7d6cb9706bfaf945847e5f18513e6f28623c6d74892d052a40ec "github.com/microsoftgraph/msgraph-sdk-go/me/contacts/item/photo"
    if2cff29c6c010fcdd117f7b6fb306b963667350bb66aaa169a0bffaa4c5db6aa "github.com/microsoftgraph/msgraph-sdk-go/me/contacts/item/singlevalueextendedproperties"
    i26ee4d399149ce9fa1f9599f3cf1e90e8738ca46419f61c3ad0cd866e26072f7 "github.com/microsoftgraph/msgraph-sdk-go/me/contacts/item/singlevalueextendedproperties/item"
    i293cd72f0585f8a5a1ed14a6a1885f09bb2e161afac43c5f2472eb8591edbf68 "github.com/microsoftgraph/msgraph-sdk-go/me/contacts/item/extensions/item"
    iea5b1c178ca0b49715ad6dd95962d56bd9009e8e6ae173d447b2c707504f6edc "github.com/microsoftgraph/msgraph-sdk-go/me/contacts/item/multivalueextendedproperties/item"
)

// ContactRequestBuilder builds and executes requests for operations under \me\contacts\{contact-id}
type ContactRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContactRequestBuilderDeleteOptions options for Delete
type ContactRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContactRequestBuilderGetOptions options for Get
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
// ContactRequestBuilderGetQueryParameters the user's contacts. Read-only. Nullable.
type ContactRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select_escaped []string;
}
// ContactRequestBuilderPatchOptions options for Patch
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
// NewContactRequestBuilderInternal instantiates a new ContactRequestBuilder and sets the default values.
func NewContactRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactRequestBuilder) {
    m := &ContactRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/contacts/{contact_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContactRequestBuilder instantiates a new ContactRequestBuilder and sets the default values.
func NewContactRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the user's contacts. Read-only. Nullable.
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
// CreateGetRequestInformation the user's contacts. Read-only. Nullable.
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
// CreatePatchRequestInformation the user's contacts. Read-only. Nullable.
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
// Delete the user's contacts. Read-only. Nullable.
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
func (m *ContactRequestBuilder) Extensions()(*i89438ec9ae223d0bedcb8242f200f32139b896bc60ecdc15e450b1129f445d59.ExtensionsRequestBuilder) {
    return i89438ec9ae223d0bedcb8242f200f32139b896bc60ecdc15e450b1129f445d59.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contacts.item.extensions.item collection
func (m *ContactRequestBuilder) ExtensionsById(id string)(*i293cd72f0585f8a5a1ed14a6a1885f09bb2e161afac43c5f2472eb8591edbf68.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i293cd72f0585f8a5a1ed14a6a1885f09bb2e161afac43c5f2472eb8591edbf68.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the user's contacts. Read-only. Nullable.
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
func (m *ContactRequestBuilder) MultiValueExtendedProperties()(*i7900dfd0bd161ba16cfab166ed3b4b73f28d6b6bdf7693846d0060cbc83fe7c6.MultiValueExtendedPropertiesRequestBuilder) {
    return i7900dfd0bd161ba16cfab166ed3b4b73f28d6b6bdf7693846d0060cbc83fe7c6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contacts.item.multiValueExtendedProperties.item collection
func (m *ContactRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iea5b1c178ca0b49715ad6dd95962d56bd9009e8e6ae173d447b2c707504f6edc.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return iea5b1c178ca0b49715ad6dd95962d56bd9009e8e6ae173d447b2c707504f6edc.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the user's contacts. Read-only. Nullable.
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
func (m *ContactRequestBuilder) Photo()(*ib7bc02a205ca7d6cb9706bfaf945847e5f18513e6f28623c6d74892d052a40ec.PhotoRequestBuilder) {
    return ib7bc02a205ca7d6cb9706bfaf945847e5f18513e6f28623c6d74892d052a40ec.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactRequestBuilder) SingleValueExtendedProperties()(*if2cff29c6c010fcdd117f7b6fb306b963667350bb66aaa169a0bffaa4c5db6aa.SingleValueExtendedPropertiesRequestBuilder) {
    return if2cff29c6c010fcdd117f7b6fb306b963667350bb66aaa169a0bffaa4c5db6aa.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contacts.item.singleValueExtendedProperties.item collection
func (m *ContactRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i26ee4d399149ce9fa1f9599f3cf1e90e8738ca46419f61c3ad0cd866e26072f7.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i26ee4d399149ce9fa1f9599f3cf1e90e8738ca46419f61c3ad0cd866e26072f7.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
