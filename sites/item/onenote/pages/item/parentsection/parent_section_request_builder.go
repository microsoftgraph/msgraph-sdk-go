package parentsection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i27b5ce8e3d5ccd754c59c5e53f83108ad49b67de1fb59505fab6b167c9aa1b92 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/parentsection/pages"
    i70efb0b6c0cb5cd9d61464740feca3cf80b411b2ecb7b8f1ffd62e0a5c731231 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/parentsection/copytosectiongroup"
    ia6f06eead351d266814f8ab7de42974e24e84676ab12ed584a198578540a8451 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/parentsection/parentnotebook"
    ida7d6aaa7398bec6015dbf60ff0b886471c733d56778be6ba839e3600f527c4f "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/parentsection/parentsectiongroup"
    if7a7b0e089096540f897f16bc90cf9c3db087b1dab1283afe2e59b057a3e1ee9 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/parentsection/copytonotebook"
    iecb005258b435242c1720a83b01c52a239c6bb2181e5f005b69445d830f93215 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/parentsection/pages/item"
)

// Builds and executes requests for operations under \sites\{site-id}\onenote\pages\{onenotePage-id}\parentSection
type ParentSectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ParentSectionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ParentSectionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ParentSectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The section that contains the page. Read-only.
type ParentSectionRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ParentSectionRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenoteSection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new ParentSectionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewParentSectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentSectionRequestBuilder) {
    m := &ParentSectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/onenote/pages/{onenotePage_id}/parentSection{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ParentSectionRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewParentSectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentSectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParentSectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ParentSectionRequestBuilder) CopyToNotebook()(*if7a7b0e089096540f897f16bc90cf9c3db087b1dab1283afe2e59b057a3e1ee9.CopyToNotebookRequestBuilder) {
    return if7a7b0e089096540f897f16bc90cf9c3db087b1dab1283afe2e59b057a3e1ee9.NewCopyToNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParentSectionRequestBuilder) CopyToSectionGroup()(*i70efb0b6c0cb5cd9d61464740feca3cf80b411b2ecb7b8f1ffd62e0a5c731231.CopyToSectionGroupRequestBuilder) {
    return i70efb0b6c0cb5cd9d61464740feca3cf80b411b2ecb7b8f1ffd62e0a5c731231.NewCopyToSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The section that contains the page. Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentSectionRequestBuilder) CreateDeleteRequestInformation(options *ParentSectionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The section that contains the page. Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentSectionRequestBuilder) CreateGetRequestInformation(options *ParentSectionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The section that contains the page. Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentSectionRequestBuilder) CreatePatchRequestInformation(options *ParentSectionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The section that contains the page. Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentSectionRequestBuilder) Delete(options *ParentSectionRequestBuilderDeleteOptions)(error) {
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
// The section that contains the page. Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentSectionRequestBuilder) Get(options *ParentSectionRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenoteSection, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOnenoteSection() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenoteSection), nil
}
func (m *ParentSectionRequestBuilder) Pages()(*i27b5ce8e3d5ccd754c59c5e53f83108ad49b67de1fb59505fab6b167c9aa1b92.PagesRequestBuilder) {
    return i27b5ce8e3d5ccd754c59c5e53f83108ad49b67de1fb59505fab6b167c9aa1b92.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.pages.item.parentSection.pages.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ParentSectionRequestBuilder) PagesById(id string)(*iecb005258b435242c1720a83b01c52a239c6bb2181e5f005b69445d830f93215.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id1"] = id
    }
    return iecb005258b435242c1720a83b01c52a239c6bb2181e5f005b69445d830f93215.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ParentSectionRequestBuilder) ParentNotebook()(*ia6f06eead351d266814f8ab7de42974e24e84676ab12ed584a198578540a8451.ParentNotebookRequestBuilder) {
    return ia6f06eead351d266814f8ab7de42974e24e84676ab12ed584a198578540a8451.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParentSectionRequestBuilder) ParentSectionGroup()(*ida7d6aaa7398bec6015dbf60ff0b886471c733d56778be6ba839e3600f527c4f.ParentSectionGroupRequestBuilder) {
    return ida7d6aaa7398bec6015dbf60ff0b886471c733d56778be6ba839e3600f527c4f.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The section that contains the page. Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentSectionRequestBuilder) Patch(options *ParentSectionRequestBuilderPatchOptions)(error) {
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
