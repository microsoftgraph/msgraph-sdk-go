package manageddeviceenrollmenttopfailureswithperiod

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder builds and executes requests for operations under \reports\microsoft.graph.managedDeviceEnrollmentTopFailures(period='{period}')
type ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderGetOptions options for Get
type ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedDeviceEnrollmentTopFailuresWithPeriodResponse union type wrapper for classes report
type ManagedDeviceEnrollmentTopFailuresWithPeriodResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type report
    report *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Report;
}
// NewManagedDeviceEnrollmentTopFailuresWithPeriodResponse instantiates a new managedDeviceEnrollmentTopFailuresWithPeriodResponse and sets the default values.
func NewManagedDeviceEnrollmentTopFailuresWithPeriodResponse()(*ManagedDeviceEnrollmentTopFailuresWithPeriodResponse) {
    m := &ManagedDeviceEnrollmentTopFailuresWithPeriodResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDeviceEnrollmentTopFailuresWithPeriodResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetReport gets the report property value. Union type representation for type report
func (m *ManagedDeviceEnrollmentTopFailuresWithPeriodResponse) GetReport()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Report) {
    if m == nil {
        return nil
    } else {
        return m.report
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceEnrollmentTopFailuresWithPeriodResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["report"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewReport() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReport(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Report))
        }
        return nil
    }
    return res
}
func (m *ManagedDeviceEnrollmentTopFailuresWithPeriodResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedDeviceEnrollmentTopFailuresWithPeriodResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("report", m.GetReport())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDeviceEnrollmentTopFailuresWithPeriodResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetReport sets the report property value. Union type representation for type report
func (m *ManagedDeviceEnrollmentTopFailuresWithPeriodResponse) SetReport(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Report)() {
    if m != nil {
        m.report = value
    }
}
// NewManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal instantiates a new ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder and sets the default values.
func NewManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, period *string)(*ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    m := &ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.managedDeviceEnrollmentTopFailures(period='{period}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if period != nil {
        urlTplParams["period"] = *period
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder instantiates a new ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder and sets the default values.
func NewManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function managedDeviceEnrollmentTopFailures
func (m *ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) CreateGetRequestInformation(options *ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function managedDeviceEnrollmentTopFailures
func (m *ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) Get(options *ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderGetOptions)(*ManagedDeviceEnrollmentTopFailuresWithPeriodResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceEnrollmentTopFailuresWithPeriodResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ManagedDeviceEnrollmentTopFailuresWithPeriodResponse), nil
}
