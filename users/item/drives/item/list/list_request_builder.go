package list

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0a2f2f73182529bc8b64f8a7561551f86018f289b0abc1f830d35cf66e13abd5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/columns"
    i0ca76db5f483d1db19e0884180e9786159b88d2ba5ba741d9de1f7305a359a50 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/items"
    i36db8193e86cf9461ce27c1050f2d6274041ce8ec754482b5ee7cbd6e0d19802 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/subscriptions"
    i9d2f3562a7d1dc275e15dc0355706aede3eab3e51167df6c8ffb776173606540 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/contenttypes"
    ie3f3789ac07f1388bb2b5674530aedbb4e0dce53003b432737b4f006585b762f "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/drive"
    i25df02055dbdbe96f9b3a308ffbcdda8b30068f7454bde6485a1e24f34d24190 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/items/item"
    i5a3f27133d05de6d006d747e000fdbacc132866443580fe76e6254bb8ca75367 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/columns/item"
    i979804fe5c4273a910e7ee1862177dd5d0b7de4ecab7b038180e38e90049961f "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/subscriptions/item"
    i98c11481076e86227e589090197425250d1f510801673c33ea2fd9321c2baa7f "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/contenttypes/item"
)

// ListRequestBuilder provides operations to manage the list property of the microsoft.graph.drive entity.
type ListRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ListRequestBuilderDeleteOptions options for Delete
type ListRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ListRequestBuilderGetOptions options for Get
type ListRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *ListRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ListRequestBuilderGetQueryParameters for drives in SharePoint, the underlying document library list. Read-only. Nullable.
type ListRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ListRequestBuilderPatchOptions options for Patch
type ListRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Listable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Columns the columns property
func (m *ListRequestBuilder) Columns()(*i0a2f2f73182529bc8b64f8a7561551f86018f289b0abc1f830d35cf66e13abd5.ColumnsRequestBuilder) {
    return i0a2f2f73182529bc8b64f8a7561551f86018f289b0abc1f830d35cf66e13abd5.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.list.columns.item collection
func (m *ListRequestBuilder) ColumnsById(id string)(*i5a3f27133d05de6d006d747e000fdbacc132866443580fe76e6254bb8ca75367.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i5a3f27133d05de6d006d747e000fdbacc132866443580fe76e6254bb8ca75367.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListRequestBuilderInternal instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/drives/{drive_id}/list{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListRequestBuilder instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentTypes the contentTypes property
func (m *ListRequestBuilder) ContentTypes()(*i9d2f3562a7d1dc275e15dc0355706aede3eab3e51167df6c8ffb776173606540.ContentTypesRequestBuilder) {
    return i9d2f3562a7d1dc275e15dc0355706aede3eab3e51167df6c8ffb776173606540.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.list.contentTypes.item collection
func (m *ListRequestBuilder) ContentTypesById(id string)(*i98c11481076e86227e589090197425250d1f510801673c33ea2fd9321c2baa7f.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return i98c11481076e86227e589090197425250d1f510801673c33ea2fd9321c2baa7f.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property list for users
func (m *ListRequestBuilder) CreateDeleteRequestInformation(options *ListRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation for drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *ListRequestBuilder) CreateGetRequestInformation(options *ListRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property list in users
func (m *ListRequestBuilder) CreatePatchRequestInformation(options *ListRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property list for users
func (m *ListRequestBuilder) Delete(options *ListRequestBuilderDeleteOptions)(error) {
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
func (m *ListRequestBuilder) Drive()(*ie3f3789ac07f1388bb2b5674530aedbb4e0dce53003b432737b4f006585b762f.DriveRequestBuilder) {
    return ie3f3789ac07f1388bb2b5674530aedbb4e0dce53003b432737b4f006585b762f.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get for drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *ListRequestBuilder) Get(options *ListRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Listable, error) {
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
func (m *ListRequestBuilder) Items()(*i0ca76db5f483d1db19e0884180e9786159b88d2ba5ba741d9de1f7305a359a50.ItemsRequestBuilder) {
    return i0ca76db5f483d1db19e0884180e9786159b88d2ba5ba741d9de1f7305a359a50.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.list.items.item collection
func (m *ListRequestBuilder) ItemsById(id string)(*i25df02055dbdbe96f9b3a308ffbcdda8b30068f7454bde6485a1e24f34d24190.ListItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem_id"] = id
    }
    return i25df02055dbdbe96f9b3a308ffbcdda8b30068f7454bde6485a1e24f34d24190.NewListItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property list in users
func (m *ListRequestBuilder) Patch(options *ListRequestBuilderPatchOptions)(error) {
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
func (m *ListRequestBuilder) Subscriptions()(*i36db8193e86cf9461ce27c1050f2d6274041ce8ec754482b5ee7cbd6e0d19802.SubscriptionsRequestBuilder) {
    return i36db8193e86cf9461ce27c1050f2d6274041ce8ec754482b5ee7cbd6e0d19802.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.list.subscriptions.item collection
func (m *ListRequestBuilder) SubscriptionsById(id string)(*i979804fe5c4273a910e7ee1862177dd5d0b7de4ecab7b038180e38e90049961f.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i979804fe5c4273a910e7ee1862177dd5d0b7de4ecab7b038180e38e90049961f.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
