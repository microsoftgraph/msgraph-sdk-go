package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i42b83c39a0076c0e0781b3b5a2bdd65960c11c81af609daf85b63f65a4bbdf56 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/directoryscope"
    i7f8de8adecdb539cfae05e3e1caeed1a0da671f22b3cd578619f93609a832d1e "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/cancel"
    i8a9242003b3db7e918b6daaf0c995000e556af520819ce7760ab6278c27d1be7 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/principal"
    i8a970ee7b4f9c0878d054b2f1a61f3bfafcff679c996655c6758d61028e77509 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/appscope"
    ibca4dcbed47532998a01d10a541a9ac08ea814c6c17a6a53dfd1fdc09d94423a "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/targetschedule"
    id2ba362c945df44ed69b09e9592d7b158ed2559c76b2621b43fccc8cabab49ce "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/roledefinition"
)

// UnifiedRoleEligibilityScheduleRequestItemRequestBuilder provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
type UnifiedRoleEligibilityScheduleRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UnifiedRoleEligibilityScheduleRequestItemRequestBuilderDeleteOptions options for Delete
type UnifiedRoleEligibilityScheduleRequestItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// UnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetOptions options for Get
type UnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// UnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetQueryParameters get roleEligibilityScheduleRequests from roleManagement
type UnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UnifiedRoleEligibilityScheduleRequestItemRequestBuilderPatchOptions options for Patch
type UnifiedRoleEligibilityScheduleRequestItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AppScope the appScope property
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) AppScope()(*i8a970ee7b4f9c0878d054b2f1a61f3bfafcff679c996655c6758d61028e77509.AppScopeRequestBuilder) {
    return i8a970ee7b4f9c0878d054b2f1a61f3bfafcff679c996655c6758d61028e77509.NewAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Cancel()(*i7f8de8adecdb539cfae05e3e1caeed1a0da671f22b3cd578619f93609a832d1e.CancelRequestBuilder) {
    return i7f8de8adecdb539cfae05e3e1caeed1a0da671f22b3cd578619f93609a832d1e.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal instantiates a new UnifiedRoleEligibilityScheduleRequestItemRequestBuilder and sets the default values.
func NewUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    m := &UnifiedRoleEligibilityScheduleRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/roleEligibilityScheduleRequests/{unifiedRoleEligibilityScheduleRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleEligibilityScheduleRequestItemRequestBuilder instantiates a new UnifiedRoleEligibilityScheduleRequestItemRequestBuilder and sets the default values.
func NewUnifiedRoleEligibilityScheduleRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleEligibilityScheduleRequests for roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRoleEligibilityScheduleRequestItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get roleEligibilityScheduleRequests from roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) CreateGetRequestInformation(options *UnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleEligibilityScheduleRequests in roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) CreatePatchRequestInformation(options *UnifiedRoleEligibilityScheduleRequestItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property roleEligibilityScheduleRequests for roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Delete(options *UnifiedRoleEligibilityScheduleRequestItemRequestBuilderDeleteOptions)(error) {
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
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) DirectoryScope()(*i42b83c39a0076c0e0781b3b5a2bdd65960c11c81af609daf85b63f65a4bbdf56.DirectoryScopeRequestBuilder) {
    return i42b83c39a0076c0e0781b3b5a2bdd65960c11c81af609daf85b63f65a4bbdf56.NewDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get roleEligibilityScheduleRequests from roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Get(options *UnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleRequestFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable), nil
}
// Patch update the navigation property roleEligibilityScheduleRequests in roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Patch(options *UnifiedRoleEligibilityScheduleRequestItemRequestBuilderPatchOptions)(error) {
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
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Principal()(*i8a9242003b3db7e918b6daaf0c995000e556af520819ce7760ab6278c27d1be7.PrincipalRequestBuilder) {
    return i8a9242003b3db7e918b6daaf0c995000e556af520819ce7760ab6278c27d1be7.NewPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinition the roleDefinition property
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) RoleDefinition()(*id2ba362c945df44ed69b09e9592d7b158ed2559c76b2621b43fccc8cabab49ce.RoleDefinitionRequestBuilder) {
    return id2ba362c945df44ed69b09e9592d7b158ed2559c76b2621b43fccc8cabab49ce.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TargetSchedule the targetSchedule property
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) TargetSchedule()(*ibca4dcbed47532998a01d10a541a9ac08ea814c6c17a6a53dfd1fdc09d94423a.TargetScheduleRequestBuilder) {
    return ibca4dcbed47532998a01d10a541a9ac08ea814c6c17a6a53dfd1fdc09d94423a.NewTargetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
