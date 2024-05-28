package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// Getoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder provides operations to call the getOffice365GroupsActivityStorage method.
type Getoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal instantiates a new Getoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*Getoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    m := &Getoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365GroupsActivityStorage(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder instantiates a new Getoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get the total storage used across all group mailboxes and group sites.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getoffice365groupsactivitystorage?view=graph-rest-1.0
func (m *Getoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *Getoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get the total storage used across all group mailboxes and group sites.
// returns a *RequestInformation when successful
func (m *Getoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *Getoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder when successful
func (m *Getoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder) WithUrl(rawUrl string)(*Getoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    return NewGetoffice365groupsactivitystoragewithperiodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
