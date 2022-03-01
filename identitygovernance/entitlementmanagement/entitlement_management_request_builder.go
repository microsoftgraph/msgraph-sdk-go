package entitlementmanagement

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i069e71eb6a0d6756f4420f4f811017c7470620ad5174a1e3aab2db0d8cb430c2 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/catalogs"
    i3e9bbff4ac7b59c7638e82a3457169acb55d826f7187ac6e78fbe6ec1bb80d7d "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/assignments"
    i441d3614c41d6a7eafb923dbc4975591570abe0ca8005d661c628e49592a07f8 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/settings"
    i458b5d51a7c8ed196efc57bee2af7eeeea44a7093ec6fe9b40b0d87560256cca "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/accesspackages"
    i4a6e5bbedb985fd4e83bec609cf8d21a776d26a1c6807eac6ae0d88110984ab9 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentapprovals"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i890fb4653f4dff2b9ab5a72705a40947b79d9bb64793d05773197622a996fc61 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations"
    i8d2bfa033b27ef0ce19de47e723f8c3475251f8917d2cbf768fc4e424505f65e "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/assignmentrequests"
    i3a731cbed9792cd58eb16385c435f4b1af4fafbdc6a965b8534ab7ff7df19e39 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item"
    i558fcddcc0d01116d05fce610754153bd32852762580e628e25c34b7d8fffa4a "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/assignments/item"
    i76acc551f9c3cb51943981936cb0b2aa4089fd504de3bf817176dbbe448d3fb7 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/assignmentrequests/item"
    i98d02a4533e64ff168c5023553f526cbe65a72cacee2308e1a90cb06055c2a93 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item"
    i9f1127e219619b20720c5350652063980d97d3210fea76590720646f27ba45bb "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/catalogs/item"
    ib6ba4a2debf33ace4ef4d7d21e119e5c8a37451c871b2845657ae921beea0814 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentapprovals/item"
)

// EntitlementManagementRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement
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
// EntitlementManagementRequestBuilderGetQueryParameters get entitlementManagement from identityGovernance
type EntitlementManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EntitlementManagementRequestBuilderPatchOptions options for Patch
type EntitlementManagementRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EntitlementManagement;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentApprovals()(*i4a6e5bbedb985fd4e83bec609cf8d21a776d26a1c6807eac6ae0d88110984ab9.AccessPackageAssignmentApprovalsRequestBuilder) {
    return i4a6e5bbedb985fd4e83bec609cf8d21a776d26a1c6807eac6ae0d88110984ab9.NewAccessPackageAssignmentApprovalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentApprovalsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentApprovals.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentApprovalsById(id string)(*ib6ba4a2debf33ace4ef4d7d21e119e5c8a37451c871b2845657ae921beea0814.ApprovalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approval_id"] = id
    }
    return ib6ba4a2debf33ace4ef4d7d21e119e5c8a37451c871b2845657ae921beea0814.NewApprovalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) AccessPackages()(*i458b5d51a7c8ed196efc57bee2af7eeeea44a7093ec6fe9b40b0d87560256cca.AccessPackagesRequestBuilder) {
    return i458b5d51a7c8ed196efc57bee2af7eeeea44a7093ec6fe9b40b0d87560256cca.NewAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackagesById(id string)(*i98d02a4533e64ff168c5023553f526cbe65a72cacee2308e1a90cb06055c2a93.AccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage_id"] = id
    }
    return i98d02a4533e64ff168c5023553f526cbe65a72cacee2308e1a90cb06055c2a93.NewAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) AssignmentRequests()(*i8d2bfa033b27ef0ce19de47e723f8c3475251f8917d2cbf768fc4e424505f65e.AssignmentRequestsRequestBuilder) {
    return i8d2bfa033b27ef0ce19de47e723f8c3475251f8917d2cbf768fc4e424505f65e.NewAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.assignmentRequests.item collection
func (m *EntitlementManagementRequestBuilder) AssignmentRequestsById(id string)(*i76acc551f9c3cb51943981936cb0b2aa4089fd504de3bf817176dbbe448d3fb7.AccessPackageAssignmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentRequest_id"] = id
    }
    return i76acc551f9c3cb51943981936cb0b2aa4089fd504de3bf817176dbbe448d3fb7.NewAccessPackageAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) Assignments()(*i3e9bbff4ac7b59c7638e82a3457169acb55d826f7187ac6e78fbe6ec1bb80d7d.AssignmentsRequestBuilder) {
    return i3e9bbff4ac7b59c7638e82a3457169acb55d826f7187ac6e78fbe6ec1bb80d7d.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.assignments.item collection
func (m *EntitlementManagementRequestBuilder) AssignmentsById(id string)(*i558fcddcc0d01116d05fce610754153bd32852762580e628e25c34b7d8fffa4a.AccessPackageAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignment_id"] = id
    }
    return i558fcddcc0d01116d05fce610754153bd32852762580e628e25c34b7d8fffa4a.NewAccessPackageAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) Catalogs()(*i069e71eb6a0d6756f4420f4f811017c7470620ad5174a1e3aab2db0d8cb430c2.CatalogsRequestBuilder) {
    return i069e71eb6a0d6756f4420f4f811017c7470620ad5174a1e3aab2db0d8cb430c2.NewCatalogsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CatalogsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.catalogs.item collection
func (m *EntitlementManagementRequestBuilder) CatalogsById(id string)(*i9f1127e219619b20720c5350652063980d97d3210fea76590720646f27ba45bb.AccessPackageCatalogItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageCatalog_id"] = id
    }
    return i9f1127e219619b20720c5350652063980d97d3210fea76590720646f27ba45bb.NewAccessPackageCatalogItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) ConnectedOrganizations()(*i890fb4653f4dff2b9ab5a72705a40947b79d9bb64793d05773197622a996fc61.ConnectedOrganizationsRequestBuilder) {
    return i890fb4653f4dff2b9ab5a72705a40947b79d9bb64793d05773197622a996fc61.NewConnectedOrganizationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectedOrganizationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.connectedOrganizations.item collection
func (m *EntitlementManagementRequestBuilder) ConnectedOrganizationsById(id string)(*i3a731cbed9792cd58eb16385c435f4b1af4fafbdc6a965b8534ab7ff7df19e39.ConnectedOrganizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectedOrganization_id"] = id
    }
    return i3a731cbed9792cd58eb16385c435f4b1af4fafbdc6a965b8534ab7ff7df19e39.NewConnectedOrganizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEntitlementManagementRequestBuilderInternal instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EntitlementManagementRequestBuilder) {
    m := &EntitlementManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property entitlementManagement for identityGovernance
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
// CreateGetRequestInformation get entitlementManagement from identityGovernance
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
// CreatePatchRequestInformation update the navigation property entitlementManagement in identityGovernance
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
// Delete delete navigation property entitlementManagement for identityGovernance
func (m *EntitlementManagementRequestBuilder) Delete(options *EntitlementManagementRequestBuilderDeleteOptions)(error) {
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
// Get get entitlementManagement from identityGovernance
func (m *EntitlementManagementRequestBuilder) Get(options *EntitlementManagementRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EntitlementManagement, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEntitlementManagement() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EntitlementManagement), nil
}
// Patch update the navigation property entitlementManagement in identityGovernance
func (m *EntitlementManagementRequestBuilder) Patch(options *EntitlementManagementRequestBuilderPatchOptions)(error) {
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
func (m *EntitlementManagementRequestBuilder) Settings()(*i441d3614c41d6a7eafb923dbc4975591570abe0ca8005d661c628e49592a07f8.SettingsRequestBuilder) {
    return i441d3614c41d6a7eafb923dbc4975591570abe0ca8005d661c628e49592a07f8.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
