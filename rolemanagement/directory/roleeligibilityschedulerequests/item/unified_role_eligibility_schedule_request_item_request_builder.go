package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i068716137e2995671c21eef72ec68f439e24e23e075243c2c047562ff086e084 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/directory/roleeligibilityschedulerequests/item/cancel"
    i1436e075ceda3211f28cd6c86c40328ef8a0635eacd86947f7d1d9da93b3d370 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/directory/roleeligibilityschedulerequests/item/appscope"
    i182d6f67ef6ebb70384f3b31a3ac9afef417e73ae4daba667f3f3c1e73861532 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/directory/roleeligibilityschedulerequests/item/directoryscope"
    i1f15baaaf0166badd21f3d7ae6da3ffbfa83697d4cf586316554095c18e97895 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/directory/roleeligibilityschedulerequests/item/principal"
    ia04d9831fcf2650e3920b34ee6c643e1c72754c63c24e7ff1d0235fc55b22ae8 "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/directory/roleeligibilityschedulerequests/item/targetschedule"
    ie85b40180c5ab54474e0625656270f71998f76706b96c9d6a50a3ae8520469ca "github.com/microsoftgraph/msgraph-sdk-go/rolemanagement/directory/roleeligibilityschedulerequests/item/roledefinition"
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
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) AppScope()(*i1436e075ceda3211f28cd6c86c40328ef8a0635eacd86947f7d1d9da93b3d370.AppScopeRequestBuilder) {
    return i1436e075ceda3211f28cd6c86c40328ef8a0635eacd86947f7d1d9da93b3d370.NewAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Cancel()(*i068716137e2995671c21eef72ec68f439e24e23e075243c2c047562ff086e084.CancelRequestBuilder) {
    return i068716137e2995671c21eef72ec68f439e24e23e075243c2c047562ff086e084.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal instantiates a new UnifiedRoleEligibilityScheduleRequestItemRequestBuilder and sets the default values.
func NewUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    m := &UnifiedRoleEligibilityScheduleRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory/roleEligibilityScheduleRequests/{unifiedRoleEligibilityScheduleRequest%2Did}{?%24select,%24expand}";
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
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) DirectoryScope()(*i182d6f67ef6ebb70384f3b31a3ac9afef417e73ae4daba667f3f3c1e73861532.DirectoryScopeRequestBuilder) {
    return i182d6f67ef6ebb70384f3b31a3ac9afef417e73ae4daba667f3f3c1e73861532.NewDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Principal()(*i1f15baaaf0166badd21f3d7ae6da3ffbfa83697d4cf586316554095c18e97895.PrincipalRequestBuilder) {
    return i1f15baaaf0166badd21f3d7ae6da3ffbfa83697d4cf586316554095c18e97895.NewPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinition the roleDefinition property
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) RoleDefinition()(*ie85b40180c5ab54474e0625656270f71998f76706b96c9d6a50a3ae8520469ca.RoleDefinitionRequestBuilder) {
    return ie85b40180c5ab54474e0625656270f71998f76706b96c9d6a50a3ae8520469ca.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TargetSchedule the targetSchedule property
func (m *UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) TargetSchedule()(*ia04d9831fcf2650e3920b34ee6c643e1c72754c63c24e7ff1d0235fc55b22ae8.TargetScheduleRequestBuilder) {
    return ia04d9831fcf2650e3920b34ee6c643e1c72754c63c24e7ff1d0235fc55b22ae8.NewTargetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
