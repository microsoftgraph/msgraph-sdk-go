package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i09fc0a778050440e74342b699ac672725441e7cf37365559ae6a2eb2457f6633 "github.com/microsoftgraph/msgraph-sdk-go/privacy/subjectrightsrequests/item/getfinalreport"
    i746ebfc2054651d0c33c207b53a9353452215cae640e13084d03015de282cf39 "github.com/microsoftgraph/msgraph-sdk-go/privacy/subjectrightsrequests/item/team"
    i793fe4cfdde8af93f5cd2898a89580574ca6d3e95e1050acfda2f70d4ef9990b "github.com/microsoftgraph/msgraph-sdk-go/privacy/subjectrightsrequests/item/notes"
    ie490ccc81823dd6adcd276d7a2ff20dfdf201522ad6cac044881cf7e0b59a420 "github.com/microsoftgraph/msgraph-sdk-go/privacy/subjectrightsrequests/item/getfinalattachment"
    i01c49ff8ee7f291e6ef03661dd6de1defbca61ef00e7cc93142b70cd237aae9e "github.com/microsoftgraph/msgraph-sdk-go/privacy/subjectrightsrequests/item/notes/item"
)

// Builds and executes requests for operations under \privacy\subjectRightsRequests\{subjectRightsRequest-id}
type SubjectRightsRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type SubjectRightsRequestRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type SubjectRightsRequestRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SubjectRightsRequestRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get subjectRightsRequests from privacy
type SubjectRightsRequestRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type SubjectRightsRequestRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SubjectRightsRequest;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new SubjectRightsRequestRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSubjectRightsRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SubjectRightsRequestRequestBuilder) {
    m := &SubjectRightsRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privacy/subjectRightsRequests/{subjectRightsRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new SubjectRightsRequestRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSubjectRightsRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SubjectRightsRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubjectRightsRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete navigation property subjectRightsRequests for privacy
// Parameters:
//  - options : Options for the request
func (m *SubjectRightsRequestRequestBuilder) CreateDeleteRequestInformation(options *SubjectRightsRequestRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get subjectRightsRequests from privacy
// Parameters:
//  - options : Options for the request
func (m *SubjectRightsRequestRequestBuilder) CreateGetRequestInformation(options *SubjectRightsRequestRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Update the navigation property subjectRightsRequests in privacy
// Parameters:
//  - options : Options for the request
func (m *SubjectRightsRequestRequestBuilder) CreatePatchRequestInformation(options *SubjectRightsRequestRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete navigation property subjectRightsRequests for privacy
// Parameters:
//  - options : Options for the request
func (m *SubjectRightsRequestRequestBuilder) Delete(options *SubjectRightsRequestRequestBuilderDeleteOptions)(error) {
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
// Get subjectRightsRequests from privacy
// Parameters:
//  - options : Options for the request
func (m *SubjectRightsRequestRequestBuilder) Get(options *SubjectRightsRequestRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SubjectRightsRequest, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSubjectRightsRequest() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SubjectRightsRequest), nil
}
// Builds and executes requests for operations under \privacy\subjectRightsRequests\{subjectRightsRequest-id}\microsoft.graph.getFinalAttachment()
func (m *SubjectRightsRequestRequestBuilder) GetFinalAttachment()(*ie490ccc81823dd6adcd276d7a2ff20dfdf201522ad6cac044881cf7e0b59a420.GetFinalAttachmentRequestBuilder) {
    return ie490ccc81823dd6adcd276d7a2ff20dfdf201522ad6cac044881cf7e0b59a420.NewGetFinalAttachmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \privacy\subjectRightsRequests\{subjectRightsRequest-id}\microsoft.graph.getFinalReport()
func (m *SubjectRightsRequestRequestBuilder) GetFinalReport()(*i09fc0a778050440e74342b699ac672725441e7cf37365559ae6a2eb2457f6633.GetFinalReportRequestBuilder) {
    return i09fc0a778050440e74342b699ac672725441e7cf37365559ae6a2eb2457f6633.NewGetFinalReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SubjectRightsRequestRequestBuilder) Notes()(*i793fe4cfdde8af93f5cd2898a89580574ca6d3e95e1050acfda2f70d4ef9990b.NotesRequestBuilder) {
    return i793fe4cfdde8af93f5cd2898a89580574ca6d3e95e1050acfda2f70d4ef9990b.NewNotesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.privacy.subjectRightsRequests.item.notes.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SubjectRightsRequestRequestBuilder) NotesById(id string)(*i01c49ff8ee7f291e6ef03661dd6de1defbca61ef00e7cc93142b70cd237aae9e.AuthoredNoteRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authoredNote_id"] = id
    }
    return i01c49ff8ee7f291e6ef03661dd6de1defbca61ef00e7cc93142b70cd237aae9e.NewAuthoredNoteRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Update the navigation property subjectRightsRequests in privacy
// Parameters:
//  - options : Options for the request
func (m *SubjectRightsRequestRequestBuilder) Patch(options *SubjectRightsRequestRequestBuilderPatchOptions)(error) {
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
func (m *SubjectRightsRequestRequestBuilder) Team()(*i746ebfc2054651d0c33c207b53a9353452215cae640e13084d03015de282cf39.TeamRequestBuilder) {
    return i746ebfc2054651d0c33c207b53a9353452215cae640e13084d03015de282cf39.NewTeamRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
