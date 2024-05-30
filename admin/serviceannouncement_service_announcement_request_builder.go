package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServiceannouncementServiceAnnouncementRequestBuilder provides operations to manage the serviceAnnouncement property of the microsoft.graph.admin entity.
type ServiceannouncementServiceAnnouncementRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceannouncementServiceAnnouncementRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementServiceAnnouncementRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServiceannouncementServiceAnnouncementRequestBuilderGetQueryParameters a container for service communications resources. Read-only.
type ServiceannouncementServiceAnnouncementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServiceannouncementServiceAnnouncementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementServiceAnnouncementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceannouncementServiceAnnouncementRequestBuilderGetQueryParameters
}
// ServiceannouncementServiceAnnouncementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementServiceAnnouncementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceannouncementServiceAnnouncementRequestBuilderInternal instantiates a new ServiceannouncementServiceAnnouncementRequestBuilder and sets the default values.
func NewServiceannouncementServiceAnnouncementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementServiceAnnouncementRequestBuilder) {
    m := &ServiceannouncementServiceAnnouncementRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/serviceAnnouncement{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewServiceannouncementServiceAnnouncementRequestBuilder instantiates a new ServiceannouncementServiceAnnouncementRequestBuilder and sets the default values.
func NewServiceannouncementServiceAnnouncementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementServiceAnnouncementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceannouncementServiceAnnouncementRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property serviceAnnouncement for admin
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceannouncementServiceAnnouncementRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServiceannouncementServiceAnnouncementRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a container for service communications resources. Read-only.
// returns a ServiceAnnouncementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceannouncementServiceAnnouncementRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceannouncementServiceAnnouncementRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceAnnouncementable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServiceAnnouncementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceAnnouncementable), nil
}
// HealthOverviews provides operations to manage the healthOverviews property of the microsoft.graph.serviceAnnouncement entity.
// returns a *ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder when successful
func (m *ServiceannouncementServiceAnnouncementRequestBuilder) HealthOverviews()(*ServiceannouncementHealthoverviewsHealthOverviewsRequestBuilder) {
    return NewServiceannouncementHealthoverviewsHealthOverviewsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Issues provides operations to manage the issues property of the microsoft.graph.serviceAnnouncement entity.
// returns a *ServiceannouncementIssuesRequestBuilder when successful
func (m *ServiceannouncementServiceAnnouncementRequestBuilder) Issues()(*ServiceannouncementIssuesRequestBuilder) {
    return NewServiceannouncementIssuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.serviceAnnouncement entity.
// returns a *ServiceannouncementMessagesRequestBuilder when successful
func (m *ServiceannouncementServiceAnnouncementRequestBuilder) Messages()(*ServiceannouncementMessagesRequestBuilder) {
    return NewServiceannouncementMessagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property serviceAnnouncement in admin
// returns a ServiceAnnouncementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceannouncementServiceAnnouncementRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceAnnouncementable, requestConfiguration *ServiceannouncementServiceAnnouncementRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceAnnouncementable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServiceAnnouncementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceAnnouncementable), nil
}
// ToDeleteRequestInformation delete navigation property serviceAnnouncement for admin
// returns a *RequestInformation when successful
func (m *ServiceannouncementServiceAnnouncementRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ServiceannouncementServiceAnnouncementRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a container for service communications resources. Read-only.
// returns a *RequestInformation when successful
func (m *ServiceannouncementServiceAnnouncementRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceannouncementServiceAnnouncementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property serviceAnnouncement in admin
// returns a *RequestInformation when successful
func (m *ServiceannouncementServiceAnnouncementRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceAnnouncementable, requestConfiguration *ServiceannouncementServiceAnnouncementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ServiceannouncementServiceAnnouncementRequestBuilder when successful
func (m *ServiceannouncementServiceAnnouncementRequestBuilder) WithUrl(rawUrl string)(*ServiceannouncementServiceAnnouncementRequestBuilder) {
    return NewServiceannouncementServiceAnnouncementRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
