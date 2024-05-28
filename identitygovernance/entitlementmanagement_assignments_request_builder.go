package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAssignmentsRequestBuilder provides operations to manage the assignments property of the microsoft.graph.entitlementManagement entity.
type EntitlementmanagementAssignmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAssignmentsRequestBuilderGetQueryParameters in Microsoft Entra entitlement management, retrieve a list of accessPackageAssignment objects. For directory-wide administrators, the resulting list includes all the assignments, current and well as expired, that the caller has access to read, across all catalogs and access packages.  If the caller is on behalf of a delegated user who is assigned only to catalog-specific delegated administrative roles, the request must supply a filter to indicate a specific access package, such as: $filter=accessPackage/id eq 'a914b616-e04e-476b-aa37-91038f0b165b'.
type EntitlementmanagementAssignmentsRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementAssignmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAssignmentsRequestBuilderGetQueryParameters
}
// EntitlementmanagementAssignmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdditionalAccess provides operations to call the additionalAccess method.
// returns a *EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder when successful
func (m *EntitlementmanagementAssignmentsRequestBuilder) AdditionalAccess()(*EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilder) {
    return NewEntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageId provides operations to call the additionalAccess method.
// returns a *EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder when successful
func (m *EntitlementmanagementAssignmentsRequestBuilder) AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageId(accessPackageId *string, incompatibleAccessPackageId *string)(*EntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) {
    return NewEntitlementmanagementAssignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, accessPackageId, incompatibleAccessPackageId)
}
// ByAccessPackageAssignmentId provides operations to manage the assignments property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder when successful
func (m *EntitlementmanagementAssignmentsRequestBuilder) ByAccessPackageAssignmentId(accessPackageAssignmentId string)(*EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageAssignmentId != "" {
        urlTplParams["accessPackageAssignment%2Did"] = accessPackageAssignmentId
    }
    return NewEntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAssignmentsRequestBuilderInternal instantiates a new EntitlementmanagementAssignmentsRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentsRequestBuilder) {
    m := &EntitlementmanagementAssignmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAssignmentsRequestBuilder instantiates a new EntitlementmanagementAssignmentsRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAssignmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementAssignmentsCountRequestBuilder when successful
func (m *EntitlementmanagementAssignmentsRequestBuilder) Count()(*EntitlementmanagementAssignmentsCountRequestBuilder) {
    return NewEntitlementmanagementAssignmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *EntitlementmanagementAssignmentsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder when successful
func (m *EntitlementmanagementAssignmentsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*EntitlementmanagementAssignmentsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    return NewEntitlementmanagementAssignmentsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get in Microsoft Entra entitlement management, retrieve a list of accessPackageAssignment objects. For directory-wide administrators, the resulting list includes all the assignments, current and well as expired, that the caller has access to read, across all catalogs and access packages.  If the caller is on behalf of a delegated user who is assigned only to catalog-specific delegated administrative roles, the request must supply a filter to indicate a specific access package, such as: $filter=accessPackage/id eq 'a914b616-e04e-476b-aa37-91038f0b165b'.
// returns a AccessPackageAssignmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/entitlementmanagement-list-assignments?view=graph-rest-1.0
func (m *EntitlementmanagementAssignmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageAssignmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentCollectionResponseable), nil
}
// Post create new navigation property to assignments for identityGovernance
// returns a AccessPackageAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAssignmentsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable, requestConfiguration *EntitlementmanagementAssignmentsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable), nil
}
// ToGetRequestInformation in Microsoft Entra entitlement management, retrieve a list of accessPackageAssignment objects. For directory-wide administrators, the resulting list includes all the assignments, current and well as expired, that the caller has access to read, across all catalogs and access packages.  If the caller is on behalf of a delegated user who is assigned only to catalog-specific delegated administrative roles, the request must supply a filter to indicate a specific access package, such as: $filter=accessPackage/id eq 'a914b616-e04e-476b-aa37-91038f0b165b'.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to assignments for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable, requestConfiguration *EntitlementmanagementAssignmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAssignmentsRequestBuilder when successful
func (m *EntitlementmanagementAssignmentsRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAssignmentsRequestBuilder) {
    return NewEntitlementmanagementAssignmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
