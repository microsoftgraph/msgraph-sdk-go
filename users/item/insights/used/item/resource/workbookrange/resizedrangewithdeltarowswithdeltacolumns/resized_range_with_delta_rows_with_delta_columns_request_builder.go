package resizedrangewithdeltarowswithdeltacolumns

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.resizedRange(deltaRows={deltaRows},deltaColumns={deltaColumns})
type ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderGetOptions options for Get
type ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ResizedRangeWithDeltaRowsWithDeltaColumnsResponse union type wrapper for classes workbookRange
type ResizedRangeWithDeltaRowsWithDeltaColumnsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type workbookRange
    workbookRange *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookRange;
}
// NewResizedRangeWithDeltaRowsWithDeltaColumnsResponse instantiates a new resizedRangeWithDeltaRowsWithDeltaColumnsResponse and sets the default values.
func NewResizedRangeWithDeltaRowsWithDeltaColumnsResponse()(*ResizedRangeWithDeltaRowsWithDeltaColumnsResponse) {
    m := &ResizedRangeWithDeltaRowsWithDeltaColumnsResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResizedRangeWithDeltaRowsWithDeltaColumnsResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetWorkbookRange gets the workbookRange property value. Union type representation for type workbookRange
func (m *ResizedRangeWithDeltaRowsWithDeltaColumnsResponse) GetWorkbookRange()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookRange) {
    if m == nil {
        return nil
    } else {
        return m.workbookRange
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResizedRangeWithDeltaRowsWithDeltaColumnsResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["workbookRange"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookRange() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkbookRange(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookRange))
        }
        return nil
    }
    return res
}
func (m *ResizedRangeWithDeltaRowsWithDeltaColumnsResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ResizedRangeWithDeltaRowsWithDeltaColumnsResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("workbookRange", m.GetWorkbookRange())
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
func (m *ResizedRangeWithDeltaRowsWithDeltaColumnsResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetWorkbookRange sets the workbookRange property value. Union type representation for type workbookRange
func (m *ResizedRangeWithDeltaRowsWithDeltaColumnsResponse) SetWorkbookRange(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookRange)() {
    if m != nil {
        m.workbookRange = value
    }
}
// NewResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal instantiates a new ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder and sets the default values.
func NewResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, deltaRows *int32, deltaColumns *int32)(*ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) {
    m := &ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights/used/{usedInsight_id}/resource/microsoft.graph.workbookRange/microsoft.graph.resizedRange(deltaRows={deltaRows},deltaColumns={deltaColumns})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if deltaRows != nil {
        urlTplParams["deltaRows"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*deltaRows), 10)
    }
    if deltaColumns != nil {
        urlTplParams["deltaColumns"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*deltaColumns), 10)
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder instantiates a new ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder and sets the default values.
func NewResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// CreateGetRequestInformation invoke function resizedRange
func (m *ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) CreateGetRequestInformation(options *ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function resizedRange
func (m *ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) Get(options *ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderGetOptions)(*ResizedRangeWithDeltaRowsWithDeltaColumnsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResizedRangeWithDeltaRowsWithDeltaColumnsResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ResizedRangeWithDeltaRowsWithDeltaColumnsResponse), nil
}
