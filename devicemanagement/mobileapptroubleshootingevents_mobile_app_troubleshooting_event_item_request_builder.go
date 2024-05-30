package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.deviceManagement entity.
type MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderGetQueryParameters read properties and relationships of the mobileAppTroubleshootingEvent object.
type MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderGetQueryParameters
}
// MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppLogCollectionRequests provides operations to manage the appLogCollectionRequests property of the microsoft.graph.mobileAppTroubleshootingEvent entity.
// returns a *MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestsRequestBuilder when successful
func (m *MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) AppLogCollectionRequests()(*MobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestsRequestBuilder) {
    return NewMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderInternal instantiates a new MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder and sets the default values.
func NewMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) {
    m := &MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/mobileAppTroubleshootingEvents/{mobileAppTroubleshootingEvent%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder instantiates a new MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder and sets the default values.
func NewMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a mobileAppTroubleshootingEvent.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-mobileapptroubleshootingevent-delete?view=graph-rest-1.0
func (m *MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read properties and relationships of the mobileAppTroubleshootingEvent object.
// returns a MobileAppTroubleshootingEventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-mobileapptroubleshootingevent-get?view=graph-rest-1.0
func (m *MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppTroubleshootingEventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileAppTroubleshootingEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppTroubleshootingEventable), nil
}
// Patch update the properties of a mobileAppTroubleshootingEvent object.
// returns a MobileAppTroubleshootingEventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-mobileapptroubleshootingevent-update?view=graph-rest-1.0
func (m *MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppTroubleshootingEventable, requestConfiguration *MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppTroubleshootingEventable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileAppTroubleshootingEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppTroubleshootingEventable), nil
}
// ToDeleteRequestInformation deletes a mobileAppTroubleshootingEvent.
// returns a *RequestInformation when successful
func (m *MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the mobileAppTroubleshootingEvent object.
// returns a *RequestInformation when successful
func (m *MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a mobileAppTroubleshootingEvent object.
// returns a *RequestInformation when successful
func (m *MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppTroubleshootingEventable, requestConfiguration *MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder when successful
func (m *MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) WithUrl(rawUrl string)(*MobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) {
    return NewMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
