package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    icc3fe1a1c8b964bc176be6a6d48621fc3548efaeadcb3b10e136af2f19764752 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/notificationmessagetemplates/item/localizednotificationmessages"
    if10e6b46a7479cb721a918b2370be92792065fdd20e80c33685b2f50841a32b9 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/notificationmessagetemplates/item/sendtestmessage"
    i079873d8184290e06e63c9a04520d4cbb886c99c078b4ea0a493ba22bc14901e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/notificationmessagetemplates/item/localizednotificationmessages/item"
)

// NotificationMessageTemplateItemRequestBuilder provides operations to manage the notificationMessageTemplates property of the microsoft.graph.deviceManagement entity.
type NotificationMessageTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NotificationMessageTemplateItemRequestBuilderDeleteOptions options for Delete
type NotificationMessageTemplateItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NotificationMessageTemplateItemRequestBuilderGetOptions options for Get
type NotificationMessageTemplateItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *NotificationMessageTemplateItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NotificationMessageTemplateItemRequestBuilderGetQueryParameters the Notification Message Templates.
type NotificationMessageTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// NotificationMessageTemplateItemRequestBuilderPatchOptions options for Patch
type NotificationMessageTemplateItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NotificationMessageTemplateable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewNotificationMessageTemplateItemRequestBuilderInternal instantiates a new NotificationMessageTemplateItemRequestBuilder and sets the default values.
func NewNotificationMessageTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NotificationMessageTemplateItemRequestBuilder) {
    m := &NotificationMessageTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/notificationMessageTemplates/{notificationMessageTemplate%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewNotificationMessageTemplateItemRequestBuilder instantiates a new NotificationMessageTemplateItemRequestBuilder and sets the default values.
func NewNotificationMessageTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NotificationMessageTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNotificationMessageTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property notificationMessageTemplates for deviceManagement
func (m *NotificationMessageTemplateItemRequestBuilder) CreateDeleteRequestInformation(options *NotificationMessageTemplateItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the Notification Message Templates.
func (m *NotificationMessageTemplateItemRequestBuilder) CreateGetRequestInformation(options *NotificationMessageTemplateItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property notificationMessageTemplates in deviceManagement
func (m *NotificationMessageTemplateItemRequestBuilder) CreatePatchRequestInformation(options *NotificationMessageTemplateItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property notificationMessageTemplates for deviceManagement
func (m *NotificationMessageTemplateItemRequestBuilder) Delete(options *NotificationMessageTemplateItemRequestBuilderDeleteOptions)(error) {
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
// Get the Notification Message Templates.
func (m *NotificationMessageTemplateItemRequestBuilder) Get(options *NotificationMessageTemplateItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NotificationMessageTemplateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateNotificationMessageTemplateFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NotificationMessageTemplateable), nil
}
// LocalizedNotificationMessages the localizedNotificationMessages property
func (m *NotificationMessageTemplateItemRequestBuilder) LocalizedNotificationMessages()(*icc3fe1a1c8b964bc176be6a6d48621fc3548efaeadcb3b10e136af2f19764752.LocalizedNotificationMessagesRequestBuilder) {
    return icc3fe1a1c8b964bc176be6a6d48621fc3548efaeadcb3b10e136af2f19764752.NewLocalizedNotificationMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocalizedNotificationMessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.notificationMessageTemplates.item.localizedNotificationMessages.item collection
func (m *NotificationMessageTemplateItemRequestBuilder) LocalizedNotificationMessagesById(id string)(*i079873d8184290e06e63c9a04520d4cbb886c99c078b4ea0a493ba22bc14901e.LocalizedNotificationMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["localizedNotificationMessage%2Did"] = id
    }
    return i079873d8184290e06e63c9a04520d4cbb886c99c078b4ea0a493ba22bc14901e.NewLocalizedNotificationMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property notificationMessageTemplates in deviceManagement
func (m *NotificationMessageTemplateItemRequestBuilder) Patch(options *NotificationMessageTemplateItemRequestBuilderPatchOptions)(error) {
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
// SendTestMessage the sendTestMessage property
func (m *NotificationMessageTemplateItemRequestBuilder) SendTestMessage()(*if10e6b46a7479cb721a918b2370be92792065fdd20e80c33685b2f50841a32b9.SendTestMessageRequestBuilder) {
    return if10e6b46a7479cb721a918b2370be92792065fdd20e80c33685b2f50841a32b9.NewSendTestMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
