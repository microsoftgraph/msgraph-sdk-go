package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder provides operations to call the exportDeviceAndAppManagementData method.
type ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilderInternal instantiates a new ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder and sets the default values.
func NewItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder) {
    m := &ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/exportDeviceAndAppManagementData()", pathParameters),
    }
    return m
}
// NewItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder instantiates a new ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder and sets the default values.
func NewItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function exportDeviceAndAppManagementData
// returns a DeviceAndAppManagementDataable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceAndAppManagementDataable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceAndAppManagementDataFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceAndAppManagementDataable), nil
}
// ToGetRequestInformation invoke function exportDeviceAndAppManagementData
// returns a *RequestInformation when successful
func (m *ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder when successful
func (m *ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder) WithUrl(rawUrl string)(*ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder) {
    return NewItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
