package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MultitenantorganizationTenantsRequestBuilder provides operations to manage the tenants property of the microsoft.graph.multiTenantOrganization entity.
type MultitenantorganizationTenantsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MultitenantorganizationTenantsRequestBuilderGetQueryParameters list the tenants and their properties in the multitenant organization.
type MultitenantorganizationTenantsRequestBuilderGetQueryParameters struct {
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
// MultitenantorganizationTenantsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MultitenantorganizationTenantsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MultitenantorganizationTenantsRequestBuilderGetQueryParameters
}
// MultitenantorganizationTenantsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MultitenantorganizationTenantsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMultiTenantOrganizationMemberId provides operations to manage the tenants property of the microsoft.graph.multiTenantOrganization entity.
// returns a *MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder when successful
func (m *MultitenantorganizationTenantsRequestBuilder) ByMultiTenantOrganizationMemberId(multiTenantOrganizationMemberId string)(*MultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if multiTenantOrganizationMemberId != "" {
        urlTplParams["multiTenantOrganizationMember%2Did"] = multiTenantOrganizationMemberId
    }
    return NewMultitenantorganizationTenantsMultiTenantOrganizationMemberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMultitenantorganizationTenantsRequestBuilderInternal instantiates a new MultitenantorganizationTenantsRequestBuilder and sets the default values.
func NewMultitenantorganizationTenantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MultitenantorganizationTenantsRequestBuilder) {
    m := &MultitenantorganizationTenantsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/multiTenantOrganization/tenants{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMultitenantorganizationTenantsRequestBuilder instantiates a new MultitenantorganizationTenantsRequestBuilder and sets the default values.
func NewMultitenantorganizationTenantsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MultitenantorganizationTenantsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMultitenantorganizationTenantsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MultitenantorganizationTenantsCountRequestBuilder when successful
func (m *MultitenantorganizationTenantsRequestBuilder) Count()(*MultitenantorganizationTenantsCountRequestBuilder) {
    return NewMultitenantorganizationTenantsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list the tenants and their properties in the multitenant organization.
// returns a MultiTenantOrganizationMemberCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/multitenantorganization-list-tenants?view=graph-rest-1.0
func (m *MultitenantorganizationTenantsRequestBuilder) Get(ctx context.Context, requestConfiguration *MultitenantorganizationTenantsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationMemberCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMultiTenantOrganizationMemberCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationMemberCollectionResponseable), nil
}
// Post add a tenant to a multitenant organization. The administrator of an owner tenant has the permissions to add tenants to the multitenant organization. The added tenant is in the pending state until the administrator of the added tenant joins the multitenant organization by submitting a join request. A tenant can be part of only one multitenant organization.
// returns a MultiTenantOrganizationMemberable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/multitenantorganization-post-tenants?view=graph-rest-1.0
func (m *MultitenantorganizationTenantsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationMemberable, requestConfiguration *MultitenantorganizationTenantsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationMemberable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation list the tenants and their properties in the multitenant organization.
// returns a *RequestInformation when successful
func (m *MultitenantorganizationTenantsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MultitenantorganizationTenantsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation add a tenant to a multitenant organization. The administrator of an owner tenant has the permissions to add tenants to the multitenant organization. The added tenant is in the pending state until the administrator of the added tenant joins the multitenant organization by submitting a join request. A tenant can be part of only one multitenant organization.
// returns a *RequestInformation when successful
func (m *MultitenantorganizationTenantsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationMemberable, requestConfiguration *MultitenantorganizationTenantsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MultitenantorganizationTenantsRequestBuilder when successful
func (m *MultitenantorganizationTenantsRequestBuilder) WithUrl(rawUrl string)(*MultitenantorganizationTenantsRequestBuilder) {
    return NewMultitenantorganizationTenantsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
