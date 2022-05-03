package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i4099812734ce25dbdeab4692470c91900c048e8dee3becb3e5c54a0326239224 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item/devicestates"
    i51b72059dfed2f421aeb43f36dd6cb3a75014ea68e29bfe9f0763a1e28dc32bd "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item/assignments"
    i571f8457b84b2b26f1f46b07c4fa37d92628066190b56c6d7b9dc33cbb426950 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item/installsummary"
    ib4714ec089652eb60ee2444506e153b0e8d535e7937b3b83850dc25a80de4be4 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item/userstatesummary"
    idc6ae391c23f1dd68e1a8ab2a8216079a7e5b650c41a8ff3082a6f856ddacad3 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item/assign"
    i162ff3d0f43f21953d43cb94a636ae231382fdb4a6c3f80ff27c8fb165d745cb "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item/assignments/item"
    i9ce747e428e80c804ab60cf526324b5166b04a2d275fe59526d2b02eb3bc9118 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item/devicestates/item"
    if265c839ceb44280882d17ad922716116a39c1ce469abc935e86c7c536ccf9e0 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item/userstatesummary/item"
)

// ManagedEBookItemRequestBuilder provides operations to manage the managedEBooks property of the microsoft.graph.deviceAppManagement entity.
type ManagedEBookItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedEBookItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedEBookItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedEBookItemRequestBuilderGetQueryParameters the Managed eBook.
type ManagedEBookItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedEBookItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedEBookItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedEBookItemRequestBuilderGetQueryParameters
}
// ManagedEBookItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedEBookItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign the assign property
func (m *ManagedEBookItemRequestBuilder) Assign()(*idc6ae391c23f1dd68e1a8ab2a8216079a7e5b650c41a8ff3082a6f856ddacad3.AssignRequestBuilder) {
    return idc6ae391c23f1dd68e1a8ab2a8216079a7e5b650c41a8ff3082a6f856ddacad3.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *ManagedEBookItemRequestBuilder) Assignments()(*i51b72059dfed2f421aeb43f36dd6cb3a75014ea68e29bfe9f0763a1e28dc32bd.AssignmentsRequestBuilder) {
    return i51b72059dfed2f421aeb43f36dd6cb3a75014ea68e29bfe9f0763a1e28dc32bd.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedEBooks.item.assignments.item collection
func (m *ManagedEBookItemRequestBuilder) AssignmentsById(id string)(*i162ff3d0f43f21953d43cb94a636ae231382fdb4a6c3f80ff27c8fb165d745cb.ManagedEBookAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedEBookAssignment%2Did"] = id
    }
    return i162ff3d0f43f21953d43cb94a636ae231382fdb4a6c3f80ff27c8fb165d745cb.NewManagedEBookAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedEBookItemRequestBuilderInternal instantiates a new ManagedEBookItemRequestBuilder and sets the default values.
func NewManagedEBookItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedEBookItemRequestBuilder) {
    m := &ManagedEBookItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedEBooks/{managedEBook%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedEBookItemRequestBuilder instantiates a new ManagedEBookItemRequestBuilder and sets the default values.
func NewManagedEBookItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedEBookItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedEBookItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedEBooks for deviceAppManagement
func (m *ManagedEBookItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property managedEBooks for deviceAppManagement
func (m *ManagedEBookItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ManagedEBookItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the Managed eBook.
func (m *ManagedEBookItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the Managed eBook.
func (m *ManagedEBookItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ManagedEBookItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property managedEBooks in deviceAppManagement
func (m *ManagedEBookItemRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property managedEBooks in deviceAppManagement
func (m *ManagedEBookItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookable, requestConfiguration *ManagedEBookItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property managedEBooks for deviceAppManagement
func (m *ManagedEBookItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property managedEBooks for deviceAppManagement
func (m *ManagedEBookItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ManagedEBookItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceStates the deviceStates property
func (m *ManagedEBookItemRequestBuilder) DeviceStates()(*i4099812734ce25dbdeab4692470c91900c048e8dee3becb3e5c54a0326239224.DeviceStatesRequestBuilder) {
    return i4099812734ce25dbdeab4692470c91900c048e8dee3becb3e5c54a0326239224.NewDeviceStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedEBooks.item.deviceStates.item collection
func (m *ManagedEBookItemRequestBuilder) DeviceStatesById(id string)(*i9ce747e428e80c804ab60cf526324b5166b04a2d275fe59526d2b02eb3bc9118.DeviceInstallStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceInstallState%2Did"] = id
    }
    return i9ce747e428e80c804ab60cf526324b5166b04a2d275fe59526d2b02eb3bc9118.NewDeviceInstallStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the Managed eBook.
func (m *ManagedEBookItemRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the Managed eBook.
func (m *ManagedEBookItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ManagedEBookItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedEBookFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookable), nil
}
// InstallSummary the installSummary property
func (m *ManagedEBookItemRequestBuilder) InstallSummary()(*i571f8457b84b2b26f1f46b07c4fa37d92628066190b56c6d7b9dc33cbb426950.InstallSummaryRequestBuilder) {
    return i571f8457b84b2b26f1f46b07c4fa37d92628066190b56c6d7b9dc33cbb426950.NewInstallSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property managedEBooks in deviceAppManagement
func (m *ManagedEBookItemRequestBuilder) Patch(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property managedEBooks in deviceAppManagement
func (m *ManagedEBookItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookable, requestConfiguration *ManagedEBookItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// UserStateSummary the userStateSummary property
func (m *ManagedEBookItemRequestBuilder) UserStateSummary()(*ib4714ec089652eb60ee2444506e153b0e8d535e7937b3b83850dc25a80de4be4.UserStateSummaryRequestBuilder) {
    return ib4714ec089652eb60ee2444506e153b0e8d535e7937b3b83850dc25a80de4be4.NewUserStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStateSummaryById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedEBooks.item.userStateSummary.item collection
func (m *ManagedEBookItemRequestBuilder) UserStateSummaryById(id string)(*if265c839ceb44280882d17ad922716116a39c1ce469abc935e86c7c536ccf9e0.UserInstallStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userInstallStateSummary%2Did"] = id
    }
    return if265c839ceb44280882d17ad922716116a39c1ce469abc935e86c7c536ccf9e0.NewUserInstallStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
