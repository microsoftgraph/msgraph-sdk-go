package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder provides operations to call the setUpResourcesFolder method.
type ClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilderInternal instantiates a new ClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder and sets the default values.
func NewClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder) {
    m := &ClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/classes/{educationClass%2Did}/modules/{educationModule%2Did}/setUpResourcesFolder", pathParameters),
    }
    return m
}
// NewClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder instantiates a new ClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder and sets the default values.
func NewClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a SharePoint folder to upload files for a given educationModule. Only teachers can perform this operation. The teacher determines what resources to upload to the SharePoint folder for the module.
// returns a EducationModuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationmodule-setupresourcesfolder?view=graph-rest-1.0
func (m *ClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder) Post(ctx context.Context, requestConfiguration *ClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationModuleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationModuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationModuleable), nil
}
// ToPostRequestInformation create a SharePoint folder to upload files for a given educationModule. Only teachers can perform this operation. The teacher determines what resources to upload to the SharePoint folder for the module.
// returns a *RequestInformation when successful
func (m *ClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder when successful
func (m *ClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder) WithUrl(rawUrl string)(*ClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder) {
    return NewClassesItemModulesItemSetupresourcesfolderSetUpResourcesFolderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
