package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder provides operations to manage the deviceStatusSummary property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
type MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderGetQueryParameters read properties and relationships of the managedDeviceMobileAppConfigurationDeviceSummary object.
type MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderGetQueryParameters
}
// MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderInternal instantiates a new MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder and sets the default values.
func NewMobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder) {
    m := &MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration%2Did}/deviceStatusSummary{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder instantiates a new MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder and sets the default values.
func NewMobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceStatusSummary for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read properties and relationships of the managedDeviceMobileAppConfigurationDeviceSummary object.
// returns a ManagedDeviceMobileAppConfigurationDeviceSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-apps-manageddevicemobileappconfigurationdevicesummary-get?view=graph-rest-1.0
func (m *MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationDeviceSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedDeviceMobileAppConfigurationDeviceSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationDeviceSummaryable), nil
}
// Patch update the properties of a managedDeviceMobileAppConfigurationDeviceSummary object.
// returns a ManagedDeviceMobileAppConfigurationDeviceSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-apps-manageddevicemobileappconfigurationdevicesummary-update?view=graph-rest-1.0
func (m *MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationDeviceSummaryable, requestConfiguration *MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationDeviceSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedDeviceMobileAppConfigurationDeviceSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationDeviceSummaryable), nil
}
// ToDeleteRequestInformation delete navigation property deviceStatusSummary for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the managedDeviceMobileAppConfigurationDeviceSummary object.
// returns a *RequestInformation when successful
func (m *MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a managedDeviceMobileAppConfigurationDeviceSummary object.
// returns a *RequestInformation when successful
func (m *MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationDeviceSummaryable, requestConfiguration *MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder when successful
func (m *MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder) WithUrl(rawUrl string)(*MobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder) {
    return NewMobileappconfigurationsItemDevicestatussummaryDeviceStatusSummaryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
