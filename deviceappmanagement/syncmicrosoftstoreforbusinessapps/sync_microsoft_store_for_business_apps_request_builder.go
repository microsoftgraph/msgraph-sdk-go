package syncmicrosoftstoreforbusinessapps

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// SyncMicrosoftStoreForBusinessAppsRequestBuilder provides operations to call the syncMicrosoftStoreForBusinessApps method.
type SyncMicrosoftStoreForBusinessAppsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SyncMicrosoftStoreForBusinessAppsRequestBuilderPostOptions options for Post
type SyncMicrosoftStoreForBusinessAppsRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal instantiates a new SyncMicrosoftStoreForBusinessAppsRequestBuilder and sets the default values.
func NewSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SyncMicrosoftStoreForBusinessAppsRequestBuilder) {
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
func NewSyncMicrosoftStoreForBusinessAppsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SyncMicrosoftStoreForBusinessAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation syncs Intune account with Microsoft Store For Business
func (m *SyncMicrosoftStoreForBusinessAppsRequestBuilder) CreatePostRequestInformation(options *SyncMicrosoftStoreForBusinessAppsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
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
