package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i3683ba7627b8cfd708db5823cf7d88bd969d10b57d9160ad4fa9a4c365901758 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedapppolicies/item/windowsinformationprotection"
    ib18aa17be07a21575f3dc1461c70fa5b8c65d4d6a8fd95cc3441984879664e20 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedapppolicies/item/targetedmanagedappprotection"
    ibab033c1a8b48bed14f99450a0356d81369bc62e3c8a8198133c8b5ead2c5d4a "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedapppolicies/item/managedappprotection"
    iece8e7ddbf399451514e3763c1358d8aada93c5bf96f4c2a68c0c6db1146ba76 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedapppolicies/item/targetapps"
)

// ManagedAppPolicyItemRequestBuilder provides operations to manage the managedAppPolicies property of the microsoft.graph.deviceAppManagement entity.
type ManagedAppPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedAppPolicyItemRequestBuilderDeleteOptions options for Delete
type ManagedAppPolicyItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ManagedAppPolicyItemRequestBuilderGetOptions options for Get
type ManagedAppPolicyItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedAppPolicyItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ManagedAppPolicyItemRequestBuilderGetQueryParameters managed app policies.
type ManagedAppPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedAppPolicyItemRequestBuilderPatchOptions options for Patch
type ManagedAppPolicyItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewManagedAppPolicyItemRequestBuilderInternal instantiates a new ManagedAppPolicyItemRequestBuilder and sets the default values.
func NewManagedAppPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedAppPolicyItemRequestBuilder) {
    m := &ManagedAppPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppPolicies/{managedAppPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedAppPolicyItemRequestBuilder instantiates a new ManagedAppPolicyItemRequestBuilder and sets the default values.
func NewManagedAppPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedAppPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedAppPolicies for deviceAppManagement
func (m *ManagedAppPolicyItemRequestBuilder) CreateDeleteRequestInformation(options *ManagedAppPolicyItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation managed app policies.
func (m *ManagedAppPolicyItemRequestBuilder) CreateGetRequestInformation(options *ManagedAppPolicyItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property managedAppPolicies in deviceAppManagement
func (m *ManagedAppPolicyItemRequestBuilder) CreatePatchRequestInformation(options *ManagedAppPolicyItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property managedAppPolicies for deviceAppManagement
func (m *ManagedAppPolicyItemRequestBuilder) Delete(options *ManagedAppPolicyItemRequestBuilderDeleteOptions)(error) {
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
// Get managed app policies.
func (m *ManagedAppPolicyItemRequestBuilder) Get(options *ManagedAppPolicyItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedAppPolicyFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyable), nil
}
// ManagedAppProtection the managedAppProtection property
func (m *ManagedAppPolicyItemRequestBuilder) ManagedAppProtection()(*ibab033c1a8b48bed14f99450a0356d81369bc62e3c8a8198133c8b5ead2c5d4a.ManagedAppProtectionRequestBuilder) {
    return ibab033c1a8b48bed14f99450a0356d81369bc62e3c8a8198133c8b5ead2c5d4a.NewManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property managedAppPolicies in deviceAppManagement
func (m *ManagedAppPolicyItemRequestBuilder) Patch(options *ManagedAppPolicyItemRequestBuilderPatchOptions)(error) {
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
// TargetApps the targetApps property
func (m *ManagedAppPolicyItemRequestBuilder) TargetApps()(*iece8e7ddbf399451514e3763c1358d8aada93c5bf96f4c2a68c0c6db1146ba76.TargetAppsRequestBuilder) {
    return iece8e7ddbf399451514e3763c1358d8aada93c5bf96f4c2a68c0c6db1146ba76.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TargetedManagedAppProtection the targetedManagedAppProtection property
func (m *ManagedAppPolicyItemRequestBuilder) TargetedManagedAppProtection()(*ib18aa17be07a21575f3dc1461c70fa5b8c65d4d6a8fd95cc3441984879664e20.TargetedManagedAppProtectionRequestBuilder) {
    return ib18aa17be07a21575f3dc1461c70fa5b8c65d4d6a8fd95cc3441984879664e20.NewTargetedManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsInformationProtection the windowsInformationProtection property
func (m *ManagedAppPolicyItemRequestBuilder) WindowsInformationProtection()(*i3683ba7627b8cfd708db5823cf7d88bd969d10b57d9160ad4fa9a4c365901758.WindowsInformationProtectionRequestBuilder) {
    return i3683ba7627b8cfd708db5823cf7d88bd969d10b57d9160ad4fa9a4c365901758.NewWindowsInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
