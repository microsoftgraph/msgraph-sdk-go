package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i8d46a6185f24acca885ef9cf0754b19ecec8c203eeb5e10885b1c3f23865a1d4 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item/scheduledactionsforrule/item/scheduledactionconfigurations"
    iec4d66d54f529f97b939337373c84e69dceab72c2e3a51099711260f728124dd "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item/scheduledactionsforrule/item/scheduledactionconfigurations/item"
)

// DeviceComplianceScheduledActionForRuleItemRequestBuilder provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceCompliancePolicy entity.
type DeviceComplianceScheduledActionForRuleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceComplianceScheduledActionForRuleItemRequestBuilderDeleteOptions options for Delete
type DeviceComplianceScheduledActionForRuleItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DeviceComplianceScheduledActionForRuleItemRequestBuilderGetOptions options for Get
type DeviceComplianceScheduledActionForRuleItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DeviceComplianceScheduledActionForRuleItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DeviceComplianceScheduledActionForRuleItemRequestBuilderGetQueryParameters the list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
type DeviceComplianceScheduledActionForRuleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceComplianceScheduledActionForRuleItemRequestBuilderPatchOptions options for Patch
type DeviceComplianceScheduledActionForRuleItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewDeviceComplianceScheduledActionForRuleItemRequestBuilderInternal instantiates a new DeviceComplianceScheduledActionForRuleItemRequestBuilder and sets the default values.
func NewDeviceComplianceScheduledActionForRuleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceComplianceScheduledActionForRuleItemRequestBuilder) {
    m := &DeviceComplianceScheduledActionForRuleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy_id}/scheduledActionsForRule/{deviceComplianceScheduledActionForRule_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceComplianceScheduledActionForRuleItemRequestBuilder instantiates a new DeviceComplianceScheduledActionForRuleItemRequestBuilder and sets the default values.
func NewDeviceComplianceScheduledActionForRuleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceComplianceScheduledActionForRuleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceComplianceScheduledActionForRuleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property scheduledActionsForRule for deviceManagement
func (m *DeviceComplianceScheduledActionForRuleItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceComplianceScheduledActionForRuleItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
func (m *DeviceComplianceScheduledActionForRuleItemRequestBuilder) CreateGetRequestInformation(options *DeviceComplianceScheduledActionForRuleItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property scheduledActionsForRule in deviceManagement
func (m *DeviceComplianceScheduledActionForRuleItemRequestBuilder) CreatePatchRequestInformation(options *DeviceComplianceScheduledActionForRuleItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property scheduledActionsForRule for deviceManagement
func (m *DeviceComplianceScheduledActionForRuleItemRequestBuilder) Delete(options *DeviceComplianceScheduledActionForRuleItemRequestBuilderDeleteOptions)(error) {
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
// Get the list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies.
func (m *DeviceComplianceScheduledActionForRuleItemRequestBuilder) Get(options *DeviceComplianceScheduledActionForRuleItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceComplianceScheduledActionForRuleFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable), nil
}
// Patch update the navigation property scheduledActionsForRule in deviceManagement
func (m *DeviceComplianceScheduledActionForRuleItemRequestBuilder) Patch(options *DeviceComplianceScheduledActionForRuleItemRequestBuilderPatchOptions)(error) {
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
// ScheduledActionConfigurations the scheduledActionConfigurations property
func (m *DeviceComplianceScheduledActionForRuleItemRequestBuilder) ScheduledActionConfigurations()(*i8d46a6185f24acca885ef9cf0754b19ecec8c203eeb5e10885b1c3f23865a1d4.ScheduledActionConfigurationsRequestBuilder) {
    return i8d46a6185f24acca885ef9cf0754b19ecec8c203eeb5e10885b1c3f23865a1d4.NewScheduledActionConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScheduledActionConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceCompliancePolicies.item.scheduledActionsForRule.item.scheduledActionConfigurations.item collection
func (m *DeviceComplianceScheduledActionForRuleItemRequestBuilder) ScheduledActionConfigurationsById(id string)(*iec4d66d54f529f97b939337373c84e69dceab72c2e3a51099711260f728124dd.DeviceComplianceActionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceActionItem_id"] = id
    }
    return iec4d66d54f529f97b939337373c84e69dceab72c2e3a51099711260f728124dd.NewDeviceComplianceActionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
