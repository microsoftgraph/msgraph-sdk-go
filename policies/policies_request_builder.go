package policies

import (
    i15c1315dca3d448080a500f31a0343661829915a3d57e0bac2b5ef3d5b702cd7 "github.com/microsoftgraph/msgraph-sdk-go/policies/identitysecuritydefaultsenforcementpolicy"
    i50fbe096cb4bcbe54ce4584948f2fa87f8486532011e2c9024a169ba2a3dacde "github.com/microsoftgraph/msgraph-sdk-go/policies/activitybasedtimeoutpolicies"
    i598aadab7f9532dcd78b5237b03ae7e6ac9390b26cc04cb85038bcb797fcb4e3 "github.com/microsoftgraph/msgraph-sdk-go/policies/featurerolloutpolicies"
    i5eff42eff2580c09a93ee6ffbdd897f38083be436dd15c59a2e4891c29fd0b5e "github.com/microsoftgraph/msgraph-sdk-go/policies/authenticationmethodspolicy"
    i684bc90d26f7d7a7b87627223964f1cc0cab2499ef34d5bdbb89a5ac5cf49209 "github.com/microsoftgraph/msgraph-sdk-go/policies/authenticationflowspolicy"
    i6894c2598f4ac73a5643fa86cbbf110868e7746e8c9ed5da452b0170eebe03b3 "github.com/microsoftgraph/msgraph-sdk-go/policies/homerealmdiscoverypolicies"
    i7cd232f5a1288d29625a9b1dc442455d81df9bedea7cc94eef2584add7856e18 "github.com/microsoftgraph/msgraph-sdk-go/policies/authorizationpolicy"
    i800d2e5efbf3f2562fb570f6abc9bec962927b27f95f6e0025aa92bfe23f13ab "github.com/microsoftgraph/msgraph-sdk-go/policies/claimsmappingpolicies"
    i8cf0127b53d26923e1649ec43c875bc6f0f70ece0a4b9d3cf3377d9e0b4720cf "github.com/microsoftgraph/msgraph-sdk-go/policies/tokenissuancepolicies"
    i94e6a9fd02466104e1f244dfcb22cf331ccf308e2f911adea2fe5a4a87bf5ccc "github.com/microsoftgraph/msgraph-sdk-go/policies/permissiongrantpolicies"
    i966a3041c5833ccf56ce1da16c013fd76d376cfc8b939cc01e633d8b495cb208 "github.com/microsoftgraph/msgraph-sdk-go/policies/conditionalaccesspolicies"
    ic72674a7b815c6dfc1c4fcd42f7ca9ca5d4b8ea0c1a65ea7317905471fa20a0f "github.com/microsoftgraph/msgraph-sdk-go/policies/adminconsentrequestpolicy"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    iebb445e03b5a1725abc955bb63d1192b435aa40b5fee4976b801435c3fc9ec06 "github.com/microsoftgraph/msgraph-sdk-go/policies/tokenlifetimepolicies"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i569717b3ca8408936df0d0a4e253024a8a7b041522350cb7c8977c2c62a9b9d6 "github.com/microsoftgraph/msgraph-sdk-go/policies/tokenlifetimepolicies/item"
    i75b801e9d4a8913922e67e9e75dd091c9280212e449c4d8e76d13ab572b48c86 "github.com/microsoftgraph/msgraph-sdk-go/policies/activitybasedtimeoutpolicies/item"
    i77b5f2108467d08d93443762e93f7967aca61b1916748814eb2ddc49bc0fb0ed "github.com/microsoftgraph/msgraph-sdk-go/policies/tokenissuancepolicies/item"
    i843df2bef70b4bcdc0e002955adffd0d0793c6957377635528f401c17a961d02 "github.com/microsoftgraph/msgraph-sdk-go/policies/conditionalaccesspolicies/item"
    ia9d7ff4c255eb343aa34f45265c76078ea351f2ba2bad54efe9863a838dd38f0 "github.com/microsoftgraph/msgraph-sdk-go/policies/permissiongrantpolicies/item"
    iae405fd168f9b22628a993c1d320fde787ac74895fc32dab558860d3df5e3cf7 "github.com/microsoftgraph/msgraph-sdk-go/policies/homerealmdiscoverypolicies/item"
    ifd481b5a9f37ded8a4d12f7c0db8d1e5cc30754c49b7cbbab66e580c09455018 "github.com/microsoftgraph/msgraph-sdk-go/policies/claimsmappingpolicies/item"
    iff6a6d8b424e8a919281280c87147525819afe04366318faafc92dbee9442123 "github.com/microsoftgraph/msgraph-sdk-go/policies/featurerolloutpolicies/item"
)

// PoliciesRequestBuilder builds and executes requests for operations under \policies
type PoliciesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PoliciesRequestBuilderGetOptions options for Get
type PoliciesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PoliciesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PoliciesRequestBuilderGetQueryParameters get policies
type PoliciesRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PoliciesRequestBuilderPatchOptions options for Patch
type PoliciesRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PolicyRoot;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PoliciesRequestBuilder) ActivityBasedTimeoutPolicies()(*i50fbe096cb4bcbe54ce4584948f2fa87f8486532011e2c9024a169ba2a3dacde.ActivityBasedTimeoutPoliciesRequestBuilder) {
    return i50fbe096cb4bcbe54ce4584948f2fa87f8486532011e2c9024a169ba2a3dacde.NewActivityBasedTimeoutPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivityBasedTimeoutPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.policies.activityBasedTimeoutPolicies.item collection
func (m *PoliciesRequestBuilder) ActivityBasedTimeoutPoliciesById(id string)(*i75b801e9d4a8913922e67e9e75dd091c9280212e449c4d8e76d13ab572b48c86.ActivityBasedTimeoutPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["activityBasedTimeoutPolicy_id"] = id
    }
    return i75b801e9d4a8913922e67e9e75dd091c9280212e449c4d8e76d13ab572b48c86.NewActivityBasedTimeoutPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) AdminConsentRequestPolicy()(*ic72674a7b815c6dfc1c4fcd42f7ca9ca5d4b8ea0c1a65ea7317905471fa20a0f.AdminConsentRequestPolicyRequestBuilder) {
    return ic72674a7b815c6dfc1c4fcd42f7ca9ca5d4b8ea0c1a65ea7317905471fa20a0f.NewAdminConsentRequestPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) AuthenticationFlowsPolicy()(*i684bc90d26f7d7a7b87627223964f1cc0cab2499ef34d5bdbb89a5ac5cf49209.AuthenticationFlowsPolicyRequestBuilder) {
    return i684bc90d26f7d7a7b87627223964f1cc0cab2499ef34d5bdbb89a5ac5cf49209.NewAuthenticationFlowsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) AuthenticationMethodsPolicy()(*i5eff42eff2580c09a93ee6ffbdd897f38083be436dd15c59a2e4891c29fd0b5e.AuthenticationMethodsPolicyRequestBuilder) {
    return i5eff42eff2580c09a93ee6ffbdd897f38083be436dd15c59a2e4891c29fd0b5e.NewAuthenticationMethodsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) AuthorizationPolicy()(*i7cd232f5a1288d29625a9b1dc442455d81df9bedea7cc94eef2584add7856e18.AuthorizationPolicyRequestBuilder) {
    return i7cd232f5a1288d29625a9b1dc442455d81df9bedea7cc94eef2584add7856e18.NewAuthorizationPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) ClaimsMappingPolicies()(*i800d2e5efbf3f2562fb570f6abc9bec962927b27f95f6e0025aa92bfe23f13ab.ClaimsMappingPoliciesRequestBuilder) {
    return i800d2e5efbf3f2562fb570f6abc9bec962927b27f95f6e0025aa92bfe23f13ab.NewClaimsMappingPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClaimsMappingPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.policies.claimsMappingPolicies.item collection
func (m *PoliciesRequestBuilder) ClaimsMappingPoliciesById(id string)(*ifd481b5a9f37ded8a4d12f7c0db8d1e5cc30754c49b7cbbab66e580c09455018.ClaimsMappingPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["claimsMappingPolicy_id"] = id
    }
    return ifd481b5a9f37ded8a4d12f7c0db8d1e5cc30754c49b7cbbab66e580c09455018.NewClaimsMappingPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) ConditionalAccessPolicies()(*i966a3041c5833ccf56ce1da16c013fd76d376cfc8b939cc01e633d8b495cb208.ConditionalAccessPoliciesRequestBuilder) {
    return i966a3041c5833ccf56ce1da16c013fd76d376cfc8b939cc01e633d8b495cb208.NewConditionalAccessPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConditionalAccessPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.policies.conditionalAccessPolicies.item collection
func (m *PoliciesRequestBuilder) ConditionalAccessPoliciesById(id string)(*i843df2bef70b4bcdc0e002955adffd0d0793c6957377635528f401c17a961d02.ConditionalAccessPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessPolicy_id"] = id
    }
    return i843df2bef70b4bcdc0e002955adffd0d0793c6957377635528f401c17a961d02.NewConditionalAccessPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPoliciesRequestBuilderInternal instantiates a new PoliciesRequestBuilder and sets the default values.
func NewPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PoliciesRequestBuilder) {
    m := &PoliciesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPoliciesRequestBuilder instantiates a new PoliciesRequestBuilder and sets the default values.
func NewPoliciesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get policies
func (m *PoliciesRequestBuilder) CreateGetRequestInformation(options *PoliciesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update policies
func (m *PoliciesRequestBuilder) CreatePatchRequestInformation(options *PoliciesRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PoliciesRequestBuilder) FeatureRolloutPolicies()(*i598aadab7f9532dcd78b5237b03ae7e6ac9390b26cc04cb85038bcb797fcb4e3.FeatureRolloutPoliciesRequestBuilder) {
    return i598aadab7f9532dcd78b5237b03ae7e6ac9390b26cc04cb85038bcb797fcb4e3.NewFeatureRolloutPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FeatureRolloutPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.policies.featureRolloutPolicies.item collection
func (m *PoliciesRequestBuilder) FeatureRolloutPoliciesById(id string)(*iff6a6d8b424e8a919281280c87147525819afe04366318faafc92dbee9442123.FeatureRolloutPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["featureRolloutPolicy_id"] = id
    }
    return iff6a6d8b424e8a919281280c87147525819afe04366318faafc92dbee9442123.NewFeatureRolloutPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get policies
func (m *PoliciesRequestBuilder) Get(options *PoliciesRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PolicyRoot, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPolicyRoot() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PolicyRoot), nil
}
func (m *PoliciesRequestBuilder) HomeRealmDiscoveryPolicies()(*i6894c2598f4ac73a5643fa86cbbf110868e7746e8c9ed5da452b0170eebe03b3.HomeRealmDiscoveryPoliciesRequestBuilder) {
    return i6894c2598f4ac73a5643fa86cbbf110868e7746e8c9ed5da452b0170eebe03b3.NewHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HomeRealmDiscoveryPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.policies.homeRealmDiscoveryPolicies.item collection
func (m *PoliciesRequestBuilder) HomeRealmDiscoveryPoliciesById(id string)(*iae405fd168f9b22628a993c1d320fde787ac74895fc32dab558860d3df5e3cf7.HomeRealmDiscoveryPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["homeRealmDiscoveryPolicy_id"] = id
    }
    return iae405fd168f9b22628a993c1d320fde787ac74895fc32dab558860d3df5e3cf7.NewHomeRealmDiscoveryPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) IdentitySecurityDefaultsEnforcementPolicy()(*i15c1315dca3d448080a500f31a0343661829915a3d57e0bac2b5ef3d5b702cd7.IdentitySecurityDefaultsEnforcementPolicyRequestBuilder) {
    return i15c1315dca3d448080a500f31a0343661829915a3d57e0bac2b5ef3d5b702cd7.NewIdentitySecurityDefaultsEnforcementPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update policies
func (m *PoliciesRequestBuilder) Patch(options *PoliciesRequestBuilderPatchOptions)(error) {
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
func (m *PoliciesRequestBuilder) PermissionGrantPolicies()(*i94e6a9fd02466104e1f244dfcb22cf331ccf308e2f911adea2fe5a4a87bf5ccc.PermissionGrantPoliciesRequestBuilder) {
    return i94e6a9fd02466104e1f244dfcb22cf331ccf308e2f911adea2fe5a4a87bf5ccc.NewPermissionGrantPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.policies.permissionGrantPolicies.item collection
func (m *PoliciesRequestBuilder) PermissionGrantPoliciesById(id string)(*ia9d7ff4c255eb343aa34f45265c76078ea351f2ba2bad54efe9863a838dd38f0.PermissionGrantPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permissionGrantPolicy_id"] = id
    }
    return ia9d7ff4c255eb343aa34f45265c76078ea351f2ba2bad54efe9863a838dd38f0.NewPermissionGrantPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) TokenIssuancePolicies()(*i8cf0127b53d26923e1649ec43c875bc6f0f70ece0a4b9d3cf3377d9e0b4720cf.TokenIssuancePoliciesRequestBuilder) {
    return i8cf0127b53d26923e1649ec43c875bc6f0f70ece0a4b9d3cf3377d9e0b4720cf.NewTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenIssuancePoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.policies.tokenIssuancePolicies.item collection
func (m *PoliciesRequestBuilder) TokenIssuancePoliciesById(id string)(*i77b5f2108467d08d93443762e93f7967aca61b1916748814eb2ddc49bc0fb0ed.TokenIssuancePolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenIssuancePolicy_id"] = id
    }
    return i77b5f2108467d08d93443762e93f7967aca61b1916748814eb2ddc49bc0fb0ed.NewTokenIssuancePolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) TokenLifetimePolicies()(*iebb445e03b5a1725abc955bb63d1192b435aa40b5fee4976b801435c3fc9ec06.TokenLifetimePoliciesRequestBuilder) {
    return iebb445e03b5a1725abc955bb63d1192b435aa40b5fee4976b801435c3fc9ec06.NewTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenLifetimePoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.policies.tokenLifetimePolicies.item collection
func (m *PoliciesRequestBuilder) TokenLifetimePoliciesById(id string)(*i569717b3ca8408936df0d0a4e253024a8a7b041522350cb7c8977c2c62a9b9d6.TokenLifetimePolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenLifetimePolicy_id"] = id
    }
    return i569717b3ca8408936df0d0a4e253024a8a7b041522350cb7c8977c2c62a9b9d6.NewTokenLifetimePolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
