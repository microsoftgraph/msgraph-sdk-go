package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualendpointDeviceimagesDeviceImagesRequestBuilder provides operations to manage the deviceImages property of the microsoft.graph.virtualEndpoint entity.
type VirtualendpointDeviceimagesDeviceImagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointDeviceimagesDeviceImagesRequestBuilderGetQueryParameters list the properties and relationships of cloudPcDeviceImage objects (operating system images) uploaded to Cloud PC.
type VirtualendpointDeviceimagesDeviceImagesRequestBuilderGetQueryParameters struct {
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
// VirtualendpointDeviceimagesDeviceImagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointDeviceimagesDeviceImagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointDeviceimagesDeviceImagesRequestBuilderGetQueryParameters
}
// VirtualendpointDeviceimagesDeviceImagesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointDeviceimagesDeviceImagesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCloudPcDeviceImageId provides operations to manage the deviceImages property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointDeviceimagesCloudPcDeviceImageItemRequestBuilder when successful
func (m *VirtualendpointDeviceimagesDeviceImagesRequestBuilder) ByCloudPcDeviceImageId(cloudPcDeviceImageId string)(*VirtualendpointDeviceimagesCloudPcDeviceImageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if cloudPcDeviceImageId != "" {
        urlTplParams["cloudPcDeviceImage%2Did"] = cloudPcDeviceImageId
    }
    return NewVirtualendpointDeviceimagesCloudPcDeviceImageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualendpointDeviceimagesDeviceImagesRequestBuilderInternal instantiates a new VirtualendpointDeviceimagesDeviceImagesRequestBuilder and sets the default values.
func NewVirtualendpointDeviceimagesDeviceImagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointDeviceimagesDeviceImagesRequestBuilder) {
    m := &VirtualendpointDeviceimagesDeviceImagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/deviceImages{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualendpointDeviceimagesDeviceImagesRequestBuilder instantiates a new VirtualendpointDeviceimagesDeviceImagesRequestBuilder and sets the default values.
func NewVirtualendpointDeviceimagesDeviceImagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointDeviceimagesDeviceImagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointDeviceimagesDeviceImagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualendpointDeviceimagesCountRequestBuilder when successful
func (m *VirtualendpointDeviceimagesDeviceImagesRequestBuilder) Count()(*VirtualendpointDeviceimagesCountRequestBuilder) {
    return NewVirtualendpointDeviceimagesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list the properties and relationships of cloudPcDeviceImage objects (operating system images) uploaded to Cloud PC.
// returns a CloudPcDeviceImageCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualendpoint-list-deviceimages?view=graph-rest-1.0
func (m *VirtualendpointDeviceimagesDeviceImagesRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointDeviceimagesDeviceImagesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcDeviceImageCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCloudPcDeviceImageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcDeviceImageCollectionResponseable), nil
}
// GetSourceImages provides operations to call the getSourceImages method.
// returns a *VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder when successful
func (m *VirtualendpointDeviceimagesDeviceImagesRequestBuilder) GetSourceImages()(*VirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilder) {
    return NewVirtualendpointDeviceimagesGetsourceimagesGetSourceImagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create a new cloudPcDeviceImage object. Upload a custom OS image that you can later provision on Cloud PCs.
// returns a CloudPcDeviceImageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualendpoint-post-deviceimages?view=graph-rest-1.0
func (m *VirtualendpointDeviceimagesDeviceImagesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcDeviceImageable, requestConfiguration *VirtualendpointDeviceimagesDeviceImagesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcDeviceImageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCloudPcDeviceImageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcDeviceImageable), nil
}
// ToGetRequestInformation list the properties and relationships of cloudPcDeviceImage objects (operating system images) uploaded to Cloud PC.
// returns a *RequestInformation when successful
func (m *VirtualendpointDeviceimagesDeviceImagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointDeviceimagesDeviceImagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new cloudPcDeviceImage object. Upload a custom OS image that you can later provision on Cloud PCs.
// returns a *RequestInformation when successful
func (m *VirtualendpointDeviceimagesDeviceImagesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcDeviceImageable, requestConfiguration *VirtualendpointDeviceimagesDeviceImagesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointDeviceimagesDeviceImagesRequestBuilder when successful
func (m *VirtualendpointDeviceimagesDeviceImagesRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointDeviceimagesDeviceImagesRequestBuilder) {
    return NewVirtualendpointDeviceimagesDeviceImagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
