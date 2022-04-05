package wipemanagedappregistrationsbydevicetag

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WipeManagedAppRegistrationsByDeviceTagRequestBuilder provides operations to call the wipeManagedAppRegistrationsByDeviceTag method.
type WipeManagedAppRegistrationsByDeviceTagRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WipeManagedAppRegistrationsByDeviceTagRequestBuilderPostOptions options for Post
type WipeManagedAppRegistrationsByDeviceTagRequestBuilderPostOptions struct {
    // 
    Body WipeManagedAppRegistrationsByDeviceTagRequestBodyable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal instantiates a new WipeManagedAppRegistrationsByDeviceTagRequestBuilder and sets the default values.
func NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    m := &WipeManagedAppRegistrationsByDeviceTagRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/microsoft.graph.wipeManagedAppRegistrationsByDeviceTag";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWipeManagedAppRegistrationsByDeviceTagRequestBuilder instantiates a new WipeManagedAppRegistrationsByDeviceTagRequestBuilder and sets the default values.
func NewWipeManagedAppRegistrationsByDeviceTagRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation issues a wipe operation on an app registration with specified device tag.
func (m *WipeManagedAppRegistrationsByDeviceTagRequestBuilder) CreatePostRequestInformation(options *WipeManagedAppRegistrationsByDeviceTagRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
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
// Post issues a wipe operation on an app registration with specified device tag.
func (m *WipeManagedAppRegistrationsByDeviceTagRequestBuilder) Post(options *WipeManagedAppRegistrationsByDeviceTagRequestBuilderPostOptions)(error) {
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
