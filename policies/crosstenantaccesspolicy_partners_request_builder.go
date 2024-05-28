package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CrosstenantaccesspolicyPartnersRequestBuilder provides operations to manage the partners property of the microsoft.graph.crossTenantAccessPolicy entity.
type CrosstenantaccesspolicyPartnersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CrosstenantaccesspolicyPartnersRequestBuilderGetQueryParameters get a list of all partner configurations within a cross-tenant access policy. You can also use the $expand parameter to list the user synchronization policy for all partner configurations.
type CrosstenantaccesspolicyPartnersRequestBuilderGetQueryParameters struct {
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
// CrosstenantaccesspolicyPartnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrosstenantaccesspolicyPartnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CrosstenantaccesspolicyPartnersRequestBuilderGetQueryParameters
}
// CrosstenantaccesspolicyPartnersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrosstenantaccesspolicyPartnersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCrossTenantAccessPolicyConfigurationPartnerTenantId provides operations to manage the partners property of the microsoft.graph.crossTenantAccessPolicy entity.
// returns a *CrosstenantaccesspolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder when successful
func (m *CrosstenantaccesspolicyPartnersRequestBuilder) ByCrossTenantAccessPolicyConfigurationPartnerTenantId(crossTenantAccessPolicyConfigurationPartnerTenantId string)(*CrosstenantaccesspolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if crossTenantAccessPolicyConfigurationPartnerTenantId != "" {
        urlTplParams["crossTenantAccessPolicyConfigurationPartner%2DtenantId"] = crossTenantAccessPolicyConfigurationPartnerTenantId
    }
    return NewCrosstenantaccesspolicyPartnersCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCrosstenantaccesspolicyPartnersRequestBuilderInternal instantiates a new CrosstenantaccesspolicyPartnersRequestBuilder and sets the default values.
func NewCrosstenantaccesspolicyPartnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrosstenantaccesspolicyPartnersRequestBuilder) {
    m := &CrosstenantaccesspolicyPartnersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/crossTenantAccessPolicy/partners{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCrosstenantaccesspolicyPartnersRequestBuilder instantiates a new CrosstenantaccesspolicyPartnersRequestBuilder and sets the default values.
func NewCrosstenantaccesspolicyPartnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrosstenantaccesspolicyPartnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCrosstenantaccesspolicyPartnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CrosstenantaccesspolicyPartnersCountRequestBuilder when successful
func (m *CrosstenantaccesspolicyPartnersRequestBuilder) Count()(*CrosstenantaccesspolicyPartnersCountRequestBuilder) {
    return NewCrosstenantaccesspolicyPartnersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of all partner configurations within a cross-tenant access policy. You can also use the $expand parameter to list the user synchronization policy for all partner configurations.
// returns a CrossTenantAccessPolicyConfigurationPartnerCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/crosstenantaccesspolicy-list-partners?view=graph-rest-1.0
func (m *CrosstenantaccesspolicyPartnersRequestBuilder) Get(ctx context.Context, requestConfiguration *CrosstenantaccesspolicyPartnersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyConfigurationPartnerCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCrossTenantAccessPolicyConfigurationPartnerCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyConfigurationPartnerCollectionResponseable), nil
}
// Post create a new partner configuration in a cross-tenant access policy.
// returns a CrossTenantAccessPolicyConfigurationPartnerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/crosstenantaccesspolicy-post-partners?view=graph-rest-1.0
func (m *CrosstenantaccesspolicyPartnersRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyConfigurationPartnerable, requestConfiguration *CrosstenantaccesspolicyPartnersRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyConfigurationPartnerable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCrossTenantAccessPolicyConfigurationPartnerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyConfigurationPartnerable), nil
}
// ToGetRequestInformation get a list of all partner configurations within a cross-tenant access policy. You can also use the $expand parameter to list the user synchronization policy for all partner configurations.
// returns a *RequestInformation when successful
func (m *CrosstenantaccesspolicyPartnersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CrosstenantaccesspolicyPartnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new partner configuration in a cross-tenant access policy.
// returns a *RequestInformation when successful
func (m *CrosstenantaccesspolicyPartnersRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyConfigurationPartnerable, requestConfiguration *CrosstenantaccesspolicyPartnersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CrosstenantaccesspolicyPartnersRequestBuilder when successful
func (m *CrosstenantaccesspolicyPartnersRequestBuilder) WithUrl(rawUrl string)(*CrosstenantaccesspolicyPartnersRequestBuilder) {
    return NewCrosstenantaccesspolicyPartnersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
