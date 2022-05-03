package syncmicrosoftstoreforbusinessapps

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SyncMicrosoftStoreForBusinessAppsRequestBuilder provides operations to call the syncMicrosoftStoreForBusinessApps method.
type SyncMicrosoftStoreForBusinessAppsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SyncMicrosoftStoreForBusinessAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SyncMicrosoftStoreForBusinessAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal instantiates a new SyncMicrosoftStoreForBusinessAppsRequestBuilder and sets the default values.
func NewSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SyncMicrosoftStoreForBusinessAppsRequestBuilder) {
    m := &SyncMicrosoftStoreForBusinessAppsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/microsoft.graph.syncMicrosoftStoreForBusinessApps";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSyncMicrosoftStoreForBusinessAppsRequestBuilder instantiates a new SyncMicrosoftStoreForBusinessAppsRequestBuilder and sets the default values.
func NewSyncMicrosoftStoreForBusinessAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SyncMicrosoftStoreForBusinessAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation syncs Intune account with Microsoft Store For Business
func (m *SyncMicrosoftStoreForBusinessAppsRequestBuilder) CreatePostRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration syncs Intune account with Microsoft Store For Business
func (m *SyncMicrosoftStoreForBusinessAppsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *SyncMicrosoftStoreForBusinessAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post syncs Intune account with Microsoft Store For Business
func (m *SyncMicrosoftStoreForBusinessAppsRequestBuilder) Post()(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler syncs Intune account with Microsoft Store For Business
func (m *SyncMicrosoftStoreForBusinessAppsRequestBuilder) PostWithRequestConfigurationAndResponseHandler(requestConfiguration *SyncMicrosoftStoreForBusinessAppsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
