package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder provides operations to call the additionalAccess method.
type EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetQueryParameters invoke function additionalAccess
type EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetQueryParameters
}
// NewEntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderInternal instantiates a new AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder and sets the default values.
func NewEntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, accessPackageId *string, incompatibleAccessPackageId *string)(*EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) {
    m := &EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignments/additionalAccess(accessPackageId='{accessPackageId}',incompatibleAccessPackageId='{incompatibleAccessPackageId}'){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    if accessPackageId != nil {
        m.BaseRequestBuilder.PathParameters["accessPackageId"] = *accessPackageId
    }
    if incompatibleAccessPackageId != nil {
        m.BaseRequestBuilder.PathParameters["incompatibleAccessPackageId"] = *incompatibleAccessPackageId
    }
    return m
}
// NewEntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder instantiates a new AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder and sets the default values.
func NewEntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function additionalAccess
func (m *EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration)(EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseable), nil
}
// ToGetRequestInformation invoke function additionalAccess
func (m *EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) {
    return NewEntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
