package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilder provides operations to call the commit method.
type MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilderInternal instantiates a new MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilder and sets the default values.
func NewMobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilder) {
    m := &MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.windowsUniversalAppX/contentVersions/{mobileAppContent%2Did}/files/{mobileAppContentFile%2Did}/commit", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilder instantiates a new MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilder and sets the default values.
func NewMobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilderInternal(urlParams, requestAdapter)
}
// Post commits a file of a given app.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilder) Post(ctx context.Context, body MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitPostRequestBodyable, requestConfiguration *MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation commits a file of a given app.
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilder) ToPostRequestInformation(ctx context.Context, body MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitPostRequestBodyable, requestConfiguration *MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilder when successful
func (m *MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilder) {
    return NewMobileappsItemGraphwindowsuniversalappxContentversionsItemFilesItemCommitRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
