package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i573e4def2e8111823894f1653301a31ec7ae07f9a4c0cd8ecab031680832c88b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/drive"
    i80ebad5d4a81635b20f1ed0f2527166bd8924c4c4c48998cae2cfd0032b93d48 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/contenttypes"
    i901dec75b2793d8f60a7087b67d833ed48b1254a9ab8c8e97cb4442e0148a8a2 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/items"
    id7d043ca8a65983c58c7e251b201e72119f874626c5be0adf785d232ddf163bf "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/subscriptions"
    if779415913e4212e9ca9af78533208bfbadb2715daa8cd231eb35ed7c011770e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/columns"
    i1f70debb81b03e21dc40d4b9610f0b304f5af25e2dadadb009789e8259efa413 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/columns/item"
    i26f59ebd4477aea7b288c6eb4e601e7de3217fb5feef3ede6d0dd025b2a1e31f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/items/item"
    i3e8954efb927f8b9d02478e7c153cf44922a2f2370bd61b14ee13c64bfff4192 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/contenttypes/item"
    i8ea87f604121fa8dc4bf1a77aec840b80d69f6919f7168703b00ef843d3d862c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item/subscriptions/item"
)

// ListItemRequestBuilder provides operations to manage the lists property of the microsoft.graph.site entity.
type ListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ListItemRequestBuilderDeleteOptions options for Delete
type ListItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ListItemRequestBuilderGetOptions options for Get
type ListItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ListItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ListItemRequestBuilderGetQueryParameters the collection of lists under this site.
type ListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ListItemRequestBuilderPatchOptions options for Patch
type ListItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Listable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Columns the columns property
func (m *ListItemRequestBuilder) Columns()(*if779415913e4212e9ca9af78533208bfbadb2715daa8cd231eb35ed7c011770e.ColumnsRequestBuilder) {
    return if779415913e4212e9ca9af78533208bfbadb2715daa8cd231eb35ed7c011770e.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.lists.item.columns.item collection
func (m *ListItemRequestBuilder) ColumnsById(id string)(*i1f70debb81b03e21dc40d4b9610f0b304f5af25e2dadadb009789e8259efa413.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i1f70debb81b03e21dc40d4b9610f0b304f5af25e2dadadb009789e8259efa413.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListItemRequestBuilder instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentTypes the contentTypes property
func (m *ListItemRequestBuilder) ContentTypes()(*i80ebad5d4a81635b20f1ed0f2527166bd8924c4c4c48998cae2cfd0032b93d48.ContentTypesRequestBuilder) {
    return i80ebad5d4a81635b20f1ed0f2527166bd8924c4c4c48998cae2cfd0032b93d48.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.lists.item.contentTypes.item collection
func (m *ListItemRequestBuilder) ContentTypesById(id string)(*i3e8954efb927f8b9d02478e7c153cf44922a2f2370bd61b14ee13c64bfff4192.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did"] = id
    }
    return i3e8954efb927f8b9d02478e7c153cf44922a2f2370bd61b14ee13c64bfff4192.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property lists for groups
func (m *ListItemRequestBuilder) CreateDeleteRequestInformation(options *ListItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of lists under this site.
func (m *ListItemRequestBuilder) CreateGetRequestInformation(options *ListItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property lists in groups
func (m *ListItemRequestBuilder) CreatePatchRequestInformation(options *ListItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property lists for groups
func (m *ListItemRequestBuilder) Delete(options *ListItemRequestBuilderDeleteOptions)(error) {
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
// Drive the drive property
func (m *ListItemRequestBuilder) Drive()(*i573e4def2e8111823894f1653301a31ec7ae07f9a4c0cd8ecab031680832c88b.DriveRequestBuilder) {
    return i573e4def2e8111823894f1653301a31ec7ae07f9a4c0cd8ecab031680832c88b.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of lists under this site.
func (m *ListItemRequestBuilder) Get(options *ListItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Listable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateListFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Listable), nil
}
// Items the items property
func (m *ListItemRequestBuilder) Items()(*i901dec75b2793d8f60a7087b67d833ed48b1254a9ab8c8e97cb4442e0148a8a2.ItemsRequestBuilder) {
    return i901dec75b2793d8f60a7087b67d833ed48b1254a9ab8c8e97cb4442e0148a8a2.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.lists.item.items.item collection
func (m *ListItemRequestBuilder) ItemsById(id string)(*i26f59ebd4477aea7b288c6eb4e601e7de3217fb5feef3ede6d0dd025b2a1e31f.ListItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem%2Did"] = id
    }
    return i26f59ebd4477aea7b288c6eb4e601e7de3217fb5feef3ede6d0dd025b2a1e31f.NewListItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property lists in groups
func (m *ListItemRequestBuilder) Patch(options *ListItemRequestBuilderPatchOptions)(error) {
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
// Subscriptions the subscriptions property
func (m *ListItemRequestBuilder) Subscriptions()(*id7d043ca8a65983c58c7e251b201e72119f874626c5be0adf785d232ddf163bf.SubscriptionsRequestBuilder) {
    return id7d043ca8a65983c58c7e251b201e72119f874626c5be0adf785d232ddf163bf.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.lists.item.subscriptions.item collection
func (m *ListItemRequestBuilder) SubscriptionsById(id string)(*i8ea87f604121fa8dc4bf1a77aec840b80d69f6919f7168703b00ef843d3d862c.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i8ea87f604121fa8dc4bf1a77aec840b80d69f6919f7168703b00ef843d3d862c.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
