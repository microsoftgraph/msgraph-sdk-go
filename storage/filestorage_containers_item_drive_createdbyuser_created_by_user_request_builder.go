package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
type FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilderGetQueryParameters identity of the user who created the item. Read-only.
type FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilderGetQueryParameters
}
// NewFilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilderInternal instantiates a new FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder) {
    m := &FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/createdByUser{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder instantiates a new FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Get identity of the user who created the item. Read-only.
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable), nil
}
// MailboxSettings the mailboxSettings property
// returns a *FilestorageContainersItemDriveCreatedbyuserMailboxsettingsMailboxSettingsRequestBuilder when successful
func (m *FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder) MailboxSettings()(*FilestorageContainersItemDriveCreatedbyuserMailboxsettingsMailboxSettingsRequestBuilder) {
    return NewFilestorageContainersItemDriveCreatedbyuserMailboxsettingsMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *FilestorageContainersItemDriveCreatedbyuserServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder) ServiceProvisioningErrors()(*FilestorageContainersItemDriveCreatedbyuserServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewFilestorageContainersItemDriveCreatedbyuserServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation identity of the user who created the item. Read-only.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder when successful
func (m *FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder) {
    return NewFilestorageContainersItemDriveCreatedbyuserCreatedByUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
