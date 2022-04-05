package contacts

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    i3c25de93fe491a45d1028201d9a97a3867c994bf94910fbb23c75589f3064367 "github.com/microsoftgraph/msgraph-sdk-go/contacts/count"
    i56415aaf9483a2e973df3b1a2fe488e7b2d8844c5d37e84bfa9789655c1de234 "github.com/microsoftgraph/msgraph-sdk-go/contacts/validateproperties"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibd6b6137a854fdc0082d98413fd3f4475189c1b69ecd839ed41d01e72ef32af8 "github.com/microsoftgraph/msgraph-sdk-go/contacts/getbyids"
    icd7c2e08bbb98a59371e3c25cbc42e41894aefcd3c07a9dda51cde0fbcbda0a1 "github.com/microsoftgraph/msgraph-sdk-go/contacts/getavailableextensionproperties"
    id6a5f8a255008661c4f3891226b4705b4315a990d81bf1bdc10e110af7ce62f9 "github.com/microsoftgraph/msgraph-sdk-go/contacts/delta"
)

// ContactsRequestBuilder provides operations to manage the collection of orgContact entities.
type ContactsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContactsRequestBuilderGetOptions options for Get
type ContactsRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *ContactsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ContactsRequestBuilderGetQueryParameters get entities from contacts
type ContactsRequestBuilderGetQueryParameters struct {
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
// ContactsRequestBuilderPostOptions options for Post
type ContactsRequestBuilderPostOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrgContactable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewContactsRequestBuilderInternal instantiates a new ContactsRequestBuilder and sets the default values.
func NewContactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContactsRequestBuilder) {
    m := &ContactsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contacts{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContactsRequestBuilder instantiates a new ContactsRequestBuilder and sets the default values.
func NewContactsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContactsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *ContactsRequestBuilder) Count()(*i3c25de93fe491a45d1028201d9a97a3867c994bf94910fbb23c75589f3064367.CountRequestBuilder) {
    return i3c25de93fe491a45d1028201d9a97a3867c994bf94910fbb23c75589f3064367.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from contacts
func (m *ContactsRequestBuilder) CreateGetRequestInformation(options *ContactsRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to contacts
func (m *ContactsRequestBuilder) CreatePostRequestInformation(options *ContactsRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delta provides operations to call the delta method.
func (m *ContactsRequestBuilder) Delta()(*id6a5f8a255008661c4f3891226b4705b4315a990d81bf1bdc10e110af7ce62f9.DeltaRequestBuilder) {
    return id6a5f8a255008661c4f3891226b4705b4315a990d81bf1bdc10e110af7ce62f9.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entities from contacts
func (m *ContactsRequestBuilder) Get(options *ContactsRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrgContactCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOrgContactCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrgContactCollectionResponseable), nil
}
// GetAvailableExtensionProperties the getAvailableExtensionProperties property
func (m *ContactsRequestBuilder) GetAvailableExtensionProperties()(*icd7c2e08bbb98a59371e3c25cbc42e41894aefcd3c07a9dda51cde0fbcbda0a1.GetAvailableExtensionPropertiesRequestBuilder) {
    return icd7c2e08bbb98a59371e3c25cbc42e41894aefcd3c07a9dda51cde0fbcbda0a1.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetByIds the getByIds property
func (m *ContactsRequestBuilder) GetByIds()(*ibd6b6137a854fdc0082d98413fd3f4475189c1b69ecd839ed41d01e72ef32af8.GetByIdsRequestBuilder) {
    return ibd6b6137a854fdc0082d98413fd3f4475189c1b69ecd839ed41d01e72ef32af8.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to contacts
func (m *ContactsRequestBuilder) Post(options *ContactsRequestBuilderPostOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrgContactable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOrgContactFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrgContactable), nil
}
// ValidateProperties the validateProperties property
func (m *ContactsRequestBuilder) ValidateProperties()(*i56415aaf9483a2e973df3b1a2fe488e7b2d8844c5d37e84bfa9789655c1de234.ValidatePropertiesRequestBuilder) {
    return i56415aaf9483a2e973df3b1a2fe488e7b2d8844c5d37e84bfa9789655c1de234.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
