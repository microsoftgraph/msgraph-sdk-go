package groupsettingtemplates

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    i5c2514cac71bd42e8d3ff0b5f2392805dbb5175ce8500f2f4066069fd3f8a025 "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/validateproperties"
    i876aea553bfa1c412214415b5c2d08dc7ae79cdcecf4bf4e6f912c57c0de078d "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/getavailableextensionproperties"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    iccb6d65dafe6f1cf697180a62c7210b026edfb6f165ab42626bb6c7cf8a7836a "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/getbyids"
    ie2feb4ec9a72c319a82ed8b82d8ae3403e819512010e7c6dd097f3426e334809 "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/count"
)

// GroupSettingTemplatesRequestBuilder provides operations to manage the collection of groupSettingTemplate entities.
type GroupSettingTemplatesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GroupSettingTemplatesRequestBuilderGetOptions options for Get
type GroupSettingTemplatesRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *GroupSettingTemplatesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// GroupSettingTemplatesRequestBuilderGetQueryParameters get entities from groupSettingTemplates
type GroupSettingTemplatesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// GroupSettingTemplatesRequestBuilderPostOptions options for Post
type GroupSettingTemplatesRequestBuilderPostOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.GroupSettingTemplateable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewGroupSettingTemplatesRequestBuilderInternal instantiates a new GroupSettingTemplatesRequestBuilder and sets the default values.
func NewGroupSettingTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupSettingTemplatesRequestBuilder) {
    m := &GroupSettingTemplatesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groupSettingTemplates{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupSettingTemplatesRequestBuilder instantiates a new GroupSettingTemplatesRequestBuilder and sets the default values.
func NewGroupSettingTemplatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupSettingTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupSettingTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *GroupSettingTemplatesRequestBuilder) Count()(*ie2feb4ec9a72c319a82ed8b82d8ae3403e819512010e7c6dd097f3426e334809.CountRequestBuilder) {
    return ie2feb4ec9a72c319a82ed8b82d8ae3403e819512010e7c6dd097f3426e334809.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from groupSettingTemplates
func (m *GroupSettingTemplatesRequestBuilder) CreateGetRequestInformation(options *GroupSettingTemplatesRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to groupSettingTemplates
func (m *GroupSettingTemplatesRequestBuilder) CreatePostRequestInformation(options *GroupSettingTemplatesRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
// Get get entities from groupSettingTemplates
func (m *GroupSettingTemplatesRequestBuilder) Get(options *GroupSettingTemplatesRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.GroupSettingTemplateCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateGroupSettingTemplateCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.GroupSettingTemplateCollectionResponseable), nil
}
// GetAvailableExtensionProperties the getAvailableExtensionProperties property
func (m *GroupSettingTemplatesRequestBuilder) GetAvailableExtensionProperties()(*i876aea553bfa1c412214415b5c2d08dc7ae79cdcecf4bf4e6f912c57c0de078d.GetAvailableExtensionPropertiesRequestBuilder) {
    return i876aea553bfa1c412214415b5c2d08dc7ae79cdcecf4bf4e6f912c57c0de078d.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetByIds the getByIds property
func (m *GroupSettingTemplatesRequestBuilder) GetByIds()(*iccb6d65dafe6f1cf697180a62c7210b026edfb6f165ab42626bb6c7cf8a7836a.GetByIdsRequestBuilder) {
    return iccb6d65dafe6f1cf697180a62c7210b026edfb6f165ab42626bb6c7cf8a7836a.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to groupSettingTemplates
func (m *GroupSettingTemplatesRequestBuilder) Post(options *GroupSettingTemplatesRequestBuilderPostOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.GroupSettingTemplateable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
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
// ValidateProperties the validateProperties property
func (m *GroupSettingTemplatesRequestBuilder) ValidateProperties()(*i5c2514cac71bd42e8d3ff0b5f2392805dbb5175ce8500f2f4066069fd3f8a025.ValidatePropertiesRequestBuilder) {
    return i5c2514cac71bd42e8d3ff0b5f2392805dbb5175ce8500f2f4066069fd3f8a025.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
