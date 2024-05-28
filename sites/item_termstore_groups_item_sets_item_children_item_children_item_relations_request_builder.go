package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
type ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilderGetQueryParameters to indicate which terms are related to the current term as either pinned or reused.
type ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilderGetQueryParameters struct {
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
// ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilderGetQueryParameters
}
// ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByRelationId provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
// returns a *ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder when successful
func (m *ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder) ByRelationId(relationId string)(*ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if relationId != "" {
        urlTplParams["relation%2Did"] = relationId
    }
    return NewItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilderInternal instantiates a new ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder and sets the default values.
func NewItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder) {
    m := &ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/termStore/groups/{group%2Did}/sets/{set%2Did}/children/{term%2Did}/children/{term%2Did1}/relations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder instantiates a new ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder and sets the default values.
func NewItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder when successful
func (m *ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder) Count()(*ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilder) {
    return NewItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get to indicate which terms are related to the current term as either pinned or reused.
// returns a RelationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilderGetRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.RelationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateRelationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.RelationCollectionResponseable), nil
}
// Post create new navigation property to relations for sites
// returns a Relationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder) Post(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, requestConfiguration *ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilderPostRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateRelationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable), nil
}
// ToGetRequestInformation to indicate which terms are related to the current term as either pinned or reused.
// returns a *RequestInformation when successful
func (m *ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to relations for sites
// returns a *RequestInformation when successful
func (m *ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, requestConfiguration *ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder when successful
func (m *ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder) WithUrl(rawUrl string)(*ItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder) {
    return NewItemTermstoreGroupsItemSetsItemChildrenItemChildrenItemRelationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
