package workbookrange

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i11b3f41350403e58533da52aa916c72c2763c27d57b3069a5fdef1d4a83372bd "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/clear"
    i279549f5e540a832cf1eccf4ef475fb9f4035660da1304afe4c9885f372ad658 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/rowsbelowwithcount"
    i2df401f0c00c83e9e9b6ad23d7fa5e26df86d1fdd7d9296f0418e48524c46150 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/intersectionwithanotherrange"
    i44b2fc492eb4dac248440e05cc96b42917abd5df39c3822184d82c91b4a3e800 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/columnsafterwithcount"
    i4f687e53d3703ee2dfa44a1e779a0695ffb746ff9bf062b1fe27a775ee3263e2 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/resizedrangewithdeltarowswithdeltacolumns"
    i51faca6105458c474c6407b39115213a62d341fa17bfba8561ddc2ff948eb3a5 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/lastcell"
    i5d97a0aab867175aaf75cbaa83fabb2e817ba4ab35a5d1a946b01213df0a3ea5 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/visibleview"
    i60ae4f0a9e392d4be1eddfcb82606a754af177eb7a114b5cd6d24da9c539bbd3 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/lastrow"
    i6ac964da8767a7dd691a4f69a99359ca56df372f8a7369034b777a810924c56a "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/lastcolumn"
    i6d051e2fd6ef4154594cef9a6cf2c7dee2f13963b4162f16d0caa473c1123669 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/offsetrangewithrowoffsetwithcolumnoffset"
    i6da1675818caf0c45320b1138556ee1a6fe70bb2238d4c467c788a980f3e83d7 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/columnwithcolumn"
    i6f3b2875acb6d9811b7227b9681fa90acdd527f8eb94ecff5f33f5eb61204e18 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/usedrangewithvaluesonly"
    i7bdd35761cad8f0b0bf7087b8c6a94b3dfe6293c02ce03a6267858eb995340fc "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/merge"
    i7e3a593cea6e23a1e383fdc2ebc388743cbb927a22bfefa0d4b07ae6dc7b5dfc "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/columnsbeforewithcount"
    i9a432892245018986eb63ef9d659205d85d7f3757b5d78549a945d4b3d3e8379 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/rowsabovewithcount"
    ia275ea063bee37f39e7df0835ca4c8c883229be21ea645b1ffe83650b7d9490b "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/insert"
    ia9323848fc8e5204ebdece64e035f4d170ff9dd59d80a85a65ff7ba901fd6e36 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/columnsafter"
    iaf472a43425018f9bbf693385da259be73c361e87dc2d99f002154b38e13edeb "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/entirecolumn"
    ib451e723175b07a08b3e964d89045cae155c873ccb855a65643a999269f7ca25 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/cellwithrowwithcolumn"
    ic3440a6d512a65eab00dac6b16eb22033fb5645646cac9ddb733cb52db651a35 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/rowsbelow"
    ic6b2ecb836907df8d01e3c2f842b1e91a7ff29b811f2030cf6f91dbb8fe11b53 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/rowsabove"
    ica0254825ac143f23d6d8617e4da9fc5077329c5bd9ed09267bc92f72515056e "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/columnsbefore"
    ie8db66c4e0024a1b0f870aa4088e5ad436b1b3b66559e5a1cece6bd720e65ba3 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/rowwithrow"
    iebe090f4e1ebf62503c6c47ad1f8f8cf461471d424872b5629cf0898ffb60d2c "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/delete"
    iecb3289dc8cf4b6962b981a2a4c1e675eaa0ac7df5e2ca1a7625e60e1cbc56ea "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/usedrange"
    if4a9f771cfe6f18065826fae1a47e2023bfa210d0bfedfdc09aa4e145c08d1b7 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/boundingrectwithanotherrange"
    if90063101947a0372e1a95a9e094ed6e71a418e0fa8cc47f07d5096b57e21e2f "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/unmerge"
    ifde463e810689c7072c7d9277a46137bfa317a65065935fd36040a8fa123d142 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrange/entirerow"
)

// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange
type WorkbookRangeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.boundingRect(anotherRange='{anotherRange}')
// Parameters:
//  - anotherRange : Usage: anotherRange={anotherRange}
func (m *WorkbookRangeRequestBuilder) BoundingRectWithAnotherRange(anotherRange *string)(*if4a9f771cfe6f18065826fae1a47e2023bfa210d0bfedfdc09aa4e145c08d1b7.BoundingRectWithAnotherRangeRequestBuilder) {
    return if4a9f771cfe6f18065826fae1a47e2023bfa210d0bfedfdc09aa4e145c08d1b7.NewBoundingRectWithAnotherRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter, anotherRange);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.cell(row={row},column={column})
// Parameters:
//  - column : Usage: column={column}
//  - row : Usage: row={row}
func (m *WorkbookRangeRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*ib451e723175b07a08b3e964d89045cae155c873ccb855a65643a999269f7ca25.CellWithRowWithColumnRequestBuilder) {
    return ib451e723175b07a08b3e964d89045cae155c873ccb855a65643a999269f7ca25.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorkbookRangeRequestBuilder) Clear()(*i11b3f41350403e58533da52aa916c72c2763c27d57b3069a5fdef1d4a83372bd.ClearRequestBuilder) {
    return i11b3f41350403e58533da52aa916c72c2763c27d57b3069a5fdef1d4a83372bd.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsAfter()
func (m *WorkbookRangeRequestBuilder) ColumnsAfter()(*ia9323848fc8e5204ebdece64e035f4d170ff9dd59d80a85a65ff7ba901fd6e36.ColumnsAfterRequestBuilder) {
    return ia9323848fc8e5204ebdece64e035f4d170ff9dd59d80a85a65ff7ba901fd6e36.NewColumnsAfterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsAfter(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) ColumnsAfterWithCount(count *int32)(*i44b2fc492eb4dac248440e05cc96b42917abd5df39c3822184d82c91b4a3e800.ColumnsAfterWithCountRequestBuilder) {
    return i44b2fc492eb4dac248440e05cc96b42917abd5df39c3822184d82c91b4a3e800.NewColumnsAfterWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsBefore()
func (m *WorkbookRangeRequestBuilder) ColumnsBefore()(*ica0254825ac143f23d6d8617e4da9fc5077329c5bd9ed09267bc92f72515056e.ColumnsBeforeRequestBuilder) {
    return ica0254825ac143f23d6d8617e4da9fc5077329c5bd9ed09267bc92f72515056e.NewColumnsBeforeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsBefore(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) ColumnsBeforeWithCount(count *int32)(*i7e3a593cea6e23a1e383fdc2ebc388743cbb927a22bfefa0d4b07ae6dc7b5dfc.ColumnsBeforeWithCountRequestBuilder) {
    return i7e3a593cea6e23a1e383fdc2ebc388743cbb927a22bfefa0d4b07ae6dc7b5dfc.NewColumnsBeforeWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.column(column={column})
// Parameters:
//  - column : Usage: column={column}
func (m *WorkbookRangeRequestBuilder) ColumnWithColumn(column *int32)(*i6da1675818caf0c45320b1138556ee1a6fe70bb2238d4c467c788a980f3e83d7.ColumnWithColumnRequestBuilder) {
    return i6da1675818caf0c45320b1138556ee1a6fe70bb2238d4c467c788a980f3e83d7.NewColumnWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, column);
}
// Instantiates a new WorkbookRangeRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookRangeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeRequestBuilder) {
    m := &WorkbookRangeRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/trending/{trending_id}/resource/microsoft.graph.workbookRange";
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
func (m *WorkbookRangeRequestBuilder) Delete()(*iebe090f4e1ebf62503c6c47ad1f8f8cf461471d424872b5629cf0898ffb60d2c.DeleteRequestBuilder) {
    return iebe090f4e1ebf62503c6c47ad1f8f8cf461471d424872b5629cf0898ffb60d2c.NewDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.entireColumn()
func (m *WorkbookRangeRequestBuilder) EntireColumn()(*iaf472a43425018f9bbf693385da259be73c361e87dc2d99f002154b38e13edeb.EntireColumnRequestBuilder) {
    return iaf472a43425018f9bbf693385da259be73c361e87dc2d99f002154b38e13edeb.NewEntireColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.entireRow()
func (m *WorkbookRangeRequestBuilder) EntireRow()(*ifde463e810689c7072c7d9277a46137bfa317a65065935fd36040a8fa123d142.EntireRowRequestBuilder) {
    return ifde463e810689c7072c7d9277a46137bfa317a65065935fd36040a8fa123d142.NewEntireRowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) Insert()(*ia275ea063bee37f39e7df0835ca4c8c883229be21ea645b1ffe83650b7d9490b.InsertRequestBuilder) {
    return ia275ea063bee37f39e7df0835ca4c8c883229be21ea645b1ffe83650b7d9490b.NewInsertRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.intersection(anotherRange='{anotherRange}')
// Parameters:
//  - anotherRange : Usage: anotherRange={anotherRange}
func (m *WorkbookRangeRequestBuilder) IntersectionWithAnotherRange(anotherRange *string)(*i2df401f0c00c83e9e9b6ad23d7fa5e26df86d1fdd7d9296f0418e48524c46150.IntersectionWithAnotherRangeRequestBuilder) {
    return i2df401f0c00c83e9e9b6ad23d7fa5e26df86d1fdd7d9296f0418e48524c46150.NewIntersectionWithAnotherRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter, anotherRange);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.lastCell()
func (m *WorkbookRangeRequestBuilder) LastCell()(*i51faca6105458c474c6407b39115213a62d341fa17bfba8561ddc2ff948eb3a5.LastCellRequestBuilder) {
    return i51faca6105458c474c6407b39115213a62d341fa17bfba8561ddc2ff948eb3a5.NewLastCellRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.lastColumn()
func (m *WorkbookRangeRequestBuilder) LastColumn()(*i6ac964da8767a7dd691a4f69a99359ca56df372f8a7369034b777a810924c56a.LastColumnRequestBuilder) {
    return i6ac964da8767a7dd691a4f69a99359ca56df372f8a7369034b777a810924c56a.NewLastColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.lastRow()
func (m *WorkbookRangeRequestBuilder) LastRow()(*i60ae4f0a9e392d4be1eddfcb82606a754af177eb7a114b5cd6d24da9c539bbd3.LastRowRequestBuilder) {
    return i60ae4f0a9e392d4be1eddfcb82606a754af177eb7a114b5cd6d24da9c539bbd3.NewLastRowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) Merge()(*i7bdd35761cad8f0b0bf7087b8c6a94b3dfe6293c02ce03a6267858eb995340fc.MergeRequestBuilder) {
    return i7bdd35761cad8f0b0bf7087b8c6a94b3dfe6293c02ce03a6267858eb995340fc.NewMergeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.offsetRange(rowOffset={rowOffset},columnOffset={columnOffset})
// Parameters:
//  - columnOffset : Usage: columnOffset={columnOffset}
//  - rowOffset : Usage: rowOffset={rowOffset}
func (m *WorkbookRangeRequestBuilder) OffsetRangeWithRowOffsetWithColumnOffset(rowOffset *int32, columnOffset *int32)(*i6d051e2fd6ef4154594cef9a6cf2c7dee2f13963b4162f16d0caa473c1123669.OffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder) {
    return i6d051e2fd6ef4154594cef9a6cf2c7dee2f13963b4162f16d0caa473c1123669.NewOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilderInternal(m.pathParameters, m.requestAdapter, rowOffset, columnOffset);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.resizedRange(deltaRows={deltaRows},deltaColumns={deltaColumns})
// Parameters:
//  - deltaColumns : Usage: deltaColumns={deltaColumns}
//  - deltaRows : Usage: deltaRows={deltaRows}
func (m *WorkbookRangeRequestBuilder) ResizedRangeWithDeltaRowsWithDeltaColumns(deltaRows *int32, deltaColumns *int32)(*i4f687e53d3703ee2dfa44a1e779a0695ffb746ff9bf062b1fe27a775ee3263e2.ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) {
    return i4f687e53d3703ee2dfa44a1e779a0695ffb746ff9bf062b1fe27a775ee3263e2.NewResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter, deltaRows, deltaColumns);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsAbove()
func (m *WorkbookRangeRequestBuilder) RowsAbove()(*ic6b2ecb836907df8d01e3c2f842b1e91a7ff29b811f2030cf6f91dbb8fe11b53.RowsAboveRequestBuilder) {
    return ic6b2ecb836907df8d01e3c2f842b1e91a7ff29b811f2030cf6f91dbb8fe11b53.NewRowsAboveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsAbove(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) RowsAboveWithCount(count *int32)(*i9a432892245018986eb63ef9d659205d85d7f3757b5d78549a945d4b3d3e8379.RowsAboveWithCountRequestBuilder) {
    return i9a432892245018986eb63ef9d659205d85d7f3757b5d78549a945d4b3d3e8379.NewRowsAboveWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsBelow()
func (m *WorkbookRangeRequestBuilder) RowsBelow()(*ic3440a6d512a65eab00dac6b16eb22033fb5645646cac9ddb733cb52db651a35.RowsBelowRequestBuilder) {
    return ic3440a6d512a65eab00dac6b16eb22033fb5645646cac9ddb733cb52db651a35.NewRowsBelowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsBelow(count={count})
// Parameters:
//  - count : Usage: count={count}
func (m *WorkbookRangeRequestBuilder) RowsBelowWithCount(count *int32)(*i279549f5e540a832cf1eccf4ef475fb9f4035660da1304afe4c9885f372ad658.RowsBelowWithCountRequestBuilder) {
    return i279549f5e540a832cf1eccf4ef475fb9f4035660da1304afe4c9885f372ad658.NewRowsBelowWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.row(row={row})
// Parameters:
//  - row : Usage: row={row}
func (m *WorkbookRangeRequestBuilder) RowWithRow(row *int32)(*ie8db66c4e0024a1b0f870aa4088e5ad436b1b3b66559e5a1cece6bd720e65ba3.RowWithRowRequestBuilder) {
    return ie8db66c4e0024a1b0f870aa4088e5ad436b1b3b66559e5a1cece6bd720e65ba3.NewRowWithRowRequestBuilderInternal(m.pathParameters, m.requestAdapter, row);
}
func (m *WorkbookRangeRequestBuilder) Unmerge()(*if90063101947a0372e1a95a9e094ed6e71a418e0fa8cc47f07d5096b57e21e2f.UnmergeRequestBuilder) {
    return if90063101947a0372e1a95a9e094ed6e71a418e0fa8cc47f07d5096b57e21e2f.NewUnmergeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.usedRange()
func (m *WorkbookRangeRequestBuilder) UsedRange()(*iecb3289dc8cf4b6962b981a2a4c1e675eaa0ac7df5e2ca1a7625e60e1cbc56ea.UsedRangeRequestBuilder) {
    return iecb3289dc8cf4b6962b981a2a4c1e675eaa0ac7df5e2ca1a7625e60e1cbc56ea.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.usedRange(valuesOnly={valuesOnly})
// Parameters:
//  - valuesOnly : Usage: valuesOnly={valuesOnly}
func (m *WorkbookRangeRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*i6f3b2875acb6d9811b7227b9681fa90acdd527f8eb94ecff5f33f5eb61204e18.UsedRangeWithValuesOnlyRequestBuilder) {
    return i6f3b2875acb6d9811b7227b9681fa90acdd527f8eb94ecff5f33f5eb61204e18.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRange\microsoft.graph.visibleView()
func (m *WorkbookRangeRequestBuilder) VisibleView()(*i5d97a0aab867175aaf75cbaa83fabb2e817ba4ab35a5d1a946b01213df0a3ea5.VisibleViewRequestBuilder) {
    return i5d97a0aab867175aaf75cbaa83fabb2e817ba4ab35a5d1a946b01213df0a3ea5.NewVisibleViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
