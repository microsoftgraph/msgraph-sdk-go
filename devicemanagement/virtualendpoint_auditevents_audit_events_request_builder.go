package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualendpointAuditeventsAuditEventsRequestBuilder provides operations to manage the auditEvents property of the microsoft.graph.virtualEndpoint entity.
type VirtualendpointAuditeventsAuditEventsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointAuditeventsAuditEventsRequestBuilderGetQueryParameters list all the cloudPcAuditEvent objects for the tenant.
type VirtualendpointAuditeventsAuditEventsRequestBuilderGetQueryParameters struct {
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
// VirtualendpointAuditeventsAuditEventsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointAuditeventsAuditEventsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointAuditeventsAuditEventsRequestBuilderGetQueryParameters
}
// VirtualendpointAuditeventsAuditEventsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointAuditeventsAuditEventsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCloudPcAuditEventId provides operations to manage the auditEvents property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointAuditeventsCloudPcAuditEventItemRequestBuilder when successful
func (m *VirtualendpointAuditeventsAuditEventsRequestBuilder) ByCloudPcAuditEventId(cloudPcAuditEventId string)(*VirtualendpointAuditeventsCloudPcAuditEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if cloudPcAuditEventId != "" {
        urlTplParams["cloudPcAuditEvent%2Did"] = cloudPcAuditEventId
    }
    return NewVirtualendpointAuditeventsCloudPcAuditEventItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualendpointAuditeventsAuditEventsRequestBuilderInternal instantiates a new VirtualendpointAuditeventsAuditEventsRequestBuilder and sets the default values.
func NewVirtualendpointAuditeventsAuditEventsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointAuditeventsAuditEventsRequestBuilder) {
    m := &VirtualendpointAuditeventsAuditEventsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/auditEvents{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualendpointAuditeventsAuditEventsRequestBuilder instantiates a new VirtualendpointAuditeventsAuditEventsRequestBuilder and sets the default values.
func NewVirtualendpointAuditeventsAuditEventsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointAuditeventsAuditEventsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointAuditeventsAuditEventsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualendpointAuditeventsCountRequestBuilder when successful
func (m *VirtualendpointAuditeventsAuditEventsRequestBuilder) Count()(*VirtualendpointAuditeventsCountRequestBuilder) {
    return NewVirtualendpointAuditeventsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list all the cloudPcAuditEvent objects for the tenant.
// returns a CloudPcAuditEventCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualendpoint-list-auditevents?view=graph-rest-1.0
func (m *VirtualendpointAuditeventsAuditEventsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointAuditeventsAuditEventsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcAuditEventCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCloudPcAuditEventCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcAuditEventCollectionResponseable), nil
}
// GetAuditActivityTypes provides operations to call the getAuditActivityTypes method.
// returns a *VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder when successful
func (m *VirtualendpointAuditeventsAuditEventsRequestBuilder) GetAuditActivityTypes()(*VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder) {
    return NewVirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to auditEvents for deviceManagement
// returns a CloudPcAuditEventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointAuditeventsAuditEventsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcAuditEventable, requestConfiguration *VirtualendpointAuditeventsAuditEventsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcAuditEventable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCloudPcAuditEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcAuditEventable), nil
}
// ToGetRequestInformation list all the cloudPcAuditEvent objects for the tenant.
// returns a *RequestInformation when successful
func (m *VirtualendpointAuditeventsAuditEventsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointAuditeventsAuditEventsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to auditEvents for deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualendpointAuditeventsAuditEventsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CloudPcAuditEventable, requestConfiguration *VirtualendpointAuditeventsAuditEventsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointAuditeventsAuditEventsRequestBuilder when successful
func (m *VirtualendpointAuditeventsAuditEventsRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointAuditeventsAuditEventsRequestBuilder) {
    return NewVirtualendpointAuditeventsAuditEventsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
