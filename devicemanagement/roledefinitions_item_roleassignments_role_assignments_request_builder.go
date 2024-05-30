package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder provides operations to manage the roleAssignments property of the microsoft.graph.roleDefinition entity.
type RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilderGetQueryParameters list properties and relationships of the roleAssignment objects.
type RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilderGetQueryParameters struct {
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
// RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilderGetQueryParameters
}
// RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByRoleAssignmentId provides operations to manage the roleAssignments property of the microsoft.graph.roleDefinition entity.
// returns a *RoledefinitionsItemRoleassignmentsRoleAssignmentItemRequestBuilder when successful
func (m *RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder) ByRoleAssignmentId(roleAssignmentId string)(*RoledefinitionsItemRoleassignmentsRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if roleAssignmentId != "" {
        urlTplParams["roleAssignment%2Did"] = roleAssignmentId
    }
    return NewRoledefinitionsItemRoleassignmentsRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewRoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilderInternal instantiates a new RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder and sets the default values.
func NewRoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder) {
    m := &RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/roleDefinitions/{roleDefinition%2Did}/roleAssignments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewRoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder instantiates a new RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder and sets the default values.
func NewRoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *RoledefinitionsItemRoleassignmentsCountRequestBuilder when successful
func (m *RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder) Count()(*RoledefinitionsItemRoleassignmentsCountRequestBuilder) {
    return NewRoledefinitionsItemRoleassignmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list properties and relationships of the roleAssignment objects.
// returns a RoleAssignmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-rbac-roleassignment-list?view=graph-rest-1.0
func (m *RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RoleAssignmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRoleAssignmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RoleAssignmentCollectionResponseable), nil
}
// Post create a new roleAssignment object.
// returns a RoleAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-rbac-roleassignment-create?view=graph-rest-1.0
func (m *RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RoleAssignmentable, requestConfiguration *RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RoleAssignmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRoleAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RoleAssignmentable), nil
}
// ToGetRequestInformation list properties and relationships of the roleAssignment objects.
// returns a *RequestInformation when successful
func (m *RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new roleAssignment object.
// returns a *RequestInformation when successful
func (m *RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RoleAssignmentable, requestConfiguration *RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder when successful
func (m *RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder) WithUrl(rawUrl string)(*RoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder) {
    return NewRoledefinitionsItemRoleassignmentsRoleAssignmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
