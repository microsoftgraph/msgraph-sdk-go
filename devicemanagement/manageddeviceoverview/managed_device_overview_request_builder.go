package manageddeviceoverview

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i8a112f9bef0ef189294c1f3ecf8b6e6f3da80cb9207f29240296ac2bc4151c97 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddeviceoverview/ref"
)

// ManagedDeviceOverviewRequestBuilder builds and executes requests for operations under \deviceManagement\managedDeviceOverview
type ManagedDeviceOverviewRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedDeviceOverviewRequestBuilderGetOptions options for Get
type ManagedDeviceOverviewRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedDeviceOverviewRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedDeviceOverviewRequestBuilderGetQueryParameters device overview
type ManagedDeviceOverviewRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewManagedDeviceOverviewRequestBuilderInternal instantiates a new ManagedDeviceOverviewRequestBuilder and sets the default values.
func NewManagedDeviceOverviewRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceOverviewRequestBuilder) {
    m := &ManagedDeviceOverviewRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managedDeviceOverview{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDeviceOverviewRequestBuilder instantiates a new ManagedDeviceOverviewRequestBuilder and sets the default values.
func NewManagedDeviceOverviewRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceOverviewRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceOverviewRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation device overview
func (m *ManagedDeviceOverviewRequestBuilder) CreateGetRequestInformation(options *ManagedDeviceOverviewRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get device overview
func (m *ManagedDeviceOverviewRequestBuilder) Get(options *ManagedDeviceOverviewRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDeviceOverview, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewManagedDeviceOverview() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedDeviceOverview), nil
}
func (m *ManagedDeviceOverviewRequestBuilder) Ref()(*i8a112f9bef0ef189294c1f3ecf8b6e6f3da80cb9207f29240296ac2bc4151c97.RefRequestBuilder) {
    return i8a112f9bef0ef189294c1f3ecf8b6e6f3da80cb9207f29240296ac2bc4151c97.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
