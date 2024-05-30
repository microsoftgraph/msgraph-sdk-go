package identityprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder provides operations to manage the riskyServicePrincipals property of the microsoft.graph.identityProtectionRoot entity.
type RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderGetQueryParameters read the properties and relationships of a riskyServicePrincipal object.
type RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderGetQueryParameters
}
// RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderInternal instantiates a new RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder and sets the default values.
func NewRiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder) {
    m := &RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityProtection/riskyServicePrincipals/{riskyServicePrincipal%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder instantiates a new RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder and sets the default values.
func NewRiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property riskyServicePrincipals for identityProtection
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a riskyServicePrincipal object.
// returns a RiskyServicePrincipalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/riskyserviceprincipal-get?view=graph-rest-1.0
func (m *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRiskyServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalable), nil
}
// History provides operations to manage the history property of the microsoft.graph.riskyServicePrincipal entity.
// returns a *RiskyserviceprincipalsItemHistoryRequestBuilder when successful
func (m *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder) History()(*RiskyserviceprincipalsItemHistoryRequestBuilder) {
    return NewRiskyserviceprincipalsItemHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property riskyServicePrincipals in identityProtection
// returns a RiskyServicePrincipalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalable, requestConfiguration *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRiskyServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalable), nil
}
// ToDeleteRequestInformation delete navigation property riskyServicePrincipals for identityProtection
// returns a *RequestInformation when successful
func (m *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a riskyServicePrincipal object.
// returns a *RequestInformation when successful
func (m *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property riskyServicePrincipals in identityProtection
// returns a *RequestInformation when successful
func (m *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalable, requestConfiguration *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder when successful
func (m *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder) WithUrl(rawUrl string)(*RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder) {
    return NewRiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
