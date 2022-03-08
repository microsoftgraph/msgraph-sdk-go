package getnotebookfromweburl

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// GetNotebookFromWebUrlRequestBuilder provides operations to call the getNotebookFromWebUrl method.
type GetNotebookFromWebUrlRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetNotebookFromWebUrlRequestBuilderPostOptions options for Post
type GetNotebookFromWebUrlRequestBuilderPostOptions struct {
    // 
    Body GetNotebookFromWebUrlRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GetNotebookFromWebUrlResponse union type wrapper for classes CopyNotebookModel
type GetNotebookFromWebUrlResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type CopyNotebookModel
    copyNotebookModel i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CopyNotebookModelable;
}
// NewGetNotebookFromWebUrlResponse instantiates a new getNotebookFromWebUrlResponse and sets the default values.
func NewGetNotebookFromWebUrlResponse()(*GetNotebookFromWebUrlResponse) {
    m := &GetNotebookFromWebUrlResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateGetNotebookFromWebUrlResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewGetNotebookFromWebUrlResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetNotebookFromWebUrlResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCopyNotebookModel gets the copyNotebookModel property value. Union type representation for type CopyNotebookModel
func (m *GetNotebookFromWebUrlResponse) GetCopyNotebookModel()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CopyNotebookModelable) {
    if m == nil {
        return nil
    } else {
        return m.copyNotebookModel
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetNotebookFromWebUrlResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["copyNotebookModel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateCopyNotebookModelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopyNotebookModel(val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CopyNotebookModelable))
        }
        return nil
    }
    return res
}
func (m *GetNotebookFromWebUrlResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetNotebookFromWebUrlResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("copyNotebookModel", m.GetCopyNotebookModel())
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
func (m *GetNotebookFromWebUrlResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCopyNotebookModel sets the copyNotebookModel property value. Union type representation for type CopyNotebookModel
func (m *GetNotebookFromWebUrlResponse) SetCopyNotebookModel(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CopyNotebookModelable)() {
    if m != nil {
        m.copyNotebookModel = value
    }
}
// NewGetNotebookFromWebUrlRequestBuilderInternal instantiates a new GetNotebookFromWebUrlRequestBuilder and sets the default values.
func NewGetNotebookFromWebUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetNotebookFromWebUrlRequestBuilder) {
    m := &GetNotebookFromWebUrlRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onenote/notebooks/microsoft.graph.getNotebookFromWebUrl";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetNotebookFromWebUrlRequestBuilder instantiates a new GetNotebookFromWebUrlRequestBuilder and sets the default values.
func NewGetNotebookFromWebUrlRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetNotebookFromWebUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetNotebookFromWebUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getNotebookFromWebUrl
func (m *GetNotebookFromWebUrlRequestBuilder) CreatePostRequestInformation(options *GetNotebookFromWebUrlRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action getNotebookFromWebUrl
func (m *GetNotebookFromWebUrlRequestBuilder) Post(options *GetNotebookFromWebUrlRequestBuilderPostOptions)(GetNotebookFromWebUrlResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetNotebookFromWebUrlResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetNotebookFromWebUrlResponseable), nil
}
