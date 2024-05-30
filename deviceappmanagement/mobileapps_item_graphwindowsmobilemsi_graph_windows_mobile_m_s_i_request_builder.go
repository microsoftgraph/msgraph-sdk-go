package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder casts the previous resource to windowsMobileMSI.
type MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderGetQueryParameters get the item of type microsoft.graph.mobileApp as microsoft.graph.windowsMobileMSI
type MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderGetQueryParameters
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphwindowsmobilemsiAssignmentsRequestBuilder when successful
func (m *MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) Assignments()(*MobileappsItemGraphwindowsmobilemsiAssignmentsRequestBuilder) {
    return NewMobileappsItemGraphwindowsmobilemsiAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphwindowsmobilemsiCategoriesRequestBuilder when successful
func (m *MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) Categories()(*MobileappsItemGraphwindowsmobilemsiCategoriesRequestBuilder) {
    return NewMobileappsItemGraphwindowsmobilemsiCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderInternal instantiates a new MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder and sets the default values.
func NewMobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) {
    m := &MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.windowsMobileMSI{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder instantiates a new MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder and sets the default values.
func NewMobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentVersions provides operations to manage the contentVersions property of the microsoft.graph.mobileLobApp entity.
// returns a *MobileappsItemGraphwindowsmobilemsiContentversionsContentVersionsRequestBuilder when successful
func (m *MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) ContentVersions()(*MobileappsItemGraphwindowsmobilemsiContentversionsContentVersionsRequestBuilder) {
    return NewMobileappsItemGraphwindowsmobilemsiContentversionsContentVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the item of type microsoft.graph.mobileApp as microsoft.graph.windowsMobileMSI
// returns a WindowsMobileMSIable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsMobileMSIable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsMobileMSIFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsMobileMSIable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.mobileApp as microsoft.graph.windowsMobileMSI
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder when successful
func (m *MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) {
    return NewMobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
