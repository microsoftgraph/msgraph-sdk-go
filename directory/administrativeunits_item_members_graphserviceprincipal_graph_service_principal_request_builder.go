package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder casts the previous resource to servicePrincipal.
type AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilderGetQueryParameters get the items of type microsoft.graph.servicePrincipal in the microsoft.graph.directoryObject collection
type AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilderGetQueryParameters struct {
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
// AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilderGetQueryParameters
}
// NewAdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal instantiates a new AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder and sets the default values.
func NewAdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    m := &AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/administrativeUnits/{administrativeUnit%2Did}/members/graph.servicePrincipal{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder instantiates a new AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder and sets the default values.
func NewAdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AdministrativeunitsItemMembersGraphserviceprincipalCountRequestBuilder when successful
func (m *AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder) Count()(*AdministrativeunitsItemMembersGraphserviceprincipalCountRequestBuilder) {
    return NewAdministrativeunitsItemMembersGraphserviceprincipalCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the items of type microsoft.graph.servicePrincipal in the microsoft.graph.directoryObject collection
// returns a ServicePrincipalCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder) Get(ctx context.Context, requestConfiguration *AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServicePrincipalCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalCollectionResponseable), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.servicePrincipal in the microsoft.graph.directoryObject collection
// returns a *RequestInformation when successful
func (m *AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder) WithUrl(rawUrl string)(*AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewAdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
