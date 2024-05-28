package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder provides operations to manage the contentVersions property of the microsoft.graph.managedMobileLobApp entity.
type MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderGetQueryParameters the list of content versions for this app.
type MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderGetQueryParameters
}
// MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderInternal instantiates a new MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder and sets the default values.
func NewMobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder) {
    m := &MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.managedIOSLobApp/contentVersions/{mobileAppContent%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder instantiates a new MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder and sets the default values.
func NewMobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContainedApps provides operations to manage the containedApps property of the microsoft.graph.mobileAppContent entity.
// returns a *MobileappsItemGraphmanagedioslobappContentversionsItemContainedappsContainedAppsRequestBuilder when successful
func (m *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder) ContainedApps()(*MobileappsItemGraphmanagedioslobappContentversionsItemContainedappsContainedAppsRequestBuilder) {
    return NewMobileappsItemGraphmanagedioslobappContentversionsItemContainedappsContainedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property contentVersions for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Files provides operations to manage the files property of the microsoft.graph.mobileAppContent entity.
// returns a *MobileappsItemGraphmanagedioslobappContentversionsItemFilesRequestBuilder when successful
func (m *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder) Files()(*MobileappsItemGraphmanagedioslobappContentversionsItemFilesRequestBuilder) {
    return NewMobileappsItemGraphmanagedioslobappContentversionsItemFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of content versions for this app.
// returns a MobileAppContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileAppContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentable), nil
}
// Patch update the navigation property contentVersions in deviceAppManagement
// returns a MobileAppContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentable, requestConfiguration *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileAppContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentable), nil
}
// ToDeleteRequestInformation delete navigation property contentVersions for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of content versions for this app.
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property contentVersions in deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentable, requestConfiguration *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder when successful
func (m *MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder) {
    return NewMobileappsItemGraphmanagedioslobappContentversionsMobileAppContentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
