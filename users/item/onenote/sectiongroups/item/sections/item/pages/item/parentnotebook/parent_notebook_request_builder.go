package parentnotebook

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1f393e7c704baee7ed544ff519b829ff6e72d511b2c1d9d72a363d2efea4cb2a "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/sectiongroups/item/sections/item/pages/item/parentnotebook/sectiongroups"
    i4fff192021fd9819f3ff0e65f610e55fa9c6ac1f934bf5d4ad3ce8115756736b "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/sectiongroups/item/sections/item/pages/item/parentnotebook/sections"
    i84f50802549187236f9b8e183ca832e601c670242da31652ae06b2343eccb272 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/sectiongroups/item/sections/item/pages/item/parentnotebook/copynotebook"
    i2ebd372bbba347419bd8e99869f799828743af3dbe0219e70d9dc7cdaf9b14c7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/sectiongroups/item/sections/item/pages/item/parentnotebook/sectiongroups/item"
    i5d9a565f2c64e009a64a20a93e50cb14b1b6b13e187dd5d027c7af2ea9f65d34 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/sectiongroups/item/sections/item/pages/item/parentnotebook/sections/item"
)

// Builds and executes requests for operations under \users\{user-id}\onenote\sectionGroups\{sectionGroup-id}\sections\{onenoteSection-id}\pages\{onenotePage-id}\parentNotebook
type ParentNotebookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ParentNotebookRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ParentNotebookRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ParentNotebookRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The notebook that contains the page.  Read-only.
type ParentNotebookRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ParentNotebookRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Notebook;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new ParentNotebookRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewParentNotebookRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentNotebookRequestBuilder) {
    m := &ParentNotebookRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/onenote/sectionGroups/{sectionGroup_id}/sections/{onenoteSection_id}/pages/{onenotePage_id}/parentNotebook{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ParentNotebookRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewParentNotebookRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentNotebookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParentNotebookRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ParentNotebookRequestBuilder) CopyNotebook()(*i84f50802549187236f9b8e183ca832e601c670242da31652ae06b2343eccb272.CopyNotebookRequestBuilder) {
    return i84f50802549187236f9b8e183ca832e601c670242da31652ae06b2343eccb272.NewCopyNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The notebook that contains the page.  Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentNotebookRequestBuilder) CreateDeleteRequestInformation(options *ParentNotebookRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The notebook that contains the page.  Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentNotebookRequestBuilder) CreateGetRequestInformation(options *ParentNotebookRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The notebook that contains the page.  Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentNotebookRequestBuilder) CreatePatchRequestInformation(options *ParentNotebookRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The notebook that contains the page.  Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentNotebookRequestBuilder) Delete(options *ParentNotebookRequestBuilderDeleteOptions)(error) {
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
// The notebook that contains the page.  Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentNotebookRequestBuilder) Get(options *ParentNotebookRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Notebook, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewNotebook() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Notebook), nil
}
// The notebook that contains the page.  Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentNotebookRequestBuilder) Patch(options *ParentNotebookRequestBuilderPatchOptions)(error) {
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
func (m *ParentNotebookRequestBuilder) SectionGroups()(*i1f393e7c704baee7ed544ff519b829ff6e72d511b2c1d9d72a363d2efea4cb2a.SectionGroupsRequestBuilder) {
    return i1f393e7c704baee7ed544ff519b829ff6e72d511b2c1d9d72a363d2efea4cb2a.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.users.item.onenote.sectionGroups.item.sections.item.pages.item.parentNotebook.sectionGroups.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ParentNotebookRequestBuilder) SectionGroupsById(id string)(*i2ebd372bbba347419bd8e99869f799828743af3dbe0219e70d9dc7cdaf9b14c7.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id1"] = id
    }
    return i2ebd372bbba347419bd8e99869f799828743af3dbe0219e70d9dc7cdaf9b14c7.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ParentNotebookRequestBuilder) Sections()(*i4fff192021fd9819f3ff0e65f610e55fa9c6ac1f934bf5d4ad3ce8115756736b.SectionsRequestBuilder) {
    return i4fff192021fd9819f3ff0e65f610e55fa9c6ac1f934bf5d4ad3ce8115756736b.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.users.item.onenote.sectionGroups.item.sections.item.pages.item.parentNotebook.sections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ParentNotebookRequestBuilder) SectionsById(id string)(*i5d9a565f2c64e009a64a20a93e50cb14b1b6b13e187dd5d027c7af2ea9f65d34.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id1"] = id
    }
    return i5d9a565f2c64e009a64a20a93e50cb14b1b6b13e187dd5d027c7af2ea9f65d34.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
