package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveListItemsListItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.list entity.
type FilestorageContainersItemDriveListItemsListItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveListItemsListItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListItemsListItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilestorageContainersItemDriveListItemsListItemItemRequestBuilderGetQueryParameters all items contained in the list.
type FilestorageContainersItemDriveListItemsListItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageContainersItemDriveListItemsListItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListItemsListItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveListItemsListItemItemRequestBuilderGetQueryParameters
}
// FilestorageContainersItemDriveListItemsListItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListItemsListItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.listItem entity.
// returns a *FilestorageContainersItemDriveListItemsItemAnalyticsRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) Analytics()(*FilestorageContainersItemDriveListItemsItemAnalyticsRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsItemAnalyticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilestorageContainersItemDriveListItemsListItemItemRequestBuilderInternal instantiates a new FilestorageContainersItemDriveListItemsListItemItemRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListItemsListItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) {
    m := &FilestorageContainersItemDriveListItemsListItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/list/items/{listItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveListItemsListItemItemRequestBuilder instantiates a new FilestorageContainersItemDriveListItemsListItemItemRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListItemsListItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveListItemsListItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *FilestorageContainersItemDriveListItemsItemCreatedbyuserCreatedByUserRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) CreatedByUser()(*FilestorageContainersItemDriveListItemsItemCreatedbyuserCreatedByUserRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsItemCreatedbyuserCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateLink provides operations to call the createLink method.
// returns a *FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) CreateLink()(*FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property items for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListItemsListItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DocumentSetVersions provides operations to manage the documentSetVersions property of the microsoft.graph.listItem entity.
// returns a *FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) DocumentSetVersions()(*FilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DriveItem provides operations to manage the driveItem property of the microsoft.graph.listItem entity.
// returns a *FilestorageContainersItemDriveListItemsItemDriveitemDriveItemRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) DriveItem()(*FilestorageContainersItemDriveListItemsItemDriveitemDriveItemRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsItemDriveitemDriveItemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Fields provides operations to manage the fields property of the microsoft.graph.listItem entity.
// returns a *FilestorageContainersItemDriveListItemsItemFieldsRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) Fields()(*FilestorageContainersItemDriveListItemsItemFieldsRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsItemFieldsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get all items contained in the list.
// returns a ListItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListItemsListItemItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ListItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateListItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ListItemable), nil
}
// GetActivitiesByInterval provides operations to call the getActivitiesByInterval method.
// returns a *FilestorageContainersItemDriveListItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) GetActivitiesByInterval()(*FilestorageContainersItemDriveListItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
// returns a *FilestorageContainersItemDriveListItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*FilestorageContainersItemDriveListItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, endDateTime, interval, startDateTime)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *FilestorageContainersItemDriveListItemsItemLastmodifiedbyuserLastModifiedByUserRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) LastModifiedByUser()(*FilestorageContainersItemDriveListItemsItemLastmodifiedbyuserLastModifiedByUserRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsItemLastmodifiedbyuserLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property items in storage
// returns a ListItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ListItemable, requestConfiguration *FilestorageContainersItemDriveListItemsListItemItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ListItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateListItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ListItemable), nil
}
// ToDeleteRequestInformation delete navigation property items for storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListItemsListItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation all items contained in the list.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListItemsListItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property items in storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ListItemable, requestConfiguration *FilestorageContainersItemDriveListItemsListItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Versions provides operations to manage the versions property of the microsoft.graph.listItem entity.
// returns a *FilestorageContainersItemDriveListItemsItemVersionsRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) Versions()(*FilestorageContainersItemDriveListItemsItemVersionsRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsItemVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveListItemsListItemItemRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsListItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
