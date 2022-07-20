package iosupdatestatuses

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i4ebb722e5d66bbe5406a3728c62628da435b125cb96d6779272defd6effb950f "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/iosupdatestatuses/count"
)

// IosUpdateStatusesRequestBuilder provides operations to manage the iosUpdateStatuses property of the microsoft.graph.deviceManagement entity.
type IosUpdateStatusesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IosUpdateStatusesRequestBuilderGetQueryParameters the IOS software update installation statuses for this account.
type IosUpdateStatusesRequestBuilderGetQueryParameters struct {
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
// IosUpdateStatusesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosUpdateStatusesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IosUpdateStatusesRequestBuilderGetQueryParameters
}
// IosUpdateStatusesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosUpdateStatusesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIosUpdateStatusesRequestBuilderInternal instantiates a new IosUpdateStatusesRequestBuilder and sets the default values.
func NewIosUpdateStatusesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosUpdateStatusesRequestBuilder) {
    m := &IosUpdateStatusesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/iosUpdateStatuses{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIosUpdateStatusesRequestBuilder instantiates a new IosUpdateStatusesRequestBuilder and sets the default values.
func NewIosUpdateStatusesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosUpdateStatusesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosUpdateStatusesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *IosUpdateStatusesRequestBuilder) Count()(*i4ebb722e5d66bbe5406a3728c62628da435b125cb96d6779272defd6effb950f.CountRequestBuilder) {
    return i4ebb722e5d66bbe5406a3728c62628da435b125cb96d6779272defd6effb950f.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the IOS software update installation statuses for this account.
func (m *IosUpdateStatusesRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the IOS software update installation statuses for this account.
func (m *IosUpdateStatusesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *IosUpdateStatusesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to iosUpdateStatuses for deviceManagement
func (m *IosUpdateStatusesRequestBuilder) CreatePostRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosUpdateDeviceStatusable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to iosUpdateStatuses for deviceManagement
func (m *IosUpdateStatusesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosUpdateDeviceStatusable, requestConfiguration *IosUpdateStatusesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get the IOS software update installation statuses for this account.
func (m *IosUpdateStatusesRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosUpdateDeviceStatusCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the IOS software update installation statuses for this account.
func (m *IosUpdateStatusesRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *IosUpdateStatusesRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosUpdateDeviceStatusCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIosUpdateDeviceStatusCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosUpdateDeviceStatusCollectionResponseable), nil
}
// Post create new navigation property to iosUpdateStatuses for deviceManagement
func (m *IosUpdateStatusesRequestBuilder) Post(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosUpdateDeviceStatusable)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosUpdateDeviceStatusable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler create new navigation property to iosUpdateStatuses for deviceManagement
func (m *IosUpdateStatusesRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosUpdateDeviceStatusable, requestConfiguration *IosUpdateStatusesRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosUpdateDeviceStatusable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIosUpdateDeviceStatusFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosUpdateDeviceStatusable), nil
}
