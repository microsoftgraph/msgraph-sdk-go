package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder provides operations to manage the files property of the microsoft.graph.mobileAppContent entity.
type MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderGetQueryParameters the list of files for this app content version.
type MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderGetQueryParameters
}
// MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Commit provides operations to call the commit method.
// returns a *MobileappsItemGraphmacosdmgappContentversionsItemFilesItemCommitRequestBuilder when successful
func (m *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) Commit()(*MobileappsItemGraphmacosdmgappContentversionsItemFilesItemCommitRequestBuilder) {
    return NewMobileappsItemGraphmacosdmgappContentversionsItemFilesItemCommitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderInternal instantiates a new MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder and sets the default values.
func NewMobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) {
    m := &MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.macOSDmgApp/contentVersions/{mobileAppContent%2Did}/files/{mobileAppContentFile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder instantiates a new MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder and sets the default values.
func NewMobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property files for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentFileable, error) {
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
func (m *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentFileable, requestConfiguration *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentFileable, error) {
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
// returns a *MobileappsItemGraphmacosdmgappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder when successful
func (m *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) RenewUpload()(*MobileappsItemGraphmacosdmgappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) {
    return NewMobileappsItemGraphmacosdmgappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property files for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppContentFileable, requestConfiguration *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder when successful
func (m *MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder) {
    return NewMobileappsItemGraphmacosdmgappContentversionsItemFilesMobileAppContentFileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
