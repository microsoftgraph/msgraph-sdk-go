package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i4f8dd6fbbbd8e5eac4c018e252d494b346d6344a7e2f2cf4d26e0787cfc50d17 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sectiongroups/item/sections"
    ib4e3139e711c23865c0e217671361df83d26fdce68565cc23cedb903bb5b8eff "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sectiongroups/item/parentsectiongroup"
    ibb6e449347a883aed7ef7d46e6fde8269b6148c357bcb3eee0fea66f175d3b29 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sectiongroups/item/sectiongroups"
    iecb427cf10e33fed277f8759f6db9ebd96d9f26a571e9a50f5f0c15563ed1f94 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sectiongroups/item/parentnotebook"
    i572c7a753a0704b855af1cd5e3531b9eaa2442243b3635858b7a519c4d054813 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sectiongroups/item/sections/item"
    i8533efee27114835d72ec3a21661ae353e4b65d8107ece04573cc5a2bac2c299 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sectiongroups/item/sectiongroups/item"
)

// SectionGroupRequestBuilder builds and executes requests for operations under \users\{user-id}\onenote\notebooks\{notebook-id}\sectionGroups\{sectionGroup-id}
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
    Select_escaped []string;
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
    m.urlTemplate = "{+baseurl}/users/{user_id}/onenote/notebooks/{notebook_id}/sectionGroups/{sectionGroup_id}{?select,expand}";
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
func (m *SectionGroupRequestBuilder) ParentNotebook()(*iecb427cf10e33fed277f8759f6db9ebd96d9f26a571e9a50f5f0c15563ed1f94.ParentNotebookRequestBuilder) {
    return iecb427cf10e33fed277f8759f6db9ebd96d9f26a571e9a50f5f0c15563ed1f94.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SectionGroupRequestBuilder) ParentSectionGroup()(*ib4e3139e711c23865c0e217671361df83d26fdce68565cc23cedb903bb5b8eff.ParentSectionGroupRequestBuilder) {
    return ib4e3139e711c23865c0e217671361df83d26fdce68565cc23cedb903bb5b8eff.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *SectionGroupRequestBuilder) SectionGroups()(*ibb6e449347a883aed7ef7d46e6fde8269b6148c357bcb3eee0fea66f175d3b29.SectionGroupsRequestBuilder) {
    return ibb6e449347a883aed7ef7d46e6fde8269b6148c357bcb3eee0fea66f175d3b29.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.onenote.notebooks.item.sectionGroups.item.sectionGroups.item collection
func (m *SectionGroupRequestBuilder) SectionGroupsById(id string)(*i8533efee27114835d72ec3a21661ae353e4b65d8107ece04573cc5a2bac2c299.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id1"] = id
    }
    return i8533efee27114835d72ec3a21661ae353e4b65d8107ece04573cc5a2bac2c299.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SectionGroupRequestBuilder) Sections()(*i4f8dd6fbbbd8e5eac4c018e252d494b346d6344a7e2f2cf4d26e0787cfc50d17.SectionsRequestBuilder) {
    return i4f8dd6fbbbd8e5eac4c018e252d494b346d6344a7e2f2cf4d26e0787cfc50d17.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.onenote.notebooks.item.sectionGroups.item.sections.item collection
func (m *SectionGroupRequestBuilder) SectionsById(id string)(*i572c7a753a0704b855af1cd5e3531b9eaa2442243b3635858b7a519c4d054813.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i572c7a753a0704b855af1cd5e3531b9eaa2442243b3635858b7a519c4d054813.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
