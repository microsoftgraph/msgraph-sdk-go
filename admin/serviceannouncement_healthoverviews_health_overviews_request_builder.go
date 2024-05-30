package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder provides operations to manage the healthOverviews property of the microsoft.graph.serviceAnnouncement entity.
type ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilderGetQueryParameters retrieve the serviceHealth resources from the healthOverviews navigation property. This operation provides the health report of all subscribed services for a tenant.
type ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilderGetQueryParameters struct {
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
// ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilderGetQueryParameters
}
// ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByServiceHealthId provides operations to manage the healthOverviews property of the microsoft.graph.serviceAnnouncement entity.
// returns a *ServiceannouncementHealthoverviewsServiceHealthItemRequestBuilder when successful
func (m *ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder) ByServiceHealthId(serviceHealthId string)(*ServiceannouncementHealthoverviewsServiceHealthItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if serviceHealthId != "" {
        urlTplParams["serviceHealth%2Did"] = serviceHealthId
    }
    return NewServiceannouncementHealthoverviewsServiceHealthItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewServiceannouncementHealthoverviewsHealthOverviewsRequestBuilderInternal instantiates a new ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder and sets the default values.
func NewServiceannouncementHealthoverviewsHealthOverviewsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder) {
    m := &ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/serviceAnnouncement/healthOverviews{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder instantiates a new ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder and sets the default values.
func NewServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceannouncementHealthoverviewsHealthOverviewsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ServiceannouncementHealthoverviewsCountRequestBuilder when successful
func (m *ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder) Count()(*ServiceannouncementHealthoverviewsCountRequestBuilder) {
    return NewServiceannouncementHealthoverviewsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the serviceHealth resources from the healthOverviews navigation property. This operation provides the health report of all subscribed services for a tenant.
// returns a ServiceHealthCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceannouncement-list-healthoverviews?view=graph-rest-1.0
func (m *ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceHealthCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServiceHealthCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceHealthCollectionResponseable), nil
}
// Post create new navigation property to healthOverviews for admin
// returns a ServiceHealthable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceHealthable, requestConfiguration *ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceHealthable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServiceHealthFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceHealthable), nil
}
// ToGetRequestInformation retrieve the serviceHealth resources from the healthOverviews navigation property. This operation provides the health report of all subscribed services for a tenant.
// returns a *RequestInformation when successful
func (m *ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to healthOverviews for admin
// returns a *RequestInformation when successful
func (m *ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceHealthable, requestConfiguration *ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder when successful
func (m *ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder) WithUrl(rawUrl string)(*ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder) {
    return NewServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
