package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder provides operations to manage the userStateSummary property of the microsoft.graph.managedEBook entity.
type ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderGetQueryParameters read properties and relationships of the userInstallStateSummary object.
type ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderGetQueryParameters
}
// ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderInternal instantiates a new ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder and sets the default values.
func NewManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder) {
    m := &ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/managedEBooks/{managedEBook%2Did}/userStateSummary/{userInstallStateSummary%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder instantiates a new ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder and sets the default values.
func NewManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a userInstallStateSummary.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-books-userinstallstatesummary-delete?view=graph-rest-1.0
func (m *ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceStates provides operations to manage the deviceStates property of the microsoft.graph.userInstallStateSummary entity.
// returns a *ManagedebooksItemUserstatesummaryItemDevicestatesDeviceStatesRequestBuilder when successful
func (m *ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder) DeviceStates()(*ManagedebooksItemUserstatesummaryItemDevicestatesDeviceStatesRequestBuilder) {
    return NewManagedebooksItemUserstatesummaryItemDevicestatesDeviceStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read properties and relationships of the userInstallStateSummary object.
// returns a UserInstallStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-books-userinstallstatesummary-get?view=graph-rest-1.0
func (m *ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserInstallStateSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserInstallStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserInstallStateSummaryable), nil
}
// Patch update the properties of a userInstallStateSummary object.
// returns a UserInstallStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-books-userinstallstatesummary-update?view=graph-rest-1.0
func (m *ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserInstallStateSummaryable, requestConfiguration *ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserInstallStateSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserInstallStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserInstallStateSummaryable), nil
}
// ToDeleteRequestInformation deletes a userInstallStateSummary.
// returns a *RequestInformation when successful
func (m *ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the userInstallStateSummary object.
// returns a *RequestInformation when successful
func (m *ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a userInstallStateSummary object.
// returns a *RequestInformation when successful
func (m *ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserInstallStateSummaryable, requestConfiguration *ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder when successful
func (m *ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder) WithUrl(rawUrl string)(*ManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder) {
    return NewManagedebooksItemUserstatesummaryUserInstallStateSummaryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
