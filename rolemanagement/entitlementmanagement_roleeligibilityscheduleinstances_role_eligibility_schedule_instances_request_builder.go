package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder provides operations to manage the roleEligibilityScheduleInstances property of the microsoft.graph.rbacApplication entity.
type EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilderGetQueryParameters instances for role eligibility requests.
type EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilderGetQueryParameters
}
// EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleEligibilityScheduleInstanceId provides operations to manage the roleEligibilityScheduleInstances property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder) ByUnifiedRoleEligibilityScheduleInstanceId(unifiedRoleEligibilityScheduleInstanceId string)(*EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleEligibilityScheduleInstanceId != "" {
        urlTplParams["unifiedRoleEligibilityScheduleInstance%2Did"] = unifiedRoleEligibilityScheduleInstanceId
    }
    return NewEntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilderInternal instantiates a new EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder) {
    m := &EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/roleEligibilityScheduleInstances{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder instantiates a new EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementRoleeligibilityscheduleinstancesCountRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder) Count()(*EntitlementmanagementRoleeligibilityscheduleinstancesCountRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityscheduleinstancesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get instances for role eligibility requests.
// returns a UnifiedRoleEligibilityScheduleInstanceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleInstanceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceCollectionResponseable), nil
}
// Post create new navigation property to roleEligibilityScheduleInstances for roleManagement
// returns a UnifiedRoleEligibilityScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceable, requestConfiguration *EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceable), nil
}
// ToGetRequestInformation instances for role eligibility requests.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to roleEligibilityScheduleInstances for roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceable, requestConfiguration *EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
