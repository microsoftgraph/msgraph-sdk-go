package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder provides operations to manage the principal property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
type DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilderGetQueryParameters the principal that's getting a role assignment through the request. Supports $expand and $select nested in $expand for id only.
type DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilderGetQueryParameters
}
// NewDirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilderInternal instantiates a new DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder and sets the default values.
func NewDirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder) {
    m := &DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleAssignmentScheduleRequests/{unifiedRoleAssignmentScheduleRequest%2Did}/principal{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder instantiates a new DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder and sets the default values.
func NewDirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the principal that's getting a role assignment through the request. Supports $expand and $select nested in $expand for id only.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}
// ToGetRequestInformation the principal that's getting a role assignment through the request. Supports $expand and $select nested in $expand for id only.
// returns a *RequestInformation when successful
func (m *DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder when successful
func (m *DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder) {
    return NewDirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
