package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder provides operations to manage the roleDefinition property of the microsoft.graph.roleAssignment entity.
type RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilderGetQueryParameters role definition this assignment is part of.
type RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilderGetQueryParameters
}
// NewRoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilderInternal instantiates a new RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder and sets the default values.
func NewRoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder) {
    m := &RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/roleDefinitions/{roleDefinition%2Did}/roleAssignments/{roleAssignment%2Did}/roleDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder instantiates a new RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder and sets the default values.
func NewRoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get role definition this assignment is part of.
// returns a RoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RoleDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRoleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RoleDefinitionable), nil
}
// ToGetRequestInformation role definition this assignment is part of.
// returns a *RequestInformation when successful
func (m *RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder when successful
func (m *RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder) WithUrl(rawUrl string)(*RoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder) {
    return NewRoledefinitionsItemRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
