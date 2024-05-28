package privacy

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SubjectrightsrequestsItemApproversUserItemRequestBuilder provides operations to manage the approvers property of the microsoft.graph.subjectRightsRequest entity.
type SubjectrightsrequestsItemApproversUserItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SubjectrightsrequestsItemApproversUserItemRequestBuilderGetQueryParameters collection of users who can approve the request. Currently only supported for requests of type delete.
type SubjectrightsrequestsItemApproversUserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SubjectrightsrequestsItemApproversUserItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectrightsrequestsItemApproversUserItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SubjectrightsrequestsItemApproversUserItemRequestBuilderGetQueryParameters
}
// NewSubjectrightsrequestsItemApproversUserItemRequestBuilderInternal instantiates a new SubjectrightsrequestsItemApproversUserItemRequestBuilder and sets the default values.
func NewSubjectrightsrequestsItemApproversUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectrightsrequestsItemApproversUserItemRequestBuilder) {
    m := &SubjectrightsrequestsItemApproversUserItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privacy/subjectRightsRequests/{subjectRightsRequest%2Did}/approvers/{user%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewSubjectrightsrequestsItemApproversUserItemRequestBuilder instantiates a new SubjectrightsrequestsItemApproversUserItemRequestBuilder and sets the default values.
func NewSubjectrightsrequestsItemApproversUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectrightsrequestsItemApproversUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubjectrightsrequestsItemApproversUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get collection of users who can approve the request. Currently only supported for requests of type delete.
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SubjectrightsrequestsItemApproversUserItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SubjectrightsrequestsItemApproversUserItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable), nil
}
// MailboxSettings the mailboxSettings property
// returns a *SubjectrightsrequestsItemApproversItemMailboxsettingsMailboxSettingsRequestBuilder when successful
func (m *SubjectrightsrequestsItemApproversUserItemRequestBuilder) MailboxSettings()(*SubjectrightsrequestsItemApproversItemMailboxsettingsMailboxSettingsRequestBuilder) {
    return NewSubjectrightsrequestsItemApproversItemMailboxsettingsMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *SubjectrightsrequestsItemApproversItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *SubjectrightsrequestsItemApproversUserItemRequestBuilder) ServiceProvisioningErrors()(*SubjectrightsrequestsItemApproversItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewSubjectrightsrequestsItemApproversItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation collection of users who can approve the request. Currently only supported for requests of type delete.
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a *RequestInformation when successful
func (m *SubjectrightsrequestsItemApproversUserItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SubjectrightsrequestsItemApproversUserItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a *SubjectrightsrequestsItemApproversUserItemRequestBuilder when successful
func (m *SubjectrightsrequestsItemApproversUserItemRequestBuilder) WithUrl(rawUrl string)(*SubjectrightsrequestsItemApproversUserItemRequestBuilder) {
    return NewSubjectrightsrequestsItemApproversUserItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
