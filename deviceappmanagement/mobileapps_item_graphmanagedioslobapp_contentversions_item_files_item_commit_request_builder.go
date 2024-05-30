package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilder provides operations to call the commit method.
type MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilderInternal instantiates a new MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilder and sets the default values.
func NewMobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilder) {
    m := &MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.managedIOSLobApp/contentVersions/{mobileAppContent%2Did}/files/{mobileAppContentFile%2Did}/commit", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilder instantiates a new MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilder and sets the default values.
func NewMobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilderInternal(urlParams, requestAdapter)
}
// Post commits a file of a given app.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilder) Post(ctx context.Context, body MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitPostRequestBodyable, requestConfiguration *MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilderPostRequestConfiguration)(error) {
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
func (m *MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilder) ToPostRequestInformation(ctx context.Context, body MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitPostRequestBodyable, requestConfiguration *MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilder when successful
func (m *MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilder) {
    return NewMobileappsItemGraphmanagedioslobappContentversionsItemFilesItemCommitRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
