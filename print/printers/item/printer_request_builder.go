package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i05f5154d5dc8e1096ef3ba18fb5bdb890712250391686b0ffd44f043e05db1fb "github.com/microsoftgraph/msgraph-sdk-go/print/printers/item/tasktriggers"
    i6ef3a73ca24cdfe4f01546d744e3ad5d5f17082ec1283caf5b2176d392cb9cd9 "github.com/microsoftgraph/msgraph-sdk-go/print/printers/item/connectors"
    i74b470ddc01590a04dfc5ba245b7ccb094a2f7b819d0a4ea94773869da1de8bc "github.com/microsoftgraph/msgraph-sdk-go/print/printers/item/shares"
    i95e425711e2e59b85807bfbd36bd18a7893a62957fccbb0dcc17647c9c3f5e90 "github.com/microsoftgraph/msgraph-sdk-go/print/printers/item/restorefactorydefaults"
    i8a35ad54635f07016096adf7157efc92ede0057cd4a790d06a32332ce94f22cc "github.com/microsoftgraph/msgraph-sdk-go/print/printers/item/tasktriggers/item"
)

// PrinterRequestBuilder builds and executes requests for operations under \print\printers\{printer-id}
type PrinterRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrinterRequestBuilderDeleteOptions options for Delete
type PrinterRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrinterRequestBuilderGetOptions options for Get
type PrinterRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrinterRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrinterRequestBuilderGetQueryParameters the list of printers registered in the tenant.
type PrinterRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PrinterRequestBuilderPatchOptions options for Patch
type PrinterRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Printer;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PrinterRequestBuilder) Connectors()(*i6ef3a73ca24cdfe4f01546d744e3ad5d5f17082ec1283caf5b2176d392cb9cd9.ConnectorsRequestBuilder) {
    return i6ef3a73ca24cdfe4f01546d744e3ad5d5f17082ec1283caf5b2176d392cb9cd9.NewConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrinterRequestBuilderInternal instantiates a new PrinterRequestBuilder and sets the default values.
func NewPrinterRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrinterRequestBuilder) {
    m := &PrinterRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/printers/{printer_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrinterRequestBuilder instantiates a new PrinterRequestBuilder and sets the default values.
func NewPrinterRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrinterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the list of printers registered in the tenant.
func (m *PrinterRequestBuilder) CreateDeleteRequestInformation(options *PrinterRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of printers registered in the tenant.
func (m *PrinterRequestBuilder) CreateGetRequestInformation(options *PrinterRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
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
// CreatePatchRequestInformation the list of printers registered in the tenant.
func (m *PrinterRequestBuilder) CreatePatchRequestInformation(options *PrinterRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the list of printers registered in the tenant.
func (m *PrinterRequestBuilder) Delete(options *PrinterRequestBuilderDeleteOptions)(error) {
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
// Get the list of printers registered in the tenant.
func (m *PrinterRequestBuilder) Get(options *PrinterRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Printer, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPrinter() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Printer), nil
}
// Patch the list of printers registered in the tenant.
func (m *PrinterRequestBuilder) Patch(options *PrinterRequestBuilderPatchOptions)(error) {
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
func (m *PrinterRequestBuilder) RestoreFactoryDefaults()(*i95e425711e2e59b85807bfbd36bd18a7893a62957fccbb0dcc17647c9c3f5e90.RestoreFactoryDefaultsRequestBuilder) {
    return i95e425711e2e59b85807bfbd36bd18a7893a62957fccbb0dcc17647c9c3f5e90.NewRestoreFactoryDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) Shares()(*i74b470ddc01590a04dfc5ba245b7ccb094a2f7b819d0a4ea94773869da1de8bc.SharesRequestBuilder) {
    return i74b470ddc01590a04dfc5ba245b7ccb094a2f7b819d0a4ea94773869da1de8bc.NewSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) TaskTriggers()(*i05f5154d5dc8e1096ef3ba18fb5bdb890712250391686b0ffd44f043e05db1fb.TaskTriggersRequestBuilder) {
    return i05f5154d5dc8e1096ef3ba18fb5bdb890712250391686b0ffd44f043e05db1fb.NewTaskTriggersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskTriggersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.print.printers.item.taskTriggers.item collection
func (m *PrinterRequestBuilder) TaskTriggersById(id string)(*i8a35ad54635f07016096adf7157efc92ede0057cd4a790d06a32332ce94f22cc.PrintTaskTriggerRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printTaskTrigger_id"] = id
    }
    return i8a35ad54635f07016096adf7157efc92ede0057cd4a790d06a32332ce94f22cc.NewPrintTaskTriggerRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
