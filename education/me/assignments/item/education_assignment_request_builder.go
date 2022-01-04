package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i608fa3df84918be40f4d961440af7a1df9cdc70cdcbcd01332c72421f74df332 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/categories"
    i6182b98a227da91892039c997d1742545d169f3ded56acfb9ebddc2b8754acb3 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/rubric"
    i710174f1350381dcd776625bf1b593be962f1d2c9b1e368ae0bd178d3495728f "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/submissions"
    i74b91723ec2102e06ba61f2eec38b8a338e711092e99e7a6121ae3287f6f5788 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/publish"
    i874a2b579f5eb892e97d9c0c157220e5118544d0dca2de101640f4b07a1c3fe3 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/resources"
    ie33a7d66ff2752d51b9a63d7758fc6c00305b565d547fa3e4e42f07c7a26bdec "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/setupresourcesfolder"
    i0a97b1183e4efc9029dbfdb7442c76c2d4fc13a243fc9089be3106c63fd1b7f2 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/submissions/item"
    i26284194e56af98cb4a7e2e347a5485f31cf2c535beac0528c6c3661c312a162 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/categories/item"
    i7b969696e68843bc2d2d30122cafa6f7eaf8a9843bf04feaf8693f3f794c0d94 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/resources/item"
)

// EducationAssignmentRequestBuilder builds and executes requests for operations under \education\me\assignments\{educationAssignment-id}
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
func (m *EducationAssignmentRequestBuilder) Categories()(*i608fa3df84918be40f4d961440af7a1df9cdc70cdcbcd01332c72421f74df332.CategoriesRequestBuilder) {
    return i608fa3df84918be40f4d961440af7a1df9cdc70cdcbcd01332c72421f74df332.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.assignments.item.categories.item collection
func (m *EducationAssignmentRequestBuilder) CategoriesById(id string)(*i26284194e56af98cb4a7e2e347a5485f31cf2c535beac0528c6c3661c312a162.EducationCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory_id"] = id
    }
    return i26284194e56af98cb4a7e2e347a5485f31cf2c535beac0528c6c3661c312a162.NewEducationCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEducationAssignmentRequestBuilderInternal instantiates a new EducationAssignmentRequestBuilder and sets the default values.
func NewEducationAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationAssignmentRequestBuilder) {
    m := &EducationAssignmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/me/assignments/{educationAssignment_id}{?select,expand}";
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
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
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
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEducationAssignment() }, nil)
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
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EducationAssignmentRequestBuilder) Publish()(*i74b91723ec2102e06ba61f2eec38b8a338e711092e99e7a6121ae3287f6f5788.PublishRequestBuilder) {
    return i74b91723ec2102e06ba61f2eec38b8a338e711092e99e7a6121ae3287f6f5788.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Resources()(*i874a2b579f5eb892e97d9c0c157220e5118544d0dca2de101640f4b07a1c3fe3.ResourcesRequestBuilder) {
    return i874a2b579f5eb892e97d9c0c157220e5118544d0dca2de101640f4b07a1c3fe3.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.assignments.item.resources.item collection
func (m *EducationAssignmentRequestBuilder) ResourcesById(id string)(*i7b969696e68843bc2d2d30122cafa6f7eaf8a9843bf04feaf8693f3f794c0d94.EducationAssignmentResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignmentResource_id"] = id
    }
    return i7b969696e68843bc2d2d30122cafa6f7eaf8a9843bf04feaf8693f3f794c0d94.NewEducationAssignmentResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Rubric()(*i6182b98a227da91892039c997d1742545d169f3ded56acfb9ebddc2b8754acb3.RubricRequestBuilder) {
    return i6182b98a227da91892039c997d1742545d169f3ded56acfb9ebddc2b8754acb3.NewRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) SetUpResourcesFolder()(*ie33a7d66ff2752d51b9a63d7758fc6c00305b565d547fa3e4e42f07c7a26bdec.SetUpResourcesFolderRequestBuilder) {
    return ie33a7d66ff2752d51b9a63d7758fc6c00305b565d547fa3e4e42f07c7a26bdec.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Submissions()(*i710174f1350381dcd776625bf1b593be962f1d2c9b1e368ae0bd178d3495728f.SubmissionsRequestBuilder) {
    return i710174f1350381dcd776625bf1b593be962f1d2c9b1e368ae0bd178d3495728f.NewSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.assignments.item.submissions.item collection
func (m *EducationAssignmentRequestBuilder) SubmissionsById(id string)(*i0a97b1183e4efc9029dbfdb7442c76c2d4fc13a243fc9089be3106c63fd1b7f2.EducationSubmissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmission_id"] = id
    }
    return i0a97b1183e4efc9029dbfdb7442c76c2d4fc13a243fc9089be3106c63fd1b7f2.NewEducationSubmissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
