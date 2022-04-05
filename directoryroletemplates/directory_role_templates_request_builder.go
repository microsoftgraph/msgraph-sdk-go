package directoryroletemplates

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    i13f35543e7fd513b7896b53959450496b4b40a0d8ed709962f6871dc35b86ec4 "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/getbyids"
    i14aa4ead11e74baa62edc9a283a7340417496098b3a9b5697fdc955d20086255 "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/validateproperties"
    i1deef55c0cacc5fe1d33525ca2ca081a8626edb9cd782f2be58da2bdabec7837 "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/getavailableextensionproperties"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ie9768601a46d4615c89194434b08253d4115813ca60fafddb482fa54036bb02d "github.com/microsoftgraph/msgraph-sdk-go/directoryroletemplates/count"
)

// DirectoryRoleTemplatesRequestBuilder provides operations to manage the collection of directoryRoleTemplate entities.
type DirectoryRoleTemplatesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryRoleTemplatesRequestBuilderGetOptions options for Get
type DirectoryRoleTemplatesRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DirectoryRoleTemplatesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DirectoryRoleTemplatesRequestBuilderGetQueryParameters get entities from directoryRoleTemplates
type DirectoryRoleTemplatesRequestBuilderGetQueryParameters struct {
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
}
// DirectoryRoleTemplatesRequestBuilderPostOptions options for Post
type DirectoryRoleTemplatesRequestBuilderPostOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryRoleTemplateable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewDirectoryRoleTemplatesRequestBuilderInternal instantiates a new DirectoryRoleTemplatesRequestBuilder and sets the default values.
func NewDirectoryRoleTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleTemplatesRequestBuilder) {
    m := &DirectoryRoleTemplatesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoleTemplates{?skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRoleTemplatesRequestBuilder instantiates a new DirectoryRoleTemplatesRequestBuilder and sets the default values.
func NewDirectoryRoleTemplatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *DirectoryRoleTemplatesRequestBuilder) Count()(*ie9768601a46d4615c89194434b08253d4115813ca60fafddb482fa54036bb02d.CountRequestBuilder) {
    return ie9768601a46d4615c89194434b08253d4115813ca60fafddb482fa54036bb02d.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from directoryRoleTemplates
func (m *DirectoryRoleTemplatesRequestBuilder) CreateGetRequestInformation(options *DirectoryRoleTemplatesRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to directoryRoleTemplates
func (m *DirectoryRoleTemplatesRequestBuilder) CreatePostRequestInformation(options *DirectoryRoleTemplatesRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get entities from directoryRoleTemplates
func (m *DirectoryRoleTemplatesRequestBuilder) Get(options *DirectoryRoleTemplatesRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryRoleTemplateCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryRoleTemplateCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryRoleTemplateCollectionResponseable), nil
}
// GetAvailableExtensionProperties the getAvailableExtensionProperties property
func (m *DirectoryRoleTemplatesRequestBuilder) GetAvailableExtensionProperties()(*i1deef55c0cacc5fe1d33525ca2ca081a8626edb9cd782f2be58da2bdabec7837.GetAvailableExtensionPropertiesRequestBuilder) {
    return i1deef55c0cacc5fe1d33525ca2ca081a8626edb9cd782f2be58da2bdabec7837.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetByIds the getByIds property
func (m *DirectoryRoleTemplatesRequestBuilder) GetByIds()(*i13f35543e7fd513b7896b53959450496b4b40a0d8ed709962f6871dc35b86ec4.GetByIdsRequestBuilder) {
    return i13f35543e7fd513b7896b53959450496b4b40a0d8ed709962f6871dc35b86ec4.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to directoryRoleTemplates
func (m *DirectoryRoleTemplatesRequestBuilder) Post(options *DirectoryRoleTemplatesRequestBuilderPostOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryRoleTemplateable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
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
// ValidateProperties the validateProperties property
func (m *DirectoryRoleTemplatesRequestBuilder) ValidateProperties()(*i14aa4ead11e74baa62edc9a283a7340417496098b3a9b5697fdc955d20086255.ValidatePropertiesRequestBuilder) {
    return i14aa4ead11e74baa62edc9a283a7340417496098b3a9b5697fdc955d20086255.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
