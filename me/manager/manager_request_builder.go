package manager

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    if299d9ef56e2af8df97202321de61e2aee15552ebb21beaaf1dc86170d772b19 "github.com/microsoftgraph/msgraph-sdk-go/contracts/getbyids"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i394eba5aaab363168091f29758854ac179599524e2675d36e01fb901cec8d9d3 "github.com/microsoftgraph/msgraph-sdk-go/me/manager/ref"
)

// ManagerRequestBuilder builds and executes requests for operations under \me\manager
type ManagerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagerRequestBuilderGetOptions options for Get
type ManagerRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagerRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagerRequestBuilderGetQueryParameters the user or contact that is this user's manager. Read-only. (HTTP Methods: GET, PUT, DELETE.). Supports $expand.
type ManagerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewManagerRequestBuilderInternal instantiates a new ManagerRequestBuilder and sets the default values.
func NewManagerRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagerRequestBuilder) {
    m := &ManagerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/manager{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagerRequestBuilder instantiates a new ManagerRequestBuilder and sets the default values.
func NewManagerRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the user or contact that is this user's manager. Read-only. (HTTP Methods: GET, PUT, DELETE.). Supports $expand.
func (m *ManagerRequestBuilder) CreateGetRequestInformation(options *ManagerRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the user or contact that is this user's manager. Read-only. (HTTP Methods: GET, PUT, DELETE.). Supports $expand.
func (m *ManagerRequestBuilder) Get(options *ManagerRequestBuilderGetOptions)(*if299d9ef56e2af8df97202321de61e2aee15552ebb21beaaf1dc86170d772b19.DirectoryObject, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return if299d9ef56e2af8df97202321de61e2aee15552ebb21beaaf1dc86170d772b19.NewDirectoryObject() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*if299d9ef56e2af8df97202321de61e2aee15552ebb21beaaf1dc86170d772b19.DirectoryObject), nil
}
func (m *ManagerRequestBuilder) Ref()(*i394eba5aaab363168091f29758854ac179599524e2675d36e01fb901cec8d9d3.RefRequestBuilder) {
    return i394eba5aaab363168091f29758854ac179599524e2675d36e01fb901cec8d9d3.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
