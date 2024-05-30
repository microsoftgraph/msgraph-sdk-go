package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder provides operations to call the filterByCurrentUser method.
type EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetQueryParameters get the instances of eligible roles for the calling principal.
type EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal instantiates a new EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, on *string)(*EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    m := &EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/roleEligibilityScheduleInstances/filterByCurrentUser(on='{on}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if on != nil {
        m.BaseRequestBuilder.PathParameters["on"] = *on
    }
    return m
}
// NewEntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder instantiates a new EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get the instances of eligible roles for the calling principal.
// Deprecated: This method is obsolete. Use GetAsFilterByCurrentUserWithOnGetResponse instead.
// returns a EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleeligibilityscheduleinstance-filterbycurrentuser?view=graph-rest-1.0
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable), nil
}
// GetAsFilterByCurrentUserWithOnGetResponse get the instances of eligible roles for the calling principal.
// returns a EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleeligibilityscheduleinstance-filterbycurrentuser?view=graph-rest-1.0
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) GetAsFilterByCurrentUserWithOnGetResponse(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable), nil
}
// ToGetRequestInformation get the instances of eligible roles for the calling principal.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
