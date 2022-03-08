package updaterecordingstatus

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// UpdateRecordingStatusRequestBuilder provides operations to call the updateRecordingStatus method.
type UpdateRecordingStatusRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UpdateRecordingStatusRequestBuilderPostOptions options for Post
type UpdateRecordingStatusRequestBuilderPostOptions struct {
    // 
    Body UpdateRecordingStatusRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UpdateRecordingStatusResponse union type wrapper for classes updateRecordingStatusOperation
type UpdateRecordingStatusResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type updateRecordingStatusOperation
    updateRecordingStatusOperation i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.UpdateRecordingStatusOperationable;
}
// NewUpdateRecordingStatusResponse instantiates a new updateRecordingStatusResponse and sets the default values.
func NewUpdateRecordingStatusResponse()(*UpdateRecordingStatusResponse) {
    m := &UpdateRecordingStatusResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateUpdateRecordingStatusResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUpdateRecordingStatusResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateRecordingStatusResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateRecordingStatusResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["updateRecordingStatusOperation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateUpdateRecordingStatusOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateRecordingStatusOperation(val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.UpdateRecordingStatusOperationable))
        }
        return nil
    }
    return res
}
// GetUpdateRecordingStatusOperation gets the updateRecordingStatusOperation property value. Union type representation for type updateRecordingStatusOperation
func (m *UpdateRecordingStatusResponse) GetUpdateRecordingStatusOperation()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.UpdateRecordingStatusOperationable) {
    if m == nil {
        return nil
    } else {
        return m.updateRecordingStatusOperation
    }
}
func (m *UpdateRecordingStatusResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UpdateRecordingStatusResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("updateRecordingStatusOperation", m.GetUpdateRecordingStatusOperation())
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
func (m *UpdateRecordingStatusResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUpdateRecordingStatusOperation sets the updateRecordingStatusOperation property value. Union type representation for type updateRecordingStatusOperation
func (m *UpdateRecordingStatusResponse) SetUpdateRecordingStatusOperation(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.UpdateRecordingStatusOperationable)() {
    if m != nil {
        m.updateRecordingStatusOperation = value
    }
}
// NewUpdateRecordingStatusRequestBuilderInternal instantiates a new UpdateRecordingStatusRequestBuilder and sets the default values.
func NewUpdateRecordingStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdateRecordingStatusRequestBuilder) {
    m := &UpdateRecordingStatusRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/{call_id}/microsoft.graph.updateRecordingStatus";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUpdateRecordingStatusRequestBuilder instantiates a new UpdateRecordingStatusRequestBuilder and sets the default values.
func NewUpdateRecordingStatusRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdateRecordingStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdateRecordingStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action updateRecordingStatus
func (m *UpdateRecordingStatusRequestBuilder) CreatePostRequestInformation(options *UpdateRecordingStatusRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action updateRecordingStatus
func (m *UpdateRecordingStatusRequestBuilder) Post(options *UpdateRecordingStatusRequestBuilderPostOptions)(UpdateRecordingStatusResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateUpdateRecordingStatusResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(UpdateRecordingStatusResponseable), nil
}
