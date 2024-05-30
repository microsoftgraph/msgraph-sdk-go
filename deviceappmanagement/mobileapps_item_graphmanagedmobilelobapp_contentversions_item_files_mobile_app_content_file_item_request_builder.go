package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder provides operations to manage the files property of the microsoft.graph.mobileAppContent entity.
type MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderGetQueryParameters the list of files for this app content version.
type MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderGetQueryParameters
}
// MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Commit provides operations to call the commit method.
// returns a *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemCommitRequestBuilder when successful
func (m *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) Commit()(*MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemCommitRequestBuilder) {
    return NewMobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemCommitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderInternal instantiates a new MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder and sets the default values.
func NewMobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) {
    m := &MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.managedMobileLobApp/contentVersions/{mobileAppContent%2Did}/files/{mobileAppContentFile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder instantiates a new MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder and sets the default values.
func NewMobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property files for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of files for this app content version.
// returns a MobileAppContentFileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentFileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileAppContentFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentFileable), nil
}
// Patch update the navigation property files in deviceAppManagement
// returns a MobileAppContentFileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentFileable, requestConfiguration *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentFileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileAppContentFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentFileable), nil
}
// RenewUpload provides operations to call the renewUpload method.
// returns a *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder when successful
func (m *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) RenewUpload()(*MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) {
    return NewMobileappsItemGraphmanagedmobilelobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property files for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of files for this app content version.
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property files in deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentFileable, requestConfiguration *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder when successful
func (m *MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) {
    return NewMobileappsItemGraphmanagedmobilelobappContentversionsItemFilesMobileAppContentFileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
