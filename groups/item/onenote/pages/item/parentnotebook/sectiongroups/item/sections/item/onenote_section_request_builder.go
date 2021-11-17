package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0522b530079a2d60e704824d3ca5a08262f82116c8a6ebae6c3ab2e18e1aa7ea "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/sections/item/pages"
    i0ba9d4676738dc937b70e9526b277207e1b0258b28d8732400fe47c0dd7f3acb "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/sections/item/copytonotebook"
    i3c40889666ae177dc72947e2d68501d7a0d1ad729468b5ade80df88ef951d4f7 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/sections/item/parentnotebook"
    ib3c2f84c73b01af0f324b10dd5f2b1c6f85906c29eb4b886862da7588d1ccd99 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/sections/item/copytosectiongroup"
    id17ee6de034982251a79a9d2533d4706e7e457fe24d3d29e3eb2ec5f846200d7 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/sections/item/parentsectiongroup"
    ia4ad9bd562460115906a9af5b391da46860e19b51a4179c99e69b7afcf9340fb "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/sections/item/pages/item"
)

// Builds and executes requests for operations under \groups\{group-id}\onenote\pages\{onenotePage-id}\parentNotebook\sectionGroups\{sectionGroup-id}\sections\{onenoteSection-id}
type OnenoteSectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type OnenoteSectionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type OnenoteSectionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnenoteSectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The sections in the section group. Read-only. Nullable.
type OnenoteSectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type OnenoteSectionRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenoteSection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new OnenoteSectionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOnenoteSectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteSectionRequestBuilder) {
    m := &OnenoteSectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/onenote/pages/{onenotePage_id}/parentNotebook/sectionGroups/{sectionGroup_id}/sections/{onenoteSection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new OnenoteSectionRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOnenoteSectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteSectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteSectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnenoteSectionRequestBuilder) CopyToNotebook()(*i0ba9d4676738dc937b70e9526b277207e1b0258b28d8732400fe47c0dd7f3acb.CopyToNotebookRequestBuilder) {
    return i0ba9d4676738dc937b70e9526b277207e1b0258b28d8732400fe47c0dd7f3acb.NewCopyToNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteSectionRequestBuilder) CopyToSectionGroup()(*ib3c2f84c73b01af0f324b10dd5f2b1c6f85906c29eb4b886862da7588d1ccd99.CopyToSectionGroupRequestBuilder) {
    return ib3c2f84c73b01af0f324b10dd5f2b1c6f85906c29eb4b886862da7588d1ccd99.NewCopyToSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The sections in the section group. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OnenoteSectionRequestBuilder) CreateDeleteRequestInformation(options *OnenoteSectionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The sections in the section group. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OnenoteSectionRequestBuilder) CreateGetRequestInformation(options *OnenoteSectionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The sections in the section group. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OnenoteSectionRequestBuilder) CreatePatchRequestInformation(options *OnenoteSectionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The sections in the section group. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OnenoteSectionRequestBuilder) Delete(options *OnenoteSectionRequestBuilderDeleteOptions)(error) {
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
// The sections in the section group. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OnenoteSectionRequestBuilder) Get(options *OnenoteSectionRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenoteSection, error) {
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
func (m *OnenoteSectionRequestBuilder) Pages()(*i0522b530079a2d60e704824d3ca5a08262f82116c8a6ebae6c3ab2e18e1aa7ea.PagesRequestBuilder) {
    return i0522b530079a2d60e704824d3ca5a08262f82116c8a6ebae6c3ab2e18e1aa7ea.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.pages.item.parentNotebook.sectionGroups.item.sections.item.pages.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteSectionRequestBuilder) PagesById(id string)(*ia4ad9bd562460115906a9af5b391da46860e19b51a4179c99e69b7afcf9340fb.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id1"] = id
    }
    return ia4ad9bd562460115906a9af5b391da46860e19b51a4179c99e69b7afcf9340fb.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteSectionRequestBuilder) ParentNotebook()(*i3c40889666ae177dc72947e2d68501d7a0d1ad729468b5ade80df88ef951d4f7.ParentNotebookRequestBuilder) {
    return i3c40889666ae177dc72947e2d68501d7a0d1ad729468b5ade80df88ef951d4f7.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteSectionRequestBuilder) ParentSectionGroup()(*id17ee6de034982251a79a9d2533d4706e7e457fe24d3d29e3eb2ec5f846200d7.ParentSectionGroupRequestBuilder) {
    return id17ee6de034982251a79a9d2533d4706e7e457fe24d3d29e3eb2ec5f846200d7.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The sections in the section group. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OnenoteSectionRequestBuilder) Patch(options *OnenoteSectionRequestBuilderPatchOptions)(error) {
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
