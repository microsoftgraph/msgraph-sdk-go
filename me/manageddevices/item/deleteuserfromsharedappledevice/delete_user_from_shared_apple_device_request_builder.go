package deleteuserfromsharedappledevice

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// deleteUserFromSharedAppleDeviceRequestBuilder builds and executes requests for operations under \me\managedDevices\{managedDevice-id}\microsoft.graph.deleteUserFromSharedAppleDevice
type DeleteUserFromSharedAppleDeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeleteUserFromSharedAppleDeviceRequestBuilderPostOptions options for Post
type DeleteUserFromSharedAppleDeviceRequestBuilderPostOptions struct {
    // 
    Body *DeleteUserFromSharedAppleDeviceRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal instantiates a new DeleteUserFromSharedAppleDeviceRequestBuilder and sets the default values.
func NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeleteUserFromSharedAppleDeviceRequestBuilder) {
    m := &DeleteUserFromSharedAppleDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices/{managedDevice_id}/microsoft.graph.deleteUserFromSharedAppleDevice";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeleteUserFromSharedAppleDeviceRequestBuilder instantiates a new DeleteUserFromSharedAppleDeviceRequestBuilder and sets the default values.
func NewDeleteUserFromSharedAppleDeviceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeleteUserFromSharedAppleDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation delete user from shared Apple device
func (m *DeleteUserFromSharedAppleDeviceRequestBuilder) CreatePostRequestInformation(options *DeleteUserFromSharedAppleDeviceRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
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
// Post delete user from shared Apple device
func (m *DeleteUserFromSharedAppleDeviceRequestBuilder) Post(options *DeleteUserFromSharedAppleDeviceRequestBuilderPostOptions)(error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
