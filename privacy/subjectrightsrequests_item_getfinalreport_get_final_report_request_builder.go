package privacy

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder provides operations to call the getFinalReport method.
type SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilderInternal instantiates a new SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder and sets the default values.
func NewSubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder) {
    m := &SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privacy/subjectRightsRequests/{subjectRightsRequest%2Did}/getFinalReport()", pathParameters),
    }
    return m
}
// NewSubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder instantiates a new SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder and sets the default values.
func NewSubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the final report for a subject rights request. The report is a text file that contains information about the files that were included by the privacy administrator.
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/subjectrightsrequest-getfinalreport?view=graph-rest-1.0
func (m *SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder) Get(ctx context.Context, requestConfiguration *SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get the final report for a subject rights request. The report is a text file that contains information about the files that were included by the privacy administrator.
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a *RequestInformation when successful
func (m *SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a *SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder when successful
func (m *SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder) WithUrl(rawUrl string)(*SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder) {
    return NewSubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
