package workbookrange

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i005c0f1e5b7636da1b81c5fb4815c470e49ce9c011bf6dd1f6b7a45e8feb037b "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/offsetrangewithrowoffsetwithcolumnoffset"
    i028bce52ac64731a93fbea7fee7604c2d6d204a8b3e1c7902ab5382d12df70ef "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/cellwithrowwithcolumn"
    i0675bf4006de8f38d581f2051bfcda596a8cd8ece0b8f0b3c44d1748c3a918d6 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/resizedrangewithdeltarowswithdeltacolumns"
    i0b1285129b5640099d4ac344a22e250eb343111f70b5370af6b36e080ad00aa9 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/visibleview"
    i0c914c0d0128d0547cbc359cc6c263e53ac2e8d8c51d9ddf638c9d1706ee3cd5 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/lastrow"
    i0d1a692b3aee49b6e69c9f5fce93622d35f98d204a0a825db39a7902063cb464 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/rowsabovewithcount"
    i1027963ad3044e24e2c3b05b05852b6bb5b0d24135b5aac4f06da3a663edc035 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/clear"
    i1034277417ccd8d7af8a852b36f875877d0f9e80870835756bd124b3e08a59b3 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/lastcell"
    i1d376d582e4c65d880add0aa304886d838b54f82dc5fe0ff0a502c2959d69932 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/columnwithcolumn"
    i255b666fd244b8db3db30fba3ebd96b57966706c0d60ebe2252c9951202aa637 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/rowwithrow"
    i2847bfbbd6454cd20ab2b6033367d62991c21b5d666bac54ef2d97f70b618137 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/lastcolumn"
    i2fb76b032c10f311cdbcc916ef296a90f295180a7ee9c0f37960c4eb5c60098d "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/unmerge"
    i3595413e03f90f7b23cd01447e06f3dfafb544bb5399098e0e1779a89116c46d "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/rowsbelow"
    i444bfb9f1e1ea1566888925ca8aef43fd33219536f9345158ecbe7a207795a66 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/rowsbelowwithcount"
    i4ec2440352592e794b584cd540d70a4f2867f7590f81786c0fc441307feadaf1 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/columnsbeforewithcount"
    i6d2c42e3dbd82afcf6c882251d192ea4f38afdd545dcb643bed37b5e756b3615 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/rowsabove"
    i75b63e47afb035b0a20bc0f8c671a1505d1daef345e07bcf4fb1524517e72e95 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/delete"
    i76586a8e4d24c42ecb2d857433efbb01063f1df154cf82919d4da30da5755ae2 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/entirerow"
    i76de5d486ac008e9df9308fa53462a42e5ef2881f43da48a3df74db83d56bfb4 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/merge"
    i77bf4d7373106dab60a0974d936b4a4f07070b726e72e8b9851540a7b6bbb594 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/columnsafter"
    i9109231e7069f278a9e3d60158c16a2f635111f9679087536064f3305d4d8635 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/columnsafterwithcount"
    i98fda923abd9a96a6e67ba6334053611cc4df4863694efe97ac7ddaa9ad58cb7 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/boundingrectwithanotherrange"
    ib18b18587e0fed18827e899c8b35502b75f0491e8053a5b77060e752c04528cd "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/columnsbefore"
    ib190a5f9eb0580cfb971a891eb73d6f54499926b2395cfe4c45a7071c544f67b "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/usedrangewithvaluesonly"
    ib7133e1768aaaa036abfc50f23092587094ce70b9df5199b513342a6a19f90e7 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/insert"
    ibfaf96acf8fbfd4177bca6e2586e8a05011bcd555e11ee6081a5d7738c75bfcf "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/usedrange"
    ic1a0affa83f6713fc56a8ed99162037b6f6fda0bd984077b03c2aff2babdd9ac "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/intersectionwithanotherrange"
    ie31775ce3f9e7268e50cc8d1d6e80b57d074c329f3a8a8f7e7bfe35a641b4b42 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrange/entirecolumn"
)

type WorkbookRangeRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *WorkbookRangeRequestBuilder) BoundingRectWithAnotherRange(anotherRange *string)(*i98fda923abd9a96a6e67ba6334053611cc4df4863694efe97ac7ddaa9ad58cb7.BoundingRectWithAnotherRangeRequestBuilder) {
    return i98fda923abd9a96a6e67ba6334053611cc4df4863694efe97ac7ddaa9ad58cb7.NewBoundingRectWithAnotherRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter, anotherRange);
}
func (m *WorkbookRangeRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*i028bce52ac64731a93fbea7fee7604c2d6d204a8b3e1c7902ab5382d12df70ef.CellWithRowWithColumnRequestBuilder) {
    return i028bce52ac64731a93fbea7fee7604c2d6d204a8b3e1c7902ab5382d12df70ef.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorkbookRangeRequestBuilder) Clear()(*i1027963ad3044e24e2c3b05b05852b6bb5b0d24135b5aac4f06da3a663edc035.ClearRequestBuilder) {
    return i1027963ad3044e24e2c3b05b05852b6bb5b0d24135b5aac4f06da3a663edc035.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) ColumnsAfter()(*i77bf4d7373106dab60a0974d936b4a4f07070b726e72e8b9851540a7b6bbb594.ColumnsAfterRequestBuilder) {
    return i77bf4d7373106dab60a0974d936b4a4f07070b726e72e8b9851540a7b6bbb594.NewColumnsAfterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) ColumnsAfterWithCount(count *int32)(*i9109231e7069f278a9e3d60158c16a2f635111f9679087536064f3305d4d8635.ColumnsAfterWithCountRequestBuilder) {
    return i9109231e7069f278a9e3d60158c16a2f635111f9679087536064f3305d4d8635.NewColumnsAfterWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
func (m *WorkbookRangeRequestBuilder) ColumnsBefore()(*ib18b18587e0fed18827e899c8b35502b75f0491e8053a5b77060e752c04528cd.ColumnsBeforeRequestBuilder) {
    return ib18b18587e0fed18827e899c8b35502b75f0491e8053a5b77060e752c04528cd.NewColumnsBeforeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) ColumnsBeforeWithCount(count *int32)(*i4ec2440352592e794b584cd540d70a4f2867f7590f81786c0fc441307feadaf1.ColumnsBeforeWithCountRequestBuilder) {
    return i4ec2440352592e794b584cd540d70a4f2867f7590f81786c0fc441307feadaf1.NewColumnsBeforeWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
func (m *WorkbookRangeRequestBuilder) ColumnWithColumn(column *int32)(*i1d376d582e4c65d880add0aa304886d838b54f82dc5fe0ff0a502c2959d69932.ColumnWithColumnRequestBuilder) {
    return i1d376d582e4c65d880add0aa304886d838b54f82dc5fe0ff0a502c2959d69932.NewColumnWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, column);
}
func NewWorkbookRangeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeRequestBuilder) {
    m := &WorkbookRangeRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/shared/{sharedInsight_id}/resource/microsoft.graph.workbookRange";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewWorkbookRangeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookRangeRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorkbookRangeRequestBuilder) Delete()(*i75b63e47afb035b0a20bc0f8c671a1505d1daef345e07bcf4fb1524517e72e95.DeleteRequestBuilder) {
    return i75b63e47afb035b0a20bc0f8c671a1505d1daef345e07bcf4fb1524517e72e95.NewDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) EntireColumn()(*ie31775ce3f9e7268e50cc8d1d6e80b57d074c329f3a8a8f7e7bfe35a641b4b42.EntireColumnRequestBuilder) {
    return ie31775ce3f9e7268e50cc8d1d6e80b57d074c329f3a8a8f7e7bfe35a641b4b42.NewEntireColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) EntireRow()(*i76586a8e4d24c42ecb2d857433efbb01063f1df154cf82919d4da30da5755ae2.EntireRowRequestBuilder) {
    return i76586a8e4d24c42ecb2d857433efbb01063f1df154cf82919d4da30da5755ae2.NewEntireRowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) Insert()(*ib7133e1768aaaa036abfc50f23092587094ce70b9df5199b513342a6a19f90e7.InsertRequestBuilder) {
    return ib7133e1768aaaa036abfc50f23092587094ce70b9df5199b513342a6a19f90e7.NewInsertRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) IntersectionWithAnotherRange(anotherRange *string)(*ic1a0affa83f6713fc56a8ed99162037b6f6fda0bd984077b03c2aff2babdd9ac.IntersectionWithAnotherRangeRequestBuilder) {
    return ic1a0affa83f6713fc56a8ed99162037b6f6fda0bd984077b03c2aff2babdd9ac.NewIntersectionWithAnotherRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter, anotherRange);
}
func (m *WorkbookRangeRequestBuilder) LastCell()(*i1034277417ccd8d7af8a852b36f875877d0f9e80870835756bd124b3e08a59b3.LastCellRequestBuilder) {
    return i1034277417ccd8d7af8a852b36f875877d0f9e80870835756bd124b3e08a59b3.NewLastCellRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) LastColumn()(*i2847bfbbd6454cd20ab2b6033367d62991c21b5d666bac54ef2d97f70b618137.LastColumnRequestBuilder) {
    return i2847bfbbd6454cd20ab2b6033367d62991c21b5d666bac54ef2d97f70b618137.NewLastColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) LastRow()(*i0c914c0d0128d0547cbc359cc6c263e53ac2e8d8c51d9ddf638c9d1706ee3cd5.LastRowRequestBuilder) {
    return i0c914c0d0128d0547cbc359cc6c263e53ac2e8d8c51d9ddf638c9d1706ee3cd5.NewLastRowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) Merge()(*i76de5d486ac008e9df9308fa53462a42e5ef2881f43da48a3df74db83d56bfb4.MergeRequestBuilder) {
    return i76de5d486ac008e9df9308fa53462a42e5ef2881f43da48a3df74db83d56bfb4.NewMergeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) OffsetRangeWithRowOffsetWithColumnOffset(rowOffset *int32, columnOffset *int32)(*i005c0f1e5b7636da1b81c5fb4815c470e49ce9c011bf6dd1f6b7a45e8feb037b.OffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder) {
    return i005c0f1e5b7636da1b81c5fb4815c470e49ce9c011bf6dd1f6b7a45e8feb037b.NewOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilderInternal(m.pathParameters, m.requestAdapter, rowOffset, columnOffset);
}
func (m *WorkbookRangeRequestBuilder) ResizedRangeWithDeltaRowsWithDeltaColumns(deltaRows *int32, deltaColumns *int32)(*i0675bf4006de8f38d581f2051bfcda596a8cd8ece0b8f0b3c44d1748c3a918d6.ResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) {
    return i0675bf4006de8f38d581f2051bfcda596a8cd8ece0b8f0b3c44d1748c3a918d6.NewResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter, deltaRows, deltaColumns);
}
func (m *WorkbookRangeRequestBuilder) RowsAbove()(*i6d2c42e3dbd82afcf6c882251d192ea4f38afdd545dcb643bed37b5e756b3615.RowsAboveRequestBuilder) {
    return i6d2c42e3dbd82afcf6c882251d192ea4f38afdd545dcb643bed37b5e756b3615.NewRowsAboveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) RowsAboveWithCount(count *int32)(*i0d1a692b3aee49b6e69c9f5fce93622d35f98d204a0a825db39a7902063cb464.RowsAboveWithCountRequestBuilder) {
    return i0d1a692b3aee49b6e69c9f5fce93622d35f98d204a0a825db39a7902063cb464.NewRowsAboveWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
func (m *WorkbookRangeRequestBuilder) RowsBelow()(*i3595413e03f90f7b23cd01447e06f3dfafb544bb5399098e0e1779a89116c46d.RowsBelowRequestBuilder) {
    return i3595413e03f90f7b23cd01447e06f3dfafb544bb5399098e0e1779a89116c46d.NewRowsBelowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) RowsBelowWithCount(count *int32)(*i444bfb9f1e1ea1566888925ca8aef43fd33219536f9345158ecbe7a207795a66.RowsBelowWithCountRequestBuilder) {
    return i444bfb9f1e1ea1566888925ca8aef43fd33219536f9345158ecbe7a207795a66.NewRowsBelowWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count);
}
func (m *WorkbookRangeRequestBuilder) RowWithRow(row *int32)(*i255b666fd244b8db3db30fba3ebd96b57966706c0d60ebe2252c9951202aa637.RowWithRowRequestBuilder) {
    return i255b666fd244b8db3db30fba3ebd96b57966706c0d60ebe2252c9951202aa637.NewRowWithRowRequestBuilderInternal(m.pathParameters, m.requestAdapter, row);
}
func (m *WorkbookRangeRequestBuilder) Unmerge()(*i2fb76b032c10f311cdbcc916ef296a90f295180a7ee9c0f37960c4eb5c60098d.UnmergeRequestBuilder) {
    return i2fb76b032c10f311cdbcc916ef296a90f295180a7ee9c0f37960c4eb5c60098d.NewUnmergeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) UsedRange()(*ibfaf96acf8fbfd4177bca6e2586e8a05011bcd555e11ee6081a5d7738c75bfcf.UsedRangeRequestBuilder) {
    return ibfaf96acf8fbfd4177bca6e2586e8a05011bcd555e11ee6081a5d7738c75bfcf.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*ib190a5f9eb0580cfb971a891eb73d6f54499926b2395cfe4c45a7071c544f67b.UsedRangeWithValuesOnlyRequestBuilder) {
    return ib190a5f9eb0580cfb971a891eb73d6f54499926b2395cfe4c45a7071c544f67b.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}
func (m *WorkbookRangeRequestBuilder) VisibleView()(*i0b1285129b5640099d4ac344a22e250eb343111f70b5370af6b36e080ad00aa9.VisibleViewRequestBuilder) {
    return i0b1285129b5640099d4ac344a22e250eb343111f70b5370af6b36e080ad00aa9.NewVisibleViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
