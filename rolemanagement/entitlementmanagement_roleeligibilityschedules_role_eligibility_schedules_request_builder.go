package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder provides operations to manage the roleEligibilitySchedules property of the microsoft.graph.rbacApplication entity.
type EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilderGetQueryParameters schedules for role eligibility operations.
type EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilderGetQueryParameters
}
// EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleEligibilityScheduleId provides operations to manage the roleEligibilitySchedules property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder) ByUnifiedRoleEligibilityScheduleId(unifiedRoleEligibilityScheduleId string)(*EntitlementmanagementRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleEligibilityScheduleId != "" {
        urlTplParams["unifiedRoleEligibilitySchedule%2Did"] = unifiedRoleEligibilityScheduleId
    }
    return NewEntitlementmanagementRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilderInternal instantiates a new EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder) {
    m := &EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/roleEligibilitySchedules{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder instantiates a new EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementRoleeligibilityschedulesCountRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder) Count()(*EntitlementmanagementRoleeligibilityschedulesCountRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *EntitlementmanagementRoleeligibilityschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*EntitlementmanagementRoleeligibilityschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get schedules for role eligibility operations.
// returns a UnifiedRoleEligibilityScheduleCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleCollectionResponseable), nil
}
// Post create new navigation property to roleEligibilitySchedules for roleManagement
// returns a UnifiedRoleEligibilityScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleable, requestConfiguration *EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleable), nil
}
// ToGetRequestInformation schedules for role eligibility operations.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to roleEligibilitySchedules for roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleable, requestConfiguration *EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
