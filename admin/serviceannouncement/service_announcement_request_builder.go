package serviceannouncement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i5a1d633dc34eb5609de7c5475bea4c4ef223726f894d6018982fd790c70858ff "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/messages"
    i64b9866ddf18d005e45ccbc515fa8deeeb01d6d4ad4b07bdd67fd26595e01984 "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/issues"
    i993eabb4740237d5836875ca4abee991a6d6f1070cef12f2f5cc6ca062b03f47 "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/healthoverviews"
    i05cb520997f5472376a5296872f494ab5bc6678ca943b0a20d8a0a735c91b06c "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/healthoverviews/item"
    i2b040537addd757f1ad668fb2339fe6e9508193ac05e02dec09691e60f070df3 "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/issues/item"
    ib231b71ad312efbdfea743bca259e9f04366799c7f96005b7055ab09c71f271c "github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/messages/item"
)

// ServiceAnnouncementRequestBuilder provides operations to manage the serviceAnnouncement property of the microsoft.graph.admin entity.
type ServiceAnnouncementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServiceAnnouncementRequestBuilderDeleteOptions options for Delete
type ServiceAnnouncementRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ServiceAnnouncementRequestBuilderGetOptions options for Get
type ServiceAnnouncementRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceAnnouncementRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ServiceAnnouncementRequestBuilderGetQueryParameters a container for service communications resources. Read-only.
type ServiceAnnouncementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServiceAnnouncementRequestBuilderPatchOptions options for Patch
type ServiceAnnouncementRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceAnnouncementable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewServiceAnnouncementRequestBuilderInternal instantiates a new ServiceAnnouncementRequestBuilder and sets the default values.
func NewServiceAnnouncementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementRequestBuilder) {
    m := &ServiceAnnouncementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/serviceAnnouncement{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServiceAnnouncementRequestBuilder instantiates a new ServiceAnnouncementRequestBuilder and sets the default values.
func NewServiceAnnouncementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceAnnouncementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property serviceAnnouncement for admin
func (m *ServiceAnnouncementRequestBuilder) CreateDeleteRequestInformation(options *ServiceAnnouncementRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a container for service communications resources. Read-only.
func (m *ServiceAnnouncementRequestBuilder) CreateGetRequestInformation(options *ServiceAnnouncementRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property serviceAnnouncement in admin
func (m *ServiceAnnouncementRequestBuilder) CreatePatchRequestInformation(options *ServiceAnnouncementRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property serviceAnnouncement for admin
func (m *ServiceAnnouncementRequestBuilder) Delete(options *ServiceAnnouncementRequestBuilderDeleteOptions)(error) {
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
// Get a container for service communications resources. Read-only.
func (m *ServiceAnnouncementRequestBuilder) Get(options *ServiceAnnouncementRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceAnnouncementable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServiceAnnouncementFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceAnnouncementable), nil
}
// HealthOverviews the healthOverviews property
func (m *ServiceAnnouncementRequestBuilder) HealthOverviews()(*i993eabb4740237d5836875ca4abee991a6d6f1070cef12f2f5cc6ca062b03f47.HealthOverviewsRequestBuilder) {
    return i993eabb4740237d5836875ca4abee991a6d6f1070cef12f2f5cc6ca062b03f47.NewHealthOverviewsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HealthOverviewsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.admin.serviceAnnouncement.healthOverviews.item collection
func (m *ServiceAnnouncementRequestBuilder) HealthOverviewsById(id string)(*i05cb520997f5472376a5296872f494ab5bc6678ca943b0a20d8a0a735c91b06c.ServiceHealthItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["serviceHealth%2Did"] = id
    }
    return i05cb520997f5472376a5296872f494ab5bc6678ca943b0a20d8a0a735c91b06c.NewServiceHealthItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Issues the issues property
func (m *ServiceAnnouncementRequestBuilder) Issues()(*i64b9866ddf18d005e45ccbc515fa8deeeb01d6d4ad4b07bdd67fd26595e01984.IssuesRequestBuilder) {
    return i64b9866ddf18d005e45ccbc515fa8deeeb01d6d4ad4b07bdd67fd26595e01984.NewIssuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IssuesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.admin.serviceAnnouncement.issues.item collection
func (m *ServiceAnnouncementRequestBuilder) IssuesById(id string)(*i2b040537addd757f1ad668fb2339fe6e9508193ac05e02dec09691e60f070df3.ServiceHealthIssueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["serviceHealthIssue%2Did"] = id
    }
    return i2b040537addd757f1ad668fb2339fe6e9508193ac05e02dec09691e60f070df3.NewServiceHealthIssueItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *ServiceAnnouncementRequestBuilder) Messages()(*i5a1d633dc34eb5609de7c5475bea4c4ef223726f894d6018982fd790c70858ff.MessagesRequestBuilder) {
    return i5a1d633dc34eb5609de7c5475bea4c4ef223726f894d6018982fd790c70858ff.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.admin.serviceAnnouncement.messages.item collection
func (m *ServiceAnnouncementRequestBuilder) MessagesById(id string)(*ib231b71ad312efbdfea743bca259e9f04366799c7f96005b7055ab09c71f271c.ServiceUpdateMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["serviceUpdateMessage%2Did"] = id
    }
    return ib231b71ad312efbdfea743bca259e9f04366799c7f96005b7055ab09c71f271c.NewServiceUpdateMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property serviceAnnouncement in admin
func (m *ServiceAnnouncementRequestBuilder) Patch(options *ServiceAnnouncementRequestBuilderPatchOptions)(error) {
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
