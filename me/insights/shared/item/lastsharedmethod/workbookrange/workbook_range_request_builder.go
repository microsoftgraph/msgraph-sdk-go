package workbookrange

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i00570291f21ef9bcce805276965080d5fe5cf59c56e2203068c6ae1ce2500117 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/unmerge"
    i1300b7eeec7e6138a3ee0a4036d6e1d3bf1a0be66c6d6ffd5cc38654d519d76f "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/resizedrangewithdeltarowswithdeltacolumns"
    i131be0f436bf168147ed562d8ee841282b99b5285b69b996411a12f41d207349 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/columnsbeforewithcount"
    i17fe3438c1fb6b98100f27f649b10ad3fbbc638156a5914b814432e9d150e400 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/rowsabovewithcount"
    i2350a9ff0ea0e38686437afca423cbbb70316955a2ba058d197f28a04837d2c4 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/lastcolumn"
    i253d00384404368963a37f0fac16c87f83f0ab4c99a88ece7305168df67deec5 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/rowsbelow"
    i3b38912249b43f0578fcbb7d85db671fcdd806bd17b2bc2cc73ef22db288d467 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/entirerow"
    i3bbb8032bdb95e1c4166d5c53b54fe701b9fda40980aed50eb2b11242cf478ae "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/usedrangewithvaluesonly"
    i4a5498fef18f91cea6d53ffd13dae2a4922992b31df8c3a374c305e91d75909e "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/intersectionwithanotherrange"
    i5771f5a9eb07d2410ee285f5c85e6ca1c8669ea4f785acc2c6e38121dc623172 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/lastcell"
    i60fe21d5d3773c8e3d8e1fb6a3eb603695b56fcacf6c9b72288e5e1e359407f7 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/rowwithrow"
    i6d4523716cead950928c25f28b83336fb35208796cb359eac3ead77d59ace0ef "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/offsetrangewithrowoffsetwithcolumnoffset"
    i7195920ce31da01a513887259851a4d8b75532503ccccb80d50f24f5b8cbfe11 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/rowsabove"
    i72cfd03c37ae740b28a12f5c66bf8e15c237ac3bfec86f21d4760a8026d332ff "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/columnsafter"
    i73647cb5cf18da60e69237942379b66c04aa16e0f8e3a82c06148a1a280afd25 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/visibleview"
    i7771f79e777df81ad2090d05500928ccf5d6213c3fa226be4aa03d3799a75fc1 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/lastrow"
    i80cf4e20f0f35e4571634be5d4c861a94c357b2ff3a7447fe1a6bb90eb6fbd37 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/usedrange"
    i8cc0d3b0b5703b55333ce1eea846f5edbdad3c95c3c8dc05cbd8843808feca0e "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/delete"
    i9c57aa6f1c30b0845e3fe1ef0e17e60f65caf5eb78cbf4c75b7e7fae7eba8356 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/merge"
    iaecb63ed1cbf2afad11f67a5c5bac7ed4677c9ae3fc0526002d80eed84ace3d4 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/rowsbelowwithcount"
    ib474875cc5f0d4d2705c2ee1ecd7cb8cc6124a2985ba942a5033b8575b7f90a7 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/cellwithrowwithcolumn"
    ib62df1d795056c7ba8d130c02918dbdba3324d3cef5bba5eeaa93af8b5e31f9c "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/entirecolumn"
    ic292eb397f1091718c85a6a298e5357f0f884c06ec3bca49f74957dbcde1e053 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/columnsbefore"
    ic54619b5323bf95e61eb2c455c1889ac37cca1f6363f37e85a62d2a300201709 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/columnwithcolumn"
    id54c1d9e5ac0d835161686f7ce9be4a2ffed194bcc648bead2caaa06e8200019 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/insert"
    id5ca968177092d230eacf79ec56912cbdd6c6dbaee2a20d351e0da665d0037ba "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/boundingrectwithanotherrange"
    ieb7c733fbd84dcd9df78320828720e07ce080c00e6b2eceee0e46e796aabd4df "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/clear"
    iefde87583ee1f0a0fc47032a7b2ee2ce5b68ae18538ee231123be2c834f492dc "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrange/columnsafterwithcount"
)

// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange
type WorkbookRangeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.boundingRect(anotherRange='{anotherRange}')
// Parameters:
//  - anotherRange : Usage: anotherRange={anotherRange}
func (m *WorkbookRangeRequestBuilder) BoundingRectWithAnotherRange(anotherRange *string)(*id5ca968177092d230eacf79ec56912cbdd6c6dbaee2a20d351e0da665d0037ba.BoundingRectWithAnotherRangeRequestBuilder) {
    return id5ca968177092d230eacf79ec56912cbdd6c6dbaee2a20d351e0da665d0037ba.NewBoundingRectWithAnotherRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter, anotherRange);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.cell(row={row},column={column})
// Parameters:
//  - column : Usage: column={column}
//  - row : Usage: row={row}
func (m *WorkbookRangeRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*ib474875cc5f0d4d2705c2ee1ecd7cb8cc6124a2985ba942a5033b8575b7f90a7.CellWithRowWithColumnRequestBuilder) {
    return ib474875cc5f0d4d2705c2ee1ecd7cb8cc6124a2985ba942a5033b8575b7f90a7.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorkbookRangeRequestBuilder) Clear()(*ieb7c733fbd84dcd9df78320828720e07ce080c00e6b2eceee0e46e796aabd4df.ClearRequestBuilder) {
    return ieb7c733fbd84dcd9df78320828720e07ce080c00e6b2eceee0e46e796aabd4df.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.columnsAfter()
func (m *WorkbookRangeRequestBuilder) ColumnsAfter()(*i72cfd03c37ae740b28a12f5c66bf8e15c237ac3bfec86f21d4760a8026d332ff.ColumnsAfterRequestBuilder) {
    return i72cfd03c37ae740b28a12f5c66bf8e15c237ac3bfec86f21d4760a8026d332ff.NewColumnsAfterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.columnsAfter(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) ColumnsAfterWithCount(count *int32)(*iefde87583ee1f0a0fc47032a7b2ee2ce5b68ae18538ee231123be2c834f492dc.ColumnsAfterWithCountRequestBuilder) {
    return iefde87583ee1f0a0fc47032a7b2ee2ce5b68ae18538ee231123be2c834f492dc.NewColumnsAfterWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.columnsBefore()
func (m *WorkbookRangeRequestBuilder) ColumnsBefore()(*ic292eb397f1091718c85a6a298e5357f0f884c06ec3bca49f74957dbcde1e053.ColumnsBeforeRequestBuilder) {
    return ic292eb397f1091718c85a6a298e5357f0f884c06ec3bca49f74957dbcde1e053.NewColumnsBeforeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.columnsBefore(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) ColumnsBeforeWithCount(count *int32)(*i131be0f436bf168147ed562d8ee841282b99b5285b69b996411a12f41d207349.ColumnsBeforeWithCountRequestBuilder) {
    return i131be0f436bf168147ed562d8ee841282b99b5285b69b996411a12f41d207349.NewColumnsBeforeWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.column(column={column})
// Parameters:
//  - column : Usage: column={column}
func (m *WorkbookRangeRequestBuilder) ColumnWithColumn(column *int32)(*ic54619b5323bf95e61eb2c455c1889ac37cca1f6363f37e85a62d2a300201709.ColumnWithColumnRequestBuilder) {
    return ic54619b5323bf95e61eb2c455c1889ac37cca1f6363f37e85a62d2a300201709.NewColumnWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, column);
}
// Instantiates a new WorkbookRangeRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookRangeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeRequestBuilder) {
    m := &WorkbookRangeRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/shared/{sharedInsight_id}/lastSharedMethod/microsoft.graph.workbookRange";
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
func (m *WorkbookRangeRequestBuilder) Delete()(*i8cc0d3b0b5703b55333ce1eea846f5edbdad3c95c3c8dc05cbd8843808feca0e.DeleteRequestBuilder) {
    return i8cc0d3b0b5703b55333ce1eea846f5edbdad3c95c3c8dc05cbd8843808feca0e.NewDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.entireColumn()
func (m *WorkbookRangeRequestBuilder) EntireColumn()(*ib62df1d795056c7ba8d130c02918dbdba3324d3cef5bba5eeaa93af8b5e31f9c.EntireColumnRequestBuilder) {
    return ib62df1d795056c7ba8d130c02918dbdba3324d3cef5bba5eeaa93af8b5e31f9c.NewEntireColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.entireRow()
func (m *WorkbookRangeRequestBuilder) EntireRow()(*i3b38912249b43f0578fcbb7d85db671fcdd806bd17b2bc2cc73ef22db288d467.EntireRowRequestBuilder) {
    return i3b38912249b43f0578fcbb7d85db671fcdd806bd17b2bc2cc73ef22db288d467.NewEntireRowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) Insert()(*id54c1d9e5ac0d835161686f7ce9be4a2ffed194bcc648bead2caaa06e8200019.InsertRequestBuilder) {
    return id54c1d9e5ac0d835161686f7ce9be4a2ffed194bcc648bead2caaa06e8200019.NewInsertRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.intersection(anotherRange='{anotherRange}')
// Parameters:
//  - anotherRange : Usage: anotherRange={anotherRange}
func (m *WorkbookRangeRequestBuilder) IntersectionWithAnotherRange(anotherRange *string)(*i4a5498fef18f91cea6d53ffd13dae2a4922992b31df8c3a374c305e91d75909e.IntersectionWithAnotherRangeRequestBuilder) {
    return i4a5498fef18f91cea6d53ffd13dae2a4922992b31df8c3a374c305e91d75909e.NewIntersectionWithAnotherRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter, anotherRange);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.lastCell()
func (m *WorkbookRangeRequestBuilder) LastCell()(*i5771f5a9eb07d2410ee285f5c85e6ca1c8669ea4f785acc2c6e38121dc623172.LastCellRequestBuilder) {
    return i5771f5a9eb07d2410ee285f5c85e6ca1c8669ea4f785acc2c6e38121dc623172.NewLastCellRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.lastColumn()
func (m *WorkbookRangeRequestBuilder) LastColumn()(*i2350a9ff0ea0e38686437afca423cbbb70316955a2ba058d197f28a04837d2c4.LastColumnRequestBuilder) {
    return i2350a9ff0ea0e38686437afca423cbbb70316955a2ba058d197f28a04837d2c4.NewLastColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.lastRow()
func (m *WorkbookRangeRequestBuilder) LastRow()(*i7771f79e777df81ad2090d05500928ccf5d6213c3fa226be4aa03d3799a75fc1.LastRowRequestBuilder) {
    return i7771f79e777df81ad2090d05500928ccf5d6213c3fa226be4aa03d3799a75fc1.NewLastRowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) Merge()(*i9c57aa6f1c30b0845e3fe1ef0e17e60f65caf5eb78cbf4c75b7e7fae7eba8356.MergeRequestBuilder) {
    return i9c57aa6f1c30b0845e3fe1ef0e17e60f65caf5eb78cbf4c75b7e7fae7eba8356.NewMergeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.offsetRange(rowOffset={rowOffset},columnOffset={columnOffset})
// Parameters:
//  - columnOffset : Usage: columnOffset={columnOffset}
//  - rowOffset : Usage: rowOffset={rowOffset}
func (m *WorkbookRangeRequestBuilder) OffsetRangeWithRowOffsetWithColumnOffset(rowOffset *int32, columnOffset *int32)(*i6d4523716cead950928c25f28b83336fb35208796cb359eac3ead77d59ace0ef.OffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder) {
    return i6d4523716cead950928c25f28b83336fb35208796cb359eac3ead77d59ace0ef.NewOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilderInternal(m.pathParameters, m.requestAdapter, rowOffset, columnOffset);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.resizedRange(deltaRows={deltaRows},deltaColumns={deltaColumns})
// Parameters:
//  - deltaColumns : Usage: deltaColumns={deltaColumns}
//  - deltaRows : Usage: deltaRows={deltaRows}
func (m *WorkbookRangeRequestBuilder) ResizedRangeWithDeltaRowsWithDeltaColumns(deltaRows *int32, deltaColumns *int32)(*i1300b7eeec7e6138a3ee0a4036d6e1d3bf1a0be66c6d6ffd5cc38654d519d76f.ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) {
    return i1300b7eeec7e6138a3ee0a4036d6e1d3bf1a0be66c6d6ffd5cc38654d519d76f.NewResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter, deltaRows, deltaColumns);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.rowsAbove()
func (m *WorkbookRangeRequestBuilder) RowsAbove()(*i7195920ce31da01a513887259851a4d8b75532503ccccb80d50f24f5b8cbfe11.RowsAboveRequestBuilder) {
    return i7195920ce31da01a513887259851a4d8b75532503ccccb80d50f24f5b8cbfe11.NewRowsAboveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.rowsAbove(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) RowsAboveWithCount(count *int32)(*i17fe3438c1fb6b98100f27f649b10ad3fbbc638156a5914b814432e9d150e400.RowsAboveWithCountRequestBuilder) {
    return i17fe3438c1fb6b98100f27f649b10ad3fbbc638156a5914b814432e9d150e400.NewRowsAboveWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.rowsBelow()
func (m *WorkbookRangeRequestBuilder) RowsBelow()(*i253d00384404368963a37f0fac16c87f83f0ab4c99a88ece7305168df67deec5.RowsBelowRequestBuilder) {
    return i253d00384404368963a37f0fac16c87f83f0ab4c99a88ece7305168df67deec5.NewRowsBelowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.rowsBelow(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) RowsBelowWithCount(count *int32)(*iaecb63ed1cbf2afad11f67a5c5bac7ed4677c9ae3fc0526002d80eed84ace3d4.RowsBelowWithCountRequestBuilder) {
    return iaecb63ed1cbf2afad11f67a5c5bac7ed4677c9ae3fc0526002d80eed84ace3d4.NewRowsBelowWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.row(row={row})
// Parameters:
//  - row : Usage: row={row}
func (m *WorkbookRangeRequestBuilder) RowWithRow(row *int32)(*i60fe21d5d3773c8e3d8e1fb6a3eb603695b56fcacf6c9b72288e5e1e359407f7.RowWithRowRequestBuilder) {
    return i60fe21d5d3773c8e3d8e1fb6a3eb603695b56fcacf6c9b72288e5e1e359407f7.NewRowWithRowRequestBuilderInternal(m.pathParameters, m.requestAdapter, row);
}
func (m *WorkbookRangeRequestBuilder) Unmerge()(*i00570291f21ef9bcce805276965080d5fe5cf59c56e2203068c6ae1ce2500117.UnmergeRequestBuilder) {
    return i00570291f21ef9bcce805276965080d5fe5cf59c56e2203068c6ae1ce2500117.NewUnmergeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.usedRange()
func (m *WorkbookRangeRequestBuilder) UsedRange()(*i80cf4e20f0f35e4571634be5d4c861a94c357b2ff3a7447fe1a6bb90eb6fbd37.UsedRangeRequestBuilder) {
    return i80cf4e20f0f35e4571634be5d4c861a94c357b2ff3a7447fe1a6bb90eb6fbd37.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.usedRange(valuesOnly={valuesOnly})
// Parameters:
//  - valuesOnly : Usage: valuesOnly={valuesOnly}
func (m *WorkbookRangeRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*i3bbb8032bdb95e1c4166d5c53b54fe701b9fda40980aed50eb2b11242cf478ae.UsedRangeWithValuesOnlyRequestBuilder) {
    return i3bbb8032bdb95e1c4166d5c53b54fe701b9fda40980aed50eb2b11242cf478ae.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.visibleView()
func (m *WorkbookRangeRequestBuilder) VisibleView()(*i73647cb5cf18da60e69237942379b66c04aa16e0f8e3a82c06148a1a280afd25.VisibleViewRequestBuilder) {
    return i73647cb5cf18da60e69237942379b66c04aa16e0f8e3a82c06148a1a280afd25.NewVisibleViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
