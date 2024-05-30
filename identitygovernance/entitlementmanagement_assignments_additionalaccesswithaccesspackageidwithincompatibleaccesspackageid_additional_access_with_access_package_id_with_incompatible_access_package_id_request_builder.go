package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder provides operations to call the additionalAccess method.
type EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetQueryParameters invoke function additionalAccess
type EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderInternal instantiates a new EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, accessPackageId *string, incompatibleAccessPackageId *string)(*EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) {
    m := &EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignments/additionalAccess(accessPackageId='{accessPackageId}',incompatibleAccessPackageId='{incompatibleAccessPackageId}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if accessPackageId != nil {
        m.BaseRequestBuilder.PathParameters["accessPackageId"] = *accessPackageId
    }
    if incompatibleAccessPackageId != nil {
        m.BaseRequestBuilder.PathParameters["incompatibleAccessPackageId"] = *incompatibleAccessPackageId
    }
    return m
}
// NewEntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder instantiates a new EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function additionalAccess
// Deprecated: This method is obsolete. Use GetAsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse instead.
// returns a EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration)(EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseable), nil
}
// GetAsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse invoke function additionalAccess
// returns a EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) GetAsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration)(EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponseable), nil
}
// ToGetRequestInformation invoke function additionalAccess
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder when successful
func (m *EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) {
    return NewEntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
