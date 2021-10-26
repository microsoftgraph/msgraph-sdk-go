package workbookrange

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i1e8d662baae26fe1ea7a0d16feeb7355ae29d36751bc62f88ff5d7719f2eb34d "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/rowsbelow"
    i1f4889c485cf003c7b4693df25cdd2a92901ea27b68668d8f51c9b2e89dd1172 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/columnsafterwithcount"
    i3e00d2d4d46dd8bca0d0f9253f22943191c05828ebf31ca61ec88932fdf392cd "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/lastcolumn"
    i4a5518b3b450ec7684b6f1e7e0990d85a43dde53f96bd294d9196fe6ab0e7169 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/offsetrangewithrowoffsetwithcolumnoffset"
    i4ae5debb30049b48551554e2e2d9ccd4e0d698441cb6197f00549e0c7e355f88 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/clear"
    i52413ac5a4295800b6ba37bf866dcac9a0c0baaf70a897cb791e502318cae298 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/entirecolumn"
    i5267d3bea1334a01bdd92dca88ef35067c5619d047e144303fbcd354cac3c1cb "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/cellwithrowwithcolumn"
    i5b1984cd842172287d9ac67d11bc6c034940e07eab228db135896f8e249d4316 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/intersectionwithanotherrange"
    i5f411d3a182cb27a2dbd9309d33508ae749d13c1d161d300fe226a4640f76185 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/usedrangewithvaluesonly"
    i71b04e6f10300fa2b50573d371eb182611d92f10704bfe0ce837d0d07bfaf993 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/rowsabove"
    i78ef0ff3c7bad1896913906521bf4940418549b0a7e4f2a47a6bbf8476055107 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/delete"
    i7c0e104600f698d5931b1c9899009c58afa5a030ce25c36052f5ffe82a5eca0f "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/columnsbeforewithcount"
    i86bef26a77058a70c5d774e94109add64b80e3686a7fd913b7f0afbcbc820d8c "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/lastrow"
    i925b21205de5a9378a69e3d49210a372f89576249c733a70590524298013de77 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/lastcell"
    ia358f9fc768288ab84f6625af4f8190cfb394b370f7f891ca4ef942259caf6fe "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/insert"
    ia4630e60212ad5b67ba052793e72b469920e0187a457791567a6c1790dc04a1d "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/unmerge"
    ib3d78911f37fbc3c9e7cdf3879fc1bccf01419728f91dd710255fc2b01e4730a "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/entirerow"
    ib4ecc4df242f347a18389172fb0ee5b487844a2f61cf869540c24d11c2014813 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/boundingrectwithanotherrange"
    ibd1cac016c78abaf8f81f918fa803b395b3be548dfce73b275768db35beea548 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/visibleview"
    ic05a7a85f38b4853c3e9a6142b62f24e8e82888579d512b52c3c82652ece3c8f "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/columnsbefore"
    ic0744fbc06721d08346cb6a0e737e8e4ead3881b84226177717f2045fdc4d3ac "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/rowwithrow"
    id5cda94a2519f449dc02309746989b0245b610a4bc5c82408851bf5a67645153 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/resizedrangewithdeltarowswithdeltacolumns"
    ie1b79187c5eccba4c1bb7e3eeed080d17fd978bf70c9542727af505e81ff9654 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/columnwithcolumn"
    ie9af03167dd204cf1a10c982f69bff9eed77010fa768d3c55c152c32a67505f5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/columnsafter"
    ied2281ca306bad5aa91e56f929366667228fdefbf5d9bcae75139dcad038b8a7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/rowsbelowwithcount"
    ieeda21696be409264c9f7e028dea483dd76be31743d706eab3468e9effdf0c46 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/rowsabovewithcount"
    if3bcbe1287fd552df598b4ae72125e67c5ac6293f30d527d0e8eefbef1291308 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/merge"
    if7d8462b49973e2b5bc7ff5b104829350d96e21c6f8f6657aba751ea60aebe8c "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange/usedrange"
)

// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange
type WorkbookRangeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.boundingRect(anotherRange='{anotherRange}')
// Parameters:
//  - anotherRange : Usage: anotherRange={anotherRange}
func (m *WorkbookRangeRequestBuilder) BoundingRectWithAnotherRange(anotherRange *string)(*ib4ecc4df242f347a18389172fb0ee5b487844a2f61cf869540c24d11c2014813.BoundingRectWithAnotherRangeRequestBuilder) {
    return ib4ecc4df242f347a18389172fb0ee5b487844a2f61cf869540c24d11c2014813.NewBoundingRectWithAnotherRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter, anotherRange);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.cell(row={row},column={column})
// Parameters:
//  - column : Usage: column={column}
//  - row : Usage: row={row}
func (m *WorkbookRangeRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*i5267d3bea1334a01bdd92dca88ef35067c5619d047e144303fbcd354cac3c1cb.CellWithRowWithColumnRequestBuilder) {
    return i5267d3bea1334a01bdd92dca88ef35067c5619d047e144303fbcd354cac3c1cb.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorkbookRangeRequestBuilder) Clear()(*i4ae5debb30049b48551554e2e2d9ccd4e0d698441cb6197f00549e0c7e355f88.ClearRequestBuilder) {
    return i4ae5debb30049b48551554e2e2d9ccd4e0d698441cb6197f00549e0c7e355f88.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.columnsAfter()
func (m *WorkbookRangeRequestBuilder) ColumnsAfter()(*ie9af03167dd204cf1a10c982f69bff9eed77010fa768d3c55c152c32a67505f5.ColumnsAfterRequestBuilder) {
    return ie9af03167dd204cf1a10c982f69bff9eed77010fa768d3c55c152c32a67505f5.NewColumnsAfterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.columnsAfter(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) ColumnsAfterWithCount(count *int32)(*i1f4889c485cf003c7b4693df25cdd2a92901ea27b68668d8f51c9b2e89dd1172.ColumnsAfterWithCountRequestBuilder) {
    return i1f4889c485cf003c7b4693df25cdd2a92901ea27b68668d8f51c9b2e89dd1172.NewColumnsAfterWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.columnsBefore()
func (m *WorkbookRangeRequestBuilder) ColumnsBefore()(*ic05a7a85f38b4853c3e9a6142b62f24e8e82888579d512b52c3c82652ece3c8f.ColumnsBeforeRequestBuilder) {
    return ic05a7a85f38b4853c3e9a6142b62f24e8e82888579d512b52c3c82652ece3c8f.NewColumnsBeforeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.columnsBefore(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) ColumnsBeforeWithCount(count *int32)(*i7c0e104600f698d5931b1c9899009c58afa5a030ce25c36052f5ffe82a5eca0f.ColumnsBeforeWithCountRequestBuilder) {
    return i7c0e104600f698d5931b1c9899009c58afa5a030ce25c36052f5ffe82a5eca0f.NewColumnsBeforeWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.column(column={column})
// Parameters:
//  - column : Usage: column={column}
func (m *WorkbookRangeRequestBuilder) ColumnWithColumn(column *int32)(*ie1b79187c5eccba4c1bb7e3eeed080d17fd978bf70c9542727af505e81ff9654.ColumnWithColumnRequestBuilder) {
    return ie1b79187c5eccba4c1bb7e3eeed080d17fd978bf70c9542727af505e81ff9654.NewColumnWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, column);
}
// Instantiates a new WorkbookRangeRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookRangeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeRequestBuilder) {
    m := &WorkbookRangeRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/shared/{sharedInsight_id}/lastSharedMethod/microsoft.graph.workbookRange";
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
func (m *WorkbookRangeRequestBuilder) Delete()(*i78ef0ff3c7bad1896913906521bf4940418549b0a7e4f2a47a6bbf8476055107.DeleteRequestBuilder) {
    return i78ef0ff3c7bad1896913906521bf4940418549b0a7e4f2a47a6bbf8476055107.NewDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.entireColumn()
func (m *WorkbookRangeRequestBuilder) EntireColumn()(*i52413ac5a4295800b6ba37bf866dcac9a0c0baaf70a897cb791e502318cae298.EntireColumnRequestBuilder) {
    return i52413ac5a4295800b6ba37bf866dcac9a0c0baaf70a897cb791e502318cae298.NewEntireColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.entireRow()
func (m *WorkbookRangeRequestBuilder) EntireRow()(*ib3d78911f37fbc3c9e7cdf3879fc1bccf01419728f91dd710255fc2b01e4730a.EntireRowRequestBuilder) {
    return ib3d78911f37fbc3c9e7cdf3879fc1bccf01419728f91dd710255fc2b01e4730a.NewEntireRowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) Insert()(*ia358f9fc768288ab84f6625af4f8190cfb394b370f7f891ca4ef942259caf6fe.InsertRequestBuilder) {
    return ia358f9fc768288ab84f6625af4f8190cfb394b370f7f891ca4ef942259caf6fe.NewInsertRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.intersection(anotherRange='{anotherRange}')
// Parameters:
//  - anotherRange : Usage: anotherRange={anotherRange}
func (m *WorkbookRangeRequestBuilder) IntersectionWithAnotherRange(anotherRange *string)(*i5b1984cd842172287d9ac67d11bc6c034940e07eab228db135896f8e249d4316.IntersectionWithAnotherRangeRequestBuilder) {
    return i5b1984cd842172287d9ac67d11bc6c034940e07eab228db135896f8e249d4316.NewIntersectionWithAnotherRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter, anotherRange);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.lastCell()
func (m *WorkbookRangeRequestBuilder) LastCell()(*i925b21205de5a9378a69e3d49210a372f89576249c733a70590524298013de77.LastCellRequestBuilder) {
    return i925b21205de5a9378a69e3d49210a372f89576249c733a70590524298013de77.NewLastCellRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.lastColumn()
func (m *WorkbookRangeRequestBuilder) LastColumn()(*i3e00d2d4d46dd8bca0d0f9253f22943191c05828ebf31ca61ec88932fdf392cd.LastColumnRequestBuilder) {
    return i3e00d2d4d46dd8bca0d0f9253f22943191c05828ebf31ca61ec88932fdf392cd.NewLastColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.lastRow()
func (m *WorkbookRangeRequestBuilder) LastRow()(*i86bef26a77058a70c5d774e94109add64b80e3686a7fd913b7f0afbcbc820d8c.LastRowRequestBuilder) {
    return i86bef26a77058a70c5d774e94109add64b80e3686a7fd913b7f0afbcbc820d8c.NewLastRowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) Merge()(*if3bcbe1287fd552df598b4ae72125e67c5ac6293f30d527d0e8eefbef1291308.MergeRequestBuilder) {
    return if3bcbe1287fd552df598b4ae72125e67c5ac6293f30d527d0e8eefbef1291308.NewMergeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.offsetRange(rowOffset={rowOffset},columnOffset={columnOffset})
// Parameters:
//  - columnOffset : Usage: columnOffset={columnOffset}
//  - rowOffset : Usage: rowOffset={rowOffset}
func (m *WorkbookRangeRequestBuilder) OffsetRangeWithRowOffsetWithColumnOffset(rowOffset *int32, columnOffset *int32)(*i4a5518b3b450ec7684b6f1e7e0990d85a43dde53f96bd294d9196fe6ab0e7169.OffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder) {
    return i4a5518b3b450ec7684b6f1e7e0990d85a43dde53f96bd294d9196fe6ab0e7169.NewOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilderInternal(m.pathParameters, m.requestAdapter, rowOffset, columnOffset);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.resizedRange(deltaRows={deltaRows},deltaColumns={deltaColumns})
// Parameters:
//  - deltaColumns : Usage: deltaColumns={deltaColumns}
//  - deltaRows : Usage: deltaRows={deltaRows}
func (m *WorkbookRangeRequestBuilder) ResizedRangeWithDeltaRowsWithDeltaColumns(deltaRows *int32, deltaColumns *int32)(*id5cda94a2519f449dc02309746989b0245b610a4bc5c82408851bf5a67645153.ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) {
    return id5cda94a2519f449dc02309746989b0245b610a4bc5c82408851bf5a67645153.NewResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter, deltaRows, deltaColumns);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.rowsAbove()
func (m *WorkbookRangeRequestBuilder) RowsAbove()(*i71b04e6f10300fa2b50573d371eb182611d92f10704bfe0ce837d0d07bfaf993.RowsAboveRequestBuilder) {
    return i71b04e6f10300fa2b50573d371eb182611d92f10704bfe0ce837d0d07bfaf993.NewRowsAboveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.rowsAbove(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) RowsAboveWithCount(count *int32)(*ieeda21696be409264c9f7e028dea483dd76be31743d706eab3468e9effdf0c46.RowsAboveWithCountRequestBuilder) {
    return ieeda21696be409264c9f7e028dea483dd76be31743d706eab3468e9effdf0c46.NewRowsAboveWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.rowsBelow()
func (m *WorkbookRangeRequestBuilder) RowsBelow()(*i1e8d662baae26fe1ea7a0d16feeb7355ae29d36751bc62f88ff5d7719f2eb34d.RowsBelowRequestBuilder) {
    return i1e8d662baae26fe1ea7a0d16feeb7355ae29d36751bc62f88ff5d7719f2eb34d.NewRowsBelowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.rowsBelow(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) RowsBelowWithCount(count *int32)(*ied2281ca306bad5aa91e56f929366667228fdefbf5d9bcae75139dcad038b8a7.RowsBelowWithCountRequestBuilder) {
    return ied2281ca306bad5aa91e56f929366667228fdefbf5d9bcae75139dcad038b8a7.NewRowsBelowWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.row(row={row})
// Parameters:
//  - row : Usage: row={row}
func (m *WorkbookRangeRequestBuilder) RowWithRow(row *int32)(*ic0744fbc06721d08346cb6a0e737e8e4ead3881b84226177717f2045fdc4d3ac.RowWithRowRequestBuilder) {
    return ic0744fbc06721d08346cb6a0e737e8e4ead3881b84226177717f2045fdc4d3ac.NewRowWithRowRequestBuilderInternal(m.pathParameters, m.requestAdapter, row);
}
func (m *WorkbookRangeRequestBuilder) Unmerge()(*ia4630e60212ad5b67ba052793e72b469920e0187a457791567a6c1790dc04a1d.UnmergeRequestBuilder) {
    return ia4630e60212ad5b67ba052793e72b469920e0187a457791567a6c1790dc04a1d.NewUnmergeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.usedRange()
func (m *WorkbookRangeRequestBuilder) UsedRange()(*if7d8462b49973e2b5bc7ff5b104829350d96e21c6f8f6657aba751ea60aebe8c.UsedRangeRequestBuilder) {
    return if7d8462b49973e2b5bc7ff5b104829350d96e21c6f8f6657aba751ea60aebe8c.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.usedRange(valuesOnly={valuesOnly})
// Parameters:
//  - valuesOnly : Usage: valuesOnly={valuesOnly}
func (m *WorkbookRangeRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*i5f411d3a182cb27a2dbd9309d33508ae749d13c1d161d300fe226a4640f76185.UsedRangeWithValuesOnlyRequestBuilder) {
    return i5f411d3a182cb27a2dbd9309d33508ae749d13c1d161d300fe226a4640f76185.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRange\microsoft.graph.visibleView()
func (m *WorkbookRangeRequestBuilder) VisibleView()(*ibd1cac016c78abaf8f81f918fa803b395b3be548dfce73b275768db35beea548.VisibleViewRequestBuilder) {
    return ibd1cac016c78abaf8f81f918fa803b395b3be548dfce73b275768db35beea548.NewVisibleViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
