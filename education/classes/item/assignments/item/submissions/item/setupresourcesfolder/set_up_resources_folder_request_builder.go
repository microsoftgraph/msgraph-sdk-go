package setupresourcesfolder

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// Builds and executes requests for operations under \education\classes\{educationClass-id}\assignments\{educationAssignment-id}\submissions\{educationSubmission-id}\microsoft.graph.setUpResourcesFolder
type SetUpResourcesFolderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Union type wrapper for classes educationSubmission
type SetUpResourcesFolderResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type educationSubmission
    educationSubmission *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationSubmission;
}
// Instantiates a new setUpResourcesFolderResponse and sets the default values.
func NewSetUpResourcesFolderResponse()(*SetUpResourcesFolderResponse) {
    m := &SetUpResourcesFolderResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetUpResourcesFolderResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the educationSubmission property value. Union type representation for type educationSubmission
func (m *SetUpResourcesFolderResponse) GetEducationSubmission()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationSubmission) {
    if m == nil {
        return nil
    } else {
        return m.educationSubmission
    }
}
// The deserialization information for the current model
func (m *SetUpResourcesFolderResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["educationSubmission"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEducationSubmission() })
        if err != nil {
            return err
        }
        m.SetEducationSubmission(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationSubmission))
        return nil
    }
    return res
}
func (m *SetUpResourcesFolderResponse) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SetUpResourcesFolderResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("educationSubmission", m.GetEducationSubmission())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SetUpResourcesFolderResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the educationSubmission property value. Union type representation for type educationSubmission
// Parameters:
//  - value : Value to set for the educationSubmission property.
func (m *SetUpResourcesFolderResponse) SetEducationSubmission(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationSubmission)() {
    m.educationSubmission = value
}
// Instantiates a new SetUpResourcesFolderRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSetUpResourcesFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetUpResourcesFolderRequestBuilder) {
    m := &SetUpResourcesFolderRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/education/classes/{educationClass_id}/assignments/{educationAssignment_id}/submissions/{educationSubmission_id}/microsoft.graph.setUpResourcesFolder";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new SetUpResourcesFolderRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSetUpResourcesFolderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetUpResourcesFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetUpResourcesFolderRequestBuilderInternal(urlParams, requestAdapter)
}
// Invoke action setUpResourcesFolder
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *SetUpResourcesFolderRequestBuilder) CreatePostRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Invoke action setUpResourcesFolder
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *SetUpResourcesFolderRequestBuilder) Post(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*SetUpResourcesFolderResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSetUpResourcesFolderResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*SetUpResourcesFolderResponse), nil
}
