package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i20a49af14049cce13bbcbf77639ba345f23a979ba735144f96caf7cc46ac52d4 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatuses"
    i6bbd09e352a0e87235e9c9775845f897603974356216a5236acfbea5f16155b0 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/assignments"
    i8d9da7a408610e6b2da98617dc04ec0862fb37929855ca63ba17948a3e7c6869 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatussummary"
    ia2a6efbe6de075696db4a436ca6bc9745df0eb596dd5de9281c86bf2bebcd940 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/assign"
    id36e539f78070cadf5b4f1fcefe8eeff25266761b7e3541b45788010e5cdfb9b "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatussummary"
    iebd1b47b377fdbd0e32c261a4ebc852aebb3266e58bde42f93309b4e2edc6d16 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatuses"
    i5644f6cefb6e37ba63ec8582d3a158ec48759832b4146d5ffd2d350288b5e5ca "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatuses/item"
    i72a4318a7e8f28adce62f7520318129090631e70fa23a9371561c34fac1f0582 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatuses/item"
    i88ff03a861b0c98567d5cfa6d93d8d9d0de34bb73fdd948e55765d9858a275c3 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item/assignments/item"
)

// ManagedDeviceMobileAppConfigurationItemRequestBuilder provides operations to manage the mobileAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
type ManagedDeviceMobileAppConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteOptions options for Delete
type ManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ManagedDeviceMobileAppConfigurationItemRequestBuilderGetOptions options for Get
type ManagedDeviceMobileAppConfigurationItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *ManagedDeviceMobileAppConfigurationItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ManagedDeviceMobileAppConfigurationItemRequestBuilderGetQueryParameters the Managed Device Mobile Application Configurations.
type ManagedDeviceMobileAppConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ManagedDeviceMobileAppConfigurationItemRequestBuilderPatchOptions options for Patch
type ManagedDeviceMobileAppConfigurationItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Assign the assign property
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) Assign()(*ia2a6efbe6de075696db4a436ca6bc9745df0eb596dd5de9281c86bf2bebcd940.AssignRequestBuilder) {
    return ia2a6efbe6de075696db4a436ca6bc9745df0eb596dd5de9281c86bf2bebcd940.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) Assignments()(*i6bbd09e352a0e87235e9c9775845f897603974356216a5236acfbea5f16155b0.AssignmentsRequestBuilder) {
    return i6bbd09e352a0e87235e9c9775845f897603974356216a5236acfbea5f16155b0.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mobileAppConfigurations.item.assignments.item collection
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) AssignmentsById(id string)(*i88ff03a861b0c98567d5cfa6d93d8d9d0de34bb73fdd948e55765d9858a275c3.ManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationAssignment_id"] = id
    }
    return i88ff03a861b0c98567d5cfa6d93d8d9d0de34bb73fdd948e55765d9858a275c3.NewManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedDeviceMobileAppConfigurationItemRequestBuilderInternal instantiates a new ManagedDeviceMobileAppConfigurationItemRequestBuilder and sets the default values.
func NewManagedDeviceMobileAppConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceMobileAppConfigurationItemRequestBuilder) {
    m := &ManagedDeviceMobileAppConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDeviceMobileAppConfigurationItemRequestBuilder instantiates a new ManagedDeviceMobileAppConfigurationItemRequestBuilder and sets the default values.
func NewManagedDeviceMobileAppConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceMobileAppConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceMobileAppConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property mobileAppConfigurations for deviceAppManagement
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) CreateDeleteRequestInformation(options *ManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the Managed Device Mobile Application Configurations.
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) CreateGetRequestInformation(options *ManagedDeviceMobileAppConfigurationItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property mobileAppConfigurations in deviceAppManagement
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) CreatePatchRequestInformation(options *ManagedDeviceMobileAppConfigurationItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property mobileAppConfigurations for deviceAppManagement
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) Delete(options *ManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteOptions)(error) {
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
// DeviceStatuses the deviceStatuses property
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) DeviceStatuses()(*iebd1b47b377fdbd0e32c261a4ebc852aebb3266e58bde42f93309b4e2edc6d16.DeviceStatusesRequestBuilder) {
    return iebd1b47b377fdbd0e32c261a4ebc852aebb3266e58bde42f93309b4e2edc6d16.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mobileAppConfigurations.item.deviceStatuses.item collection
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) DeviceStatusesById(id string)(*i5644f6cefb6e37ba63ec8582d3a158ec48759832b4146d5ffd2d350288b5e5ca.ManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationDeviceStatus_id"] = id
    }
    return i5644f6cefb6e37ba63ec8582d3a158ec48759832b4146d5ffd2d350288b5e5ca.NewManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceStatusSummary the deviceStatusSummary property
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) DeviceStatusSummary()(*i8d9da7a408610e6b2da98617dc04ec0862fb37929855ca63ba17948a3e7c6869.DeviceStatusSummaryRequestBuilder) {
    return i8d9da7a408610e6b2da98617dc04ec0862fb37929855ca63ba17948a3e7c6869.NewDeviceStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the Managed Device Mobile Application Configurations.
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) Get(options *ManagedDeviceMobileAppConfigurationItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedDeviceMobileAppConfigurationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationable), nil
}
// Patch update the navigation property mobileAppConfigurations in deviceAppManagement
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) Patch(options *ManagedDeviceMobileAppConfigurationItemRequestBuilderPatchOptions)(error) {
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
// UserStatuses the userStatuses property
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) UserStatuses()(*i20a49af14049cce13bbcbf77639ba345f23a979ba735144f96caf7cc46ac52d4.UserStatusesRequestBuilder) {
    return i20a49af14049cce13bbcbf77639ba345f23a979ba735144f96caf7cc46ac52d4.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mobileAppConfigurations.item.userStatuses.item collection
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) UserStatusesById(id string)(*i72a4318a7e8f28adce62f7520318129090631e70fa23a9371561c34fac1f0582.ManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationUserStatus_id"] = id
    }
    return i72a4318a7e8f28adce62f7520318129090631e70fa23a9371561c34fac1f0582.NewManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserStatusSummary the userStatusSummary property
func (m *ManagedDeviceMobileAppConfigurationItemRequestBuilder) UserStatusSummary()(*id36e539f78070cadf5b4f1fcefe8eeff25266761b7e3541b45788010e5cdfb9b.UserStatusSummaryRequestBuilder) {
    return id36e539f78070cadf5b4f1fcefe8eeff25266761b7e3541b45788010e5cdfb9b.NewUserStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
