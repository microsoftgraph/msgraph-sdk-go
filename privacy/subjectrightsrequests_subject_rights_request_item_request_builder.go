package privacy

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder provides operations to manage the subjectRightsRequests property of the microsoft.graph.privacy entity.
type SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SubjectrightsrequestsSubjectRightsRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectrightsrequestsSubjectRightsRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SubjectrightsrequestsSubjectRightsRequestItemRequestBuilderGetQueryParameters read the properties and relationships of a subjectRightsRequest object.
type SubjectrightsrequestsSubjectRightsRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SubjectrightsrequestsSubjectRightsRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectrightsrequestsSubjectRightsRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilderGetQueryParameters
}
// SubjectrightsrequestsSubjectRightsRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectrightsrequestsSubjectRightsRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Approvers provides operations to manage the approvers property of the microsoft.graph.subjectRightsRequest entity.
// returns a *SubjectrightsrequestsItemApproversRequestBuilder when successful
func (m *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) Approvers()(*SubjectrightsrequestsItemApproversRequestBuilder) {
    return NewSubjectrightsrequestsItemApproversRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Collaborators provides operations to manage the collaborators property of the microsoft.graph.subjectRightsRequest entity.
// returns a *SubjectrightsrequestsItemCollaboratorsRequestBuilder when successful
func (m *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) Collaborators()(*SubjectrightsrequestsItemCollaboratorsRequestBuilder) {
    return NewSubjectrightsrequestsItemCollaboratorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewSubjectrightsrequestsSubjectRightsRequestItemRequestBuilderInternal instantiates a new SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder and sets the default values.
func NewSubjectrightsrequestsSubjectRightsRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) {
    m := &SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privacy/subjectRightsRequests/{subjectRightsRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewSubjectrightsrequestsSubjectRightsRequestItemRequestBuilder instantiates a new SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder and sets the default values.
func NewSubjectrightsrequestsSubjectRightsRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubjectrightsrequestsSubjectRightsRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property subjectRightsRequests for privacy
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a subjectRightsRequest object.
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a SubjectRightsRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/subjectrightsrequest-get?view=graph-rest-1.0
func (m *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSubjectRightsRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable), nil
}
// GetFinalAttachment provides operations to call the getFinalAttachment method.
// returns a *SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder when successful
func (m *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) GetFinalAttachment()(*SubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilder) {
    return NewSubjectrightsrequestsItemGetfinalattachmentGetFinalAttachmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetFinalReport provides operations to call the getFinalReport method.
// returns a *SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder when successful
func (m *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) GetFinalReport()(*SubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilder) {
    return NewSubjectrightsrequestsItemGetfinalreportGetFinalReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Notes provides operations to manage the notes property of the microsoft.graph.subjectRightsRequest entity.
// returns a *SubjectrightsrequestsItemNotesRequestBuilder when successful
func (m *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) Notes()(*SubjectrightsrequestsItemNotesRequestBuilder) {
    return NewSubjectrightsrequestsItemNotesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a subjectRightsRequest object.
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a SubjectRightsRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/subjectrightsrequest-update?view=graph-rest-1.0
func (m *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable, requestConfiguration *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSubjectRightsRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable), nil
}
// Team provides operations to manage the team property of the microsoft.graph.subjectRightsRequest entity.
// returns a *SubjectrightsrequestsItemTeamRequestBuilder when successful
func (m *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) Team()(*SubjectrightsrequestsItemTeamRequestBuilder) {
    return NewSubjectrightsrequestsItemTeamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property subjectRightsRequests for privacy
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a *RequestInformation when successful
func (m *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a subjectRightsRequest object.
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a *RequestInformation when successful
func (m *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a subjectRightsRequest object.
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a *RequestInformation when successful
func (m *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SubjectRightsRequestable, requestConfiguration *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder when successful
func (m *SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) WithUrl(rawUrl string)(*SubjectrightsrequestsSubjectRightsRequestItemRequestBuilder) {
    return NewSubjectrightsrequestsSubjectRightsRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
