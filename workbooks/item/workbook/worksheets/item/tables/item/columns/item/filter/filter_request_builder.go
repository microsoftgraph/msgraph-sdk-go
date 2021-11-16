package filter

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i074b29baa79648bf3b3aebfd606afaa237c5250e2efcf8b6b816daf4e3b39638 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/clear"
    i12e9e1694686ad3970c526b0cf764ea76c5e01a368b2bd52cdcd916292a79806 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applybottomitemsfilter"
    i20bd35e3d1e15e53b89433876cf620911614ee247e61c2f12d3912193fc08bef "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applyiconfilter"
    i2df76a49aaca1109a284d636d3cbe9c7e2da19d673dfd4fcd9fd5129182bc9cc "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applycellcolorfilter"
    i7ac70352599c75c430c483c6e4772e922e1dced44e9d06b913f3d0262c1639f2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applycustomfilter"
    i8fa4cd2d67454e174ebeca06aa59d1a4f6042a474c8373f4a54e3e91c6625eaf "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applytopitemsfilter"
    ia7ea04d1f4e7819640e9a6aa90d1e304b408bedc7e2a2e922f10f9dbf026bc75 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applytoppercentfilter"
    ib0e2b3d1228cf3527b4e57814bc8199316b43365c556bc1f85d3490124a61b11 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applybottompercentfilter"
    ib416ce06f4b78621e7f88b5df2d4132756fce9d6763dd96c3b76d239a83a7403 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applydynamicfilter"
    ibc22e77a492c42f558ed799b77c3cd3f5d094eec69998ed94fad96e9ee94f93d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/apply"
    idd197cb3bcd87f64f3171a2593fa379ece591bae21771270cb2318471a09e5ff "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applyfontcolorfilter"
    ide05357466b976f2da283467d98fd838d0a2030c91e163259dcef1ee880c811a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/worksheets/item/tables/item/columns/item/filter/applyvaluesfilter"
)

// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\filter
type FilterRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type FilterRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type FilterRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *FilterRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Retrieve the filter applied to the column. Read-only.
type FilterRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type FilterRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookFilter;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *FilterRequestBuilder) Apply()(*ibc22e77a492c42f558ed799b77c3cd3f5d094eec69998ed94fad96e9ee94f93d.ApplyRequestBuilder) {
    return ibc22e77a492c42f558ed799b77c3cd3f5d094eec69998ed94fad96e9ee94f93d.NewApplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyBottomItemsFilter()(*i12e9e1694686ad3970c526b0cf764ea76c5e01a368b2bd52cdcd916292a79806.ApplyBottomItemsFilterRequestBuilder) {
    return i12e9e1694686ad3970c526b0cf764ea76c5e01a368b2bd52cdcd916292a79806.NewApplyBottomItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyBottomPercentFilter()(*ib0e2b3d1228cf3527b4e57814bc8199316b43365c556bc1f85d3490124a61b11.ApplyBottomPercentFilterRequestBuilder) {
    return ib0e2b3d1228cf3527b4e57814bc8199316b43365c556bc1f85d3490124a61b11.NewApplyBottomPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyCellColorFilter()(*i2df76a49aaca1109a284d636d3cbe9c7e2da19d673dfd4fcd9fd5129182bc9cc.ApplyCellColorFilterRequestBuilder) {
    return i2df76a49aaca1109a284d636d3cbe9c7e2da19d673dfd4fcd9fd5129182bc9cc.NewApplyCellColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyCustomFilter()(*i7ac70352599c75c430c483c6e4772e922e1dced44e9d06b913f3d0262c1639f2.ApplyCustomFilterRequestBuilder) {
    return i7ac70352599c75c430c483c6e4772e922e1dced44e9d06b913f3d0262c1639f2.NewApplyCustomFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyDynamicFilter()(*ib416ce06f4b78621e7f88b5df2d4132756fce9d6763dd96c3b76d239a83a7403.ApplyDynamicFilterRequestBuilder) {
    return ib416ce06f4b78621e7f88b5df2d4132756fce9d6763dd96c3b76d239a83a7403.NewApplyDynamicFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyFontColorFilter()(*idd197cb3bcd87f64f3171a2593fa379ece591bae21771270cb2318471a09e5ff.ApplyFontColorFilterRequestBuilder) {
    return idd197cb3bcd87f64f3171a2593fa379ece591bae21771270cb2318471a09e5ff.NewApplyFontColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyIconFilter()(*i20bd35e3d1e15e53b89433876cf620911614ee247e61c2f12d3912193fc08bef.ApplyIconFilterRequestBuilder) {
    return i20bd35e3d1e15e53b89433876cf620911614ee247e61c2f12d3912193fc08bef.NewApplyIconFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyTopItemsFilter()(*i8fa4cd2d67454e174ebeca06aa59d1a4f6042a474c8373f4a54e3e91c6625eaf.ApplyTopItemsFilterRequestBuilder) {
    return i8fa4cd2d67454e174ebeca06aa59d1a4f6042a474c8373f4a54e3e91c6625eaf.NewApplyTopItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyTopPercentFilter()(*ia7ea04d1f4e7819640e9a6aa90d1e304b408bedc7e2a2e922f10f9dbf026bc75.ApplyTopPercentFilterRequestBuilder) {
    return ia7ea04d1f4e7819640e9a6aa90d1e304b408bedc7e2a2e922f10f9dbf026bc75.NewApplyTopPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyValuesFilter()(*ide05357466b976f2da283467d98fd838d0a2030c91e163259dcef1ee880c811a.ApplyValuesFilterRequestBuilder) {
    return ide05357466b976f2da283467d98fd838d0a2030c91e163259dcef1ee880c811a.NewApplyValuesFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) Clear()(*i074b29baa79648bf3b3aebfd606afaa237c5250e2efcf8b6b816daf4e3b39638.ClearRequestBuilder) {
    return i074b29baa79648bf3b3aebfd606afaa237c5250e2efcf8b6b816daf4e3b39638.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new FilterRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FilterRequestBuilder) {
    m := &FilterRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}/tables/{workbookTable_id}/columns/{workbookTableColumn_id}/filter{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new FilterRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewFilterRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilterRequestBuilderInternal(urlParams, requestAdapter)
}
// Retrieve the filter applied to the column. Read-only.
// Parameters:
//  - options : Options for the request
func (m *FilterRequestBuilder) CreateDeleteRequestInformation(options *FilterRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Retrieve the filter applied to the column. Read-only.
// Parameters:
//  - options : Options for the request
func (m *FilterRequestBuilder) CreateGetRequestInformation(options *FilterRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Retrieve the filter applied to the column. Read-only.
// Parameters:
//  - options : Options for the request
func (m *FilterRequestBuilder) CreatePatchRequestInformation(options *FilterRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Retrieve the filter applied to the column. Read-only.
// Parameters:
//  - options : Options for the request
func (m *FilterRequestBuilder) Delete(options *FilterRequestBuilderDeleteOptions)(error) {
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
// Retrieve the filter applied to the column. Read-only.
// Parameters:
//  - options : Options for the request
func (m *FilterRequestBuilder) Get(options *FilterRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookFilter, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookFilter() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookFilter), nil
}
// Retrieve the filter applied to the column. Read-only.
// Parameters:
//  - options : Options for the request
func (m *FilterRequestBuilder) Patch(options *FilterRequestBuilderPatchOptions)(error) {
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
