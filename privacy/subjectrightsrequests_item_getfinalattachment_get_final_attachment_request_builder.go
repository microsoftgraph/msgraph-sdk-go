package privacy

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder provides operations to call the getFinalAttachment method.
type SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilderInternal instantiates a new SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder and sets the default values.
func NewSubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder) {
    m := &SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privacy/subjectRightsRequests/{subjectRightsRequest%2Did}/getFinalAttachment()", pathParameters),
    }
    return m
}
// NewSubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder instantiates a new SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder and sets the default values.
func NewSubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the final attachment for a subject rights request. The attachment is a zip file that contains all the files that were included by the privacy administrator.
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/subjectrightsrequest-getfinalattachment?view=graph-rest-1.0
func (m *SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder) Get(ctx context.Context, requestConfiguration *SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get the final attachment for a subject rights request. The attachment is a zip file that contains all the files that were included by the privacy administrator.
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a *RequestInformation when successful
func (m *SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder when successful
func (m *SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder) WithUrl(rawUrl string)(*SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder) {
    return NewSubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
