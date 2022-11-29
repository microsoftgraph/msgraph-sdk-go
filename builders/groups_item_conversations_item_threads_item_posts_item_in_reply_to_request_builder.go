package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder provides operations to manage the inReplyTo property of the microsoft.graph.post entity.
type GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetQueryParameters read-only. Supports $expand.
type GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.post entity.
func (m *GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) Attachments()(*GroupsItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsRequestBuilder) {
    return NewGroupsItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.post entity.
func (m *GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) AttachmentsById(id string)(*GroupsItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewGroupsItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewGroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderInternal instantiates a new InReplyToRequestBuilder and sets the default values.
func NewGroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) {
    m := &GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/conversations/{conversation%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/inReplyTo{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder instantiates a new InReplyToRequestBuilder and sets the default values.
func NewGroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation read-only. Supports $expand.
func (m *GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) Extensions()(*GroupsItemConversationsItemThreadsItemPostsItemInReplyToExtensionsRequestBuilder) {
    return NewGroupsItemConversationsItemThreadsItemPostsItemInReplyToExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.post entity.
func (m *GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) ExtensionsById(id string)(*GroupsItemConversationsItemThreadsItemPostsItemInReplyToExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewGroupsItemConversationsItemThreadsItemPostsItemInReplyToExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) Forward()(*GroupsItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilder) {
    return NewGroupsItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get read-only. Supports $expand.
func (m *GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Postable, error) {
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
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.post entity.
func (m *GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) MultiValueExtendedProperties()(*GroupsItemConversationsItemThreadsItemPostsItemInReplyToMultiValueExtendedPropertiesRequestBuilder) {
    return NewGroupsItemConversationsItemThreadsItemPostsItemInReplyToMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.post entity.
func (m *GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) MultiValueExtendedPropertiesById(id string)(*GroupsItemConversationsItemThreadsItemPostsItemInReplyToMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewGroupsItemConversationsItemThreadsItemPostsItemInReplyToMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Reply provides operations to call the reply method.
func (m *GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) Reply()(*GroupsItemConversationsItemThreadsItemPostsItemInReplyToReplyRequestBuilder) {
    return NewGroupsItemConversationsItemThreadsItemPostsItemInReplyToReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.post entity.
func (m *GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) SingleValueExtendedProperties()(*GroupsItemConversationsItemThreadsItemPostsItemInReplyToSingleValueExtendedPropertiesRequestBuilder) {
    return NewGroupsItemConversationsItemThreadsItemPostsItemInReplyToSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.post entity.
func (m *GroupsItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) SingleValueExtendedPropertiesById(id string)(*GroupsItemConversationsItemThreadsItemPostsItemInReplyToSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewGroupsItemConversationsItemThreadsItemPostsItemInReplyToSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
