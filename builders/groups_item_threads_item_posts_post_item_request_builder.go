package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GroupsItemThreadsItemPostsPostItemRequestBuilder provides operations to manage the posts property of the microsoft.graph.conversationThread entity.
type GroupsItemThreadsItemPostsPostItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemThreadsItemPostsPostItemRequestBuilderGetQueryParameters get posts from groups
type GroupsItemThreadsItemPostsPostItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemThreadsItemPostsPostItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemThreadsItemPostsPostItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemThreadsItemPostsPostItemRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.post entity.
func (m *GroupsItemThreadsItemPostsPostItemRequestBuilder) Attachments()(*GroupsItemThreadsItemPostsItemAttachmentsRequestBuilder) {
    return NewGroupsItemThreadsItemPostsItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.post entity.
func (m *GroupsItemThreadsItemPostsPostItemRequestBuilder) AttachmentsById(id string)(*GroupsItemThreadsItemPostsItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewGroupsItemThreadsItemPostsItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewGroupsItemThreadsItemPostsPostItemRequestBuilderInternal instantiates a new PostItemRequestBuilder and sets the default values.
func NewGroupsItemThreadsItemPostsPostItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemThreadsItemPostsPostItemRequestBuilder) {
    m := &GroupsItemThreadsItemPostsPostItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemThreadsItemPostsPostItemRequestBuilder instantiates a new PostItemRequestBuilder and sets the default values.
func NewGroupsItemThreadsItemPostsPostItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemThreadsItemPostsPostItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemThreadsItemPostsPostItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get posts from groups
func (m *GroupsItemThreadsItemPostsPostItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemThreadsItemPostsPostItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.post entity.
func (m *GroupsItemThreadsItemPostsPostItemRequestBuilder) Extensions()(*GroupsItemThreadsItemPostsItemExtensionsRequestBuilder) {
    return NewGroupsItemThreadsItemPostsItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.post entity.
func (m *GroupsItemThreadsItemPostsPostItemRequestBuilder) ExtensionsById(id string)(*GroupsItemThreadsItemPostsItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewGroupsItemThreadsItemPostsItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *GroupsItemThreadsItemPostsPostItemRequestBuilder) Forward()(*GroupsItemThreadsItemPostsItemForwardRequestBuilder) {
    return NewGroupsItemThreadsItemPostsItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get posts from groups
func (m *GroupsItemThreadsItemPostsPostItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemThreadsItemPostsPostItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Postable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePostFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Postable), nil
}
// InReplyTo provides operations to manage the inReplyTo property of the microsoft.graph.post entity.
func (m *GroupsItemThreadsItemPostsPostItemRequestBuilder) InReplyTo()(*GroupsItemThreadsItemPostsItemInReplyToRequestBuilder) {
    return NewGroupsItemThreadsItemPostsItemInReplyToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.post entity.
func (m *GroupsItemThreadsItemPostsPostItemRequestBuilder) MultiValueExtendedProperties()(*GroupsItemThreadsItemPostsItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewGroupsItemThreadsItemPostsItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.post entity.
func (m *GroupsItemThreadsItemPostsPostItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*GroupsItemThreadsItemPostsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewGroupsItemThreadsItemPostsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Reply provides operations to call the reply method.
func (m *GroupsItemThreadsItemPostsPostItemRequestBuilder) Reply()(*GroupsItemThreadsItemPostsItemReplyRequestBuilder) {
    return NewGroupsItemThreadsItemPostsItemReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.post entity.
func (m *GroupsItemThreadsItemPostsPostItemRequestBuilder) SingleValueExtendedProperties()(*GroupsItemThreadsItemPostsItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewGroupsItemThreadsItemPostsItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.post entity.
func (m *GroupsItemThreadsItemPostsPostItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*GroupsItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewGroupsItemThreadsItemPostsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
