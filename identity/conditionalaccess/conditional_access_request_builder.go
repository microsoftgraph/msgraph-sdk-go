package conditionalaccess

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i27d3a1bae6f4577a4de2876ba58fad41be78be972958441916bfc0b15d805933 "github.com/microsoftgraph/msgraph-sdk-go/identity/conditionalaccess/policies"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    ie167212f1f78fb7bf5e6a91e4a3df660b8a5499daeed0e885580a0dde04d710e "github.com/microsoftgraph/msgraph-sdk-go/identity/conditionalaccess/namedlocations"
    i652ce65e13fdceb3d1ed10efc4a8f6c0d893a47d549accb475d48e05eedf3251 "github.com/microsoftgraph/msgraph-sdk-go/identity/conditionalaccess/namedlocations/item"
    i91801d979eb0028b76d066e7195d63b5ed00d456ccdceaa9f6bb9a091401a467 "github.com/microsoftgraph/msgraph-sdk-go/identity/conditionalaccess/policies/item"
)

// Builds and executes requests for operations under \identity\conditionalAccess
type ConditionalAccessRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ConditionalAccessRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ConditionalAccessRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ConditionalAccessRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// the entry point for the Conditional Access (CA) object model.
type ConditionalAccessRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ConditionalAccessRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ConditionalAccessRoot;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new ConditionalAccessRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewConditionalAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ConditionalAccessRequestBuilder) {
    m := &ConditionalAccessRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/conditionalAccess{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ConditionalAccessRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewConditionalAccessRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ConditionalAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalAccessRequestBuilderInternal(urlParams, requestAdapter)
}
// the entry point for the Conditional Access (CA) object model.
// Parameters:
//  - options : Options for the request
func (m *ConditionalAccessRequestBuilder) CreateDeleteRequestInformation(options *ConditionalAccessRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// the entry point for the Conditional Access (CA) object model.
// Parameters:
//  - options : Options for the request
func (m *ConditionalAccessRequestBuilder) CreateGetRequestInformation(options *ConditionalAccessRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// the entry point for the Conditional Access (CA) object model.
// Parameters:
//  - options : Options for the request
func (m *ConditionalAccessRequestBuilder) CreatePatchRequestInformation(options *ConditionalAccessRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// the entry point for the Conditional Access (CA) object model.
// Parameters:
//  - options : Options for the request
func (m *ConditionalAccessRequestBuilder) Delete(options *ConditionalAccessRequestBuilderDeleteOptions)(error) {
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
// the entry point for the Conditional Access (CA) object model.
// Parameters:
//  - options : Options for the request
func (m *ConditionalAccessRequestBuilder) Get(options *ConditionalAccessRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ConditionalAccessRoot, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewConditionalAccessRoot() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ConditionalAccessRoot), nil
}
func (m *ConditionalAccessRequestBuilder) NamedLocations()(*ie167212f1f78fb7bf5e6a91e4a3df660b8a5499daeed0e885580a0dde04d710e.NamedLocationsRequestBuilder) {
    return ie167212f1f78fb7bf5e6a91e4a3df660b8a5499daeed0e885580a0dde04d710e.NewNamedLocationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.conditionalAccess.namedLocations.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ConditionalAccessRequestBuilder) NamedLocationsById(id string)(*i652ce65e13fdceb3d1ed10efc4a8f6c0d893a47d549accb475d48e05eedf3251.NamedLocationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["namedLocation_id"] = id
    }
    return i652ce65e13fdceb3d1ed10efc4a8f6c0d893a47d549accb475d48e05eedf3251.NewNamedLocationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// the entry point for the Conditional Access (CA) object model.
// Parameters:
//  - options : Options for the request
func (m *ConditionalAccessRequestBuilder) Patch(options *ConditionalAccessRequestBuilderPatchOptions)(error) {
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
func (m *ConditionalAccessRequestBuilder) Policies()(*i27d3a1bae6f4577a4de2876ba58fad41be78be972958441916bfc0b15d805933.PoliciesRequestBuilder) {
    return i27d3a1bae6f4577a4de2876ba58fad41be78be972958441916bfc0b15d805933.NewPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.conditionalAccess.policies.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ConditionalAccessRequestBuilder) PoliciesById(id string)(*i91801d979eb0028b76d066e7195d63b5ed00d456ccdceaa9f6bb9a091401a467.ConditionalAccessPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessPolicy_id"] = id
    }
    return i91801d979eb0028b76d066e7195d63b5ed00d456ccdceaa9f6bb9a091401a467.NewConditionalAccessPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
