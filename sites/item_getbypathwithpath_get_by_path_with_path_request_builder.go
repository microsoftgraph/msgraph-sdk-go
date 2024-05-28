package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemGetbypathwithpathGetByPathWithPathRequestBuilder provides operations to call the getByPath method.
type ItemGetbypathwithpathGetByPathWithPathRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetbypathwithpathGetByPathWithPathRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetbypathwithpathGetByPathWithPathRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathAnalyticsRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) Analytics()(*ItemGetbypathwithpathAnalyticsRequestBuilder) {
    return NewItemGetbypathwithpathAnalyticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathColumnsRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) Columns()(*ItemGetbypathwithpathColumnsRequestBuilder) {
    return NewItemGetbypathwithpathColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemGetbypathwithpathGetByPathWithPathRequestBuilderInternal instantiates a new ItemGetbypathwithpathGetByPathWithPathRequestBuilder and sets the default values.
func NewItemGetbypathwithpathGetByPathWithPathRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, path *string)(*ItemGetbypathwithpathGetByPathWithPathRequestBuilder) {
    m := &ItemGetbypathwithpathGetByPathWithPathRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/getByPath(path='{path}')", pathParameters),
    }
    if path != nil {
        m.BaseRequestBuilder.PathParameters["path"] = *path
    }
    return m
}
// NewItemGetbypathwithpathGetByPathWithPathRequestBuilder instantiates a new ItemGetbypathwithpathGetByPathWithPathRequestBuilder and sets the default values.
func NewItemGetbypathwithpathGetByPathWithPathRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetbypathwithpathGetByPathWithPathRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetbypathwithpathGetByPathWithPathRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// ContentTypes provides operations to manage the contentTypes property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathContenttypesContentTypesRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) ContentTypes()(*ItemGetbypathwithpathContenttypesContentTypesRequestBuilder) {
    return NewItemGetbypathwithpathContenttypesContentTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *ItemGetbypathwithpathCreatedbyuserCreatedByUserRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) CreatedByUser()(*ItemGetbypathwithpathCreatedbyuserCreatedByUserRequestBuilder) {
    return NewItemGetbypathwithpathCreatedbyuserCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Drive provides operations to manage the drive property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathDriveRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) Drive()(*ItemGetbypathwithpathDriveRequestBuilder) {
    return NewItemGetbypathwithpathDriveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Drives provides operations to manage the drives property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathDrivesRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) Drives()(*ItemGetbypathwithpathDrivesRequestBuilder) {
    return NewItemGetbypathwithpathDrivesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExternalColumns provides operations to manage the externalColumns property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathExternalcolumnsExternalColumnsRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) ExternalColumns()(*ItemGetbypathwithpathExternalcolumnsExternalColumnsRequestBuilder) {
    return NewItemGetbypathwithpathExternalcolumnsExternalColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get invoke function getByPath
// returns a Siteable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGetbypathwithpathGetByPathWithPathRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, error) {
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
// returns a *ItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) GetActivitiesByInterval()(*ItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) {
    return NewItemGetbypathwithpathGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
// returns a *ItemGetbypathwithpathGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ItemGetbypathwithpathGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewItemGetbypathwithpathGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, endDateTime, interval, startDateTime)
}
// GetApplicableContentTypesForListWithListId provides operations to call the getApplicableContentTypesForList method.
// returns a *ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) GetApplicableContentTypesForListWithListId(listId *string)(*ItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilder) {
    return NewItemGetbypathwithpathGetapplicablecontenttypesforlistwithlistidGetApplicableContentTypesForListWithListIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, listId)
}
// GetByPathWithPath1 provides operations to call the getByPath method.
// returns a *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) GetByPathWithPath1(path1 *string)(*ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, path1)
}
// Items provides operations to manage the items property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathItemsRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) Items()(*ItemGetbypathwithpathItemsRequestBuilder) {
    return NewItemGetbypathwithpathItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *ItemGetbypathwithpathLastmodifiedbyuserLastModifiedByUserRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) LastModifiedByUser()(*ItemGetbypathwithpathLastmodifiedbyuserLastModifiedByUserRequestBuilder) {
    return NewItemGetbypathwithpathLastmodifiedbyuserLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Lists provides operations to manage the lists property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathListsRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) Lists()(*ItemGetbypathwithpathListsRequestBuilder) {
    return NewItemGetbypathwithpathListsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathOnenoteRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) Onenote()(*ItemGetbypathwithpathOnenoteRequestBuilder) {
    return NewItemGetbypathwithpathOnenoteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathOperationsRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) Operations()(*ItemGetbypathwithpathOperationsRequestBuilder) {
    return NewItemGetbypathwithpathOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pages provides operations to manage the pages property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathPagesRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) Pages()(*ItemGetbypathwithpathPagesRequestBuilder) {
    return NewItemGetbypathwithpathPagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathPermissionsRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) Permissions()(*ItemGetbypathwithpathPermissionsRequestBuilder) {
    return NewItemGetbypathwithpathPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sites provides operations to manage the sites property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathSitesRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) Sites()(*ItemGetbypathwithpathSitesRequestBuilder) {
    return NewItemGetbypathwithpathSitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TermStore provides operations to manage the termStore property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathTermstoreTermStoreRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) TermStore()(*ItemGetbypathwithpathTermstoreTermStoreRequestBuilder) {
    return NewItemGetbypathwithpathTermstoreTermStoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TermStores provides operations to manage the termStores property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathTermstoresTermStoresRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) TermStores()(*ItemGetbypathwithpathTermstoresTermStoresRequestBuilder) {
    return NewItemGetbypathwithpathTermstoresTermStoresRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation invoke function getByPath
// returns a *RequestInformation when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGetbypathwithpathGetByPathWithPathRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemGetbypathwithpathGetByPathWithPathRequestBuilder when successful
func (m *ItemGetbypathwithpathGetByPathWithPathRequestBuilder) WithUrl(rawUrl string)(*ItemGetbypathwithpathGetByPathWithPathRequestBuilder) {
    return NewItemGetbypathwithpathGetByPathWithPathRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
