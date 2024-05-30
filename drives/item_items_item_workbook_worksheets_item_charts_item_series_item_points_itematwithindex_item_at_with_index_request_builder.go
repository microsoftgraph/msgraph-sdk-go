package drives

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilder provides operations to call the itemAt method.
type ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, index *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}/series/{workbookChartSeries%2Did}/points/itemAt(index={index})", pathParameters),
    }
    if index != nil {
        m.BaseRequestBuilder.PathParameters["index"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*index), 10)
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Format provides operations to manage the format property of the microsoft.graph.workbookChartPoint entity.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexFormatRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilder) Format()(*ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexFormatRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexFormatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a point based on its position within the series.
// returns a WorkbookChartPointable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chartpointscollection-itemat?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartPointable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookChartPointFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartPointable), nil
}
// ToGetRequestInformation retrieve a point based on its position within the series.
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsItematwithindexItemAtWithIndexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
