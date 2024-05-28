package domains

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemFederationconfigurationFederationConfigurationRequestBuilder provides operations to manage the federationConfiguration property of the microsoft.graph.domain entity.
type ItemFederationconfigurationFederationConfigurationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemFederationconfigurationFederationConfigurationRequestBuilderGetQueryParameters read the properties of the internalDomainFederation objects for the domain. This API returns only one object in the collection.
type ItemFederationconfigurationFederationConfigurationRequestBuilderGetQueryParameters struct {
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
// ItemFederationconfigurationFederationConfigurationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFederationconfigurationFederationConfigurationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemFederationconfigurationFederationConfigurationRequestBuilderGetQueryParameters
}
// ItemFederationconfigurationFederationConfigurationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFederationconfigurationFederationConfigurationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByInternalDomainFederationId provides operations to manage the federationConfiguration property of the microsoft.graph.domain entity.
// returns a *ItemFederationconfigurationInternalDomainFederationItemRequestBuilder when successful
func (m *ItemFederationconfigurationFederationConfigurationRequestBuilder) ByInternalDomainFederationId(internalDomainFederationId string)(*ItemFederationconfigurationInternalDomainFederationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if internalDomainFederationId != "" {
        urlTplParams["internalDomainFederation%2Did"] = internalDomainFederationId
    }
    return NewItemFederationconfigurationInternalDomainFederationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemFederationconfigurationFederationConfigurationRequestBuilderInternal instantiates a new ItemFederationconfigurationFederationConfigurationRequestBuilder and sets the default values.
func NewItemFederationconfigurationFederationConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFederationconfigurationFederationConfigurationRequestBuilder) {
    m := &ItemFederationconfigurationFederationConfigurationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/domains/{domain%2Did}/federationConfiguration{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemFederationconfigurationFederationConfigurationRequestBuilder instantiates a new ItemFederationconfigurationFederationConfigurationRequestBuilder and sets the default values.
func NewItemFederationconfigurationFederationConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFederationconfigurationFederationConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFederationconfigurationFederationConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemFederationconfigurationCountRequestBuilder when successful
func (m *ItemFederationconfigurationFederationConfigurationRequestBuilder) Count()(*ItemFederationconfigurationCountRequestBuilder) {
    return NewItemFederationconfigurationCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties of the internalDomainFederation objects for the domain. This API returns only one object in the collection.
// returns a InternalDomainFederationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/domain-list-federationconfiguration?view=graph-rest-1.0
func (m *ItemFederationconfigurationFederationConfigurationRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemFederationconfigurationFederationConfigurationRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InternalDomainFederationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateInternalDomainFederationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InternalDomainFederationCollectionResponseable), nil
}
// Post create a new internalDomainFederation object.
// returns a InternalDomainFederationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/domain-post-federationconfiguration?view=graph-rest-1.0
func (m *ItemFederationconfigurationFederationConfigurationRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InternalDomainFederationable, requestConfiguration *ItemFederationconfigurationFederationConfigurationRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InternalDomainFederationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateInternalDomainFederationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InternalDomainFederationable), nil
}
// ToGetRequestInformation read the properties of the internalDomainFederation objects for the domain. This API returns only one object in the collection.
// returns a *RequestInformation when successful
func (m *ItemFederationconfigurationFederationConfigurationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemFederationconfigurationFederationConfigurationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new internalDomainFederation object.
// returns a *RequestInformation when successful
func (m *ItemFederationconfigurationFederationConfigurationRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InternalDomainFederationable, requestConfiguration *ItemFederationconfigurationFederationConfigurationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemFederationconfigurationFederationConfigurationRequestBuilder when successful
func (m *ItemFederationconfigurationFederationConfigurationRequestBuilder) WithUrl(rawUrl string)(*ItemFederationconfigurationFederationConfigurationRequestBuilder) {
    return NewItemFederationconfigurationFederationConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
