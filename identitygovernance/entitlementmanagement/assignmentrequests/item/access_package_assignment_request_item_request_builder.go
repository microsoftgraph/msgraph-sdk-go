package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i6fcfba7c3cf2f80ff7eb56c464bf4c01dcb8779eda51a4d5b0aa111086c6f1d9 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/assignmentrequests/item/reprocess"
    i748dee09a8f1f14c9b723a15a9da9cdd5b2b3bca36dcdaaf8cbca1ebda06a99a "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/assignmentrequests/item/requestor"
    ib498f7ef8ff5f910a4aa1ac1da3f19d7068012349151e82a66f3b85b4ee5c843 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/assignmentrequests/item/assignment"
    if6c1379b58039723b5baf775bd869822006180e1ec4a8ed31e699e042488794e "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/assignmentrequests/item/accesspackage"
    if98f644e25be228d6e957c4164f416b338a91937249fd745a801630f4c1beb41 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/assignmentrequests/item/cancel"
)

// AccessPackageAssignmentRequestItemRequestBuilder provides operations to manage the assignmentRequests property of the microsoft.graph.entitlementManagement entity.
type AccessPackageAssignmentRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackageAssignmentRequestItemRequestBuilderDeleteOptions options for Delete
type AccessPackageAssignmentRequestItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AccessPackageAssignmentRequestItemRequestBuilderGetOptions options for Get
type AccessPackageAssignmentRequestItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *AccessPackageAssignmentRequestItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AccessPackageAssignmentRequestItemRequestBuilderGetQueryParameters represents access package assignment requests created by or on behalf of a user.
type AccessPackageAssignmentRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessPackageAssignmentRequestItemRequestBuilderPatchOptions options for Patch
type AccessPackageAssignmentRequestItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentRequestable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AccessPackage the accessPackage property
func (m *AccessPackageAssignmentRequestItemRequestBuilder) AccessPackage()(*if6c1379b58039723b5baf775bd869822006180e1ec4a8ed31e699e042488794e.AccessPackageRequestBuilder) {
    return if6c1379b58039723b5baf775bd869822006180e1ec4a8ed31e699e042488794e.NewAccessPackageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignment the assignment property
func (m *AccessPackageAssignmentRequestItemRequestBuilder) Assignment()(*ib498f7ef8ff5f910a4aa1ac1da3f19d7068012349151e82a66f3b85b4ee5c843.AssignmentRequestBuilder) {
    return ib498f7ef8ff5f910a4aa1ac1da3f19d7068012349151e82a66f3b85b4ee5c843.NewAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *AccessPackageAssignmentRequestItemRequestBuilder) Cancel()(*if98f644e25be228d6e957c4164f416b338a91937249fd745a801630f4c1beb41.CancelRequestBuilder) {
    return if98f644e25be228d6e957c4164f416b338a91937249fd745a801630f4c1beb41.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessPackageAssignmentRequestItemRequestBuilderInternal instantiates a new AccessPackageAssignmentRequestItemRequestBuilder and sets the default values.
func NewAccessPackageAssignmentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageAssignmentRequestItemRequestBuilder) {
    m := &AccessPackageAssignmentRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/assignmentRequests/{accessPackageAssignmentRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageAssignmentRequestItemRequestBuilder instantiates a new AccessPackageAssignmentRequestItemRequestBuilder and sets the default values.
func NewAccessPackageAssignmentRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageAssignmentRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageAssignmentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property assignmentRequests for identityGovernance
func (m *AccessPackageAssignmentRequestItemRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageAssignmentRequestItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation represents access package assignment requests created by or on behalf of a user.
func (m *AccessPackageAssignmentRequestItemRequestBuilder) CreateGetRequestInformation(options *AccessPackageAssignmentRequestItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property assignmentRequests in identityGovernance
func (m *AccessPackageAssignmentRequestItemRequestBuilder) CreatePatchRequestInformation(options *AccessPackageAssignmentRequestItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property assignmentRequests for identityGovernance
func (m *AccessPackageAssignmentRequestItemRequestBuilder) Delete(options *AccessPackageAssignmentRequestItemRequestBuilderDeleteOptions)(error) {
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
// Get represents access package assignment requests created by or on behalf of a user.
func (m *AccessPackageAssignmentRequestItemRequestBuilder) Get(options *AccessPackageAssignmentRequestItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentRequestable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageAssignmentRequestFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentRequestable), nil
}
// Patch update the navigation property assignmentRequests in identityGovernance
func (m *AccessPackageAssignmentRequestItemRequestBuilder) Patch(options *AccessPackageAssignmentRequestItemRequestBuilderPatchOptions)(error) {
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
func (m *AccessPackageAssignmentRequestItemRequestBuilder) Reprocess()(*i6fcfba7c3cf2f80ff7eb56c464bf4c01dcb8779eda51a4d5b0aa111086c6f1d9.ReprocessRequestBuilder) {
    return i6fcfba7c3cf2f80ff7eb56c464bf4c01dcb8779eda51a4d5b0aa111086c6f1d9.NewReprocessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Requestor the requestor property
func (m *AccessPackageAssignmentRequestItemRequestBuilder) Requestor()(*i748dee09a8f1f14c9b723a15a9da9cdd5b2b3bca36dcdaaf8cbca1ebda06a99a.RequestorRequestBuilder) {
    return i748dee09a8f1f14c9b723a15a9da9cdd5b2b3bca36dcdaaf8cbca1ebda06a99a.NewRequestorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
