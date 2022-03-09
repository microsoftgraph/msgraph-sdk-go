package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i21e37932240ac829addd6426fbe4f398cd6c1253e77d14a3a2d0ed32b9860f0d "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/languages"
    i33898ff3b4aa69a4b3cd382359fade5f14f2042fa16a4ceea1285f765e40c1eb "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/identityproviders"
    i71e405a8d1fa3f6a758b702c04166c77cbb98f401dadb3fa682f3e2ed2f5589c "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/userattributeassignments"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    ie5d8c8cff8f0ae37cced67067dc8d6c60197107961396166a82aa25c36193142 "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/userflowidentityproviders"
    i03e59e65e6bc6e940d593a951d05a91effb188992892de8b3543badf8989532d "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/identityproviders/item"
    i4e7268fbaade322b200bff897ba7a9ec706172d25d1fac40768df711c1624a0f "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/userattributeassignments/item"
    i8eca55aa158e5bb2109c9b40a378aea7b85dc3cd79c49336bb5abb4e5653f8d0 "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/userflowidentityproviders/item"
    iffeb8fd6812998e372f8b86fce623d5b4b055347b27a5e10d253ff43d246f33d "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/languages/item"
)

// B2xIdentityUserFlowItemRequestBuilder provides operations to manage the b2xUserFlows property of the microsoft.graph.identityContainer entity.
type B2xIdentityUserFlowItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// B2xIdentityUserFlowItemRequestBuilderDeleteOptions options for Delete
type B2xIdentityUserFlowItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// B2xIdentityUserFlowItemRequestBuilderGetOptions options for Get
type B2xIdentityUserFlowItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *B2xIdentityUserFlowItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// B2xIdentityUserFlowItemRequestBuilderGetQueryParameters represents entry point for B2X/self-service sign-up identity userflows.
type B2xIdentityUserFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// B2xIdentityUserFlowItemRequestBuilderPatchOptions options for Patch
type B2xIdentityUserFlowItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.B2xIdentityUserFlowable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewB2xIdentityUserFlowItemRequestBuilderInternal instantiates a new B2xIdentityUserFlowItemRequestBuilder and sets the default values.
func NewB2xIdentityUserFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*B2xIdentityUserFlowItemRequestBuilder) {
    m := &B2xIdentityUserFlowItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewB2xIdentityUserFlowItemRequestBuilder instantiates a new B2xIdentityUserFlowItemRequestBuilder and sets the default values.
func NewB2xIdentityUserFlowItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*B2xIdentityUserFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xIdentityUserFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property b2xUserFlows for identity
func (m *B2xIdentityUserFlowItemRequestBuilder) CreateDeleteRequestInformation(options *B2xIdentityUserFlowItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *B2xIdentityUserFlowItemRequestBuilder) CreateGetRequestInformation(options *B2xIdentityUserFlowItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property b2xUserFlows in identity
func (m *B2xIdentityUserFlowItemRequestBuilder) CreatePatchRequestInformation(options *B2xIdentityUserFlowItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property b2xUserFlows for identity
func (m *B2xIdentityUserFlowItemRequestBuilder) Delete(options *B2xIdentityUserFlowItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowItemRequestBuilder) Get(options *B2xIdentityUserFlowItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.B2xIdentityUserFlowable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateB2xIdentityUserFlowFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.B2xIdentityUserFlowable), nil
}
func (m *B2xIdentityUserFlowItemRequestBuilder) IdentityProviders()(*i33898ff3b4aa69a4b3cd382359fade5f14f2042fa16a4ceea1285f765e40c1eb.IdentityProvidersRequestBuilder) {
    return i33898ff3b4aa69a4b3cd382359fade5f14f2042fa16a4ceea1285f765e40c1eb.NewIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityProvidersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item.identityProviders.item collection
func (m *B2xIdentityUserFlowItemRequestBuilder) IdentityProvidersById(id string)(*i03e59e65e6bc6e940d593a951d05a91effb188992892de8b3543badf8989532d.IdentityProviderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProvider_id"] = id
    }
    return i03e59e65e6bc6e940d593a951d05a91effb188992892de8b3543badf8989532d.NewIdentityProviderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *B2xIdentityUserFlowItemRequestBuilder) Languages()(*i21e37932240ac829addd6426fbe4f398cd6c1253e77d14a3a2d0ed32b9860f0d.LanguagesRequestBuilder) {
    return i21e37932240ac829addd6426fbe4f398cd6c1253e77d14a3a2d0ed32b9860f0d.NewLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LanguagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item.languages.item collection
func (m *B2xIdentityUserFlowItemRequestBuilder) LanguagesById(id string)(*iffeb8fd6812998e372f8b86fce623d5b4b055347b27a5e10d253ff43d246f33d.UserFlowLanguageConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguageConfiguration_id"] = id
    }
    return iffeb8fd6812998e372f8b86fce623d5b4b055347b27a5e10d253ff43d246f33d.NewUserFlowLanguageConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property b2xUserFlows in identity
func (m *B2xIdentityUserFlowItemRequestBuilder) Patch(options *B2xIdentityUserFlowItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *B2xIdentityUserFlowItemRequestBuilder) UserAttributeAssignments()(*i71e405a8d1fa3f6a758b702c04166c77cbb98f401dadb3fa682f3e2ed2f5589c.UserAttributeAssignmentsRequestBuilder) {
    return i71e405a8d1fa3f6a758b702c04166c77cbb98f401dadb3fa682f3e2ed2f5589c.NewUserAttributeAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserAttributeAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item.userAttributeAssignments.item collection
func (m *B2xIdentityUserFlowItemRequestBuilder) UserAttributeAssignmentsById(id string)(*i4e7268fbaade322b200bff897ba7a9ec706172d25d1fac40768df711c1624a0f.IdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlowAttributeAssignment_id"] = id
    }
    return i4e7268fbaade322b200bff897ba7a9ec706172d25d1fac40768df711c1624a0f.NewIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *B2xIdentityUserFlowItemRequestBuilder) UserFlowIdentityProviders()(*ie5d8c8cff8f0ae37cced67067dc8d6c60197107961396166a82aa25c36193142.UserFlowIdentityProvidersRequestBuilder) {
    return ie5d8c8cff8f0ae37cced67067dc8d6c60197107961396166a82aa25c36193142.NewUserFlowIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserFlowIdentityProvidersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item.userFlowIdentityProviders.item collection
func (m *B2xIdentityUserFlowItemRequestBuilder) UserFlowIdentityProvidersById(id string)(*i8eca55aa158e5bb2109c9b40a378aea7b85dc3cd79c49336bb5abb4e5653f8d0.IdentityProviderBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProviderBase_id"] = id
    }
    return i8eca55aa158e5bb2109c9b40a378aea7b85dc3cd79c49336bb5abb4e5653f8d0.NewIdentityProviderBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
