package wipemanagedappregistrationsbydevicetag

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// WipeManagedAppRegistrationsByDeviceTagRequestBuilder provides operations to call the wipeManagedAppRegistrationsByDeviceTag method.
type WipeManagedAppRegistrationsByDeviceTagRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WipeManagedAppRegistrationsByDeviceTagRequestBuilderPostOptions options for Post
type WipeManagedAppRegistrationsByDeviceTagRequestBuilderPostOptions struct {
    // 
    Body WipeManagedAppRegistrationsByDeviceTagRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal instantiates a new WipeManagedAppRegistrationsByDeviceTagRequestBuilder and sets the default values.
func NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    m := &WipeManagedAppRegistrationsByDeviceTagRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/microsoft.graph.wipeManagedAppRegistrationsByDeviceTag";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWipeManagedAppRegistrationsByDeviceTagRequestBuilder instantiates a new WipeManagedAppRegistrationsByDeviceTagRequestBuilder and sets the default values.
func NewWipeManagedAppRegistrationsByDeviceTagRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation issues a wipe operation on an app registration with specified device tag.
func (m *WipeManagedAppRegistrationsByDeviceTagRequestBuilder) CreatePostRequestInformation(options *WipeManagedAppRegistrationsByDeviceTagRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
