package onenote

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i025d995a3eab0912bb44d350ddf5d7a3b0aeb9332e0629260ad912238bc917e4 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages"
    i0decb69f831ce545caa107a838115b23576ee1efac361ea22cf80d81cb4ded00 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/resources"
    id056e9dcc48f2e5766390396802781a6f6e4a85d729a1d7f65531745ed03bd15 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/notebooks"
    ifc9a684d2dbdb1fe7d1eae49e0a8c30c188888963093335a2e0b25cb77cacb10 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/sectiongroups"
    ifedb0f50ef2ef408f7a3fbfc146cacec13291103da0bba733f4b2757f5d4de71 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/operations"
    iff11e6eb619f9d0b688f685398702dc1c9c46658026db69f2e1441e49a1f8f3b "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/sections"
    i2ba021dd48264681e34929a50215efa4ffab4a38e1e781d5129ff35bc76ada21 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/notebooks/item"
    i2e9176f71b9cfb53869eedb784edba98a6ec7f334a527247849d1da13f2a52a6 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/sectiongroups/item"
    ib6e5547c6ec7db85ec94ebc2072c06598dc08ff7d88cc9151447095852200162 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/resources/item"
    ida022400bfcca3ce3366c6354f31a6136d19a7aa86866cfd1c28151b05cbee04 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/sections/item"
    ie1466039fa513108cc7b0f3d5f87933999e122789d21aa1fd0e611cacc94cd21 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item"
    ie9c35a4053aa937b7cbd985746f2c474e4823ad0293ce536d11acb1d37b1f790 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/operations/item"
)

// OnenoteRequestBuilder builds and executes requests for operations under \sites\{site-id}\onenote
type OnenoteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenoteRequestBuilderDeleteOptions options for Delete
type OnenoteRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenoteRequestBuilderGetOptions options for Get
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
// OnenoteRequestBuilderGetQueryParameters calls the OneNote service for notebook related operations.
type OnenoteRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenoteRequestBuilderPatchOptions options for Patch
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
// NewOnenoteRequestBuilderInternal instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteRequestBuilder) {
    m := &OnenoteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/onenote{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteRequestBuilder instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation calls the OneNote service for notebook related operations.
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
// CreateGetRequestInformation calls the OneNote service for notebook related operations.
func (m *OnenoteRequestBuilder) CreateGetRequestInformation(options *OnenoteRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation calls the OneNote service for notebook related operations.
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
// Delete calls the OneNote service for notebook related operations.
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
// Get calls the OneNote service for notebook related operations.
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
func (m *OnenoteRequestBuilder) Notebooks()(*id056e9dcc48f2e5766390396802781a6f6e4a85d729a1d7f65531745ed03bd15.NotebooksRequestBuilder) {
    return id056e9dcc48f2e5766390396802781a6f6e4a85d729a1d7f65531745ed03bd15.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotebooksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.notebooks.item collection
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*i2ba021dd48264681e34929a50215efa4ffab4a38e1e781d5129ff35bc76ada21.NotebookRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook_id"] = id
    }
    return i2ba021dd48264681e34929a50215efa4ffab4a38e1e781d5129ff35bc76ada21.NewNotebookRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Operations()(*ifedb0f50ef2ef408f7a3fbfc146cacec13291103da0bba733f4b2757f5d4de71.OperationsRequestBuilder) {
    return ifedb0f50ef2ef408f7a3fbfc146cacec13291103da0bba733f4b2757f5d4de71.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.operations.item collection
func (m *OnenoteRequestBuilder) OperationsById(id string)(*ie9c35a4053aa937b7cbd985746f2c474e4823ad0293ce536d11acb1d37b1f790.OnenoteOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation_id"] = id
    }
    return ie9c35a4053aa937b7cbd985746f2c474e4823ad0293ce536d11acb1d37b1f790.NewOnenoteOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Pages()(*i025d995a3eab0912bb44d350ddf5d7a3b0aeb9332e0629260ad912238bc917e4.PagesRequestBuilder) {
    return i025d995a3eab0912bb44d350ddf5d7a3b0aeb9332e0629260ad912238bc917e4.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.pages.item collection
func (m *OnenoteRequestBuilder) PagesById(id string)(*ie1466039fa513108cc7b0f3d5f87933999e122789d21aa1fd0e611cacc94cd21.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return ie1466039fa513108cc7b0f3d5f87933999e122789d21aa1fd0e611cacc94cd21.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch calls the OneNote service for notebook related operations.
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
func (m *OnenoteRequestBuilder) Resources()(*i0decb69f831ce545caa107a838115b23576ee1efac361ea22cf80d81cb4ded00.ResourcesRequestBuilder) {
    return i0decb69f831ce545caa107a838115b23576ee1efac361ea22cf80d81cb4ded00.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.resources.item collection
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*ib6e5547c6ec7db85ec94ebc2072c06598dc08ff7d88cc9151447095852200162.OnenoteResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource_id"] = id
    }
    return ib6e5547c6ec7db85ec94ebc2072c06598dc08ff7d88cc9151447095852200162.NewOnenoteResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) SectionGroups()(*ifc9a684d2dbdb1fe7d1eae49e0a8c30c188888963093335a2e0b25cb77cacb10.SectionGroupsRequestBuilder) {
    return ifc9a684d2dbdb1fe7d1eae49e0a8c30c188888963093335a2e0b25cb77cacb10.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.sectionGroups.item collection
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*i2e9176f71b9cfb53869eedb784edba98a6ec7f334a527247849d1da13f2a52a6.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i2e9176f71b9cfb53869eedb784edba98a6ec7f334a527247849d1da13f2a52a6.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Sections()(*iff11e6eb619f9d0b688f685398702dc1c9c46658026db69f2e1441e49a1f8f3b.SectionsRequestBuilder) {
    return iff11e6eb619f9d0b688f685398702dc1c9c46658026db69f2e1441e49a1f8f3b.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.sections.item collection
func (m *OnenoteRequestBuilder) SectionsById(id string)(*ida022400bfcca3ce3366c6354f31a6136d19a7aa86866cfd1c28151b05cbee04.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return ida022400bfcca3ce3366c6354f31a6136d19a7aa86866cfd1c28151b05cbee04.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
