package registeredusers

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i31b110b87a150128d0be2fb38c3d1a41d8789a05c1bdb153ed1383f78798d28a "github.com/microsoftgraph/msgraph-sdk-go/devices/item/registeredusers/ref"
)

// RegisteredUsersRequestBuilder builds and executes requests for operations under \devices\{device-id}\registeredUsers
type RegisteredUsersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RegisteredUsersRequestBuilderGetOptions options for Get
type RegisteredUsersRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RegisteredUsersRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RegisteredUsersRequestBuilderGetQueryParameters collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
type RegisteredUsersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// NewRegisteredUsersRequestBuilderInternal instantiates a new RegisteredUsersRequestBuilder and sets the default values.
func NewRegisteredUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RegisteredUsersRequestBuilder) {
    m := &RegisteredUsersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/devices/{device_id}/registeredUsers{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRegisteredUsersRequestBuilder instantiates a new RegisteredUsersRequestBuilder and sets the default values.
func NewRegisteredUsersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RegisteredUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRegisteredUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
func (m *RegisteredUsersRequestBuilder) CreateGetRequestInformation(options *RegisteredUsersRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
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
// Get collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
func (m *RegisteredUsersRequestBuilder) Get(options *RegisteredUsersRequestBuilderGetOptions)(*RegisteredUsersResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRegisteredUsersResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*RegisteredUsersResponse), nil
}
func (m *RegisteredUsersRequestBuilder) Ref()(*i31b110b87a150128d0be2fb38c3d1a41d8789a05c1bdb153ed1383f78798d28a.RefRequestBuilder) {
    return i31b110b87a150128d0be2fb38c3d1a41d8789a05c1bdb153ed1383f78798d28a.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
