package workbookrange

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i0010ef2cf68233909d7215b6b06bf7fcd334d647aa00f76f2777da9431ef6e43 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/merge"
    i0e39f083bcabfdaa2415b0b0ae0049b42085360c5c63353b5b2cbfc350e295f3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/unmerge"
    i1d08359c1719e2e3f0d27c9d8b0f624dcae0787fd9696ff091e139d9b83cef90 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/rowsbelow"
    i2ac6777b97449ca55e3b754fbba1fa8ac022f8056fe5953f0a1bb092dc5bb83f "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/columnsafter"
    i374c3a48e7e1022d3f55c1ce446dc3b42f5f93c5819a5d3a88ea48ec90250dc1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/columnwithcolumn"
    i3de1e53c37e476a8c467186268ff5a74f269b42cfcb2993b3bd04355f3c7ca3a "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/cellwithrowwithcolumn"
    i4f28bdb13c2fd87a1c75d1674b19b69336205b48b40a9ef98194b874484c6c8d "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/clear"
    i5295c3d4ef5ab9459939d7e800a6cfb76b1d590012cdec1a1523f6137050cf98 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/rowsabove"
    i5891dc4908554789336d71bb8a65a6b67e9b611e2332618b592f1a6feb29e398 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/columnsbefore"
    i5fe862c067b6cc2c27917673dd5fab9db7c08de2de174fc42a10ee6b2298ef61 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/intersectionwithanotherrange"
    i63cb89fa25ec69b92bc6cdd164a1e9c462bb6f08fa36c574f9328930f0d72d0e "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/boundingrectwithanotherrange"
    i6d1656b1396ecc59813b56bbe5daad6d441282e4e47d172200bc8a8f19e8e1b6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/lastcell"
    i876fdf124520867ecbee0b6620986c96df58a25a67d40f602277cb87b7d41f83 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/columnsbeforewithcount"
    i88aa7bc69dc13f22f247a1c8d93721c463b6a86440f438c818aeb856510c9cfa "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/offsetrangewithrowoffsetwithcolumnoffset"
    i942c32ac556f3909d7571488492b3874c802686325750e11c8b8eccfa7283864 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/visibleview"
    i99c8806041762bb035823314fbbceb9f25bdc714f0cd3306b9a95aceb24e107e "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/insert"
    i9bf3609161722f4644897e31b9db264e25c5ed26f65f65a2b085152d0457740f "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/delete"
    ia08cce7e1ad054d61e205c56260ef4c56783fa73c6695bdb925338eb8399a4b1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/entirerow"
    ia3470598b61abab2b75498261b6ad02dc4b604622ff8b42ad365c0436e13ea30 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/resizedrangewithdeltarowswithdeltacolumns"
    ia9c0e7344dd3703657c02127ca85fc9c0646d10b70d26d50f0ca1eaab398596f "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/rowwithrow"
    iba6c6ad0d4a446035d66aed6ad32b5c388c58ccb2769b2dc62823d6a45edd6f1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/columnsafterwithcount"
    ibe4897318ce16aa0018e1ce62a365fe45895080bf2bce955cebfdb26bcfa510e "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/lastcolumn"
    ibf792d16f01ee042af7083b1cc76530625bf1898cba49c69d9b5a48347663fa7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/entirecolumn"
    icfcd42608e5e5f366fbb3898fd9ae0095930ae84a165d052b2b7e6bd9c1fbd57 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/rowsabovewithcount"
    id8b964f5c3a9bce56843fa9b7cf3ac36f882352361759999106eac4496b35817 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/rowsbelowwithcount"
    ide4ab96c3ebd5530d22dcdb72c8c6f70098d4a7f4d6c97505409e3e77c2d1777 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/lastrow"
    idfdd10f48cce18d5abb8766176682f43cae164de5bce459bb4256daf48cae615 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/usedrange"
    ieb975fb512abf3424ad8d1953634695043f973a4b7470ad74ff9a8429749e4c2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrange/usedrangewithvaluesonly"
)

// workbookRangeRequestBuilder builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange
type WorkbookRangeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BoundingRectWithAnotherRange builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.boundingRect(anotherRange='{anotherRange}')
func (m *WorkbookRangeRequestBuilder) BoundingRectWithAnotherRange(anotherRange *string)(*i63cb89fa25ec69b92bc6cdd164a1e9c462bb6f08fa36c574f9328930f0d72d0e.BoundingRectWithAnotherRangeRequestBuilder) {
    return i63cb89fa25ec69b92bc6cdd164a1e9c462bb6f08fa36c574f9328930f0d72d0e.NewBoundingRectWithAnotherRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter, anotherRange);
}
// CellWithRowWithColumn builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.cell(row={row},column={column})
func (m *WorkbookRangeRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*i3de1e53c37e476a8c467186268ff5a74f269b42cfcb2993b3bd04355f3c7ca3a.CellWithRowWithColumnRequestBuilder) {
    return i3de1e53c37e476a8c467186268ff5a74f269b42cfcb2993b3bd04355f3c7ca3a.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorkbookRangeRequestBuilder) Clear()(*i4f28bdb13c2fd87a1c75d1674b19b69336205b48b40a9ef98194b874484c6c8d.ClearRequestBuilder) {
    return i4f28bdb13c2fd87a1c75d1674b19b69336205b48b40a9ef98194b874484c6c8d.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsAfter builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsAfter()
func (m *WorkbookRangeRequestBuilder) ColumnsAfter()(*i2ac6777b97449ca55e3b754fbba1fa8ac022f8056fe5953f0a1bb092dc5bb83f.ColumnsAfterRequestBuilder) {
    return i2ac6777b97449ca55e3b754fbba1fa8ac022f8056fe5953f0a1bb092dc5bb83f.NewColumnsAfterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsAfterWithCount builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsAfter(count={count})
func (m *WorkbookRangeRequestBuilder) ColumnsAfterWithCount(count *int32)(*iba6c6ad0d4a446035d66aed6ad32b5c388c58ccb2769b2dc62823d6a45edd6f1.ColumnsAfterWithCountRequestBuilder) {
    return iba6c6ad0d4a446035d66aed6ad32b5c388c58ccb2769b2dc62823d6a45edd6f1.NewColumnsAfterWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// ColumnsBefore builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsBefore()
func (m *WorkbookRangeRequestBuilder) ColumnsBefore()(*i5891dc4908554789336d71bb8a65a6b67e9b611e2332618b592f1a6feb29e398.ColumnsBeforeRequestBuilder) {
    return i5891dc4908554789336d71bb8a65a6b67e9b611e2332618b592f1a6feb29e398.NewColumnsBeforeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsBeforeWithCount builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsBefore(count={count})
func (m *WorkbookRangeRequestBuilder) ColumnsBeforeWithCount(count *int32)(*i876fdf124520867ecbee0b6620986c96df58a25a67d40f602277cb87b7d41f83.ColumnsBeforeWithCountRequestBuilder) {
    return i876fdf124520867ecbee0b6620986c96df58a25a67d40f602277cb87b7d41f83.NewColumnsBeforeWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// ColumnWithColumn builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.column(column={column})
func (m *WorkbookRangeRequestBuilder) ColumnWithColumn(column *int32)(*i374c3a48e7e1022d3f55c1ce446dc3b42f5f93c5819a5d3a88ea48ec90250dc1.ColumnWithColumnRequestBuilder) {
    return i374c3a48e7e1022d3f55c1ce446dc3b42f5f93c5819a5d3a88ea48ec90250dc1.NewColumnWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, column);
}
// NewWorkbookRangeRequestBuilderInternal instantiates a new WorkbookRangeRequestBuilder and sets the default values.
func NewWorkbookRangeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeRequestBuilder) {
    m := &WorkbookRangeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights/trending/{trending_id}/resource/microsoft.graph.workbookRange";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorkbookRangeRequestBuilder instantiates a new WorkbookRangeRequestBuilder and sets the default values.
func NewWorkbookRangeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookRangeRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorkbookRangeRequestBuilder) Delete()(*i9bf3609161722f4644897e31b9db264e25c5ed26f65f65a2b085152d0457740f.DeleteRequestBuilder) {
    return i9bf3609161722f4644897e31b9db264e25c5ed26f65f65a2b085152d0457740f.NewDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EntireColumn builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.entireColumn()
func (m *WorkbookRangeRequestBuilder) EntireColumn()(*ibf792d16f01ee042af7083b1cc76530625bf1898cba49c69d9b5a48347663fa7.EntireColumnRequestBuilder) {
    return ibf792d16f01ee042af7083b1cc76530625bf1898cba49c69d9b5a48347663fa7.NewEntireColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EntireRow builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.entireRow()
func (m *WorkbookRangeRequestBuilder) EntireRow()(*ia08cce7e1ad054d61e205c56260ef4c56783fa73c6695bdb925338eb8399a4b1.EntireRowRequestBuilder) {
    return ia08cce7e1ad054d61e205c56260ef4c56783fa73c6695bdb925338eb8399a4b1.NewEntireRowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) Insert()(*i99c8806041762bb035823314fbbceb9f25bdc714f0cd3306b9a95aceb24e107e.InsertRequestBuilder) {
    return i99c8806041762bb035823314fbbceb9f25bdc714f0cd3306b9a95aceb24e107e.NewInsertRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IntersectionWithAnotherRange builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.intersection(anotherRange='{anotherRange}')
func (m *WorkbookRangeRequestBuilder) IntersectionWithAnotherRange(anotherRange *string)(*i5fe862c067b6cc2c27917673dd5fab9db7c08de2de174fc42a10ee6b2298ef61.IntersectionWithAnotherRangeRequestBuilder) {
    return i5fe862c067b6cc2c27917673dd5fab9db7c08de2de174fc42a10ee6b2298ef61.NewIntersectionWithAnotherRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter, anotherRange);
}
// LastCell builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.lastCell()
func (m *WorkbookRangeRequestBuilder) LastCell()(*i6d1656b1396ecc59813b56bbe5daad6d441282e4e47d172200bc8a8f19e8e1b6.LastCellRequestBuilder) {
    return i6d1656b1396ecc59813b56bbe5daad6d441282e4e47d172200bc8a8f19e8e1b6.NewLastCellRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LastColumn builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.lastColumn()
func (m *WorkbookRangeRequestBuilder) LastColumn()(*ibe4897318ce16aa0018e1ce62a365fe45895080bf2bce955cebfdb26bcfa510e.LastColumnRequestBuilder) {
    return ibe4897318ce16aa0018e1ce62a365fe45895080bf2bce955cebfdb26bcfa510e.NewLastColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LastRow builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.lastRow()
func (m *WorkbookRangeRequestBuilder) LastRow()(*ide4ab96c3ebd5530d22dcdb72c8c6f70098d4a7f4d6c97505409e3e77c2d1777.LastRowRequestBuilder) {
    return ide4ab96c3ebd5530d22dcdb72c8c6f70098d4a7f4d6c97505409e3e77c2d1777.NewLastRowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) Merge()(*i0010ef2cf68233909d7215b6b06bf7fcd334d647aa00f76f2777da9431ef6e43.MergeRequestBuilder) {
    return i0010ef2cf68233909d7215b6b06bf7fcd334d647aa00f76f2777da9431ef6e43.NewMergeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OffsetRangeWithRowOffsetWithColumnOffset builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.offsetRange(rowOffset={rowOffset},columnOffset={columnOffset})
func (m *WorkbookRangeRequestBuilder) OffsetRangeWithRowOffsetWithColumnOffset(rowOffset *int32, columnOffset *int32)(*i88aa7bc69dc13f22f247a1c8d93721c463b6a86440f438c818aeb856510c9cfa.OffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder) {
    return i88aa7bc69dc13f22f247a1c8d93721c463b6a86440f438c818aeb856510c9cfa.NewOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilderInternal(m.pathParameters, m.requestAdapter, rowOffset, columnOffset);
}
// ResizedRangeWithDeltaRowsWithDeltaColumns builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.resizedRange(deltaRows={deltaRows},deltaColumns={deltaColumns})
func (m *WorkbookRangeRequestBuilder) ResizedRangeWithDeltaRowsWithDeltaColumns(deltaRows *int32, deltaColumns *int32)(*ia3470598b61abab2b75498261b6ad02dc4b604622ff8b42ad365c0436e13ea30.ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) {
    return ia3470598b61abab2b75498261b6ad02dc4b604622ff8b42ad365c0436e13ea30.NewResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter, deltaRows, deltaColumns);
}
// RowsAbove builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsAbove()
func (m *WorkbookRangeRequestBuilder) RowsAbove()(*i5295c3d4ef5ab9459939d7e800a6cfb76b1d590012cdec1a1523f6137050cf98.RowsAboveRequestBuilder) {
    return i5295c3d4ef5ab9459939d7e800a6cfb76b1d590012cdec1a1523f6137050cf98.NewRowsAboveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RowsAboveWithCount builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsAbove(count={count})
func (m *WorkbookRangeRequestBuilder) RowsAboveWithCount(count *int32)(*icfcd42608e5e5f366fbb3898fd9ae0095930ae84a165d052b2b7e6bd9c1fbd57.RowsAboveWithCountRequestBuilder) {
    return icfcd42608e5e5f366fbb3898fd9ae0095930ae84a165d052b2b7e6bd9c1fbd57.NewRowsAboveWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// RowsBelow builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsBelow()
func (m *WorkbookRangeRequestBuilder) RowsBelow()(*i1d08359c1719e2e3f0d27c9d8b0f624dcae0787fd9696ff091e139d9b83cef90.RowsBelowRequestBuilder) {
    return i1d08359c1719e2e3f0d27c9d8b0f624dcae0787fd9696ff091e139d9b83cef90.NewRowsBelowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RowsBelowWithCount builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsBelow(count={count})
func (m *WorkbookRangeRequestBuilder) RowsBelowWithCount(count *int32)(*id8b964f5c3a9bce56843fa9b7cf3ac36f882352361759999106eac4496b35817.RowsBelowWithCountRequestBuilder) {
    return id8b964f5c3a9bce56843fa9b7cf3ac36f882352361759999106eac4496b35817.NewRowsBelowWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// RowWithRow builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.row(row={row})
func (m *WorkbookRangeRequestBuilder) RowWithRow(row *int32)(*ia9c0e7344dd3703657c02127ca85fc9c0646d10b70d26d50f0ca1eaab398596f.RowWithRowRequestBuilder) {
    return ia9c0e7344dd3703657c02127ca85fc9c0646d10b70d26d50f0ca1eaab398596f.NewRowWithRowRequestBuilderInternal(m.pathParameters, m.requestAdapter, row);
}
func (m *WorkbookRangeRequestBuilder) Unmerge()(*i0e39f083bcabfdaa2415b0b0ae0049b42085360c5c63353b5b2cbfc350e295f3.UnmergeRequestBuilder) {
    return i0e39f083bcabfdaa2415b0b0ae0049b42085360c5c63353b5b2cbfc350e295f3.NewUnmergeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsedRange builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.usedRange()
func (m *WorkbookRangeRequestBuilder) UsedRange()(*idfdd10f48cce18d5abb8766176682f43cae164de5bce459bb4256daf48cae615.UsedRangeRequestBuilder) {
    return idfdd10f48cce18d5abb8766176682f43cae164de5bce459bb4256daf48cae615.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsedRangeWithValuesOnly builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.usedRange(valuesOnly={valuesOnly})
func (m *WorkbookRangeRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*ieb975fb512abf3424ad8d1953634695043f973a4b7470ad74ff9a8429749e4c2.UsedRangeWithValuesOnlyRequestBuilder) {
    return ieb975fb512abf3424ad8d1953634695043f973a4b7470ad74ff9a8429749e4c2.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
// VisibleView builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.visibleView()
func (m *WorkbookRangeRequestBuilder) VisibleView()(*i942c32ac556f3909d7571488492b3874c802686325750e11c8b8eccfa7283864.VisibleViewRequestBuilder) {
    return i942c32ac556f3909d7571488492b3874c802686325750e11c8b8eccfa7283864.NewVisibleViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
