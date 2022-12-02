package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder provides operations to manage the managedAppRegistrations property of the microsoft.graph.deviceAppManagement entity.
type DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetQueryParameters the managed app registrations.
type DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetQueryParameters
}
// DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppliedPolicies provides operations to manage the appliedPolicies property of the microsoft.graph.managedAppRegistration entity.
func (m *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) AppliedPolicies()(*DeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesRequestBuilder) {
    return NewDeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppliedPoliciesById provides operations to manage the appliedPolicies property of the microsoft.graph.managedAppRegistration entity.
func (m *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) AppliedPoliciesById(id string)(*DeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesManagedAppPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy%2Did"] = id
    }
    return NewDeviceAppManagementManagedAppRegistrationsItemAppliedPoliciesManagedAppPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderInternal instantiates a new ManagedAppRegistrationItemRequestBuilder and sets the default values.
func NewDeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) {
    m := &DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder{
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
// NewDeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder instantiates a new ManagedAppRegistrationItemRequestBuilder and sets the default values.
func NewDeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedAppRegistrations for deviceAppManagement
func (m *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the managed app registrations.
func (m *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property managedAppRegistrations in deviceAppManagement
func (m *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable, requestConfiguration *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property managedAppRegistrations for deviceAppManagement
func (m *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the managed app registrations.
func (m *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedAppRegistrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable), nil
}
// IntendedPolicies provides operations to manage the intendedPolicies property of the microsoft.graph.managedAppRegistration entity.
func (m *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) IntendedPolicies()(*DeviceAppManagementManagedAppRegistrationsItemIntendedPoliciesRequestBuilder) {
    return NewDeviceAppManagementManagedAppRegistrationsItemIntendedPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IntendedPoliciesById provides operations to manage the intendedPolicies property of the microsoft.graph.managedAppRegistration entity.
func (m *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) IntendedPoliciesById(id string)(*DeviceAppManagementManagedAppRegistrationsItemIntendedPoliciesManagedAppPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy%2Did"] = id
    }
    return NewDeviceAppManagementManagedAppRegistrationsItemIntendedPoliciesManagedAppPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.managedAppRegistration entity.
func (m *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) Operations()(*DeviceAppManagementManagedAppRegistrationsItemOperationsRequestBuilder) {
    return NewDeviceAppManagementManagedAppRegistrationsItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.managedAppRegistration entity.
func (m *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) OperationsById(id string)(*DeviceAppManagementManagedAppRegistrationsItemOperationsManagedAppOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppOperation%2Did"] = id
    }
    return NewDeviceAppManagementManagedAppRegistrationsItemOperationsManagedAppOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property managedAppRegistrations in deviceAppManagement
func (m *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable, requestConfiguration *DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedAppRegistrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable), nil
}
