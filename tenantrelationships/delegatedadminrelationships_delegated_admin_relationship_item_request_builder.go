package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder provides operations to manage the delegatedAdminRelationships property of the microsoft.graph.tenantRelationship entity.
type DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderGetQueryParameters read the properties of a delegatedAdminRelationship object.
type DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderGetQueryParameters
}
// DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessAssignments provides operations to manage the accessAssignments property of the microsoft.graph.delegatedAdminRelationship entity.
// returns a *DelegatedadminrelationshipsItemAccessassignmentsAccessAssignmentsRequestBuilder when successful
func (m *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder) AccessAssignments()(*DelegatedadminrelationshipsItemAccessassignmentsAccessAssignmentsRequestBuilder) {
    return NewDelegatedadminrelationshipsItemAccessassignmentsAccessAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderInternal instantiates a new DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder and sets the default values.
func NewDelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder) {
    m := &DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/delegatedAdminRelationships/{delegatedAdminRelationship%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder instantiates a new DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder and sets the default values.
func NewDelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a delegatedAdminRelationship object. A relationship can only be deleted if it's in the 'created' status. 
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/delegatedadminrelationship-delete?view=graph-rest-1.0
func (m *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties of a delegatedAdminRelationship object.
// returns a DelegatedAdminRelationshipable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/delegatedadminrelationship-get?view=graph-rest-1.0
func (m *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedAdminRelationshipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipable), nil
}
// Operations provides operations to manage the operations property of the microsoft.graph.delegatedAdminRelationship entity.
// returns a *DelegatedadminrelationshipsItemOperationsRequestBuilder when successful
func (m *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder) Operations()(*DelegatedadminrelationshipsItemOperationsRequestBuilder) {
    return NewDelegatedadminrelationshipsItemOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a delegatedAdminRelationship object.  The following restrictions apply:- You can update this relationship when its status property is created.- You can update the autoExtendDuration property when status is either created or active.- You can only remove the Microsoft Entra Global Administrator role when the status property is active, which indicates a long-running operation.
// returns a DelegatedAdminRelationshipable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/delegatedadminrelationship-update?view=graph-rest-1.0
func (m *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipable, requestConfiguration *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedAdminRelationshipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipable), nil
}
// Requests provides operations to manage the requests property of the microsoft.graph.delegatedAdminRelationship entity.
// returns a *DelegatedadminrelationshipsItemRequestsRequestBuilder when successful
func (m *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder) Requests()(*DelegatedadminrelationshipsItemRequestsRequestBuilder) {
    return NewDelegatedadminrelationshipsItemRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a delegatedAdminRelationship object. A relationship can only be deleted if it's in the 'created' status. 
// returns a *RequestInformation when successful
func (m *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties of a delegatedAdminRelationship object.
// returns a *RequestInformation when successful
func (m *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a delegatedAdminRelationship object.  The following restrictions apply:- You can update this relationship when its status property is created.- You can update the autoExtendDuration property when status is either created or active.- You can only remove the Microsoft Entra Global Administrator role when the status property is active, which indicates a long-running operation.
// returns a *RequestInformation when successful
func (m *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipable, requestConfiguration *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder when successful
func (m *DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder) WithUrl(rawUrl string)(*DelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder) {
    return NewDelegatedadminrelationshipsDelegatedAdminRelationshipItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
