package workbookrange

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i01a6ac6abcd930672aee5206abe7a41a3a0df732d02b0b33d1f14535db152d6b "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/offsetrangewithrowoffsetwithcolumnoffset"
    i03480ca880efabe11e1b7365b63d362b3bf8f87f370753912b22c647f1986970 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/clear"
    i08d6c176f22afc86bc7479df10d36128521583763ce85b91a7ac9790e023dc9d "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/lastcell"
    i171bd06f7f7612ddc969c9f4cdcaa87b819270288bc7f7d1f037998484f9892a "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/entirecolumn"
    i25fb10375154c97e26f50a33859645c6f5aafef746d6c6031f52cff3ef7221f0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/merge"
    i26cf31baad65fdb9d17303a5c661926c2cf3fef6a6a94e5f303a462b6210d55c "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/boundingrectwithanotherrange"
    i31b6945f9747094e604633de571ed46fc9a6941061abfc81c222b077bc90ced0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/rowsabovewithcount"
    i33ec2740019b8fa444273561c3f9262f47129ae8748264a5cd2cf29cbe8bc021 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/insert"
    i4a10b3104fb69bbfed19693a228ea3244148dad964ac9c7107edaab2b1e2153a "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/entirerow"
    i61fbc6095b6a3f0d7e7882dc19fd2712ede7e352a8c8696128a8e87402b3c186 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/rowwithrow"
    i65783c8d11d930a960df63614ef3009276773f269fe821aadcefa2e56494f3d1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/usedrange"
    i7f8a6e3fc0ef3113e8566e2989b6d447a1d04ee7e2f6fe1f2d15dd9e530c2bd1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/visibleview"
    i858fdbbd596afaf465437fa2301b4a91dfcf17e28fe4bdd23e5b627b9071a3c6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/lastcolumn"
    i897def161808cce8e54f7ddf5d062f0c9a93926b1b8ab0fe0ec2f0015e183c15 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/resizedrangewithdeltarowswithdeltacolumns"
    i8e265314ef374156ccee87bc526bcfc8e8f29ceedc982581cbf49f9f1098e8df "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/columnsbefore"
    i8eeaff30b5ff88d3e2edb3ce43c02b2665a183bec6da3f0a1aa9f5fbc1d60743 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/intersectionwithanotherrange"
    ia358f89504c007c9f749e41e4f15398cd1e54797872e044a95d0a784031ab907 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/usedrangewithvaluesonly"
    ia39ba689b46ce026dc4ebfe96e7da226daa330b94edf1f968401408c4ea36855 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/columnsafterwithcount"
    ia73293ecacc3045e9cde771d0a5fe7245734159742fb84451d08dd9a8ff6a30d "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/rowsbelow"
    iad438744f32df291867afd3cb89d30cb96dbdd970b5b82570c9f5eb7b15541a4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/columnsbeforewithcount"
    ib6348d6ac1e1f2b29f8e41eb36961e479de52779e64edefed26fb3e84e90ddbe "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/cellwithrowwithcolumn"
    ibd31d9e7d4988dd3f055287ff34f1724847bd79d1e1dc926de2b843bc6fadad7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/delete"
    icff8ecafbc5d6f09ff9258220aec08ce3150318b7f7305014894e9f657c988cb "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/rowsbelowwithcount"
    ie08995876b2e2f52a52c2b69aaa57390ad1eb04ca7b66cad3154a65b0578583b "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/columnwithcolumn"
    ie1648a89e1e2322978ca232dca1a66425eb36f6d387029b0089703d2e9e217d9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/unmerge"
    ie1667c27818b20b435a41fa0c35b0a1cfa78cd596bfe9055b1398a0d0badee5b "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/rowsabove"
    ie43eb65f3d96b65be3745d627726cbec24b8bf4660549411c1eada8ca04a52e5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/columnsafter"
    if8efb353e33d9509a82e882b71ea8fb9e197eda2dcdd0038da402313bbf1bcbe "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrange/lastrow"
)

// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange
type WorkbookRangeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.boundingRect(anotherRange='{anotherRange}')
// Parameters:
//  - anotherRange : Usage: anotherRange={anotherRange}
func (m *WorkbookRangeRequestBuilder) BoundingRectWithAnotherRange(anotherRange *string)(*i26cf31baad65fdb9d17303a5c661926c2cf3fef6a6a94e5f303a462b6210d55c.BoundingRectWithAnotherRangeRequestBuilder) {
    return i26cf31baad65fdb9d17303a5c661926c2cf3fef6a6a94e5f303a462b6210d55c.NewBoundingRectWithAnotherRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter, anotherRange);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.cell(row={row},column={column})
// Parameters:
//  - column : Usage: column={column}
//  - row : Usage: row={row}
func (m *WorkbookRangeRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*ib6348d6ac1e1f2b29f8e41eb36961e479de52779e64edefed26fb3e84e90ddbe.CellWithRowWithColumnRequestBuilder) {
    return ib6348d6ac1e1f2b29f8e41eb36961e479de52779e64edefed26fb3e84e90ddbe.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorkbookRangeRequestBuilder) Clear()(*i03480ca880efabe11e1b7365b63d362b3bf8f87f370753912b22c647f1986970.ClearRequestBuilder) {
    return i03480ca880efabe11e1b7365b63d362b3bf8f87f370753912b22c647f1986970.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsAfter()
func (m *WorkbookRangeRequestBuilder) ColumnsAfter()(*ie43eb65f3d96b65be3745d627726cbec24b8bf4660549411c1eada8ca04a52e5.ColumnsAfterRequestBuilder) {
    return ie43eb65f3d96b65be3745d627726cbec24b8bf4660549411c1eada8ca04a52e5.NewColumnsAfterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsAfter(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) ColumnsAfterWithCount(count *int32)(*ia39ba689b46ce026dc4ebfe96e7da226daa330b94edf1f968401408c4ea36855.ColumnsAfterWithCountRequestBuilder) {
    return ia39ba689b46ce026dc4ebfe96e7da226daa330b94edf1f968401408c4ea36855.NewColumnsAfterWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsBefore()
func (m *WorkbookRangeRequestBuilder) ColumnsBefore()(*i8e265314ef374156ccee87bc526bcfc8e8f29ceedc982581cbf49f9f1098e8df.ColumnsBeforeRequestBuilder) {
    return i8e265314ef374156ccee87bc526bcfc8e8f29ceedc982581cbf49f9f1098e8df.NewColumnsBeforeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsBefore(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) ColumnsBeforeWithCount(count *int32)(*iad438744f32df291867afd3cb89d30cb96dbdd970b5b82570c9f5eb7b15541a4.ColumnsBeforeWithCountRequestBuilder) {
    return iad438744f32df291867afd3cb89d30cb96dbdd970b5b82570c9f5eb7b15541a4.NewColumnsBeforeWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.column(column={column})
// Parameters:
//  - column : Usage: column={column}
func (m *WorkbookRangeRequestBuilder) ColumnWithColumn(column *int32)(*ie08995876b2e2f52a52c2b69aaa57390ad1eb04ca7b66cad3154a65b0578583b.ColumnWithColumnRequestBuilder) {
    return ie08995876b2e2f52a52c2b69aaa57390ad1eb04ca7b66cad3154a65b0578583b.NewColumnWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, column);
}
// Instantiates a new WorkbookRangeRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookRangeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeRequestBuilder) {
    m := &WorkbookRangeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights/shared/{sharedInsight_id}/resource/microsoft.graph.workbookRange";
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
func (m *WorkbookRangeRequestBuilder) Delete()(*ibd31d9e7d4988dd3f055287ff34f1724847bd79d1e1dc926de2b843bc6fadad7.DeleteRequestBuilder) {
    return ibd31d9e7d4988dd3f055287ff34f1724847bd79d1e1dc926de2b843bc6fadad7.NewDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.entireColumn()
func (m *WorkbookRangeRequestBuilder) EntireColumn()(*i171bd06f7f7612ddc969c9f4cdcaa87b819270288bc7f7d1f037998484f9892a.EntireColumnRequestBuilder) {
    return i171bd06f7f7612ddc969c9f4cdcaa87b819270288bc7f7d1f037998484f9892a.NewEntireColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.entireRow()
func (m *WorkbookRangeRequestBuilder) EntireRow()(*i4a10b3104fb69bbfed19693a228ea3244148dad964ac9c7107edaab2b1e2153a.EntireRowRequestBuilder) {
    return i4a10b3104fb69bbfed19693a228ea3244148dad964ac9c7107edaab2b1e2153a.NewEntireRowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) Insert()(*i33ec2740019b8fa444273561c3f9262f47129ae8748264a5cd2cf29cbe8bc021.InsertRequestBuilder) {
    return i33ec2740019b8fa444273561c3f9262f47129ae8748264a5cd2cf29cbe8bc021.NewInsertRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.intersection(anotherRange='{anotherRange}')
// Parameters:
//  - anotherRange : Usage: anotherRange={anotherRange}
func (m *WorkbookRangeRequestBuilder) IntersectionWithAnotherRange(anotherRange *string)(*i8eeaff30b5ff88d3e2edb3ce43c02b2665a183bec6da3f0a1aa9f5fbc1d60743.IntersectionWithAnotherRangeRequestBuilder) {
    return i8eeaff30b5ff88d3e2edb3ce43c02b2665a183bec6da3f0a1aa9f5fbc1d60743.NewIntersectionWithAnotherRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter, anotherRange);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.lastCell()
func (m *WorkbookRangeRequestBuilder) LastCell()(*i08d6c176f22afc86bc7479df10d36128521583763ce85b91a7ac9790e023dc9d.LastCellRequestBuilder) {
    return i08d6c176f22afc86bc7479df10d36128521583763ce85b91a7ac9790e023dc9d.NewLastCellRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.lastColumn()
func (m *WorkbookRangeRequestBuilder) LastColumn()(*i858fdbbd596afaf465437fa2301b4a91dfcf17e28fe4bdd23e5b627b9071a3c6.LastColumnRequestBuilder) {
    return i858fdbbd596afaf465437fa2301b4a91dfcf17e28fe4bdd23e5b627b9071a3c6.NewLastColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.lastRow()
func (m *WorkbookRangeRequestBuilder) LastRow()(*if8efb353e33d9509a82e882b71ea8fb9e197eda2dcdd0038da402313bbf1bcbe.LastRowRequestBuilder) {
    return if8efb353e33d9509a82e882b71ea8fb9e197eda2dcdd0038da402313bbf1bcbe.NewLastRowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) Merge()(*i25fb10375154c97e26f50a33859645c6f5aafef746d6c6031f52cff3ef7221f0.MergeRequestBuilder) {
    return i25fb10375154c97e26f50a33859645c6f5aafef746d6c6031f52cff3ef7221f0.NewMergeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.offsetRange(rowOffset={rowOffset},columnOffset={columnOffset})
// Parameters:
//  - columnOffset : Usage: columnOffset={columnOffset}
//  - rowOffset : Usage: rowOffset={rowOffset}
func (m *WorkbookRangeRequestBuilder) OffsetRangeWithRowOffsetWithColumnOffset(rowOffset *int32, columnOffset *int32)(*i01a6ac6abcd930672aee5206abe7a41a3a0df732d02b0b33d1f14535db152d6b.OffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder) {
    return i01a6ac6abcd930672aee5206abe7a41a3a0df732d02b0b33d1f14535db152d6b.NewOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilderInternal(m.pathParameters, m.requestAdapter, rowOffset, columnOffset);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.resizedRange(deltaRows={deltaRows},deltaColumns={deltaColumns})
// Parameters:
//  - deltaColumns : Usage: deltaColumns={deltaColumns}
//  - deltaRows : Usage: deltaRows={deltaRows}
func (m *WorkbookRangeRequestBuilder) ResizedRangeWithDeltaRowsWithDeltaColumns(deltaRows *int32, deltaColumns *int32)(*i897def161808cce8e54f7ddf5d062f0c9a93926b1b8ab0fe0ec2f0015e183c15.ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) {
    return i897def161808cce8e54f7ddf5d062f0c9a93926b1b8ab0fe0ec2f0015e183c15.NewResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter, deltaRows, deltaColumns);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsAbove()
func (m *WorkbookRangeRequestBuilder) RowsAbove()(*ie1667c27818b20b435a41fa0c35b0a1cfa78cd596bfe9055b1398a0d0badee5b.RowsAboveRequestBuilder) {
    return ie1667c27818b20b435a41fa0c35b0a1cfa78cd596bfe9055b1398a0d0badee5b.NewRowsAboveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsAbove(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) RowsAboveWithCount(count *int32)(*i31b6945f9747094e604633de571ed46fc9a6941061abfc81c222b077bc90ced0.RowsAboveWithCountRequestBuilder) {
    return i31b6945f9747094e604633de571ed46fc9a6941061abfc81c222b077bc90ced0.NewRowsAboveWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsBelow()
func (m *WorkbookRangeRequestBuilder) RowsBelow()(*ia73293ecacc3045e9cde771d0a5fe7245734159742fb84451d08dd9a8ff6a30d.RowsBelowRequestBuilder) {
    return ia73293ecacc3045e9cde771d0a5fe7245734159742fb84451d08dd9a8ff6a30d.NewRowsBelowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsBelow(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) RowsBelowWithCount(count *int32)(*icff8ecafbc5d6f09ff9258220aec08ce3150318b7f7305014894e9f657c988cb.RowsBelowWithCountRequestBuilder) {
    return icff8ecafbc5d6f09ff9258220aec08ce3150318b7f7305014894e9f657c988cb.NewRowsBelowWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.row(row={row})
// Parameters:
//  - row : Usage: row={row}
func (m *WorkbookRangeRequestBuilder) RowWithRow(row *int32)(*i61fbc6095b6a3f0d7e7882dc19fd2712ede7e352a8c8696128a8e87402b3c186.RowWithRowRequestBuilder) {
    return i61fbc6095b6a3f0d7e7882dc19fd2712ede7e352a8c8696128a8e87402b3c186.NewRowWithRowRequestBuilderInternal(m.pathParameters, m.requestAdapter, row);
}
func (m *WorkbookRangeRequestBuilder) Unmerge()(*ie1648a89e1e2322978ca232dca1a66425eb36f6d387029b0089703d2e9e217d9.UnmergeRequestBuilder) {
    return ie1648a89e1e2322978ca232dca1a66425eb36f6d387029b0089703d2e9e217d9.NewUnmergeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.usedRange()
func (m *WorkbookRangeRequestBuilder) UsedRange()(*i65783c8d11d930a960df63614ef3009276773f269fe821aadcefa2e56494f3d1.UsedRangeRequestBuilder) {
    return i65783c8d11d930a960df63614ef3009276773f269fe821aadcefa2e56494f3d1.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.usedRange(valuesOnly={valuesOnly})
// Parameters:
//  - valuesOnly : Usage: valuesOnly={valuesOnly}
func (m *WorkbookRangeRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*ia358f89504c007c9f749e41e4f15398cd1e54797872e044a95d0a784031ab907.UsedRangeWithValuesOnlyRequestBuilder) {
    return ia358f89504c007c9f749e41e4f15398cd1e54797872e044a95d0a784031ab907.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.visibleView()
func (m *WorkbookRangeRequestBuilder) VisibleView()(*i7f8a6e3fc0ef3113e8566e2989b6d447a1d04ee7e2f6fe1f2d15dd9e530c2bd1.VisibleViewRequestBuilder) {
    return i7f8a6e3fc0ef3113e8566e2989b6d447a1d04ee7e2f6fe1f2d15dd9e530c2bd1.NewVisibleViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
