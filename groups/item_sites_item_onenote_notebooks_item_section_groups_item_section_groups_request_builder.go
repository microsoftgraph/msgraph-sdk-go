package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilder provides operations to manage the sectionGroups property of the microsoft.graph.sectionGroup entity.
type ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilderGetQueryParameters retrieve a list of section groups from the specified section group.
type ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilderGetQueryParameters
}
// BySectionGroupId1 provides operations to manage the sectionGroups property of the microsoft.graph.sectionGroup entity.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilder) BySectionGroupId1(sectionGroupId1 string)(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if sectionGroupId1 != "" {
        urlTplParams["sectionGroup%2Did1"] = sectionGroupId1
    }
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilderInternal instantiates a new SectionGroupsRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilder) {
    m := &ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}/sectionGroups{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilder instantiates a new SectionGroupsRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilder) Count()(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsCountRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of section groups from the specified section group.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/sectiongroup-list-sectiongroups?view=graph-rest-1.0
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSectionGroupCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupCollectionResponseable), nil
}
// ToGetRequestInformation retrieve a list of section groups from the specified section group.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
