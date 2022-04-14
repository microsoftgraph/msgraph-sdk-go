package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i46b288d9d8f774d7fa38d373a828c5bc7cf81c90735be3014be7f4d14860c0a6 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/externalsponsors"
    i6f5b9f8ecd653f7004e3a87b6fd9f1edd541123ee01576fc601583fd7a8ee2f8 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/internalsponsors"
    i38c48f98eb917d486190d72c8f1c4ae0a96a04a29df568065e9cdfbc7a1bcc7a "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/externalsponsors/item"
    i6a8a0317ba69a6948d7f14baa6206d53d8ce7fc2cbce939e0ac4ece8b4a838aa "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/internalsponsors/item"
)

// ConnectedOrganizationItemRequestBuilder provides operations to manage the connectedOrganizations property of the microsoft.graph.entitlementManagement entity.
type ConnectedOrganizationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ConnectedOrganizationItemRequestBuilderDeleteOptions options for Delete
type ConnectedOrganizationItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ConnectedOrganizationItemRequestBuilderGetOptions options for Get
type ConnectedOrganizationItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConnectedOrganizationItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ConnectedOrganizationItemRequestBuilderGetQueryParameters represents references to a directory or domain of another organization whose users can request access.
type ConnectedOrganizationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConnectedOrganizationItemRequestBuilderPatchOptions options for Patch
type ConnectedOrganizationItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConnectedOrganizationable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewConnectedOrganizationItemRequestBuilderInternal instantiates a new ConnectedOrganizationItemRequestBuilder and sets the default values.
func NewConnectedOrganizationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectedOrganizationItemRequestBuilder) {
    m := &ConnectedOrganizationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewConnectedOrganizationItemRequestBuilder instantiates a new ConnectedOrganizationItemRequestBuilder and sets the default values.
func NewConnectedOrganizationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectedOrganizationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectedOrganizationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property connectedOrganizations for identityGovernance
func (m *ConnectedOrganizationItemRequestBuilder) CreateDeleteRequestInformation(options *ConnectedOrganizationItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation represents references to a directory or domain of another organization whose users can request access.
func (m *ConnectedOrganizationItemRequestBuilder) CreateGetRequestInformation(options *ConnectedOrganizationItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property connectedOrganizations in identityGovernance
func (m *ConnectedOrganizationItemRequestBuilder) CreatePatchRequestInformation(options *ConnectedOrganizationItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property connectedOrganizations for identityGovernance
func (m *ConnectedOrganizationItemRequestBuilder) Delete(options *ConnectedOrganizationItemRequestBuilderDeleteOptions)(error) {
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
// ExternalSponsors the externalSponsors property
func (m *ConnectedOrganizationItemRequestBuilder) ExternalSponsors()(*i46b288d9d8f774d7fa38d373a828c5bc7cf81c90735be3014be7f4d14860c0a6.ExternalSponsorsRequestBuilder) {
    return i46b288d9d8f774d7fa38d373a828c5bc7cf81c90735be3014be7f4d14860c0a6.NewExternalSponsorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalSponsorsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.connectedOrganizations.item.externalSponsors.item collection
func (m *ConnectedOrganizationItemRequestBuilder) ExternalSponsorsById(id string)(*i38c48f98eb917d486190d72c8f1c4ae0a96a04a29df568065e9cdfbc7a1bcc7a.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i38c48f98eb917d486190d72c8f1c4ae0a96a04a29df568065e9cdfbc7a1bcc7a.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get represents references to a directory or domain of another organization whose users can request access.
func (m *ConnectedOrganizationItemRequestBuilder) Get(options *ConnectedOrganizationItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConnectedOrganizationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConnectedOrganizationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConnectedOrganizationable), nil
}
// InternalSponsors the internalSponsors property
func (m *ConnectedOrganizationItemRequestBuilder) InternalSponsors()(*i6f5b9f8ecd653f7004e3a87b6fd9f1edd541123ee01576fc601583fd7a8ee2f8.InternalSponsorsRequestBuilder) {
    return i6f5b9f8ecd653f7004e3a87b6fd9f1edd541123ee01576fc601583fd7a8ee2f8.NewInternalSponsorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InternalSponsorsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.connectedOrganizations.item.internalSponsors.item collection
func (m *ConnectedOrganizationItemRequestBuilder) InternalSponsorsById(id string)(*i6a8a0317ba69a6948d7f14baa6206d53d8ce7fc2cbce939e0ac4ece8b4a838aa.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i6a8a0317ba69a6948d7f14baa6206d53d8ce7fc2cbce939e0ac4ece8b4a838aa.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property connectedOrganizations in identityGovernance
func (m *ConnectedOrganizationItemRequestBuilder) Patch(options *ConnectedOrganizationItemRequestBuilderPatchOptions)(error) {
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
