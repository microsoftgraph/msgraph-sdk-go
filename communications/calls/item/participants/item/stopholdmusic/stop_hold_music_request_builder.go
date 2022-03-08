package stopholdmusic

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// StopHoldMusicRequestBuilder provides operations to call the stopHoldMusic method.
type StopHoldMusicRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// StopHoldMusicRequestBuilderPostOptions options for Post
type StopHoldMusicRequestBuilderPostOptions struct {
    // 
    Body StopHoldMusicRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// StopHoldMusicResponse union type wrapper for classes stopHoldMusicOperation
type StopHoldMusicResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type stopHoldMusicOperation
    stopHoldMusicOperation i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.StopHoldMusicOperationable;
}
// NewStopHoldMusicResponse instantiates a new stopHoldMusicResponse and sets the default values.
func NewStopHoldMusicResponse()(*StopHoldMusicResponse) {
    m := &StopHoldMusicResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateStopHoldMusicResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewStopHoldMusicResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StopHoldMusicResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StopHoldMusicResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["stopHoldMusicOperation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateStopHoldMusicOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStopHoldMusicOperation(val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.StopHoldMusicOperationable))
        }
        return nil
    }
    return res
}
// GetStopHoldMusicOperation gets the stopHoldMusicOperation property value. Union type representation for type stopHoldMusicOperation
func (m *StopHoldMusicResponse) GetStopHoldMusicOperation()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.StopHoldMusicOperationable) {
    if m == nil {
        return nil
    } else {
        return m.stopHoldMusicOperation
    }
}
func (m *StopHoldMusicResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *StopHoldMusicResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("stopHoldMusicOperation", m.GetStopHoldMusicOperation())
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
func (m *StopHoldMusicResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetStopHoldMusicOperation sets the stopHoldMusicOperation property value. Union type representation for type stopHoldMusicOperation
func (m *StopHoldMusicResponse) SetStopHoldMusicOperation(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.StopHoldMusicOperationable)() {
    if m != nil {
        m.stopHoldMusicOperation = value
    }
}
// NewStopHoldMusicRequestBuilderInternal instantiates a new StopHoldMusicRequestBuilder and sets the default values.
func NewStopHoldMusicRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*StopHoldMusicRequestBuilder) {
    m := &StopHoldMusicRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/{call_id}/participants/{participant_id}/microsoft.graph.stopHoldMusic";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewStopHoldMusicRequestBuilder instantiates a new StopHoldMusicRequestBuilder and sets the default values.
func NewStopHoldMusicRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*StopHoldMusicRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStopHoldMusicRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action stopHoldMusic
func (m *StopHoldMusicRequestBuilder) CreatePostRequestInformation(options *StopHoldMusicRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action stopHoldMusic
func (m *StopHoldMusicRequestBuilder) Post(options *StopHoldMusicRequestBuilderPostOptions)(StopHoldMusicResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateStopHoldMusicResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(StopHoldMusicResponseable), nil
}
