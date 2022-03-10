package generatedownloaduri

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// GenerateDownloadUriRequestBuilder provides operations to call the generateDownloadUri method.
type GenerateDownloadUriRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GenerateDownloadUriRequestBuilderPostOptions options for Post
type GenerateDownloadUriRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GenerateDownloadUriResponse union type wrapper for classes accessReviewHistoryInstance
type GenerateDownloadUriResponse struct {
    // Union type representation for type accessReviewHistoryInstance
    accessReviewHistoryInstance i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AccessReviewHistoryInstanceable;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
}
// NewGenerateDownloadUriResponse instantiates a new generateDownloadUriResponse and sets the default values.
func NewGenerateDownloadUriResponse()(*GenerateDownloadUriResponse) {
    m := &GenerateDownloadUriResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateGenerateDownloadUriResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewGenerateDownloadUriResponse(), nil
}
// GetAccessReviewHistoryInstance gets the accessReviewHistoryInstance property value. Union type representation for type accessReviewHistoryInstance
func (m *GenerateDownloadUriResponse) GetAccessReviewHistoryInstance()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AccessReviewHistoryInstanceable) {
    if m == nil {
        return nil
    } else {
        return m.accessReviewHistoryInstance
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GenerateDownloadUriResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GenerateDownloadUriResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessReviewHistoryInstance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateAccessReviewHistoryInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessReviewHistoryInstance(val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AccessReviewHistoryInstanceable))
        }
        return nil
    }
    return res
}
func (m *GenerateDownloadUriResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GenerateDownloadUriResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accessReviewHistoryInstance", m.GetAccessReviewHistoryInstance())
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
// SetAccessReviewHistoryInstance sets the accessReviewHistoryInstance property value. Union type representation for type accessReviewHistoryInstance
func (m *GenerateDownloadUriResponse) SetAccessReviewHistoryInstance(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AccessReviewHistoryInstanceable)() {
    if m != nil {
        m.accessReviewHistoryInstance = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GenerateDownloadUriResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// GenerateDownloadUriResponseable 
type GenerateDownloadUriResponseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessReviewHistoryInstance()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AccessReviewHistoryInstanceable)
    SetAccessReviewHistoryInstance(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AccessReviewHistoryInstanceable)()
}
// NewGenerateDownloadUriRequestBuilderInternal instantiates a new GenerateDownloadUriRequestBuilder and sets the default values.
func NewGenerateDownloadUriRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GenerateDownloadUriRequestBuilder) {
    m := &GenerateDownloadUriRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/historyDefinitions/{accessReviewHistoryDefinition_id}/instances/{accessReviewHistoryInstance_id}/microsoft.graph.generateDownloadUri";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGenerateDownloadUriRequestBuilder instantiates a new GenerateDownloadUriRequestBuilder and sets the default values.
func NewGenerateDownloadUriRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GenerateDownloadUriRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGenerateDownloadUriRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action generateDownloadUri
func (m *GenerateDownloadUriRequestBuilder) CreatePostRequestInformation(options *GenerateDownloadUriRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action generateDownloadUri
func (m *GenerateDownloadUriRequestBuilder) Post(options *GenerateDownloadUriRequestBuilderPostOptions)(GenerateDownloadUriResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGenerateDownloadUriResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GenerateDownloadUriResponseable), nil
}
