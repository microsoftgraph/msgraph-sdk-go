package workbook

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0903ad19dd9574dfbb4bed167da14643c2228c1dea871561fb358665afbb3ca6 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names"
    i1eca07bc58ad01f3faa752b735e09e50ab7683422b4e83226505d100e95ee50c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets"
    i1fe31042b8e2ca156ea1194a9d9ff819cefef5e881b2ad8b7e157690bc2cde4d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/application"
    i4b302c1274ce37d0c4ddda1cc08f3afdf0512a0fbf76035100627221a347cf0a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/closesession"
    i57a806e6c9fd10c20a29800182ad7f2f5148df028028093cef93c98220b8169b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tablerowoperationresultwithkey"
    i5f803d3a424a23fce972e96641273fa544aa01981cf2d99fda7e49efa665af3f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions"
    icc3ddbd086b90086e0cc449c4727a2d955b5fa6e180fc36ba2ff568fd7726c93 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/operations"
    idf6798b505adccc099c94c60811b9d6c499a7568c1af8d4aa9244555e95d4904 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables"
    ie05bda438f4ce09e0d311b6ea61d5242f2a8a003f28f8c6e94cc5c1668b41fbf "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/sessioninforesourcewithkey"
    ie2cc01c1583220350b936e632410ebd1c5d0ccd0cab73c17650c9bd54bde96be "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/createsession"
    ied7e07d513586a96205d8ae9f0876cae0217399e1898581f1a059df4292376dc "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/comments"
    iee77dcfa41a1d1e752c266cbc0f329cab7791c1dc70b104aeda598ce43b190cc "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/refreshsession"
    i867aea5d1d9d669532c2772fe62859a2f33577fdfb784d5e4dfb9e0ee0a76e43 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item"
    ia41154e3202e47bd63acff7870d9d6636dc792d816181e05a5856d39b83bec9f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/comments/item"
    ia867d3dcb91a7bd46ca5aaeb6bb772aca9c6f09cef40078b4d04158f64420354 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/tables/item"
    ib4734e530e4717b1ad6655524718a151560482e22b338d1d804a63a8a2052198 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/operations/item"
    icab31dc579f88cafdb4bcc1f326af353da1fc9fc046edacfaab9d24362f5130c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item"
)

// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook
type WorkbookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type WorkbookRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type WorkbookRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WorkbookRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// For files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
type WorkbookRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
}
// Options for Patch
type WorkbookRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Workbook;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *WorkbookRequestBuilder) Application()(*i1fe31042b8e2ca156ea1194a9d9ff819cefef5e881b2ad8b7e157690bc2cde4d.ApplicationRequestBuilder) {
    return i1fe31042b8e2ca156ea1194a9d9ff819cefef5e881b2ad8b7e157690bc2cde4d.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) CloseSession()(*i4b302c1274ce37d0c4ddda1cc08f3afdf0512a0fbf76035100627221a347cf0a.CloseSessionRequestBuilder) {
    return i4b302c1274ce37d0c4ddda1cc08f3afdf0512a0fbf76035100627221a347cf0a.NewCloseSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) Comments()(*ied7e07d513586a96205d8ae9f0876cae0217399e1898581f1a059df4292376dc.CommentsRequestBuilder) {
    return ied7e07d513586a96205d8ae9f0876cae0217399e1898581f1a059df4292376dc.NewCommentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.comments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WorkbookRequestBuilder) CommentsById(id string)(*ia41154e3202e47bd63acff7870d9d6636dc792d816181e05a5856d39b83bec9f.WorkbookCommentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookComment_id"] = id
    }
    return ia41154e3202e47bd63acff7870d9d6636dc792d816181e05a5856d39b83bec9f.NewWorkbookCommentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new WorkbookRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRequestBuilder) {
    m := &WorkbookRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook{?expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new WorkbookRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookRequestBuilderInternal(urlParams, requestAdapter)
}
// For files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
// Parameters:
//  - options : Options for the request
func (m *WorkbookRequestBuilder) CreateDeleteRequestInformation(options *WorkbookRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// For files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
// Parameters:
//  - options : Options for the request
func (m *WorkbookRequestBuilder) CreateGetRequestInformation(options *WorkbookRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(options.Q)
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// For files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
// Parameters:
//  - options : Options for the request
func (m *WorkbookRequestBuilder) CreatePatchRequestInformation(options *WorkbookRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *WorkbookRequestBuilder) CreateSession()(*ie2cc01c1583220350b936e632410ebd1c5d0ccd0cab73c17650c9bd54bde96be.CreateSessionRequestBuilder) {
    return ie2cc01c1583220350b936e632410ebd1c5d0ccd0cab73c17650c9bd54bde96be.NewCreateSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// For files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
// Parameters:
//  - options : Options for the request
func (m *WorkbookRequestBuilder) Delete(options *WorkbookRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookRequestBuilder) Functions()(*i5f803d3a424a23fce972e96641273fa544aa01981cf2d99fda7e49efa665af3f.FunctionsRequestBuilder) {
    return i5f803d3a424a23fce972e96641273fa544aa01981cf2d99fda7e49efa665af3f.NewFunctionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// For files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
// Parameters:
//  - options : Options for the request
func (m *WorkbookRequestBuilder) Get(options *WorkbookRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Workbook, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbook() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Workbook), nil
}
func (m *WorkbookRequestBuilder) Names()(*i0903ad19dd9574dfbb4bed167da14643c2228c1dea871561fb358665afbb3ca6.NamesRequestBuilder) {
    return i0903ad19dd9574dfbb4bed167da14643c2228c1dea871561fb358665afbb3ca6.NewNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.names.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WorkbookRequestBuilder) NamesById(id string)(*i867aea5d1d9d669532c2772fe62859a2f33577fdfb784d5e4dfb9e0ee0a76e43.WorkbookNamedItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookNamedItem_id"] = id
    }
    return i867aea5d1d9d669532c2772fe62859a2f33577fdfb784d5e4dfb9e0ee0a76e43.NewWorkbookNamedItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) Operations()(*icc3ddbd086b90086e0cc449c4727a2d955b5fa6e180fc36ba2ff568fd7726c93.OperationsRequestBuilder) {
    return icc3ddbd086b90086e0cc449c4727a2d955b5fa6e180fc36ba2ff568fd7726c93.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.operations.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WorkbookRequestBuilder) OperationsById(id string)(*ib4734e530e4717b1ad6655524718a151560482e22b338d1d804a63a8a2052198.WorkbookOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookOperation_id"] = id
    }
    return ib4734e530e4717b1ad6655524718a151560482e22b338d1d804a63a8a2052198.NewWorkbookOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// For files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
// Parameters:
//  - options : Options for the request
func (m *WorkbookRequestBuilder) Patch(options *WorkbookRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookRequestBuilder) RefreshSession()(*iee77dcfa41a1d1e752c266cbc0f329cab7791c1dc70b104aeda598ce43b190cc.RefreshSessionRequestBuilder) {
    return iee77dcfa41a1d1e752c266cbc0f329cab7791c1dc70b104aeda598ce43b190cc.NewRefreshSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\microsoft.graph.sessionInfoResource(key='{key}')
// Parameters:
//  - key : Usage: key={key}
func (m *WorkbookRequestBuilder) SessionInfoResourceWithKey(key *string)(*ie05bda438f4ce09e0d311b6ea61d5242f2a8a003f28f8c6e94cc5c1668b41fbf.SessionInfoResourceWithKeyRequestBuilder) {
    return ie05bda438f4ce09e0d311b6ea61d5242f2a8a003f28f8c6e94cc5c1668b41fbf.NewSessionInfoResourceWithKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, key);
}
// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\microsoft.graph.tableRowOperationResult(key='{key}')
// Parameters:
//  - key : Usage: key={key}
func (m *WorkbookRequestBuilder) TableRowOperationResultWithKey(key *string)(*i57a806e6c9fd10c20a29800182ad7f2f5148df028028093cef93c98220b8169b.TableRowOperationResultWithKeyRequestBuilder) {
    return i57a806e6c9fd10c20a29800182ad7f2f5148df028028093cef93c98220b8169b.NewTableRowOperationResultWithKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, key);
}
func (m *WorkbookRequestBuilder) Tables()(*idf6798b505adccc099c94c60811b9d6c499a7568c1af8d4aa9244555e95d4904.TablesRequestBuilder) {
    return idf6798b505adccc099c94c60811b9d6c499a7568c1af8d4aa9244555e95d4904.NewTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.tables.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WorkbookRequestBuilder) TablesById(id string)(*ia867d3dcb91a7bd46ca5aaeb6bb772aca9c6f09cef40078b4d04158f64420354.WorkbookTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTable_id"] = id
    }
    return ia867d3dcb91a7bd46ca5aaeb6bb772aca9c6f09cef40078b4d04158f64420354.NewWorkbookTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookRequestBuilder) Worksheets()(*i1eca07bc58ad01f3faa752b735e09e50ab7683422b4e83226505d100e95ee50c.WorksheetsRequestBuilder) {
    return i1eca07bc58ad01f3faa752b735e09e50ab7683422b4e83226505d100e95ee50c.NewWorksheetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.workbooks.item.workbook.worksheets.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WorkbookRequestBuilder) WorksheetsById(id string)(*icab31dc579f88cafdb4bcc1f326af353da1fc9fc046edacfaab9d24362f5130c.WorkbookWorksheetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookWorksheet_id"] = id
    }
    return icab31dc579f88cafdb4bcc1f326af353da1fc9fc046edacfaab9d24362f5130c.NewWorkbookWorksheetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
