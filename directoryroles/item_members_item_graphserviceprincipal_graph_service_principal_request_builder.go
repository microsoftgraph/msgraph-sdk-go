package directoryroles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder casts the previous resource to servicePrincipal.
type ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
type ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilderGetQueryParameters
}
// NewItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal instantiates a new ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder and sets the default values.
func NewItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    m := &ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directoryRoles/{directoryRole%2Did}/members/{directoryObject%2Did}/graph.servicePrincipal{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder instantiates a new ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder and sets the default values.
func NewItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
// returns a ServicePrincipalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
// returns a *RequestInformation when successful
func (m *ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder) WithUrl(rawUrl string)(*ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
