package parentsectiongroup

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i9a843897b1fb3a57d124d108800320e52be3032be7574454b0f2aa90dc51473b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sections/item/parentsectiongroup/parentnotebook"
    id6fbe22e0bfe57670b9d5c58a0749ad13935375c357442a5d966f47e487a3797 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sections/item/parentsectiongroup/sectiongroups"
    ie614bbbfadcc39c98723b67ff3bfa79df1d04523b2d89c504ba7c7c0ae110bf7 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sections/item/parentsectiongroup/sections"
    i4e09baf5a51ded2625cd42e5c813e1e1935cd7c946ad04074c038dc8b9708ae3 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sections/item/parentsectiongroup/sections/item"
    ifd3defddd2ce7046aa1f60101051a5e952ca0620aaa0ff18ac6d5ef23879cb17 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sections/item/parentsectiongroup/sectiongroups/item"
)

// Builds and executes requests for operations under \groups\{group-id}\onenote\pages\{onenotePage-id}\parentNotebook\sections\{onenoteSection-id}\parentSectionGroup
type ParentSectionGroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ParentSectionGroupRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
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
// The section group that contains the section.  Read-only.
type ParentSectionGroupRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
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
// Instantiates a new ParentSectionGroupRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewParentSectionGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentSectionGroupRequestBuilder) {
    m := &ParentSectionGroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/onenote/pages/{onenotePage_id}/parentNotebook/sections/{onenoteSection_id}/parentSectionGroup{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ParentSectionGroupRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewParentSectionGroupRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentSectionGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParentSectionGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// The section group that contains the section.  Read-only.
// Parameters:
//  - options : Options for the request
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
// The section group that contains the section.  Read-only.
// Parameters:
//  - options : Options for the request
func (m *ParentSectionGroupRequestBuilder) CreateGetRequestInformation(options *ParentSectionGroupRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The section group that contains the section.  Read-only.
// Parameters:
//  - options : Options for the request
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
// The section group that contains the section.  Read-only.
// Parameters:
//  - options : Options for the request
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
// The section group that contains the section.  Read-only.
// Parameters:
//  - options : Options for the request
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
func (m *ParentSectionGroupRequestBuilder) ParentNotebook()(*i9a843897b1fb3a57d124d108800320e52be3032be7574454b0f2aa90dc51473b.ParentNotebookRequestBuilder) {
    return i9a843897b1fb3a57d124d108800320e52be3032be7574454b0f2aa90dc51473b.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParentSectionGroupRequestBuilder) ParentSectionGroup()(*ParentSectionGroupRequestBuilder) {
    return NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The section group that contains the section.  Read-only.
// Parameters:
//  - options : Options for the request
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
func (m *ParentSectionGroupRequestBuilder) SectionGroups()(*id6fbe22e0bfe57670b9d5c58a0749ad13935375c357442a5d966f47e487a3797.SectionGroupsRequestBuilder) {
    return id6fbe22e0bfe57670b9d5c58a0749ad13935375c357442a5d966f47e487a3797.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.groups.item.onenote.pages.item.parentNotebook.sections.item.parentSectionGroup.sectionGroups.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ParentSectionGroupRequestBuilder) SectionGroupsById(id string)(*ifd3defddd2ce7046aa1f60101051a5e952ca0620aaa0ff18ac6d5ef23879cb17.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return ifd3defddd2ce7046aa1f60101051a5e952ca0620aaa0ff18ac6d5ef23879cb17.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ParentSectionGroupRequestBuilder) Sections()(*ie614bbbfadcc39c98723b67ff3bfa79df1d04523b2d89c504ba7c7c0ae110bf7.SectionsRequestBuilder) {
    return ie614bbbfadcc39c98723b67ff3bfa79df1d04523b2d89c504ba7c7c0ae110bf7.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.groups.item.onenote.pages.item.parentNotebook.sections.item.parentSectionGroup.sections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ParentSectionGroupRequestBuilder) SectionsById(id string)(*i4e09baf5a51ded2625cd42e5c813e1e1935cd7c946ad04074c038dc8b9708ae3.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id1"] = id
    }
    return i4e09baf5a51ded2625cd42e5c813e1e1935cd7c946ad04074c038dc8b9708ae3.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
