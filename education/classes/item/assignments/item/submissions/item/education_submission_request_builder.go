package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i3c6015cf8457d8d8ec01d019f3dbc202765651c2130e73641cf6b06caeb2d228 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/submissions/item/setupresourcesfolder"
    i4897ecfe261ff150792387d2c4b3c93a4a0c1fa8d17924ec2e00dd6cf21d022d "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/submissions/item/return_escaped"
    i685e1486255b2d524092d2de81f1a7d1d38632cbb99a29975e93cc433375ed77 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/submissions/item/submit"
    ic9e6f044df077c5c49203af4cbb1cb4bdd990da9a1e6a2c50a189335ae8c62fc "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/submissions/item/outcomes"
    id25fa828362bc547da297fba40d2cdbe268a1fd50d31a4ead0dd41ffa77fdfc6 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/submissions/item/resources"
    id45b66ebeec3910c7bc64c11a6669bb5fd10605cbdd08bde69bcaa49dd9ab2a7 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/submissions/item/unsubmit"
    id595ffd3ed593bffbb73e849bb31881e531d536f5c57d0b37c7933bb34ad1ad0 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/submissions/item/submittedresources"
    i9a0443e0b2ffd58714a9cae0542bdec49c7a1c21111a0cba91c45aa013014f81 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/submissions/item/outcomes/item"
    idf6da69a1128e56db2a075a35c90dfb9d769744b063f36a369e943680fa35795 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/submissions/item/resources/item"
    if2b3a1bb63b3f0f978027b12a06272583c8a6c2ebe1a551f1e1b72a751ef4004 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/submissions/item/submittedresources/item"
)

// Builds and executes requests for operations under \education\classes\{educationClass-id}\assignments\{educationAssignment-id}\submissions\{educationSubmission-id}
type EducationSubmissionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type EducationSubmissionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
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
// Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
type EducationSubmissionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
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
// Instantiates a new EducationSubmissionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEducationSubmissionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSubmissionRequestBuilder) {
    m := &EducationSubmissionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass_id}/assignments/{educationAssignment_id}/submissions/{educationSubmission_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new EducationSubmissionRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEducationSubmissionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSubmissionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationSubmissionRequestBuilderInternal(urlParams, requestAdapter)
}
// Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
// Parameters:
//  - options : Options for the request
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
// Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
// Parameters:
//  - options : Options for the request
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
// Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
// Parameters:
//  - options : Options for the request
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
// Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
// Parameters:
//  - options : Options for the request
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
// Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
// Parameters:
//  - options : Options for the request
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
func (m *EducationSubmissionRequestBuilder) Outcomes()(*ic9e6f044df077c5c49203af4cbb1cb4bdd990da9a1e6a2c50a189335ae8c62fc.OutcomesRequestBuilder) {
    return ic9e6f044df077c5c49203af4cbb1cb4bdd990da9a1e6a2c50a189335ae8c62fc.NewOutcomesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.assignments.item.submissions.item.outcomes.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EducationSubmissionRequestBuilder) OutcomesById(id string)(*i9a0443e0b2ffd58714a9cae0542bdec49c7a1c21111a0cba91c45aa013014f81.EducationOutcomeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationOutcome_id"] = id
    }
    return i9a0443e0b2ffd58714a9cae0542bdec49c7a1c21111a0cba91c45aa013014f81.NewEducationOutcomeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
// Parameters:
//  - options : Options for the request
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
func (m *EducationSubmissionRequestBuilder) Resources()(*id25fa828362bc547da297fba40d2cdbe268a1fd50d31a4ead0dd41ffa77fdfc6.ResourcesRequestBuilder) {
    return id25fa828362bc547da297fba40d2cdbe268a1fd50d31a4ead0dd41ffa77fdfc6.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.assignments.item.submissions.item.resources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EducationSubmissionRequestBuilder) ResourcesById(id string)(*idf6da69a1128e56db2a075a35c90dfb9d769744b063f36a369e943680fa35795.EducationSubmissionResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return idf6da69a1128e56db2a075a35c90dfb9d769744b063f36a369e943680fa35795.NewEducationSubmissionResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Return_escaped()(*i4897ecfe261ff150792387d2c4b3c93a4a0c1fa8d17924ec2e00dd6cf21d022d.ReturnRequestBuilder) {
    return i4897ecfe261ff150792387d2c4b3c93a4a0c1fa8d17924ec2e00dd6cf21d022d.NewReturnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) SetUpResourcesFolder()(*i3c6015cf8457d8d8ec01d019f3dbc202765651c2130e73641cf6b06caeb2d228.SetUpResourcesFolderRequestBuilder) {
    return i3c6015cf8457d8d8ec01d019f3dbc202765651c2130e73641cf6b06caeb2d228.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Submit()(*i685e1486255b2d524092d2de81f1a7d1d38632cbb99a29975e93cc433375ed77.SubmitRequestBuilder) {
    return i685e1486255b2d524092d2de81f1a7d1d38632cbb99a29975e93cc433375ed77.NewSubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) SubmittedResources()(*id595ffd3ed593bffbb73e849bb31881e531d536f5c57d0b37c7933bb34ad1ad0.SubmittedResourcesRequestBuilder) {
    return id595ffd3ed593bffbb73e849bb31881e531d536f5c57d0b37c7933bb34ad1ad0.NewSubmittedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.assignments.item.submissions.item.submittedResources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EducationSubmissionRequestBuilder) SubmittedResourcesById(id string)(*if2b3a1bb63b3f0f978027b12a06272583c8a6c2ebe1a551f1e1b72a751ef4004.EducationSubmissionResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return if2b3a1bb63b3f0f978027b12a06272583c8a6c2ebe1a551f1e1b72a751ef4004.NewEducationSubmissionResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Unsubmit()(*id45b66ebeec3910c7bc64c11a6669bb5fd10605cbdd08bde69bcaa49dd9ab2a7.UnsubmitRequestBuilder) {
    return id45b66ebeec3910c7bc64c11a6669bb5fd10605cbdd08bde69bcaa49dd9ab2a7.NewUnsubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
