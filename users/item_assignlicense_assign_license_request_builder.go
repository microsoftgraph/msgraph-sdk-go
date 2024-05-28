package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAssignlicenseAssignLicenseRequestBuilder provides operations to call the assignLicense method.
type ItemAssignlicenseAssignLicenseRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAssignlicenseAssignLicenseRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAssignlicenseAssignLicenseRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAssignlicenseAssignLicenseRequestBuilderInternal instantiates a new ItemAssignlicenseAssignLicenseRequestBuilder and sets the default values.
func NewItemAssignlicenseAssignLicenseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAssignlicenseAssignLicenseRequestBuilder) {
    m := &ItemAssignlicenseAssignLicenseRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/assignLicense", pathParameters),
    }
    return m
}
// NewItemAssignlicenseAssignLicenseRequestBuilder instantiates a new ItemAssignlicenseAssignLicenseRequestBuilder and sets the default values.
func NewItemAssignlicenseAssignLicenseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAssignlicenseAssignLicenseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAssignlicenseAssignLicenseRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add or remove subscriptions for the user. You can also enable and disable specific plans associated with a subscription.
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-assignlicense?view=graph-rest-1.0
func (m *ItemAssignlicenseAssignLicenseRequestBuilder) Post(ctx context.Context, body ItemAssignlicenseAssignLicensePostRequestBodyable, requestConfiguration *ItemAssignlicenseAssignLicenseRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable), nil
}
// ToPostRequestInformation add or remove subscriptions for the user. You can also enable and disable specific plans associated with a subscription.
// returns a *RequestInformation when successful
func (m *ItemAssignlicenseAssignLicenseRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemAssignlicenseAssignLicensePostRequestBodyable, requestConfiguration *ItemAssignlicenseAssignLicenseRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAssignlicenseAssignLicenseRequestBuilder when successful
func (m *ItemAssignlicenseAssignLicenseRequestBuilder) WithUrl(rawUrl string)(*ItemAssignlicenseAssignLicenseRequestBuilder) {
    return NewItemAssignlicenseAssignLicenseRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
