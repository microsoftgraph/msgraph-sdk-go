package bypassactivationlock

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// BypassActivationLockRequestBuilder provides operations to call the bypassActivationLock method.
type BypassActivationLockRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BypassActivationLockRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BypassActivationLockRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBypassActivationLockRequestBuilderInternal instantiates a new BypassActivationLockRequestBuilder and sets the default values.
func NewBypassActivationLockRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BypassActivationLockRequestBuilder) {
    m := &BypassActivationLockRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/microsoft.graph.bypassActivationLock";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBypassActivationLockRequestBuilder instantiates a new BypassActivationLockRequestBuilder and sets the default values.
func NewBypassActivationLockRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BypassActivationLockRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBypassActivationLockRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation bypass activation lock
func (m *BypassActivationLockRequestBuilder) CreatePostRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration bypass activation lock
func (m *BypassActivationLockRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *BypassActivationLockRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post bypass activation lock
func (m *BypassActivationLockRequestBuilder) Post()(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler bypass activation lock
func (m *BypassActivationLockRequestBuilder) PostWithRequestConfigurationAndResponseHandler(requestConfiguration *BypassActivationLockRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
