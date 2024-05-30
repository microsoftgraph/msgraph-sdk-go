package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder provides operations to manage the horizontalSections property of the microsoft.graph.canvasLayout entity.
type ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderGetQueryParameters collection of horizontal sections on the SharePoint page.
type ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderGetQueryParameters
}
// ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByHorizontalSectionId provides operations to manage the horizontalSections property of the microsoft.graph.canvasLayout entity.
// returns a *ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) ByHorizontalSectionId(horizontalSectionId string)(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if horizontalSectionId != "" {
        urlTplParams["horizontalSection%2Did"] = horizontalSectionId
    }
    return NewItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderInternal instantiates a new ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) {
    m := &ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout/horizontalSections{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder instantiates a new ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsCountRequestBuilder when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) Count()(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsCountRequestBuilder) {
    return NewItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of horizontal sections on the SharePoint page.
// returns a HorizontalSectionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateHorizontalSectionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionCollectionResponseable), nil
}
// Post create new navigation property to horizontalSections for groups
// returns a HorizontalSectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionable, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateHorizontalSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionable), nil
}
// ToGetRequestInformation collection of horizontal sections on the SharePoint page.
// returns a *RequestInformation when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to horizontalSections for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionable, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) {
    return NewItemSitesItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
