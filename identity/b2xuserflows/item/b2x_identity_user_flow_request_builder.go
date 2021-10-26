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

// Builds and executes requests for operations under \identity\b2xUserFlows\{b2xIdentityUserFlow-id}
type B2xIdentityUserFlowRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Represents entry point for B2X/self-service sign-up identity userflows.
type B2xIdentityUserFlowRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new B2xIdentityUserFlowRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewB2xIdentityUserFlowRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*B2xIdentityUserFlowRequestBuilder) {
    m := &B2xIdentityUserFlowRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/identity/b2xUserFlows/{b2xIdentityUserFlow_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new B2xIdentityUserFlowRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewB2xIdentityUserFlowRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*B2xIdentityUserFlowRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xIdentityUserFlowRequestBuilderInternal(urlParams, requestAdapter)
}
// Represents entry point for B2X/self-service sign-up identity userflows.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *B2xIdentityUserFlowRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Represents entry point for B2X/self-service sign-up identity userflows.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *B2xIdentityUserFlowRequestBuilder) CreateGetRequestInformation(q func (value *B2xIdentityUserFlowRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(B2xIdentityUserFlowRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Represents entry point for B2X/self-service sign-up identity userflows.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *B2xIdentityUserFlowRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.B2xIdentityUserFlow, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Represents entry point for B2X/self-service sign-up identity userflows.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *B2xIdentityUserFlowRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
// Represents entry point for B2X/self-service sign-up identity userflows.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *B2xIdentityUserFlowRequestBuilder) Get(q func (value *B2xIdentityUserFlowRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.B2xIdentityUserFlow, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewB2xIdentityUserFlow() }, responseHandler)
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
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item.languages.item collection
// Parameters:
//  - id : Unique identifier of the item
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
// Represents entry point for B2X/self-service sign-up identity userflows.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *B2xIdentityUserFlowRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.B2xIdentityUserFlow, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *B2xIdentityUserFlowRequestBuilder) UserAttributeAssignments()(*i71e405a8d1fa3f6a758b702c04166c77cbb98f401dadb3fa682f3e2ed2f5589c.UserAttributeAssignmentsRequestBuilder) {
    return i71e405a8d1fa3f6a758b702c04166c77cbb98f401dadb3fa682f3e2ed2f5589c.NewUserAttributeAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item.userAttributeAssignments.item collection
// Parameters:
//  - id : Unique identifier of the item
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
