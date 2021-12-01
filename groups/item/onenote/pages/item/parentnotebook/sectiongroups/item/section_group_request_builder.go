package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i2d4405d61073097820d21feaeb8cfb4b00b8b97ca196c8d61adaaa139a98553e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/sectiongroups"
    ia63f4929dba5b38acb07057cefc6c982492d9f0f8bd23fbf594b9af33e358615 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/sections"
    iabd06dfa023f86494ccd0adf4612084319849eebeafc8626f13f6e8e3a901114 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/parentnotebook"
    ie54d3ae495ee37efadd6e2191a33ee0ce5b6fac69255f0ba77b835f821831d1a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/parentsectiongroup"
    i1d1d7ffa0230017a6a201b1784298e9991cc81e8f030794f7db0f9755ec106ef "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/sections/item"
    i29f4f1b45400549b046daaacd804ee1174916b917848e3c7a941a72a9026979d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/sectiongroups/item"
)

// SectionGroupRequestBuilder builds and executes requests for operations under \groups\{group-id}\onenote\pages\{onenotePage-id}\parentNotebook\sectionGroups\{sectionGroup-id}
type SectionGroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SectionGroupRequestBuilderDeleteOptions options for Delete
type SectionGroupRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SectionGroupRequestBuilderGetOptions options for Get
type SectionGroupRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SectionGroupRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SectionGroupRequestBuilderGetQueryParameters the section groups in the notebook. Read-only. Nullable.
type SectionGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SectionGroupRequestBuilderPatchOptions options for Patch
type SectionGroupRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SectionGroup;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSectionGroupRequestBuilderInternal instantiates a new SectionGroupRequestBuilder and sets the default values.
func NewSectionGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SectionGroupRequestBuilder) {
    m := &SectionGroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/onenote/pages/{onenotePage_id}/parentNotebook/sectionGroups/{sectionGroup_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSectionGroupRequestBuilder instantiates a new SectionGroupRequestBuilder and sets the default values.
func NewSectionGroupRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SectionGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSectionGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the section groups in the notebook. Read-only. Nullable.
func (m *SectionGroupRequestBuilder) CreateDeleteRequestInformation(options *SectionGroupRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the section groups in the notebook. Read-only. Nullable.
func (m *SectionGroupRequestBuilder) CreateGetRequestInformation(options *SectionGroupRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the section groups in the notebook. Read-only. Nullable.
func (m *SectionGroupRequestBuilder) CreatePatchRequestInformation(options *SectionGroupRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the section groups in the notebook. Read-only. Nullable.
func (m *SectionGroupRequestBuilder) Delete(options *SectionGroupRequestBuilderDeleteOptions)(error) {
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
// Get the section groups in the notebook. Read-only. Nullable.
func (m *SectionGroupRequestBuilder) Get(options *SectionGroupRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SectionGroup, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSectionGroup() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SectionGroup), nil
}
func (m *SectionGroupRequestBuilder) ParentNotebook()(*iabd06dfa023f86494ccd0adf4612084319849eebeafc8626f13f6e8e3a901114.ParentNotebookRequestBuilder) {
    return iabd06dfa023f86494ccd0adf4612084319849eebeafc8626f13f6e8e3a901114.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SectionGroupRequestBuilder) ParentSectionGroup()(*ie54d3ae495ee37efadd6e2191a33ee0ce5b6fac69255f0ba77b835f821831d1a.ParentSectionGroupRequestBuilder) {
    return ie54d3ae495ee37efadd6e2191a33ee0ce5b6fac69255f0ba77b835f821831d1a.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the section groups in the notebook. Read-only. Nullable.
func (m *SectionGroupRequestBuilder) Patch(options *SectionGroupRequestBuilderPatchOptions)(error) {
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
func (m *SectionGroupRequestBuilder) SectionGroups()(*i2d4405d61073097820d21feaeb8cfb4b00b8b97ca196c8d61adaaa139a98553e.SectionGroupsRequestBuilder) {
    return i2d4405d61073097820d21feaeb8cfb4b00b8b97ca196c8d61adaaa139a98553e.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.pages.item.parentNotebook.sectionGroups.item.sectionGroups.item collection
func (m *SectionGroupRequestBuilder) SectionGroupsById(id string)(*i29f4f1b45400549b046daaacd804ee1174916b917848e3c7a941a72a9026979d.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id1"] = id
    }
    return i29f4f1b45400549b046daaacd804ee1174916b917848e3c7a941a72a9026979d.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SectionGroupRequestBuilder) Sections()(*ia63f4929dba5b38acb07057cefc6c982492d9f0f8bd23fbf594b9af33e358615.SectionsRequestBuilder) {
    return ia63f4929dba5b38acb07057cefc6c982492d9f0f8bd23fbf594b9af33e358615.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.pages.item.parentNotebook.sectionGroups.item.sections.item collection
func (m *SectionGroupRequestBuilder) SectionsById(id string)(*i1d1d7ffa0230017a6a201b1784298e9991cc81e8f030794f7db0f9755ec106ef.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i1d1d7ffa0230017a6a201b1784298e9991cc81e8f030794f7db0f9755ec106ef.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
