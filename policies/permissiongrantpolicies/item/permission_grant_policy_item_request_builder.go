package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i93d6ed019a5688a4e8e0075b0389cf2e1ef569a5a88c931a9822c1aad16943ed "github.com/microsoftgraph/msgraph-sdk-go/policies/permissiongrantpolicies/item/includes"
    icde6cf197f013af55fc32ae765cd6718c8283b86b87581a48e55e766347fa252 "github.com/microsoftgraph/msgraph-sdk-go/policies/permissiongrantpolicies/item/excludes"
    i02a33916c6ad92e2af8e1a1c3c5bb0b0e4ed56d4de269096ae95977b3c84a597 "github.com/microsoftgraph/msgraph-sdk-go/policies/permissiongrantpolicies/item/excludes/item"
    ia8e0f04a057817a8b64c99bfcfceca507e3ffc64a0ce4ab2cfcd9fc69349259f "github.com/microsoftgraph/msgraph-sdk-go/policies/permissiongrantpolicies/item/includes/item"
)

// PermissionGrantPolicyItemRequestBuilder provides operations to manage the permissionGrantPolicies property of the microsoft.graph.policyRoot entity.
type PermissionGrantPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PermissionGrantPolicyItemRequestBuilderDeleteOptions options for Delete
type PermissionGrantPolicyItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PermissionGrantPolicyItemRequestBuilderGetOptions options for Get
type PermissionGrantPolicyItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PermissionGrantPolicyItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PermissionGrantPolicyItemRequestBuilderGetQueryParameters the policy that specifies the conditions under which consent can be granted.
type PermissionGrantPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PermissionGrantPolicyItemRequestBuilderPatchOptions options for Patch
type PermissionGrantPolicyItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PermissionGrantPolicyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewPermissionGrantPolicyItemRequestBuilderInternal instantiates a new PermissionGrantPolicyItemRequestBuilder and sets the default values.
func NewPermissionGrantPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PermissionGrantPolicyItemRequestBuilder) {
    m := &PermissionGrantPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/permissionGrantPolicies/{permissionGrantPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPermissionGrantPolicyItemRequestBuilder instantiates a new PermissionGrantPolicyItemRequestBuilder and sets the default values.
func NewPermissionGrantPolicyItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PermissionGrantPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionGrantPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property permissionGrantPolicies for policies
func (m *PermissionGrantPolicyItemRequestBuilder) CreateDeleteRequestInformation(options *PermissionGrantPolicyItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the policy that specifies the conditions under which consent can be granted.
func (m *PermissionGrantPolicyItemRequestBuilder) CreateGetRequestInformation(options *PermissionGrantPolicyItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property permissionGrantPolicies in policies
func (m *PermissionGrantPolicyItemRequestBuilder) CreatePatchRequestInformation(options *PermissionGrantPolicyItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property permissionGrantPolicies for policies
func (m *PermissionGrantPolicyItemRequestBuilder) Delete(options *PermissionGrantPolicyItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *PermissionGrantPolicyItemRequestBuilder) Excludes()(*icde6cf197f013af55fc32ae765cd6718c8283b86b87581a48e55e766347fa252.ExcludesRequestBuilder) {
    return icde6cf197f013af55fc32ae765cd6718c8283b86b87581a48e55e766347fa252.NewExcludesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExcludesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.policies.permissionGrantPolicies.item.excludes.item collection
func (m *PermissionGrantPolicyItemRequestBuilder) ExcludesById(id string)(*i02a33916c6ad92e2af8e1a1c3c5bb0b0e4ed56d4de269096ae95977b3c84a597.PermissionGrantConditionSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permissionGrantConditionSet_id"] = id
    }
    return i02a33916c6ad92e2af8e1a1c3c5bb0b0e4ed56d4de269096ae95977b3c84a597.NewPermissionGrantConditionSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the policy that specifies the conditions under which consent can be granted.
func (m *PermissionGrantPolicyItemRequestBuilder) Get(options *PermissionGrantPolicyItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PermissionGrantPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreatePermissionGrantPolicyFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PermissionGrantPolicyable), nil
}
func (m *PermissionGrantPolicyItemRequestBuilder) Includes()(*i93d6ed019a5688a4e8e0075b0389cf2e1ef569a5a88c931a9822c1aad16943ed.IncludesRequestBuilder) {
    return i93d6ed019a5688a4e8e0075b0389cf2e1ef569a5a88c931a9822c1aad16943ed.NewIncludesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncludesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.policies.permissionGrantPolicies.item.includes.item collection
func (m *PermissionGrantPolicyItemRequestBuilder) IncludesById(id string)(*ia8e0f04a057817a8b64c99bfcfceca507e3ffc64a0ce4ab2cfcd9fc69349259f.PermissionGrantConditionSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permissionGrantConditionSet_id"] = id
    }
    return ia8e0f04a057817a8b64c99bfcfceca507e3ffc64a0ce4ab2cfcd9fc69349259f.NewPermissionGrantConditionSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property permissionGrantPolicies in policies
func (m *PermissionGrantPolicyItemRequestBuilder) Patch(options *PermissionGrantPolicyItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
