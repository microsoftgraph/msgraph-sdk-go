package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrintPrintersPrinterItemRequestBuilder provides operations to manage the printers property of the microsoft.graph.print entity.
type PrintPrintersPrinterItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrintPrintersPrinterItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintPrintersPrinterItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrintPrintersPrinterItemRequestBuilderGetQueryParameters the list of printers registered in the tenant.
type PrintPrintersPrinterItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrintPrintersPrinterItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintPrintersPrinterItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrintPrintersPrinterItemRequestBuilderGetQueryParameters
}
// PrintPrintersPrinterItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintPrintersPrinterItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Connectors provides operations to manage the connectors property of the microsoft.graph.printer entity.
func (m *PrintPrintersPrinterItemRequestBuilder) Connectors()(*PrintPrintersItemConnectorsRequestBuilder) {
    return NewPrintPrintersItemConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectorsById provides operations to manage the connectors property of the microsoft.graph.printer entity.
func (m *PrintPrintersPrinterItemRequestBuilder) ConnectorsById(id string)(*PrintPrintersItemConnectorsPrintConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printConnector%2Did"] = id
    }
    return NewPrintPrintersItemConnectorsPrintConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPrintPrintersPrinterItemRequestBuilderInternal instantiates a new PrinterItemRequestBuilder and sets the default values.
func NewPrintPrintersPrinterItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintPrintersPrinterItemRequestBuilder) {
    m := &PrintPrintersPrinterItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/printers/{printer%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrintPrintersPrinterItemRequestBuilder instantiates a new PrinterItemRequestBuilder and sets the default values.
func NewPrintPrintersPrinterItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintPrintersPrinterItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintPrintersPrinterItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property printers for print
func (m *PrintPrintersPrinterItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *PrintPrintersPrinterItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the list of printers registered in the tenant.
func (m *PrintPrintersPrinterItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PrintPrintersPrinterItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property printers in print
func (m *PrintPrintersPrinterItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Printerable, requestConfiguration *PrintPrintersPrinterItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property printers for print
func (m *PrintPrintersPrinterItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrintPrintersPrinterItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the list of printers registered in the tenant.
func (m *PrintPrintersPrinterItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrintPrintersPrinterItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Printerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrinterFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Printerable), nil
}
// Patch update the navigation property printers in print
func (m *PrintPrintersPrinterItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Printerable, requestConfiguration *PrintPrintersPrinterItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Printerable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrinterFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Printerable), nil
}
// RestoreFactoryDefaults provides operations to call the restoreFactoryDefaults method.
func (m *PrintPrintersPrinterItemRequestBuilder) RestoreFactoryDefaults()(*PrintPrintersItemRestoreFactoryDefaultsRequestBuilder) {
    return NewPrintPrintersItemRestoreFactoryDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Shares provides operations to manage the shares property of the microsoft.graph.printer entity.
func (m *PrintPrintersPrinterItemRequestBuilder) Shares()(*PrintPrintersItemSharesRequestBuilder) {
    return NewPrintPrintersItemSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharesById provides operations to manage the shares property of the microsoft.graph.printer entity.
func (m *PrintPrintersPrinterItemRequestBuilder) SharesById(id string)(*PrintPrintersItemSharesPrinterShareItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printerShare%2Did"] = id
    }
    return NewPrintPrintersItemSharesPrinterShareItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaskTriggers provides operations to manage the taskTriggers property of the microsoft.graph.printer entity.
func (m *PrintPrintersPrinterItemRequestBuilder) TaskTriggers()(*PrintPrintersItemTaskTriggersRequestBuilder) {
    return NewPrintPrintersItemTaskTriggersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskTriggersById provides operations to manage the taskTriggers property of the microsoft.graph.printer entity.
func (m *PrintPrintersPrinterItemRequestBuilder) TaskTriggersById(id string)(*PrintPrintersItemTaskTriggersPrintTaskTriggerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printTaskTrigger%2Did"] = id
    }
    return NewPrintPrintersItemTaskTriggersPrintTaskTriggerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
