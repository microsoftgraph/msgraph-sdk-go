package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder provides operations to call the renewUpload method.
type MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderInternal instantiates a new MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder and sets the default values.
func NewMobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) {
    m := &MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.windowsUniversalAppX/contentVersions/{mobileAppContent%2Did}/files/{mobileAppContentFile%2Did}/renewUpload", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder instantiates a new MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder and sets the default values.
func NewMobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderInternal(urlParams, requestAdapter)
}
// Post renews the SAS URI for an application file upload.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) Post(ctx context.Context, requestConfiguration *MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderPostRequestConfiguration)(error) {
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
func (m *MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder when successful
func (m *MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder) {
    return NewMobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemRenewuploadRenewUploadRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
