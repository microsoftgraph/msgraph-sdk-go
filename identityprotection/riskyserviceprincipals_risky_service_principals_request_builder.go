package identityprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder provides operations to manage the riskyServicePrincipals property of the microsoft.graph.identityProtectionRoot entity.
type RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilderGetQueryParameters retrieve the properties and relationships of riskyServicePrincipal objects.
type RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilderGetQueryParameters struct {
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
// RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilderGetQueryParameters
}
// RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByRiskyServicePrincipalId provides operations to manage the riskyServicePrincipals property of the microsoft.graph.identityProtectionRoot entity.
// returns a *RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder when successful
func (m *RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder) ByRiskyServicePrincipalId(riskyServicePrincipalId string)(*RiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if riskyServicePrincipalId != "" {
        urlTplParams["riskyServicePrincipal%2Did"] = riskyServicePrincipalId
    }
    return NewRiskyserviceprincipalsRiskyServicePrincipalItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ConfirmCompromised provides operations to call the confirmCompromised method.
// returns a *RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder when successful
func (m *RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder) ConfirmCompromised()(*RiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilder) {
    return NewRiskyserviceprincipalsConfirmcompromisedConfirmCompromisedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewRiskyserviceprincipalsRiskyServicePrincipalsRequestBuilderInternal instantiates a new RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder and sets the default values.
func NewRiskyserviceprincipalsRiskyServicePrincipalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder) {
    m := &RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityProtection/riskyServicePrincipals{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewRiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder instantiates a new RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder and sets the default values.
func NewRiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyserviceprincipalsRiskyServicePrincipalsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *RiskyserviceprincipalsCountRequestBuilder when successful
func (m *RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder) Count()(*RiskyserviceprincipalsCountRequestBuilder) {
    return NewRiskyserviceprincipalsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dismiss provides operations to call the dismiss method.
// returns a *RiskyserviceprincipalsDismissRequestBuilder when successful
func (m *RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder) Dismiss()(*RiskyserviceprincipalsDismissRequestBuilder) {
    return NewRiskyserviceprincipalsDismissRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the properties and relationships of riskyServicePrincipal objects.
// returns a RiskyServicePrincipalCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityprotectionroot-list-riskyserviceprincipals?view=graph-rest-1.0
func (m *RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder) Get(ctx context.Context, requestConfiguration *RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRiskyServicePrincipalCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalCollectionResponseable), nil
}
// Post create new navigation property to riskyServicePrincipals for identityProtection
// returns a RiskyServicePrincipalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalable, requestConfiguration *RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation retrieve the properties and relationships of riskyServicePrincipal objects.
// returns a *RequestInformation when successful
func (m *RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to riskyServicePrincipals for identityProtection
// returns a *RequestInformation when successful
func (m *RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RiskyServicePrincipalable, requestConfiguration *RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder when successful
func (m *RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder) WithUrl(rawUrl string)(*RiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder) {
    return NewRiskyserviceprincipalsRiskyServicePrincipalsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
