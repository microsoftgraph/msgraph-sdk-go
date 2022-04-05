package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i414bcfa81cf62cb492a10891e6ade220e793a207338da61bbf717173d599c787 "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/item/checkmemberobjects"
    iaa01455e2185fdd184cd4993202a2e1f3eec805d38e4ba431cce0c3144cca8f4 "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/item/getmemberobjects"
    ic28c47ff94ea6e8ffffd08e5644e712f6344e2bc49deb132bfac7a2d04ceb662 "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/item/restore"
    iec1a8ab976accbeb3945c4eeab0a0c4f560960c9c1f9f5ebd942d29059b40629 "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/item/getmembergroups"
    iff1e45e5e14f79eaaea630c1948f4ffdd5e56ce522e6fc233cd011734be4178b "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/item/checkmembergroups"
)

// DirectoryRoleTemplateItemRequestBuilder provides operations to manage the collection of directoryRoleTemplate entities.
type DirectoryRoleTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryRoleTemplateItemRequestBuilderDeleteOptions options for Delete
type DirectoryRoleTemplateItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DirectoryRoleTemplateItemRequestBuilderGetOptions options for Get
type DirectoryRoleTemplateItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DirectoryRoleTemplateItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DirectoryRoleTemplateItemRequestBuilderGetQueryParameters get entity from directoryRoleTemplates by key
type DirectoryRoleTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DirectoryRoleTemplateItemRequestBuilderPatchOptions options for Patch
type DirectoryRoleTemplateItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryRoleTemplateable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// CheckMemberGroups the checkMemberGroups property
func (m *DirectoryRoleTemplateItemRequestBuilder) CheckMemberGroups()(*iff1e45e5e14f79eaaea630c1948f4ffdd5e56ce522e6fc233cd011734be4178b.CheckMemberGroupsRequestBuilder) {
    return iff1e45e5e14f79eaaea630c1948f4ffdd5e56ce522e6fc233cd011734be4178b.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *DirectoryRoleTemplateItemRequestBuilder) CheckMemberObjects()(*i414bcfa81cf62cb492a10891e6ade220e793a207338da61bbf717173d599c787.CheckMemberObjectsRequestBuilder) {
    return i414bcfa81cf62cb492a10891e6ade220e793a207338da61bbf717173d599c787.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryRoleTemplateItemRequestBuilderInternal instantiates a new DirectoryRoleTemplateItemRequestBuilder and sets the default values.
func NewDirectoryRoleTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleTemplateItemRequestBuilder) {
    m := &DirectoryRoleTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoleTemplates/{directoryRoleTemplate_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRoleTemplateItemRequestBuilder instantiates a new DirectoryRoleTemplateItemRequestBuilder and sets the default values.
func NewDirectoryRoleTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from directoryRoleTemplates
func (m *DirectoryRoleTemplateItemRequestBuilder) CreateDeleteRequestInformation(options *DirectoryRoleTemplateItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get entity from directoryRoleTemplates by key
func (m *DirectoryRoleTemplateItemRequestBuilder) CreateGetRequestInformation(options *DirectoryRoleTemplateItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in directoryRoleTemplates
func (m *DirectoryRoleTemplateItemRequestBuilder) CreatePatchRequestInformation(options *DirectoryRoleTemplateItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete entity from directoryRoleTemplates
func (m *DirectoryRoleTemplateItemRequestBuilder) Delete(options *DirectoryRoleTemplateItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entity from directoryRoleTemplates by key
func (m *DirectoryRoleTemplateItemRequestBuilder) Get(options *DirectoryRoleTemplateItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryRoleTemplateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryRoleTemplateFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryRoleTemplateable), nil
}
// GetMemberGroups the getMemberGroups property
func (m *DirectoryRoleTemplateItemRequestBuilder) GetMemberGroups()(*iec1a8ab976accbeb3945c4eeab0a0c4f560960c9c1f9f5ebd942d29059b40629.GetMemberGroupsRequestBuilder) {
    return iec1a8ab976accbeb3945c4eeab0a0c4f560960c9c1f9f5ebd942d29059b40629.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *DirectoryRoleTemplateItemRequestBuilder) GetMemberObjects()(*iaa01455e2185fdd184cd4993202a2e1f3eec805d38e4ba431cce0c3144cca8f4.GetMemberObjectsRequestBuilder) {
    return iaa01455e2185fdd184cd4993202a2e1f3eec805d38e4ba431cce0c3144cca8f4.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in directoryRoleTemplates
func (m *DirectoryRoleTemplateItemRequestBuilder) Patch(options *DirectoryRoleTemplateItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Restore the restore property
func (m *DirectoryRoleTemplateItemRequestBuilder) Restore()(*ic28c47ff94ea6e8ffffd08e5644e712f6344e2bc49deb132bfac7a2d04ceb662.RestoreRequestBuilder) {
    return ic28c47ff94ea6e8ffffd08e5644e712f6344e2bc49deb132bfac7a2d04ceb662.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
