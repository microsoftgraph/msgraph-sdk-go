package orgcontact

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i5307d39e321084657374420197d6bad76704cd8af975569f63dfd909cb2d4ee2 "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/directreports/item/orgcontact/getmemberobjects"
    i5a5287510ff323763268cf5d31e52be81f704086cc7f8bb24d9313bebcbeb5df "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/directreports/item/orgcontact/restore"
    i9808996e3f1d9376a42b86dd31652ad859e36735b7abb0bafcf166f5627cf391 "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/directreports/item/orgcontact/checkmemberobjects"
    ieadfc03fcf73c63ca038ff5ecd23e02d91716df5c7218c891fec5f01f7d177b4 "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/directreports/item/orgcontact/getmembergroups"
    ife1369acc17607af23215a7733c7fd659130e6c869f913924647b537f9f8d3f9 "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/directreports/item/orgcontact/checkmembergroups"
)

// OrgContactRequestBuilder casts the previous resource to orgContact.
type OrgContactRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OrgContactRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.orgContact
type OrgContactRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OrgContactRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrgContactRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OrgContactRequestBuilderGetQueryParameters
}
// CheckMemberGroups the checkMemberGroups property
func (m *OrgContactRequestBuilder) CheckMemberGroups()(*ife1369acc17607af23215a7733c7fd659130e6c869f913924647b537f9f8d3f9.CheckMemberGroupsRequestBuilder) {
    return ife1369acc17607af23215a7733c7fd659130e6c869f913924647b537f9f8d3f9.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *OrgContactRequestBuilder) CheckMemberObjects()(*i9808996e3f1d9376a42b86dd31652ad859e36735b7abb0bafcf166f5627cf391.CheckMemberObjectsRequestBuilder) {
    return i9808996e3f1d9376a42b86dd31652ad859e36735b7abb0bafcf166f5627cf391.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOrgContactRequestBuilderInternal instantiates a new OrgContactRequestBuilder and sets the default values.
func NewOrgContactRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrgContactRequestBuilder) {
    m := &OrgContactRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contacts/{orgContact%2Did}/directReports/{directoryObject%2Did}/microsoft.graph.orgContact{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOrgContactRequestBuilder instantiates a new OrgContactRequestBuilder and sets the default values.
func NewOrgContactRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrgContactRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrgContactRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.orgContact
func (m *OrgContactRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get the item of type microsoft.graph.directoryObject as microsoft.graph.orgContact
func (m *OrgContactRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OrgContactRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.orgContact
func (m *OrgContactRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrgContactable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *OrgContactRequestBuilder) GetMemberGroups()(*ieadfc03fcf73c63ca038ff5ecd23e02d91716df5c7218c891fec5f01f7d177b4.GetMemberGroupsRequestBuilder) {
    return ieadfc03fcf73c63ca038ff5ecd23e02d91716df5c7218c891fec5f01f7d177b4.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *OrgContactRequestBuilder) GetMemberObjects()(*i5307d39e321084657374420197d6bad76704cd8af975569f63dfd909cb2d4ee2.GetMemberObjectsRequestBuilder) {
    return i5307d39e321084657374420197d6bad76704cd8af975569f63dfd909cb2d4ee2.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler get the item of type microsoft.graph.directoryObject as microsoft.graph.orgContact
func (m *OrgContactRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *OrgContactRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrgContactable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOrgContactFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrgContactable), nil
}
// Restore the restore property
func (m *OrgContactRequestBuilder) Restore()(*i5a5287510ff323763268cf5d31e52be81f704086cc7f8bb24d9313bebcbeb5df.RestoreRequestBuilder) {
    return i5a5287510ff323763268cf5d31e52be81f704086cc7f8bb24d9313bebcbeb5df.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
