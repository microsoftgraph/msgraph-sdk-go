package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappconfigurationsMobileAppConfigurationsRequestBuilder provides operations to manage the mobileAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
type MobileappconfigurationsMobileAppConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappconfigurationsMobileAppConfigurationsRequestBuilderGetQueryParameters list properties and relationships of the iosMobileAppConfiguration objects.
type MobileappconfigurationsMobileAppConfigurationsRequestBuilderGetQueryParameters struct {
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
// MobileappconfigurationsMobileAppConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappconfigurationsMobileAppConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappconfigurationsMobileAppConfigurationsRequestBuilderGetQueryParameters
}
// MobileappconfigurationsMobileAppConfigurationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappconfigurationsMobileAppConfigurationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByManagedDeviceMobileAppConfigurationId provides operations to manage the mobileAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
// returns a *MobileappconfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder when successful
func (m *MobileappconfigurationsMobileAppConfigurationsRequestBuilder) ByManagedDeviceMobileAppConfigurationId(managedDeviceMobileAppConfigurationId string)(*MobileappconfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managedDeviceMobileAppConfigurationId != "" {
        urlTplParams["managedDeviceMobileAppConfiguration%2Did"] = managedDeviceMobileAppConfigurationId
    }
    return NewMobileappconfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappconfigurationsMobileAppConfigurationsRequestBuilderInternal instantiates a new MobileappconfigurationsMobileAppConfigurationsRequestBuilder and sets the default values.
func NewMobileappconfigurationsMobileAppConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappconfigurationsMobileAppConfigurationsRequestBuilder) {
    m := &MobileappconfigurationsMobileAppConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileAppConfigurations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMobileappconfigurationsMobileAppConfigurationsRequestBuilder instantiates a new MobileappconfigurationsMobileAppConfigurationsRequestBuilder and sets the default values.
func NewMobileappconfigurationsMobileAppConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappconfigurationsMobileAppConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappconfigurationsMobileAppConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MobileappconfigurationsCountRequestBuilder when successful
func (m *MobileappconfigurationsMobileAppConfigurationsRequestBuilder) Count()(*MobileappconfigurationsCountRequestBuilder) {
    return NewMobileappconfigurationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list properties and relationships of the iosMobileAppConfiguration objects.
// returns a ManagedDeviceMobileAppConfigurationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-apps-iosmobileappconfiguration-list?view=graph-rest-1.0
func (m *MobileappconfigurationsMobileAppConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappconfigurationsMobileAppConfigurationsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedDeviceMobileAppConfigurationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationCollectionResponseable), nil
}
// Post create a new iosMobileAppConfiguration object.
// returns a ManagedDeviceMobileAppConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-apps-iosmobileappconfiguration-create?view=graph-rest-1.0
func (m *MobileappconfigurationsMobileAppConfigurationsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationable, requestConfiguration *MobileappconfigurationsMobileAppConfigurationsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedDeviceMobileAppConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationable), nil
}
// ToGetRequestInformation list properties and relationships of the iosMobileAppConfiguration objects.
// returns a *RequestInformation when successful
func (m *MobileappconfigurationsMobileAppConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappconfigurationsMobileAppConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new iosMobileAppConfiguration object.
// returns a *RequestInformation when successful
func (m *MobileappconfigurationsMobileAppConfigurationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationable, requestConfiguration *MobileappconfigurationsMobileAppConfigurationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappconfigurationsMobileAppConfigurationsRequestBuilder when successful
func (m *MobileappconfigurationsMobileAppConfigurationsRequestBuilder) WithUrl(rawUrl string)(*MobileappconfigurationsMobileAppConfigurationsRequestBuilder) {
    return NewMobileappconfigurationsMobileAppConfigurationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
