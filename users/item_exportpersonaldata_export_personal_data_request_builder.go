package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemExportpersonaldataExportPersonalDataRequestBuilder provides operations to call the exportPersonalData method.
type ItemExportpersonaldataExportPersonalDataRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemExportpersonaldataExportPersonalDataRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemExportpersonaldataExportPersonalDataRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemExportpersonaldataExportPersonalDataRequestBuilderInternal instantiates a new ItemExportpersonaldataExportPersonalDataRequestBuilder and sets the default values.
func NewItemExportpersonaldataExportPersonalDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemExportpersonaldataExportPersonalDataRequestBuilder) {
    m := &ItemExportpersonaldataExportPersonalDataRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/exportPersonalData", pathParameters),
    }
    return m
}
// NewItemExportpersonaldataExportPersonalDataRequestBuilder instantiates a new ItemExportpersonaldataExportPersonalDataRequestBuilder and sets the default values.
func NewItemExportpersonaldataExportPersonalDataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemExportpersonaldataExportPersonalDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemExportpersonaldataExportPersonalDataRequestBuilderInternal(urlParams, requestAdapter)
}
// Post submit a data policy operation request from a company administrator or an application to export an organizational user's data. This data includes the user's data stored in OneDrive and their activity reports. For more information about exporting data while complying with regulations, see Data Subject Requests and the GDPR and CCPA.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-exportpersonaldata?view=graph-rest-1.0
func (m *ItemExportpersonaldataExportPersonalDataRequestBuilder) Post(ctx context.Context, body ItemExportpersonaldataExportPersonalDataPostRequestBodyable, requestConfiguration *ItemExportpersonaldataExportPersonalDataRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation submit a data policy operation request from a company administrator or an application to export an organizational user's data. This data includes the user's data stored in OneDrive and their activity reports. For more information about exporting data while complying with regulations, see Data Subject Requests and the GDPR and CCPA.
// returns a *RequestInformation when successful
func (m *ItemExportpersonaldataExportPersonalDataRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemExportpersonaldataExportPersonalDataPostRequestBodyable, requestConfiguration *ItemExportpersonaldataExportPersonalDataRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemExportpersonaldataExportPersonalDataRequestBuilder when successful
func (m *ItemExportpersonaldataExportPersonalDataRequestBuilder) WithUrl(rawUrl string)(*ItemExportpersonaldataExportPersonalDataRequestBuilder) {
    return NewItemExportpersonaldataExportPersonalDataRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
