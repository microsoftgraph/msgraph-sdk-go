package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i1348b4589375e5abe2aadb202f6333594b5d1751d170bf9ca36b5dbb12d0f961 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/submissions/item/reassign"
    i22de02c8ebcc8858f70a597c50e10aaabef478e949434d07071c431cf794e48b "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/submissions/item/return_escaped"
    i7d297e6b19b3a526e5ff3a4e05a3c51ebaa8b75db93203b5b44c25ca525e63b7 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/submissions/item/unsubmit"
    i88b145b736a4defd7f959baef7d247b6f9780a06624a878b6eb6fe328be9e03e "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/submissions/item/resources"
    ia875d5e362b74f9d96b74a33d07714a43cac0c9f46e9d284c4a3b9563b2648be "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/submissions/item/submittedresources"
    ibb5e09f845b30b48220a8c2c52e0003759df19239c2b699bcd05f8770b8e2736 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/submissions/item/submit"
    ic61ae04c6c878f0744cbb425fafa87a99592b6ff25d78fd0e783439f81c21ca9 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/submissions/item/outcomes"
    ie7b1f1bad52c20f6a6c0c1705d25531013c14a3669432c5a69724bd67b0509fa "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/submissions/item/setupresourcesfolder"
    i081bee1a5d3b8c208a665428688091c98598de447b3cb5f38b1ba25e23db8a6b "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/submissions/item/outcomes/item"
    i25aed5774acd9ec91ff415fc80f838cc483c55d18701532dc2e7f9480f6170bc "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/submissions/item/submittedresources/item"
    if70d650fff38d98927e0802e08152564ebdba3fae5438f03c5412678c4471192 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/submissions/item/resources/item"
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
    m.urlTemplate = "{+baseurl}/education/me/assignments/{educationAssignment_id}/submissions/{educationSubmission_id}{?select,expand}";
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
func (m *EducationSubmissionItemRequestBuilder) Outcomes()(*ic61ae04c6c878f0744cbb425fafa87a99592b6ff25d78fd0e783439f81c21ca9.OutcomesRequestBuilder) {
    return ic61ae04c6c878f0744cbb425fafa87a99592b6ff25d78fd0e783439f81c21ca9.NewOutcomesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OutcomesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.assignments.item.submissions.item.outcomes.item collection
func (m *EducationSubmissionItemRequestBuilder) OutcomesById(id string)(*i081bee1a5d3b8c208a665428688091c98598de447b3cb5f38b1ba25e23db8a6b.EducationOutcomeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationOutcome_id"] = id
    }
    return i081bee1a5d3b8c208a665428688091c98598de447b3cb5f38b1ba25e23db8a6b.NewEducationOutcomeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EducationSubmissionItemRequestBuilder) Reassign()(*i1348b4589375e5abe2aadb202f6333594b5d1751d170bf9ca36b5dbb12d0f961.ReassignRequestBuilder) {
    return i1348b4589375e5abe2aadb202f6333594b5d1751d170bf9ca36b5dbb12d0f961.NewReassignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Resources the resources property
func (m *EducationSubmissionItemRequestBuilder) Resources()(*i88b145b736a4defd7f959baef7d247b6f9780a06624a878b6eb6fe328be9e03e.ResourcesRequestBuilder) {
    return i88b145b736a4defd7f959baef7d247b6f9780a06624a878b6eb6fe328be9e03e.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.assignments.item.submissions.item.resources.item collection
func (m *EducationSubmissionItemRequestBuilder) ResourcesById(id string)(*if70d650fff38d98927e0802e08152564ebdba3fae5438f03c5412678c4471192.EducationSubmissionResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return if70d650fff38d98927e0802e08152564ebdba3fae5438f03c5412678c4471192.NewEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Return_escaped the return property
func (m *EducationSubmissionItemRequestBuilder) Return_escaped()(*i22de02c8ebcc8858f70a597c50e10aaabef478e949434d07071c431cf794e48b.ReturnRequestBuilder) {
    return i22de02c8ebcc8858f70a597c50e10aaabef478e949434d07071c431cf794e48b.NewReturnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetUpResourcesFolder the setUpResourcesFolder property
func (m *EducationSubmissionItemRequestBuilder) SetUpResourcesFolder()(*ie7b1f1bad52c20f6a6c0c1705d25531013c14a3669432c5a69724bd67b0509fa.SetUpResourcesFolderRequestBuilder) {
    return ie7b1f1bad52c20f6a6c0c1705d25531013c14a3669432c5a69724bd67b0509fa.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Submit the submit property
func (m *EducationSubmissionItemRequestBuilder) Submit()(*ibb5e09f845b30b48220a8c2c52e0003759df19239c2b699bcd05f8770b8e2736.SubmitRequestBuilder) {
    return ibb5e09f845b30b48220a8c2c52e0003759df19239c2b699bcd05f8770b8e2736.NewSubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmittedResources the submittedResources property
func (m *EducationSubmissionItemRequestBuilder) SubmittedResources()(*ia875d5e362b74f9d96b74a33d07714a43cac0c9f46e9d284c4a3b9563b2648be.SubmittedResourcesRequestBuilder) {
    return ia875d5e362b74f9d96b74a33d07714a43cac0c9f46e9d284c4a3b9563b2648be.NewSubmittedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmittedResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.assignments.item.submissions.item.submittedResources.item collection
func (m *EducationSubmissionItemRequestBuilder) SubmittedResourcesById(id string)(*i25aed5774acd9ec91ff415fc80f838cc483c55d18701532dc2e7f9480f6170bc.EducationSubmissionResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return i25aed5774acd9ec91ff415fc80f838cc483c55d18701532dc2e7f9480f6170bc.NewEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unsubmit the unsubmit property
func (m *EducationSubmissionItemRequestBuilder) Unsubmit()(*i7d297e6b19b3a526e5ff3a4e05a3c51ebaa8b75db93203b5b44c25ca525e63b7.UnsubmitRequestBuilder) {
    return i7d297e6b19b3a526e5ff3a4e05a3c51ebaa8b75db93203b5b44c25ca525e63b7.NewUnsubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
