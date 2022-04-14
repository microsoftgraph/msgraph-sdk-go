package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i3eed464a4a4f7a056ae78f7d5084484b8e4c7071ced5cd16d8978c75a211418d "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies"
    i798793d33f3b0c4349f4a5beaee290369942f7f1d78a4207726bf30670bbf0d0 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/operations"
    ifa8f167e8cac58c1a5037bc99e8f6f25e58a3455715484a9419cebb030ee4304 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies"
    i0a22e7eeb322ca1994c33c38e16e5fe50b47a7810496c52fd2e2c94195f6c08f "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/operations/item"
    i0d4a89b01c459a241cf1b150996622b188ac6eefcbb7901958e20402e3fbda2f "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item"
    i8c014196246f79aab4e3ce9c3fa6da286aeeb8d4af36d5f9b7fe584d08030fe6 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies/item"
)

// ManagedAppRegistrationItemRequestBuilder provides operations to manage the managedAppRegistrations property of the microsoft.graph.deviceAppManagement entity.
type ManagedAppRegistrationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedAppRegistrationItemRequestBuilderDeleteOptions options for Delete
type ManagedAppRegistrationItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ManagedAppRegistrationItemRequestBuilderGetOptions options for Get
type ManagedAppRegistrationItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedAppRegistrationItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ManagedAppRegistrationItemRequestBuilderGetQueryParameters the managed app registrations.
type ManagedAppRegistrationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedAppRegistrationItemRequestBuilderPatchOptions options for Patch
type ManagedAppRegistrationItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AppliedPolicies the appliedPolicies property
func (m *ManagedAppRegistrationItemRequestBuilder) AppliedPolicies()(*ifa8f167e8cac58c1a5037bc99e8f6f25e58a3455715484a9419cebb030ee4304.AppliedPoliciesRequestBuilder) {
    return ifa8f167e8cac58c1a5037bc99e8f6f25e58a3455715484a9419cebb030ee4304.NewAppliedPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppliedPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedAppRegistrations.item.appliedPolicies.item collection
func (m *ManagedAppRegistrationItemRequestBuilder) AppliedPoliciesById(id string)(*i8c014196246f79aab4e3ce9c3fa6da286aeeb8d4af36d5f9b7fe584d08030fe6.ManagedAppPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy%2Did"] = id
    }
    return i8c014196246f79aab4e3ce9c3fa6da286aeeb8d4af36d5f9b7fe584d08030fe6.NewManagedAppPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedAppRegistrationItemRequestBuilderInternal instantiates a new ManagedAppRegistrationItemRequestBuilder and sets the default values.
func NewManagedAppRegistrationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedAppRegistrationItemRequestBuilder) {
    m := &ManagedAppRegistrationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedAppRegistrationItemRequestBuilder instantiates a new ManagedAppRegistrationItemRequestBuilder and sets the default values.
func NewManagedAppRegistrationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedAppRegistrationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppRegistrationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedAppRegistrations for deviceAppManagement
func (m *ManagedAppRegistrationItemRequestBuilder) CreateDeleteRequestInformation(options *ManagedAppRegistrationItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the managed app registrations.
func (m *ManagedAppRegistrationItemRequestBuilder) CreateGetRequestInformation(options *ManagedAppRegistrationItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property managedAppRegistrations in deviceAppManagement
func (m *ManagedAppRegistrationItemRequestBuilder) CreatePatchRequestInformation(options *ManagedAppRegistrationItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property managedAppRegistrations for deviceAppManagement
func (m *ManagedAppRegistrationItemRequestBuilder) Delete(options *ManagedAppRegistrationItemRequestBuilderDeleteOptions)(error) {
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
// Get the managed app registrations.
func (m *ManagedAppRegistrationItemRequestBuilder) Get(options *ManagedAppRegistrationItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedAppRegistrationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable), nil
}
// IntendedPolicies the intendedPolicies property
func (m *ManagedAppRegistrationItemRequestBuilder) IntendedPolicies()(*i3eed464a4a4f7a056ae78f7d5084484b8e4c7071ced5cd16d8978c75a211418d.IntendedPoliciesRequestBuilder) {
    return i3eed464a4a4f7a056ae78f7d5084484b8e4c7071ced5cd16d8978c75a211418d.NewIntendedPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IntendedPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedAppRegistrations.item.intendedPolicies.item collection
func (m *ManagedAppRegistrationItemRequestBuilder) IntendedPoliciesById(id string)(*i0d4a89b01c459a241cf1b150996622b188ac6eefcbb7901958e20402e3fbda2f.ManagedAppPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy%2Did"] = id
    }
    return i0d4a89b01c459a241cf1b150996622b188ac6eefcbb7901958e20402e3fbda2f.NewManagedAppPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *ManagedAppRegistrationItemRequestBuilder) Operations()(*i798793d33f3b0c4349f4a5beaee290369942f7f1d78a4207726bf30670bbf0d0.OperationsRequestBuilder) {
    return i798793d33f3b0c4349f4a5beaee290369942f7f1d78a4207726bf30670bbf0d0.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedAppRegistrations.item.operations.item collection
func (m *ManagedAppRegistrationItemRequestBuilder) OperationsById(id string)(*i0a22e7eeb322ca1994c33c38e16e5fe50b47a7810496c52fd2e2c94195f6c08f.ManagedAppOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppOperation%2Did"] = id
    }
    return i0a22e7eeb322ca1994c33c38e16e5fe50b47a7810496c52fd2e2c94195f6c08f.NewManagedAppOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property managedAppRegistrations in deviceAppManagement
func (m *ManagedAppRegistrationItemRequestBuilder) Patch(options *ManagedAppRegistrationItemRequestBuilderPatchOptions)(error) {
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
