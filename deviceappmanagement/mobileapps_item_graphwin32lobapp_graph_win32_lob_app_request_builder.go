package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder casts the previous resource to win32LobApp.
type MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilderGetQueryParameters get the item of type microsoft.graph.mobileApp as microsoft.graph.win32LobApp
type MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilderGetQueryParameters
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphwin32lobappAssignmentsRequestBuilder when successful
func (m *MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder) Assignments()(*MobileappsItemGraphwin32lobappAssignmentsRequestBuilder) {
    return NewMobileappsItemGraphwin32lobappAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphwin32lobappCategoriesRequestBuilder when successful
func (m *MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder) Categories()(*MobileappsItemGraphwin32lobappCategoriesRequestBuilder) {
    return NewMobileappsItemGraphwin32lobappCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilderInternal instantiates a new MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder) {
    m := &MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.win32LobApp{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder instantiates a new MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentVersions provides operations to manage the contentVersions property of the microsoft.graph.mobileLobApp entity.
// returns a *MobileappsItemGraphwin32lobappContentversionsContentVersionsRequestBuilder when successful
func (m *MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder) ContentVersions()(*MobileappsItemGraphwin32lobappContentversionsContentVersionsRequestBuilder) {
    return NewMobileappsItemGraphwin32lobappContentversionsContentVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the item of type microsoft.graph.mobileApp as microsoft.graph.win32LobApp
// returns a Win32LobAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Win32LobAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWin32LobAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Win32LobAppable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.mobileApp as microsoft.graph.win32LobApp
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder when successful
func (m *MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder) {
    return NewMobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
