package onenote

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
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

// OnenoteRequestBuilder provides operations to manage the onenote property of the microsoft.graph.site entity.
type OnenoteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnenoteRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenoteRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnenoteRequestBuilderGetQueryParameters calls the OneNote service for notebook related operations.
type OnenoteRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnenoteRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenoteRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnenoteRequestBuilderGetQueryParameters
}
// OnenoteRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenoteRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnenoteRequestBuilderInternal instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteRequestBuilder) {
    m := &OnenoteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/onenote{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteRequestBuilder instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property onenote for sites
func (m *OnenoteRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property onenote for sites
func (m *OnenoteRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *OnenoteRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation calls the OneNote service for notebook related operations.
func (m *OnenoteRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration calls the OneNote service for notebook related operations.
func (m *OnenoteRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OnenoteRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property onenote in sites
func (m *OnenoteRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Onenoteable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property onenote in sites
func (m *OnenoteRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Onenoteable, requestConfiguration *OnenoteRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property onenote for sites
func (m *OnenoteRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property onenote for sites
func (m *OnenoteRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *OnenoteRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get calls the OneNote service for notebook related operations.
func (m *OnenoteRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Onenoteable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler calls the OneNote service for notebook related operations.
func (m *OnenoteRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *OnenoteRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Onenoteable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenoteFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Onenoteable), nil
}
// Notebooks the notebooks property
func (m *OnenoteRequestBuilder) Notebooks()(*id056e9dcc48f2e5766390396802781a6f6e4a85d729a1d7f65531745ed03bd15.NotebooksRequestBuilder) {
    return id056e9dcc48f2e5766390396802781a6f6e4a85d729a1d7f65531745ed03bd15.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotebooksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.notebooks.item collection
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*i2ba021dd48264681e34929a50215efa4ffab4a38e1e781d5129ff35bc76ada21.NotebookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook%2Did"] = id
    }
    return i2ba021dd48264681e34929a50215efa4ffab4a38e1e781d5129ff35bc76ada21.NewNotebookItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *OnenoteRequestBuilder) Operations()(*ifedb0f50ef2ef408f7a3fbfc146cacec13291103da0bba733f4b2757f5d4de71.OperationsRequestBuilder) {
    return ifedb0f50ef2ef408f7a3fbfc146cacec13291103da0bba733f4b2757f5d4de71.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.operations.item collection
func (m *OnenoteRequestBuilder) OperationsById(id string)(*ie9c35a4053aa937b7cbd985746f2c474e4823ad0293ce536d11acb1d37b1f790.OnenoteOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation%2Did"] = id
    }
    return ie9c35a4053aa937b7cbd985746f2c474e4823ad0293ce536d11acb1d37b1f790.NewOnenoteOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Pages the pages property
func (m *OnenoteRequestBuilder) Pages()(*i025d995a3eab0912bb44d350ddf5d7a3b0aeb9332e0629260ad912238bc917e4.PagesRequestBuilder) {
    return i025d995a3eab0912bb44d350ddf5d7a3b0aeb9332e0629260ad912238bc917e4.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.pages.item collection
func (m *OnenoteRequestBuilder) PagesById(id string)(*ie1466039fa513108cc7b0f3d5f87933999e122789d21aa1fd0e611cacc94cd21.OnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage%2Did"] = id
    }
    return ie1466039fa513108cc7b0f3d5f87933999e122789d21aa1fd0e611cacc94cd21.NewOnenotePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property onenote in sites
func (m *OnenoteRequestBuilder) Patch(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Onenoteable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property onenote in sites
func (m *OnenoteRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Onenoteable, requestConfiguration *OnenoteRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Resources the resources property
func (m *OnenoteRequestBuilder) Resources()(*i0decb69f831ce545caa107a838115b23576ee1efac361ea22cf80d81cb4ded00.ResourcesRequestBuilder) {
    return i0decb69f831ce545caa107a838115b23576ee1efac361ea22cf80d81cb4ded00.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.resources.item collection
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*ib6e5547c6ec7db85ec94ebc2072c06598dc08ff7d88cc9151447095852200162.OnenoteResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource%2Did"] = id
    }
    return ib6e5547c6ec7db85ec94ebc2072c06598dc08ff7d88cc9151447095852200162.NewOnenoteResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SectionGroups the sectionGroups property
func (m *OnenoteRequestBuilder) SectionGroups()(*ifc9a684d2dbdb1fe7d1eae49e0a8c30c188888963093335a2e0b25cb77cacb10.SectionGroupsRequestBuilder) {
    return ifc9a684d2dbdb1fe7d1eae49e0a8c30c188888963093335a2e0b25cb77cacb10.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.sectionGroups.item collection
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*i2e9176f71b9cfb53869eedb784edba98a6ec7f334a527247849d1da13f2a52a6.SectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup%2Did"] = id
    }
    return i2e9176f71b9cfb53869eedb784edba98a6ec7f334a527247849d1da13f2a52a6.NewSectionGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sections the sections property
func (m *OnenoteRequestBuilder) Sections()(*iff11e6eb619f9d0b688f685398702dc1c9c46658026db69f2e1441e49a1f8f3b.SectionsRequestBuilder) {
    return iff11e6eb619f9d0b688f685398702dc1c9c46658026db69f2e1441e49a1f8f3b.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.sections.item collection
func (m *OnenoteRequestBuilder) SectionsById(id string)(*ida022400bfcca3ce3366c6354f31a6136d19a7aa86866cfd1c28151b05cbee04.OnenoteSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection%2Did"] = id
    }
    return ida022400bfcca3ce3366c6354f31a6136d19a7aa86866cfd1c28151b05cbee04.NewOnenoteSectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
