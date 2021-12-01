package parentsectiongroup

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i403140d9240d2ca0c26468d4f845b04a9cfc07fb733b2069eb35ed6420df4e5d "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/pages/item/parentsection/parentsectiongroup/sections"
    i70572942e4e26e792b9abc42d8bc31a1bcca1e13ca65ef659d066a5c75c04240 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/pages/item/parentsection/parentsectiongroup/sectiongroups"
    i9ddad1432f81a4ddcd2c39cde80038892fede4ab32b3a9eb39efaa237d2eb60f "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/pages/item/parentsection/parentsectiongroup/parentnotebook"
    i368717425ffc1da9f8a2b8325a5237916cb553ed28f3bec4aea68c82d0e90dbd "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/pages/item/parentsection/parentsectiongroup/sectiongroups/item"
    ie39e2121bdc448dff7d5ee49d99641f2f6320492542c54d790ff62fd91056010 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/pages/item/parentsection/parentsectiongroup/sections/item"
)

// ParentSectionGroupRequestBuilder builds and executes requests for operations under \users\{user-id}\onenote\pages\{onenotePage-id}\parentSection\parentSectionGroup
type ParentSectionGroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ParentSectionGroupRequestBuilderDeleteOptions options for Delete
type ParentSectionGroupRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ParentSectionGroupRequestBuilderGetOptions options for Get
type ParentSectionGroupRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ParentSectionGroupRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ParentSectionGroupRequestBuilderGetQueryParameters the section group that contains the section.  Read-only.
type ParentSectionGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ParentSectionGroupRequestBuilderPatchOptions options for Patch
type ParentSectionGroupRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SectionGroup;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewParentSectionGroupRequestBuilderInternal instantiates a new ParentSectionGroupRequestBuilder and sets the default values.
func NewParentSectionGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentSectionGroupRequestBuilder) {
    m := &ParentSectionGroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/onenote/pages/{onenotePage_id}/parentSection/parentSectionGroup{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewParentSectionGroupRequestBuilder instantiates a new ParentSectionGroupRequestBuilder and sets the default values.
func NewParentSectionGroupRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentSectionGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParentSectionGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the section group that contains the section.  Read-only.
func (m *ParentSectionGroupRequestBuilder) CreateDeleteRequestInformation(options *ParentSectionGroupRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the section group that contains the section.  Read-only.
func (m *ParentSectionGroupRequestBuilder) CreateGetRequestInformation(options *ParentSectionGroupRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the section group that contains the section.  Read-only.
func (m *ParentSectionGroupRequestBuilder) CreatePatchRequestInformation(options *ParentSectionGroupRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the section group that contains the section.  Read-only.
func (m *ParentSectionGroupRequestBuilder) Delete(options *ParentSectionGroupRequestBuilderDeleteOptions)(error) {
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
// Get the section group that contains the section.  Read-only.
func (m *ParentSectionGroupRequestBuilder) Get(options *ParentSectionGroupRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SectionGroup, error) {
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
func (m *ParentSectionGroupRequestBuilder) ParentNotebook()(*i9ddad1432f81a4ddcd2c39cde80038892fede4ab32b3a9eb39efaa237d2eb60f.ParentNotebookRequestBuilder) {
    return i9ddad1432f81a4ddcd2c39cde80038892fede4ab32b3a9eb39efaa237d2eb60f.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParentSectionGroupRequestBuilder) ParentSectionGroup()(*ParentSectionGroupRequestBuilder) {
    return NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the section group that contains the section.  Read-only.
func (m *ParentSectionGroupRequestBuilder) Patch(options *ParentSectionGroupRequestBuilderPatchOptions)(error) {
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
func (m *ParentSectionGroupRequestBuilder) SectionGroups()(*i70572942e4e26e792b9abc42d8bc31a1bcca1e13ca65ef659d066a5c75c04240.SectionGroupsRequestBuilder) {
    return i70572942e4e26e792b9abc42d8bc31a1bcca1e13ca65ef659d066a5c75c04240.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.onenote.pages.item.parentSection.parentSectionGroup.sectionGroups.item collection
func (m *ParentSectionGroupRequestBuilder) SectionGroupsById(id string)(*i368717425ffc1da9f8a2b8325a5237916cb553ed28f3bec4aea68c82d0e90dbd.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i368717425ffc1da9f8a2b8325a5237916cb553ed28f3bec4aea68c82d0e90dbd.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ParentSectionGroupRequestBuilder) Sections()(*i403140d9240d2ca0c26468d4f845b04a9cfc07fb733b2069eb35ed6420df4e5d.SectionsRequestBuilder) {
    return i403140d9240d2ca0c26468d4f845b04a9cfc07fb733b2069eb35ed6420df4e5d.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.onenote.pages.item.parentSection.parentSectionGroup.sections.item collection
func (m *ParentSectionGroupRequestBuilder) SectionsById(id string)(*ie39e2121bdc448dff7d5ee49d99641f2f6320492542c54d790ff62fd91056010.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return ie39e2121bdc448dff7d5ee49d99641f2f6320492542c54d790ff62fd91056010.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
