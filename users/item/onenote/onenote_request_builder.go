package onenote

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i3521324ea2fa03aac1ba1029060d9326077d780f41956c3996cc7c08c74175d5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/operations"
    i6ab0f57415f0f91846673c7144401d1ecb437b2ed9583ff3a324fce02d79d6ce "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/sectiongroups"
    i89a64f4378390a1ff2c4709e25458a2484ce3ed41135ba5fb5e45f1f2c8feecd "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/sections"
    iac5f1f7ab3fe7f049908ebaa2118d2ae66da6365ddcc42c944df9fa3747dd82e "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/resources"
    ic1143c927ee5b190553f9b473aa6e674cb9c24f3a9e2793ca61c4126357f44d4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/pages"
    ieb61df1d5a7fd9e6e860ccddd58f15f53dabcafd6e50ba4704ac752bb5e98897 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks"
    i1f7cd27e6d43acad4f65fde049e01af4f5167bbae163aa9f54817ce16f3eb17e "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/pages/item"
    i2fd4361716803984f5766f676870e98432266f948b1f5e8bffb51c4c87176198 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/sections/item"
    i5ce07e3fde797b20643c20f1577367403edd721d4ddd3efd58a478430089aa94 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/operations/item"
    ie9d18cbc4daa0c9f216a92bd04bfbd2e093269deb326c4499728b346448951ef "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/resources/item"
    ief43f129f43e4e2e00209d40c6ab20f225d4a6be2f614f3de5c37ad2d36f6397 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/sectiongroups/item"
    ifa98a498cfd880b015264285e62bc7fee30c1a6cc2150f3d064f69448f9f5b27 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item"
)

// OnenoteRequestBuilder builds and executes requests for operations under \users\{user-id}\onenote
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
// OnenoteRequestBuilderGetQueryParameters read-only.
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
    m.urlTemplate = "{+baseurl}/users/{user_id}/onenote{?select,expand}";
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
// CreateDeleteRequestInformation read-only.
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
// CreateGetRequestInformation read-only.
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
// CreatePatchRequestInformation read-only.
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
// Delete read-only.
func (m *OnenoteRequestBuilder) Delete(options *OnenoteRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get read-only.
func (m *OnenoteRequestBuilder) Get(options *OnenoteRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Onenote, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOnenote() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Onenote), nil
}
func (m *OnenoteRequestBuilder) Notebooks()(*ieb61df1d5a7fd9e6e860ccddd58f15f53dabcafd6e50ba4704ac752bb5e98897.NotebooksRequestBuilder) {
    return ieb61df1d5a7fd9e6e860ccddd58f15f53dabcafd6e50ba4704ac752bb5e98897.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotebooksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.onenote.notebooks.item collection
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*ifa98a498cfd880b015264285e62bc7fee30c1a6cc2150f3d064f69448f9f5b27.NotebookRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook_id"] = id
    }
    return ifa98a498cfd880b015264285e62bc7fee30c1a6cc2150f3d064f69448f9f5b27.NewNotebookRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Operations()(*i3521324ea2fa03aac1ba1029060d9326077d780f41956c3996cc7c08c74175d5.OperationsRequestBuilder) {
    return i3521324ea2fa03aac1ba1029060d9326077d780f41956c3996cc7c08c74175d5.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.onenote.operations.item collection
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i5ce07e3fde797b20643c20f1577367403edd721d4ddd3efd58a478430089aa94.OnenoteOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation_id"] = id
    }
    return i5ce07e3fde797b20643c20f1577367403edd721d4ddd3efd58a478430089aa94.NewOnenoteOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Pages()(*ic1143c927ee5b190553f9b473aa6e674cb9c24f3a9e2793ca61c4126357f44d4.PagesRequestBuilder) {
    return ic1143c927ee5b190553f9b473aa6e674cb9c24f3a9e2793ca61c4126357f44d4.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.onenote.pages.item collection
func (m *OnenoteRequestBuilder) PagesById(id string)(*i1f7cd27e6d43acad4f65fde049e01af4f5167bbae163aa9f54817ce16f3eb17e.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i1f7cd27e6d43acad4f65fde049e01af4f5167bbae163aa9f54817ce16f3eb17e.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch read-only.
func (m *OnenoteRequestBuilder) Patch(options *OnenoteRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *OnenoteRequestBuilder) Resources()(*iac5f1f7ab3fe7f049908ebaa2118d2ae66da6365ddcc42c944df9fa3747dd82e.ResourcesRequestBuilder) {
    return iac5f1f7ab3fe7f049908ebaa2118d2ae66da6365ddcc42c944df9fa3747dd82e.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.onenote.resources.item collection
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*ie9d18cbc4daa0c9f216a92bd04bfbd2e093269deb326c4499728b346448951ef.OnenoteResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource_id"] = id
    }
    return ie9d18cbc4daa0c9f216a92bd04bfbd2e093269deb326c4499728b346448951ef.NewOnenoteResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) SectionGroups()(*i6ab0f57415f0f91846673c7144401d1ecb437b2ed9583ff3a324fce02d79d6ce.SectionGroupsRequestBuilder) {
    return i6ab0f57415f0f91846673c7144401d1ecb437b2ed9583ff3a324fce02d79d6ce.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.onenote.sectionGroups.item collection
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*ief43f129f43e4e2e00209d40c6ab20f225d4a6be2f614f3de5c37ad2d36f6397.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return ief43f129f43e4e2e00209d40c6ab20f225d4a6be2f614f3de5c37ad2d36f6397.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Sections()(*i89a64f4378390a1ff2c4709e25458a2484ce3ed41135ba5fb5e45f1f2c8feecd.SectionsRequestBuilder) {
    return i89a64f4378390a1ff2c4709e25458a2484ce3ed41135ba5fb5e45f1f2c8feecd.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.onenote.sections.item collection
func (m *OnenoteRequestBuilder) SectionsById(id string)(*i2fd4361716803984f5766f676870e98432266f948b1f5e8bffb51c4c87176198.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i2fd4361716803984f5766f676870e98432266f948b1f5e8bffb51c4c87176198.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
