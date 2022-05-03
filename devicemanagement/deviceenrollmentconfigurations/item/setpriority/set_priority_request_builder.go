package setpriority

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SetPriorityRequestBuilder provides operations to call the setPriority method.
type SetPriorityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SetPriorityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SetPriorityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSetPriorityRequestBuilderInternal instantiates a new SetPriorityRequestBuilder and sets the default values.
func NewSetPriorityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetPriorityRequestBuilder) {
    m := &SetPriorityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration%2Did}/microsoft.graph.setPriority";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetPriorityRequestBuilder instantiates a new SetPriorityRequestBuilder and sets the default values.
func NewSetPriorityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetPriorityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetPriorityRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action setPriority
func (m *SetPriorityRequestBuilder) CreatePostRequestInformation(body SetPriorityRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action setPriority
func (m *SetPriorityRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body SetPriorityRequestBodyable, requestConfiguration *SetPriorityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action setPriority
func (m *SetPriorityRequestBuilder) Post(body SetPriorityRequestBodyable)(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action setPriority
func (m *SetPriorityRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body SetPriorityRequestBodyable, requestConfiguration *SetPriorityRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
