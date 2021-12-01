package addformulalocal

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// AddFormulaLocalRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\names\microsoft.graph.addFormulaLocal
type AddFormulaLocalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AddFormulaLocalRequestBuilderPostOptions options for Post
type AddFormulaLocalRequestBuilderPostOptions struct {
    // 
    Body *AddFormulaLocalRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AddFormulaLocalResponse union type wrapper for classes workbookNamedItem
type AddFormulaLocalResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type workbookNamedItem
    workbookNamedItem *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookNamedItem;
}
// NewAddFormulaLocalResponse instantiates a new addFormulaLocalResponse and sets the default values.
func NewAddFormulaLocalResponse()(*AddFormulaLocalResponse) {
    m := &AddFormulaLocalResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AddFormulaLocalResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetWorkbookNamedItem gets the workbookNamedItem property value. Union type representation for type workbookNamedItem
func (m *AddFormulaLocalResponse) GetWorkbookNamedItem()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookNamedItem) {
    if m == nil {
        return nil
    } else {
        return m.workbookNamedItem
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddFormulaLocalResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["workbookNamedItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookNamedItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkbookNamedItem(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookNamedItem))
        }
        return nil
    }
    return res
}
func (m *AddFormulaLocalResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AddFormulaLocalResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("workbookNamedItem", m.GetWorkbookNamedItem())
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
func (m *AddFormulaLocalResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetWorkbookNamedItem sets the workbookNamedItem property value. Union type representation for type workbookNamedItem
func (m *AddFormulaLocalResponse) SetWorkbookNamedItem(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookNamedItem)() {
    if m != nil {
        m.workbookNamedItem = value
    }
}
// NewAddFormulaLocalRequestBuilderInternal instantiates a new AddFormulaLocalRequestBuilder and sets the default values.
func NewAddFormulaLocalRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AddFormulaLocalRequestBuilder) {
    m := &AddFormulaLocalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}/names/microsoft.graph.addFormulaLocal";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAddFormulaLocalRequestBuilder instantiates a new AddFormulaLocalRequestBuilder and sets the default values.
func NewAddFormulaLocalRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AddFormulaLocalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAddFormulaLocalRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action addFormulaLocal
func (m *AddFormulaLocalRequestBuilder) CreatePostRequestInformation(options *AddFormulaLocalRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action addFormulaLocal
func (m *AddFormulaLocalRequestBuilder) Post(options *AddFormulaLocalRequestBuilderPostOptions)(*AddFormulaLocalResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAddFormulaLocalResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*AddFormulaLocalResponse), nil
}
