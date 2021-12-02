package uploadclientcertificate

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// UploadClientCertificateRequestBuilder builds and executes requests for operations under \identity\apiConnectors\{identityApiConnector-id}\microsoft.graph.uploadClientCertificate
type UploadClientCertificateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UploadClientCertificateRequestBuilderPostOptions options for Post
type UploadClientCertificateRequestBuilderPostOptions struct {
    // 
    Body *UploadClientCertificateRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UploadClientCertificateResponse union type wrapper for classes identityApiConnector
type UploadClientCertificateResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type identityApiConnector
    identityApiConnector *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IdentityApiConnector;
}
// NewUploadClientCertificateResponse instantiates a new uploadClientCertificateResponse and sets the default values.
func NewUploadClientCertificateResponse()(*UploadClientCertificateResponse) {
    m := &UploadClientCertificateResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UploadClientCertificateResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIdentityApiConnector gets the identityApiConnector property value. Union type representation for type identityApiConnector
func (m *UploadClientCertificateResponse) GetIdentityApiConnector()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IdentityApiConnector) {
    if m == nil {
        return nil
    } else {
        return m.identityApiConnector
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UploadClientCertificateResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["identityApiConnector"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewIdentityApiConnector() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityApiConnector(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IdentityApiConnector))
        }
        return nil
    }
    return res
}
func (m *UploadClientCertificateResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UploadClientCertificateResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("identityApiConnector", m.GetIdentityApiConnector())
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
func (m *UploadClientCertificateResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIdentityApiConnector sets the identityApiConnector property value. Union type representation for type identityApiConnector
func (m *UploadClientCertificateResponse) SetIdentityApiConnector(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IdentityApiConnector)() {
    if m != nil {
        m.identityApiConnector = value
    }
}
// NewUploadClientCertificateRequestBuilderInternal instantiates a new UploadClientCertificateRequestBuilder and sets the default values.
func NewUploadClientCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UploadClientCertificateRequestBuilder) {
    m := &UploadClientCertificateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/apiConnectors/{identityApiConnector_id}/microsoft.graph.uploadClientCertificate";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUploadClientCertificateRequestBuilder instantiates a new UploadClientCertificateRequestBuilder and sets the default values.
func NewUploadClientCertificateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UploadClientCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUploadClientCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action uploadClientCertificate
func (m *UploadClientCertificateRequestBuilder) CreatePostRequestInformation(options *UploadClientCertificateRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action uploadClientCertificate
func (m *UploadClientCertificateRequestBuilder) Post(options *UploadClientCertificateRequestBuilderPostOptions)(*UploadClientCertificateResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUploadClientCertificateResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*UploadClientCertificateResponse), nil
}
