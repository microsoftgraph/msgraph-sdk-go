package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder provides operations to manage the scheduledActionConfigurations property of the microsoft.graph.deviceComplianceScheduledActionForRule entity.
type DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilderGetQueryParameters list properties and relationships of the deviceComplianceActionItem objects.
type DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilderGetQueryParameters
}
// DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceComplianceActionItemId provides operations to manage the scheduledActionConfigurations property of the microsoft.graph.deviceComplianceScheduledActionForRule entity.
// returns a *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder when successful
func (m *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder) ByDeviceComplianceActionItemId(deviceComplianceActionItemId string)(*DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceComplianceActionItemId != "" {
        urlTplParams["deviceComplianceActionItem%2Did"] = deviceComplianceActionItemId
    }
    return NewDevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceComplianceActionItemItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilderInternal instantiates a new DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder) {
    m := &DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy%2Did}/scheduledActionsForRule/{deviceComplianceScheduledActionForRule%2Did}/scheduledActionConfigurations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder instantiates a new DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsCountRequestBuilder when successful
func (m *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder) Count()(*DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsCountRequestBuilder) {
    return NewDevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list properties and relationships of the deviceComplianceActionItem objects.
// returns a DeviceComplianceActionItemCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecomplianceactionitem-list?view=graph-rest-1.0
func (m *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceActionItemCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceComplianceActionItemCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceActionItemCollectionResponseable), nil
}
// Post create a new deviceComplianceActionItem object.
// returns a DeviceComplianceActionItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecomplianceactionitem-create?view=graph-rest-1.0
func (m *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceActionItemable, requestConfiguration *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceActionItemable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceComplianceActionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceActionItemable), nil
}
// ToGetRequestInformation list properties and relationships of the deviceComplianceActionItem objects.
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create a new deviceComplianceActionItem object.
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceActionItemable, requestConfiguration *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder when successful
func (m *DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder) WithUrl(rawUrl string)(*DevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder) {
    return NewDevicecompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsScheduledActionConfigurationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
