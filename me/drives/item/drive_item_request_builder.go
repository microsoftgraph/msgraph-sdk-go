package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i091165ee47a53365087a9b57bf31e9ca719ff7771179173713ec1e124a4a753d "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/following"
    i0dae3c20cb0fc75243bf5f3fe7a800224db7512d907490fbb718cb2eecffc1f6 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/items"
    i14f7d808b16988af7c461c10ab368dc33cf12f6dd059a02532983cbe76ecb3cf "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list"
    i71a78fc710509695f355c29c03de4fc7f02c5982fcd79825f756f25d45ea9021 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/root"
    ib3a7e73bc5ddfc97d824cd78d5150cb9dfa843ede58050738ee6654be3931ce2 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/special"
    ib9ba40d643265a3cd641fc287bf1cdfba67d4ecd3b7e85f186bdc1e2719a8508 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/bundles"
    i02963047ed83bfaaa492b881f396af2b691df3601a94c1338b7f279004d8950c "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/special/item"
    i904578fd49676fdd39eec5e2a3b32acf6b122da24f860ba553c0f41b6d894de0 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/items/item"
    ie35b5069eeea0359b3dc0a8264d4aa578446aa78e86d3bf1873094d7bfbafc39 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/bundles/item"
    ie5906441abbc73187c4eeab57ee5b669993d8fe5f2b96ae57c4a1f3ea9cbbfd0 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/following/item"
)

// DriveItemRequestBuilder provides operations to manage the drives property of the microsoft.graph.user entity.
type DriveItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DriveItemRequestBuilderDeleteOptions options for Delete
type DriveItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// DriveItemRequestBuilderGetOptions options for Get
type DriveItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DriveItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// DriveItemRequestBuilderGetQueryParameters a collection of drives available for this user. Read-only.
type DriveItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DriveItemRequestBuilderPatchOptions options for Patch
type DriveItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Bundles the bundles property
func (m *DriveItemRequestBuilder) Bundles()(*ib9ba40d643265a3cd641fc287bf1cdfba67d4ecd3b7e85f186bdc1e2719a8508.BundlesRequestBuilder) {
    return ib9ba40d643265a3cd641fc287bf1cdfba67d4ecd3b7e85f186bdc1e2719a8508.NewBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BundlesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.bundles.item collection
func (m *DriveItemRequestBuilder) BundlesById(id string)(*ie35b5069eeea0359b3dc0a8264d4aa578446aa78e86d3bf1873094d7bfbafc39.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return ie35b5069eeea0359b3dc0a8264d4aa578446aa78e86d3bf1873094d7bfbafc39.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemRequestBuilderInternal instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemRequestBuilder) {
    m := &DriveItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveItemRequestBuilder instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property drives for me
func (m *DriveItemRequestBuilder) CreateDeleteRequestInformation(options *DriveItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of drives available for this user. Read-only.
func (m *DriveItemRequestBuilder) CreateGetRequestInformation(options *DriveItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property drives in me
func (m *DriveItemRequestBuilder) CreatePatchRequestInformation(options *DriveItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property drives for me
func (m *DriveItemRequestBuilder) Delete(options *DriveItemRequestBuilderDeleteOptions)(error) {
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
// Following the following property
func (m *DriveItemRequestBuilder) Following()(*i091165ee47a53365087a9b57bf31e9ca719ff7771179173713ec1e124a4a753d.FollowingRequestBuilder) {
    return i091165ee47a53365087a9b57bf31e9ca719ff7771179173713ec1e124a4a753d.NewFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowingById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.following.item collection
func (m *DriveItemRequestBuilder) FollowingById(id string)(*ie5906441abbc73187c4eeab57ee5b669993d8fe5f2b96ae57c4a1f3ea9cbbfd0.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return ie5906441abbc73187c4eeab57ee5b669993d8fe5f2b96ae57c4a1f3ea9cbbfd0.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get a collection of drives available for this user. Read-only.
func (m *DriveItemRequestBuilder) Get(options *DriveItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable), nil
}
// Items the items property
func (m *DriveItemRequestBuilder) Items()(*i0dae3c20cb0fc75243bf5f3fe7a800224db7512d907490fbb718cb2eecffc1f6.ItemsRequestBuilder) {
    return i0dae3c20cb0fc75243bf5f3fe7a800224db7512d907490fbb718cb2eecffc1f6.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.items.item collection
func (m *DriveItemRequestBuilder) ItemsById(id string)(*i904578fd49676fdd39eec5e2a3b32acf6b122da24f860ba553c0f41b6d894de0.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i904578fd49676fdd39eec5e2a3b32acf6b122da24f860ba553c0f41b6d894de0.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// List the list property
func (m *DriveItemRequestBuilder) List()(*i14f7d808b16988af7c461c10ab368dc33cf12f6dd059a02532983cbe76ecb3cf.ListRequestBuilder) {
    return i14f7d808b16988af7c461c10ab368dc33cf12f6dd059a02532983cbe76ecb3cf.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property drives in me
func (m *DriveItemRequestBuilder) Patch(options *DriveItemRequestBuilderPatchOptions)(error) {
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
// Root the root property
func (m *DriveItemRequestBuilder) Root()(*i71a78fc710509695f355c29c03de4fc7f02c5982fcd79825f756f25d45ea9021.RootRequestBuilder) {
    return i71a78fc710509695f355c29c03de4fc7f02c5982fcd79825f756f25d45ea9021.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Special the special property
func (m *DriveItemRequestBuilder) Special()(*ib3a7e73bc5ddfc97d824cd78d5150cb9dfa843ede58050738ee6654be3931ce2.SpecialRequestBuilder) {
    return ib3a7e73bc5ddfc97d824cd78d5150cb9dfa843ede58050738ee6654be3931ce2.NewSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SpecialById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.special.item collection
func (m *DriveItemRequestBuilder) SpecialById(id string)(*i02963047ed83bfaaa492b881f396af2b691df3601a94c1338b7f279004d8950c.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i02963047ed83bfaaa492b881f396af2b691df3601a94c1338b7f279004d8950c.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
