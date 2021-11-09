package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1b63d3f13e84ba5e6ea683797730a779bc0a9065ea44069338926a19812ad239 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/noncustodialdatasources/item/updateindex"
    i4c3d3b6907651fb3d7dd7d818a9c692008d404244415bf3c118fce7524aa40a0 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/noncustodialdatasources/item/datasource"
    ifae0b845c7ebf0ec6d5b74823e20670fc850fa370f6a7a747b6ae7c02b8a9b5d "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/noncustodialdatasources/item/release"
)

// Builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\noncustodialDataSources\{noncustodialDataSource-id}
type NoncustodialDataSourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type NoncustodialDataSourceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type NoncustodialDataSourceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *NoncustodialDataSourceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Returns a list of case noncustodialDataSource objects for this case.  Nullable.
type NoncustodialDataSourceRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type NoncustodialDataSourceRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NoncustodialDataSource;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new NoncustodialDataSourceRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewNoncustodialDataSourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NoncustodialDataSourceRequestBuilder) {
    m := &NoncustodialDataSourceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/noncustodialDataSources/{noncustodialDataSource_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new NoncustodialDataSourceRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewNoncustodialDataSourceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NoncustodialDataSourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNoncustodialDataSourceRequestBuilderInternal(urlParams, requestAdapter)
}
// Returns a list of case noncustodialDataSource objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *NoncustodialDataSourceRequestBuilder) CreateDeleteRequestInformation(options *NoncustodialDataSourceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns a list of case noncustodialDataSource objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *NoncustodialDataSourceRequestBuilder) CreateGetRequestInformation(options *NoncustodialDataSourceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns a list of case noncustodialDataSource objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *NoncustodialDataSourceRequestBuilder) CreatePatchRequestInformation(options *NoncustodialDataSourceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *NoncustodialDataSourceRequestBuilder) DataSource()(*i4c3d3b6907651fb3d7dd7d818a9c692008d404244415bf3c118fce7524aa40a0.DataSourceRequestBuilder) {
    return i4c3d3b6907651fb3d7dd7d818a9c692008d404244415bf3c118fce7524aa40a0.NewDataSourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Returns a list of case noncustodialDataSource objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *NoncustodialDataSourceRequestBuilder) Delete(options *NoncustodialDataSourceRequestBuilderDeleteOptions)(error) {
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
// Returns a list of case noncustodialDataSource objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *NoncustodialDataSourceRequestBuilder) Get(options *NoncustodialDataSourceRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NoncustodialDataSource, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewNoncustodialDataSource() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NoncustodialDataSource), nil
}
// Returns a list of case noncustodialDataSource objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *NoncustodialDataSourceRequestBuilder) Patch(options *NoncustodialDataSourceRequestBuilderPatchOptions)(error) {
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
func (m *NoncustodialDataSourceRequestBuilder) Release()(*ifae0b845c7ebf0ec6d5b74823e20670fc850fa370f6a7a747b6ae7c02b8a9b5d.ReleaseRequestBuilder) {
    return ifae0b845c7ebf0ec6d5b74823e20670fc850fa370f6a7a747b6ae7c02b8a9b5d.NewReleaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *NoncustodialDataSourceRequestBuilder) UpdateIndex()(*i1b63d3f13e84ba5e6ea683797730a779bc0a9065ea44069338926a19812ad239.UpdateIndexRequestBuilder) {
    return i1b63d3f13e84ba5e6ea683797730a779bc0a9065ea44069338926a19812ad239.NewUpdateIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
