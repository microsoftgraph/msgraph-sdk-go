package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i28469022081e100ad5feec64a0c99129191736547aadc119e5eacd4b98954919 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/targetschedule"
    i38ca64b8b058140def4153684e693311e56f0603177bcec236c50544a07a446a "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/activatedusing"
    i43519742abaa7fdd577f8211a5a4de7a5a3ce6b75a0a3ecb63c73732a8129a36 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/principal"
    i49779d099fba6b7fdca0cde9f16a0863487cced229750a243d32f7caa3b02192 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/roledefinition"
    i61bcc2628b8c7f8d8cf13f3357710e8ffb34ade9f23e103c82cf9bb901980c35 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/directoryscope"
    ibb3650c856589e0d8dd876ef57453246087e344aabb2cf09a6d4791fff2a9d45 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/appscope"
    ic1b7c918ee37633806f85e686be26c0f33a0767bb232031438f3dc8763be87ff "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/cancel"
)

// UnifiedRoleAssignmentScheduleRequestItemRequestBuilder provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteOptions options for Delete
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetOptions options for Get
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters get roleAssignmentScheduleRequests from roleManagement
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchOptions options for Patch
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ActivatedUsing the activatedUsing property
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ActivatedUsing()(*i38ca64b8b058140def4153684e693311e56f0603177bcec236c50544a07a446a.ActivatedUsingRequestBuilder) {
    return i38ca64b8b058140def4153684e693311e56f0603177bcec236c50544a07a446a.NewActivatedUsingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppScope the appScope property
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) AppScope()(*ibb3650c856589e0d8dd876ef57453246087e344aabb2cf09a6d4791fff2a9d45.AppScopeRequestBuilder) {
    return ibb3650c856589e0d8dd876ef57453246087e344aabb2cf09a6d4791fff2a9d45.NewAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Cancel()(*ic1b7c918ee37633806f85e686be26c0f33a0767bb232031438f3dc8763be87ff.CancelRequestBuilder) {
    return ic1b7c918ee37633806f85e686be26c0f33a0767bb232031438f3dc8763be87ff.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal instantiates a new UnifiedRoleAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    m := &UnifiedRoleAssignmentScheduleRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/roleAssignmentScheduleRequests/{unifiedRoleAssignmentScheduleRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilder instantiates a new UnifiedRoleAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get roleAssignmentScheduleRequests from roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreateGetRequestInformation(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreatePatchRequestInformation(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Delete(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteOptions)(error) {
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
// DirectoryScope the directoryScope property
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) DirectoryScope()(*i61bcc2628b8c7f8d8cf13f3357710e8ffb34ade9f23e103c82cf9bb901980c35.DirectoryScopeRequestBuilder) {
    return i61bcc2628b8c7f8d8cf13f3357710e8ffb34ade9f23e103c82cf9bb901980c35.NewDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get roleAssignmentScheduleRequests from roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Get(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleAssignmentScheduleRequestFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable), nil
}
// Patch update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Patch(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchOptions)(error) {
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
// Principal the principal property
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Principal()(*i43519742abaa7fdd577f8211a5a4de7a5a3ce6b75a0a3ecb63c73732a8129a36.PrincipalRequestBuilder) {
    return i43519742abaa7fdd577f8211a5a4de7a5a3ce6b75a0a3ecb63c73732a8129a36.NewPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinition the roleDefinition property
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) RoleDefinition()(*i49779d099fba6b7fdca0cde9f16a0863487cced229750a243d32f7caa3b02192.RoleDefinitionRequestBuilder) {
    return i49779d099fba6b7fdca0cde9f16a0863487cced229750a243d32f7caa3b02192.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TargetSchedule the targetSchedule property
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) TargetSchedule()(*i28469022081e100ad5feec64a0c99129191736547aadc119e5eacd4b98954919.TargetScheduleRequestBuilder) {
    return i28469022081e100ad5feec64a0c99129191736547aadc119e5eacd4b98954919.NewTargetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
