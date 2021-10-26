package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i93d6ed019a5688a4e8e0075b0389cf2e1ef569a5a88c931a9822c1aad16943ed "github.com/microsoftgraph/msgraph-sdk-go/policies/permissiongrantpolicies/item/includes"
    icde6cf197f013af55fc32ae765cd6718c8283b86b87581a48e55e766347fa252 "github.com/microsoftgraph/msgraph-sdk-go/policies/permissiongrantpolicies/item/excludes"
    i02a33916c6ad92e2af8e1a1c3c5bb0b0e4ed56d4de269096ae95977b3c84a597 "github.com/microsoftgraph/msgraph-sdk-go/policies/permissiongrantpolicies/item/excludes/item"
    ia8e0f04a057817a8b64c99bfcfceca507e3ffc64a0ce4ab2cfcd9fc69349259f "github.com/microsoftgraph/msgraph-sdk-go/policies/permissiongrantpolicies/item/includes/item"
)

// Builds and executes requests for operations under \policies\permissionGrantPolicies\{permissionGrantPolicy-id}
type PermissionGrantPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The policy that specifies the conditions under which consent can be granted.
type PermissionGrantPolicyRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new PermissionGrantPolicyRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPermissionGrantPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PermissionGrantPolicyRequestBuilder) {
    m := &PermissionGrantPolicyRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/policies/permissionGrantPolicies/{permissionGrantPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new PermissionGrantPolicyRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPermissionGrantPolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PermissionGrantPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionGrantPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// The policy that specifies the conditions under which consent can be granted.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *PermissionGrantPolicyRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The policy that specifies the conditions under which consent can be granted.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *PermissionGrantPolicyRequestBuilder) CreateGetRequestInformation(q func (value *PermissionGrantPolicyRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(PermissionGrantPolicyRequestBuilderGetQueryParameters)
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
// The policy that specifies the conditions under which consent can be granted.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *PermissionGrantPolicyRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PermissionGrantPolicy, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The policy that specifies the conditions under which consent can be granted.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *PermissionGrantPolicyRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PermissionGrantPolicyRequestBuilder) Excludes()(*icde6cf197f013af55fc32ae765cd6718c8283b86b87581a48e55e766347fa252.ExcludesRequestBuilder) {
    return icde6cf197f013af55fc32ae765cd6718c8283b86b87581a48e55e766347fa252.NewExcludesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.policies.permissionGrantPolicies.item.excludes.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PermissionGrantPolicyRequestBuilder) ExcludesById(id string)(*i02a33916c6ad92e2af8e1a1c3c5bb0b0e4ed56d4de269096ae95977b3c84a597.PermissionGrantConditionSetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permissionGrantConditionSet_id"] = id
    }
    return i02a33916c6ad92e2af8e1a1c3c5bb0b0e4ed56d4de269096ae95977b3c84a597.NewPermissionGrantConditionSetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The policy that specifies the conditions under which consent can be granted.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *PermissionGrantPolicyRequestBuilder) Get(q func (value *PermissionGrantPolicyRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PermissionGrantPolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPermissionGrantPolicy() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PermissionGrantPolicy), nil
}
func (m *PermissionGrantPolicyRequestBuilder) Includes()(*i93d6ed019a5688a4e8e0075b0389cf2e1ef569a5a88c931a9822c1aad16943ed.IncludesRequestBuilder) {
    return i93d6ed019a5688a4e8e0075b0389cf2e1ef569a5a88c931a9822c1aad16943ed.NewIncludesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.policies.permissionGrantPolicies.item.includes.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PermissionGrantPolicyRequestBuilder) IncludesById(id string)(*ia8e0f04a057817a8b64c99bfcfceca507e3ffc64a0ce4ab2cfcd9fc69349259f.PermissionGrantConditionSetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permissionGrantConditionSet_id"] = id
    }
    return ia8e0f04a057817a8b64c99bfcfceca507e3ffc64a0ce4ab2cfcd9fc69349259f.NewPermissionGrantConditionSetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The policy that specifies the conditions under which consent can be granted.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *PermissionGrantPolicyRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PermissionGrantPolicy, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
