package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder provides operations to manage the identitySecurityDefaultsEnforcementPolicy property of the microsoft.graph.policyRoot entity.
type IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderGetQueryParameters retrieve the properties of an identitySecurityDefaultsEnforcementPolicy object.
type IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderGetQueryParameters
}
// IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderInternal instantiates a new IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder and sets the default values.
func NewIdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) {
    m := &IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/identitySecurityDefaultsEnforcementPolicy{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder instantiates a new IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder and sets the default values.
func NewIdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property identitySecurityDefaultsEnforcementPolicy for policies
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties of an identitySecurityDefaultsEnforcementPolicy object.
// returns a IdentitySecurityDefaultsEnforcementPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitysecuritydefaultsenforcementpolicy-get?view=graph-rest-1.0
func (m *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentitySecurityDefaultsEnforcementPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentitySecurityDefaultsEnforcementPolicyable), nil
}
// Patch update the properties of an identitySecurityDefaultsEnforcementPolicy object.
// returns a IdentitySecurityDefaultsEnforcementPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitysecuritydefaultsenforcementpolicy-update?view=graph-rest-1.0
func (m *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentitySecurityDefaultsEnforcementPolicyable, requestConfiguration *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentitySecurityDefaultsEnforcementPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentitySecurityDefaultsEnforcementPolicyable), nil
}
// ToDeleteRequestInformation delete navigation property identitySecurityDefaultsEnforcementPolicy for policies
// returns a *RequestInformation when successful
func (m *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties of an identitySecurityDefaultsEnforcementPolicy object.
// returns a *RequestInformation when successful
func (m *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an identitySecurityDefaultsEnforcementPolicy object.
// returns a *RequestInformation when successful
func (m *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentitySecurityDefaultsEnforcementPolicyable, requestConfiguration *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder when successful
func (m *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) WithUrl(rawUrl string)(*IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) {
    return NewIdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
