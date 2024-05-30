package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsGraphioslobappGraphIosLobAppRequestBuilder casts the previous resource to iosLobApp.
type MobileappsGraphioslobappGraphIosLobAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsGraphioslobappGraphIosLobAppRequestBuilderGetQueryParameters get the items of type microsoft.graph.iosLobApp in the microsoft.graph.mobileApp collection
type MobileappsGraphioslobappGraphIosLobAppRequestBuilderGetQueryParameters struct {
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
// MobileappsGraphioslobappGraphIosLobAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsGraphioslobappGraphIosLobAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsGraphioslobappGraphIosLobAppRequestBuilderGetQueryParameters
}
// NewMobileappsGraphioslobappGraphIosLobAppRequestBuilderInternal instantiates a new MobileappsGraphioslobappGraphIosLobAppRequestBuilder and sets the default values.
func NewMobileappsGraphioslobappGraphIosLobAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsGraphioslobappGraphIosLobAppRequestBuilder) {
    m := &MobileappsGraphioslobappGraphIosLobAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/graph.iosLobApp{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMobileappsGraphioslobappGraphIosLobAppRequestBuilder instantiates a new MobileappsGraphioslobappGraphIosLobAppRequestBuilder and sets the default values.
func NewMobileappsGraphioslobappGraphIosLobAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsGraphioslobappGraphIosLobAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsGraphioslobappGraphIosLobAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MobileappsGraphioslobappCountRequestBuilder when successful
func (m *MobileappsGraphioslobappGraphIosLobAppRequestBuilder) Count()(*MobileappsGraphioslobappCountRequestBuilder) {
    return NewMobileappsGraphioslobappCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the items of type microsoft.graph.iosLobApp in the microsoft.graph.mobileApp collection
// returns a IosLobAppCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsGraphioslobappGraphIosLobAppRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsGraphioslobappGraphIosLobAppRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosLobAppCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIosLobAppCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosLobAppCollectionResponseable), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.iosLobApp in the microsoft.graph.mobileApp collection
// returns a *RequestInformation when successful
func (m *MobileappsGraphioslobappGraphIosLobAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsGraphioslobappGraphIosLobAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsGraphioslobappGraphIosLobAppRequestBuilder when successful
func (m *MobileappsGraphioslobappGraphIosLobAppRequestBuilder) WithUrl(rawUrl string)(*MobileappsGraphioslobappGraphIosLobAppRequestBuilder) {
    return NewMobileappsGraphioslobappGraphIosLobAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
