package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder provides operations to manage the terms property of the microsoft.graph.termStore.set entity.
type ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderGetQueryParameters all the terms under the set.
type ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderGetQueryParameters
}
// ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Children provides operations to manage the children property of the microsoft.graph.termStore.term entity.
// returns a *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsItemChildrenRequestBuilder when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder) Children()(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsItemChildrenRequestBuilder) {
    return NewItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsItemChildrenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderInternal instantiates a new ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder and sets the default values.
func NewItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder) {
    m := &ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/terms/{term%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder instantiates a new ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder and sets the default values.
func NewItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property terms for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get all the terms under the set.
// returns a Termable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderGetRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable), nil
}
// Patch update the navigation property terms in groups
// returns a Termable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder) Patch(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, requestConfiguration *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderPatchRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable), nil
}
// Relations provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
// returns a *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsItemRelationsRequestBuilder when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder) Relations()(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsItemRelationsRequestBuilder) {
    return NewItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsItemRelationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Set provides operations to manage the set property of the microsoft.graph.termStore.term entity.
// returns a *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsItemSetRequestBuilder when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder) Set()(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsItemSetRequestBuilder) {
    return NewItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsItemSetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property terms for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation all the terms under the set.
// returns a *RequestInformation when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property terms in groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, requestConfiguration *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder) {
    return NewItemSitesItemTermstoreSetsItemParentgroupSetsItemTermsTermItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
