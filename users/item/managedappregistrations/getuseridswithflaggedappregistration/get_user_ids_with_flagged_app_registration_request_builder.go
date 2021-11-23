package getuseridswithflaggedappregistration

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// getUserIdsWithFlaggedAppRegistrationRequestBuilder builds and executes requests for operations under \users\{user-id}\managedAppRegistrations\microsoft.graph.getUserIdsWithFlaggedAppRegistration()
type GetUserIdsWithFlaggedAppRegistrationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetUserIdsWithFlaggedAppRegistrationRequestBuilderGetOptions options for Get
type GetUserIdsWithFlaggedAppRegistrationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetUserIdsWithFlaggedAppRegistrationRequestBuilderInternal instantiates a new GetUserIdsWithFlaggedAppRegistrationRequestBuilder and sets the default values.
func NewGetUserIdsWithFlaggedAppRegistrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetUserIdsWithFlaggedAppRegistrationRequestBuilder) {
    m := &GetUserIdsWithFlaggedAppRegistrationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/managedAppRegistrations/microsoft.graph.getUserIdsWithFlaggedAppRegistration()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetUserIdsWithFlaggedAppRegistrationRequestBuilder instantiates a new GetUserIdsWithFlaggedAppRegistrationRequestBuilder and sets the default values.
func NewGetUserIdsWithFlaggedAppRegistrationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetUserIdsWithFlaggedAppRegistrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetUserIdsWithFlaggedAppRegistrationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getUserIdsWithFlaggedAppRegistration
func (m *GetUserIdsWithFlaggedAppRegistrationRequestBuilder) CreateGetRequestInformation(options *GetUserIdsWithFlaggedAppRegistrationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Get invoke function getUserIdsWithFlaggedAppRegistration
func (m *GetUserIdsWithFlaggedAppRegistrationRequestBuilder) Get(options *GetUserIdsWithFlaggedAppRegistrationRequestBuilderGetOptions)([]string, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveCollectionAsync(*requestInfo, "string", nil)
    if err != nil {
        return nil, err
    }
    val := make([]string, len(res))
    for i, v := range res {
        val[i] = *(v.(*string))
    }
    return val, nil
}
