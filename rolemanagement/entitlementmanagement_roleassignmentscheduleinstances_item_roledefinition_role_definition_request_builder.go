package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
type EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilderGetQueryParameters detailed information for the roleDefinition object that is referenced through the roleDefinitionId property.
type EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilderInternal instantiates a new EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder) {
    m := &EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/roleAssignmentScheduleInstances/{unifiedRoleAssignmentScheduleInstance%2Did}/roleDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder instantiates a new EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get detailed information for the roleDefinition object that is referenced through the roleDefinitionId property.
// returns a UnifiedRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleDefinitionable), nil
}
// ToGetRequestInformation detailed information for the roleDefinition object that is referenced through the roleDefinitionId property.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
