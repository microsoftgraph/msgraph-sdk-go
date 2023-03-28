package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SyncMicrosoftStoreForBusinessAppsRequestBuilder provides operations to call the syncMicrosoftStoreForBusinessApps method.
type SyncMicrosoftStoreForBusinessAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SyncMicrosoftStoreForBusinessAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SyncMicrosoftStoreForBusinessAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal instantiates a new SyncMicrosoftStoreForBusinessAppsRequestBuilder and sets the default values.
func NewSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SyncMicrosoftStoreForBusinessAppsRequestBuilder) {
    m := &SyncMicrosoftStoreForBusinessAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/syncMicrosoftStoreForBusinessApps", pathParameters),
    }
    return m
}
// NewSyncMicrosoftStoreForBusinessAppsRequestBuilder instantiates a new SyncMicrosoftStoreForBusinessAppsRequestBuilder and sets the default values.
func NewSyncMicrosoftStoreForBusinessAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SyncMicrosoftStoreForBusinessAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post syncs Intune account with Microsoft Store For Business
func (m *SyncMicrosoftStoreForBusinessAppsRequestBuilder) Post(ctx context.Context, requestConfiguration *SyncMicrosoftStoreForBusinessAppsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation syncs Intune account with Microsoft Store For Business
func (m *SyncMicrosoftStoreForBusinessAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *SyncMicrosoftStoreForBusinessAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
