package setmobiledevicemanagementauthority

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// SetMobileDeviceManagementAuthorityRequestBuilder builds and executes requests for operations under \organization\{organization-id}\microsoft.graph.setMobileDeviceManagementAuthority
type SetMobileDeviceManagementAuthorityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SetMobileDeviceManagementAuthorityRequestBuilderPostOptions options for Post
type SetMobileDeviceManagementAuthorityRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSetMobileDeviceManagementAuthorityRequestBuilderInternal instantiates a new SetMobileDeviceManagementAuthorityRequestBuilder and sets the default values.
func NewSetMobileDeviceManagementAuthorityRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetMobileDeviceManagementAuthorityRequestBuilder) {
    m := &SetMobileDeviceManagementAuthorityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization_id}/microsoft.graph.setMobileDeviceManagementAuthority";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetMobileDeviceManagementAuthorityRequestBuilder instantiates a new SetMobileDeviceManagementAuthorityRequestBuilder and sets the default values.
func NewSetMobileDeviceManagementAuthorityRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetMobileDeviceManagementAuthorityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetMobileDeviceManagementAuthorityRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation set mobile device management authority
func (m *SetMobileDeviceManagementAuthorityRequestBuilder) CreatePostRequestInformation(options *SetMobileDeviceManagementAuthorityRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post set mobile device management authority
func (m *SetMobileDeviceManagementAuthorityRequestBuilder) Post(options *SetMobileDeviceManagementAuthorityRequestBuilderPostOptions)(*int32, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(*requestInfo, "int32", nil)
    if err != nil {
        return nil, err
    }
    return res.(*int32), nil
}
