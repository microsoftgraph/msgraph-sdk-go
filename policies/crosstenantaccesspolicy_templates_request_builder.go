package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CrosstenantaccesspolicyTemplatesRequestBuilder provides operations to manage the templates property of the microsoft.graph.crossTenantAccessPolicy entity.
type CrosstenantaccesspolicyTemplatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CrosstenantaccesspolicyTemplatesRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrosstenantaccesspolicyTemplatesRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CrosstenantaccesspolicyTemplatesRequestBuilderGetQueryParameters represents the base policy in the directory for multitenant organization settings.
type CrosstenantaccesspolicyTemplatesRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CrosstenantaccesspolicyTemplatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrosstenantaccesspolicyTemplatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CrosstenantaccesspolicyTemplatesRequestBuilderGetQueryParameters
}
// CrosstenantaccesspolicyTemplatesRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrosstenantaccesspolicyTemplatesRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCrosstenantaccesspolicyTemplatesRequestBuilderInternal instantiates a new CrosstenantaccesspolicyTemplatesRequestBuilder and sets the default values.
func NewCrosstenantaccesspolicyTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrosstenantaccesspolicyTemplatesRequestBuilder) {
    m := &CrosstenantaccesspolicyTemplatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/crossTenantAccessPolicy/templates{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCrosstenantaccesspolicyTemplatesRequestBuilder instantiates a new CrosstenantaccesspolicyTemplatesRequestBuilder and sets the default values.
func NewCrosstenantaccesspolicyTemplatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrosstenantaccesspolicyTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCrosstenantaccesspolicyTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property templates for policies
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CrosstenantaccesspolicyTemplatesRequestBuilder) Delete(ctx context.Context, requestConfiguration *CrosstenantaccesspolicyTemplatesRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents the base policy in the directory for multitenant organization settings.
// returns a PolicyTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CrosstenantaccesspolicyTemplatesRequestBuilder) Get(ctx context.Context, requestConfiguration *CrosstenantaccesspolicyTemplatesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PolicyTemplateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePolicyTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PolicyTemplateable), nil
}
// MultiTenantOrganizationIdentitySynchronization provides operations to manage the multiTenantOrganizationIdentitySynchronization property of the microsoft.graph.policyTemplate entity.
// returns a *CrosstenantaccesspolicyTemplatesMultitenantorganizationidentitysynchronizationMultiTenantOrganizationIdentitySynchronizationRequestBuilder when successful
func (m *CrosstenantaccesspolicyTemplatesRequestBuilder) MultiTenantOrganizationIdentitySynchronization()(*CrosstenantaccesspolicyTemplatesMultitenantorganizationidentitysynchronizationMultiTenantOrganizationIdentitySynchronizationRequestBuilder) {
    return NewCrosstenantaccesspolicyTemplatesMultitenantorganizationidentitysynchronizationMultiTenantOrganizationIdentitySynchronizationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MultiTenantOrganizationPartnerConfiguration provides operations to manage the multiTenantOrganizationPartnerConfiguration property of the microsoft.graph.policyTemplate entity.
// returns a *CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationMultiTenantOrganizationPartnerConfigurationRequestBuilder when successful
func (m *CrosstenantaccesspolicyTemplatesRequestBuilder) MultiTenantOrganizationPartnerConfiguration()(*CrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationMultiTenantOrganizationPartnerConfigurationRequestBuilder) {
    return NewCrosstenantaccesspolicyTemplatesMultitenantorganizationpartnerconfigurationMultiTenantOrganizationPartnerConfigurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property templates in policies
// returns a PolicyTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CrosstenantaccesspolicyTemplatesRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PolicyTemplateable, requestConfiguration *CrosstenantaccesspolicyTemplatesRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PolicyTemplateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePolicyTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PolicyTemplateable), nil
}
// ToDeleteRequestInformation delete navigation property templates for policies
// returns a *RequestInformation when successful
func (m *CrosstenantaccesspolicyTemplatesRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CrosstenantaccesspolicyTemplatesRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents the base policy in the directory for multitenant organization settings.
// returns a *RequestInformation when successful
func (m *CrosstenantaccesspolicyTemplatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CrosstenantaccesspolicyTemplatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property templates in policies
// returns a *RequestInformation when successful
func (m *CrosstenantaccesspolicyTemplatesRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PolicyTemplateable, requestConfiguration *CrosstenantaccesspolicyTemplatesRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CrosstenantaccesspolicyTemplatesRequestBuilder when successful
func (m *CrosstenantaccesspolicyTemplatesRequestBuilder) WithUrl(rawUrl string)(*CrosstenantaccesspolicyTemplatesRequestBuilder) {
    return NewCrosstenantaccesspolicyTemplatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
