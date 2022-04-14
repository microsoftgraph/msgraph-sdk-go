package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i19fcb5942b41d05bc8e1c7d346744a769f0dcba6d3dfabd7abaecada60eb08e6 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/termstores/item/groups/item/sets/item/parentgroup"
    i2e200efe1145dadaf91fc2971972adf40635376e9d313f760aa4f671d4eda785 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/termstores/item/groups/item/sets/item/relations"
    ia0ad2eae4bf805bcb47202b33867afd6988b74b8a78a7ee875e146b601ab589f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/termstores/item/groups/item/sets/item/terms"
    ia80ca37edc2e8f343c9a828a29244438eb659c3d663b4309e5c9e4c3d1276748 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/termstores/item/groups/item/sets/item/children"
    i6c5105d33e0ae38adcea2e9e98a433443ab0d547d5cb40c689cbe279de6634e9 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/termstores/item/groups/item/sets/item/relations/item"
    ia2baa1b4705026e634ef4a1f93753b00111c9846fbf1fb7129a6a574ea125ba1 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/termstores/item/groups/item/sets/item/children/item"
    ifdc840e8cc75ed4281fa25878e2aeb0c352b6b48e54b9e4b22a4ab0de02edbc0 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/termstores/item/groups/item/sets/item/terms/item"
)

// SetItemRequestBuilder provides operations to manage the sets property of the microsoft.graph.termStore.group entity.
type SetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SetItemRequestBuilderDeleteOptions options for Delete
type SetItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SetItemRequestBuilderGetOptions options for Get
type SetItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SetItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SetItemRequestBuilderGetQueryParameters all sets under the group in a term [store].
type SetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SetItemRequestBuilderPatchOptions options for Patch
type SetItemRequestBuilderPatchOptions struct {
    // 
    Body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Children the children property
func (m *SetItemRequestBuilder) Children()(*ia80ca37edc2e8f343c9a828a29244438eb659c3d663b4309e5c9e4c3d1276748.ChildrenRequestBuilder) {
    return ia80ca37edc2e8f343c9a828a29244438eb659c3d663b4309e5c9e4c3d1276748.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.termStores.item.groups.item.sets.item.children.item collection
func (m *SetItemRequestBuilder) ChildrenById(id string)(*ia2baa1b4705026e634ef4a1f93753b00111c9846fbf1fb7129a6a574ea125ba1.TermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did"] = id
    }
    return ia2baa1b4705026e634ef4a1f93753b00111c9846fbf1fb7129a6a574ea125ba1.NewTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSetItemRequestBuilderInternal instantiates a new SetItemRequestBuilder and sets the default values.
func NewSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetItemRequestBuilder) {
    m := &SetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStores/{store%2Did}/groups/{group%2Did1}/sets/{set%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetItemRequestBuilder instantiates a new SetItemRequestBuilder and sets the default values.
func NewSetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sets for groups
func (m *SetItemRequestBuilder) CreateDeleteRequestInformation(options *SetItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation all sets under the group in a term [store].
func (m *SetItemRequestBuilder) CreateGetRequestInformation(options *SetItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property sets in groups
func (m *SetItemRequestBuilder) CreatePatchRequestInformation(options *SetItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property sets for groups
func (m *SetItemRequestBuilder) Delete(options *SetItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get all sets under the group in a term [store].
func (m *SetItemRequestBuilder) Get(options *SetItemRequestBuilderGetOptions)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateSetFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable), nil
}
// ParentGroup the parentGroup property
func (m *SetItemRequestBuilder) ParentGroup()(*i19fcb5942b41d05bc8e1c7d346744a769f0dcba6d3dfabd7abaecada60eb08e6.ParentGroupRequestBuilder) {
    return i19fcb5942b41d05bc8e1c7d346744a769f0dcba6d3dfabd7abaecada60eb08e6.NewParentGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property sets in groups
func (m *SetItemRequestBuilder) Patch(options *SetItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Relations the relations property
func (m *SetItemRequestBuilder) Relations()(*i2e200efe1145dadaf91fc2971972adf40635376e9d313f760aa4f671d4eda785.RelationsRequestBuilder) {
    return i2e200efe1145dadaf91fc2971972adf40635376e9d313f760aa4f671d4eda785.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.termStores.item.groups.item.sets.item.relations.item collection
func (m *SetItemRequestBuilder) RelationsById(id string)(*i6c5105d33e0ae38adcea2e9e98a433443ab0d547d5cb40c689cbe279de6634e9.RelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation%2Did"] = id
    }
    return i6c5105d33e0ae38adcea2e9e98a433443ab0d547d5cb40c689cbe279de6634e9.NewRelationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Terms the terms property
func (m *SetItemRequestBuilder) Terms()(*ia0ad2eae4bf805bcb47202b33867afd6988b74b8a78a7ee875e146b601ab589f.TermsRequestBuilder) {
    return ia0ad2eae4bf805bcb47202b33867afd6988b74b8a78a7ee875e146b601ab589f.NewTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.termStores.item.groups.item.sets.item.terms.item collection
func (m *SetItemRequestBuilder) TermsById(id string)(*ifdc840e8cc75ed4281fa25878e2aeb0c352b6b48e54b9e4b22a4ab0de02edbc0.TermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did"] = id
    }
    return ifdc840e8cc75ed4281fa25878e2aeb0c352b6b48e54b9e4b22a4ab0de02edbc0.NewTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
