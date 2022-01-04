package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1548a312864f6a8093d6647950c34611e4ac02220716d0d28a7ed6ea68563e15 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/submissions/item/setupresourcesfolder"
    i3b9a3bc2cbb04d5774a8df86388bde86947a7b7d7c62e9d673dc25fa1f129f54 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/submissions/item/submittedresources"
    i4e04ce4811df69ad23b6cc639626ab84d384f66c5cbdc3157ea4eafa7746a09c "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/submissions/item/return_escaped"
    ia8cb4af48769798156b6d22780d66d72a7e7275c6bc9e7bc60a98007f81716fa "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/submissions/item/resources"
    ia9b02e54f8f0a390b69bcaf5d845b0f358c5f9201a3ec774316fdc87c9d21ca7 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/submissions/item/unsubmit"
    id00761bb21207f8c90cc6b9e657e4d6592705ed171fec38e1c0ce6a4e6b0fa7d "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/submissions/item/submit"
    id2c2532be9783e0b7091664a606705196b5b520c9fd05db1ea40f6cdf4bc5944 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/submissions/item/outcomes"
    id76ff41c49808e9ab3c419e39394c8b5a49d6375f73b75e7363fb291151a2bed "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/submissions/item/reassign"
    i363dd35cda83aa9a020591fb0f4fd184033fd6ab5dbc6a251a3af8f414725817 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/submissions/item/submittedresources/item"
    i71aa111a74a29d833b5fd0653f7a2d786fe91774ffbbc1f1ed0064bc02923991 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/submissions/item/outcomes/item"
    ia7af5a2c3efa196597d6805e5921ce27fb2948e0a38a819bfc0b0c6cf8eef9da "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/submissions/item/resources/item"
)

// EducationSubmissionRequestBuilder builds and executes requests for operations under \education\users\{educationUser-id}\assignments\{educationAssignment-id}\submissions\{educationSubmission-id}
type EducationSubmissionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationSubmissionRequestBuilderDeleteOptions options for Delete
type EducationSubmissionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationSubmissionRequestBuilderGetOptions options for Get
type EducationSubmissionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EducationSubmissionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationSubmissionRequestBuilderGetQueryParameters once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
type EducationSubmissionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationSubmissionRequestBuilderPatchOptions options for Patch
type EducationSubmissionRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationSubmission;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewEducationSubmissionRequestBuilderInternal instantiates a new EducationSubmissionRequestBuilder and sets the default values.
func NewEducationSubmissionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSubmissionRequestBuilder) {
    m := &EducationSubmissionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser_id}/assignments/{educationAssignment_id}/submissions/{educationSubmission_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationSubmissionRequestBuilder instantiates a new EducationSubmissionRequestBuilder and sets the default values.
func NewEducationSubmissionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSubmissionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationSubmissionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionRequestBuilder) CreateDeleteRequestInformation(options *EducationSubmissionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionRequestBuilder) CreateGetRequestInformation(options *EducationSubmissionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionRequestBuilder) CreatePatchRequestInformation(options *EducationSubmissionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionRequestBuilder) Delete(options *EducationSubmissionRequestBuilderDeleteOptions)(error) {
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
// Get once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionRequestBuilder) Get(options *EducationSubmissionRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationSubmission, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEducationSubmission() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationSubmission), nil
}
func (m *EducationSubmissionRequestBuilder) Outcomes()(*id2c2532be9783e0b7091664a606705196b5b520c9fd05db1ea40f6cdf4bc5944.OutcomesRequestBuilder) {
    return id2c2532be9783e0b7091664a606705196b5b520c9fd05db1ea40f6cdf4bc5944.NewOutcomesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OutcomesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.assignments.item.submissions.item.outcomes.item collection
func (m *EducationSubmissionRequestBuilder) OutcomesById(id string)(*i71aa111a74a29d833b5fd0653f7a2d786fe91774ffbbc1f1ed0064bc02923991.EducationOutcomeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationOutcome_id"] = id
    }
    return i71aa111a74a29d833b5fd0653f7a2d786fe91774ffbbc1f1ed0064bc02923991.NewEducationOutcomeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionRequestBuilder) Patch(options *EducationSubmissionRequestBuilderPatchOptions)(error) {
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
func (m *EducationSubmissionRequestBuilder) Reassign()(*id76ff41c49808e9ab3c419e39394c8b5a49d6375f73b75e7363fb291151a2bed.ReassignRequestBuilder) {
    return id76ff41c49808e9ab3c419e39394c8b5a49d6375f73b75e7363fb291151a2bed.NewReassignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Resources()(*ia8cb4af48769798156b6d22780d66d72a7e7275c6bc9e7bc60a98007f81716fa.ResourcesRequestBuilder) {
    return ia8cb4af48769798156b6d22780d66d72a7e7275c6bc9e7bc60a98007f81716fa.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.assignments.item.submissions.item.resources.item collection
func (m *EducationSubmissionRequestBuilder) ResourcesById(id string)(*ia7af5a2c3efa196597d6805e5921ce27fb2948e0a38a819bfc0b0c6cf8eef9da.EducationSubmissionResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return ia7af5a2c3efa196597d6805e5921ce27fb2948e0a38a819bfc0b0c6cf8eef9da.NewEducationSubmissionResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Return_escaped()(*i4e04ce4811df69ad23b6cc639626ab84d384f66c5cbdc3157ea4eafa7746a09c.ReturnRequestBuilder) {
    return i4e04ce4811df69ad23b6cc639626ab84d384f66c5cbdc3157ea4eafa7746a09c.NewReturnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) SetUpResourcesFolder()(*i1548a312864f6a8093d6647950c34611e4ac02220716d0d28a7ed6ea68563e15.SetUpResourcesFolderRequestBuilder) {
    return i1548a312864f6a8093d6647950c34611e4ac02220716d0d28a7ed6ea68563e15.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Submit()(*id00761bb21207f8c90cc6b9e657e4d6592705ed171fec38e1c0ce6a4e6b0fa7d.SubmitRequestBuilder) {
    return id00761bb21207f8c90cc6b9e657e4d6592705ed171fec38e1c0ce6a4e6b0fa7d.NewSubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) SubmittedResources()(*i3b9a3bc2cbb04d5774a8df86388bde86947a7b7d7c62e9d673dc25fa1f129f54.SubmittedResourcesRequestBuilder) {
    return i3b9a3bc2cbb04d5774a8df86388bde86947a7b7d7c62e9d673dc25fa1f129f54.NewSubmittedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmittedResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.assignments.item.submissions.item.submittedResources.item collection
func (m *EducationSubmissionRequestBuilder) SubmittedResourcesById(id string)(*i363dd35cda83aa9a020591fb0f4fd184033fd6ab5dbc6a251a3af8f414725817.EducationSubmissionResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return i363dd35cda83aa9a020591fb0f4fd184033fd6ab5dbc6a251a3af8f414725817.NewEducationSubmissionResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Unsubmit()(*ia9b02e54f8f0a390b69bcaf5d845b0f358c5f9201a3ec774316fdc87c9d21ca7.UnsubmitRequestBuilder) {
    return ia9b02e54f8f0a390b69bcaf5d845b0f358c5f9201a3ec774316fdc87c9d21ca7.NewUnsubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
