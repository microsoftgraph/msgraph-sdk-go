package homerealmdiscoverypolicies

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i9403f2966154900b82f3ba4891bee0f59f41e6ede23c4d59c8d30438f363d122 "github.com/microsoftgraph/msgraph-sdk-go/applications/item/homerealmdiscoverypolicies/count"
)

// HomeRealmDiscoveryPoliciesRequestBuilder provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.application entity.
type HomeRealmDiscoveryPoliciesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// HomeRealmDiscoveryPoliciesRequestBuilderGetQueryParameters get homeRealmDiscoveryPolicies from applications
type HomeRealmDiscoveryPoliciesRequestBuilderGetQueryParameters struct {
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
// HomeRealmDiscoveryPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HomeRealmDiscoveryPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *HomeRealmDiscoveryPoliciesRequestBuilderGetQueryParameters
}
// NewHomeRealmDiscoveryPoliciesRequestBuilderInternal instantiates a new HomeRealmDiscoveryPoliciesRequestBuilder and sets the default values.
func NewHomeRealmDiscoveryPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HomeRealmDiscoveryPoliciesRequestBuilder) {
    m := &HomeRealmDiscoveryPoliciesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/homeRealmDiscoveryPolicies{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewHomeRealmDiscoveryPoliciesRequestBuilder instantiates a new HomeRealmDiscoveryPoliciesRequestBuilder and sets the default values.
func NewHomeRealmDiscoveryPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HomeRealmDiscoveryPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHomeRealmDiscoveryPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *HomeRealmDiscoveryPoliciesRequestBuilder) Count()(*i9403f2966154900b82f3ba4891bee0f59f41e6ede23c4d59c8d30438f363d122.CountRequestBuilder) {
    return i9403f2966154900b82f3ba4891bee0f59f41e6ede23c4d59c8d30438f363d122.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get homeRealmDiscoveryPolicies from applications
func (m *HomeRealmDiscoveryPoliciesRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get homeRealmDiscoveryPolicies from applications
func (m *HomeRealmDiscoveryPoliciesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *HomeRealmDiscoveryPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get homeRealmDiscoveryPolicies from applications
func (m *HomeRealmDiscoveryPoliciesRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HomeRealmDiscoveryPolicyCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get homeRealmDiscoveryPolicies from applications
func (m *HomeRealmDiscoveryPoliciesRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *HomeRealmDiscoveryPoliciesRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HomeRealmDiscoveryPolicyCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateHomeRealmDiscoveryPolicyCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HomeRealmDiscoveryPolicyCollectionResponseable), nil
}
