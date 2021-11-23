package workbookrange

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i0ea4bcd7e635a8802e89c2dd9157a074e67b05935e907e5d3cdd2f9e3dace2a2 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/visibleview"
    i359117e4d8e74fe7d3fd0f9e4267ef1a2bbaf3acdea9e71a8f74c31db7c2f2c6 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/offsetrangewithrowoffsetwithcolumnoffset"
    i398c9197d49e978a425e6e166a439c91e943085b3c29df60fade6a53a56d3bea "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/intersectionwithanotherrange"
    i5eee546d61a8059cc8b574d94106d3118e92a659c738cf949653af82e935c9c1 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/rowsabovewithcount"
    i603b4b1621895babe2ff4605207157c6ab511e7403741acbc2a5af131d559677 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/cellwithrowwithcolumn"
    i67446cc532cd2af00c4afb8b6cc0bf4878b07667853200124a4cc344421039dd "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/rowsbelowwithcount"
    i6d440179f2e09d2eb73f72ff1e1138db4735d962c7f4472657b82bf32af6fdd0 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/usedrange"
    i704493c210cd774ad34da83cdfa405e67d5873b8142812f13bdc39eccdd00105 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/columnsbefore"
    i707592ebdaee8ad116953ddcf15d60c493cb5b13677239916db552b127077a8d "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/columnwithcolumn"
    i71efea2760fea561c43b056072a6ea7012d3a34ecc5eefdd5ab6e365471c5d37 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/resizedrangewithdeltarowswithdeltacolumns"
    i731ef5c5606a4ee848f3114cd118420aff1ab21a583a4c9eac78afee8f1335a2 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/columnsafterwithcount"
    i7cc2ada747f5265a2c92f25819a1789adffd3ed5e3a3b0c1472249d6a2f2c971 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/boundingrectwithanotherrange"
    i7dd9304ebf4108b1d9598494713fd624a9eca0bbdb300e1fabbe2cf7ed6b75ea "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/columnsbeforewithcount"
    i8e624f2b1ffdefbcbaed0a9601ad7640734f43f983ce0266006dcf541b0cc6ed "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/unmerge"
    i90af1ea1c0e5be11d83d032a403ccf7ebd8f4d6c0046f1219fd0eb18efdd8792 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/usedrangewithvaluesonly"
    i914203dec0c3eda21f2ec826e40caae5f8efaab58eed92e4fc33c7c799a998ca "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/delete"
    i9258743acdba337d7e577dc349b78db67394ffdf808a65b33621ceb3127f6109 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/lastrow"
    i98a746c25a303fbe9a3bcf24ca716f99a8c6a7af8030534352a2552b678e5020 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/lastcolumn"
    ia6b891b58dc3fb5bd19386a056d2f4b4ed6fa494742efdce5fd363b6eb6df857 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/entirerow"
    ibdf1fd776abfc0f5dd028c61dd99cfce69ad1b03ffb4e607bda4139a9be957d4 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/merge"
    ibee914a0814b89067e1173af2e01826ce6a9bccda1c0626d5fedd20bd0a15a16 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/insert"
    ic5d0a7ae11e0045cfe57b0627d0d9a4ea70afccb09590eba99dd52338f11112b "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/clear"
    id18e5819e187482dece793b951f8e7694c42c6d6d555162b3eb4e6aeb1749af5 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/rowwithrow"
    id9c4050eb8bdc8107ecffc02fea319c8b505f0bfc08b4059d41863057e60e4fa "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/columnsafter"
    if01004b447c656e7c1d02f1df17a9b01c4e1afc7f3c7d5420cd6d6a0541b0390 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/lastcell"
    ifa7383d153ce16044b2d5d96c8b48f327360bdac28227e47ae6e9ecc0036c5b5 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/rowsabove"
    ifa8ecc65656629d47271b7e8908ad3758d4f278efd1af4b253a70c26d3c3a0a3 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/rowsbelow"
    ifc399aa733120668ba9af582c106b5d4b087bc7861bb94067311c6fa09714cae "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrange/entirecolumn"
)

// WorkbookRangeRequestBuilder builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange
type WorkbookRangeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BoundingRectWithAnotherRange builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.boundingRect(anotherRange='{anotherRange}')
func (m *WorkbookRangeRequestBuilder) BoundingRectWithAnotherRange(anotherRange *string)(*i7cc2ada747f5265a2c92f25819a1789adffd3ed5e3a3b0c1472249d6a2f2c971.BoundingRectWithAnotherRangeRequestBuilder) {
    return i7cc2ada747f5265a2c92f25819a1789adffd3ed5e3a3b0c1472249d6a2f2c971.NewBoundingRectWithAnotherRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter, anotherRange);
}
// CellWithRowWithColumn builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.cell(row={row},column={column})
func (m *WorkbookRangeRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*i603b4b1621895babe2ff4605207157c6ab511e7403741acbc2a5af131d559677.CellWithRowWithColumnRequestBuilder) {
    return i603b4b1621895babe2ff4605207157c6ab511e7403741acbc2a5af131d559677.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorkbookRangeRequestBuilder) Clear()(*ic5d0a7ae11e0045cfe57b0627d0d9a4ea70afccb09590eba99dd52338f11112b.ClearRequestBuilder) {
    return ic5d0a7ae11e0045cfe57b0627d0d9a4ea70afccb09590eba99dd52338f11112b.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsAfter builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsAfter()
func (m *WorkbookRangeRequestBuilder) ColumnsAfter()(*id9c4050eb8bdc8107ecffc02fea319c8b505f0bfc08b4059d41863057e60e4fa.ColumnsAfterRequestBuilder) {
    return id9c4050eb8bdc8107ecffc02fea319c8b505f0bfc08b4059d41863057e60e4fa.NewColumnsAfterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsAfterWithCount builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsAfter(count={count})
func (m *WorkbookRangeRequestBuilder) ColumnsAfterWithCount(count *int32)(*i731ef5c5606a4ee848f3114cd118420aff1ab21a583a4c9eac78afee8f1335a2.ColumnsAfterWithCountRequestBuilder) {
    return i731ef5c5606a4ee848f3114cd118420aff1ab21a583a4c9eac78afee8f1335a2.NewColumnsAfterWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// ColumnsBefore builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsBefore()
func (m *WorkbookRangeRequestBuilder) ColumnsBefore()(*i704493c210cd774ad34da83cdfa405e67d5873b8142812f13bdc39eccdd00105.ColumnsBeforeRequestBuilder) {
    return i704493c210cd774ad34da83cdfa405e67d5873b8142812f13bdc39eccdd00105.NewColumnsBeforeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsBeforeWithCount builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.columnsBefore(count={count})
func (m *WorkbookRangeRequestBuilder) ColumnsBeforeWithCount(count *int32)(*i7dd9304ebf4108b1d9598494713fd624a9eca0bbdb300e1fabbe2cf7ed6b75ea.ColumnsBeforeWithCountRequestBuilder) {
    return i7dd9304ebf4108b1d9598494713fd624a9eca0bbdb300e1fabbe2cf7ed6b75ea.NewColumnsBeforeWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// ColumnWithColumn builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.column(column={column})
func (m *WorkbookRangeRequestBuilder) ColumnWithColumn(column *int32)(*i707592ebdaee8ad116953ddcf15d60c493cb5b13677239916db552b127077a8d.ColumnWithColumnRequestBuilder) {
    return i707592ebdaee8ad116953ddcf15d60c493cb5b13677239916db552b127077a8d.NewColumnWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, column);
}
// NewWorkbookRangeRequestBuilderInternal instantiates a new WorkbookRangeRequestBuilder and sets the default values.
func NewWorkbookRangeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeRequestBuilder) {
    m := &WorkbookRangeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/used/{usedInsight_id}/resource/microsoft.graph.workbookRange";
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
func (m *WorkbookRangeRequestBuilder) Delete()(*i914203dec0c3eda21f2ec826e40caae5f8efaab58eed92e4fc33c7c799a998ca.DeleteRequestBuilder) {
    return i914203dec0c3eda21f2ec826e40caae5f8efaab58eed92e4fc33c7c799a998ca.NewDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EntireColumn builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.entireColumn()
func (m *WorkbookRangeRequestBuilder) EntireColumn()(*ifc399aa733120668ba9af582c106b5d4b087bc7861bb94067311c6fa09714cae.EntireColumnRequestBuilder) {
    return ifc399aa733120668ba9af582c106b5d4b087bc7861bb94067311c6fa09714cae.NewEntireColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EntireRow builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.entireRow()
func (m *WorkbookRangeRequestBuilder) EntireRow()(*ia6b891b58dc3fb5bd19386a056d2f4b4ed6fa494742efdce5fd363b6eb6df857.EntireRowRequestBuilder) {
    return ia6b891b58dc3fb5bd19386a056d2f4b4ed6fa494742efdce5fd363b6eb6df857.NewEntireRowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) Insert()(*ibee914a0814b89067e1173af2e01826ce6a9bccda1c0626d5fedd20bd0a15a16.InsertRequestBuilder) {
    return ibee914a0814b89067e1173af2e01826ce6a9bccda1c0626d5fedd20bd0a15a16.NewInsertRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IntersectionWithAnotherRange builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.intersection(anotherRange='{anotherRange}')
func (m *WorkbookRangeRequestBuilder) IntersectionWithAnotherRange(anotherRange *string)(*i398c9197d49e978a425e6e166a439c91e943085b3c29df60fade6a53a56d3bea.IntersectionWithAnotherRangeRequestBuilder) {
    return i398c9197d49e978a425e6e166a439c91e943085b3c29df60fade6a53a56d3bea.NewIntersectionWithAnotherRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter, anotherRange);
}
// LastCell builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.lastCell()
func (m *WorkbookRangeRequestBuilder) LastCell()(*if01004b447c656e7c1d02f1df17a9b01c4e1afc7f3c7d5420cd6d6a0541b0390.LastCellRequestBuilder) {
    return if01004b447c656e7c1d02f1df17a9b01c4e1afc7f3c7d5420cd6d6a0541b0390.NewLastCellRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LastColumn builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.lastColumn()
func (m *WorkbookRangeRequestBuilder) LastColumn()(*i98a746c25a303fbe9a3bcf24ca716f99a8c6a7af8030534352a2552b678e5020.LastColumnRequestBuilder) {
    return i98a746c25a303fbe9a3bcf24ca716f99a8c6a7af8030534352a2552b678e5020.NewLastColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LastRow builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.lastRow()
func (m *WorkbookRangeRequestBuilder) LastRow()(*i9258743acdba337d7e577dc349b78db67394ffdf808a65b33621ceb3127f6109.LastRowRequestBuilder) {
    return i9258743acdba337d7e577dc349b78db67394ffdf808a65b33621ceb3127f6109.NewLastRowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) Merge()(*ibdf1fd776abfc0f5dd028c61dd99cfce69ad1b03ffb4e607bda4139a9be957d4.MergeRequestBuilder) {
    return ibdf1fd776abfc0f5dd028c61dd99cfce69ad1b03ffb4e607bda4139a9be957d4.NewMergeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OffsetRangeWithRowOffsetWithColumnOffset builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.offsetRange(rowOffset={rowOffset},columnOffset={columnOffset})
func (m *WorkbookRangeRequestBuilder) OffsetRangeWithRowOffsetWithColumnOffset(rowOffset *int32, columnOffset *int32)(*i359117e4d8e74fe7d3fd0f9e4267ef1a2bbaf3acdea9e71a8f74c31db7c2f2c6.OffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder) {
    return i359117e4d8e74fe7d3fd0f9e4267ef1a2bbaf3acdea9e71a8f74c31db7c2f2c6.NewOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilderInternal(m.pathParameters, m.requestAdapter, rowOffset, columnOffset);
}
// ResizedRangeWithDeltaRowsWithDeltaColumns builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.resizedRange(deltaRows={deltaRows},deltaColumns={deltaColumns})
func (m *WorkbookRangeRequestBuilder) ResizedRangeWithDeltaRowsWithDeltaColumns(deltaRows *int32, deltaColumns *int32)(*i71efea2760fea561c43b056072a6ea7012d3a34ecc5eefdd5ab6e365471c5d37.ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) {
    return i71efea2760fea561c43b056072a6ea7012d3a34ecc5eefdd5ab6e365471c5d37.NewResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter, deltaRows, deltaColumns);
}
// RowsAbove builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsAbove()
func (m *WorkbookRangeRequestBuilder) RowsAbove()(*ifa7383d153ce16044b2d5d96c8b48f327360bdac28227e47ae6e9ecc0036c5b5.RowsAboveRequestBuilder) {
    return ifa7383d153ce16044b2d5d96c8b48f327360bdac28227e47ae6e9ecc0036c5b5.NewRowsAboveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RowsAboveWithCount builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsAbove(count={count})
func (m *WorkbookRangeRequestBuilder) RowsAboveWithCount(count *int32)(*i5eee546d61a8059cc8b574d94106d3118e92a659c738cf949653af82e935c9c1.RowsAboveWithCountRequestBuilder) {
    return i5eee546d61a8059cc8b574d94106d3118e92a659c738cf949653af82e935c9c1.NewRowsAboveWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// RowsBelow builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsBelow()
func (m *WorkbookRangeRequestBuilder) RowsBelow()(*ifa8ecc65656629d47271b7e8908ad3758d4f278efd1af4b253a70c26d3c3a0a3.RowsBelowRequestBuilder) {
    return ifa8ecc65656629d47271b7e8908ad3758d4f278efd1af4b253a70c26d3c3a0a3.NewRowsBelowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RowsBelowWithCount builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.rowsBelow(count={count})
func (m *WorkbookRangeRequestBuilder) RowsBelowWithCount(count *int32)(*i67446cc532cd2af00c4afb8b6cc0bf4878b07667853200124a4cc344421039dd.RowsBelowWithCountRequestBuilder) {
    return i67446cc532cd2af00c4afb8b6cc0bf4878b07667853200124a4cc344421039dd.NewRowsBelowWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
// RowWithRow builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.row(row={row})
func (m *WorkbookRangeRequestBuilder) RowWithRow(row *int32)(*id18e5819e187482dece793b951f8e7694c42c6d6d555162b3eb4e6aeb1749af5.RowWithRowRequestBuilder) {
    return id18e5819e187482dece793b951f8e7694c42c6d6d555162b3eb4e6aeb1749af5.NewRowWithRowRequestBuilderInternal(m.pathParameters, m.requestAdapter, row);
}
func (m *WorkbookRangeRequestBuilder) Unmerge()(*i8e624f2b1ffdefbcbaed0a9601ad7640734f43f983ce0266006dcf541b0cc6ed.UnmergeRequestBuilder) {
    return i8e624f2b1ffdefbcbaed0a9601ad7640734f43f983ce0266006dcf541b0cc6ed.NewUnmergeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsedRange builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.usedRange()
func (m *WorkbookRangeRequestBuilder) UsedRange()(*i6d440179f2e09d2eb73f72ff1e1138db4735d962c7f4472657b82bf32af6fdd0.UsedRangeRequestBuilder) {
    return i6d440179f2e09d2eb73f72ff1e1138db4735d962c7f4472657b82bf32af6fdd0.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsedRangeWithValuesOnly builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.usedRange(valuesOnly={valuesOnly})
func (m *WorkbookRangeRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*i90af1ea1c0e5be11d83d032a403ccf7ebd8f4d6c0046f1219fd0eb18efdd8792.UsedRangeWithValuesOnlyRequestBuilder) {
    return i90af1ea1c0e5be11d83d032a403ccf7ebd8f4d6c0046f1219fd0eb18efdd8792.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
// VisibleView builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRange\microsoft.graph.visibleView()
func (m *WorkbookRangeRequestBuilder) VisibleView()(*i0ea4bcd7e635a8802e89c2dd9157a074e67b05935e907e5d3cdd2f9e3dace2a2.VisibleViewRequestBuilder) {
    return i0ea4bcd7e635a8802e89c2dd9157a074e67b05935e907e5d3cdd2f9e3dace2a2.NewVisibleViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
