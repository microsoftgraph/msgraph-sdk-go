package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i3d50a0849e57022c8564fae5dbbb2b471978d26a2db68975ed257f10e693e7cd "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/submissions"
    i42850d7acdbd3a8e01158ce3f91031a284414d79a1b8423753fe9921702335f6 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/setupresourcesfolder"
    i560e0f51eb0b15ee54800eeae829fe0350e27c3694c083160844bef71444f4f0 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/rubric"
    i6f1671e5ccc67a30dbb510d7c37809232aa6c2612b707440d4a323edf5cdceb1 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/resources"
    i87cf461bbbb0a4957678adb2652f5d882f0056914766c119c39f374d5e351968 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/categories"
    i9952290c89ab227cf2447a4dcefa0a460e4f57f5f1fa1a544211747659e33ab5 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/publish"
    i1683357a3b1fcd506250383b696db7903aedda5919d87479cc06dbe504619859 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/submissions/item"
    ic1a63ca1d60be548c163492f15ab00755d812529295e8c920991227ea011d404 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/categories/item"
    ieb0811fa685bb225bf1a7f78082f07573e816696c309fce9d12e426d809e5daa "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/resources/item"
)

// EducationAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.educationClass entity.
type EducationAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationAssignmentItemRequestBuilderDeleteOptions options for Delete
type EducationAssignmentItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationAssignmentItemRequestBuilderGetOptions options for Get
type EducationAssignmentItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EducationAssignmentItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationAssignmentItemRequestBuilderGetQueryParameters all assignments associated with this class. Nullable.
type EducationAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationAssignmentItemRequestBuilderPatchOptions options for Patch
type EducationAssignmentItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationAssignmentable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EducationAssignmentItemRequestBuilder) Categories()(*i87cf461bbbb0a4957678adb2652f5d882f0056914766c119c39f374d5e351968.CategoriesRequestBuilder) {
    return i87cf461bbbb0a4957678adb2652f5d882f0056914766c119c39f374d5e351968.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.assignments.item.categories.item collection
func (m *EducationAssignmentItemRequestBuilder) CategoriesById(id string)(*ic1a63ca1d60be548c163492f15ab00755d812529295e8c920991227ea011d404.EducationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory_id"] = id
    }
    return ic1a63ca1d60be548c163492f15ab00755d812529295e8c920991227ea011d404.NewEducationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEducationAssignmentItemRequestBuilderInternal instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewEducationAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationAssignmentItemRequestBuilder) {
    m := &EducationAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass_id}/assignments/{educationAssignment_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationAssignmentItemRequestBuilder instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewEducationAssignmentItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property assignments for education
func (m *EducationAssignmentItemRequestBuilder) CreateDeleteRequestInformation(options *EducationAssignmentItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation all assignments associated with this class. Nullable.
func (m *EducationAssignmentItemRequestBuilder) CreateGetRequestInformation(options *EducationAssignmentItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property assignments in education
func (m *EducationAssignmentItemRequestBuilder) CreatePatchRequestInformation(options *EducationAssignmentItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property assignments for education
func (m *EducationAssignmentItemRequestBuilder) Delete(options *EducationAssignmentItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get all assignments associated with this class. Nullable.
func (m *EducationAssignmentItemRequestBuilder) Get(options *EducationAssignmentItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationAssignmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateEducationAssignmentFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationAssignmentable), nil
}
// Patch update the navigation property assignments in education
func (m *EducationAssignmentItemRequestBuilder) Patch(options *EducationAssignmentItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EducationAssignmentItemRequestBuilder) Publish()(*i9952290c89ab227cf2447a4dcefa0a460e4f57f5f1fa1a544211747659e33ab5.PublishRequestBuilder) {
    return i9952290c89ab227cf2447a4dcefa0a460e4f57f5f1fa1a544211747659e33ab5.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentItemRequestBuilder) Resources()(*i6f1671e5ccc67a30dbb510d7c37809232aa6c2612b707440d4a323edf5cdceb1.ResourcesRequestBuilder) {
    return i6f1671e5ccc67a30dbb510d7c37809232aa6c2612b707440d4a323edf5cdceb1.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.assignments.item.resources.item collection
func (m *EducationAssignmentItemRequestBuilder) ResourcesById(id string)(*ieb0811fa685bb225bf1a7f78082f07573e816696c309fce9d12e426d809e5daa.EducationAssignmentResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignmentResource_id"] = id
    }
    return ieb0811fa685bb225bf1a7f78082f07573e816696c309fce9d12e426d809e5daa.NewEducationAssignmentResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationAssignmentItemRequestBuilder) Rubric()(*i560e0f51eb0b15ee54800eeae829fe0350e27c3694c083160844bef71444f4f0.RubricRequestBuilder) {
    return i560e0f51eb0b15ee54800eeae829fe0350e27c3694c083160844bef71444f4f0.NewRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentItemRequestBuilder) SetUpResourcesFolder()(*i42850d7acdbd3a8e01158ce3f91031a284414d79a1b8423753fe9921702335f6.SetUpResourcesFolderRequestBuilder) {
    return i42850d7acdbd3a8e01158ce3f91031a284414d79a1b8423753fe9921702335f6.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentItemRequestBuilder) Submissions()(*i3d50a0849e57022c8564fae5dbbb2b471978d26a2db68975ed257f10e693e7cd.SubmissionsRequestBuilder) {
    return i3d50a0849e57022c8564fae5dbbb2b471978d26a2db68975ed257f10e693e7cd.NewSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.assignments.item.submissions.item collection
func (m *EducationAssignmentItemRequestBuilder) SubmissionsById(id string)(*i1683357a3b1fcd506250383b696db7903aedda5919d87479cc06dbe504619859.EducationSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmission_id"] = id
    }
    return i1683357a3b1fcd506250383b696db7903aedda5919d87479cc06dbe504619859.NewEducationSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
