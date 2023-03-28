package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder provides operations to call the getOmaSettingPlainTextValue method.
type DeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal instantiates a new GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder and sets the default values.
func NewDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, secretReferenceValueId *string)(*DeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    m := &DeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}/getOmaSettingPlainTextValue(secretReferenceValueId='{secretReferenceValueId}')", pathParameters),
    }
    if secretReferenceValueId != nil {
        m.BaseRequestBuilder.PathParameters["secretReferenceValueId"] = *secretReferenceValueId
    }
    return m
}
// NewDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder instantiates a new GetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder and sets the default values.
func NewDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOmaSettingPlainTextValue
func (m *DeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration)(DeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseable), nil
}
// ToGetRequestInformation invoke function getOmaSettingPlainTextValue
func (m *DeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
