package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i1b63f23b60a34f4a2123b7bb0bda1a81794c3af78660cac406cd4a3b1073a36e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/termsandconditions/item/acceptancestatuses"
    i311c330be28d96ff5663f2910da29ba0b1a78acd2fc73a31a5135e243041cb6e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/termsandconditions/item/assignments"
    i2fc7a2614754963997f99635ae3e50b962e993d7cc3fdeb1c79dc9fa08d7674c "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/termsandconditions/item/assignments/item"
    i4153c8e1ab57c8baaa98c7558ada5d29fa34dd83a3f935d56846eea107484e94 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/termsandconditions/item/acceptancestatuses/item"
)

// TermsAndConditionsItemRequestBuilder provides operations to manage the termsAndConditions property of the microsoft.graph.deviceManagement entity.
type TermsAndConditionsItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TermsAndConditionsItemRequestBuilderDeleteOptions options for Delete
type TermsAndConditionsItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// TermsAndConditionsItemRequestBuilderGetOptions options for Get
type TermsAndConditionsItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TermsAndConditionsItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// TermsAndConditionsItemRequestBuilderGetQueryParameters the terms and conditions associated with device management of the company.
type TermsAndConditionsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TermsAndConditionsItemRequestBuilderPatchOptions options for Patch
type TermsAndConditionsItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TermsAndConditionsable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AcceptanceStatuses the acceptanceStatuses property
func (m *TermsAndConditionsItemRequestBuilder) AcceptanceStatuses()(*i1b63f23b60a34f4a2123b7bb0bda1a81794c3af78660cac406cd4a3b1073a36e.AcceptanceStatusesRequestBuilder) {
    return i1b63f23b60a34f4a2123b7bb0bda1a81794c3af78660cac406cd4a3b1073a36e.NewAcceptanceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcceptanceStatusesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.termsAndConditions.item.acceptanceStatuses.item collection
func (m *TermsAndConditionsItemRequestBuilder) AcceptanceStatusesById(id string)(*i4153c8e1ab57c8baaa98c7558ada5d29fa34dd83a3f935d56846eea107484e94.TermsAndConditionsAcceptanceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditionsAcceptanceStatus%2Did"] = id
    }
    return i4153c8e1ab57c8baaa98c7558ada5d29fa34dd83a3f935d56846eea107484e94.NewTermsAndConditionsAcceptanceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Assignments the assignments property
func (m *TermsAndConditionsItemRequestBuilder) Assignments()(*i311c330be28d96ff5663f2910da29ba0b1a78acd2fc73a31a5135e243041cb6e.AssignmentsRequestBuilder) {
    return i311c330be28d96ff5663f2910da29ba0b1a78acd2fc73a31a5135e243041cb6e.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.termsAndConditions.item.assignments.item collection
func (m *TermsAndConditionsItemRequestBuilder) AssignmentsById(id string)(*i2fc7a2614754963997f99635ae3e50b962e993d7cc3fdeb1c79dc9fa08d7674c.TermsAndConditionsAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditionsAssignment%2Did"] = id
    }
    return i2fc7a2614754963997f99635ae3e50b962e993d7cc3fdeb1c79dc9fa08d7674c.NewTermsAndConditionsAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTermsAndConditionsItemRequestBuilderInternal instantiates a new TermsAndConditionsItemRequestBuilder and sets the default values.
func NewTermsAndConditionsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsAndConditionsItemRequestBuilder) {
    m := &TermsAndConditionsItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/termsAndConditions/{termsAndConditions%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTermsAndConditionsItemRequestBuilder instantiates a new TermsAndConditionsItemRequestBuilder and sets the default values.
func NewTermsAndConditionsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsAndConditionsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsAndConditionsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property termsAndConditions for deviceManagement
func (m *TermsAndConditionsItemRequestBuilder) CreateDeleteRequestInformation(options *TermsAndConditionsItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsItemRequestBuilder) CreateGetRequestInformation(options *TermsAndConditionsItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property termsAndConditions in deviceManagement
func (m *TermsAndConditionsItemRequestBuilder) CreatePatchRequestInformation(options *TermsAndConditionsItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property termsAndConditions for deviceManagement
func (m *TermsAndConditionsItemRequestBuilder) Delete(options *TermsAndConditionsItemRequestBuilderDeleteOptions)(error) {
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
// Get the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsItemRequestBuilder) Get(options *TermsAndConditionsItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TermsAndConditionsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTermsAndConditionsFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TermsAndConditionsable), nil
}
// Patch update the navigation property termsAndConditions in deviceManagement
func (m *TermsAndConditionsItemRequestBuilder) Patch(options *TermsAndConditionsItemRequestBuilderPatchOptions)(error) {
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
