package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i740a32cda16cd8a7f35813fa81f1d9cca14e072757e244629824f58d5deabdf5 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/assignments/item/reprocess"
    i82c767b5012f97dcb46f02ff991f0c26a1a7967f84309e34ede909edfa1546e0 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/assignments/item/target"
    i89ce7f3453a99fc00ceb72ee142034423113f7282418e19187b392e880621914 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/assignments/item/assignmentpolicy"
    i8da3a952221aa4c7bde05a62c631de52f912d9f6898d19d970a88752a1c845e5 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/assignments/item/accesspackage"
)

// AccessPackageAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.entitlementManagement entity.
type AccessPackageAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackageAssignmentItemRequestBuilderDeleteOptions options for Delete
type AccessPackageAssignmentItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AccessPackageAssignmentItemRequestBuilderGetOptions options for Get
type AccessPackageAssignmentItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *AccessPackageAssignmentItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AccessPackageAssignmentItemRequestBuilderGetQueryParameters represents the grant of an access package to a subject (user or group).
type AccessPackageAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessPackageAssignmentItemRequestBuilderPatchOptions options for Patch
type AccessPackageAssignmentItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AccessPackage the accessPackage property
func (m *AccessPackageAssignmentItemRequestBuilder) AccessPackage()(*i8da3a952221aa4c7bde05a62c631de52f912d9f6898d19d970a88752a1c845e5.AccessPackageRequestBuilder) {
    return i8da3a952221aa4c7bde05a62c631de52f912d9f6898d19d970a88752a1c845e5.NewAccessPackageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentPolicy the assignmentPolicy property
func (m *AccessPackageAssignmentItemRequestBuilder) AssignmentPolicy()(*i89ce7f3453a99fc00ceb72ee142034423113f7282418e19187b392e880621914.AssignmentPolicyRequestBuilder) {
    return i89ce7f3453a99fc00ceb72ee142034423113f7282418e19187b392e880621914.NewAssignmentPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessPackageAssignmentItemRequestBuilderInternal instantiates a new AccessPackageAssignmentItemRequestBuilder and sets the default values.
func NewAccessPackageAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageAssignmentItemRequestBuilder) {
    m := &AccessPackageAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/assignments/{accessPackageAssignment_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageAssignmentItemRequestBuilder instantiates a new AccessPackageAssignmentItemRequestBuilder and sets the default values.
func NewAccessPackageAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property assignments for identityGovernance
func (m *AccessPackageAssignmentItemRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageAssignmentItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation represents the grant of an access package to a subject (user or group).
func (m *AccessPackageAssignmentItemRequestBuilder) CreateGetRequestInformation(options *AccessPackageAssignmentItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property assignments in identityGovernance
func (m *AccessPackageAssignmentItemRequestBuilder) CreatePatchRequestInformation(options *AccessPackageAssignmentItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property assignments for identityGovernance
func (m *AccessPackageAssignmentItemRequestBuilder) Delete(options *AccessPackageAssignmentItemRequestBuilderDeleteOptions)(error) {
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
// Get represents the grant of an access package to a subject (user or group).
func (m *AccessPackageAssignmentItemRequestBuilder) Get(options *AccessPackageAssignmentItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageAssignmentFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable), nil
}
// Patch update the navigation property assignments in identityGovernance
func (m *AccessPackageAssignmentItemRequestBuilder) Patch(options *AccessPackageAssignmentItemRequestBuilderPatchOptions)(error) {
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
// Reprocess the reprocess property
func (m *AccessPackageAssignmentItemRequestBuilder) Reprocess()(*i740a32cda16cd8a7f35813fa81f1d9cca14e072757e244629824f58d5deabdf5.ReprocessRequestBuilder) {
    return i740a32cda16cd8a7f35813fa81f1d9cca14e072757e244629824f58d5deabdf5.NewReprocessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Target the target property
func (m *AccessPackageAssignmentItemRequestBuilder) Target()(*i82c767b5012f97dcb46f02ff991f0c26a1a7967f84309e34ede909edfa1546e0.TargetRequestBuilder) {
    return i82c767b5012f97dcb46f02ff991f0c26a1a7967f84309e34ede909edfa1546e0.NewTargetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
