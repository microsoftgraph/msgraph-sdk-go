package startholdmusic

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// StartHoldMusicRequestBuilder provides operations to call the startHoldMusic method.
type StartHoldMusicRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// StartHoldMusicRequestBuilderPostOptions options for Post
type StartHoldMusicRequestBuilderPostOptions struct {
    // 
    Body StartHoldMusicRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// StartHoldMusicResponse union type wrapper for classes startHoldMusicOperation
type StartHoldMusicResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type startHoldMusicOperation
    startHoldMusicOperation i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.StartHoldMusicOperationable;
}
// NewStartHoldMusicResponse instantiates a new startHoldMusicResponse and sets the default values.
func NewStartHoldMusicResponse()(*StartHoldMusicResponse) {
    m := &StartHoldMusicResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateStartHoldMusicResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewStartHoldMusicResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StartHoldMusicResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StartHoldMusicResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["startHoldMusicOperation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateStartHoldMusicOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartHoldMusicOperation(val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.StartHoldMusicOperationable))
        }
        return nil
    }
    return res
}
// GetStartHoldMusicOperation gets the startHoldMusicOperation property value. Union type representation for type startHoldMusicOperation
func (m *StartHoldMusicResponse) GetStartHoldMusicOperation()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.StartHoldMusicOperationable) {
    if m == nil {
        return nil
    } else {
        return m.startHoldMusicOperation
    }
}
func (m *StartHoldMusicResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *StartHoldMusicResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("startHoldMusicOperation", m.GetStartHoldMusicOperation())
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
func (m *StartHoldMusicResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetStartHoldMusicOperation sets the startHoldMusicOperation property value. Union type representation for type startHoldMusicOperation
func (m *StartHoldMusicResponse) SetStartHoldMusicOperation(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.StartHoldMusicOperationable)() {
    if m != nil {
        m.startHoldMusicOperation = value
    }
}
// StartHoldMusicResponseable 
type StartHoldMusicResponseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetStartHoldMusicOperation()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.StartHoldMusicOperationable)
    SetStartHoldMusicOperation(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.StartHoldMusicOperationable)()
}
// NewStartHoldMusicRequestBuilderInternal instantiates a new StartHoldMusicRequestBuilder and sets the default values.
func NewStartHoldMusicRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*StartHoldMusicRequestBuilder) {
    m := &StartHoldMusicRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/{call_id}/participants/{participant_id}/microsoft.graph.startHoldMusic";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewStartHoldMusicRequestBuilder instantiates a new StartHoldMusicRequestBuilder and sets the default values.
func NewStartHoldMusicRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*StartHoldMusicRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStartHoldMusicRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action startHoldMusic
func (m *StartHoldMusicRequestBuilder) CreatePostRequestInformation(options *StartHoldMusicRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action startHoldMusic
func (m *StartHoldMusicRequestBuilder) Post(options *StartHoldMusicRequestBuilderPostOptions)(StartHoldMusicResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateStartHoldMusicResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(StartHoldMusicResponseable), nil
}
