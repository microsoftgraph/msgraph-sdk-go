package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i492f2c0deb9e84f7d9a7743fda4e1d1f3a1791c71cda353ebb7023d385689e39 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/children"
    i62a9e863e325b0aabc2b399fe176787f5e007e5d5ebde360ce5496bb73b60b1f "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/terms"
    i6bc6f7534ec2cd0ea31ebe97205d014f30992fabe1575742019d5cb73a56227b "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/parentgroup"
    ifb3c619491821827a3e3ee9e96a8b7fee32dcae5099fa00e741a8272d0353454 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/relations"
    i2b9fdf09ba8d0e3f8e4ec726b485908e535dd730cd31df3088c246390559cf0c "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/children/item"
    id29b486b9de74a5e228634cbc0b78a3e93bf36959bcca334777fa52e4dd5ba76 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/relations/item"
    id6ef0c2679ba27db5107619f8479aebb6116d87d37a91efddc176e4a725bef98 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/termstores/item/sets/item/terms/item"
)

// SetItemRequestBuilder provides operations to manage the sets property of the microsoft.graph.termStore.store entity.
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
// SetItemRequestBuilderGetQueryParameters collection of all sets available in the term store.
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
func (m *SetItemRequestBuilder) Children()(*i492f2c0deb9e84f7d9a7743fda4e1d1f3a1791c71cda353ebb7023d385689e39.ChildrenRequestBuilder) {
    return i492f2c0deb9e84f7d9a7743fda4e1d1f3a1791c71cda353ebb7023d385689e39.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStores.item.sets.item.children.item collection
func (m *SetItemRequestBuilder) ChildrenById(id string)(*i2b9fdf09ba8d0e3f8e4ec726b485908e535dd730cd31df3088c246390559cf0c.TermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did"] = id
    }
    return i2b9fdf09ba8d0e3f8e4ec726b485908e535dd730cd31df3088c246390559cf0c.NewTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSetItemRequestBuilderInternal instantiates a new SetItemRequestBuilder and sets the default values.
func NewSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetItemRequestBuilder) {
    m := &SetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/termStores/{store%2Did}/sets/{set%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property sets for sites
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
// CreateGetRequestInformation collection of all sets available in the term store.
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
// CreatePatchRequestInformation update the navigation property sets in sites
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
// Delete delete navigation property sets for sites
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
// Get collection of all sets available in the term store.
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
func (m *SetItemRequestBuilder) ParentGroup()(*i6bc6f7534ec2cd0ea31ebe97205d014f30992fabe1575742019d5cb73a56227b.ParentGroupRequestBuilder) {
    return i6bc6f7534ec2cd0ea31ebe97205d014f30992fabe1575742019d5cb73a56227b.NewParentGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property sets in sites
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
func (m *SetItemRequestBuilder) Relations()(*ifb3c619491821827a3e3ee9e96a8b7fee32dcae5099fa00e741a8272d0353454.RelationsRequestBuilder) {
    return ifb3c619491821827a3e3ee9e96a8b7fee32dcae5099fa00e741a8272d0353454.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStores.item.sets.item.relations.item collection
func (m *SetItemRequestBuilder) RelationsById(id string)(*id29b486b9de74a5e228634cbc0b78a3e93bf36959bcca334777fa52e4dd5ba76.RelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation%2Did"] = id
    }
    return id29b486b9de74a5e228634cbc0b78a3e93bf36959bcca334777fa52e4dd5ba76.NewRelationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Terms the terms property
func (m *SetItemRequestBuilder) Terms()(*i62a9e863e325b0aabc2b399fe176787f5e007e5d5ebde360ce5496bb73b60b1f.TermsRequestBuilder) {
    return i62a9e863e325b0aabc2b399fe176787f5e007e5d5ebde360ce5496bb73b60b1f.NewTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.termStores.item.sets.item.terms.item collection
func (m *SetItemRequestBuilder) TermsById(id string)(*id6ef0c2679ba27db5107619f8479aebb6116d87d37a91efddc176e4a725bef98.TermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did"] = id
    }
    return id6ef0c2679ba27db5107619f8479aebb6116d87d37a91efddc176e4a725bef98.NewTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
