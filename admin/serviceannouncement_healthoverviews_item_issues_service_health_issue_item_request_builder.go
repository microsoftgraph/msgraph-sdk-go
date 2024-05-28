package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder provides operations to manage the issues property of the microsoft.graph.serviceHealth entity.
type ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderGetQueryParameters a collection of issues that happened on the service, with detailed information for each issue.
type ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderGetQueryParameters
}
// ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderInternal instantiates a new ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder and sets the default values.
func NewServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder) {
    m := &ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/serviceAnnouncement/healthOverviews/{serviceHealth%2Did}/issues/{serviceHealthIssue%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder instantiates a new ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder and sets the default values.
func NewServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property issues for admin
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of issues that happened on the service, with detailed information for each issue.
// returns a ServiceHealthIssueable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceHealthIssueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServiceHealthIssueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceHealthIssueable), nil
}
// IncidentReport provides operations to call the incidentReport method.
// returns a *ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder when successful
func (m *ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder) IncidentReport()(*ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder) {
    return NewServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property issues in admin
// returns a ServiceHealthIssueable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceHealthIssueable, requestConfiguration *ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceHealthIssueable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServiceHealthIssueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceHealthIssueable), nil
}
// ToDeleteRequestInformation delete navigation property issues for admin
// returns a *RequestInformation when successful
func (m *ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of issues that happened on the service, with detailed information for each issue.
// returns a *RequestInformation when successful
func (m *ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property issues in admin
// returns a *RequestInformation when successful
func (m *ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceHealthIssueable, requestConfiguration *ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder when successful
func (m *ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder) WithUrl(rawUrl string)(*ServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder) {
    return NewServiceannouncementHealthoverviewsItemIssuesServiceHealthIssueItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
