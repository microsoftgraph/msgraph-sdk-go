package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i00b23f7f86a982af8a15ee36120329e64b13d88be4f288d635f6ae65ae966124 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list"
    i05fb43f9021a0913de359c2cfdf1fcad6b4de0310980145d3c6276c460270205 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/items"
    i2cbdf6d3a4841b26281917b6289b085e9cd675537f09928d44582bcd4533f5d2 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/root"
    i4c63c9b02f3ed101ac621ec88441d97587c19d3efeb627897c02cbe1ca7c6af4 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/listitem"
    i8c15d6975df59e7bff68ceaff8e2e7e5fb860064ce0c3fe87881dfa452f8d0f3 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/site"
    ia21810eb35790209720396d6baefa7bad436f4af7de39d51a975ac8ed17d9b90 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/permission"
    ic21e3d9f3eaf63bbdce62fd14ebed43e30e12830422fa068b6c26780d7bb64a6 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/driveitem"
    i645f1e1a516727ccf24f1ba68b8f4137ee52e397b6e58438bccce4f0cb3d74b5 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/items/item"
)

// SharedDriveItemItemRequestBuilder provides operations to manage the collection of sharedDriveItem entities.
type SharedDriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SharedDriveItemItemRequestBuilderDeleteOptions options for Delete
type SharedDriveItemItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// SharedDriveItemItemRequestBuilderGetOptions options for Get
type SharedDriveItemItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *SharedDriveItemItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// SharedDriveItemItemRequestBuilderGetQueryParameters get entity from shares by key
type SharedDriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SharedDriveItemItemRequestBuilderPatchOptions options for Patch
type SharedDriveItemItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SharedDriveItemable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewSharedDriveItemItemRequestBuilderInternal instantiates a new SharedDriveItemItemRequestBuilder and sets the default values.
func NewSharedDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharedDriveItemItemRequestBuilder) {
    m := &SharedDriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSharedDriveItemItemRequestBuilder instantiates a new SharedDriveItemItemRequestBuilder and sets the default values.
func NewSharedDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharedDriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSharedDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from shares
func (m *SharedDriveItemItemRequestBuilder) CreateDeleteRequestInformation(options *SharedDriveItemItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from shares by key
func (m *SharedDriveItemItemRequestBuilder) CreateGetRequestInformation(options *SharedDriveItemItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in shares
func (m *SharedDriveItemItemRequestBuilder) CreatePatchRequestInformation(options *SharedDriveItemItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete entity from shares
func (m *SharedDriveItemItemRequestBuilder) Delete(options *SharedDriveItemItemRequestBuilderDeleteOptions)(error) {
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
// DriveItem the driveItem property
func (m *SharedDriveItemItemRequestBuilder) DriveItem()(*ic21e3d9f3eaf63bbdce62fd14ebed43e30e12830422fa068b6c26780d7bb64a6.DriveItemRequestBuilder) {
    return ic21e3d9f3eaf63bbdce62fd14ebed43e30e12830422fa068b6c26780d7bb64a6.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entity from shares by key
func (m *SharedDriveItemItemRequestBuilder) Get(options *SharedDriveItemItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SharedDriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSharedDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SharedDriveItemable), nil
}
// Items the items property
func (m *SharedDriveItemItemRequestBuilder) Items()(*i05fb43f9021a0913de359c2cfdf1fcad6b4de0310980145d3c6276c460270205.ItemsRequestBuilder) {
    return i05fb43f9021a0913de359c2cfdf1fcad6b4de0310980145d3c6276c460270205.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.items.item collection
func (m *SharedDriveItemItemRequestBuilder) ItemsById(id string)(*i645f1e1a516727ccf24f1ba68b8f4137ee52e397b6e58438bccce4f0cb3d74b5.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i645f1e1a516727ccf24f1ba68b8f4137ee52e397b6e58438bccce4f0cb3d74b5.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// List the list property
func (m *SharedDriveItemItemRequestBuilder) List()(*i00b23f7f86a982af8a15ee36120329e64b13d88be4f288d635f6ae65ae966124.ListRequestBuilder) {
    return i00b23f7f86a982af8a15ee36120329e64b13d88be4f288d635f6ae65ae966124.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem the listItem property
func (m *SharedDriveItemItemRequestBuilder) ListItem()(*i4c63c9b02f3ed101ac621ec88441d97587c19d3efeb627897c02cbe1ca7c6af4.ListItemRequestBuilder) {
    return i4c63c9b02f3ed101ac621ec88441d97587c19d3efeb627897c02cbe1ca7c6af4.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in shares
func (m *SharedDriveItemItemRequestBuilder) Patch(options *SharedDriveItemItemRequestBuilderPatchOptions)(error) {
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
// Permission the permission property
func (m *SharedDriveItemItemRequestBuilder) Permission()(*ia21810eb35790209720396d6baefa7bad436f4af7de39d51a975ac8ed17d9b90.PermissionRequestBuilder) {
    return ia21810eb35790209720396d6baefa7bad436f4af7de39d51a975ac8ed17d9b90.NewPermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Root the root property
func (m *SharedDriveItemItemRequestBuilder) Root()(*i2cbdf6d3a4841b26281917b6289b085e9cd675537f09928d44582bcd4533f5d2.RootRequestBuilder) {
    return i2cbdf6d3a4841b26281917b6289b085e9cd675537f09928d44582bcd4533f5d2.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Site the site property
func (m *SharedDriveItemItemRequestBuilder) Site()(*i8c15d6975df59e7bff68ceaff8e2e7e5fb860064ce0c3fe87881dfa452f8d0f3.SiteRequestBuilder) {
    return i8c15d6975df59e7bff68ceaff8e2e7e5fb860064ce0c3fe87881dfa452f8d0f3.NewSiteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
