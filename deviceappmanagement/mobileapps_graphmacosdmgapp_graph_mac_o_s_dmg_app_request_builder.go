package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder casts the previous resource to macOSDmgApp.
type MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilderGetQueryParameters get the items of type microsoft.graph.macOSDmgApp in the microsoft.graph.mobileApp collection
type MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilderGetQueryParameters struct {
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
// MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilderGetQueryParameters
}
// NewMobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilderInternal instantiates a new MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder and sets the default values.
func NewMobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder) {
    m := &MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/graph.macOSDmgApp{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder instantiates a new MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder and sets the default values.
func NewMobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MobileappsGraphmacosdmgappCountRequestBuilder when successful
func (m *MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder) Count()(*MobileappsGraphmacosdmgappCountRequestBuilder) {
    return NewMobileappsGraphmacosdmgappCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the items of type microsoft.graph.macOSDmgApp in the microsoft.graph.mobileApp collection
// returns a MacOSDmgAppCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MacOSDmgAppCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMacOSDmgAppCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MacOSDmgAppCollectionResponseable), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.macOSDmgApp in the microsoft.graph.mobileApp collection
// returns a *RequestInformation when successful
func (m *MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder when successful
func (m *MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder) WithUrl(rawUrl string)(*MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder) {
    return NewMobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
