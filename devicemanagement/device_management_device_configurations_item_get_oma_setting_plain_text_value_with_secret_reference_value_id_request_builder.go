package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder provides operations to call the getOmaSettingPlainTextValue method.
type DeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal instantiates a new GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder and sets the default values.
func NewDeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, secretReferenceValueId *string)(*DeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    m := &DeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}/microsoft.graph.getOmaSettingPlainTextValue(secretReferenceValueId='{secretReferenceValueId}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if secretReferenceValueId != nil {
        urlTplParams["secretReferenceValueId"] = *secretReferenceValueId
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder instantiates a new GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder and sets the default values.
func NewDeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getOmaSettingPlainTextValue
func (m *DeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function getOmaSettingPlainTextValue
func (m *DeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseable), nil
}
