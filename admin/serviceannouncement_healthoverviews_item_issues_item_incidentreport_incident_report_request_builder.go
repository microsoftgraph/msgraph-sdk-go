package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder provides operations to call the incidentReport method.
type ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilderInternal instantiates a new ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder and sets the default values.
func NewServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder) {
    m := &ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/serviceAnnouncement/healthOverviews/{serviceHealth%2Did}/issues/{serviceHealthIssue%2Did}/incidentReport()", pathParameters),
    }
    return m
}
// NewServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder instantiates a new ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder and sets the default values.
func NewServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Get provide the Post-Incident Review (PIR) document of a specified service issue for tenant.  An issue only with status of PostIncidentReviewPublished indicates that the PIR document exists for the issue. The operation returns an error if the specified issue doesn't exist for the tenant or if PIR document does not exist for the issue.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation provide the Post-Incident Review (PIR) document of a specified service issue for tenant.  An issue only with status of PostIncidentReviewPublished indicates that the PIR document exists for the issue. The operation returns an error if the specified issue doesn't exist for the tenant or if PIR document does not exist for the issue.
// returns a *RequestInformation when successful
func (m *ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder when successful
func (m *ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder) WithUrl(rawUrl string)(*ServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder) {
    return NewServiceannouncementHealthoverviewsItemIssuesItemIncidentreportIncidentReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
