package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i4c90a5b5af67ed60ab6adf5613edbcb1da1c9dcd2ea15d3712d431fabae295d2 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/permissions"
    i6d048e988ecd7a469df6e77aceb776b58b4b96a45ec000a5ade64874f4ef46af "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/subscriptions"
    i8de7c056a112237127b0dcc4ea8d6badfd1aba50e70ba16d561038ac278428e1 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/thumbnails"
    i8ed13773ce508ebebaa60ab9a3303a3ca02078bf14b6afa3519a1293144b0a57 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/versions"
    i977cc3182db94374759f872d55be1bd164691f41d3ef9934b2a1b2ee3a5022b9 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/listitem"
    ibad2dbd640f8438d505ec69c54f24d1b51b0428157e3df1a56b77ca2b617c307 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/content"
    ic187abdb8f33a4cb81256b996a715ff597202bed85f8f4d64b6e8319f5022059 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/children"
    if6e4b55b696c1aaead2907344730bf6b12310b52afb14a9ffd1ce42161325da9 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/analytics"
    i01b95b3d57747e53543e8dadcce70427ca7796d7df1fa1a45329a65eae314015 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/versions/item"
    i0a1467d2d422249bfaa9b52d3bde594deab07162016808ecd41602cbd68904ce "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/permissions/item"
    i170fcbbf415e96234b1aa6424878a65c5d7b82ec53c144e5203f73656e625891 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/subscriptions/item"
    i33b94f1f1fae365d235ffc10e8627c15caaca7be62b15dfd7f8cb6a23342e856 "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/children/item"
    icf7102c66807ffb1835310fb1cb3f4be4c0c9d50a0fe651aac0d3390144ccfad "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item/thumbnails/item"
)

// DriveItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.drive entity.
type DriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DriveItemItemRequestBuilderDeleteOptions options for Delete
type DriveItemItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DriveItemItemRequestBuilderGetOptions options for Get
type DriveItemItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DriveItemItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DriveItemItemRequestBuilderGetQueryParameters all items contained in the drive. Read-only. Nullable.
type DriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DriveItemItemRequestBuilderPatchOptions options for Patch
type DriveItemItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Analytics the analytics property
func (m *DriveItemItemRequestBuilder) Analytics()(*if6e4b55b696c1aaead2907344730bf6b12310b52afb14a9ffd1ce42161325da9.AnalyticsRequestBuilder) {
    return if6e4b55b696c1aaead2907344730bf6b12310b52afb14a9ffd1ce42161325da9.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *DriveItemItemRequestBuilder) Children()(*ic187abdb8f33a4cb81256b996a715ff597202bed85f8f4d64b6e8319f5022059.ChildrenRequestBuilder) {
    return ic187abdb8f33a4cb81256b996a715ff597202bed85f8f4d64b6e8319f5022059.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.items.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i33b94f1f1fae365d235ffc10e8627c15caaca7be62b15dfd7f8cb6a23342e856.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i33b94f1f1fae365d235ffc10e8627c15caaca7be62b15dfd7f8cb6a23342e856.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/items/{driveItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveItemItemRequestBuilder instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the content property
func (m *DriveItemItemRequestBuilder) Content()(*ibad2dbd640f8438d505ec69c54f24d1b51b0428157e3df1a56b77ca2b617c307.ContentRequestBuilder) {
    return ibad2dbd640f8438d505ec69c54f24d1b51b0428157e3df1a56b77ca2b617c307.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property items for drive
func (m *DriveItemItemRequestBuilder) CreateDeleteRequestInformation(options *DriveItemItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) CreateGetRequestInformation(options *DriveItemItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property items in drive
func (m *DriveItemItemRequestBuilder) CreatePatchRequestInformation(options *DriveItemItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property items for drive
func (m *DriveItemItemRequestBuilder) Delete(options *DriveItemItemRequestBuilderDeleteOptions)(error) {
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
// Get all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) Get(options *DriveItemItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable), nil
}
// ListItem the listItem property
func (m *DriveItemItemRequestBuilder) ListItem()(*i977cc3182db94374759f872d55be1bd164691f41d3ef9934b2a1b2ee3a5022b9.ListItemRequestBuilder) {
    return i977cc3182db94374759f872d55be1bd164691f41d3ef9934b2a1b2ee3a5022b9.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property items in drive
func (m *DriveItemItemRequestBuilder) Patch(options *DriveItemItemRequestBuilderPatchOptions)(error) {
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
// Permissions the permissions property
func (m *DriveItemItemRequestBuilder) Permissions()(*i4c90a5b5af67ed60ab6adf5613edbcb1da1c9dcd2ea15d3712d431fabae295d2.PermissionsRequestBuilder) {
    return i4c90a5b5af67ed60ab6adf5613edbcb1da1c9dcd2ea15d3712d431fabae295d2.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.items.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i0a1467d2d422249bfaa9b52d3bde594deab07162016808ecd41602cbd68904ce.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i0a1467d2d422249bfaa9b52d3bde594deab07162016808ecd41602cbd68904ce.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Subscriptions the subscriptions property
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i6d048e988ecd7a469df6e77aceb776b58b4b96a45ec000a5ade64874f4ef46af.SubscriptionsRequestBuilder) {
    return i6d048e988ecd7a469df6e77aceb776b58b4b96a45ec000a5ade64874f4ef46af.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.items.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i170fcbbf415e96234b1aa6424878a65c5d7b82ec53c144e5203f73656e625891.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i170fcbbf415e96234b1aa6424878a65c5d7b82ec53c144e5203f73656e625891.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i8de7c056a112237127b0dcc4ea8d6badfd1aba50e70ba16d561038ac278428e1.ThumbnailsRequestBuilder) {
    return i8de7c056a112237127b0dcc4ea8d6badfd1aba50e70ba16d561038ac278428e1.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.items.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*icf7102c66807ffb1835310fb1cb3f4be4c0c9d50a0fe651aac0d3390144ccfad.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return icf7102c66807ffb1835310fb1cb3f4be4c0c9d50a0fe651aac0d3390144ccfad.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions the versions property
func (m *DriveItemItemRequestBuilder) Versions()(*i8ed13773ce508ebebaa60ab9a3303a3ca02078bf14b6afa3519a1293144b0a57.VersionsRequestBuilder) {
    return i8ed13773ce508ebebaa60ab9a3303a3ca02078bf14b6afa3519a1293144b0a57.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.items.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i01b95b3d57747e53543e8dadcce70427ca7796d7df1fa1a45329a65eae314015.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i01b95b3d57747e53543e8dadcce70427ca7796d7df1fa1a45329a65eae314015.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
