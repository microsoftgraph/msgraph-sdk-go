package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder provides operations to manage the exemptAppLockerFiles property of the microsoft.graph.windowsInformationProtection entity.
type WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetQueryParameters read properties and relationships of the windowsInformationProtectionAppLockerFile object.
type WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetQueryParameters
}
// WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderInternal instantiates a new WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder and sets the default values.
func NewWindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) {
    m := &WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/windowsInformationProtectionPolicies/{windowsInformationProtectionPolicy%2Did}/exemptAppLockerFiles/{windowsInformationProtectionAppLockerFile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder instantiates a new WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder and sets the default values.
func NewWindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a windowsInformationProtectionAppLockerFile.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-windowsinformationprotectionapplockerfile-delete?view=graph-rest-1.0
func (m *WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read properties and relationships of the windowsInformationProtectionAppLockerFile object.
// returns a WindowsInformationProtectionAppLockerFileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-windowsinformationprotectionapplockerfile-get?view=graph-rest-1.0
func (m *WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsInformationProtectionAppLockerFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable), nil
}
// Patch update the properties of a windowsInformationProtectionAppLockerFile object.
// returns a WindowsInformationProtectionAppLockerFileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-windowsinformationprotectionapplockerfile-update?view=graph-rest-1.0
func (m *WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, requestConfiguration *WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsInformationProtectionAppLockerFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable), nil
}
// ToDeleteRequestInformation deletes a windowsInformationProtectionAppLockerFile.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the windowsInformationProtectionAppLockerFile object.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a windowsInformationProtectionAppLockerFile object.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, requestConfiguration *WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder when successful
func (m *WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) WithUrl(rawUrl string)(*WindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) {
    return NewWindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
