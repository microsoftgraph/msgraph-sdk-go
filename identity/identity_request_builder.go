package identity

import (
    i376515807007b3519d82a5a4098cd62c89440575d253961ff5f37dfa7e2c4cf9 "github.com/microsoftgraph/msgraph-sdk-go/identity/userflowattributes"
    i39f9aa738972c580cae1de0901f292bddc11d0b2dfb37b93a8b3bf482b3425b6 "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows"
    i6d5364f470a0f1371f3cf8eca71916ebe1920a271c423b65a9c2d093319d8c92 "github.com/microsoftgraph/msgraph-sdk-go/identity/conditionalaccess"
    ia730aeb1863e64ad647ac03b721c109d7b3275e7f50b8ee4b3c6d898a27d4ae4 "github.com/microsoftgraph/msgraph-sdk-go/identity/apiconnectors"
    id5876d25bb1b14598cec17c6b206e367f3fec3f5b79166c83290f819d9677a47 "github.com/microsoftgraph/msgraph-sdk-go/identity/identityproviders"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2507f16e0650c8332777de1431c18c6feb0e9e9b6447d81e77fe14756eb86ba5 "github.com/microsoftgraph/msgraph-sdk-go/identity/userflowattributes/item"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i96f1cf51060a4b7eb52e2dfb3dda5198789ea46f04c6ab2c2f846bbeb4889b16 "github.com/microsoftgraph/msgraph-sdk-go/identity/apiconnectors/item"
    i97c8ca8708adfdaff0fc183949b8d2204a16d0d0d8e4cf7437d167ee9b2963cf "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item"
    ib0a61c24c37402097c09c6d80120c946edfc11c02e39bf8abbccceafe4660080 "github.com/microsoftgraph/msgraph-sdk-go/identity/identityproviders/item"
)

// IdentityRequestBuilder builds and executes requests for operations under \identity
type IdentityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// IdentityRequestBuilderGetOptions options for Get
type IdentityRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *IdentityRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IdentityRequestBuilderGetQueryParameters get identity
type IdentityRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// IdentityRequestBuilderPatchOptions options for Patch
type IdentityRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IdentityContainer;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *IdentityRequestBuilder) ApiConnectors()(*ia730aeb1863e64ad647ac03b721c109d7b3275e7f50b8ee4b3c6d898a27d4ae4.ApiConnectorsRequestBuilder) {
    return ia730aeb1863e64ad647ac03b721c109d7b3275e7f50b8ee4b3c6d898a27d4ae4.NewApiConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApiConnectorsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.apiConnectors.item collection
func (m *IdentityRequestBuilder) ApiConnectorsById(id string)(*i96f1cf51060a4b7eb52e2dfb3dda5198789ea46f04c6ab2c2f846bbeb4889b16.IdentityApiConnectorRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityApiConnector_id"] = id
    }
    return i96f1cf51060a4b7eb52e2dfb3dda5198789ea46f04c6ab2c2f846bbeb4889b16.NewIdentityApiConnectorRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *IdentityRequestBuilder) B2xUserFlows()(*i39f9aa738972c580cae1de0901f292bddc11d0b2dfb37b93a8b3bf482b3425b6.B2xUserFlowsRequestBuilder) {
    return i39f9aa738972c580cae1de0901f292bddc11d0b2dfb37b93a8b3bf482b3425b6.NewB2xUserFlowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// B2xUserFlowsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item collection
func (m *IdentityRequestBuilder) B2xUserFlowsById(id string)(*i97c8ca8708adfdaff0fc183949b8d2204a16d0d0d8e4cf7437d167ee9b2963cf.B2xIdentityUserFlowRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["b2xIdentityUserFlow_id"] = id
    }
    return i97c8ca8708adfdaff0fc183949b8d2204a16d0d0d8e4cf7437d167ee9b2963cf.NewB2xIdentityUserFlowRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *IdentityRequestBuilder) ConditionalAccess()(*i6d5364f470a0f1371f3cf8eca71916ebe1920a271c423b65a9c2d093319d8c92.ConditionalAccessRequestBuilder) {
    return i6d5364f470a0f1371f3cf8eca71916ebe1920a271c423b65a9c2d093319d8c92.NewConditionalAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIdentityRequestBuilderInternal instantiates a new IdentityRequestBuilder and sets the default values.
func NewIdentityRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IdentityRequestBuilder) {
    m := &IdentityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityRequestBuilder instantiates a new IdentityRequestBuilder and sets the default values.
func NewIdentityRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IdentityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get identity
func (m *IdentityRequestBuilder) CreateGetRequestInformation(options *IdentityRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update identity
func (m *IdentityRequestBuilder) CreatePatchRequestInformation(options *IdentityRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get identity
func (m *IdentityRequestBuilder) Get(options *IdentityRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IdentityContainer, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewIdentityContainer() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IdentityContainer), nil
}
func (m *IdentityRequestBuilder) IdentityProviders()(*id5876d25bb1b14598cec17c6b206e367f3fec3f5b79166c83290f819d9677a47.IdentityProvidersRequestBuilder) {
    return id5876d25bb1b14598cec17c6b206e367f3fec3f5b79166c83290f819d9677a47.NewIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityProvidersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.identityProviders.item collection
func (m *IdentityRequestBuilder) IdentityProvidersById(id string)(*ib0a61c24c37402097c09c6d80120c946edfc11c02e39bf8abbccceafe4660080.IdentityProviderBaseRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProviderBase_id"] = id
    }
    return ib0a61c24c37402097c09c6d80120c946edfc11c02e39bf8abbccceafe4660080.NewIdentityProviderBaseRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update identity
func (m *IdentityRequestBuilder) Patch(options *IdentityRequestBuilderPatchOptions)(error) {
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
func (m *IdentityRequestBuilder) UserFlowAttributes()(*i376515807007b3519d82a5a4098cd62c89440575d253961ff5f37dfa7e2c4cf9.UserFlowAttributesRequestBuilder) {
    return i376515807007b3519d82a5a4098cd62c89440575d253961ff5f37dfa7e2c4cf9.NewUserFlowAttributesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserFlowAttributesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.userFlowAttributes.item collection
func (m *IdentityRequestBuilder) UserFlowAttributesById(id string)(*i2507f16e0650c8332777de1431c18c6feb0e9e9b6447d81e77fe14756eb86ba5.IdentityUserFlowAttributeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlowAttribute_id"] = id
    }
    return i2507f16e0650c8332777de1431c18c6feb0e9e9b6447d81e77fe14756eb86ba5.NewIdentityUserFlowAttributeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
