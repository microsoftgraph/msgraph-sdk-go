package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder provides operations to call the getByPath method.
type ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathAnalyticsRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) Analytics()(*ItemSitesItemGetbypathwithpathAnalyticsRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathAnalyticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathColumnsRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) Columns()(*ItemSitesItemGetbypathwithpathColumnsRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilderInternal instantiates a new ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder and sets the default values.
func NewItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, path *string)(*ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) {
    m := &ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/getByPath(path='{path}')", pathParameters),
    }
    if path != nil {
        m.BaseRequestBuilder.PathParameters["path"] = *path
    }
    return m
}
// NewItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder instantiates a new ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder and sets the default values.
func NewItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// ContentTypes provides operations to manage the contentTypes property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathContenttypesContentTypesRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) ContentTypes()(*ItemSitesItemGetbypathwithpathContenttypesContentTypesRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathContenttypesContentTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *ItemSitesItemGetbypathwithpathCreatedbyuserCreatedByUserRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) CreatedByUser()(*ItemSitesItemGetbypathwithpathCreatedbyuserCreatedByUserRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathCreatedbyuserCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Drive provides operations to manage the drive property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathDriveRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) Drive()(*ItemSitesItemGetbypathwithpathDriveRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathDriveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Drives provides operations to manage the drives property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathDrivesRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) Drives()(*ItemSitesItemGetbypathwithpathDrivesRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathDrivesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExternalColumns provides operations to manage the externalColumns property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathExternalcolumnsExternalColumnsRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) ExternalColumns()(*ItemSitesItemGetbypathwithpathExternalcolumnsExternalColumnsRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathExternalcolumnsExternalColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get invoke function getByPath
// returns a Siteable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable), nil
}
// GetActivitiesByInterval provides operations to call the getActivitiesByInterval method.
// returns a *ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) GetActivitiesByInterval()(*ItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
// returns a *ItemSitesItemGetbypathwithpathGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ItemSitesItemGetbypathwithpathGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, endDateTime, interval, startDateTime)
}
// GetApplicableContentTypesForListWithListId provides operations to call the getApplicableContentTypesForList method.
// returns a *ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) GetApplicableContentTypesForListWithListId(listId *string)(*ItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, listId)
}
// GetByPathWithPath1 provides operations to call the getByPath method.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) GetByPathWithPath1(path1 *string)(*ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, path1)
}
// Items provides operations to manage the items property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathItemsRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) Items()(*ItemSitesItemGetbypathwithpathItemsRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *ItemSitesItemGetbypathwithpathLastmodifiedbyuserLastModifiedByUserRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) LastModifiedByUser()(*ItemSitesItemGetbypathwithpathLastmodifiedbyuserLastModifiedByUserRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathLastmodifiedbyuserLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Lists provides operations to manage the lists property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathListsRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) Lists()(*ItemSitesItemGetbypathwithpathListsRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathListsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathOnenoteRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) Onenote()(*ItemSitesItemGetbypathwithpathOnenoteRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathOnenoteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathOperationsRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) Operations()(*ItemSitesItemGetbypathwithpathOperationsRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pages provides operations to manage the pages property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathPagesRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) Pages()(*ItemSitesItemGetbypathwithpathPagesRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathPagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathPermissionsRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) Permissions()(*ItemSitesItemGetbypathwithpathPermissionsRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sites provides operations to manage the sites property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathSitesRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) Sites()(*ItemSitesItemGetbypathwithpathSitesRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathSitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TermStore provides operations to manage the termStore property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathTermstoreTermStoreRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) TermStore()(*ItemSitesItemGetbypathwithpathTermstoreTermStoreRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathTermstoreTermStoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TermStores provides operations to manage the termStores property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathTermstoresTermStoresRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) TermStores()(*ItemSitesItemGetbypathwithpathTermstoresTermStoresRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathTermstoresTermStoresRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation invoke function getByPath
// returns a *RequestInformation when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetByPathWithPathRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
