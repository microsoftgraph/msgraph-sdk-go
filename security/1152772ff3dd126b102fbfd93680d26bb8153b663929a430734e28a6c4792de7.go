// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilder provides operations to call the computeRightsAndInheritance method.
type DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilderInternal instantiates a new DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilder and sets the default values.
func NewDataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilder) {
    m := &DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/dataSecurityAndGovernance/sensitivityLabels/{sensitivityLabel%2Did}/sublabels/computeRightsAndInheritance", pathParameters),
    }
    return m
}
// NewDataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilder instantiates a new DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilder and sets the default values.
func NewDataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post computes the rights and inheritance for sensitivity labels based on the input content and labels.
// returns a ComputeRightsAndInheritanceResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/sensitivitylabel-computerightsandinheritance?view=graph-rest-1.0
func (m *DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilder) Post(ctx context.Context, body DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritancePostRequestBodyable, requestConfiguration *DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ComputeRightsAndInheritanceResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateComputeRightsAndInheritanceResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ComputeRightsAndInheritanceResultable), nil
}
// ToPostRequestInformation computes the rights and inheritance for sensitivity labels based on the input content and labels.
// returns a *RequestInformation when successful
func (m *DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilder) ToPostRequestInformation(ctx context.Context, body DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritancePostRequestBodyable, requestConfiguration *DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilder when successful
func (m *DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilder) WithUrl(rawUrl string)(*DataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilder) {
    return NewDataSecurityAndGovernanceSensitivityLabelsItemSublabelsComputeRightsAndInheritanceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
