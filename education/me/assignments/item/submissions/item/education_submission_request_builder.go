package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
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

// EducationSubmissionRequestBuilder builds and executes requests for operations under \education\me\assignments\{educationAssignment-id}\submissions\{educationSubmission-id}
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
    m.urlTemplate = "{+baseurl}/education/me/assignments/{educationAssignment_id}/submissions/{educationSubmission_id}{?select,expand}";
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
func (m *EducationSubmissionRequestBuilder) Outcomes()(*ic61ae04c6c878f0744cbb425fafa87a99592b6ff25d78fd0e783439f81c21ca9.OutcomesRequestBuilder) {
    return ic61ae04c6c878f0744cbb425fafa87a99592b6ff25d78fd0e783439f81c21ca9.NewOutcomesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OutcomesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.assignments.item.submissions.item.outcomes.item collection
func (m *EducationSubmissionRequestBuilder) OutcomesById(id string)(*i081bee1a5d3b8c208a665428688091c98598de447b3cb5f38b1ba25e23db8a6b.EducationOutcomeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationOutcome_id"] = id
    }
    return i081bee1a5d3b8c208a665428688091c98598de447b3cb5f38b1ba25e23db8a6b.NewEducationOutcomeRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EducationSubmissionRequestBuilder) Reassign()(*i1348b4589375e5abe2aadb202f6333594b5d1751d170bf9ca36b5dbb12d0f961.ReassignRequestBuilder) {
    return i1348b4589375e5abe2aadb202f6333594b5d1751d170bf9ca36b5dbb12d0f961.NewReassignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Resources()(*i88b145b736a4defd7f959baef7d247b6f9780a06624a878b6eb6fe328be9e03e.ResourcesRequestBuilder) {
    return i88b145b736a4defd7f959baef7d247b6f9780a06624a878b6eb6fe328be9e03e.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.assignments.item.submissions.item.resources.item collection
func (m *EducationSubmissionRequestBuilder) ResourcesById(id string)(*if70d650fff38d98927e0802e08152564ebdba3fae5438f03c5412678c4471192.EducationSubmissionResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return if70d650fff38d98927e0802e08152564ebdba3fae5438f03c5412678c4471192.NewEducationSubmissionResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Return_escaped()(*i22de02c8ebcc8858f70a597c50e10aaabef478e949434d07071c431cf794e48b.ReturnRequestBuilder) {
    return i22de02c8ebcc8858f70a597c50e10aaabef478e949434d07071c431cf794e48b.NewReturnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) SetUpResourcesFolder()(*ie7b1f1bad52c20f6a6c0c1705d25531013c14a3669432c5a69724bd67b0509fa.SetUpResourcesFolderRequestBuilder) {
    return ie7b1f1bad52c20f6a6c0c1705d25531013c14a3669432c5a69724bd67b0509fa.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Submit()(*ibb5e09f845b30b48220a8c2c52e0003759df19239c2b699bcd05f8770b8e2736.SubmitRequestBuilder) {
    return ibb5e09f845b30b48220a8c2c52e0003759df19239c2b699bcd05f8770b8e2736.NewSubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) SubmittedResources()(*ia875d5e362b74f9d96b74a33d07714a43cac0c9f46e9d284c4a3b9563b2648be.SubmittedResourcesRequestBuilder) {
    return ia875d5e362b74f9d96b74a33d07714a43cac0c9f46e9d284c4a3b9563b2648be.NewSubmittedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmittedResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.assignments.item.submissions.item.submittedResources.item collection
func (m *EducationSubmissionRequestBuilder) SubmittedResourcesById(id string)(*i25aed5774acd9ec91ff415fc80f838cc483c55d18701532dc2e7f9480f6170bc.EducationSubmissionResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return i25aed5774acd9ec91ff415fc80f838cc483c55d18701532dc2e7f9480f6170bc.NewEducationSubmissionResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Unsubmit()(*i7d297e6b19b3a526e5ff3a4e05a3c51ebaa8b75db93203b5b44c25ca525e63b7.UnsubmitRequestBuilder) {
    return i7d297e6b19b3a526e5ff3a4e05a3c51ebaa8b75db93203b5b44c25ca525e63b7.NewUnsubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
