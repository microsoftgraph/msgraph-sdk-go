package filter

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i361c59c66bb7e477a88ead7c5a67567e4e997f4d894a54ba9e072484ba7b83e6 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applytopitemsfilter"
    i37e17448ebc3c88d0d568891a6a197d87b83301cfd1f5ac2e5ecd4cbd27fead1 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applyfontcolorfilter"
    i4e34fe940e728eb7850075738d7e53292cccb6a4e341ef12bf32be59cec52a59 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applybottomitemsfilter"
    i6e765936f672b9f83a474ef2df99b618cc00b4b8f9611c0f0bd2c71bf9d996f4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/clear"
    i7b8e11c3e13cc29e60e05c7a93c0e721e488611ea6b275c7f14bd9eb2024424d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applybottompercentfilter"
    i7ec0979b141deaf9e8c06a54957e0b767a0d4073438da611add9d8be6674e3d4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applycellcolorfilter"
    i83798e0dc53921050c45b8378700a0300a0aa620603edffa729c9154aeb12ae9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applydynamicfilter"
    i86c63d77e75b907a65db65d104488e99c50d2ffd1b18402c6778fd4eee11c529 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applycustomfilter"
    i9351ada88d8f833003cd009a69dfaa9bf39e91b79b90be867afa22f4211d3fc4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applyiconfilter"
    ib9b678a6142ab4d7270bcaba151ff35c7971bd42a6caa29906cd0ceb2c4a15b2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/apply"
    ibbdc929a240a59a714772b7f8e7e089397d74208a4f2d7454e1c55c145dbb5e0 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applytoppercentfilter"
    ief856482f2d134a9efe38be849364abfebe3d62f31aa9dc78422dc6b1f07a7d3 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/columns/item/filter/applyvaluesfilter"
)

// Builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\columns\{workbookTableColumn-id}\filter
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
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
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
func (m *FilterRequestBuilder) Apply()(*ib9b678a6142ab4d7270bcaba151ff35c7971bd42a6caa29906cd0ceb2c4a15b2.ApplyRequestBuilder) {
    return ib9b678a6142ab4d7270bcaba151ff35c7971bd42a6caa29906cd0ceb2c4a15b2.NewApplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyBottomItemsFilter()(*i4e34fe940e728eb7850075738d7e53292cccb6a4e341ef12bf32be59cec52a59.ApplyBottomItemsFilterRequestBuilder) {
    return i4e34fe940e728eb7850075738d7e53292cccb6a4e341ef12bf32be59cec52a59.NewApplyBottomItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyBottomPercentFilter()(*i7b8e11c3e13cc29e60e05c7a93c0e721e488611ea6b275c7f14bd9eb2024424d.ApplyBottomPercentFilterRequestBuilder) {
    return i7b8e11c3e13cc29e60e05c7a93c0e721e488611ea6b275c7f14bd9eb2024424d.NewApplyBottomPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyCellColorFilter()(*i7ec0979b141deaf9e8c06a54957e0b767a0d4073438da611add9d8be6674e3d4.ApplyCellColorFilterRequestBuilder) {
    return i7ec0979b141deaf9e8c06a54957e0b767a0d4073438da611add9d8be6674e3d4.NewApplyCellColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyCustomFilter()(*i86c63d77e75b907a65db65d104488e99c50d2ffd1b18402c6778fd4eee11c529.ApplyCustomFilterRequestBuilder) {
    return i86c63d77e75b907a65db65d104488e99c50d2ffd1b18402c6778fd4eee11c529.NewApplyCustomFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyDynamicFilter()(*i83798e0dc53921050c45b8378700a0300a0aa620603edffa729c9154aeb12ae9.ApplyDynamicFilterRequestBuilder) {
    return i83798e0dc53921050c45b8378700a0300a0aa620603edffa729c9154aeb12ae9.NewApplyDynamicFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyFontColorFilter()(*i37e17448ebc3c88d0d568891a6a197d87b83301cfd1f5ac2e5ecd4cbd27fead1.ApplyFontColorFilterRequestBuilder) {
    return i37e17448ebc3c88d0d568891a6a197d87b83301cfd1f5ac2e5ecd4cbd27fead1.NewApplyFontColorFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyIconFilter()(*i9351ada88d8f833003cd009a69dfaa9bf39e91b79b90be867afa22f4211d3fc4.ApplyIconFilterRequestBuilder) {
    return i9351ada88d8f833003cd009a69dfaa9bf39e91b79b90be867afa22f4211d3fc4.NewApplyIconFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyTopItemsFilter()(*i361c59c66bb7e477a88ead7c5a67567e4e997f4d894a54ba9e072484ba7b83e6.ApplyTopItemsFilterRequestBuilder) {
    return i361c59c66bb7e477a88ead7c5a67567e4e997f4d894a54ba9e072484ba7b83e6.NewApplyTopItemsFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyTopPercentFilter()(*ibbdc929a240a59a714772b7f8e7e089397d74208a4f2d7454e1c55c145dbb5e0.ApplyTopPercentFilterRequestBuilder) {
    return ibbdc929a240a59a714772b7f8e7e089397d74208a4f2d7454e1c55c145dbb5e0.NewApplyTopPercentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) ApplyValuesFilter()(*ief856482f2d134a9efe38be849364abfebe3d62f31aa9dc78422dc6b1f07a7d3.ApplyValuesFilterRequestBuilder) {
    return ief856482f2d134a9efe38be849364abfebe3d62f31aa9dc78422dc6b1f07a7d3.NewApplyValuesFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FilterRequestBuilder) Clear()(*i6e765936f672b9f83a474ef2df99b618cc00b4b8f9611c0f0bd2c71bf9d996f4.ClearRequestBuilder) {
    return i6e765936f672b9f83a474ef2df99b618cc00b4b8f9611c0f0bd2c71bf9d996f4.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new FilterRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FilterRequestBuilder) {
    m := &FilterRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet/tables/{workbookTable_id}/columns/{workbookTableColumn_id}/filter{?select,expand}";
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
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
