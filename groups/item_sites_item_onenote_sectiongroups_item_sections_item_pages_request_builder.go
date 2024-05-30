package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder provides operations to manage the pages property of the microsoft.graph.onenoteSection entity.
type ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilderGetQueryParameters the collection of pages in the section.  Read-only. Nullable.
type ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilderGetQueryParameters
}
// ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByOnenotePageId provides operations to manage the pages property of the microsoft.graph.onenoteSection entity.
// returns a *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder when successful
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder) ByOnenotePageId(onenotePageId string)(*ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if onenotePageId != "" {
        urlTplParams["onenotePage%2Did"] = onenotePageId
    }
    return NewItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesOnenotePageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilderInternal instantiates a new ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder) {
    m := &ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/pages{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder instantiates a new ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesCountRequestBuilder when successful
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder) Count()(*ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesCountRequestBuilder) {
    return NewItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of pages in the section.  Read-only. Nullable.
// returns a OnenotePageCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenotePageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageCollectionResponseable), nil
}
// Post create new navigation property to pages for groups
// returns a OnenotePageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, requestConfiguration *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenotePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable), nil
}
// ToGetRequestInformation the collection of pages in the section.  Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to pages for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePageable, requestConfiguration *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder when successful
func (m *ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder) {
    return NewItemSitesItemOnenoteSectiongroupsItemSectionsItemPagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
