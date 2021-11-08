package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i776e11034305b48e5290b2e53dd2f7575359c486273d18f2353ff65faa9063b1 "github.com/microsoftgraph/msgraph-sdk-go/directory/administrativeunits/item/scopedrolemembers"
    i94253098475db605e990260e13febd658ee5c0709ceaf573e3f903be950c52b6 "github.com/microsoftgraph/msgraph-sdk-go/directory/administrativeunits/item/extensions"
    ia454d456d99bb0ead0e1bfb50c53f817055b0e68ca4356296e7f406ee25d3d77 "github.com/microsoftgraph/msgraph-sdk-go/directory/administrativeunits/item/members"
    ia9df9f379ea5fc92e0383f29dfd050c3b88591b8f475f586b856fc130a4fd053 "github.com/microsoftgraph/msgraph-sdk-go/directory/administrativeunits/item/scopedrolemembers/item"
    if71e90a642a2c3f3e90bf7d0f3cbe5d719d6ca30c2312110465fc6e587f8ec38 "github.com/microsoftgraph/msgraph-sdk-go/directory/administrativeunits/item/extensions/item"
)

// Builds and executes requests for operations under \directory\administrativeUnits\{administrativeUnit-id}
type AdministrativeUnitRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type AdministrativeUnitRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type AdministrativeUnitRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AdministrativeUnitRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Conceptual container for user and group directory objects.
type AdministrativeUnitRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type AdministrativeUnitRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AdministrativeUnit;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new AdministrativeUnitRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAdministrativeUnitRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdministrativeUnitRequestBuilder) {
    m := &AdministrativeUnitRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/administrativeUnits/{administrativeUnit_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new AdministrativeUnitRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAdministrativeUnitRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdministrativeUnitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdministrativeUnitRequestBuilderInternal(urlParams, requestAdapter)
}
// Conceptual container for user and group directory objects.
// Parameters:
//  - options : Options for the request
func (m *AdministrativeUnitRequestBuilder) CreateDeleteRequestInformation(options *AdministrativeUnitRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Conceptual container for user and group directory objects.
// Parameters:
//  - options : Options for the request
func (m *AdministrativeUnitRequestBuilder) CreateGetRequestInformation(options *AdministrativeUnitRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Conceptual container for user and group directory objects.
// Parameters:
//  - options : Options for the request
func (m *AdministrativeUnitRequestBuilder) CreatePatchRequestInformation(options *AdministrativeUnitRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Conceptual container for user and group directory objects.
// Parameters:
//  - options : Options for the request
func (m *AdministrativeUnitRequestBuilder) Delete(options *AdministrativeUnitRequestBuilderDeleteOptions)(error) {
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
func (m *AdministrativeUnitRequestBuilder) Extensions()(*i94253098475db605e990260e13febd658ee5c0709ceaf573e3f903be950c52b6.ExtensionsRequestBuilder) {
    return i94253098475db605e990260e13febd658ee5c0709ceaf573e3f903be950c52b6.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.directory.administrativeUnits.item.extensions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AdministrativeUnitRequestBuilder) ExtensionsById(id string)(*if71e90a642a2c3f3e90bf7d0f3cbe5d719d6ca30c2312110465fc6e587f8ec38.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return if71e90a642a2c3f3e90bf7d0f3cbe5d719d6ca30c2312110465fc6e587f8ec38.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Conceptual container for user and group directory objects.
// Parameters:
//  - options : Options for the request
func (m *AdministrativeUnitRequestBuilder) Get(options *AdministrativeUnitRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AdministrativeUnit, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewAdministrativeUnit() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AdministrativeUnit), nil
}
func (m *AdministrativeUnitRequestBuilder) Members()(*ia454d456d99bb0ead0e1bfb50c53f817055b0e68ca4356296e7f406ee25d3d77.MembersRequestBuilder) {
    return ia454d456d99bb0ead0e1bfb50c53f817055b0e68ca4356296e7f406ee25d3d77.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Conceptual container for user and group directory objects.
// Parameters:
//  - options : Options for the request
func (m *AdministrativeUnitRequestBuilder) Patch(options *AdministrativeUnitRequestBuilderPatchOptions)(error) {
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
func (m *AdministrativeUnitRequestBuilder) ScopedRoleMembers()(*i776e11034305b48e5290b2e53dd2f7575359c486273d18f2353ff65faa9063b1.ScopedRoleMembersRequestBuilder) {
    return i776e11034305b48e5290b2e53dd2f7575359c486273d18f2353ff65faa9063b1.NewScopedRoleMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.directory.administrativeUnits.item.scopedRoleMembers.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AdministrativeUnitRequestBuilder) ScopedRoleMembersById(id string)(*ia9df9f379ea5fc92e0383f29dfd050c3b88591b8f475f586b856fc130a4fd053.ScopedRoleMembershipRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership_id"] = id
    }
    return ia9df9f379ea5fc92e0383f29dfd050c3b88591b8f475f586b856fc130a4fd053.NewScopedRoleMembershipRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
