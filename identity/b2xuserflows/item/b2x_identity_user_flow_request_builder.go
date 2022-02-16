package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i21e37932240ac829addd6426fbe4f398cd6c1253e77d14a3a2d0ed32b9860f0d "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/languages"
    i33898ff3b4aa69a4b3cd382359fade5f14f2042fa16a4ceea1285f765e40c1eb "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/identityproviders"
    i71e405a8d1fa3f6a758b702c04166c77cbb98f401dadb3fa682f3e2ed2f5589c "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/userattributeassignments"
    ie5d8c8cff8f0ae37cced67067dc8d6c60197107961396166a82aa25c36193142 "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/userflowidentityproviders"
    i4e7268fbaade322b200bff897ba7a9ec706172d25d1fac40768df711c1624a0f "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/userattributeassignments/item"
    iffeb8fd6812998e372f8b86fce623d5b4b055347b27a5e10d253ff43d246f33d "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/languages/item"
)

// B2xIdentityUserFlowRequestBuilder builds and executes requests for operations under \identity\b2xUserFlows\{b2xIdentityUserFlow-id}
type B2xIdentityUserFlowRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// B2xIdentityUserFlowRequestBuilderDeleteOptions options for Delete
type B2xIdentityUserFlowRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// B2xIdentityUserFlowRequestBuilderGetOptions options for Get
type B2xIdentityUserFlowRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *B2xIdentityUserFlowRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// B2xIdentityUserFlowRequestBuilderGetQueryParameters represents entry point for B2X/self-service sign-up identity userflows.
type B2xIdentityUserFlowRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// B2xIdentityUserFlowRequestBuilderPatchOptions options for Patch
type B2xIdentityUserFlowRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.B2xIdentityUserFlow;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewB2xIdentityUserFlowRequestBuilderInternal instantiates a new B2xIdentityUserFlowRequestBuilder and sets the default values.
func NewB2xIdentityUserFlowRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*B2xIdentityUserFlowRequestBuilder) {
    m := &B2xIdentityUserFlowRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewB2xIdentityUserFlowRequestBuilder instantiates a new B2xIdentityUserFlowRequestBuilder and sets the default values.
func NewB2xIdentityUserFlowRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*B2xIdentityUserFlowRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xIdentityUserFlowRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowRequestBuilder) CreateDeleteRequestInformation(options *B2xIdentityUserFlowRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowRequestBuilder) CreateGetRequestInformation(options *B2xIdentityUserFlowRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowRequestBuilder) CreatePatchRequestInformation(options *B2xIdentityUserFlowRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowRequestBuilder) Delete(options *B2xIdentityUserFlowRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowRequestBuilder) Get(options *B2xIdentityUserFlowRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.B2xIdentityUserFlow, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewB2xIdentityUserFlow() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.B2xIdentityUserFlow), nil
}
func (m *B2xIdentityUserFlowRequestBuilder) IdentityProviders()(*i33898ff3b4aa69a4b3cd382359fade5f14f2042fa16a4ceea1285f765e40c1eb.IdentityProvidersRequestBuilder) {
    return i33898ff3b4aa69a4b3cd382359fade5f14f2042fa16a4ceea1285f765e40c1eb.NewIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *B2xIdentityUserFlowRequestBuilder) Languages()(*i21e37932240ac829addd6426fbe4f398cd6c1253e77d14a3a2d0ed32b9860f0d.LanguagesRequestBuilder) {
    return i21e37932240ac829addd6426fbe4f398cd6c1253e77d14a3a2d0ed32b9860f0d.NewLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LanguagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item.languages.item collection
func (m *B2xIdentityUserFlowRequestBuilder) LanguagesById(id string)(*iffeb8fd6812998e372f8b86fce623d5b4b055347b27a5e10d253ff43d246f33d.UserFlowLanguageConfigurationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguageConfiguration_id"] = id
    }
    return iffeb8fd6812998e372f8b86fce623d5b4b055347b27a5e10d253ff43d246f33d.NewUserFlowLanguageConfigurationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowRequestBuilder) Patch(options *B2xIdentityUserFlowRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *B2xIdentityUserFlowRequestBuilder) UserAttributeAssignments()(*i71e405a8d1fa3f6a758b702c04166c77cbb98f401dadb3fa682f3e2ed2f5589c.UserAttributeAssignmentsRequestBuilder) {
    return i71e405a8d1fa3f6a758b702c04166c77cbb98f401dadb3fa682f3e2ed2f5589c.NewUserAttributeAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserAttributeAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item.userAttributeAssignments.item collection
func (m *B2xIdentityUserFlowRequestBuilder) UserAttributeAssignmentsById(id string)(*i4e7268fbaade322b200bff897ba7a9ec706172d25d1fac40768df711c1624a0f.IdentityUserFlowAttributeAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlowAttributeAssignment_id"] = id
    }
    return i4e7268fbaade322b200bff897ba7a9ec706172d25d1fac40768df711c1624a0f.NewIdentityUserFlowAttributeAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *B2xIdentityUserFlowRequestBuilder) UserFlowIdentityProviders()(*ie5d8c8cff8f0ae37cced67067dc8d6c60197107961396166a82aa25c36193142.UserFlowIdentityProvidersRequestBuilder) {
    return ie5d8c8cff8f0ae37cced67067dc8d6c60197107961396166a82aa25c36193142.NewUserFlowIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
