package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.servicePrincipal entity.
type ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderGetQueryParameters federated identities for a specific type of service principal - managed identity. Supports $expand and $filter (/$count eq 0, /$count ne 0).
type ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderGetQueryParameters
}
// ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderInternal instantiates a new ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder and sets the default values.
func NewItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, name *string)(*ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder) {
    m := &ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/federatedIdentityCredentials(name='{name}'){?%24expand,%24select}", pathParameters),
    }
    if name != nil {
        m.BaseRequestBuilder.PathParameters["name"] = *name
    }
    return m
}
// NewItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder instantiates a new ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder and sets the default values.
func NewItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete navigation property federatedIdentityCredentials for servicePrincipals
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get federated identities for a specific type of service principal - managed identity. Supports $expand and $filter (/$count eq 0, /$count ne 0).
// returns a FederatedIdentityCredentialable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FederatedIdentityCredentialable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateFederatedIdentityCredentialFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FederatedIdentityCredentialable), nil
}
// Patch update the navigation property federatedIdentityCredentials in servicePrincipals
// returns a FederatedIdentityCredentialable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FederatedIdentityCredentialable, requestConfiguration *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FederatedIdentityCredentialable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateFederatedIdentityCredentialFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FederatedIdentityCredentialable), nil
}
// ToDeleteRequestInformation delete navigation property federatedIdentityCredentials for servicePrincipals
// returns a *RequestInformation when successful
func (m *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation federated identities for a specific type of service principal - managed identity. Supports $expand and $filter (/$count eq 0, /$count ne 0).
// returns a *RequestInformation when successful
func (m *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property federatedIdentityCredentials in servicePrincipals
// returns a *RequestInformation when successful
func (m *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FederatedIdentityCredentialable, requestConfiguration *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder when successful
func (m *ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder) WithUrl(rawUrl string)(*ItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder) {
    return NewItemFederatedidentitycredentialswithnameFederatedIdentityCredentialsWithNameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
