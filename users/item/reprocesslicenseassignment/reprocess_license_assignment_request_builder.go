package reprocesslicenseassignment

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// ReprocessLicenseAssignmentRequestBuilder builds and executes requests for operations under \users\{user-id}\microsoft.graph.reprocessLicenseAssignment
type ReprocessLicenseAssignmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ReprocessLicenseAssignmentRequestBuilderPostOptions options for Post
type ReprocessLicenseAssignmentRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ReprocessLicenseAssignmentResponse union type wrapper for classes user
type ReprocessLicenseAssignmentResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type user
    user *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.User;
}
// NewReprocessLicenseAssignmentResponse instantiates a new reprocessLicenseAssignmentResponse and sets the default values.
func NewReprocessLicenseAssignmentResponse()(*ReprocessLicenseAssignmentResponse) {
    m := &ReprocessLicenseAssignmentResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReprocessLicenseAssignmentResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetUser gets the user property value. Union type representation for type user
func (m *ReprocessLicenseAssignmentResponse) GetUser()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.User) {
    if m == nil {
        return nil
    } else {
        return m.user
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReprocessLicenseAssignmentResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["user"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewUser() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.User))
        }
        return nil
    }
    return res
}
func (m *ReprocessLicenseAssignmentResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ReprocessLicenseAssignmentResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *ReprocessLicenseAssignmentResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUser sets the user property value. Union type representation for type user
func (m *ReprocessLicenseAssignmentResponse) SetUser(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.User)() {
    if m != nil {
        m.user = value
    }
}
// NewReprocessLicenseAssignmentRequestBuilderInternal instantiates a new ReprocessLicenseAssignmentRequestBuilder and sets the default values.
func NewReprocessLicenseAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ReprocessLicenseAssignmentRequestBuilder) {
    m := &ReprocessLicenseAssignmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/microsoft.graph.reprocessLicenseAssignment";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewReprocessLicenseAssignmentRequestBuilder instantiates a new ReprocessLicenseAssignmentRequestBuilder and sets the default values.
func NewReprocessLicenseAssignmentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ReprocessLicenseAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReprocessLicenseAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action reprocessLicenseAssignment
func (m *ReprocessLicenseAssignmentRequestBuilder) CreatePostRequestInformation(options *ReprocessLicenseAssignmentRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action reprocessLicenseAssignment
func (m *ReprocessLicenseAssignmentRequestBuilder) Post(options *ReprocessLicenseAssignmentRequestBuilderPostOptions)(*ReprocessLicenseAssignmentResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewReprocessLicenseAssignmentResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ReprocessLicenseAssignmentResponse), nil
}
