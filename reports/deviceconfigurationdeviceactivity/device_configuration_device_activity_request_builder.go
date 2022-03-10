package deviceconfigurationdeviceactivity

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// DeviceConfigurationDeviceActivityRequestBuilder provides operations to call the deviceConfigurationDeviceActivity method.
type DeviceConfigurationDeviceActivityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceConfigurationDeviceActivityRequestBuilderGetOptions options for Get
type DeviceConfigurationDeviceActivityRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceConfigurationDeviceActivityResponse union type wrapper for classes report
type DeviceConfigurationDeviceActivityResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type report
    report i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Reportable;
}
// NewDeviceConfigurationDeviceActivityResponse instantiates a new deviceConfigurationDeviceActivityResponse and sets the default values.
func NewDeviceConfigurationDeviceActivityResponse()(*DeviceConfigurationDeviceActivityResponse) {
    m := &DeviceConfigurationDeviceActivityResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateDeviceConfigurationDeviceActivityResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceConfigurationDeviceActivityResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceConfigurationDeviceActivityResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationDeviceActivityResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["report"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateReportFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReport(val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Reportable))
        }
        return nil
    }
    return res
}
// GetReport gets the report property value. Union type representation for type report
func (m *DeviceConfigurationDeviceActivityResponse) GetReport()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Reportable) {
    if m == nil {
        return nil
    } else {
        return m.report
    }
}
func (m *DeviceConfigurationDeviceActivityResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceConfigurationDeviceActivityResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceConfigurationDeviceActivityResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetReport sets the report property value. Union type representation for type report
func (m *DeviceConfigurationDeviceActivityResponse) SetReport(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Reportable)() {
    if m != nil {
        m.report = value
    }
}
// DeviceConfigurationDeviceActivityResponseable 
type DeviceConfigurationDeviceActivityResponseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    DeviceConfigurationDeviceActivityResponseable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetReport()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Reportable)
    SetReport(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Reportable)()
}
// NewDeviceConfigurationDeviceActivityRequestBuilderInternal instantiates a new DeviceConfigurationDeviceActivityRequestBuilder and sets the default values.
func NewDeviceConfigurationDeviceActivityRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceConfigurationDeviceActivityRequestBuilder) {
    m := &DeviceConfigurationDeviceActivityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.deviceConfigurationDeviceActivity()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceConfigurationDeviceActivityRequestBuilder instantiates a new DeviceConfigurationDeviceActivityRequestBuilder and sets the default values.
func NewDeviceConfigurationDeviceActivityRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceConfigurationDeviceActivityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationDeviceActivityRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation metadata for the device configuration device activity report
func (m *DeviceConfigurationDeviceActivityRequestBuilder) CreateGetRequestInformation(options *DeviceConfigurationDeviceActivityRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get metadata for the device configuration device activity report
func (m *DeviceConfigurationDeviceActivityRequestBuilder) Get(options *DeviceConfigurationDeviceActivityRequestBuilderGetOptions)(DeviceConfigurationDeviceActivityResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateDeviceConfigurationDeviceActivityResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(DeviceConfigurationDeviceActivityResponseable), nil
}
