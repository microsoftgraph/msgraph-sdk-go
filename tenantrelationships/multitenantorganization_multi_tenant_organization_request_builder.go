package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MultitenantorganizationMultiTenantOrganizationRequestBuilder provides operations to manage the multiTenantOrganization property of the microsoft.graph.tenantRelationship entity.
type MultitenantorganizationMultiTenantOrganizationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MultitenantorganizationMultiTenantOrganizationRequestBuilderGetQueryParameters get properties of the multitenant organization.
type MultitenantorganizationMultiTenantOrganizationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MultitenantorganizationMultiTenantOrganizationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MultitenantorganizationMultiTenantOrganizationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MultitenantorganizationMultiTenantOrganizationRequestBuilderGetQueryParameters
}
// MultitenantorganizationMultiTenantOrganizationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MultitenantorganizationMultiTenantOrganizationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMultitenantorganizationMultiTenantOrganizationRequestBuilderInternal instantiates a new MultitenantorganizationMultiTenantOrganizationRequestBuilder and sets the default values.
func NewMultitenantorganizationMultiTenantOrganizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MultitenantorganizationMultiTenantOrganizationRequestBuilder) {
    m := &MultitenantorganizationMultiTenantOrganizationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/multiTenantOrganization{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMultitenantorganizationMultiTenantOrganizationRequestBuilder instantiates a new MultitenantorganizationMultiTenantOrganizationRequestBuilder and sets the default values.
func NewMultitenantorganizationMultiTenantOrganizationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MultitenantorganizationMultiTenantOrganizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMultitenantorganizationMultiTenantOrganizationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get properties of the multitenant organization.
// returns a MultiTenantOrganizationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/multitenantorganization-get?view=graph-rest-1.0
func (m *MultitenantorganizationMultiTenantOrganizationRequestBuilder) Get(ctx context.Context, requestConfiguration *MultitenantorganizationMultiTenantOrganizationRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMultiTenantOrganizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationable), nil
}
// JoinRequest provides operations to manage the joinRequest property of the microsoft.graph.multiTenantOrganization entity.
// returns a *MultitenantorganizationJoinrequestJoinRequestRequestBuilder when successful
func (m *MultitenantorganizationMultiTenantOrganizationRequestBuilder) JoinRequest()(*MultitenantorganizationJoinrequestJoinRequestRequestBuilder) {
    return NewMultitenantorganizationJoinrequestJoinRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch create a new multitenant organization. By default, the creator tenant becomes an owner tenant upon successful creation. Only owner tenants can manage a multitenant organization.
// returns a MultiTenantOrganizationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tenantrelationship-put-multitenantorganization?view=graph-rest-1.0
func (m *MultitenantorganizationMultiTenantOrganizationRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationable, requestConfiguration *MultitenantorganizationMultiTenantOrganizationRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMultiTenantOrganizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationable), nil
}
// Tenants provides operations to manage the tenants property of the microsoft.graph.multiTenantOrganization entity.
// returns a *MultitenantorganizationTenantsRequestBuilder when successful
func (m *MultitenantorganizationMultiTenantOrganizationRequestBuilder) Tenants()(*MultitenantorganizationTenantsRequestBuilder) {
    return NewMultitenantorganizationTenantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get properties of the multitenant organization.
// returns a *RequestInformation when successful
func (m *MultitenantorganizationMultiTenantOrganizationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MultitenantorganizationMultiTenantOrganizationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation create a new multitenant organization. By default, the creator tenant becomes an owner tenant upon successful creation. Only owner tenants can manage a multitenant organization.
// returns a *RequestInformation when successful
func (m *MultitenantorganizationMultiTenantOrganizationRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MultiTenantOrganizationable, requestConfiguration *MultitenantorganizationMultiTenantOrganizationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MultitenantorganizationMultiTenantOrganizationRequestBuilder when successful
func (m *MultitenantorganizationMultiTenantOrganizationRequestBuilder) WithUrl(rawUrl string)(*MultitenantorganizationMultiTenantOrganizationRequestBuilder) {
    return NewMultitenantorganizationMultiTenantOrganizationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
