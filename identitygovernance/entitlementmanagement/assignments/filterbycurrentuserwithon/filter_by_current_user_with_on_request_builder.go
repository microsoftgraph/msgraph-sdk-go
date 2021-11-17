package filterbycurrentuserwithon

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Builds and executes requests for operations under \identityGovernance\entitlementManagement\assignments\microsoft.graph.filterByCurrentUser(on={on})
type FilterByCurrentUserWithOnRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type FilterByCurrentUserWithOnRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new FilterByCurrentUserWithOnRequestBuilder and sets the default values.
// Parameters:
//  - on : Usage: on={on}
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewFilterByCurrentUserWithOnRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, on *string)(*FilterByCurrentUserWithOnRequestBuilder) {
    m := &FilterByCurrentUserWithOnRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/assignments/microsoft.graph.filterByCurrentUser(on={on})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if on != nil {
        urlTplParams["on"] = *on
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new FilterByCurrentUserWithOnRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewFilterByCurrentUserWithOnRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FilterByCurrentUserWithOnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilterByCurrentUserWithOnRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Invoke function filterByCurrentUser
// Parameters:
//  - options : Options for the request
func (m *FilterByCurrentUserWithOnRequestBuilder) CreateGetRequestInformation(options *FilterByCurrentUserWithOnRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Invoke function filterByCurrentUser
// Parameters:
//  - options : Options for the request
func (m *FilterByCurrentUserWithOnRequestBuilder) Get(options *FilterByCurrentUserWithOnRequestBuilderGetOptions)([]FilterByCurrentUserWithOn, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFilterByCurrentUserWithOn() }, nil)
    if err != nil {
        return nil, err
    }
    val := make([]FilterByCurrentUserWithOn, len(res))
    for i, v := range res {
        val[i] = *(v.(*FilterByCurrentUserWithOn))
    }
    return val, nil
}
