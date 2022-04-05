package list

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i4ceb3442437d25b4d65a5fb368c3ea2eebf89199b8598c6e3c12a21be842c870 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes"
    ibba07da9108e84b41a3d09d01bb4a6886a66e86df95f8f0730fa1c61201467ba "github.com/microsoftgraph/msgraph-sdk-go/drive/list/drive"
    ibf5d29d5462efb100f3ac502d1eba6060e878b0461baab8d107bba2657e634d5 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/subscriptions"
    ibf8d356e7a219821b2879666575a0d558fe0723087550da01c802cb67c2ea4a9 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/items"
    id137f289f6d68aa0ce87e57630353e341f93dbb3fb993d0ec54f3e9944c8865c "github.com/microsoftgraph/msgraph-sdk-go/drive/list/columns"
    i05799d6c1d82ecbf851acb48a1cb741eeecec1479585ec331eb151a1eb12fef7 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/columns/item"
    i0cc711a2d9e9934918c44b76b7f313fb885f4cc911346a17e775fac04e5518ac "github.com/microsoftgraph/msgraph-sdk-go/drive/list/subscriptions/item"
    i666016b9de0fea6c6ca42db4cdc518e2bbe5b22999b3927c4fc4035901b10667 "github.com/microsoftgraph/msgraph-sdk-go/drive/list/contenttypes/item"
    if1844020acb9c67e872cff770020bec5e9fbe0ec4bdb729561240d6a686db31c "github.com/microsoftgraph/msgraph-sdk-go/drive/list/items/item"
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
func (m *ListRequestBuilder) Columns()(*id137f289f6d68aa0ce87e57630353e341f93dbb3fb993d0ec54f3e9944c8865c.ColumnsRequestBuilder) {
    return id137f289f6d68aa0ce87e57630353e341f93dbb3fb993d0ec54f3e9944c8865c.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.list.columns.item collection
func (m *ListRequestBuilder) ColumnsById(id string)(*i05799d6c1d82ecbf851acb48a1cb741eeecec1479585ec331eb151a1eb12fef7.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i05799d6c1d82ecbf851acb48a1cb741eeecec1479585ec331eb151a1eb12fef7.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListRequestBuilderInternal instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/list{?select,expand}";
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
func (m *ListRequestBuilder) ContentTypes()(*i4ceb3442437d25b4d65a5fb368c3ea2eebf89199b8598c6e3c12a21be842c870.ContentTypesRequestBuilder) {
    return i4ceb3442437d25b4d65a5fb368c3ea2eebf89199b8598c6e3c12a21be842c870.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.list.contentTypes.item collection
func (m *ListRequestBuilder) ContentTypesById(id string)(*i666016b9de0fea6c6ca42db4cdc518e2bbe5b22999b3927c4fc4035901b10667.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return i666016b9de0fea6c6ca42db4cdc518e2bbe5b22999b3927c4fc4035901b10667.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property list for drive
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
// CreatePatchRequestInformation update the navigation property list in drive
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
// Delete delete navigation property list for drive
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
func (m *ListRequestBuilder) Drive()(*ibba07da9108e84b41a3d09d01bb4a6886a66e86df95f8f0730fa1c61201467ba.DriveRequestBuilder) {
    return ibba07da9108e84b41a3d09d01bb4a6886a66e86df95f8f0730fa1c61201467ba.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ListRequestBuilder) Items()(*ibf8d356e7a219821b2879666575a0d558fe0723087550da01c802cb67c2ea4a9.ItemsRequestBuilder) {
    return ibf8d356e7a219821b2879666575a0d558fe0723087550da01c802cb67c2ea4a9.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.list.items.item collection
func (m *ListRequestBuilder) ItemsById(id string)(*if1844020acb9c67e872cff770020bec5e9fbe0ec4bdb729561240d6a686db31c.ListItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem_id"] = id
    }
    return if1844020acb9c67e872cff770020bec5e9fbe0ec4bdb729561240d6a686db31c.NewListItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property list in drive
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
func (m *ListRequestBuilder) Subscriptions()(*ibf5d29d5462efb100f3ac502d1eba6060e878b0461baab8d107bba2657e634d5.SubscriptionsRequestBuilder) {
    return ibf5d29d5462efb100f3ac502d1eba6060e878b0461baab8d107bba2657e634d5.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.list.subscriptions.item collection
func (m *ListRequestBuilder) SubscriptionsById(id string)(*i0cc711a2d9e9934918c44b76b7f313fb885f4cc911346a17e775fac04e5518ac.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i0cc711a2d9e9934918c44b76b7f313fb885f4cc911346a17e775fac04e5518ac.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
