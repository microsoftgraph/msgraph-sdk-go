package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
type ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters to indicate which terms are related to the current term as either pinned or reused.
type ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters
}
// ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderInternal instantiates a new ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder and sets the default values.
func NewItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder) {
    m := &ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStore/sets/{set%2Did}/children/{term%2Did}/relations/{relation%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder instantiates a new ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder and sets the default values.
func NewItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property relations for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// FromTerm provides operations to manage the fromTerm property of the microsoft.graph.termStore.relation entity.
// returns a *ItemSitesItemTermstoreSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder when successful
func (m *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder) FromTerm()(*ItemSitesItemTermstoreSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder) {
    return NewItemSitesItemTermstoreSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get to indicate which terms are related to the current term as either pinned or reused.
// returns a Relationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property relations in groups
// returns a Relationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder) Patch(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, requestConfiguration *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// Set provides operations to manage the set property of the microsoft.graph.termStore.relation entity.
// returns a *ItemSitesItemTermstoreSetsItemChildrenItemRelationsItemSetRequestBuilder when successful
func (m *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder) Set()(*ItemSitesItemTermstoreSetsItemChildrenItemRelationsItemSetRequestBuilder) {
    return NewItemSitesItemTermstoreSetsItemChildrenItemRelationsItemSetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property relations for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation to indicate which terms are related to the current term as either pinned or reused.
// returns a *RequestInformation when successful
func (m *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property relations in groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, requestConfiguration *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToTerm provides operations to manage the toTerm property of the microsoft.graph.termStore.relation entity.
// returns a *ItemSitesItemTermstoreSetsItemChildrenItemRelationsItemTotermToTermRequestBuilder when successful
func (m *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder) ToTerm()(*ItemSitesItemTermstoreSetsItemChildrenItemRelationsItemTotermToTermRequestBuilder) {
    return NewItemSitesItemTermstoreSetsItemChildrenItemRelationsItemTotermToTermRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder when successful
func (m *ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder) {
    return NewItemSitesItemTermstoreSetsItemChildrenItemRelationsRelationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
