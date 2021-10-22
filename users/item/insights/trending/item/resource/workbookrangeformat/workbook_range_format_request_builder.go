package workbookrangeformat

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i6943613ac6960f17a5722ec03609ded3537a3525e07ee3487cee340c42b72b73 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrangeformat/autofitrows"
    i9c1ef975a6c10fb0c286d7de5e138d5c0008cc03c17cd251be31de5efd10ce91 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/workbookrangeformat/autofitcolumns"
)

type WorkbookRangeFormatRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitColumns()(*i9c1ef975a6c10fb0c286d7de5e138d5c0008cc03c17cd251be31de5efd10ce91.AutofitColumnsRequestBuilder) {
    return i9c1ef975a6c10fb0c286d7de5e138d5c0008cc03c17cd251be31de5efd10ce91.NewAutofitColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitRows()(*i6943613ac6960f17a5722ec03609ded3537a3525e07ee3487cee340c42b72b73.AutofitRowsRequestBuilder) {
    return i6943613ac6960f17a5722ec03609ded3537a3525e07ee3487cee340c42b72b73.NewAutofitRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewWorkbookRangeFormatRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFormatRequestBuilder) {
    m := &WorkbookRangeFormatRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/trending/{trending_id}/resource/microsoft.graph.workbookRangeFormat";
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
func NewWorkbookRangeFormatRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFormatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookRangeFormatRequestBuilderInternal(urlParams, requestAdapter)
}
