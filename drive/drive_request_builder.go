package drive

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DriveRequestBuilder provides operations to manage the drive singleton.
type DriveRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DriveRequestBuilderGetQueryParameters get drive
type DriveRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DriveRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DriveRequestBuilderGetQueryParameters
}
// DriveRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Bundles provides operations to manage the bundles property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) Bundles()(*BundlesRequestBuilder) {
    return NewBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BundlesById provides operations to manage the bundles property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) BundlesById(id string)(*DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveRequestBuilderInternal instantiates a new DriveRequestBuilder and sets the default values.
func NewDriveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveRequestBuilder) {
    m := &DriveRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveRequestBuilder instantiates a new DriveRequestBuilder and sets the default values.
func NewDriveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get drive
func (m *DriveRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DriveRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update drive
func (m *DriveRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, requestConfiguration *DriveRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Following provides operations to manage the following property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) Following()(*FollowingRequestBuilder) {
    return NewFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowingById provides operations to manage the following property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) FollowingById(id string)(*DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get drive
func (m *DriveRequestBuilder) Get(ctx context.Context, requestConfiguration *DriveRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, error) {
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
func (m *DriveRequestBuilder) Items()(*ItemsRequestBuilder) {
    return NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById provides operations to manage the items property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) ItemsById(id string)(*DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// List provides operations to manage the list property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) List()(*ListRequestBuilder) {
    return NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update drive
func (m *DriveRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, requestConfiguration *DriveRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, error) {
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
func (m *DriveRequestBuilder) Recent()(*RecentRequestBuilder) {
    return NewRecentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Root provides operations to manage the root property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) Root()(*RootRequestBuilder) {
    return NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DriveRequestBuilder) SearchWithQ(q *string)(*SearchWithQRequestBuilder) {
    return NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// SharedWithMe provides operations to call the sharedWithMe method.
func (m *DriveRequestBuilder) SharedWithMe()(*SharedWithMeRequestBuilder) {
    return NewSharedWithMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Special provides operations to manage the special property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) Special()(*SpecialRequestBuilder) {
    return NewSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SpecialById provides operations to manage the special property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) SpecialById(id string)(*DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
