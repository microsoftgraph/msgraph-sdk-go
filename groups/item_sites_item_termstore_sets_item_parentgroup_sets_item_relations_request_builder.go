package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder provides operations to manage the relations property of the microsoft.graph.termStore.set entity.
type ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderGetQueryParameters indicates which terms have been pinned or reused directly under the set.
type ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderGetQueryParameters
}
// ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByRelationId provides operations to manage the relations property of the microsoft.graph.termStore.set entity.
// returns a *ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRelationItemRequestBuilder when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) ByRelationId(relationId string)(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if relationId != "" {
        urlTplParams["relation%2Did"] = relationId
    }
    return NewItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRelationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderInternal instantiates a new ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder and sets the default values.
func NewItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) {
    m := &ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/relations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder instantiates a new ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder and sets the default values.
func NewItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsCountRequestBuilder when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) Count()(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsCountRequestBuilder) {
    return NewItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get indicates which terms have been pinned or reused directly under the set.
// returns a RelationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderGetRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.RelationCollectionResponseable, error) {
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
// Post create new navigation property to relations for groups
// returns a Relationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) Post(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, requestConfiguration *ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderPostRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, error) {
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
// ToGetRequestInformation indicates which terms have been pinned or reused directly under the set.
// returns a *RequestInformation when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to relations for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, requestConfiguration *ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) {
    return NewItemSitesItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
