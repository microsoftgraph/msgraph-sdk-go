package provisionemail

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// ProvisionEmailRequestBuilder provides operations to call the provisionEmail method.
type ProvisionEmailRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ProvisionEmailRequestBuilderPostOptions options for Post
type ProvisionEmailRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ProvisionEmailResponse union type wrapper for classes provisionChannelEmailResult
type ProvisionEmailResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type provisionChannelEmailResult
    provisionChannelEmailResult i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ProvisionChannelEmailResultable;
}
// NewProvisionEmailResponse instantiates a new provisionEmailResponse and sets the default values.
func NewProvisionEmailResponse()(*ProvisionEmailResponse) {
    m := &ProvisionEmailResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateProvisionEmailResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewProvisionEmailResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisionEmailResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisionEmailResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["provisionChannelEmailResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateProvisionChannelEmailResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisionChannelEmailResult(val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ProvisionChannelEmailResultable))
        }
        return nil
    }
    return res
}
// GetProvisionChannelEmailResult gets the provisionChannelEmailResult property value. Union type representation for type provisionChannelEmailResult
func (m *ProvisionEmailResponse) GetProvisionChannelEmailResult()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ProvisionChannelEmailResultable) {
    if m == nil {
        return nil
    } else {
        return m.provisionChannelEmailResult
    }
}
func (m *ProvisionEmailResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ProvisionEmailResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("provisionChannelEmailResult", m.GetProvisionChannelEmailResult())
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
func (m *ProvisionEmailResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetProvisionChannelEmailResult sets the provisionChannelEmailResult property value. Union type representation for type provisionChannelEmailResult
func (m *ProvisionEmailResponse) SetProvisionChannelEmailResult(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ProvisionChannelEmailResultable)() {
    if m != nil {
        m.provisionChannelEmailResult = value
    }
}
// ProvisionEmailResponseable 
type ProvisionEmailResponseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    ProvisionEmailResponseable
    GetProvisionChannelEmailResult()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ProvisionChannelEmailResultable)
    SetProvisionChannelEmailResult(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ProvisionChannelEmailResultable)()
}
// NewProvisionEmailRequestBuilderInternal instantiates a new ProvisionEmailRequestBuilder and sets the default values.
func NewProvisionEmailRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ProvisionEmailRequestBuilder) {
    m := &ProvisionEmailRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/joinedTeams/{team_id}/primaryChannel/microsoft.graph.provisionEmail";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewProvisionEmailRequestBuilder instantiates a new ProvisionEmailRequestBuilder and sets the default values.
func NewProvisionEmailRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ProvisionEmailRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProvisionEmailRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action provisionEmail
func (m *ProvisionEmailRequestBuilder) CreatePostRequestInformation(options *ProvisionEmailRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Post invoke action provisionEmail
func (m *ProvisionEmailRequestBuilder) Post(options *ProvisionEmailRequestBuilderPostOptions)(ProvisionEmailResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateProvisionEmailResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(ProvisionEmailResponseable), nil
}
