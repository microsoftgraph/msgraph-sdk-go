package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder provides operations to manage the horizontalSections property of the microsoft.graph.canvasLayout entity.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderGetQueryParameters collection of horizontal sections on the SharePoint page.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderGetQueryParameters struct {
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
// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderGetQueryParameters
}
// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByHorizontalSectionId provides operations to manage the horizontalSections property of the microsoft.graph.canvasLayout entity.
// returns a *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) ByHorizontalSectionId(horizontalSectionId string)(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if horizontalSectionId != "" {
        urlTplParams["horizontalSection%2Did"] = horizontalSectionId
    }
    return NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderInternal instantiates a new ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder and sets the default values.
func NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) {
    m := &ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout/horizontalSections{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder instantiates a new ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder and sets the default values.
func NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsCountRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) Count()(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsCountRequestBuilder) {
    return NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of horizontal sections on the SharePoint page.
// returns a HorizontalSectionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionCollectionResponseable, error) {
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
// Post create new navigation property to horizontalSections for sites
// returns a HorizontalSectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionable, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionable, error) {
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
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to horizontalSections for sites
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionable, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) WithUrl(rawUrl string)(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder) {
    return NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
