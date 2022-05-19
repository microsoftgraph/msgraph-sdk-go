package orgcontact

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0df8517747613c63d0d2637a830081981a5e8217a56df7bba05da406a1838189 "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/memberof/item/orgcontact/getmembergroups"
    i351e7d23bf2cd1ce0319bea6531909167917735cf052450e60898fd212aae4a8 "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/memberof/item/orgcontact/checkmembergroups"
    i50c7500d19cd648b663331aed6a98e9c6c7e8c8d46c5b0e850e7dcaca85e3905 "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/memberof/item/orgcontact/checkmemberobjects"
    i5e0b205e76f9ebcf139676d6b209993bde9a95dadeb9d9d814723c2b3d40b938 "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/memberof/item/orgcontact/restore"
    i8499db48fda1a3977204844ab4b80cd074f7d3b3e4d0b05af7af3dec95079eaf "github.com/microsoftgraph/msgraph-sdk-go/contacts/item/memberof/item/orgcontact/getmemberobjects"
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
func (m *OrgContactRequestBuilder) CheckMemberGroups()(*i351e7d23bf2cd1ce0319bea6531909167917735cf052450e60898fd212aae4a8.CheckMemberGroupsRequestBuilder) {
    return i351e7d23bf2cd1ce0319bea6531909167917735cf052450e60898fd212aae4a8.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *OrgContactRequestBuilder) CheckMemberObjects()(*i50c7500d19cd648b663331aed6a98e9c6c7e8c8d46c5b0e850e7dcaca85e3905.CheckMemberObjectsRequestBuilder) {
    return i50c7500d19cd648b663331aed6a98e9c6c7e8c8d46c5b0e850e7dcaca85e3905.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOrgContactRequestBuilderInternal instantiates a new OrgContactRequestBuilder and sets the default values.
func NewOrgContactRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrgContactRequestBuilder) {
    m := &OrgContactRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contacts/{orgContact%2Did}/memberOf/{directoryObject%2Did}/microsoft.graph.orgContact{?%24select,%24expand}";
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
func (m *OrgContactRequestBuilder) GetMemberGroups()(*i0df8517747613c63d0d2637a830081981a5e8217a56df7bba05da406a1838189.GetMemberGroupsRequestBuilder) {
    return i0df8517747613c63d0d2637a830081981a5e8217a56df7bba05da406a1838189.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *OrgContactRequestBuilder) GetMemberObjects()(*i8499db48fda1a3977204844ab4b80cd074f7d3b3e4d0b05af7af3dec95079eaf.GetMemberObjectsRequestBuilder) {
    return i8499db48fda1a3977204844ab4b80cd074f7d3b3e4d0b05af7af3dec95079eaf.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *OrgContactRequestBuilder) Restore()(*i5e0b205e76f9ebcf139676d6b209993bde9a95dadeb9d9d814723c2b3d40b938.RestoreRequestBuilder) {
    return i5e0b205e76f9ebcf139676d6b209993bde9a95dadeb9d9d814723c2b3d40b938.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
