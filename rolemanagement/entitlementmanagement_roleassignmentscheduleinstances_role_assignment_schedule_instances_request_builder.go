package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder provides operations to manage the roleAssignmentScheduleInstances property of the microsoft.graph.rbacApplication entity.
type EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilderGetQueryParameters instances for active role assignments.
type EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilderGetQueryParameters
}
// EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleAssignmentScheduleInstanceId provides operations to manage the roleAssignmentScheduleInstances property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder) ByUnifiedRoleAssignmentScheduleInstanceId(unifiedRoleAssignmentScheduleInstanceId string)(*EntitlementmanagementRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleAssignmentScheduleInstanceId != "" {
        urlTplParams["unifiedRoleAssignmentScheduleInstance%2Did"] = unifiedRoleAssignmentScheduleInstanceId
    }
    return NewEntitlementmanagementRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilderInternal instantiates a new EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder) {
    m := &EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/roleAssignmentScheduleInstances{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder instantiates a new EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementRoleassignmentscheduleinstancesCountRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder) Count()(*EntitlementmanagementRoleassignmentscheduleinstancesCountRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentscheduleinstancesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *EntitlementmanagementRoleassignmentscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*EntitlementmanagementRoleassignmentscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get instances for active role assignments.
// returns a UnifiedRoleAssignmentScheduleInstanceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleInstanceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleAssignmentScheduleInstanceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleInstanceCollectionResponseable), nil
}
// Post create new navigation property to roleAssignmentScheduleInstances for roleManagement
// returns a UnifiedRoleAssignmentScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleInstanceable, requestConfiguration *EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleInstanceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleAssignmentScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleInstanceable), nil
}
// ToGetRequestInformation instances for active role assignments.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to roleAssignmentScheduleInstances for roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleInstanceable, requestConfiguration *EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
