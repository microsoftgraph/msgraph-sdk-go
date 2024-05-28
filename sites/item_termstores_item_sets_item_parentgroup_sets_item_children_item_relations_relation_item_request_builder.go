package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
type ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters to indicate which terms are related to the current term as either pinned or reused.
type ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters
}
// ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderInternal instantiates a new ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder and sets the default values.
func NewItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) {
    m := &ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/termStores/{store%2Did}/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/children/{term%2Did}/relations/{relation%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder instantiates a new ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder and sets the default values.
func NewItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property relations for sites
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder when successful
func (m *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) FromTerm()(*ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilder) {
    return NewItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsItemFromtermFromTermRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get to indicate which terms are related to the current term as either pinned or reused.
// returns a Relationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, error) {
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
// Patch update the navigation property relations in sites
// returns a Relationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) Patch(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, requestConfiguration *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, error) {
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
// returns a *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsItemSetRequestBuilder when successful
func (m *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) Set()(*ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsItemSetRequestBuilder) {
    return NewItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsItemSetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property relations for sites
// returns a *RequestInformation when successful
func (m *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property relations in sites
// returns a *RequestInformation when successful
func (m *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, requestConfiguration *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsItemTotermToTermRequestBuilder when successful
func (m *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) ToTerm()(*ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsItemTotermToTermRequestBuilder) {
    return NewItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsItemTotermToTermRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder when successful
func (m *ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) WithUrl(rawUrl string)(*ItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder) {
    return NewItemTermstoresItemSetsItemParentgroupSetsItemChildrenItemRelationsRelationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
