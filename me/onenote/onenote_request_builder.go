package onenote

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i54a3bfcdd648d6ee47e87c311695b29c4737b8921bb8f2a506c0864844df5dba "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sections"
    i97f70d655aa735279e41729aa6800f01397d2d7e8d41164003eff9b47387d3a5 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/notebooks"
    iac72d35f9c20cb360839134d86d535f87e1cdbe10b719f8c792d8017c06edd1d "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/operations"
    iae2d6340b64994bda52b72186cf936e932db0b6d72196a85d99b8d63141db19b "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/pages"
    ic51747a34f2e47436a56cc441d712037e581ac2bd968645ce4fe67a401a2f4bf "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/resources"
    id71911f826ddd0ad6edad424ad81ad5b822e25039b7e26730ec6b2f6de37ad3b "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups"
    i1927275756e9aeb82a189ce572418c67c8a998c4353590bc94f2a62a876dae2d "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/notebooks/item"
    i1941888ff368eb1fd130194825ce82fb010616c26e7af09e43370fa6bb2cfd71 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item"
    i2b8ad7bd2ae87c92c8852bbeef996852d7685df93cb12b50d58ecc094ecfd124 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/operations/item"
    i9ffb03d414d6abd9f024bf2d1f19d7b7fa97959d1245465fdaef4d163ddd6de4 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/pages/item"
    ibe60e1baafd5c0c7f830c971df51a300df66822516c0cfcd8a43fd2adc385c80 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sections/item"
    id67049e3fa5525a6ba50b6b8bc9f60107ac06f61256f8cf55bb655b4552ecd21 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/resources/item"
)

// Builds and executes requests for operations under \me\onenote
type OnenoteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type OnenoteRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type OnenoteRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnenoteRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Read-only.
type OnenoteRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type OnenoteRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Onenote;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new OnenoteRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOnenoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteRequestBuilder) {
    m := &OnenoteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onenote{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new OnenoteRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOnenoteRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteRequestBuilderInternal(urlParams, requestAdapter)
}
// Read-only.
// Parameters:
//  - options : Options for the request
func (m *OnenoteRequestBuilder) CreateDeleteRequestInformation(options *OnenoteRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only.
// Parameters:
//  - options : Options for the request
func (m *OnenoteRequestBuilder) CreateGetRequestInformation(options *OnenoteRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only.
// Parameters:
//  - options : Options for the request
func (m *OnenoteRequestBuilder) CreatePatchRequestInformation(options *OnenoteRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only.
// Parameters:
//  - options : Options for the request
func (m *OnenoteRequestBuilder) Delete(options *OnenoteRequestBuilderDeleteOptions)(error) {
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
// Read-only.
// Parameters:
//  - options : Options for the request
func (m *OnenoteRequestBuilder) Get(options *OnenoteRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Onenote, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOnenote() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Onenote), nil
}
func (m *OnenoteRequestBuilder) Notebooks()(*i97f70d655aa735279e41729aa6800f01397d2d7e8d41164003eff9b47387d3a5.NotebooksRequestBuilder) {
    return i97f70d655aa735279e41729aa6800f01397d2d7e8d41164003eff9b47387d3a5.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.onenote.notebooks.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*i1927275756e9aeb82a189ce572418c67c8a998c4353590bc94f2a62a876dae2d.NotebookRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook_id"] = id
    }
    return i1927275756e9aeb82a189ce572418c67c8a998c4353590bc94f2a62a876dae2d.NewNotebookRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Operations()(*iac72d35f9c20cb360839134d86d535f87e1cdbe10b719f8c792d8017c06edd1d.OperationsRequestBuilder) {
    return iac72d35f9c20cb360839134d86d535f87e1cdbe10b719f8c792d8017c06edd1d.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.onenote.operations.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i2b8ad7bd2ae87c92c8852bbeef996852d7685df93cb12b50d58ecc094ecfd124.OnenoteOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation_id"] = id
    }
    return i2b8ad7bd2ae87c92c8852bbeef996852d7685df93cb12b50d58ecc094ecfd124.NewOnenoteOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Pages()(*iae2d6340b64994bda52b72186cf936e932db0b6d72196a85d99b8d63141db19b.PagesRequestBuilder) {
    return iae2d6340b64994bda52b72186cf936e932db0b6d72196a85d99b8d63141db19b.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.onenote.pages.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) PagesById(id string)(*i9ffb03d414d6abd9f024bf2d1f19d7b7fa97959d1245465fdaef4d163ddd6de4.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i9ffb03d414d6abd9f024bf2d1f19d7b7fa97959d1245465fdaef4d163ddd6de4.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Read-only.
// Parameters:
//  - options : Options for the request
func (m *OnenoteRequestBuilder) Patch(options *OnenoteRequestBuilderPatchOptions)(error) {
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
func (m *OnenoteRequestBuilder) Resources()(*ic51747a34f2e47436a56cc441d712037e581ac2bd968645ce4fe67a401a2f4bf.ResourcesRequestBuilder) {
    return ic51747a34f2e47436a56cc441d712037e581ac2bd968645ce4fe67a401a2f4bf.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.onenote.resources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*id67049e3fa5525a6ba50b6b8bc9f60107ac06f61256f8cf55bb655b4552ecd21.OnenoteResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource_id"] = id
    }
    return id67049e3fa5525a6ba50b6b8bc9f60107ac06f61256f8cf55bb655b4552ecd21.NewOnenoteResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) SectionGroups()(*id71911f826ddd0ad6edad424ad81ad5b822e25039b7e26730ec6b2f6de37ad3b.SectionGroupsRequestBuilder) {
    return id71911f826ddd0ad6edad424ad81ad5b822e25039b7e26730ec6b2f6de37ad3b.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.onenote.sectionGroups.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*i1941888ff368eb1fd130194825ce82fb010616c26e7af09e43370fa6bb2cfd71.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i1941888ff368eb1fd130194825ce82fb010616c26e7af09e43370fa6bb2cfd71.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Sections()(*i54a3bfcdd648d6ee47e87c311695b29c4737b8921bb8f2a506c0864844df5dba.SectionsRequestBuilder) {
    return i54a3bfcdd648d6ee47e87c311695b29c4737b8921bb8f2a506c0864844df5dba.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.onenote.sections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) SectionsById(id string)(*ibe60e1baafd5c0c7f830c971df51a300df66822516c0cfcd8a43fd2adc385c80.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return ibe60e1baafd5c0c7f830c971df51a300df66822516c0cfcd8a43fd2adc385c80.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
