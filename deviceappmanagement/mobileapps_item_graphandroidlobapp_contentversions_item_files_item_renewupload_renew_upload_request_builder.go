package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder provides operations to call the renewUpload method.
type MobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderInternal instantiates a new MobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder and sets the default values.
func NewMobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) {
    m := &MobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.androidLobApp/contentVersions/{mobileAppContent%2Did}/files/{mobileAppContentFile%2Did}/renewUpload", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder instantiates a new MobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder and sets the default values.
func NewMobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderInternal(urlParams, requestAdapter)
}
// Post renews the SAS URI for an application file upload.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) Post(ctx context.Context, requestConfiguration *MobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation renews the SAS URI for an application file upload.
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder when successful
func (m *MobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) {
    return NewMobileappsItemGraphandroidlobappContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
