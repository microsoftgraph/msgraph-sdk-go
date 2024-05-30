package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder provides operations to call the reprocessLicenseAssignment method.
type ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilderInternal instantiates a new ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder and sets the default values.
func NewItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder) {
    m := &ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/reprocessLicenseAssignment", pathParameters),
    }
    return m
}
// NewItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder instantiates a new ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder and sets the default values.
func NewItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reprocess all group-based license assignments for the user. To learn more about group-based licensing, see What is group-based licensing in Microsoft Entra ID. Also see Identify and resolve license assignment problems for a group in Microsoft Entra ID for more details.
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-reprocesslicenseassignment?view=graph-rest-1.0
func (m *ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation reprocess all group-based license assignments for the user. To learn more about group-based licensing, see What is group-based licensing in Microsoft Entra ID. Also see Identify and resolve license assignment problems for a group in Microsoft Entra ID for more details.
// returns a *RequestInformation when successful
func (m *ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder when successful
func (m *ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder) WithUrl(rawUrl string)(*ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder) {
    return NewItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
