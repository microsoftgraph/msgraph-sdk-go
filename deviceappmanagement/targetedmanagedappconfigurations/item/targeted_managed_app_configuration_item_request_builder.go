package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i10df9e1d49ddd8e40cfbeb83f29b10a26d0a9bd66f497b5f5a6f0cac3ad0a238 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/deploymentsummary"
    i1bc0e847e41ebdc918b382210b618f57eecf34594ab8b85ee5961792f938a28b "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assign"
    i2e2aa2658be52cafac5b62880a98fc25ab296fd6408d08702a2e240e8e1ef1af "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assignments"
    ibacc3af011a2e5c778e54915b40576462c31c68b752e932193de352b09cc3d55 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/targetapps"
    if5f3e25c7fb7036c0005255cf83f278d6dbdc54efe6806e0f73b94dcc8769f54 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/apps"
    i0c58f5c03e1e1e1b14874efa5218a60bb1595c685ea667497c77cd45c45299ab "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/apps/item"
    iae5cd0b807da9eeca622ea354242f60add20264ab977061a2f592e1e2bb4e2ef "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assignments/item"
)

// TargetedManagedAppConfigurationItemRequestBuilder provides operations to manage the targetedManagedAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
type TargetedManagedAppConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TargetedManagedAppConfigurationItemRequestBuilderDeleteOptions options for Delete
type TargetedManagedAppConfigurationItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// TargetedManagedAppConfigurationItemRequestBuilderGetOptions options for Get
type TargetedManagedAppConfigurationItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// TargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters targeted managed app configurations.
type TargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TargetedManagedAppConfigurationItemRequestBuilderPatchOptions options for Patch
type TargetedManagedAppConfigurationItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppConfigurationable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Apps the apps property
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Apps()(*if5f3e25c7fb7036c0005255cf83f278d6dbdc54efe6806e0f73b94dcc8769f54.AppsRequestBuilder) {
    return if5f3e25c7fb7036c0005255cf83f278d6dbdc54efe6806e0f73b94dcc8769f54.NewAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.targetedManagedAppConfigurations.item.apps.item collection
func (m *TargetedManagedAppConfigurationItemRequestBuilder) AppsById(id string)(*i0c58f5c03e1e1e1b14874efa5218a60bb1595c685ea667497c77cd45c45299ab.ManagedMobileAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedMobileApp%2Did"] = id
    }
    return i0c58f5c03e1e1e1b14874efa5218a60bb1595c685ea667497c77cd45c45299ab.NewManagedMobileAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Assign the assign property
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Assign()(*i1bc0e847e41ebdc918b382210b618f57eecf34594ab8b85ee5961792f938a28b.AssignRequestBuilder) {
    return i1bc0e847e41ebdc918b382210b618f57eecf34594ab8b85ee5961792f938a28b.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Assignments()(*i2e2aa2658be52cafac5b62880a98fc25ab296fd6408d08702a2e240e8e1ef1af.AssignmentsRequestBuilder) {
    return i2e2aa2658be52cafac5b62880a98fc25ab296fd6408d08702a2e240e8e1ef1af.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.targetedManagedAppConfigurations.item.assignments.item collection
func (m *TargetedManagedAppConfigurationItemRequestBuilder) AssignmentsById(id string)(*iae5cd0b807da9eeca622ea354242f60add20264ab977061a2f592e1e2bb4e2ef.TargetedManagedAppPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["targetedManagedAppPolicyAssignment%2Did"] = id
    }
    return iae5cd0b807da9eeca622ea354242f60add20264ab977061a2f592e1e2bb4e2ef.NewTargetedManagedAppPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTargetedManagedAppConfigurationItemRequestBuilderInternal instantiates a new TargetedManagedAppConfigurationItemRequestBuilder and sets the default values.
func NewTargetedManagedAppConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedManagedAppConfigurationItemRequestBuilder) {
    m := &TargetedManagedAppConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTargetedManagedAppConfigurationItemRequestBuilder instantiates a new TargetedManagedAppConfigurationItemRequestBuilder and sets the default values.
func NewTargetedManagedAppConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedManagedAppConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property targetedManagedAppConfigurations for deviceAppManagement
func (m *TargetedManagedAppConfigurationItemRequestBuilder) CreateDeleteRequestInformation(options *TargetedManagedAppConfigurationItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation targeted managed app configurations.
func (m *TargetedManagedAppConfigurationItemRequestBuilder) CreateGetRequestInformation(options *TargetedManagedAppConfigurationItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property targetedManagedAppConfigurations in deviceAppManagement
func (m *TargetedManagedAppConfigurationItemRequestBuilder) CreatePatchRequestInformation(options *TargetedManagedAppConfigurationItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property targetedManagedAppConfigurations for deviceAppManagement
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Delete(options *TargetedManagedAppConfigurationItemRequestBuilderDeleteOptions)(error) {
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
// DeploymentSummary the deploymentSummary property
func (m *TargetedManagedAppConfigurationItemRequestBuilder) DeploymentSummary()(*i10df9e1d49ddd8e40cfbeb83f29b10a26d0a9bd66f497b5f5a6f0cac3ad0a238.DeploymentSummaryRequestBuilder) {
    return i10df9e1d49ddd8e40cfbeb83f29b10a26d0a9bd66f497b5f5a6f0cac3ad0a238.NewDeploymentSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get targeted managed app configurations.
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Get(options *TargetedManagedAppConfigurationItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTargetedManagedAppConfigurationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppConfigurationable), nil
}
// Patch update the navigation property targetedManagedAppConfigurations in deviceAppManagement
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Patch(options *TargetedManagedAppConfigurationItemRequestBuilderPatchOptions)(error) {
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
func (m *TargetedManagedAppConfigurationItemRequestBuilder) TargetApps()(*ibacc3af011a2e5c778e54915b40576462c31c68b752e932193de352b09cc3d55.TargetAppsRequestBuilder) {
    return ibacc3af011a2e5c778e54915b40576462c31c68b752e932193de352b09cc3d55.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
