package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder provides operations to manage the terms property of the microsoft.graph.termStore.set entity.
type ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderGetQueryParameters all the terms under the set.
type ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderGetQueryParameters
}
// ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Children provides operations to manage the children property of the microsoft.graph.termStore.term entity.
// returns a *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsItemChildrenRequestBuilder when successful
func (m *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder) Children()(*ItemSitesItemTermstoresItemGroupsItemSetsItemTermsItemChildrenRequestBuilder) {
    return NewItemSitesItemTermstoresItemGroupsItemSetsItemTermsItemChildrenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderInternal instantiates a new ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder and sets the default values.
func NewItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder) {
    m := &ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStores/{store%2Did}/groups/{group%2Did1}/sets/{set%2Did}/terms/{term%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder instantiates a new ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder and sets the default values.
func NewItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property terms for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderGetRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, error) {
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
func (m *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder) Patch(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, requestConfiguration *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderPatchRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, error) {
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
// returns a *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsItemRelationsRequestBuilder when successful
func (m *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder) Relations()(*ItemSitesItemTermstoresItemGroupsItemSetsItemTermsItemRelationsRequestBuilder) {
    return NewItemSitesItemTermstoresItemGroupsItemSetsItemTermsItemRelationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Set provides operations to manage the set property of the microsoft.graph.termStore.term entity.
// returns a *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsItemSetRequestBuilder when successful
func (m *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder) Set()(*ItemSitesItemTermstoresItemGroupsItemSetsItemTermsItemSetRequestBuilder) {
    return NewItemSitesItemTermstoresItemGroupsItemSetsItemTermsItemSetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property terms for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, requestConfiguration *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder when successful
func (m *ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder) {
    return NewItemSitesItemTermstoresItemGroupsItemSetsItemTermsTermItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
