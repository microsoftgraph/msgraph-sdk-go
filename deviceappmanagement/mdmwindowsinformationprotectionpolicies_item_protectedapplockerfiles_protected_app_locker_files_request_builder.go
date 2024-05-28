package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder provides operations to manage the protectedAppLockerFiles property of the microsoft.graph.windowsInformationProtection entity.
type MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderGetQueryParameters another way to input protected apps through xml files
type MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderGetQueryParameters
}
// MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWindowsInformationProtectionAppLockerFileId provides operations to manage the protectedAppLockerFiles property of the microsoft.graph.windowsInformationProtection entity.
// returns a *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder when successful
func (m *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) ByWindowsInformationProtectionAppLockerFileId(windowsInformationProtectionAppLockerFileId string)(*MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if windowsInformationProtectionAppLockerFileId != "" {
        urlTplParams["windowsInformationProtectionAppLockerFile%2Did"] = windowsInformationProtectionAppLockerFileId
    }
    return NewMdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesWindowsInformationProtectionAppLockerFileItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderInternal instantiates a new MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder and sets the default values.
func NewMdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) {
    m := &MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mdmWindowsInformationProtectionPolicies/{mdmWindowsInformationProtectionPolicy%2Did}/protectedAppLockerFiles{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder instantiates a new MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder and sets the default values.
func NewMdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder when successful
func (m *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) Count()(*MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilder) {
    return NewMdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get another way to input protected apps through xml files
// returns a WindowsInformationProtectionAppLockerFileCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) Get(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWindowsInformationProtectionAppLockerFileCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileCollectionResponseable), nil
}
// Post create new navigation property to protectedAppLockerFiles for deviceAppManagement
// returns a WindowsInformationProtectionAppLockerFileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation another way to input protected apps through xml files
// returns a *RequestInformation when successful
func (m *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to protectedAppLockerFiles for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WindowsInformationProtectionAppLockerFileable, requestConfiguration *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder when successful
func (m *MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) WithUrl(rawUrl string)(*MdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder) {
    return NewMdmwindowsinformationprotectionpoliciesItemProtectedapplockerfilesProtectedAppLockerFilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
