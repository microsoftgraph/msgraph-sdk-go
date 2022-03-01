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

// SubjectRightsRequestItemRequestBuilder builds and executes requests for operations under \privacy\subjectRightsRequests\{subjectRightsRequest-id}
type SubjectRightsRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SubjectRightsRequestItemRequestBuilderDeleteOptions options for Delete
type SubjectRightsRequestItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SubjectRightsRequestItemRequestBuilderGetOptions options for Get
type SubjectRightsRequestItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SubjectRightsRequestItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SubjectRightsRequestItemRequestBuilderGetQueryParameters get subjectRightsRequests from privacy
type SubjectRightsRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SubjectRightsRequestItemRequestBuilderPatchOptions options for Patch
type SubjectRightsRequestItemRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SubjectRightsRequest;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSubjectRightsRequestItemRequestBuilderInternal instantiates a new SubjectRightsRequestItemRequestBuilder and sets the default values.
func NewSubjectRightsRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SubjectRightsRequestItemRequestBuilder) {
    m := &SubjectRightsRequestItemRequestBuilder{
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
// NewSubjectRightsRequestItemRequestBuilder instantiates a new SubjectRightsRequestItemRequestBuilder and sets the default values.
func NewSubjectRightsRequestItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SubjectRightsRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubjectRightsRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property subjectRightsRequests for privacy
func (m *SubjectRightsRequestItemRequestBuilder) CreateDeleteRequestInformation(options *SubjectRightsRequestItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get subjectRightsRequests from privacy
func (m *SubjectRightsRequestItemRequestBuilder) CreateGetRequestInformation(options *SubjectRightsRequestItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property subjectRightsRequests in privacy
func (m *SubjectRightsRequestItemRequestBuilder) CreatePatchRequestInformation(options *SubjectRightsRequestItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property subjectRightsRequests for privacy
func (m *SubjectRightsRequestItemRequestBuilder) Delete(options *SubjectRightsRequestItemRequestBuilderDeleteOptions)(error) {
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
// Get get subjectRightsRequests from privacy
func (m *SubjectRightsRequestItemRequestBuilder) Get(options *SubjectRightsRequestItemRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SubjectRightsRequest, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSubjectRightsRequest() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SubjectRightsRequest), nil
}
// GetFinalAttachment builds and executes requests for operations under \privacy\subjectRightsRequests\{subjectRightsRequest-id}\microsoft.graph.getFinalAttachment()
func (m *SubjectRightsRequestItemRequestBuilder) GetFinalAttachment()(*ie490ccc81823dd6adcd276d7a2ff20dfdf201522ad6cac044881cf7e0b59a420.GetFinalAttachmentRequestBuilder) {
    return ie490ccc81823dd6adcd276d7a2ff20dfdf201522ad6cac044881cf7e0b59a420.NewGetFinalAttachmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFinalReport builds and executes requests for operations under \privacy\subjectRightsRequests\{subjectRightsRequest-id}\microsoft.graph.getFinalReport()
func (m *SubjectRightsRequestItemRequestBuilder) GetFinalReport()(*i09fc0a778050440e74342b699ac672725441e7cf37365559ae6a2eb2457f6633.GetFinalReportRequestBuilder) {
    return i09fc0a778050440e74342b699ac672725441e7cf37365559ae6a2eb2457f6633.NewGetFinalReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SubjectRightsRequestItemRequestBuilder) Notes()(*i793fe4cfdde8af93f5cd2898a89580574ca6d3e95e1050acfda2f70d4ef9990b.NotesRequestBuilder) {
    return i793fe4cfdde8af93f5cd2898a89580574ca6d3e95e1050acfda2f70d4ef9990b.NewNotesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.privacy.subjectRightsRequests.item.notes.item collection
func (m *SubjectRightsRequestItemRequestBuilder) NotesById(id string)(*i01c49ff8ee7f291e6ef03661dd6de1defbca61ef00e7cc93142b70cd237aae9e.AuthoredNoteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authoredNote_id"] = id
    }
    return i01c49ff8ee7f291e6ef03661dd6de1defbca61ef00e7cc93142b70cd237aae9e.NewAuthoredNoteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property subjectRightsRequests in privacy
func (m *SubjectRightsRequestItemRequestBuilder) Patch(options *SubjectRightsRequestItemRequestBuilderPatchOptions)(error) {
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
func (m *SubjectRightsRequestItemRequestBuilder) Team()(*i746ebfc2054651d0c33c207b53a9353452215cae640e13084d03015de282cf39.TeamRequestBuilder) {
    return i746ebfc2054651d0c33c207b53a9353452215cae640e13084d03015de282cf39.NewTeamRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
