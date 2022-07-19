package registereddevices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i24e3e639b9f5b06148c9ec40d1d11974c2ee1dbd8182f991533e15fb20dbb9ee "github.com/microsoftgraph/msgraph-sdk-go/users/item/registereddevices/device"
    i274eb2c99a6073f368e48449fad87a320ad578ec991c48ae7df3005f922c0350 "github.com/microsoftgraph/msgraph-sdk-go/users/item/registereddevices/count"
    i3ca670a56332cf26c92419001b7da9370dce9185612ec29584dc6a5affa20aa1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/registereddevices/approleassignment"
    ib80af5593da2ad299e2c660a7b282b160f877f832195f3ae92a57cc844a6ef2f "github.com/microsoftgraph/msgraph-sdk-go/users/item/registereddevices/endpoint"
)

// RegisteredDevicesRequestBuilder provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
type RegisteredDevicesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RegisteredDevicesRequestBuilderGetQueryParameters devices that are registered for the user. Read-only. Nullable. Supports $expand.
type RegisteredDevicesRequestBuilderGetQueryParameters struct {
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
// RegisteredDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RegisteredDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RegisteredDevicesRequestBuilderGetQueryParameters
}
// AppRoleAssignment the appRoleAssignment property
func (m *RegisteredDevicesRequestBuilder) AppRoleAssignment()(*i3ca670a56332cf26c92419001b7da9370dce9185612ec29584dc6a5affa20aa1.AppRoleAssignmentRequestBuilder) {
    return i3ca670a56332cf26c92419001b7da9370dce9185612ec29584dc6a5affa20aa1.NewAppRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRegisteredDevicesRequestBuilderInternal instantiates a new RegisteredDevicesRequestBuilder and sets the default values.
func NewRegisteredDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RegisteredDevicesRequestBuilder) {
    m := &RegisteredDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/registeredDevices{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRegisteredDevicesRequestBuilder instantiates a new RegisteredDevicesRequestBuilder and sets the default values.
func NewRegisteredDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RegisteredDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRegisteredDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *RegisteredDevicesRequestBuilder) Count()(*i274eb2c99a6073f368e48449fad87a320ad578ec991c48ae7df3005f922c0350.CountRequestBuilder) {
    return i274eb2c99a6073f368e48449fad87a320ad578ec991c48ae7df3005f922c0350.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *RegisteredDevicesRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *RegisteredDevicesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RegisteredDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Device the device property
func (m *RegisteredDevicesRequestBuilder) Device()(*i24e3e639b9f5b06148c9ec40d1d11974c2ee1dbd8182f991533e15fb20dbb9ee.DeviceRequestBuilder) {
    return i24e3e639b9f5b06148c9ec40d1d11974c2ee1dbd8182f991533e15fb20dbb9ee.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Endpoint the endpoint property
func (m *RegisteredDevicesRequestBuilder) Endpoint()(*ib80af5593da2ad299e2c660a7b282b160f877f832195f3ae92a57cc844a6ef2f.EndpointRequestBuilder) {
    return ib80af5593da2ad299e2c660a7b282b160f877f832195f3ae92a57cc844a6ef2f.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *RegisteredDevicesRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *RegisteredDevicesRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *RegisteredDevicesRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
