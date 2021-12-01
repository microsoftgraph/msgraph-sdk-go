package entitlementmanagement

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i235730479c1a6cf49ab5074a2df5fee811eeaf00c02b8796d08bae7d0b1ac121 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roledefinitions"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    iab5b1a7cf751122475647c02429db57a9207d8d85016b356d65e7aeb3a05d352 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignments"
    i02dcd7ccbb395e0de15ed121c8c0a0de1ce9b7389ee22ee420b13b1115a5b1d4 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignments/item"
    i450e6fa82c54d6f54b7abb6a155c2b0b0cb9aced086261fdf11e3d7d814c6286 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roledefinitions/item"
)

// EntitlementManagementRequestBuilder builds and executes requests for operations under \roleManagement\entitlementManagement
type EntitlementManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EntitlementManagementRequestBuilderDeleteOptions options for Delete
type EntitlementManagementRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EntitlementManagementRequestBuilderGetOptions options for Get
type EntitlementManagementRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EntitlementManagementRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EntitlementManagementRequestBuilderGetQueryParameters the RbacApplication for Entitlement Management
type EntitlementManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EntitlementManagementRequestBuilderPatchOptions options for Patch
type EntitlementManagementRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.RbacApplication;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewEntitlementManagementRequestBuilderInternal instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EntitlementManagementRequestBuilder) {
    m := &EntitlementManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEntitlementManagementRequestBuilder instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementManagementRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EntitlementManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the RbacApplication for Entitlement Management
func (m *EntitlementManagementRequestBuilder) CreateDeleteRequestInformation(options *EntitlementManagementRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the RbacApplication for Entitlement Management
func (m *EntitlementManagementRequestBuilder) CreateGetRequestInformation(options *EntitlementManagementRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the RbacApplication for Entitlement Management
func (m *EntitlementManagementRequestBuilder) CreatePatchRequestInformation(options *EntitlementManagementRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the RbacApplication for Entitlement Management
func (m *EntitlementManagementRequestBuilder) Delete(options *EntitlementManagementRequestBuilderDeleteOptions)(error) {
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
// Get the RbacApplication for Entitlement Management
func (m *EntitlementManagementRequestBuilder) Get(options *EntitlementManagementRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.RbacApplication, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewRbacApplication() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.RbacApplication), nil
}
// Patch the RbacApplication for Entitlement Management
func (m *EntitlementManagementRequestBuilder) Patch(options *EntitlementManagementRequestBuilderPatchOptions)(error) {
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
func (m *EntitlementManagementRequestBuilder) RoleAssignments()(*iab5b1a7cf751122475647c02429db57a9207d8d85016b356d65e7aeb3a05d352.RoleAssignmentsRequestBuilder) {
    return iab5b1a7cf751122475647c02429db57a9207d8d85016b356d65e7aeb3a05d352.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.roleManagement.entitlementManagement.roleAssignments.item collection
func (m *EntitlementManagementRequestBuilder) RoleAssignmentsById(id string)(*i02dcd7ccbb395e0de15ed121c8c0a0de1ce9b7389ee22ee420b13b1115a5b1d4.UnifiedRoleAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignment_id"] = id
    }
    return i02dcd7ccbb395e0de15ed121c8c0a0de1ce9b7389ee22ee420b13b1115a5b1d4.NewUnifiedRoleAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) RoleDefinitions()(*i235730479c1a6cf49ab5074a2df5fee811eeaf00c02b8796d08bae7d0b1ac121.RoleDefinitionsRequestBuilder) {
    return i235730479c1a6cf49ab5074a2df5fee811eeaf00c02b8796d08bae7d0b1ac121.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.roleManagement.entitlementManagement.roleDefinitions.item collection
func (m *EntitlementManagementRequestBuilder) RoleDefinitionsById(id string)(*i450e6fa82c54d6f54b7abb6a155c2b0b0cb9aced086261fdf11e3d7d814c6286.UnifiedRoleDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleDefinition_id"] = id
    }
    return i450e6fa82c54d6f54b7abb6a155c2b0b0cb9aced086261fdf11e3d7d814c6286.NewUnifiedRoleDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
