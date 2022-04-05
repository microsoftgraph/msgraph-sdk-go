package syncmicrosoftstoreforbusinessapps

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SyncMicrosoftStoreForBusinessAppsRequestBuilder provides operations to call the syncMicrosoftStoreForBusinessApps method.
type SyncMicrosoftStoreForBusinessAppsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SyncMicrosoftStoreForBusinessAppsRequestBuilderPostOptions options for Post
type SyncMicrosoftStoreForBusinessAppsRequestBuilderPostOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
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
func (m *SyncMicrosoftStoreForBusinessAppsRequestBuilder) CreatePostRequestInformation(options *SyncMicrosoftStoreForBusinessAppsRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Post syncs Intune account with Microsoft Store For Business
func (m *SyncMicrosoftStoreForBusinessAppsRequestBuilder) Post(options *SyncMicrosoftStoreForBusinessAppsRequestBuilderPostOptions)(error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
