package cancelmediaprocessing

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// CancelMediaProcessingRequestBuilder provides operations to call the cancelMediaProcessing method.
type CancelMediaProcessingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CancelMediaProcessingRequestBuilderPostOptions options for Post
type CancelMediaProcessingRequestBuilderPostOptions struct {
    // 
    Body CancelMediaProcessingRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CancelMediaProcessingResponse union type wrapper for classes cancelMediaProcessingOperation
type CancelMediaProcessingResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type cancelMediaProcessingOperation
    cancelMediaProcessingOperation i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CancelMediaProcessingOperationable;
}
// NewCancelMediaProcessingResponse instantiates a new cancelMediaProcessingResponse and sets the default values.
func NewCancelMediaProcessingResponse()(*CancelMediaProcessingResponse) {
    m := &CancelMediaProcessingResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateCancelMediaProcessingResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCancelMediaProcessingResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CancelMediaProcessingResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCancelMediaProcessingOperation gets the cancelMediaProcessingOperation property value. Union type representation for type cancelMediaProcessingOperation
func (m *CancelMediaProcessingResponse) GetCancelMediaProcessingOperation()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CancelMediaProcessingOperationable) {
    if m == nil {
        return nil
    } else {
        return m.cancelMediaProcessingOperation
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CancelMediaProcessingResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["cancelMediaProcessingOperation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateCancelMediaProcessingOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCancelMediaProcessingOperation(val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CancelMediaProcessingOperationable))
        }
        return nil
    }
    return res
}
func (m *CancelMediaProcessingResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CancelMediaProcessingResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cancelMediaProcessingOperation", m.GetCancelMediaProcessingOperation())
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
func (m *CancelMediaProcessingResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCancelMediaProcessingOperation sets the cancelMediaProcessingOperation property value. Union type representation for type cancelMediaProcessingOperation
func (m *CancelMediaProcessingResponse) SetCancelMediaProcessingOperation(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CancelMediaProcessingOperationable)() {
    if m != nil {
        m.cancelMediaProcessingOperation = value
    }
}
// CancelMediaProcessingResponseable 
type CancelMediaProcessingResponseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    CancelMediaProcessingResponseable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCancelMediaProcessingOperation()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CancelMediaProcessingOperationable)
    SetCancelMediaProcessingOperation(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CancelMediaProcessingOperationable)()
}
// NewCancelMediaProcessingRequestBuilderInternal instantiates a new CancelMediaProcessingRequestBuilder and sets the default values.
func NewCancelMediaProcessingRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CancelMediaProcessingRequestBuilder) {
    m := &CancelMediaProcessingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/{call_id}/microsoft.graph.cancelMediaProcessing";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCancelMediaProcessingRequestBuilder instantiates a new CancelMediaProcessingRequestBuilder and sets the default values.
func NewCancelMediaProcessingRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CancelMediaProcessingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCancelMediaProcessingRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action cancelMediaProcessing
func (m *CancelMediaProcessingRequestBuilder) CreatePostRequestInformation(options *CancelMediaProcessingRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
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
// Post invoke action cancelMediaProcessing
func (m *CancelMediaProcessingRequestBuilder) Post(options *CancelMediaProcessingRequestBuilderPostOptions)(CancelMediaProcessingResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateCancelMediaProcessingResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(CancelMediaProcessingResponseable), nil
}
