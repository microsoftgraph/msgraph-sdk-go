package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i535340b77c27094ee236bb42cd6d6ecb0741e1f8ed70e47b43e92472c1f9535c "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/item/checkmemberobjects"
    ib24705b5e21850aefb572da669580b25cff74a7513ba53674efeff14fe51e5cb "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/item/getmemberobjects"
    ic9284ba20a018bdf9cc611bd4cc87612370fd87e308417912c51e24a09a84b7b "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/item/checkmembergroups"
    id343ec0d598ad136300b8599ebf6da41179c7224734d67f8c9ca6914ab993a9c "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/item/getmembergroups"
    if0de78395fe1dccb1dbdb7b7fe16ca5dd70edd86fe4f481d0b82bca471e231f0 "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/item/restore"
)

// GroupSettingTemplateItemRequestBuilder provides operations to manage the collection of groupSettingTemplate entities.
type GroupSettingTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupSettingTemplateItemRequestBuilderDeleteOptions options for Delete
type GroupSettingTemplateItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// GroupSettingTemplateItemRequestBuilderGetOptions options for Get
type GroupSettingTemplateItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupSettingTemplateItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// GroupSettingTemplateItemRequestBuilderGetQueryParameters get entity from groupSettingTemplates by key
type GroupSettingTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupSettingTemplateItemRequestBuilderPatchOptions options for Patch
type GroupSettingTemplateItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.GroupSettingTemplateable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CheckMemberGroups the checkMemberGroups property
func (m *GroupSettingTemplateItemRequestBuilder) CheckMemberGroups()(*ic9284ba20a018bdf9cc611bd4cc87612370fd87e308417912c51e24a09a84b7b.CheckMemberGroupsRequestBuilder) {
    return ic9284ba20a018bdf9cc611bd4cc87612370fd87e308417912c51e24a09a84b7b.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *GroupSettingTemplateItemRequestBuilder) CheckMemberObjects()(*i535340b77c27094ee236bb42cd6d6ecb0741e1f8ed70e47b43e92472c1f9535c.CheckMemberObjectsRequestBuilder) {
    return i535340b77c27094ee236bb42cd6d6ecb0741e1f8ed70e47b43e92472c1f9535c.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupSettingTemplateItemRequestBuilderInternal instantiates a new GroupSettingTemplateItemRequestBuilder and sets the default values.
func NewGroupSettingTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupSettingTemplateItemRequestBuilder) {
    m := &GroupSettingTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groupSettingTemplates/{groupSettingTemplate%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupSettingTemplateItemRequestBuilder instantiates a new GroupSettingTemplateItemRequestBuilder and sets the default values.
func NewGroupSettingTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupSettingTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupSettingTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from groupSettingTemplates
func (m *GroupSettingTemplateItemRequestBuilder) CreateDeleteRequestInformation(options *GroupSettingTemplateItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from groupSettingTemplates by key
func (m *GroupSettingTemplateItemRequestBuilder) CreateGetRequestInformation(options *GroupSettingTemplateItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in groupSettingTemplates
func (m *GroupSettingTemplateItemRequestBuilder) CreatePatchRequestInformation(options *GroupSettingTemplateItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete entity from groupSettingTemplates
func (m *GroupSettingTemplateItemRequestBuilder) Delete(options *GroupSettingTemplateItemRequestBuilderDeleteOptions)(error) {
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
// Get get entity from groupSettingTemplates by key
func (m *GroupSettingTemplateItemRequestBuilder) Get(options *GroupSettingTemplateItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.GroupSettingTemplateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateGroupSettingTemplateFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.GroupSettingTemplateable), nil
}
// GetMemberGroups the getMemberGroups property
func (m *GroupSettingTemplateItemRequestBuilder) GetMemberGroups()(*id343ec0d598ad136300b8599ebf6da41179c7224734d67f8c9ca6914ab993a9c.GetMemberGroupsRequestBuilder) {
    return id343ec0d598ad136300b8599ebf6da41179c7224734d67f8c9ca6914ab993a9c.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *GroupSettingTemplateItemRequestBuilder) GetMemberObjects()(*ib24705b5e21850aefb572da669580b25cff74a7513ba53674efeff14fe51e5cb.GetMemberObjectsRequestBuilder) {
    return ib24705b5e21850aefb572da669580b25cff74a7513ba53674efeff14fe51e5cb.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in groupSettingTemplates
func (m *GroupSettingTemplateItemRequestBuilder) Patch(options *GroupSettingTemplateItemRequestBuilderPatchOptions)(error) {
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
func (m *GroupSettingTemplateItemRequestBuilder) Restore()(*if0de78395fe1dccb1dbdb7b7fe16ca5dd70edd86fe4f481d0b82bca471e231f0.RestoreRequestBuilder) {
    return if0de78395fe1dccb1dbdb7b7fe16ca5dd70edd86fe4f481d0b82bca471e231f0.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
