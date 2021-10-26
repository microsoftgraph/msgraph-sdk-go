package onenote

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
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

// Builds and executes requests for operations under \groups\{group-id}\onenote
type OnenoteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Read-only.
type OnenoteRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new OnenoteRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOnenoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteRequestBuilder) {
    m := &OnenoteRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/groups/{group_id}/onenote{?select,expand}";
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
//  - h : Request headers
//  - o : Request options
func (m *OnenoteRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *OnenoteRequestBuilder) CreateGetRequestInformation(q func (value *OnenoteRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(OnenoteRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *OnenoteRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Onenote, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *OnenoteRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
// Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *OnenoteRequestBuilder) Get(q func (value *OnenoteRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Onenote, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOnenote() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Onenote), nil
}
func (m *OnenoteRequestBuilder) Notebooks()(*i3e5ab660d6ac3db9593f97ea896aa48566cf72944732b076c4a9d061912a6c04.NotebooksRequestBuilder) {
    return i3e5ab660d6ac3db9593f97ea896aa48566cf72944732b076c4a9d061912a6c04.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.notebooks.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*icf88544b86cbd77f33fc028168410031e602c247e6524b0ec76c9ec5301c1b8c.NotebookRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook_id"] = id
    }
    return icf88544b86cbd77f33fc028168410031e602c247e6524b0ec76c9ec5301c1b8c.NewNotebookRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Operations()(*ic4302b238aa48da7d88cf2ada9f17f6eed828e36e8f54efabf0292996911ef1a.OperationsRequestBuilder) {
    return ic4302b238aa48da7d88cf2ada9f17f6eed828e36e8f54efabf0292996911ef1a.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.operations.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i15f518181fda87a8cdf2c10fc20b18c70bacacca5e08a1dc09a45eb124da89ad.OnenoteOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation_id"] = id
    }
    return i15f518181fda87a8cdf2c10fc20b18c70bacacca5e08a1dc09a45eb124da89ad.NewOnenoteOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Pages()(*id7a0200daa8eeafd045771e775fa3934eb9a171334b734bc94cdc8222f05a916.PagesRequestBuilder) {
    return id7a0200daa8eeafd045771e775fa3934eb9a171334b734bc94cdc8222f05a916.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.pages.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) PagesById(id string)(*i0a0eb5562b0dd916c680d0a554cb25b167a2ebe18945dc353bdaa39a7ec4fa14.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i0a0eb5562b0dd916c680d0a554cb25b167a2ebe18945dc353bdaa39a7ec4fa14.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *OnenoteRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Onenote, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *OnenoteRequestBuilder) Resources()(*i56afdd94701540796ba18ff7485293b0edb3d0ff46c482e936c1a66059939a7d.ResourcesRequestBuilder) {
    return i56afdd94701540796ba18ff7485293b0edb3d0ff46c482e936c1a66059939a7d.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.resources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*ic982193c51308e3c875bd915544c104ad98988f51026864216f9494619169091.OnenoteResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource_id"] = id
    }
    return ic982193c51308e3c875bd915544c104ad98988f51026864216f9494619169091.NewOnenoteResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) SectionGroups()(*i392b78ec87e6862ef0bb0965913ffdca97eaf69084704fd05b09f2a82f718e11.SectionGroupsRequestBuilder) {
    return i392b78ec87e6862ef0bb0965913ffdca97eaf69084704fd05b09f2a82f718e11.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.sectionGroups.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*i032333af94d0562d937990b412828a992cf6dcd9300441d921dd966122896df0.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i032333af94d0562d937990b412828a992cf6dcd9300441d921dd966122896df0.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Sections()(*id5409a02c0495e9c5e6b38f89b83f82f93291dc5e79dd29c377d082d3a67b973.SectionsRequestBuilder) {
    return id5409a02c0495e9c5e6b38f89b83f82f93291dc5e79dd29c377d082d3a67b973.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.sections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) SectionsById(id string)(*i18fb32726d606974074adcc29ef533aa0b798a82757b5717d4487c0bec214a39.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i18fb32726d606974074adcc29ef533aa0b798a82757b5717d4487c0bec214a39.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
