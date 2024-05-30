package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder provides operations to call the getSourceImages method.
type VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilderGetQueryParameters get cloudPcSourceDeviceImage objects that can be uploaded and used on Cloud PCs. View a list of all the managed image resources from your Microsoft Entra subscriptions.
type VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilderGetQueryParameters
}
// NewVirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilderInternal instantiates a new VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder and sets the default values.
func NewVirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder) {
    m := &VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/deviceImages/getSourceImages(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder instantiates a new VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder and sets the default values.
func NewVirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get cloudPcSourceDeviceImage objects that can be uploaded and used on Cloud PCs. View a list of all the managed image resources from your Microsoft Entra subscriptions.
// Deprecated: This method is obsolete. Use GetAsGetSourceImagesGetResponse instead.
// returns a VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcdeviceimage-getsourceimages?view=graph-rest-1.0
func (m *VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilderGetRequestConfiguration)(VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualendpointDeviceimagesGetsourceimagesGetSourceImagesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesResponseable), nil
}
// GetAsGetSourceImagesGetResponse get cloudPcSourceDeviceImage objects that can be uploaded and used on Cloud PCs. View a list of all the managed image resources from your Microsoft Entra subscriptions.
// returns a VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcdeviceimage-getsourceimages?view=graph-rest-1.0
func (m *VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder) GetAsGetSourceImagesGetResponse(ctx context.Context, requestConfiguration *VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilderGetRequestConfiguration)(VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualendpointDeviceimagesGetsourceimagesGetSourceImagesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesGetResponseable), nil
}
// ToGetRequestInformation get cloudPcSourceDeviceImage objects that can be uploaded and used on Cloud PCs. View a list of all the managed image resources from your Microsoft Entra subscriptions.
// returns a *RequestInformation when successful
func (m *VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder when successful
func (m *VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder) {
    return NewVirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
