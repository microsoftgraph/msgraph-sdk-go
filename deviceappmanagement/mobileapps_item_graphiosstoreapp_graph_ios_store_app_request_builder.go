package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder casts the previous resource to iosStoreApp.
type MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilderGetQueryParameters get the item of type microsoft.graph.mobileApp as microsoft.graph.iosStoreApp
type MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilderGetQueryParameters
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphiosstoreappAssignmentsRequestBuilder when successful
func (m *MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder) Assignments()(*MobileappsItemGraphiosstoreappAssignmentsRequestBuilder) {
    return NewMobileappsItemGraphiosstoreappAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphiosstoreappCategoriesRequestBuilder when successful
func (m *MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder) Categories()(*MobileappsItemGraphiosstoreappCategoriesRequestBuilder) {
    return NewMobileappsItemGraphiosstoreappCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilderInternal instantiates a new MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder) {
    m := &MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.iosStoreApp{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder instantiates a new MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.mobileApp as microsoft.graph.iosStoreApp
// returns a IosStoreAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosStoreAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIosStoreAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosStoreAppable), nil
}
// ToGetRequestInformation get the item of type microsoft.graph.mobileApp as microsoft.graph.iosStoreApp
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder when successful
func (m *MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder) {
    return NewMobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
