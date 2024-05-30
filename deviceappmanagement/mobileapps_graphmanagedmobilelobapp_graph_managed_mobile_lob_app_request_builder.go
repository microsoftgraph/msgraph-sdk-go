package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder casts the previous resource to managedMobileLobApp.
type MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilderGetQueryParameters get the items of type microsoft.graph.managedMobileLobApp in the microsoft.graph.mobileApp collection
type MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilderGetQueryParameters struct {
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
// MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilderGetQueryParameters
}
// NewMobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilderInternal instantiates a new MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder and sets the default values.
func NewMobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder) {
    m := &MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/graph.managedMobileLobApp{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder instantiates a new MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder and sets the default values.
func NewMobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MobileappsGraphmanagedmobilelobappCountRequestBuilder when successful
func (m *MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder) Count()(*MobileappsGraphmanagedmobilelobappCountRequestBuilder) {
    return NewMobileappsGraphmanagedmobilelobappCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the items of type microsoft.graph.managedMobileLobApp in the microsoft.graph.mobileApp collection
// returns a ManagedMobileLobAppCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedMobileLobAppCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedMobileLobAppCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedMobileLobAppCollectionResponseable), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.managedMobileLobApp in the microsoft.graph.mobileApp collection
// returns a *RequestInformation when successful
func (m *MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder when successful
func (m *MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder) WithUrl(rawUrl string)(*MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder) {
    return NewMobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
