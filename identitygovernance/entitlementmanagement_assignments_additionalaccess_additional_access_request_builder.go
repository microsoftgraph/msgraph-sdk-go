package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder provides operations to call the additionalAccess method.
type EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilderGetQueryParameters in Microsoft Entra Entitlement Management, retrieve a collection of accessPackageAssignment objects that indicate a target user has an assignment to a specified access package and also an assignment to another, potentially incompatible, access package.  This can be used to prepare to configure the incompatible access packages for a specific access package.
type EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilderInternal instantiates a new EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder) {
    m := &EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignments/additionalAccess(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder instantiates a new EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilderInternal(urlParams, requestAdapter)
}
// Get in Microsoft Entra Entitlement Management, retrieve a collection of accessPackageAssignment objects that indicate a target user has an assignment to a specified access package and also an assignment to another, potentially incompatible, access package.  This can be used to prepare to configure the incompatible access packages for a specific access package.
// Deprecated: This method is obsolete. Use GetAsAdditionalAccessGetResponse instead.
// returns a EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignment-additionalaccess?view=graph-rest-1.0
func (m *EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilderGetRequestConfiguration)(EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessResponseable), nil
}
// GetAsAdditionalAccessGetResponse in Microsoft Entra Entitlement Management, retrieve a collection of accessPackageAssignment objects that indicate a target user has an assignment to a specified access package and also an assignment to another, potentially incompatible, access package.  This can be used to prepare to configure the incompatible access packages for a specific access package.
// returns a EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignment-additionalaccess?view=graph-rest-1.0
func (m *EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder) GetAsAdditionalAccessGetResponse(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilderGetRequestConfiguration)(EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessGetResponseable), nil
}
// ToGetRequestInformation in Microsoft Entra Entitlement Management, retrieve a collection of accessPackageAssignment objects that indicate a target user has an assignment to a specified access package and also an assignment to another, potentially incompatible, access package.  This can be used to prepare to configure the incompatible access packages for a specific access package.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder when successful
func (m *EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder) {
    return NewEntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
