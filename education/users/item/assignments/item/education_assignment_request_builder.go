package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i4a60e659b3e1427ff2f92b71f44e26c8ac1ac819ee646533021117deeeb6b88a "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/rubric"
    i7b1c280606a6564228514ab9e1b7f963f46c753907436ec516230878d54adc8f "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/categories"
    i88086f8017be5762df8cc58a527cb8b9ece6f57e0c9ce6c6a7b1682cb8a238cc "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/resources"
    ib979357b11252b944c152b3aa6f90a6415f1f60345dd67d0804b319a1f2123ba "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/submissions"
    ic8a76de5ffb5d6e2de378abe7f4341d25a11d4ea88cad7b35c31b72f63d17e7a "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/setupresourcesfolder"
    ied86aa882920629ba539d3c8055e54c93f45dfa9b4f250108d4b4d3bbc80975a "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/publish"
    i2d4596e7585085aa1bcbd1dfc3b2ce4826b6afe4b02ed3c80e40d2c1de2dd769 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/submissions/item"
    i37bd40aa3d07e5fb03511b981c264fa6dcad03889d85b2483191049178483b0d "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/resources/item"
    i5914468bbfd7cf156e785aa5b7c676dd4a5c8f64f204e13bca5907a261b89c6c "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/categories/item"
)

// EducationAssignmentRequestBuilder builds and executes requests for operations under \education\users\{educationUser-id}\assignments\{educationAssignment-id}
type EducationAssignmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationAssignmentRequestBuilderDeleteOptions options for Delete
type EducationAssignmentRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationAssignmentRequestBuilderGetOptions options for Get
type EducationAssignmentRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EducationAssignmentRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationAssignmentRequestBuilderGetQueryParameters assignments belonging to the user.
type EducationAssignmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationAssignmentRequestBuilderPatchOptions options for Patch
type EducationAssignmentRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationAssignment;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EducationAssignmentRequestBuilder) Categories()(*i7b1c280606a6564228514ab9e1b7f963f46c753907436ec516230878d54adc8f.CategoriesRequestBuilder) {
    return i7b1c280606a6564228514ab9e1b7f963f46c753907436ec516230878d54adc8f.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.assignments.item.categories.item collection
func (m *EducationAssignmentRequestBuilder) CategoriesById(id string)(*i5914468bbfd7cf156e785aa5b7c676dd4a5c8f64f204e13bca5907a261b89c6c.EducationCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory_id"] = id
    }
    return i5914468bbfd7cf156e785aa5b7c676dd4a5c8f64f204e13bca5907a261b89c6c.NewEducationCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEducationAssignmentRequestBuilderInternal instantiates a new EducationAssignmentRequestBuilder and sets the default values.
func NewEducationAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationAssignmentRequestBuilder) {
    m := &EducationAssignmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser_id}/assignments/{educationAssignment_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationAssignmentRequestBuilder instantiates a new EducationAssignmentRequestBuilder and sets the default values.
func NewEducationAssignmentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation assignments belonging to the user.
func (m *EducationAssignmentRequestBuilder) CreateDeleteRequestInformation(options *EducationAssignmentRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation assignments belonging to the user.
func (m *EducationAssignmentRequestBuilder) CreateGetRequestInformation(options *EducationAssignmentRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
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
// CreatePatchRequestInformation assignments belonging to the user.
func (m *EducationAssignmentRequestBuilder) CreatePatchRequestInformation(options *EducationAssignmentRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
// Delete assignments belonging to the user.
func (m *EducationAssignmentRequestBuilder) Delete(options *EducationAssignmentRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get assignments belonging to the user.
func (m *EducationAssignmentRequestBuilder) Get(options *EducationAssignmentRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationAssignment, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEducationAssignment() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationAssignment), nil
}
// Patch assignments belonging to the user.
func (m *EducationAssignmentRequestBuilder) Patch(options *EducationAssignmentRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EducationAssignmentRequestBuilder) Publish()(*ied86aa882920629ba539d3c8055e54c93f45dfa9b4f250108d4b4d3bbc80975a.PublishRequestBuilder) {
    return ied86aa882920629ba539d3c8055e54c93f45dfa9b4f250108d4b4d3bbc80975a.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Resources()(*i88086f8017be5762df8cc58a527cb8b9ece6f57e0c9ce6c6a7b1682cb8a238cc.ResourcesRequestBuilder) {
    return i88086f8017be5762df8cc58a527cb8b9ece6f57e0c9ce6c6a7b1682cb8a238cc.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.assignments.item.resources.item collection
func (m *EducationAssignmentRequestBuilder) ResourcesById(id string)(*i37bd40aa3d07e5fb03511b981c264fa6dcad03889d85b2483191049178483b0d.EducationAssignmentResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignmentResource_id"] = id
    }
    return i37bd40aa3d07e5fb03511b981c264fa6dcad03889d85b2483191049178483b0d.NewEducationAssignmentResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Rubric()(*i4a60e659b3e1427ff2f92b71f44e26c8ac1ac819ee646533021117deeeb6b88a.RubricRequestBuilder) {
    return i4a60e659b3e1427ff2f92b71f44e26c8ac1ac819ee646533021117deeeb6b88a.NewRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) SetUpResourcesFolder()(*ic8a76de5ffb5d6e2de378abe7f4341d25a11d4ea88cad7b35c31b72f63d17e7a.SetUpResourcesFolderRequestBuilder) {
    return ic8a76de5ffb5d6e2de378abe7f4341d25a11d4ea88cad7b35c31b72f63d17e7a.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Submissions()(*ib979357b11252b944c152b3aa6f90a6415f1f60345dd67d0804b319a1f2123ba.SubmissionsRequestBuilder) {
    return ib979357b11252b944c152b3aa6f90a6415f1f60345dd67d0804b319a1f2123ba.NewSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.assignments.item.submissions.item collection
func (m *EducationAssignmentRequestBuilder) SubmissionsById(id string)(*i2d4596e7585085aa1bcbd1dfc3b2ce4826b6afe4b02ed3c80e40d2c1de2dd769.EducationSubmissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmission_id"] = id
    }
    return i2d4596e7585085aa1bcbd1dfc3b2ce4826b6afe4b02ed3c80e40d2c1de2dd769.NewEducationSubmissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
