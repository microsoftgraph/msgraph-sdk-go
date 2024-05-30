package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder provides operations to manage the crossTenantAccessPolicy property of the microsoft.graph.policyRoot entity.
type CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderGetQueryParameters read the properties and relationships of a crossTenantAccessPolicy object.
type CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderGetQueryParameters
}
// CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderInternal instantiates a new CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder and sets the default values.
func NewCrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder) {
    m := &CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/crossTenantAccessPolicy{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder instantiates a new CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder and sets the default values.
func NewCrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// DefaultEscaped provides operations to manage the default property of the microsoft.graph.crossTenantAccessPolicy entity.
// returns a *CrosstenantaccesspolicyDefaultRequestBuilder when successful
func (m *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder) DefaultEscaped()(*CrosstenantaccesspolicyDefaultRequestBuilder) {
    return NewCrosstenantaccesspolicyDefaultRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property crossTenantAccessPolicy for policies
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder) Delete(ctx context.Context, requestConfiguration *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a crossTenantAccessPolicy object.
// returns a CrossTenantAccessPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/crosstenantaccesspolicy-get?view=graph-rest-1.0
func (m *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCrossTenantAccessPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyable), nil
}
// Partners provides operations to manage the partners property of the microsoft.graph.crossTenantAccessPolicy entity.
// returns a *CrosstenantaccesspolicyPartnersRequestBuilder when successful
func (m *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder) Partners()(*CrosstenantaccesspolicyPartnersRequestBuilder) {
    return NewCrosstenantaccesspolicyPartnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a cross-tenant access policy.
// returns a CrossTenantAccessPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/crosstenantaccesspolicy-update?view=graph-rest-1.0
func (m *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyable, requestConfiguration *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCrossTenantAccessPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyable), nil
}
// Templates provides operations to manage the templates property of the microsoft.graph.crossTenantAccessPolicy entity.
// returns a *CrosstenantaccesspolicyTemplatesRequestBuilder when successful
func (m *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder) Templates()(*CrosstenantaccesspolicyTemplatesRequestBuilder) {
    return NewCrosstenantaccesspolicyTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property crossTenantAccessPolicy for policies
// returns a *RequestInformation when successful
func (m *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a crossTenantAccessPolicy object.
// returns a *RequestInformation when successful
func (m *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a cross-tenant access policy.
// returns a *RequestInformation when successful
func (m *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantAccessPolicyable, requestConfiguration *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder when successful
func (m *CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder) WithUrl(rawUrl string)(*CrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder) {
    return NewCrosstenantaccesspolicyCrossTenantAccessPolicyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
