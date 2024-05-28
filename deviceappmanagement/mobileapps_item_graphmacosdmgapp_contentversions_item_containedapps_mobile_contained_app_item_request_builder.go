package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder provides operations to manage the containedApps property of the microsoft.graph.mobileAppContent entity.
type MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderGetQueryParameters the collection of contained apps in a MobileLobApp acting as a package.
type MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderGetQueryParameters
}
// MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderInternal instantiates a new MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder and sets the default values.
func NewMobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder) {
    m := &MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.macOSDmgApp/contentVersions/{mobileAppContent%2Did}/containedApps/{mobileContainedApp%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder instantiates a new MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder and sets the default values.
func NewMobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property containedApps for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of contained apps in a MobileLobApp acting as a package.
// returns a MobileContainedAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileContainedAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileContainedAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileContainedAppable), nil
}
// Patch update the navigation property containedApps in deviceAppManagement
// returns a MobileContainedAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileContainedAppable, requestConfiguration *MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileContainedAppable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileContainedAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileContainedAppable), nil
}
// ToDeleteRequestInformation delete navigation property containedApps for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of contained apps in a MobileLobApp acting as a package.
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property containedApps in deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileContainedAppable, requestConfiguration *MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder when successful
func (m *MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder) {
    return NewMobileappsItemGraphmacosdmgappContentversionsItemContainedappsMobileContainedAppItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
