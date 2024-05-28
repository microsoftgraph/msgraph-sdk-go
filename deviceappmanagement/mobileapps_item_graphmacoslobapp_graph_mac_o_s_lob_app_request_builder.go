package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder casts the previous resource to macOSLobApp.
type MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilderGetQueryParameters get the item of type microsoft.graph.mobileApp as microsoft.graph.macOSLobApp
type MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilderGetQueryParameters
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphmacoslobappAssignmentsRequestBuilder when successful
func (m *MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder) Assignments()(*MobileappsItemGraphmacoslobappAssignmentsRequestBuilder) {
    return NewMobileappsItemGraphmacoslobappAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphmacoslobappCategoriesRequestBuilder when successful
func (m *MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder) Categories()(*MobileappsItemGraphmacoslobappCategoriesRequestBuilder) {
    return NewMobileappsItemGraphmacoslobappCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilderInternal instantiates a new MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder) {
    m := &MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.macOSLobApp{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder instantiates a new MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentVersions provides operations to manage the contentVersions property of the microsoft.graph.mobileLobApp entity.
// returns a *MobileappsItemGraphmacoslobappContentversionsContentVersionsRequestBuilder when successful
func (m *MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder) ContentVersions()(*MobileappsItemGraphmacoslobappContentversionsContentVersionsRequestBuilder) {
    return NewMobileappsItemGraphmacoslobappContentversionsContentVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the item of type microsoft.graph.mobileApp as microsoft.graph.macOSLobApp
// returns a MacOSLobAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MacOSLobAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMacOSLobAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MacOSLobAppable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.mobileApp as microsoft.graph.macOSLobApp
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder when successful
func (m *MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder) {
    return NewMobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
