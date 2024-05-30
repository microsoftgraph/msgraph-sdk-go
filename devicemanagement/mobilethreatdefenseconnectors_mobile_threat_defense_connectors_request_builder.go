package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder provides operations to manage the mobileThreatDefenseConnectors property of the microsoft.graph.deviceManagement entity.
type MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilderGetQueryParameters list properties and relationships of the mobileThreatDefenseConnector objects.
type MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilderGetQueryParameters struct {
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
// MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilderGetQueryParameters
}
// MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMobileThreatDefenseConnectorId provides operations to manage the mobileThreatDefenseConnectors property of the microsoft.graph.deviceManagement entity.
// returns a *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder when successful
func (m *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder) ByMobileThreatDefenseConnectorId(mobileThreatDefenseConnectorId string)(*MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mobileThreatDefenseConnectorId != "" {
        urlTplParams["mobileThreatDefenseConnector%2Did"] = mobileThreatDefenseConnectorId
    }
    return NewMobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilderInternal instantiates a new MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder and sets the default values.
func NewMobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder) {
    m := &MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/mobileThreatDefenseConnectors{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder instantiates a new MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder and sets the default values.
func NewMobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MobilethreatdefenseconnectorsCountRequestBuilder when successful
func (m *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder) Count()(*MobilethreatdefenseconnectorsCountRequestBuilder) {
    return NewMobilethreatdefenseconnectorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list properties and relationships of the mobileThreatDefenseConnector objects.
// returns a MobileThreatDefenseConnectorCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-onboarding-mobilethreatdefenseconnector-list?view=graph-rest-1.0
func (m *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder) Get(ctx context.Context, requestConfiguration *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileThreatDefenseConnectorCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileThreatDefenseConnectorCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileThreatDefenseConnectorCollectionResponseable), nil
}
// Post create a new mobileThreatDefenseConnector object.
// returns a MobileThreatDefenseConnectorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-onboarding-mobilethreatdefenseconnector-create?view=graph-rest-1.0
func (m *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileThreatDefenseConnectorable, requestConfiguration *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileThreatDefenseConnectorable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileThreatDefenseConnectorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileThreatDefenseConnectorable), nil
}
// ToGetRequestInformation list properties and relationships of the mobileThreatDefenseConnector objects.
// returns a *RequestInformation when successful
func (m *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new mobileThreatDefenseConnector object.
// returns a *RequestInformation when successful
func (m *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileThreatDefenseConnectorable, requestConfiguration *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder when successful
func (m *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder) WithUrl(rawUrl string)(*MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder) {
    return NewMobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
