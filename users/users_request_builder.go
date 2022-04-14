package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    i27af31442f114a2bae9be3024351e4faa9f5343a89396b4ddd33b593dab2accb "github.com/microsoftgraph/msgraph-sdk-go/users/count"
    i36fc530779481c7e7fe57e514468d077da1dea7070a4bc1d2219c7a5565008ce "github.com/microsoftgraph/msgraph-sdk-go/users/delta"
    i4d70decf9494085a15e038ef9296d113e35f32634231ffa1f8b3948883f1768e "github.com/microsoftgraph/msgraph-sdk-go/users/getavailableextensionproperties"
    i699fa55acd10fc78e979097bfc57c32ad76275d14ca7f9670c41500cf15f0183 "github.com/microsoftgraph/msgraph-sdk-go/users/validateproperties"
    ia25a03c7bdb06bb4807fe5a1efe7c3631ba16f526700224ff61a9cbe9fb09770 "github.com/microsoftgraph/msgraph-sdk-go/users/getbyids"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UsersRequestBuilder provides operations to manage the collection of user entities.
type UsersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersRequestBuilderGetOptions options for Get
type UsersRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// UsersRequestBuilderGetQueryParameters get entities from users
type UsersRequestBuilderGetQueryParameters struct {
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
// UsersRequestBuilderPostOptions options for Post
type UsersRequestBuilderPostOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewUsersRequestBuilderInternal instantiates a new UsersRequestBuilder and sets the default values.
func NewUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersRequestBuilder) {
    m := &UsersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersRequestBuilder instantiates a new UsersRequestBuilder and sets the default values.
func NewUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *UsersRequestBuilder) Count()(*i27af31442f114a2bae9be3024351e4faa9f5343a89396b4ddd33b593dab2accb.CountRequestBuilder) {
    return i27af31442f114a2bae9be3024351e4faa9f5343a89396b4ddd33b593dab2accb.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from users
func (m *UsersRequestBuilder) CreateGetRequestInformation(options *UsersRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to users
func (m *UsersRequestBuilder) CreatePostRequestInformation(options *UsersRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersRequestBuilder) Delta()(*i36fc530779481c7e7fe57e514468d077da1dea7070a4bc1d2219c7a5565008ce.DeltaRequestBuilder) {
    return i36fc530779481c7e7fe57e514468d077da1dea7070a4bc1d2219c7a5565008ce.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entities from users
func (m *UsersRequestBuilder) Get(options *UsersRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserCollectionResponseable), nil
}
// GetAvailableExtensionProperties the getAvailableExtensionProperties property
func (m *UsersRequestBuilder) GetAvailableExtensionProperties()(*i4d70decf9494085a15e038ef9296d113e35f32634231ffa1f8b3948883f1768e.GetAvailableExtensionPropertiesRequestBuilder) {
    return i4d70decf9494085a15e038ef9296d113e35f32634231ffa1f8b3948883f1768e.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetByIds the getByIds property
func (m *UsersRequestBuilder) GetByIds()(*ia25a03c7bdb06bb4807fe5a1efe7c3631ba16f526700224ff61a9cbe9fb09770.GetByIdsRequestBuilder) {
    return ia25a03c7bdb06bb4807fe5a1efe7c3631ba16f526700224ff61a9cbe9fb09770.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to users
func (m *UsersRequestBuilder) Post(options *UsersRequestBuilderPostOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable), nil
}
// ValidateProperties the validateProperties property
func (m *UsersRequestBuilder) ValidateProperties()(*i699fa55acd10fc78e979097bfc57c32ad76275d14ca7f9670c41500cf15f0183.ValidatePropertiesRequestBuilder) {
    return i699fa55acd10fc78e979097bfc57c32ad76275d14ca7f9670c41500cf15f0183.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
