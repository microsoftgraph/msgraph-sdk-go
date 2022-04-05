package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i09fc0a778050440e74342b699ac672725441e7cf37365559ae6a2eb2457f6633 "github.com/microsoftgraph/msgraph-sdk-go/privacy/subjectrightsrequests/item/getfinalreport"
    i746ebfc2054651d0c33c207b53a9353452215cae640e13084d03015de282cf39 "github.com/microsoftgraph/msgraph-sdk-go/privacy/subjectrightsrequests/item/team"
    i793fe4cfdde8af93f5cd2898a89580574ca6d3e95e1050acfda2f70d4ef9990b "github.com/microsoftgraph/msgraph-sdk-go/privacy/subjectrightsrequests/item/notes"
    ie490ccc81823dd6adcd276d7a2ff20dfdf201522ad6cac044881cf7e0b59a420 "github.com/microsoftgraph/msgraph-sdk-go/privacy/subjectrightsrequests/item/getfinalattachment"
    i01c49ff8ee7f291e6ef03661dd6de1defbca61ef00e7cc93142b70cd237aae9e "github.com/microsoftgraph/msgraph-sdk-go/privacy/subjectrightsrequests/item/notes/item"
)

// SubjectRightsRequestItemRequestBuilder provides operations to manage the subjectRightsRequests property of the microsoft.graph.privacy entity.
type SubjectRightsRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SubjectRightsRequestItemRequestBuilderDeleteOptions options for Delete
type SubjectRightsRequestItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// SubjectRightsRequestItemRequestBuilderGetOptions options for Get
type SubjectRightsRequestItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *SubjectRightsRequestItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
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
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewSubjectRightsRequestItemRequestBuilderInternal instantiates a new SubjectRightsRequestItemRequestBuilder and sets the default values.
func NewSubjectRightsRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectRightsRequestItemRequestBuilder) {
    m := &SubjectRightsRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privacy/subjectRightsRequests/{subjectRightsRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSubjectRightsRequestItemRequestBuilder instantiates a new SubjectRightsRequestItemRequestBuilder and sets the default values.
func NewSubjectRightsRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectRightsRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubjectRightsRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property subjectRightsRequests for privacy
func (m *SubjectRightsRequestItemRequestBuilder) CreateDeleteRequestInformation(options *SubjectRightsRequestItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get subjectRightsRequests from privacy
func (m *SubjectRightsRequestItemRequestBuilder) CreateGetRequestInformation(options *SubjectRightsRequestItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property subjectRightsRequests in privacy
func (m *SubjectRightsRequestItemRequestBuilder) CreatePatchRequestInformation(options *SubjectRightsRequestItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property subjectRightsRequests for privacy
func (m *SubjectRightsRequestItemRequestBuilder) Delete(options *SubjectRightsRequestItemRequestBuilderDeleteOptions)(error) {
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
// Get get subjectRightsRequests from privacy
func (m *SubjectRightsRequestItemRequestBuilder) Get(options *SubjectRightsRequestItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSubjectRightsRequestFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable), nil
}
// GetFinalAttachment provides operations to call the getFinalAttachment method.
func (m *SubjectRightsRequestItemRequestBuilder) GetFinalAttachment()(*ie490ccc81823dd6adcd276d7a2ff20dfdf201522ad6cac044881cf7e0b59a420.GetFinalAttachmentRequestBuilder) {
    return ie490ccc81823dd6adcd276d7a2ff20dfdf201522ad6cac044881cf7e0b59a420.NewGetFinalAttachmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFinalReport provides operations to call the getFinalReport method.
func (m *SubjectRightsRequestItemRequestBuilder) GetFinalReport()(*i09fc0a778050440e74342b699ac672725441e7cf37365559ae6a2eb2457f6633.GetFinalReportRequestBuilder) {
    return i09fc0a778050440e74342b699ac672725441e7cf37365559ae6a2eb2457f6633.NewGetFinalReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Notes the notes property
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
// Team the team property
func (m *SubjectRightsRequestItemRequestBuilder) Team()(*i746ebfc2054651d0c33c207b53a9353452215cae640e13084d03015de282cf39.TeamRequestBuilder) {
    return i746ebfc2054651d0c33c207b53a9353452215cae640e13084d03015de282cf39.NewTeamRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
