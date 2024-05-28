package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder provides operations to manage the tenants property of the microsoft.graph.multiTenantOrganization entity.
type MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderGetQueryParameters get a tenant and its properties in the multitenant organization.
type MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderGetQueryParameters
}
// MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderInternal instantiates a new MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder and sets the default values.
func NewMultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) {
    m := &MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/multiTenantOrganization/tenants/{multiTenantOrganizationMember%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder instantiates a new MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder and sets the default values.
func NewMultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove a tenant from a multitenant organization. A tenant can be removed in the following scenarios:
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/multitenantorganization-delete-tenants?view=graph-rest-1.0
func (m *MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get a tenant and its properties in the multitenant organization.
// returns a MultiTenantOrganizationMemberable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/multitenantorganizationmember-get?view=graph-rest-1.0
func (m *MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationMemberable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMultiTenantOrganizationMemberFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationMemberable), nil
}
// Patch update the navigation property tenants in tenantRelationships
// returns a MultiTenantOrganizationMemberable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationMemberable, requestConfiguration *MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationMemberable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMultiTenantOrganizationMemberFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationMemberable), nil
}
// ToDeleteRequestInformation remove a tenant from a multitenant organization. A tenant can be removed in the following scenarios:
// returns a *RequestInformation when successful
func (m *MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get a tenant and its properties in the multitenant organization.
// returns a *RequestInformation when successful
func (m *MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property tenants in tenantRelationships
// returns a *RequestInformation when successful
func (m *MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationMemberable, requestConfiguration *MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder when successful
func (m *MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) WithUrl(rawUrl string)(*MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) {
    return NewMultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
