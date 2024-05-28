package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder provides operations to call the findTenantInformationByDomainName method.
type FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilderInternal instantiates a new FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder and sets the default values.
func NewFindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, domainName *string)(*FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder) {
    m := &FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/findTenantInformationByDomainName(domainName='{domainName}')", pathParameters),
    }
    if domainName != nil {
        m.BaseRequestBuilder.PathParameters["domainName"] = *domainName
    }
    return m
}
// NewFindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder instantiates a new FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder and sets the default values.
func NewFindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get given a domain name, search for a tenant and read its tenantInformation. You can use this API to validate tenant information and use the tenantId to configure cross-tenant access settings between you and the tenant.
// returns a TenantInformationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tenantrelationship-findtenantinformationbydomainname?view=graph-rest-1.0
func (m *FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder) Get(ctx context.Context, requestConfiguration *FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TenantInformationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTenantInformationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TenantInformationable), nil
}
// ToGetRequestInformation given a domain name, search for a tenant and read its tenantInformation. You can use this API to validate tenant information and use the tenantId to configure cross-tenant access settings between you and the tenant.
// returns a *RequestInformation when successful
func (m *FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder when successful
func (m *FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder) WithUrl(rawUrl string)(*FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder) {
    return NewFindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
