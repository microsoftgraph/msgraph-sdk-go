package workbookrange

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i0137e90048b8192c70d2981f12f5825552abe1d7253d2de5c5cf983392b76f2e "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/intersectionwithanotherrange"
    i02a5c56469b42bdc0ce5f430a5f0612795fe6b4fd97997dcd4b447be7b5ba861 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/rowsabovewithcount"
    i03d1153ceb9ce7e039b43442ba57e1648785a12bdd09b592ed3f4e656e311e5a "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/entirerow"
    i19f450b56b11336764c237e45c777ddc376b16b5af511cf03565ca8c35e4a75a "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/boundingrectwithanotherrange"
    i2b7ce90d80d4123dfb21e7605c2fb8994d11bd546fef034f7b458fbe88efcb27 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/columnsafter"
    i4a3b7d2a8e3f4908a7e8203dc20e24926de35a06af94c2ff174624744a5cb191 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/lastcell"
    i4c53bde08a7009a7c80c7c2b20e6f4fb19dde1d576571b2ef6de2b99e885617c "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/cellwithrowwithcolumn"
    i5f67ab708a56b6bfa64606d70ab83ed8f85a5ad38d56b02a59e4c519e8fc4be1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/offsetrangewithrowoffsetwithcolumnoffset"
    i63f4f8673d2c326dbfa2211124f841d7bb63294db97302c6b4c3cb1ab15465f1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/merge"
    i7aecccf6a4efcedf729c75ef5d6f8c686d20d07d04ec6bc6f9d4ab9f04c41fad "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/resizedrangewithdeltarowswithdeltacolumns"
    i7d62b3f6ca1d7d823fee23e1c2f2c436e7330247eff72970e4399e294d915ec7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/rowwithrow"
    i8a5311b23354ddf82bcd1b646f3eb63adfe8d773b4606708ccb7c539b3faec3c "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/lastcolumn"
    i8cda08c52c0efdece117e0fb29abf49f789fbe1d34cc095dda7cc6cb29253967 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/visibleview"
    i91ca2d416bc8c7860958a099103a9545cd2b8519b5cfac9ab760a45877ebfea1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/lastrow"
    i9622dcedf3086727437cd583adeda01e9c0afda0e71bc76178c7fc044e008f1b "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/columnsafterwithcount"
    i97f27ba112adda84a04bdf4eac84a02427339ff52514510201e60df9eebe4389 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/usedrange"
    iaa6d3acf8e841006360a3954e830cf9c5b5832d607648a97383b999ae3b3f331 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/rowsabove"
    ib0248ff240f5b5ad9f326a9b5746a4baa736d280af427f68c75e2ac2f8562659 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/rowsbelowwithcount"
    ib299e74ddaae1d25144b4eabfedc925b93bad3997b891249ba5522c4be887ea5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/delete"
    ib48c4a281d002ad30af16576e123a5717e6b413a11f4aa3480d03f4f22ca1f00 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/rowsbelow"
    ib8133d3713a6a700af7f7ac07448caddea3e14b46ef7b4119fb5db8fa90cfc42 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/usedrangewithvaluesonly"
    ibc4295fac6d454478f729b247eb1bb6c6df70cd4620de9a1d702ae8731648405 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/columnsbeforewithcount"
    ic8daa19dcb2cb23d286fa4b858f2b7483e3277177c165ac869f177dc87477318 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/columnwithcolumn"
    icaf0bc4e754c5f4f6c34786a769aa4a76b1547e8bbb75bbf625caab2fb29a1b2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/clear"
    icc8c73281fa3f1a351101ae8c02b7004fbedc8894e64f758040b955c094190a9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/insert"
    id4aa6042f1f790a2c032e89a3bf64f66ea3fe7b73ddd8907b120b374d8b9e58d "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/entirecolumn"
    ie138fb5e04a8819c265fe49eb6107d57526826809fc0ff38bc8744e46ab167a6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/columnsbefore"
    iefe5f75095bc26602696a7643255e6c4c3650219bb4d1eb8c3361ae039b10826 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrange/unmerge"
)

// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange
type WorkbookRangeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.boundingRect(anotherRange='{anotherRange}')
// Parameters:
//  - anotherRange : Usage: anotherRange={anotherRange}
func (m *WorkbookRangeRequestBuilder) BoundingRectWithAnotherRange(anotherRange *string)(*i19f450b56b11336764c237e45c777ddc376b16b5af511cf03565ca8c35e4a75a.BoundingRectWithAnotherRangeRequestBuilder) {
    return i19f450b56b11336764c237e45c777ddc376b16b5af511cf03565ca8c35e4a75a.NewBoundingRectWithAnotherRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter, anotherRange);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.cell(row={row},column={column})
// Parameters:
//  - column : Usage: column={column}
//  - row : Usage: row={row}
func (m *WorkbookRangeRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*i4c53bde08a7009a7c80c7c2b20e6f4fb19dde1d576571b2ef6de2b99e885617c.CellWithRowWithColumnRequestBuilder) {
    return i4c53bde08a7009a7c80c7c2b20e6f4fb19dde1d576571b2ef6de2b99e885617c.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorkbookRangeRequestBuilder) Clear()(*icaf0bc4e754c5f4f6c34786a769aa4a76b1547e8bbb75bbf625caab2fb29a1b2.ClearRequestBuilder) {
    return icaf0bc4e754c5f4f6c34786a769aa4a76b1547e8bbb75bbf625caab2fb29a1b2.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsAfter()
func (m *WorkbookRangeRequestBuilder) ColumnsAfter()(*i2b7ce90d80d4123dfb21e7605c2fb8994d11bd546fef034f7b458fbe88efcb27.ColumnsAfterRequestBuilder) {
    return i2b7ce90d80d4123dfb21e7605c2fb8994d11bd546fef034f7b458fbe88efcb27.NewColumnsAfterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsAfter(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) ColumnsAfterWithCount(count *int32)(*i9622dcedf3086727437cd583adeda01e9c0afda0e71bc76178c7fc044e008f1b.ColumnsAfterWithCountRequestBuilder) {
    return i9622dcedf3086727437cd583adeda01e9c0afda0e71bc76178c7fc044e008f1b.NewColumnsAfterWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsBefore()
func (m *WorkbookRangeRequestBuilder) ColumnsBefore()(*ie138fb5e04a8819c265fe49eb6107d57526826809fc0ff38bc8744e46ab167a6.ColumnsBeforeRequestBuilder) {
    return ie138fb5e04a8819c265fe49eb6107d57526826809fc0ff38bc8744e46ab167a6.NewColumnsBeforeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsBefore(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) ColumnsBeforeWithCount(count *int32)(*ibc4295fac6d454478f729b247eb1bb6c6df70cd4620de9a1d702ae8731648405.ColumnsBeforeWithCountRequestBuilder) {
    return ibc4295fac6d454478f729b247eb1bb6c6df70cd4620de9a1d702ae8731648405.NewColumnsBeforeWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.column(column={column})
// Parameters:
//  - column : Usage: column={column}
func (m *WorkbookRangeRequestBuilder) ColumnWithColumn(column *int32)(*ic8daa19dcb2cb23d286fa4b858f2b7483e3277177c165ac869f177dc87477318.ColumnWithColumnRequestBuilder) {
    return ic8daa19dcb2cb23d286fa4b858f2b7483e3277177c165ac869f177dc87477318.NewColumnWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, column);
}
// Instantiates a new WorkbookRangeRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookRangeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeRequestBuilder) {
    m := &WorkbookRangeRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/used/{usedInsight_id}/resource/microsoft.graph.workbookRange";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new WorkbookRangeRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookRangeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookRangeRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorkbookRangeRequestBuilder) Delete()(*ib299e74ddaae1d25144b4eabfedc925b93bad3997b891249ba5522c4be887ea5.DeleteRequestBuilder) {
    return ib299e74ddaae1d25144b4eabfedc925b93bad3997b891249ba5522c4be887ea5.NewDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.entireColumn()
func (m *WorkbookRangeRequestBuilder) EntireColumn()(*id4aa6042f1f790a2c032e89a3bf64f66ea3fe7b73ddd8907b120b374d8b9e58d.EntireColumnRequestBuilder) {
    return id4aa6042f1f790a2c032e89a3bf64f66ea3fe7b73ddd8907b120b374d8b9e58d.NewEntireColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.entireRow()
func (m *WorkbookRangeRequestBuilder) EntireRow()(*i03d1153ceb9ce7e039b43442ba57e1648785a12bdd09b592ed3f4e656e311e5a.EntireRowRequestBuilder) {
    return i03d1153ceb9ce7e039b43442ba57e1648785a12bdd09b592ed3f4e656e311e5a.NewEntireRowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) Insert()(*icc8c73281fa3f1a351101ae8c02b7004fbedc8894e64f758040b955c094190a9.InsertRequestBuilder) {
    return icc8c73281fa3f1a351101ae8c02b7004fbedc8894e64f758040b955c094190a9.NewInsertRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.intersection(anotherRange='{anotherRange}')
// Parameters:
//  - anotherRange : Usage: anotherRange={anotherRange}
func (m *WorkbookRangeRequestBuilder) IntersectionWithAnotherRange(anotherRange *string)(*i0137e90048b8192c70d2981f12f5825552abe1d7253d2de5c5cf983392b76f2e.IntersectionWithAnotherRangeRequestBuilder) {
    return i0137e90048b8192c70d2981f12f5825552abe1d7253d2de5c5cf983392b76f2e.NewIntersectionWithAnotherRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter, anotherRange);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.lastCell()
func (m *WorkbookRangeRequestBuilder) LastCell()(*i4a3b7d2a8e3f4908a7e8203dc20e24926de35a06af94c2ff174624744a5cb191.LastCellRequestBuilder) {
    return i4a3b7d2a8e3f4908a7e8203dc20e24926de35a06af94c2ff174624744a5cb191.NewLastCellRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.lastColumn()
func (m *WorkbookRangeRequestBuilder) LastColumn()(*i8a5311b23354ddf82bcd1b646f3eb63adfe8d773b4606708ccb7c539b3faec3c.LastColumnRequestBuilder) {
    return i8a5311b23354ddf82bcd1b646f3eb63adfe8d773b4606708ccb7c539b3faec3c.NewLastColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.lastRow()
func (m *WorkbookRangeRequestBuilder) LastRow()(*i91ca2d416bc8c7860958a099103a9545cd2b8519b5cfac9ab760a45877ebfea1.LastRowRequestBuilder) {
    return i91ca2d416bc8c7860958a099103a9545cd2b8519b5cfac9ab760a45877ebfea1.NewLastRowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) Merge()(*i63f4f8673d2c326dbfa2211124f841d7bb63294db97302c6b4c3cb1ab15465f1.MergeRequestBuilder) {
    return i63f4f8673d2c326dbfa2211124f841d7bb63294db97302c6b4c3cb1ab15465f1.NewMergeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.offsetRange(rowOffset={rowOffset},columnOffset={columnOffset})
// Parameters:
//  - columnOffset : Usage: columnOffset={columnOffset}
//  - rowOffset : Usage: rowOffset={rowOffset}
func (m *WorkbookRangeRequestBuilder) OffsetRangeWithRowOffsetWithColumnOffset(rowOffset *int32, columnOffset *int32)(*i5f67ab708a56b6bfa64606d70ab83ed8f85a5ad38d56b02a59e4c519e8fc4be1.OffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder) {
    return i5f67ab708a56b6bfa64606d70ab83ed8f85a5ad38d56b02a59e4c519e8fc4be1.NewOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilderInternal(m.pathParameters, m.requestAdapter, rowOffset, columnOffset);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.resizedRange(deltaRows={deltaRows},deltaColumns={deltaColumns})
// Parameters:
//  - deltaColumns : Usage: deltaColumns={deltaColumns}
//  - deltaRows : Usage: deltaRows={deltaRows}
func (m *WorkbookRangeRequestBuilder) ResizedRangeWithDeltaRowsWithDeltaColumns(deltaRows *int32, deltaColumns *int32)(*i7aecccf6a4efcedf729c75ef5d6f8c686d20d07d04ec6bc6f9d4ab9f04c41fad.ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) {
    return i7aecccf6a4efcedf729c75ef5d6f8c686d20d07d04ec6bc6f9d4ab9f04c41fad.NewResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter, deltaRows, deltaColumns);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsAbove()
func (m *WorkbookRangeRequestBuilder) RowsAbove()(*iaa6d3acf8e841006360a3954e830cf9c5b5832d607648a97383b999ae3b3f331.RowsAboveRequestBuilder) {
    return iaa6d3acf8e841006360a3954e830cf9c5b5832d607648a97383b999ae3b3f331.NewRowsAboveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsAbove(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) RowsAboveWithCount(count *int32)(*i02a5c56469b42bdc0ce5f430a5f0612795fe6b4fd97997dcd4b447be7b5ba861.RowsAboveWithCountRequestBuilder) {
    return i02a5c56469b42bdc0ce5f430a5f0612795fe6b4fd97997dcd4b447be7b5ba861.NewRowsAboveWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsBelow()
func (m *WorkbookRangeRequestBuilder) RowsBelow()(*ib48c4a281d002ad30af16576e123a5717e6b413a11f4aa3480d03f4f22ca1f00.RowsBelowRequestBuilder) {
    return ib48c4a281d002ad30af16576e123a5717e6b413a11f4aa3480d03f4f22ca1f00.NewRowsBelowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsBelow(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) RowsBelowWithCount(count *int32)(*ib0248ff240f5b5ad9f326a9b5746a4baa736d280af427f68c75e2ac2f8562659.RowsBelowWithCountRequestBuilder) {
    return ib0248ff240f5b5ad9f326a9b5746a4baa736d280af427f68c75e2ac2f8562659.NewRowsBelowWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.row(row={row})
// Parameters:
//  - row : Usage: row={row}
func (m *WorkbookRangeRequestBuilder) RowWithRow(row *int32)(*i7d62b3f6ca1d7d823fee23e1c2f2c436e7330247eff72970e4399e294d915ec7.RowWithRowRequestBuilder) {
    return i7d62b3f6ca1d7d823fee23e1c2f2c436e7330247eff72970e4399e294d915ec7.NewRowWithRowRequestBuilderInternal(m.pathParameters, m.requestAdapter, row);
}
func (m *WorkbookRangeRequestBuilder) Unmerge()(*iefe5f75095bc26602696a7643255e6c4c3650219bb4d1eb8c3361ae039b10826.UnmergeRequestBuilder) {
    return iefe5f75095bc26602696a7643255e6c4c3650219bb4d1eb8c3361ae039b10826.NewUnmergeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.usedRange()
func (m *WorkbookRangeRequestBuilder) UsedRange()(*i97f27ba112adda84a04bdf4eac84a02427339ff52514510201e60df9eebe4389.UsedRangeRequestBuilder) {
    return i97f27ba112adda84a04bdf4eac84a02427339ff52514510201e60df9eebe4389.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.usedRange(valuesOnly={valuesOnly})
// Parameters:
//  - valuesOnly : Usage: valuesOnly={valuesOnly}
func (m *WorkbookRangeRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*ib8133d3713a6a700af7f7ac07448caddea3e14b46ef7b4119fb5db8fa90cfc42.UsedRangeWithValuesOnlyRequestBuilder) {
    return ib8133d3713a6a700af7f7ac07448caddea3e14b46ef7b4119fb5db8fa90cfc42.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.visibleView()
func (m *WorkbookRangeRequestBuilder) VisibleView()(*i8cda08c52c0efdece117e0fb29abf49f789fbe1d34cc095dda7cc6cb29253967.VisibleViewRequestBuilder) {
    return i8cda08c52c0efdece117e0fb29abf49f789fbe1d34cc095dda7cc6cb29253967.NewVisibleViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
