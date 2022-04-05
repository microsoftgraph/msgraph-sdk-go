package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
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

// EducationSubmissionItemRequestBuilder provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
type EducationSubmissionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationSubmissionItemRequestBuilderDeleteOptions options for Delete
type EducationSubmissionItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// EducationSubmissionItemRequestBuilderGetOptions options for Get
type EducationSubmissionItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *EducationSubmissionItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// EducationSubmissionItemRequestBuilderGetQueryParameters once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
type EducationSubmissionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationSubmissionItemRequestBuilderPatchOptions options for Patch
type EducationSubmissionItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewEducationSubmissionItemRequestBuilderInternal instantiates a new EducationSubmissionItemRequestBuilder and sets the default values.
func NewEducationSubmissionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationSubmissionItemRequestBuilder) {
    m := &EducationSubmissionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser_id}/assignments/{educationAssignment_id}/submissions/{educationSubmission_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationSubmissionItemRequestBuilder instantiates a new EducationSubmissionItemRequestBuilder and sets the default values.
func NewEducationSubmissionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationSubmissionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationSubmissionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property submissions for education
func (m *EducationSubmissionItemRequestBuilder) CreateDeleteRequestInformation(options *EducationSubmissionItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionItemRequestBuilder) CreateGetRequestInformation(options *EducationSubmissionItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property submissions in education
func (m *EducationSubmissionItemRequestBuilder) CreatePatchRequestInformation(options *EducationSubmissionItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property submissions for education
func (m *EducationSubmissionItemRequestBuilder) Delete(options *EducationSubmissionItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionItemRequestBuilder) Get(options *EducationSubmissionItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationSubmissionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable), nil
}
// Outcomes the outcomes property
func (m *EducationSubmissionItemRequestBuilder) Outcomes()(*id2c2532be9783e0b7091664a606705196b5b520c9fd05db1ea40f6cdf4bc5944.OutcomesRequestBuilder) {
    return id2c2532be9783e0b7091664a606705196b5b520c9fd05db1ea40f6cdf4bc5944.NewOutcomesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OutcomesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.assignments.item.submissions.item.outcomes.item collection
func (m *EducationSubmissionItemRequestBuilder) OutcomesById(id string)(*i71aa111a74a29d833b5fd0653f7a2d786fe91774ffbbc1f1ed0064bc02923991.EducationOutcomeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationOutcome_id"] = id
    }
    return i71aa111a74a29d833b5fd0653f7a2d786fe91774ffbbc1f1ed0064bc02923991.NewEducationOutcomeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property submissions in education
func (m *EducationSubmissionItemRequestBuilder) Patch(options *EducationSubmissionItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Reassign the reassign property
func (m *EducationSubmissionItemRequestBuilder) Reassign()(*id76ff41c49808e9ab3c419e39394c8b5a49d6375f73b75e7363fb291151a2bed.ReassignRequestBuilder) {
    return id76ff41c49808e9ab3c419e39394c8b5a49d6375f73b75e7363fb291151a2bed.NewReassignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Resources the resources property
func (m *EducationSubmissionItemRequestBuilder) Resources()(*ia8cb4af48769798156b6d22780d66d72a7e7275c6bc9e7bc60a98007f81716fa.ResourcesRequestBuilder) {
    return ia8cb4af48769798156b6d22780d66d72a7e7275c6bc9e7bc60a98007f81716fa.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.assignments.item.submissions.item.resources.item collection
func (m *EducationSubmissionItemRequestBuilder) ResourcesById(id string)(*ia7af5a2c3efa196597d6805e5921ce27fb2948e0a38a819bfc0b0c6cf8eef9da.EducationSubmissionResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return ia7af5a2c3efa196597d6805e5921ce27fb2948e0a38a819bfc0b0c6cf8eef9da.NewEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Return_escaped the return property
func (m *EducationSubmissionItemRequestBuilder) Return_escaped()(*i4e04ce4811df69ad23b6cc639626ab84d384f66c5cbdc3157ea4eafa7746a09c.ReturnRequestBuilder) {
    return i4e04ce4811df69ad23b6cc639626ab84d384f66c5cbdc3157ea4eafa7746a09c.NewReturnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetUpResourcesFolder the setUpResourcesFolder property
func (m *EducationSubmissionItemRequestBuilder) SetUpResourcesFolder()(*i1548a312864f6a8093d6647950c34611e4ac02220716d0d28a7ed6ea68563e15.SetUpResourcesFolderRequestBuilder) {
    return i1548a312864f6a8093d6647950c34611e4ac02220716d0d28a7ed6ea68563e15.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Submit the submit property
func (m *EducationSubmissionItemRequestBuilder) Submit()(*id00761bb21207f8c90cc6b9e657e4d6592705ed171fec38e1c0ce6a4e6b0fa7d.SubmitRequestBuilder) {
    return id00761bb21207f8c90cc6b9e657e4d6592705ed171fec38e1c0ce6a4e6b0fa7d.NewSubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmittedResources the submittedResources property
func (m *EducationSubmissionItemRequestBuilder) SubmittedResources()(*i3b9a3bc2cbb04d5774a8df86388bde86947a7b7d7c62e9d673dc25fa1f129f54.SubmittedResourcesRequestBuilder) {
    return i3b9a3bc2cbb04d5774a8df86388bde86947a7b7d7c62e9d673dc25fa1f129f54.NewSubmittedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmittedResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.assignments.item.submissions.item.submittedResources.item collection
func (m *EducationSubmissionItemRequestBuilder) SubmittedResourcesById(id string)(*i363dd35cda83aa9a020591fb0f4fd184033fd6ab5dbc6a251a3af8f414725817.EducationSubmissionResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return i363dd35cda83aa9a020591fb0f4fd184033fd6ab5dbc6a251a3af8f414725817.NewEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unsubmit the unsubmit property
func (m *EducationSubmissionItemRequestBuilder) Unsubmit()(*ia9b02e54f8f0a390b69bcaf5d845b0f358c5f9201a3ec774316fdc87c9d21ca7.UnsubmitRequestBuilder) {
    return ia9b02e54f8f0a390b69bcaf5d845b0f358c5f9201a3ec774316fdc87c9d21ca7.NewUnsubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
