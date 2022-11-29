package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UsersUserItemRequestBuilder provides operations to manage the collection of user entities.
type UsersUserItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersUserItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersUserItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersUserItemRequestBuilderGetQueryParameters retrieve the properties and relationships of user object.
type UsersUserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersUserItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersUserItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersUserItemRequestBuilderGetQueryParameters
}
// UsersUserItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersUserItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Activities()(*UsersItemActivitiesRequestBuilder) {
    return NewUsersItemActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ActivitiesById(id string)(*UsersItemActivitiesUserActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userActivity%2Did"] = id
    }
    return NewUsersItemActivitiesUserActivityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AgreementAcceptances provides operations to manage the agreementAcceptances property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) AgreementAcceptances()(*UsersItemAgreementAcceptancesRequestBuilder) {
    return NewUsersItemAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgreementAcceptancesById provides operations to manage the agreementAcceptances property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) AgreementAcceptancesById(id string)(*UsersItemAgreementAcceptancesAgreementAcceptanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance%2Did"] = id
    }
    return NewUsersItemAgreementAcceptancesAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AppRoleAssignments provides operations to manage the appRoleAssignments property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) AppRoleAssignments()(*UsersItemAppRoleAssignmentsRequestBuilder) {
    return NewUsersItemAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById provides operations to manage the appRoleAssignments property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) AppRoleAssignmentsById(id string)(*UsersItemAppRoleAssignmentsAppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return NewUsersItemAppRoleAssignmentsAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignLicense provides operations to call the assignLicense method.
func (m *UsersUserItemRequestBuilder) AssignLicense()(*UsersItemAssignLicenseRequestBuilder) {
    return NewUsersItemAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Authentication provides operations to manage the authentication property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Authentication()(*UsersItemAuthenticationRequestBuilder) {
    return NewUsersItemAuthenticationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Calendar()(*UsersItemCalendarRequestBuilder) {
    return NewUsersItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarGroups provides operations to manage the calendarGroups property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) CalendarGroups()(*UsersItemCalendarGroupsRequestBuilder) {
    return NewUsersItemCalendarGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarGroupsById provides operations to manage the calendarGroups property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) CalendarGroupsById(id string)(*UsersItemCalendarGroupsCalendarGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarGroup%2Did"] = id
    }
    return NewUsersItemCalendarGroupsCalendarGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendars provides operations to manage the calendars property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Calendars()(*UsersItemCalendarsRequestBuilder) {
    return NewUsersItemCalendarsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarsById provides operations to manage the calendars property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) CalendarsById(id string)(*UsersItemCalendarsCalendarItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendar%2Did"] = id
    }
    return NewUsersItemCalendarsCalendarItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) CalendarView()(*UsersItemCalendarViewRequestBuilder) {
    return NewUsersItemCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById provides operations to manage the calendarView property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) CalendarViewById(id string)(*UsersItemCalendarViewEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewUsersItemCalendarViewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ChangePassword provides operations to call the changePassword method.
func (m *UsersUserItemRequestBuilder) ChangePassword()(*UsersItemChangePasswordRequestBuilder) {
    return NewUsersItemChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Chats provides operations to manage the chats property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Chats()(*UsersItemChatsRequestBuilder) {
    return NewUsersItemChatsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChatsById provides operations to manage the chats property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ChatsById(id string)(*UsersItemChatsChatItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chat%2Did"] = id
    }
    return NewUsersItemChatsChatItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *UsersUserItemRequestBuilder) CheckMemberGroups()(*UsersItemCheckMemberGroupsRequestBuilder) {
    return NewUsersItemCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *UsersUserItemRequestBuilder) CheckMemberObjects()(*UsersItemCheckMemberObjectsRequestBuilder) {
    return NewUsersItemCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUsersUserItemRequestBuilderInternal instantiates a new UserItemRequestBuilder and sets the default values.
func NewUsersUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersUserItemRequestBuilder) {
    m := &UsersUserItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersUserItemRequestBuilder instantiates a new UserItemRequestBuilder and sets the default values.
func NewUsersUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactFolders provides operations to manage the contactFolders property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ContactFolders()(*UsersItemContactFoldersRequestBuilder) {
    return NewUsersItemContactFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactFoldersById provides operations to manage the contactFolders property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ContactFoldersById(id string)(*UsersItemContactFoldersContactFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contactFolder%2Did"] = id
    }
    return NewUsersItemContactFoldersContactFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Contacts provides operations to manage the contacts property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Contacts()(*UsersItemContactsRequestBuilder) {
    return NewUsersItemContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactsById provides operations to manage the contacts property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ContactsById(id string)(*UsersItemContactsContactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact%2Did"] = id
    }
    return NewUsersItemContactsContactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete user.   When deleted, user resources are moved to a temporary container and can be restored within 30 days.  After that time, they are permanently deleted.  To learn more, see deletedItems.
func (m *UsersUserItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersUserItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatedObjects provides operations to manage the createdObjects property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) CreatedObjects()(*UsersItemCreatedObjectsRequestBuilder) {
    return NewUsersItemCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatedObjectsById provides operations to manage the createdObjects property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) CreatedObjectsById(id string)(*UsersItemCreatedObjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemCreatedObjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateGetRequestInformation retrieve the properties and relationships of user object.
func (m *UsersUserItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersUserItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
func (m *UsersUserItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, requestConfiguration *UsersUserItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete user.   When deleted, user resources are moved to a temporary container and can be restored within 30 days.  After that time, they are permanently deleted.  To learn more, see deletedItems.
func (m *UsersUserItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersUserItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceManagementTroubleshootingEvents provides operations to manage the deviceManagementTroubleshootingEvents property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) DeviceManagementTroubleshootingEvents()(*UsersItemDeviceManagementTroubleshootingEventsRequestBuilder) {
    return NewUsersItemDeviceManagementTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceManagementTroubleshootingEventsById provides operations to manage the deviceManagementTroubleshootingEvents property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) DeviceManagementTroubleshootingEventsById(id string)(*UsersItemDeviceManagementTroubleshootingEventsDeviceManagementTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTroubleshootingEvent%2Did"] = id
    }
    return NewUsersItemDeviceManagementTroubleshootingEventsDeviceManagementTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DirectReports provides operations to manage the directReports property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) DirectReports()(*UsersItemDirectReportsRequestBuilder) {
    return NewUsersItemDirectReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectReportsById provides operations to manage the directReports property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) DirectReportsById(id string)(*UsersItemDirectReportsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemDirectReportsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Drive provides operations to manage the drive property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Drive()(*UsersItemDriveRequestBuilder) {
    return NewUsersItemDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives provides operations to manage the drives property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Drives()(*UsersItemDrivesRequestBuilder) {
    return NewUsersItemDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById provides operations to manage the drives property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) DrivesById(id string)(*UsersItemDrivesDriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return NewUsersItemDrivesDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Events provides operations to manage the events property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Events()(*UsersItemEventsRequestBuilder) {
    return NewUsersItemEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById provides operations to manage the events property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) EventsById(id string)(*UsersItemEventsEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewUsersItemEventsEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExportPersonalData provides operations to call the exportPersonalData method.
func (m *UsersUserItemRequestBuilder) ExportPersonalData()(*UsersItemExportPersonalDataRequestBuilder) {
    return NewUsersItemExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Extensions()(*UsersItemExtensionsRequestBuilder) {
    return NewUsersItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ExtensionsById(id string)(*UsersItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewUsersItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FindMeetingTimes provides operations to call the findMeetingTimes method.
func (m *UsersUserItemRequestBuilder) FindMeetingTimes()(*UsersItemFindMeetingTimesRequestBuilder) {
    return NewUsersItemFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowedSites provides operations to manage the followedSites property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) FollowedSites()(*UsersItemFollowedSitesRequestBuilder) {
    return NewUsersItemFollowedSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowedSitesById provides operations to manage the followedSites property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) FollowedSitesById(id string)(*UsersItemFollowedSitesSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return NewUsersItemFollowedSitesSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get retrieve the properties and relationships of user object.
func (m *UsersUserItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersUserItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable), nil
}
// GetMailTips provides operations to call the getMailTips method.
func (m *UsersUserItemRequestBuilder) GetMailTips()(*UsersItemGetMailTipsRequestBuilder) {
    return NewUsersItemGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UsersUserItemRequestBuilder) GetManagedAppDiagnosticStatuses()(*UsersItemGetManagedAppDiagnosticStatusesRequestBuilder) {
    return NewUsersItemGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UsersUserItemRequestBuilder) GetManagedAppPolicies()(*UsersItemGetManagedAppPoliciesRequestBuilder) {
    return NewUsersItemGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *UsersUserItemRequestBuilder) GetMemberGroups()(*UsersItemGetMemberGroupsRequestBuilder) {
    return NewUsersItemGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *UsersUserItemRequestBuilder) GetMemberObjects()(*UsersItemGetMemberObjectsRequestBuilder) {
    return NewUsersItemGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InferenceClassification provides operations to manage the inferenceClassification property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) InferenceClassification()(*UsersItemInferenceClassificationRequestBuilder) {
    return NewUsersItemInferenceClassificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Insights provides operations to manage the insights property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Insights()(*UsersItemInsightsRequestBuilder) {
    return NewUsersItemInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedTeams provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) JoinedTeams()(*UsersItemJoinedTeamsRequestBuilder) {
    return NewUsersItemJoinedTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedTeamsById provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) JoinedTeamsById(id string)(*UsersItemJoinedTeamsTeamItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["team%2Did"] = id
    }
    return NewUsersItemJoinedTeamsTeamItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LicenseDetails provides operations to manage the licenseDetails property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) LicenseDetails()(*UsersItemLicenseDetailsRequestBuilder) {
    return NewUsersItemLicenseDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LicenseDetailsById provides operations to manage the licenseDetails property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) LicenseDetailsById(id string)(*UsersItemLicenseDetailsLicenseDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["licenseDetails%2Did"] = id
    }
    return NewUsersItemLicenseDetailsLicenseDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MailFolders provides operations to manage the mailFolders property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) MailFolders()(*UsersItemMailFoldersRequestBuilder) {
    return NewUsersItemMailFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MailFoldersById provides operations to manage the mailFolders property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) MailFoldersById(id string)(*UsersItemMailFoldersMailFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder%2Did"] = id
    }
    return NewUsersItemMailFoldersMailFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedAppRegistrations provides operations to manage the managedAppRegistrations property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ManagedAppRegistrations()(*UsersItemManagedAppRegistrationsRequestBuilder) {
    return NewUsersItemManagedAppRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedAppRegistrationsById provides operations to manage the managedAppRegistrations property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ManagedAppRegistrationsById(id string)(*UsersItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppRegistration%2Did"] = id
    }
    return NewUsersItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedDevices provides operations to manage the managedDevices property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ManagedDevices()(*UsersItemManagedDevicesRequestBuilder) {
    return NewUsersItemManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDevicesById provides operations to manage the managedDevices property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ManagedDevicesById(id string)(*UsersItemManagedDevicesManagedDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDevice%2Did"] = id
    }
    return NewUsersItemManagedDevicesManagedDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Manager provides operations to manage the manager property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Manager()(*UsersItemManagerRequestBuilder) {
    return NewUsersItemManagerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) MemberOf()(*UsersItemMemberOfRequestBuilder) {
    return NewUsersItemMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById provides operations to manage the memberOf property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) MemberOfById(id string)(*UsersItemMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages provides operations to manage the messages property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Messages()(*UsersItemMessagesRequestBuilder) {
    return NewUsersItemMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) MessagesById(id string)(*UsersItemMessagesMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message%2Did"] = id
    }
    return NewUsersItemMessagesMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Oauth2PermissionGrants provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Oauth2PermissionGrants()(*UsersItemOauth2PermissionGrantsRequestBuilder) {
    return NewUsersItemOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Oauth2PermissionGrantsById provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Oauth2PermissionGrantsById(id string)(*UsersItemOauth2PermissionGrantsOAuth2PermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["oAuth2PermissionGrant%2Did"] = id
    }
    return NewUsersItemOauth2PermissionGrantsOAuth2PermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Onenote()(*UsersItemOnenoteRequestBuilder) {
    return NewUsersItemOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetings provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) OnlineMeetings()(*UsersItemOnlineMeetingsRequestBuilder) {
    return NewUsersItemOnlineMeetingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetingsById provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) OnlineMeetingsById(id string)(*UsersItemOnlineMeetingsOnlineMeetingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onlineMeeting%2Did"] = id
    }
    return NewUsersItemOnlineMeetingsOnlineMeetingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Outlook provides operations to manage the outlook property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Outlook()(*UsersItemOutlookRequestBuilder) {
    return NewUsersItemOutlookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedDevices provides operations to manage the ownedDevices property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) OwnedDevices()(*UsersItemOwnedDevicesRequestBuilder) {
    return NewUsersItemOwnedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedDevicesById provides operations to manage the ownedDevices property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) OwnedDevicesById(id string)(*UsersItemOwnedDevicesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemOwnedDevicesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OwnedObjects provides operations to manage the ownedObjects property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) OwnedObjects()(*UsersItemOwnedObjectsRequestBuilder) {
    return NewUsersItemOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedObjectsById provides operations to manage the ownedObjects property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) OwnedObjectsById(id string)(*UsersItemOwnedObjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemOwnedObjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
func (m *UsersUserItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, requestConfiguration *UsersUserItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable), nil
}
// People provides operations to manage the people property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) People()(*UsersItemPeopleRequestBuilder) {
    return NewUsersItemPeopleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PeopleById provides operations to manage the people property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) PeopleById(id string)(*UsersItemPeoplePersonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["person%2Did"] = id
    }
    return NewUsersItemPeoplePersonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo provides operations to manage the photo property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Photo()(*UsersItemPhotoRequestBuilder) {
    return NewUsersItemPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Photos provides operations to manage the photos property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Photos()(*UsersItemPhotosRequestBuilder) {
    return NewUsersItemPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhotosById provides operations to manage the photos property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) PhotosById(id string)(*UsersItemPhotosProfilePhotoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profilePhoto%2Did"] = id
    }
    return NewUsersItemPhotosProfilePhotoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Planner provides operations to manage the planner property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Planner()(*UsersItemPlannerRequestBuilder) {
    return NewUsersItemPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Presence provides operations to manage the presence property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Presence()(*UsersItemPresenceRequestBuilder) {
    return NewUsersItemPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredDevices provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) RegisteredDevices()(*UsersItemRegisteredDevicesRequestBuilder) {
    return NewUsersItemRegisteredDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredDevicesById provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) RegisteredDevicesById(id string)(*UsersItemRegisteredDevicesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemRegisteredDevicesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UsersUserItemRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*UsersItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewUsersItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement provides operations to call the removeAllDevicesFromManagement method.
func (m *UsersUserItemRequestBuilder) RemoveAllDevicesFromManagement()(*UsersItemRemoveAllDevicesFromManagementRequestBuilder) {
    return NewUsersItemRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment provides operations to call the reprocessLicenseAssignment method.
func (m *UsersUserItemRequestBuilder) ReprocessLicenseAssignment()(*UsersItemReprocessLicenseAssignmentRequestBuilder) {
    return NewUsersItemReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *UsersUserItemRequestBuilder) Restore()(*UsersItemRestoreRequestBuilder) {
    return NewUsersItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions provides operations to call the revokeSignInSessions method.
func (m *UsersUserItemRequestBuilder) RevokeSignInSessions()(*UsersItemRevokeSignInSessionsRequestBuilder) {
    return NewUsersItemRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMemberOf provides operations to manage the scopedRoleMemberOf property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ScopedRoleMemberOf()(*UsersItemScopedRoleMemberOfRequestBuilder) {
    return NewUsersItemScopedRoleMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMemberOfById provides operations to manage the scopedRoleMemberOf property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ScopedRoleMemberOfById(id string)(*UsersItemScopedRoleMemberOfScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership%2Did"] = id
    }
    return NewUsersItemScopedRoleMemberOfScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SendMail provides operations to call the sendMail method.
func (m *UsersUserItemRequestBuilder) SendMail()(*UsersItemSendMailRequestBuilder) {
    return NewUsersItemSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings provides operations to manage the settings property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Settings()(*UsersItemSettingsRequestBuilder) {
    return NewUsersItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Teamwork provides operations to manage the teamwork property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Teamwork()(*UsersItemTeamworkRequestBuilder) {
    return NewUsersItemTeamworkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Todo provides operations to manage the todo property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Todo()(*UsersItemTodoRequestBuilder) {
    return NewUsersItemTodoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) TransitiveMemberOf()(*UsersItemTransitiveMemberOfRequestBuilder) {
    return NewUsersItemTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) TransitiveMemberOfById(id string)(*UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TranslateExchangeIds provides operations to call the translateExchangeIds method.
func (m *UsersUserItemRequestBuilder) TranslateExchangeIds()(*UsersItemTranslateExchangeIdsRequestBuilder) {
    return NewUsersItemTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag provides operations to call the wipeManagedAppRegistrationsByDeviceTag method.
func (m *UsersUserItemRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*UsersItemWipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return NewUsersItemWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
