package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilder provides operations to manage the collaborators property of the microsoft.graph.subjectRightsRequest entity.
type SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilderGetQueryParameters get collaborators from security
type SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilderGetQueryParameters
}
// NewSubjectRightsRequestsItemCollaboratorsUserItemRequestBuilderInternal instantiates a new UserItemRequestBuilder and sets the default values.
func NewSubjectRightsRequestsItemCollaboratorsUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilder) {
    m := &SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/subjectRightsRequests/{subjectRightsRequest%2Did}/collaborators/{user%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewSubjectRightsRequestsItemCollaboratorsUserItemRequestBuilder instantiates a new UserItemRequestBuilder and sets the default values.
func NewSubjectRightsRequestsItemCollaboratorsUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubjectRightsRequestsItemCollaboratorsUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get collaborators from security
func (m *SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilder) MailboxSettings()(*SubjectRightsRequestsItemCollaboratorsItemMailboxSettingsRequestBuilder) {
    return NewSubjectRightsRequestsItemCollaboratorsItemMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get collaborators from security
func (m *SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilder) WithUrl(rawUrl string)(*SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilder) {
    return NewSubjectRightsRequestsItemCollaboratorsUserItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
