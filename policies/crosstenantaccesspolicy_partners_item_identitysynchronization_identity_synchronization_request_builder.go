package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder provides operations to manage the identitySynchronization property of the microsoft.graph.crossTenantAccessPolicyConfigurationPartner entity.
type CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderGetQueryParameters get the user synchronization policy of a partner-specific configuration.
type CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderGetQueryParameters
}
// CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderInternal instantiates a new CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder and sets the default values.
func NewCrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder) {
    m := &CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/crossTenantAccessPolicy/partners/{crossTenantAccessPolicyConfigurationPartner%2DtenantId}/identitySynchronization{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder instantiates a new CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder and sets the default values.
func NewCrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete the user synchronization policy for a partner-specific configuration.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/crosstenantidentitysyncpolicypartner-delete?view=graph-rest-1.0
func (m *CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder) Delete(ctx context.Context, requestConfiguration *CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the user synchronization policy of a partner-specific configuration.
// returns a CrossTenantIdentitySyncPolicyPartnerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/crosstenantidentitysyncpolicypartner-get?view=graph-rest-1.0
func (m *CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder) Get(ctx context.Context, requestConfiguration *CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantIdentitySyncPolicyPartnerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCrossTenantIdentitySyncPolicyPartnerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantIdentitySyncPolicyPartnerable), nil
}
// Put create a cross-tenant user synchronization policy for a partner-specific configuration.
// returns a CrossTenantIdentitySyncPolicyPartnerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/crosstenantaccesspolicyconfigurationpartner-put-identitysynchronization?view=graph-rest-1.0
func (m *CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder) Put(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantIdentitySyncPolicyPartnerable, requestConfiguration *CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderPutRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantIdentitySyncPolicyPartnerable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCrossTenantIdentitySyncPolicyPartnerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantIdentitySyncPolicyPartnerable), nil
}
// ToDeleteRequestInformation delete the user synchronization policy for a partner-specific configuration.
// returns a *RequestInformation when successful
func (m *CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the user synchronization policy of a partner-specific configuration.
// returns a *RequestInformation when successful
func (m *CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation create a cross-tenant user synchronization policy for a partner-specific configuration.
// returns a *RequestInformation when successful
func (m *CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder) ToPutRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CrossTenantIdentitySyncPolicyPartnerable, requestConfiguration *CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder when successful
func (m *CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder) WithUrl(rawUrl string)(*CrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder) {
    return NewCrosstenantaccesspolicyPartnersItemIdentitysynchronizationIdentitySynchronizationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
