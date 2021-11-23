package administrativeunit

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i44ee87fbaba09f09109fbc423f1bfe93712a1a875a4af8f8a1f7f62d2db4c490 "github.com/microsoftgraph/msgraph-sdk-go/education/schools/item/administrativeunit/ref"
)

// AdministrativeUnitRequestBuilder builds and executes requests for operations under \education\schools\{educationSchool-id}\administrativeUnit
type AdministrativeUnitRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AdministrativeUnitRequestBuilderGetOptions options for Get
type AdministrativeUnitRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AdministrativeUnitRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AdministrativeUnitRequestBuilderGetQueryParameters the underlying administrativeUnit for this school.
type AdministrativeUnitRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// NewAdministrativeUnitRequestBuilderInternal instantiates a new AdministrativeUnitRequestBuilder and sets the default values.
func NewAdministrativeUnitRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdministrativeUnitRequestBuilder) {
    m := &AdministrativeUnitRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/schools/{educationSchool_id}/administrativeUnit{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdministrativeUnitRequestBuilder instantiates a new AdministrativeUnitRequestBuilder and sets the default values.
func NewAdministrativeUnitRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdministrativeUnitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdministrativeUnitRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the underlying administrativeUnit for this school.
func (m *AdministrativeUnitRequestBuilder) CreateGetRequestInformation(options *AdministrativeUnitRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the underlying administrativeUnit for this school.
func (m *AdministrativeUnitRequestBuilder) Get(options *AdministrativeUnitRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AdministrativeUnit, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewAdministrativeUnit() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AdministrativeUnit), nil
}
func (m *AdministrativeUnitRequestBuilder) Ref()(*i44ee87fbaba09f09109fbc423f1bfe93712a1a875a4af8f8a1f7f62d2db4c490.RefRequestBuilder) {
    return i44ee87fbaba09f09109fbc423f1bfe93712a1a875a4af8f8a1f7f62d2db4c490.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
