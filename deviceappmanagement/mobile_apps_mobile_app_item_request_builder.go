package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileAppsMobileAppItemRequestBuilder provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
type MobileAppsMobileAppItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsMobileAppItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsMobileAppItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileAppsMobileAppItemRequestBuilderGetQueryParameters read properties and relationships of the managedAndroidStoreApp object.
type MobileAppsMobileAppItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppsMobileAppItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsMobileAppItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsMobileAppItemRequestBuilderGetQueryParameters
}
// MobileAppsMobileAppItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsMobileAppItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *MobileAppsMobileAppItemRequestBuilder) Assign()(*MobileAppsItemAssignRequestBuilder) {
    return NewMobileAppsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsMobileAppItemRequestBuilder) Assignments()(*MobileAppsItemAssignmentsRequestBuilder) {
    return NewMobileAppsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsMobileAppItemRequestBuilder) Categories()(*MobileAppsItemCategoriesRequestBuilder) {
    return NewMobileAppsItemCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileAppsMobileAppItemRequestBuilderInternal instantiates a new MobileAppItemRequestBuilder and sets the default values.
func NewMobileAppsMobileAppItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsMobileAppItemRequestBuilder) {
    m := &MobileAppsMobileAppItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewMobileAppsMobileAppItemRequestBuilder instantiates a new MobileAppItemRequestBuilder and sets the default values.
func NewMobileAppsMobileAppItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsMobileAppItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsMobileAppItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a macOSOfficeSuiteApp.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-apps-macosofficesuiteapp-delete?view=graph-rest-1.0
func (m *MobileAppsMobileAppItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileAppsMobileAppItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read properties and relationships of the managedAndroidStoreApp object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-apps-managedandroidstoreapp-get?view=graph-rest-1.0
func (m *MobileAppsMobileAppItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsMobileAppItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppable), nil
}
// GraphAndroidLobApp casts the previous resource to androidLobApp.
func (m *MobileAppsMobileAppItemRequestBuilder) GraphAndroidLobApp()(*MobileAppsItemGraphAndroidLobAppRequestBuilder) {
    return NewMobileAppsItemGraphAndroidLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphAndroidStoreApp casts the previous resource to androidStoreApp.
func (m *MobileAppsMobileAppItemRequestBuilder) GraphAndroidStoreApp()(*MobileAppsItemGraphAndroidStoreAppRequestBuilder) {
    return NewMobileAppsItemGraphAndroidStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphIosLobApp casts the previous resource to iosLobApp.
func (m *MobileAppsMobileAppItemRequestBuilder) GraphIosLobApp()(*MobileAppsItemGraphIosLobAppRequestBuilder) {
    return NewMobileAppsItemGraphIosLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphIosStoreApp casts the previous resource to iosStoreApp.
func (m *MobileAppsMobileAppItemRequestBuilder) GraphIosStoreApp()(*MobileAppsItemGraphIosStoreAppRequestBuilder) {
    return NewMobileAppsItemGraphIosStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphIosVppApp casts the previous resource to iosVppApp.
func (m *MobileAppsMobileAppItemRequestBuilder) GraphIosVppApp()(*MobileAppsItemGraphIosVppAppRequestBuilder) {
    return NewMobileAppsItemGraphIosVppAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMacOSDmgApp casts the previous resource to macOSDmgApp.
func (m *MobileAppsMobileAppItemRequestBuilder) GraphMacOSDmgApp()(*MobileAppsItemGraphMacOSDmgAppRequestBuilder) {
    return NewMobileAppsItemGraphMacOSDmgAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMacOSLobApp casts the previous resource to macOSLobApp.
func (m *MobileAppsMobileAppItemRequestBuilder) GraphMacOSLobApp()(*MobileAppsItemGraphMacOSLobAppRequestBuilder) {
    return NewMobileAppsItemGraphMacOSLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphManagedAndroidLobApp casts the previous resource to managedAndroidLobApp.
func (m *MobileAppsMobileAppItemRequestBuilder) GraphManagedAndroidLobApp()(*MobileAppsItemGraphManagedAndroidLobAppRequestBuilder) {
    return NewMobileAppsItemGraphManagedAndroidLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphManagedIOSLobApp casts the previous resource to managedIOSLobApp.
func (m *MobileAppsMobileAppItemRequestBuilder) GraphManagedIOSLobApp()(*MobileAppsItemGraphManagedIOSLobAppRequestBuilder) {
    return NewMobileAppsItemGraphManagedIOSLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphManagedMobileLobApp casts the previous resource to managedMobileLobApp.
func (m *MobileAppsMobileAppItemRequestBuilder) GraphManagedMobileLobApp()(*MobileAppsItemGraphManagedMobileLobAppRequestBuilder) {
    return NewMobileAppsItemGraphManagedMobileLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMicrosoftStoreForBusinessApp casts the previous resource to microsoftStoreForBusinessApp.
func (m *MobileAppsMobileAppItemRequestBuilder) GraphMicrosoftStoreForBusinessApp()(*MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder) {
    return NewMobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWin32LobApp casts the previous resource to win32LobApp.
func (m *MobileAppsMobileAppItemRequestBuilder) GraphWin32LobApp()(*MobileAppsItemGraphWin32LobAppRequestBuilder) {
    return NewMobileAppsItemGraphWin32LobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsAppX casts the previous resource to windowsAppX.
func (m *MobileAppsMobileAppItemRequestBuilder) GraphWindowsAppX()(*MobileAppsItemGraphWindowsAppXRequestBuilder) {
    return NewMobileAppsItemGraphWindowsAppXRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsMobileMSI casts the previous resource to windowsMobileMSI.
func (m *MobileAppsMobileAppItemRequestBuilder) GraphWindowsMobileMSI()(*MobileAppsItemGraphWindowsMobileMSIRequestBuilder) {
    return NewMobileAppsItemGraphWindowsMobileMSIRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsUniversalAppX casts the previous resource to windowsUniversalAppX.
func (m *MobileAppsMobileAppItemRequestBuilder) GraphWindowsUniversalAppX()(*MobileAppsItemGraphWindowsUniversalAppXRequestBuilder) {
    return NewMobileAppsItemGraphWindowsUniversalAppXRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsWebApp casts the previous resource to windowsWebApp.
func (m *MobileAppsMobileAppItemRequestBuilder) GraphWindowsWebApp()(*MobileAppsItemGraphWindowsWebAppRequestBuilder) {
    return NewMobileAppsItemGraphWindowsWebAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a macOSLobApp object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-apps-macoslobapp-update?view=graph-rest-1.0
func (m *MobileAppsMobileAppItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppable, requestConfiguration *MobileAppsMobileAppItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppable), nil
}
// ToDeleteRequestInformation deletes a macOSOfficeSuiteApp.
func (m *MobileAppsMobileAppItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileAppsMobileAppItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the managedAndroidStoreApp object.
func (m *MobileAppsMobileAppItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsMobileAppItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a macOSLobApp object.
func (m *MobileAppsMobileAppItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppable, requestConfiguration *MobileAppsMobileAppItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MobileAppsMobileAppItemRequestBuilder) WithUrl(rawUrl string)(*MobileAppsMobileAppItemRequestBuilder) {
    return NewMobileAppsMobileAppItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
