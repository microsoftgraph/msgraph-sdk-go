package instantiate

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// InstantiateRequestBuilder provides operations to call the instantiate method.
type InstantiateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// InstantiateRequestBuilderPostOptions options for Post
type InstantiateRequestBuilderPostOptions struct {
    // 
    Body InstantiateRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// InstantiateResponse union type wrapper for classes applicationServicePrincipal
type InstantiateResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type applicationServicePrincipal
    applicationServicePrincipal i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ApplicationServicePrincipalable;
}
// NewInstantiateResponse instantiates a new instantiateResponse and sets the default values.
func NewInstantiateResponse()(*InstantiateResponse) {
    m := &InstantiateResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateInstantiateResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewInstantiateResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InstantiateResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplicationServicePrincipal gets the applicationServicePrincipal property value. Union type representation for type applicationServicePrincipal
func (m *InstantiateResponse) GetApplicationServicePrincipal()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ApplicationServicePrincipalable) {
    if m == nil {
        return nil
    } else {
        return m.applicationServicePrincipal
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InstantiateResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applicationServicePrincipal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateApplicationServicePrincipalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationServicePrincipal(val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ApplicationServicePrincipalable))
        }
        return nil
    }
    return res
}
func (m *InstantiateResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *InstantiateResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("applicationServicePrincipal", m.GetApplicationServicePrincipal())
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
func (m *InstantiateResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplicationServicePrincipal sets the applicationServicePrincipal property value. Union type representation for type applicationServicePrincipal
func (m *InstantiateResponse) SetApplicationServicePrincipal(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ApplicationServicePrincipalable)() {
    if m != nil {
        m.applicationServicePrincipal = value
    }
}
// InstantiateResponseable 
type InstantiateResponseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    InstantiateResponseable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplicationServicePrincipal()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ApplicationServicePrincipalable)
    SetApplicationServicePrincipal(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ApplicationServicePrincipalable)()
}
// NewInstantiateRequestBuilderInternal instantiates a new InstantiateRequestBuilder and sets the default values.
func NewInstantiateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InstantiateRequestBuilder) {
    m := &InstantiateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applicationTemplates/{applicationTemplate_id}/microsoft.graph.instantiate";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInstantiateRequestBuilder instantiates a new InstantiateRequestBuilder and sets the default values.
func NewInstantiateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InstantiateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInstantiateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action instantiate
func (m *InstantiateRequestBuilder) CreatePostRequestInformation(options *InstantiateRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action instantiate
func (m *InstantiateRequestBuilder) Post(options *InstantiateRequestBuilderPostOptions)(InstantiateResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateInstantiateResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(InstantiateResponseable), nil
}
