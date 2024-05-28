package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder provides operations to manage the sectionGroups property of the microsoft.graph.sectionGroup entity.
type ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderGetQueryParameters the section groups in the section. Read-only. Nullable.
type ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderGetQueryParameters struct {
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
// ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderGetQueryParameters
}
// BySectionGroupId1 provides operations to manage the sectionGroups property of the microsoft.graph.sectionGroup entity.
// returns a *ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupItemRequestBuilder when successful
func (m *ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) BySectionGroupId1(sectionGroupId1 string)(*ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if sectionGroupId1 != "" {
        urlTplParams["sectionGroup%2Did1"] = sectionGroupId1
    }
    return NewItemOnenoteSectiongroupsItemSectiongroupsSectionGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderInternal instantiates a new ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder and sets the default values.
func NewItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) {
    m := &ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onenote/sectionGroups/{sectionGroup%2Did}/sectionGroups{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder instantiates a new ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder and sets the default values.
func NewItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemOnenoteSectiongroupsItemSectiongroupsCountRequestBuilder when successful
func (m *ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) Count()(*ItemOnenoteSectiongroupsItemSectiongroupsCountRequestBuilder) {
    return NewItemOnenoteSectiongroupsItemSectiongroupsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the section groups in the section. Read-only. Nullable.
// returns a SectionGroupCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation the section groups in the section. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder when successful
func (m *ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) WithUrl(rawUrl string)(*ItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder) {
    return NewItemOnenoteSectiongroupsItemSectiongroupsSectionGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
