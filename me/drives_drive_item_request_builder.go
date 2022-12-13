package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DrivesDriveItemRequestBuilder provides operations to manage the drives property of the microsoft.graph.user entity.
type DrivesDriveItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DrivesDriveItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DrivesDriveItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DrivesDriveItemRequestBuilderGetQueryParameters a collection of drives available for this user. Read-only.
type DrivesDriveItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DrivesDriveItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DrivesDriveItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DrivesDriveItemRequestBuilderGetQueryParameters
}
// DrivesDriveItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DrivesDriveItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Bundles provides operations to manage the bundles property of the microsoft.graph.drive entity.
func (m *DrivesDriveItemRequestBuilder) Bundles()(*DrivesItemBundlesRequestBuilder) {
    return NewDrivesItemBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BundlesById provides operations to manage the bundles property of the microsoft.graph.drive entity.
func (m *DrivesDriveItemRequestBuilder) BundlesById(id string)(*DrivesItemBundlesDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewDrivesItemBundlesDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDrivesDriveItemRequestBuilderInternal instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDrivesDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DrivesDriveItemRequestBuilder) {
    m := &DrivesDriveItemRequestBuilder{
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
// NewDrivesDriveItemRequestBuilder instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDrivesDriveItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DrivesDriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDrivesDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property drives for me
func (m *DrivesDriveItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DrivesDriveItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation a collection of drives available for this user. Read-only.
func (m *DrivesDriveItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DrivesDriveItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property drives in me
func (m *DrivesDriveItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, requestConfiguration *DrivesDriveItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property drives for me
func (m *DrivesDriveItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DrivesDriveItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Following provides operations to manage the following property of the microsoft.graph.drive entity.
func (m *DrivesDriveItemRequestBuilder) Following()(*DrivesItemFollowingRequestBuilder) {
    return NewDrivesItemFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowingById provides operations to manage the following property of the microsoft.graph.drive entity.
func (m *DrivesDriveItemRequestBuilder) FollowingById(id string)(*DrivesItemFollowingDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewDrivesItemFollowingDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get a collection of drives available for this user. Read-only.
func (m *DrivesDriveItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DrivesDriveItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable), nil
}
// Items provides operations to manage the items property of the microsoft.graph.drive entity.
func (m *DrivesDriveItemRequestBuilder) Items()(*DrivesItemItemsRequestBuilder) {
    return NewDrivesItemItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById provides operations to manage the items property of the microsoft.graph.drive entity.
func (m *DrivesDriveItemRequestBuilder) ItemsById(id string)(*DrivesItemItemsDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewDrivesItemItemsDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// List provides operations to manage the list property of the microsoft.graph.drive entity.
func (m *DrivesDriveItemRequestBuilder) List()(*DrivesItemListRequestBuilder) {
    return NewDrivesItemListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property drives in me
func (m *DrivesDriveItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, requestConfiguration *DrivesDriveItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable), nil
}
// Recent provides operations to call the recent method.
func (m *DrivesDriveItemRequestBuilder) Recent()(*DrivesItemRecentRequestBuilder) {
    return NewDrivesItemRecentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Root provides operations to manage the root property of the microsoft.graph.drive entity.
func (m *DrivesDriveItemRequestBuilder) Root()(*DrivesItemRootRequestBuilder) {
    return NewDrivesItemRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DrivesDriveItemRequestBuilder) SearchWithQ(q *string)(*DrivesItemSearchWithQRequestBuilder) {
    return NewDrivesItemSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// SharedWithMe provides operations to call the sharedWithMe method.
func (m *DrivesDriveItemRequestBuilder) SharedWithMe()(*DrivesItemSharedWithMeRequestBuilder) {
    return NewDrivesItemSharedWithMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Special provides operations to manage the special property of the microsoft.graph.drive entity.
func (m *DrivesDriveItemRequestBuilder) Special()(*DrivesItemSpecialRequestBuilder) {
    return NewDrivesItemSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SpecialById provides operations to manage the special property of the microsoft.graph.drive entity.
func (m *DrivesDriveItemRequestBuilder) SpecialById(id string)(*DrivesItemSpecialDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewDrivesItemSpecialDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
