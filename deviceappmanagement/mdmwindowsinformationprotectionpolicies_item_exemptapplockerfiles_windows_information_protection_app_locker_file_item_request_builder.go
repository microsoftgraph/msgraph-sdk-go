package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder provides operations to manage the exemptAppLockerFiles property of the microsoft.graph.windowsInformationProtection entity.
type MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetQueryParameters another way to input exempt apps through xml files
type MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetQueryParameters
}
// MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderInternal instantiates a new MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder and sets the default values.
func NewMdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) {
    m := &MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mdmWindowsInformationProtectionPolicies/{mdmWindowsInformationProtectionPolicy%2Did}/exemptAppLockerFiles/{windowsInformationProtectionAppLockerFile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder instantiates a new MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder and sets the default values.
func NewMdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property exemptAppLockerFiles for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get another way to input exempt apps through xml files
// returns a WindowsInformationProtectionAppLockerFileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, error) {
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
// Patch update the navigation property exemptAppLockerFiles in deviceAppManagement
// returns a WindowsInformationProtectionAppLockerFileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, error) {
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
// ToDeleteRequestInformation delete navigation property exemptAppLockerFiles for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation another way to input exempt apps through xml files
// returns a *RequestInformation when successful
func (m *MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property exemptAppLockerFiles in deviceAppManagement
// returns a *RequestInformation when successful
func (m *MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder when successful
func (m *MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) WithUrl(rawUrl string)(*MdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) {
    return NewMdmwindowsinformationprotectionpoliciesItemExemptapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
