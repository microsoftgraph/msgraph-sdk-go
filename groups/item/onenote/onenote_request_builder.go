package onenote

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i392b78ec87e6862ef0bb0965913ffdca97eaf69084704fd05b09f2a82f718e11 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/sectiongroups"
    i3e5ab660d6ac3db9593f97ea896aa48566cf72944732b076c4a9d061912a6c04 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/notebooks"
    i56afdd94701540796ba18ff7485293b0edb3d0ff46c482e936c1a66059939a7d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/resources"
    ic4302b238aa48da7d88cf2ada9f17f6eed828e36e8f54efabf0292996911ef1a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/operations"
    id5409a02c0495e9c5e6b38f89b83f82f93291dc5e79dd29c377d082d3a67b973 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/sections"
    id7a0200daa8eeafd045771e775fa3934eb9a171334b734bc94cdc8222f05a916 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages"
    i032333af94d0562d937990b412828a992cf6dcd9300441d921dd966122896df0 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/sectiongroups/item"
    i0a0eb5562b0dd916c680d0a554cb25b167a2ebe18945dc353bdaa39a7ec4fa14 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item"
    i15f518181fda87a8cdf2c10fc20b18c70bacacca5e08a1dc09a45eb124da89ad "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/operations/item"
    i18fb32726d606974074adcc29ef533aa0b798a82757b5717d4487c0bec214a39 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/sections/item"
    ic982193c51308e3c875bd915544c104ad98988f51026864216f9494619169091 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/resources/item"
    icf88544b86cbd77f33fc028168410031e602c247e6524b0ec76c9ec5301c1b8c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/notebooks/item"
)

// OnenoteRequestBuilder provides operations to manage the onenote property of the microsoft.graph.group entity.
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
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Onenoteable;
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
    m.urlTemplate = "{+baseurl}/groups/{group_id}/onenote{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteRequestBuilder instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property onenote for groups
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
// CreatePatchRequestInformation update the navigation property onenote in groups
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
// Delete delete navigation property onenote for groups
func (m *OnenoteRequestBuilder) Delete(options *OnenoteRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read-only.
func (m *OnenoteRequestBuilder) Get(options *OnenoteRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Onenoteable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateOnenoteFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Onenoteable), nil
}
func (m *OnenoteRequestBuilder) Notebooks()(*i3e5ab660d6ac3db9593f97ea896aa48566cf72944732b076c4a9d061912a6c04.NotebooksRequestBuilder) {
    return i3e5ab660d6ac3db9593f97ea896aa48566cf72944732b076c4a9d061912a6c04.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotebooksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.notebooks.item collection
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*icf88544b86cbd77f33fc028168410031e602c247e6524b0ec76c9ec5301c1b8c.NotebookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook_id"] = id
    }
    return icf88544b86cbd77f33fc028168410031e602c247e6524b0ec76c9ec5301c1b8c.NewNotebookItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Operations()(*ic4302b238aa48da7d88cf2ada9f17f6eed828e36e8f54efabf0292996911ef1a.OperationsRequestBuilder) {
    return ic4302b238aa48da7d88cf2ada9f17f6eed828e36e8f54efabf0292996911ef1a.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.operations.item collection
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i15f518181fda87a8cdf2c10fc20b18c70bacacca5e08a1dc09a45eb124da89ad.OnenoteOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation_id"] = id
    }
    return i15f518181fda87a8cdf2c10fc20b18c70bacacca5e08a1dc09a45eb124da89ad.NewOnenoteOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Pages()(*id7a0200daa8eeafd045771e775fa3934eb9a171334b734bc94cdc8222f05a916.PagesRequestBuilder) {
    return id7a0200daa8eeafd045771e775fa3934eb9a171334b734bc94cdc8222f05a916.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.pages.item collection
func (m *OnenoteRequestBuilder) PagesById(id string)(*i0a0eb5562b0dd916c680d0a554cb25b167a2ebe18945dc353bdaa39a7ec4fa14.OnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i0a0eb5562b0dd916c680d0a554cb25b167a2ebe18945dc353bdaa39a7ec4fa14.NewOnenotePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property onenote in groups
func (m *OnenoteRequestBuilder) Patch(options *OnenoteRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *OnenoteRequestBuilder) Resources()(*i56afdd94701540796ba18ff7485293b0edb3d0ff46c482e936c1a66059939a7d.ResourcesRequestBuilder) {
    return i56afdd94701540796ba18ff7485293b0edb3d0ff46c482e936c1a66059939a7d.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.resources.item collection
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*ic982193c51308e3c875bd915544c104ad98988f51026864216f9494619169091.OnenoteResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource_id"] = id
    }
    return ic982193c51308e3c875bd915544c104ad98988f51026864216f9494619169091.NewOnenoteResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) SectionGroups()(*i392b78ec87e6862ef0bb0965913ffdca97eaf69084704fd05b09f2a82f718e11.SectionGroupsRequestBuilder) {
    return i392b78ec87e6862ef0bb0965913ffdca97eaf69084704fd05b09f2a82f718e11.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.sectionGroups.item collection
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*i032333af94d0562d937990b412828a992cf6dcd9300441d921dd966122896df0.SectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i032333af94d0562d937990b412828a992cf6dcd9300441d921dd966122896df0.NewSectionGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Sections()(*id5409a02c0495e9c5e6b38f89b83f82f93291dc5e79dd29c377d082d3a67b973.SectionsRequestBuilder) {
    return id5409a02c0495e9c5e6b38f89b83f82f93291dc5e79dd29c377d082d3a67b973.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.sections.item collection
func (m *OnenoteRequestBuilder) SectionsById(id string)(*i18fb32726d606974074adcc29ef533aa0b798a82757b5717d4487c0bec214a39.OnenoteSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i18fb32726d606974074adcc29ef533aa0b798a82757b5717d4487c0bec214a39.NewOnenoteSectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
