package deleteuserfromsharedappledevice

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeleteUserFromSharedAppleDeviceRequestBuilder provides operations to call the deleteUserFromSharedAppleDevice method.
type DeleteUserFromSharedAppleDeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeleteUserFromSharedAppleDeviceRequestBuilderPostOptions options for Post
type DeleteUserFromSharedAppleDeviceRequestBuilderPostOptions struct {
    // 
    Body DeleteUserFromSharedAppleDeviceRequestBodyable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal instantiates a new DeleteUserFromSharedAppleDeviceRequestBuilder and sets the default values.
func NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeleteUserFromSharedAppleDeviceRequestBuilder) {
    m := &DeleteUserFromSharedAppleDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managedDevices/{managedDevice_id}/microsoft.graph.deleteUserFromSharedAppleDevice";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeleteUserFromSharedAppleDeviceRequestBuilder instantiates a new DeleteUserFromSharedAppleDeviceRequestBuilder and sets the default values.
func NewDeleteUserFromSharedAppleDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeleteUserFromSharedAppleDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation delete user from shared Apple device
func (m *DeleteUserFromSharedAppleDeviceRequestBuilder) CreatePostRequestInformation(options *DeleteUserFromSharedAppleDeviceRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post delete user from shared Apple device
func (m *DeleteUserFromSharedAppleDeviceRequestBuilder) Post(options *DeleteUserFromSharedAppleDeviceRequestBuilderPostOptions)(error) {
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
