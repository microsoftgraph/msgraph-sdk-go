package entitlementmanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i235730479c1a6cf49ab5074a2df5fee811eeaf00c02b8796d08bae7d0b1ac121 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roledefinitions"
    i3a967725e663f75605977fcee02eb4b2642ee31d0648ec4e2ed0f800882b32bc "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedules"
    i6105172c68f51bb2e2927309c74148ed2bf9eaebff5fecc8144191485bb58a1c "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests"
    i75e0187968a6021e302f1aa7a3eb2a876441acad1f351cff3d51d6d0c4cd52a4 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityscheduleinstances"
    i7fdde2c600ba1f94ea8ccc2b16da2d73b29fabb1bdfe5db21564f182303fa289 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignmentscheduleinstances"
    ia00666d9c15220a15dcb7fcfb7faecbb7b75cddefc298ac4a352e53501419482 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedules"
    iab5b1a7cf751122475647c02429db57a9207d8d85016b356d65e7aeb3a05d352 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignments"
    iffc6a561bcf2356c17f04c85d6602846a02a99b2f78d88074e35ba71edbddf9f "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests"
    i02dcd7ccbb395e0de15ed121c8c0a0de1ce9b7389ee22ee420b13b1115a5b1d4 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignments/item"
    i13c7c694392be2fc5947090399fe93a9fdab4dedaabe0939405a5313dd75bdae "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item"
    i1e88f04b4de5bc9a33401e3c9b33e94251a138355ac41b0e707b0340e1fb8409 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignmentscheduleinstances/item"
    i3b9e25b9426e31e4bcd53952d1b0865afd529e64b7d980cc74d8ef468d55a6ca "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item"
    i450e6fa82c54d6f54b7abb6a155c2b0b0cb9aced086261fdf11e3d7d814c6286 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roledefinitions/item"
    i7c97a82666b511908e752f14b6d01e35e5f2565606757e69c0f6835e4ec31b4c "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedules/item"
    i9737b224e822cfafa077e03d8ff571fad57ee6c6761c45f56482ea39d6af9c92 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedules/item"
    i987a32583667f0a5ef515c4455c9a8fa39e772a7676688f15e1cc180c13ae1c9 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityscheduleinstances/item"
)

// EntitlementManagementRequestBuilder provides operations to manage the entitlementManagement property of the microsoft.graph.roleManagement entity.
type EntitlementManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementRequestBuilderDeleteOptions options for Delete
type EntitlementManagementRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// EntitlementManagementRequestBuilderGetOptions options for Get
type EntitlementManagementRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// EntitlementManagementRequestBuilderGetQueryParameters container for roles and assignments for entitlement management resources.
type EntitlementManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementRequestBuilderPatchOptions options for Patch
type EntitlementManagementRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RbacApplicationable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewEntitlementManagementRequestBuilderInternal instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementRequestBuilder) {
    m := &EntitlementManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEntitlementManagementRequestBuilder instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property entitlementManagement for roleManagement
func (m *EntitlementManagementRequestBuilder) CreateDeleteRequestInformation(options *EntitlementManagementRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation container for roles and assignments for entitlement management resources.
func (m *EntitlementManagementRequestBuilder) CreateGetRequestInformation(options *EntitlementManagementRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property entitlementManagement in roleManagement
func (m *EntitlementManagementRequestBuilder) CreatePatchRequestInformation(options *EntitlementManagementRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property entitlementManagement for roleManagement
func (m *EntitlementManagementRequestBuilder) Delete(options *EntitlementManagementRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get container for roles and assignments for entitlement management resources.
func (m *EntitlementManagementRequestBuilder) Get(options *EntitlementManagementRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RbacApplicationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRbacApplicationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RbacApplicationable), nil
}
// Patch update the navigation property entitlementManagement in roleManagement
func (m *EntitlementManagementRequestBuilder) Patch(options *EntitlementManagementRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// RoleAssignments the roleAssignments property
func (m *EntitlementManagementRequestBuilder) RoleAssignments()(*iab5b1a7cf751122475647c02429db57a9207d8d85016b356d65e7aeb3a05d352.RoleAssignmentsRequestBuilder) {
    return iab5b1a7cf751122475647c02429db57a9207d8d85016b356d65e7aeb3a05d352.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.roleManagement.entitlementManagement.roleAssignments.item collection
func (m *EntitlementManagementRequestBuilder) RoleAssignmentsById(id string)(*i02dcd7ccbb395e0de15ed121c8c0a0de1ce9b7389ee22ee420b13b1115a5b1d4.UnifiedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignment%2Did"] = id
    }
    return i02dcd7ccbb395e0de15ed121c8c0a0de1ce9b7389ee22ee420b13b1115a5b1d4.NewUnifiedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentScheduleInstances the roleAssignmentScheduleInstances property
func (m *EntitlementManagementRequestBuilder) RoleAssignmentScheduleInstances()(*i7fdde2c600ba1f94ea8ccc2b16da2d73b29fabb1bdfe5db21564f182303fa289.RoleAssignmentScheduleInstancesRequestBuilder) {
    return i7fdde2c600ba1f94ea8ccc2b16da2d73b29fabb1bdfe5db21564f182303fa289.NewRoleAssignmentScheduleInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentScheduleInstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.roleManagement.entitlementManagement.roleAssignmentScheduleInstances.item collection
func (m *EntitlementManagementRequestBuilder) RoleAssignmentScheduleInstancesById(id string)(*i1e88f04b4de5bc9a33401e3c9b33e94251a138355ac41b0e707b0340e1fb8409.UnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentScheduleInstance%2Did"] = id
    }
    return i1e88f04b4de5bc9a33401e3c9b33e94251a138355ac41b0e707b0340e1fb8409.NewUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentScheduleRequests the roleAssignmentScheduleRequests property
func (m *EntitlementManagementRequestBuilder) RoleAssignmentScheduleRequests()(*iffc6a561bcf2356c17f04c85d6602846a02a99b2f78d88074e35ba71edbddf9f.RoleAssignmentScheduleRequestsRequestBuilder) {
    return iffc6a561bcf2356c17f04c85d6602846a02a99b2f78d88074e35ba71edbddf9f.NewRoleAssignmentScheduleRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentScheduleRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.roleManagement.entitlementManagement.roleAssignmentScheduleRequests.item collection
func (m *EntitlementManagementRequestBuilder) RoleAssignmentScheduleRequestsById(id string)(*i13c7c694392be2fc5947090399fe93a9fdab4dedaabe0939405a5313dd75bdae.UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentScheduleRequest%2Did"] = id
    }
    return i13c7c694392be2fc5947090399fe93a9fdab4dedaabe0939405a5313dd75bdae.NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentSchedules the roleAssignmentSchedules property
func (m *EntitlementManagementRequestBuilder) RoleAssignmentSchedules()(*ia00666d9c15220a15dcb7fcfb7faecbb7b75cddefc298ac4a352e53501419482.RoleAssignmentSchedulesRequestBuilder) {
    return ia00666d9c15220a15dcb7fcfb7faecbb7b75cddefc298ac4a352e53501419482.NewRoleAssignmentSchedulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentSchedulesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.roleManagement.entitlementManagement.roleAssignmentSchedules.item collection
func (m *EntitlementManagementRequestBuilder) RoleAssignmentSchedulesById(id string)(*i7c97a82666b511908e752f14b6d01e35e5f2565606757e69c0f6835e4ec31b4c.UnifiedRoleAssignmentScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentSchedule%2Did"] = id
    }
    return i7c97a82666b511908e752f14b6d01e35e5f2565606757e69c0f6835e4ec31b4c.NewUnifiedRoleAssignmentScheduleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleDefinitions the roleDefinitions property
func (m *EntitlementManagementRequestBuilder) RoleDefinitions()(*i235730479c1a6cf49ab5074a2df5fee811eeaf00c02b8796d08bae7d0b1ac121.RoleDefinitionsRequestBuilder) {
    return i235730479c1a6cf49ab5074a2df5fee811eeaf00c02b8796d08bae7d0b1ac121.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.roleManagement.entitlementManagement.roleDefinitions.item collection
func (m *EntitlementManagementRequestBuilder) RoleDefinitionsById(id string)(*i450e6fa82c54d6f54b7abb6a155c2b0b0cb9aced086261fdf11e3d7d814c6286.UnifiedRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleDefinition%2Did"] = id
    }
    return i450e6fa82c54d6f54b7abb6a155c2b0b0cb9aced086261fdf11e3d7d814c6286.NewUnifiedRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleEligibilityScheduleInstances the roleEligibilityScheduleInstances property
func (m *EntitlementManagementRequestBuilder) RoleEligibilityScheduleInstances()(*i75e0187968a6021e302f1aa7a3eb2a876441acad1f351cff3d51d6d0c4cd52a4.RoleEligibilityScheduleInstancesRequestBuilder) {
    return i75e0187968a6021e302f1aa7a3eb2a876441acad1f351cff3d51d6d0c4cd52a4.NewRoleEligibilityScheduleInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilityScheduleInstancesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.roleManagement.entitlementManagement.roleEligibilityScheduleInstances.item collection
func (m *EntitlementManagementRequestBuilder) RoleEligibilityScheduleInstancesById(id string)(*i987a32583667f0a5ef515c4455c9a8fa39e772a7676688f15e1cc180c13ae1c9.UnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilityScheduleInstance%2Did"] = id
    }
    return i987a32583667f0a5ef515c4455c9a8fa39e772a7676688f15e1cc180c13ae1c9.NewUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleEligibilityScheduleRequests the roleEligibilityScheduleRequests property
func (m *EntitlementManagementRequestBuilder) RoleEligibilityScheduleRequests()(*i6105172c68f51bb2e2927309c74148ed2bf9eaebff5fecc8144191485bb58a1c.RoleEligibilityScheduleRequestsRequestBuilder) {
    return i6105172c68f51bb2e2927309c74148ed2bf9eaebff5fecc8144191485bb58a1c.NewRoleEligibilityScheduleRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilityScheduleRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.roleManagement.entitlementManagement.roleEligibilityScheduleRequests.item collection
func (m *EntitlementManagementRequestBuilder) RoleEligibilityScheduleRequestsById(id string)(*i3b9e25b9426e31e4bcd53952d1b0865afd529e64b7d980cc74d8ef468d55a6ca.UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilityScheduleRequest%2Did"] = id
    }
    return i3b9e25b9426e31e4bcd53952d1b0865afd529e64b7d980cc74d8ef468d55a6ca.NewUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleEligibilitySchedules the roleEligibilitySchedules property
func (m *EntitlementManagementRequestBuilder) RoleEligibilitySchedules()(*i3a967725e663f75605977fcee02eb4b2642ee31d0648ec4e2ed0f800882b32bc.RoleEligibilitySchedulesRequestBuilder) {
    return i3a967725e663f75605977fcee02eb4b2642ee31d0648ec4e2ed0f800882b32bc.NewRoleEligibilitySchedulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilitySchedulesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.roleManagement.entitlementManagement.roleEligibilitySchedules.item collection
func (m *EntitlementManagementRequestBuilder) RoleEligibilitySchedulesById(id string)(*i9737b224e822cfafa077e03d8ff571fad57ee6c6761c45f56482ea39d6af9c92.UnifiedRoleEligibilityScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilitySchedule%2Did"] = id
    }
    return i9737b224e822cfafa077e03d8ff571fad57ee6c6761c45f56482ea39d6af9c92.NewUnifiedRoleEligibilityScheduleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
