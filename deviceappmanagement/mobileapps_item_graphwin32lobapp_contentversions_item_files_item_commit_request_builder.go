package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilder provides operations to call the commit method.
type MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilderInternal instantiates a new MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilder and sets the default values.
func NewMobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilder) {
    m := &MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.win32LobApp/contentVersions/{mobileAppContent%2Did}/files/{mobileAppContentFile%2Did}/commit", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilder instantiates a new MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilder and sets the default values.
func NewMobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilderInternal(urlParams, requestAdapter)
}
// Post commits a file of a given app.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilder) Post(ctx context.Context, body MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitPostRequestBodyable, requestConfiguration *MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilderPostRequestConfiguration)(error) {
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
func (m *MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilder) ToPostRequestInformation(ctx context.Context, body MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitPostRequestBodyable, requestConfiguration *MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilder when successful
func (m *MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilder) {
    return NewMobileappsItemGraphwin32lobappContentversionsItemFilesItemCommitRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
