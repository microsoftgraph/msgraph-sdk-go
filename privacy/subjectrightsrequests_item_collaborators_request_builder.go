package privacy

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SubjectrightsrequestsItemCollaboratorsRequestBuilder provides operations to manage the collaborators property of the microsoft.graph.subjectRightsRequest entity.
type SubjectrightsrequestsItemCollaboratorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SubjectrightsrequestsItemCollaboratorsRequestBuilderGetQueryParameters collection of users who can collaborate on the request.
type SubjectrightsrequestsItemCollaboratorsRequestBuilderGetQueryParameters struct {
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
// SubjectrightsrequestsItemCollaboratorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectrightsrequestsItemCollaboratorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SubjectrightsrequestsItemCollaboratorsRequestBuilderGetQueryParameters
}
// ByUserId provides operations to manage the collaborators property of the microsoft.graph.subjectRightsRequest entity.
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a *SubjectrightsrequestsItemCollaboratorsUserItemRequestBuilder when successful
func (m *SubjectrightsrequestsItemCollaboratorsRequestBuilder) ByUserId(userId string)(*SubjectrightsrequestsItemCollaboratorsUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userId != "" {
        urlTplParams["user%2Did"] = userId
    }
    return NewSubjectrightsrequestsItemCollaboratorsUserItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewSubjectrightsrequestsItemCollaboratorsRequestBuilderInternal instantiates a new SubjectrightsrequestsItemCollaboratorsRequestBuilder and sets the default values.
func NewSubjectrightsrequestsItemCollaboratorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectrightsrequestsItemCollaboratorsRequestBuilder) {
    m := &SubjectrightsrequestsItemCollaboratorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privacy/subjectRightsRequests/{subjectRightsRequest%2Did}/collaborators{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewSubjectrightsrequestsItemCollaboratorsRequestBuilder instantiates a new SubjectrightsrequestsItemCollaboratorsRequestBuilder and sets the default values.
func NewSubjectrightsrequestsItemCollaboratorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectrightsrequestsItemCollaboratorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubjectrightsrequestsItemCollaboratorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *SubjectrightsrequestsItemCollaboratorsCountRequestBuilder when successful
func (m *SubjectrightsrequestsItemCollaboratorsRequestBuilder) Count()(*SubjectrightsrequestsItemCollaboratorsCountRequestBuilder) {
    return NewSubjectrightsrequestsItemCollaboratorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of users who can collaborate on the request.
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a UserCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SubjectrightsrequestsItemCollaboratorsRequestBuilder) Get(ctx context.Context, requestConfiguration *SubjectrightsrequestsItemCollaboratorsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserCollectionResponseable), nil
}
// ToGetRequestInformation collection of users who can collaborate on the request.
// Deprecated: The subject rights request API under Privacy is deprecated and will stop working on  March 22, 2025. Please use the new API under Security. as of 2022-02/PrivacyDeprecate
// returns a *RequestInformation when successful
func (m *SubjectrightsrequestsItemCollaboratorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SubjectrightsrequestsItemCollaboratorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SubjectrightsrequestsItemCollaboratorsRequestBuilder when successful
func (m *SubjectrightsrequestsItemCollaboratorsRequestBuilder) WithUrl(rawUrl string)(*SubjectrightsrequestsItemCollaboratorsRequestBuilder) {
    return NewSubjectrightsrequestsItemCollaboratorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
