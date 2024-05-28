package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveRequestBuilder provides operations to manage the drive property of the microsoft.graph.fileStorageContainer entity.
type FilestorageContainersItemDriveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilestorageContainersItemDriveRequestBuilderGetQueryParameters the drive of the resource fileStorageContainer. Read-only.
type FilestorageContainersItemDriveRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageContainersItemDriveRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveRequestBuilderGetQueryParameters
}
// FilestorageContainersItemDriveRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Bundles provides operations to manage the bundles property of the microsoft.graph.drive entity.
// returns a *FilestorageContainersItemDriveBundlesRequestBuilder when successful
func (m *FilestorageContainersItemDriveRequestBuilder) Bundles()(*FilestorageContainersItemDriveBundlesRequestBuilder) {
    return NewFilestorageContainersItemDriveBundlesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilestorageContainersItemDriveRequestBuilderInternal instantiates a new FilestorageContainersItemDriveRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveRequestBuilder) {
    m := &FilestorageContainersItemDriveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveRequestBuilder instantiates a new FilestorageContainersItemDriveRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder when successful
func (m *FilestorageContainersItemDriveRequestBuilder) CreatedByUser()(*FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder) {
    return NewFilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property drive for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Following provides operations to manage the following property of the microsoft.graph.drive entity.
// returns a *FilestorageContainersItemDriveFollowingRequestBuilder when successful
func (m *FilestorageContainersItemDriveRequestBuilder) Following()(*FilestorageContainersItemDriveFollowingRequestBuilder) {
    return NewFilestorageContainersItemDriveFollowingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the drive of the resource fileStorageContainer. Read-only.
// returns a Driveable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable), nil
}
// Items provides operations to manage the items property of the microsoft.graph.drive entity.
// returns a *FilestorageContainersItemDriveItemsRequestBuilder when successful
func (m *FilestorageContainersItemDriveRequestBuilder) Items()(*FilestorageContainersItemDriveItemsRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *FilestorageContainersItemDriveLastmodifiedbyuserLastModifiedByUserRequestBuilder when successful
func (m *FilestorageContainersItemDriveRequestBuilder) LastModifiedByUser()(*FilestorageContainersItemDriveLastmodifiedbyuserLastModifiedByUserRequestBuilder) {
    return NewFilestorageContainersItemDriveLastmodifiedbyuserLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// List provides operations to manage the list property of the microsoft.graph.drive entity.
// returns a *FilestorageContainersItemDriveListRequestBuilder when successful
func (m *FilestorageContainersItemDriveRequestBuilder) List()(*FilestorageContainersItemDriveListRequestBuilder) {
    return NewFilestorageContainersItemDriveListRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property drive in storage
// returns a Driveable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, requestConfiguration *FilestorageContainersItemDriveRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable), nil
}
// Recent provides operations to call the recent method.
// returns a *FilestorageContainersItemDriveRecentRequestBuilder when successful
func (m *FilestorageContainersItemDriveRequestBuilder) Recent()(*FilestorageContainersItemDriveRecentRequestBuilder) {
    return NewFilestorageContainersItemDriveRecentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Root provides operations to manage the root property of the microsoft.graph.drive entity.
// returns a *FilestorageContainersItemDriveRootRequestBuilder when successful
func (m *FilestorageContainersItemDriveRequestBuilder) Root()(*FilestorageContainersItemDriveRootRequestBuilder) {
    return NewFilestorageContainersItemDriveRootRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SearchWithQ provides operations to call the search method.
// returns a *FilestorageContainersItemDriveSearchwithqSearchWithQRequestBuilder when successful
func (m *FilestorageContainersItemDriveRequestBuilder) SearchWithQ(q *string)(*FilestorageContainersItemDriveSearchwithqSearchWithQRequestBuilder) {
    return NewFilestorageContainersItemDriveSearchwithqSearchWithQRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, q)
}
// SharedWithMe provides operations to call the sharedWithMe method.
// returns a *FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder when successful
func (m *FilestorageContainersItemDriveRequestBuilder) SharedWithMe()(*FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder) {
    return NewFilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Special provides operations to manage the special property of the microsoft.graph.drive entity.
// returns a *FilestorageContainersItemDriveSpecialRequestBuilder when successful
func (m *FilestorageContainersItemDriveRequestBuilder) Special()(*FilestorageContainersItemDriveSpecialRequestBuilder) {
    return NewFilestorageContainersItemDriveSpecialRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property drive for storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the drive of the resource fileStorageContainer. Read-only.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property drive in storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, requestConfiguration *FilestorageContainersItemDriveRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveRequestBuilder when successful
func (m *FilestorageContainersItemDriveRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveRequestBuilder) {
    return NewFilestorageContainersItemDriveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
