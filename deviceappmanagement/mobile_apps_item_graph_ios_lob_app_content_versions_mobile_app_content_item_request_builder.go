package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder provides operations to manage the contentVersions property of the microsoft.graph.mobileLobApp entity.
type MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderGetQueryParameters the list of content versions for this app.
type MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderGetQueryParameters
}
// MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderInternal instantiates a new MobileAppContentItemRequestBuilder and sets the default values.
func NewMobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder) {
    m := &MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.iosLobApp/contentVersions/{mobileAppContent%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewMobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder instantiates a new MobileAppContentItemRequestBuilder and sets the default values.
func NewMobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContainedApps provides operations to manage the containedApps property of the microsoft.graph.mobileAppContent entity.
func (m *MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder) ContainedApps()(*MobileAppsItemGraphIosLobAppContentVersionsItemContainedAppsRequestBuilder) {
    return NewMobileAppsItemGraphIosLobAppContentVersionsItemContainedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property contentVersions for deviceAppManagement
func (m *MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Files provides operations to manage the files property of the microsoft.graph.mobileAppContent entity.
func (m *MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder) Files()(*MobileAppsItemGraphIosLobAppContentVersionsItemFilesRequestBuilder) {
    return NewMobileAppsItemGraphIosLobAppContentVersionsItemFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of content versions for this app.
func (m *MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentable, requestConfiguration *MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the list of content versions for this app.
func (m *MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property contentVersions in deviceAppManagement
func (m *MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentable, requestConfiguration *MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder) WithUrl(rawUrl string)(*MobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder) {
    return NewMobileAppsItemGraphIosLobAppContentVersionsMobileAppContentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
