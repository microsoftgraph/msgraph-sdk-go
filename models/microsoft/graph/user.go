package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// User 
type User struct {
    DirectoryObject
    // A freeform text entry field for the user to describe themselves. Returned only on $select.
    aboutMe *string;
    // true if the account is enabled; otherwise, false. This property is required when a user is created. Returned only on $select. Supports $filter (eq, ne, not, and in).
    accountEnabled *bool;
    // The user's activities across devices. Read-only. Nullable.
    activities []UserActivity;
    // Sets the age group of the user. Allowed values: null, minor, notAdult and adult. Refer to the legal age group property definitions for further information. Returned only on $select. Supports $filter (eq, ne, not, and in).
    ageGroup *string;
    // The user's terms of use acceptance statuses. Read-only. Nullable.
    agreementAcceptances []AgreementAcceptance;
    // Represents the app roles a user has been granted for an application. Supports $expand.
    appRoleAssignments []AppRoleAssignment;
    // The licenses that are assigned to the user, including inherited (group-based) licenses.  Not nullable. Returned only on $select. Supports $filter (eq and not).
    assignedLicenses []AssignedLicense;
    // The plans that are assigned to the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq and not).
    assignedPlans []AssignedPlan;
    // 
    authentication *Authentication;
    // The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.
    birthday *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The telephone numbers for the user. NOTE: Although this is a string collection, only one number can be set for this property. Read-only for users synced from on-premises directory. Returned by default. Supports $filter (eq, not, ge, le, startsWith).
    businessPhones []string;
    // The user's primary calendar. Read-only.
    calendar *Calendar;
    // The user's calendar groups. Read-only. Nullable.
    calendarGroups []CalendarGroup;
    // The user's calendars. Read-only. Nullable.
    calendars []Calendar;
    // The calendar view for the calendar. Read-only. Nullable.
    calendarView []Event;
    // 
    chats []Chat;
    // The city in which the user is located. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    city *string;
    // The company name which the user is associated. This property can be useful for describing the company that an external user comes from. The maximum length of the company name is 64 characters.Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    companyName *string;
    // Sets whether consent has been obtained for minors. Allowed values: null, granted, denied and notRequired. Refer to the legal age group property definitions for further information. Returned only on $select. Supports $filter (eq, ne, not, and in).
    consentProvidedForMinor *string;
    // The user's contacts folders. Read-only. Nullable.
    contactFolders []ContactFolder;
    // The user's contacts. Read-only. Nullable.
    contacts []Contact;
    // The country/region in which the user is located; for example, US or UK. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    country *string;
    // The created date of the user object. Read-only. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Directory objects that were created by the user. Read-only. Nullable.
    createdObjects []DirectoryObject;
    // Indicates whether the user account was created through one of the following methods:  As a regular school or work account (null). As an external account (Invitation). As a local account for an Azure Active Directory B2C tenant (LocalAccount). Through self-service sign-up by an internal user using email verification (EmailVerified). Through self-service sign-up by an external user signing up through a link that is part of a user flow (SelfServiceSignUp). Read-only.Returned only on $select. Supports $filter (eq, ne, not, in).
    creationType *string;
    // The name for the department in which the user works. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, and eq on null values).
    department *string;
    // The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000.
    deviceEnrollmentLimit *int32;
    // The list of troubleshooting events for this user.
    deviceManagementTroubleshootingEvents []DeviceManagementTroubleshootingEvent;
    // The users and contacts that report to the user. (The users and contacts that have their manager property set to this user.) Read-only. Nullable. Supports $expand.
    directReports []DirectoryObject;
    // The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial and last name. This property is required when a user is created and it cannot be cleared during updates. Maximum length is 256 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values), $orderBy, and $search.
    displayName *string;
    // The user's OneDrive. Read-only.
    drive *Drive;
    // A collection of drives available for this user. Read-only.
    drives []Drive;
    // The date and time when the user was hired or will start work in case of a future hire. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
    employeeHireDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The employee identifier assigned to the user by the organization. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
    employeeId *string;
    // Represents organization data (e.g. division and costCenter) associated with a user. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
    employeeOrgData *EmployeeOrgData;
    // Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith).
    employeeType *string;
    // The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
    events []Event;
    // The collection of open extensions defined for the user. Read-only. Nullable.
    extensions []Extension;
    // For an external user invited to the tenant using the invitation API, this property represents the invited user's invitation status. For invited users, the state can be PendingAcceptance or Accepted, or null for all other users. Returned only on $select. Supports $filter (eq, ne, not , in).
    externalUserState *string;
    // Shows the timestamp for the latest change to the externalUserState property. Returned only on $select. Supports $filter (eq, ne, not , in).
    externalUserStateChangeDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The fax number of the user. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
    faxNumber *string;
    // 
    followedSites []Site;
    // The given name (first name) of the user. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
    givenName *string;
    // The hire date of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.  Note: This property is specific to SharePoint Online. We recommend using the native employeeHireDate property to set and update hire date values using Microsoft Graph APIs.
    hireDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Represents the identities that can be used to sign in to this user account. An identity can be provided by Microsoft (also known as a local account), by organizations, or by social identity providers such as Facebook, Google, and Microsoft, and tied to a user account. May contain multiple items with the same signInType value. Returned only on $select. Supports $filter (eq) including on null values, only where the signInType is not userPrincipalName.
    identities []ObjectIdentity;
    // The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user. Read-only. Returned only on $select. Supports $filter (eq, not, ge, le, startsWith).
    imAddresses []string;
    // Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
    inferenceClassification *InferenceClassification;
    // Read-only. Nullable.
    insights *OfficeGraphInsights;
    // A list for the user to describe their interests. Returned only on $select.
    interests []string;
    // Do not use – reserved for future use.
    isResourceAccount *bool;
    // The user's job title. Maximum length is 128 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
    jobTitle *string;
    // The Microsoft Teams teams that the user is a member of. Read-only. Nullable.
    joinedTeams []Team;
    // The time when this Azure AD user last changed their password or when their password was created, whichever date the latest action was performed. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.
    lastPasswordChangeDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Used by enterprise applications to determine the legal age group of the user. This property is read-only and calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, minorWithOutParentalConsent, minorWithParentalConsent, minorNoParentalConsentRequired, notAdult and adult. Refer to the legal age group property definitions for further information. Returned only on $select.
    legalAgeGroupClassification *string;
    // State of license assignments for this user. Read-only. Returned only on $select.
    licenseAssignmentStates []LicenseAssignmentState;
    // A collection of this user's license details. Read-only.
    licenseDetails []LicenseDetails;
    // The SMTP address for the user, for example, jeff@contoso.onmicrosoft.com.Changes to this property will also update the user's proxyAddresses collection to include the value as an SMTP address. For Azure AD B2C accounts, this property can be updated up to only ten times with unique SMTP addresses. This property cannot contain accent characters.Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith, and eq on null values).
    mail *string;
    // Settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale and time zone.Returned only on $select.
    mailboxSettings *MailboxSettings;
    // The user's mail folders. Read-only. Nullable.
    mailFolders []MailFolder;
    // The mail alias for the user. This property must be specified when a user is created. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    mailNickname *string;
    // Zero or more managed app registrations that belong to the user.
    managedAppRegistrations []ManagedAppRegistration;
    // The managed devices associated with the user.
    managedDevices []ManagedDevice;
    // The user or contact that is this user's manager. Read-only. (HTTP Methods: GET, PUT, DELETE.). Supports $expand.
    manager *DirectoryObject;
    // The groups and directory roles that the user is a member of. Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObject;
    // The messages in a mailbox or folder. Read-only. Nullable.
    messages []Message;
    // The primary cellular telephone number for the user. Read-only for users synced from on-premises directory. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    mobilePhone *string;
    // The URL for the user's personal site. Returned only on $select.
    mySite *string;
    // 
    oauth2PermissionGrants []OAuth2PermissionGrant;
    // The office location in the user's place of business. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    officeLocation *string;
    // Read-only.
    onenote *Onenote;
    // 
    onlineMeetings []OnlineMeeting;
    // Contains the on-premises Active Directory distinguished name or DN. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select.
    onPremisesDistinguishedName *string;
    // Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select.
    onPremisesDomainName *string;
    // Contains extensionAttributes1-15 for the user. The individual extension attributes are neither selectable nor filterable. For an onPremisesSyncEnabled user, the source of authority for this set of properties is the on-premises and is read-only. For a cloud-only user (where onPremisesSyncEnabled is false), these properties can be set during creation or update of a user object.  For a cloud-only user previously synced from on-premises Active Directory, these properties are read-only in Microsoft Graph but can be fully managed through the Exchange Admin Center or the Exchange Online V2 module in PowerShell. These extension attributes are also known as Exchange custom attributes 1-15.
    onPremisesExtensionAttributes *OnPremisesExtensionAttributes;
    // This property is used to associate an on-premises Active Directory user account to their Azure AD user object. This property must be specified when creating a new user account in the Graph if you are using a federated domain for the user's userPrincipalName (UPN) property. NOTE: The $ and _ characters cannot be used when specifying this property. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in)..
    onPremisesImmutableId *string;
    // Indicates the last time at which the object was synced with the on-premises directory; for example: 2013-02-16T03:04:54Z. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in).
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Errors when using Microsoft synchronization product during provisioning. Returned only on $select. Supports $filter (eq, not, ge, le).
    onPremisesProvisioningErrors []OnPremisesProvisioningError;
    // Contains the on-premises samAccountName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith).
    onPremisesSamAccountName *string;
    // Contains the on-premises security identifier (SID) for the user that was synchronized from on-premises to the cloud. Read-only. Returned only on $select. Supports $filter (eq) on null values only.
    onPremisesSecurityIdentifier *string;
    // true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values).
    onPremisesSyncEnabled *bool;
    // Contains the on-premises userPrincipalName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith).
    onPremisesUserPrincipalName *string;
    // A list of additional email addresses for the user; for example: ['bob@contoso.com', 'Robert@fabrikam.com']. NOTE: This property cannot contain accent characters. Returned only on $select. Supports $filter (eq, not, ge, le, in, startsWith).
    otherMails []string;
    // Read-only.
    outlook *OutlookUser;
    // Devices that are owned by the user. Read-only. Nullable. Supports $expand.
    ownedDevices []DirectoryObject;
    // Directory objects that are owned by the user. Read-only. Nullable. Supports $expand.
    ownedObjects []DirectoryObject;
    // Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two may be specified together; for example: DisablePasswordExpiration, DisableStrongPassword. Returned only on $select. For more information on the default password policies, see Azure AD pasword policies. Supports $filter (ne, not, and eq on null values).
    passwordPolicies *string;
    // Specifies the password profile for the user. The profile contains the user’s password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required. NOTE: For Azure B2C tenants, the forceChangePasswordNextSignIn property should be set to false and instead use custom policies and user flows to force password reset at first logon. See Force password reset at first logon.Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values).
    passwordProfile *PasswordProfile;
    // A list for the user to enumerate their past projects. Returned only on $select.
    pastProjects []string;
    // People that are relevant to the user. Read-only. Nullable.
    people []Person;
    // The user's profile photo. Read-only.
    photo *ProfilePhoto;
    // Read-only. Nullable.
    photos []ProfilePhoto;
    // Entry-point to the Planner resource that might exist for a user. Read-only.
    planner *PlannerUser;
    // The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code. Maximum length is 40 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    postalCode *string;
    // The preferred language for the user. Should follow ISO 639-1 Code; for example en-US. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values)
    preferredLanguage *string;
    // The preferred name for the user. Returned only on $select.
    preferredName *string;
    // 
    presence *Presence;
    // The plans that are provisioned for the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq, not, ge, le).
    provisionedPlans []ProvisionedPlan;
    // For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. For Azure AD B2C accounts, this property has a limit of ten unique addresses. Read-only, Not nullable. Returned only on $select. Supports $filter (eq, not, ge, le, startsWith).
    proxyAddresses []string;
    // Devices that are registered for the user. Read-only. Nullable. Supports $expand.
    registeredDevices []DirectoryObject;
    // A list for the user to enumerate their responsibilities. Returned only on $select.
    responsibilities []string;
    // A list for the user to enumerate the schools they have attended. Returned only on $select.
    schools []string;
    // The scoped-role administrative unit memberships for this user. Read-only. Nullable.
    scopedRoleMemberOf []ScopedRoleMembership;
    // Read-only. Nullable.
    settings *UserSettings;
    // true if the Outlook global address list should contain this user, otherwise false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false. Returned only on $select. Supports $filter (eq, ne, not, in).
    showInAddressList *bool;
    // Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).  If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint. Read-only. Use revokeSignInSessions to reset. Returned only on $select.
    signInSessionsValidFromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // A list for the user to enumerate their skills. Returned only on $select.
    skills []string;
    // The state or province in the user's address. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    state *string;
    // The street address of the user's place of business. Maximum length is 1024 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    streetAddress *string;
    // The user's surname (family name or last name). Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    surname *string;
    // A container for Microsoft Teams features available for the user. Read-only. Nullable.
    teamwork *UserTeamwork;
    // Represents the To Do services available to a user.
    todo *Todo;
    // 
    transitiveMemberOf []DirectoryObject;
    // A two letter country code (ISO standard 3166). Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.  Examples include: US, JP, and GB. Not nullable. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
    usageLocation *string;
    // The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization.NOTE: This property cannot contain accent characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith) and $orderBy.
    userPrincipalName *string;
    // A string value that can be used to classify user types in your directory, such as Member and Guest. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values). NOTE: For more information about the permissions for member and guest users, see What are the default user permissions in Azure Active Directory?
    userType *string;
}
// NewUser instantiates a new user and sets the default values.
func NewUser()(*User) {
    m := &User{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// GetAboutMe gets the aboutMe property value. A freeform text entry field for the user to describe themselves. Returned only on $select.
func (m *User) GetAboutMe()(*string) {
    if m == nil {
        return nil
    } else {
        return m.aboutMe
    }
}
// GetAccountEnabled gets the accountEnabled property value. true if the account is enabled; otherwise, false. This property is required when a user is created. Returned only on $select. Supports $filter (eq, ne, not, and in).
func (m *User) GetAccountEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accountEnabled
    }
}
// GetActivities gets the activities property value. The user's activities across devices. Read-only. Nullable.
func (m *User) GetActivities()([]UserActivity) {
    if m == nil {
        return nil
    } else {
        return m.activities
    }
}
// GetAgeGroup gets the ageGroup property value. Sets the age group of the user. Allowed values: null, minor, notAdult and adult. Refer to the legal age group property definitions for further information. Returned only on $select. Supports $filter (eq, ne, not, and in).
func (m *User) GetAgeGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ageGroup
    }
}
// GetAgreementAcceptances gets the agreementAcceptances property value. The user's terms of use acceptance statuses. Read-only. Nullable.
func (m *User) GetAgreementAcceptances()([]AgreementAcceptance) {
    if m == nil {
        return nil
    } else {
        return m.agreementAcceptances
    }
}
// GetAppRoleAssignments gets the appRoleAssignments property value. Represents the app roles a user has been granted for an application. Supports $expand.
func (m *User) GetAppRoleAssignments()([]AppRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.appRoleAssignments
    }
}
// GetAssignedLicenses gets the assignedLicenses property value. The licenses that are assigned to the user, including inherited (group-based) licenses.  Not nullable. Returned only on $select. Supports $filter (eq and not).
func (m *User) GetAssignedLicenses()([]AssignedLicense) {
    if m == nil {
        return nil
    } else {
        return m.assignedLicenses
    }
}
// GetAssignedPlans gets the assignedPlans property value. The plans that are assigned to the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq and not).
func (m *User) GetAssignedPlans()([]AssignedPlan) {
    if m == nil {
        return nil
    } else {
        return m.assignedPlans
    }
}
// GetAuthentication gets the authentication property value. 
func (m *User) GetAuthentication()(*Authentication) {
    if m == nil {
        return nil
    } else {
        return m.authentication
    }
}
// GetBirthday gets the birthday property value. The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.
func (m *User) GetBirthday()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.birthday
    }
}
// GetBusinessPhones gets the businessPhones property value. The telephone numbers for the user. NOTE: Although this is a string collection, only one number can be set for this property. Read-only for users synced from on-premises directory. Returned by default. Supports $filter (eq, not, ge, le, startsWith).
func (m *User) GetBusinessPhones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.businessPhones
    }
}
// GetCalendar gets the calendar property value. The user's primary calendar. Read-only.
func (m *User) GetCalendar()(*Calendar) {
    if m == nil {
        return nil
    } else {
        return m.calendar
    }
}
// GetCalendarGroups gets the calendarGroups property value. The user's calendar groups. Read-only. Nullable.
func (m *User) GetCalendarGroups()([]CalendarGroup) {
    if m == nil {
        return nil
    } else {
        return m.calendarGroups
    }
}
// GetCalendars gets the calendars property value. The user's calendars. Read-only. Nullable.
func (m *User) GetCalendars()([]Calendar) {
    if m == nil {
        return nil
    } else {
        return m.calendars
    }
}
// GetCalendarView gets the calendarView property value. The calendar view for the calendar. Read-only. Nullable.
func (m *User) GetCalendarView()([]Event) {
    if m == nil {
        return nil
    } else {
        return m.calendarView
    }
}
// GetChats gets the chats property value. 
func (m *User) GetChats()([]Chat) {
    if m == nil {
        return nil
    } else {
        return m.chats
    }
}
// GetCity gets the city property value. The city in which the user is located. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// GetCompanyName gets the companyName property value. The company name which the user is associated. This property can be useful for describing the company that an external user comes from. The maximum length of the company name is 64 characters.Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetCompanyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.companyName
    }
}
// GetConsentProvidedForMinor gets the consentProvidedForMinor property value. Sets whether consent has been obtained for minors. Allowed values: null, granted, denied and notRequired. Refer to the legal age group property definitions for further information. Returned only on $select. Supports $filter (eq, ne, not, and in).
func (m *User) GetConsentProvidedForMinor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.consentProvidedForMinor
    }
}
// GetContactFolders gets the contactFolders property value. The user's contacts folders. Read-only. Nullable.
func (m *User) GetContactFolders()([]ContactFolder) {
    if m == nil {
        return nil
    } else {
        return m.contactFolders
    }
}
// GetContacts gets the contacts property value. The user's contacts. Read-only. Nullable.
func (m *User) GetContacts()([]Contact) {
    if m == nil {
        return nil
    } else {
        return m.contacts
    }
}
// GetCountry gets the country property value. The country/region in which the user is located; for example, US or UK. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetCountry()(*string) {
    if m == nil {
        return nil
    } else {
        return m.country
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The created date of the user object. Read-only. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetCreatedObjects gets the createdObjects property value. Directory objects that were created by the user. Read-only. Nullable.
func (m *User) GetCreatedObjects()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.createdObjects
    }
}
// GetCreationType gets the creationType property value. Indicates whether the user account was created through one of the following methods:  As a regular school or work account (null). As an external account (Invitation). As a local account for an Azure Active Directory B2C tenant (LocalAccount). Through self-service sign-up by an internal user using email verification (EmailVerified). Through self-service sign-up by an external user signing up through a link that is part of a user flow (SelfServiceSignUp). Read-only.Returned only on $select. Supports $filter (eq, ne, not, in).
func (m *User) GetCreationType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.creationType
    }
}
// GetDepartment gets the department property value. The name for the department in which the user works. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, and eq on null values).
func (m *User) GetDepartment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.department
    }
}
// GetDeviceEnrollmentLimit gets the deviceEnrollmentLimit property value. The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000.
func (m *User) GetDeviceEnrollmentLimit()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceEnrollmentLimit
    }
}
// GetDeviceManagementTroubleshootingEvents gets the deviceManagementTroubleshootingEvents property value. The list of troubleshooting events for this user.
func (m *User) GetDeviceManagementTroubleshootingEvents()([]DeviceManagementTroubleshootingEvent) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementTroubleshootingEvents
    }
}
// GetDirectReports gets the directReports property value. The users and contacts that report to the user. (The users and contacts that have their manager property set to this user.) Read-only. Nullable. Supports $expand.
func (m *User) GetDirectReports()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.directReports
    }
}
// GetDisplayName gets the displayName property value. The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial and last name. This property is required when a user is created and it cannot be cleared during updates. Maximum length is 256 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values), $orderBy, and $search.
func (m *User) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDrive gets the drive property value. The user's OneDrive. Read-only.
func (m *User) GetDrive()(*Drive) {
    if m == nil {
        return nil
    } else {
        return m.drive
    }
}
// GetDrives gets the drives property value. A collection of drives available for this user. Read-only.
func (m *User) GetDrives()([]Drive) {
    if m == nil {
        return nil
    } else {
        return m.drives
    }
}
// GetEmployeeHireDate gets the employeeHireDate property value. The date and time when the user was hired or will start work in case of a future hire. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) GetEmployeeHireDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.employeeHireDate
    }
}
// GetEmployeeId gets the employeeId property value. The employee identifier assigned to the user by the organization. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) GetEmployeeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.employeeId
    }
}
// GetEmployeeOrgData gets the employeeOrgData property value. Represents organization data (e.g. division and costCenter) associated with a user. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) GetEmployeeOrgData()(*EmployeeOrgData) {
    if m == nil {
        return nil
    } else {
        return m.employeeOrgData
    }
}
// GetEmployeeType gets the employeeType property value. Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith).
func (m *User) GetEmployeeType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.employeeType
    }
}
// GetEvents gets the events property value. The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *User) GetEvents()([]Event) {
    if m == nil {
        return nil
    } else {
        return m.events
    }
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the user. Read-only. Nullable.
func (m *User) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetExternalUserState gets the externalUserState property value. For an external user invited to the tenant using the invitation API, this property represents the invited user's invitation status. For invited users, the state can be PendingAcceptance or Accepted, or null for all other users. Returned only on $select. Supports $filter (eq, ne, not , in).
func (m *User) GetExternalUserState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalUserState
    }
}
// GetExternalUserStateChangeDateTime gets the externalUserStateChangeDateTime property value. Shows the timestamp for the latest change to the externalUserState property. Returned only on $select. Supports $filter (eq, ne, not , in).
func (m *User) GetExternalUserStateChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.externalUserStateChangeDateTime
    }
}
// GetFaxNumber gets the faxNumber property value. The fax number of the user. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) GetFaxNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.faxNumber
    }
}
// GetFollowedSites gets the followedSites property value. 
func (m *User) GetFollowedSites()([]Site) {
    if m == nil {
        return nil
    } else {
        return m.followedSites
    }
}
// GetGivenName gets the givenName property value. The given name (first name) of the user. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) GetGivenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.givenName
    }
}
// GetHireDate gets the hireDate property value. The hire date of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.  Note: This property is specific to SharePoint Online. We recommend using the native employeeHireDate property to set and update hire date values using Microsoft Graph APIs.
func (m *User) GetHireDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.hireDate
    }
}
// GetIdentities gets the identities property value. Represents the identities that can be used to sign in to this user account. An identity can be provided by Microsoft (also known as a local account), by organizations, or by social identity providers such as Facebook, Google, and Microsoft, and tied to a user account. May contain multiple items with the same signInType value. Returned only on $select. Supports $filter (eq) including on null values, only where the signInType is not userPrincipalName.
func (m *User) GetIdentities()([]ObjectIdentity) {
    if m == nil {
        return nil
    } else {
        return m.identities
    }
}
// GetImAddresses gets the imAddresses property value. The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user. Read-only. Returned only on $select. Supports $filter (eq, not, ge, le, startsWith).
func (m *User) GetImAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.imAddresses
    }
}
// GetInferenceClassification gets the inferenceClassification property value. Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *User) GetInferenceClassification()(*InferenceClassification) {
    if m == nil {
        return nil
    } else {
        return m.inferenceClassification
    }
}
// GetInsights gets the insights property value. Read-only. Nullable.
func (m *User) GetInsights()(*OfficeGraphInsights) {
    if m == nil {
        return nil
    } else {
        return m.insights
    }
}
// GetInterests gets the interests property value. A list for the user to describe their interests. Returned only on $select.
func (m *User) GetInterests()([]string) {
    if m == nil {
        return nil
    } else {
        return m.interests
    }
}
// GetIsResourceAccount gets the isResourceAccount property value. Do not use – reserved for future use.
func (m *User) GetIsResourceAccount()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isResourceAccount
    }
}
// GetJobTitle gets the jobTitle property value. The user's job title. Maximum length is 128 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) GetJobTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jobTitle
    }
}
// GetJoinedTeams gets the joinedTeams property value. The Microsoft Teams teams that the user is a member of. Read-only. Nullable.
func (m *User) GetJoinedTeams()([]Team) {
    if m == nil {
        return nil
    } else {
        return m.joinedTeams
    }
}
// GetLastPasswordChangeDateTime gets the lastPasswordChangeDateTime property value. The time when this Azure AD user last changed their password or when their password was created, whichever date the latest action was performed. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.
func (m *User) GetLastPasswordChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastPasswordChangeDateTime
    }
}
// GetLegalAgeGroupClassification gets the legalAgeGroupClassification property value. Used by enterprise applications to determine the legal age group of the user. This property is read-only and calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, minorWithOutParentalConsent, minorWithParentalConsent, minorNoParentalConsentRequired, notAdult and adult. Refer to the legal age group property definitions for further information. Returned only on $select.
func (m *User) GetLegalAgeGroupClassification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.legalAgeGroupClassification
    }
}
// GetLicenseAssignmentStates gets the licenseAssignmentStates property value. State of license assignments for this user. Read-only. Returned only on $select.
func (m *User) GetLicenseAssignmentStates()([]LicenseAssignmentState) {
    if m == nil {
        return nil
    } else {
        return m.licenseAssignmentStates
    }
}
// GetLicenseDetails gets the licenseDetails property value. A collection of this user's license details. Read-only.
func (m *User) GetLicenseDetails()([]LicenseDetails) {
    if m == nil {
        return nil
    } else {
        return m.licenseDetails
    }
}
// GetMail gets the mail property value. The SMTP address for the user, for example, jeff@contoso.onmicrosoft.com.Changes to this property will also update the user's proxyAddresses collection to include the value as an SMTP address. For Azure AD B2C accounts, this property can be updated up to only ten times with unique SMTP addresses. This property cannot contain accent characters.Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith, and eq on null values).
func (m *User) GetMail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mail
    }
}
// GetMailboxSettings gets the mailboxSettings property value. Settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale and time zone.Returned only on $select.
func (m *User) GetMailboxSettings()(*MailboxSettings) {
    if m == nil {
        return nil
    } else {
        return m.mailboxSettings
    }
}
// GetMailFolders gets the mailFolders property value. The user's mail folders. Read-only. Nullable.
func (m *User) GetMailFolders()([]MailFolder) {
    if m == nil {
        return nil
    } else {
        return m.mailFolders
    }
}
// GetMailNickname gets the mailNickname property value. The mail alias for the user. This property must be specified when a user is created. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetMailNickname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mailNickname
    }
}
// GetManagedAppRegistrations gets the managedAppRegistrations property value. Zero or more managed app registrations that belong to the user.
func (m *User) GetManagedAppRegistrations()([]ManagedAppRegistration) {
    if m == nil {
        return nil
    } else {
        return m.managedAppRegistrations
    }
}
// GetManagedDevices gets the managedDevices property value. The managed devices associated with the user.
func (m *User) GetManagedDevices()([]ManagedDevice) {
    if m == nil {
        return nil
    } else {
        return m.managedDevices
    }
}
// GetManager gets the manager property value. The user or contact that is this user's manager. Read-only. (HTTP Methods: GET, PUT, DELETE.). Supports $expand.
func (m *User) GetManager()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.manager
    }
}
// GetMemberOf gets the memberOf property value. The groups and directory roles that the user is a member of. Read-only. Nullable. Supports $expand.
func (m *User) GetMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.memberOf
    }
}
// GetMessages gets the messages property value. The messages in a mailbox or folder. Read-only. Nullable.
func (m *User) GetMessages()([]Message) {
    if m == nil {
        return nil
    } else {
        return m.messages
    }
}
// GetMobilePhone gets the mobilePhone property value. The primary cellular telephone number for the user. Read-only for users synced from on-premises directory. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetMobilePhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mobilePhone
    }
}
// GetMySite gets the mySite property value. The URL for the user's personal site. Returned only on $select.
func (m *User) GetMySite()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mySite
    }
}
// GetOauth2PermissionGrants gets the oauth2PermissionGrants property value. 
func (m *User) GetOauth2PermissionGrants()([]OAuth2PermissionGrant) {
    if m == nil {
        return nil
    } else {
        return m.oauth2PermissionGrants
    }
}
// GetOfficeLocation gets the officeLocation property value. The office location in the user's place of business. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetOfficeLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.officeLocation
    }
}
// GetOnenote gets the onenote property value. Read-only.
func (m *User) GetOnenote()(*Onenote) {
    if m == nil {
        return nil
    } else {
        return m.onenote
    }
}
// GetOnlineMeetings gets the onlineMeetings property value. 
func (m *User) GetOnlineMeetings()([]OnlineMeeting) {
    if m == nil {
        return nil
    } else {
        return m.onlineMeetings
    }
}
// GetOnPremisesDistinguishedName gets the onPremisesDistinguishedName property value. Contains the on-premises Active Directory distinguished name or DN. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select.
func (m *User) GetOnPremisesDistinguishedName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesDistinguishedName
    }
}
// GetOnPremisesDomainName gets the onPremisesDomainName property value. Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select.
func (m *User) GetOnPremisesDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesDomainName
    }
}
// GetOnPremisesExtensionAttributes gets the onPremisesExtensionAttributes property value. Contains extensionAttributes1-15 for the user. The individual extension attributes are neither selectable nor filterable. For an onPremisesSyncEnabled user, the source of authority for this set of properties is the on-premises and is read-only. For a cloud-only user (where onPremisesSyncEnabled is false), these properties can be set during creation or update of a user object.  For a cloud-only user previously synced from on-premises Active Directory, these properties are read-only in Microsoft Graph but can be fully managed through the Exchange Admin Center or the Exchange Online V2 module in PowerShell. These extension attributes are also known as Exchange custom attributes 1-15.
func (m *User) GetOnPremisesExtensionAttributes()(*OnPremisesExtensionAttributes) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesExtensionAttributes
    }
}
// GetOnPremisesImmutableId gets the onPremisesImmutableId property value. This property is used to associate an on-premises Active Directory user account to their Azure AD user object. This property must be specified when creating a new user account in the Graph if you are using a federated domain for the user's userPrincipalName (UPN) property. NOTE: The $ and _ characters cannot be used when specifying this property. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in)..
func (m *User) GetOnPremisesImmutableId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesImmutableId
    }
}
// GetOnPremisesLastSyncDateTime gets the onPremisesLastSyncDateTime property value. Indicates the last time at which the object was synced with the on-premises directory; for example: 2013-02-16T03:04:54Z. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in).
func (m *User) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesLastSyncDateTime
    }
}
// GetOnPremisesProvisioningErrors gets the onPremisesProvisioningErrors property value. Errors when using Microsoft synchronization product during provisioning. Returned only on $select. Supports $filter (eq, not, ge, le).
func (m *User) GetOnPremisesProvisioningErrors()([]OnPremisesProvisioningError) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesProvisioningErrors
    }
}
// GetOnPremisesSamAccountName gets the onPremisesSamAccountName property value. Contains the on-premises samAccountName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith).
func (m *User) GetOnPremisesSamAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSamAccountName
    }
}
// GetOnPremisesSecurityIdentifier gets the onPremisesSecurityIdentifier property value. Contains the on-premises security identifier (SID) for the user that was synchronized from on-premises to the cloud. Read-only. Returned only on $select. Supports $filter (eq) on null values only.
func (m *User) GetOnPremisesSecurityIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSecurityIdentifier
    }
}
// GetOnPremisesSyncEnabled gets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *User) GetOnPremisesSyncEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSyncEnabled
    }
}
// GetOnPremisesUserPrincipalName gets the onPremisesUserPrincipalName property value. Contains the on-premises userPrincipalName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith).
func (m *User) GetOnPremisesUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesUserPrincipalName
    }
}
// GetOtherMails gets the otherMails property value. A list of additional email addresses for the user; for example: ['bob@contoso.com', 'Robert@fabrikam.com']. NOTE: This property cannot contain accent characters. Returned only on $select. Supports $filter (eq, not, ge, le, in, startsWith).
func (m *User) GetOtherMails()([]string) {
    if m == nil {
        return nil
    } else {
        return m.otherMails
    }
}
// GetOutlook gets the outlook property value. Read-only.
func (m *User) GetOutlook()(*OutlookUser) {
    if m == nil {
        return nil
    } else {
        return m.outlook
    }
}
// GetOwnedDevices gets the ownedDevices property value. Devices that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *User) GetOwnedDevices()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.ownedDevices
    }
}
// GetOwnedObjects gets the ownedObjects property value. Directory objects that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *User) GetOwnedObjects()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.ownedObjects
    }
}
// GetPasswordPolicies gets the passwordPolicies property value. Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two may be specified together; for example: DisablePasswordExpiration, DisableStrongPassword. Returned only on $select. For more information on the default password policies, see Azure AD pasword policies. Supports $filter (ne, not, and eq on null values).
func (m *User) GetPasswordPolicies()(*string) {
    if m == nil {
        return nil
    } else {
        return m.passwordPolicies
    }
}
// GetPasswordProfile gets the passwordProfile property value. Specifies the password profile for the user. The profile contains the user’s password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required. NOTE: For Azure B2C tenants, the forceChangePasswordNextSignIn property should be set to false and instead use custom policies and user flows to force password reset at first logon. See Force password reset at first logon.Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *User) GetPasswordProfile()(*PasswordProfile) {
    if m == nil {
        return nil
    } else {
        return m.passwordProfile
    }
}
// GetPastProjects gets the pastProjects property value. A list for the user to enumerate their past projects. Returned only on $select.
func (m *User) GetPastProjects()([]string) {
    if m == nil {
        return nil
    } else {
        return m.pastProjects
    }
}
// GetPeople gets the people property value. People that are relevant to the user. Read-only. Nullable.
func (m *User) GetPeople()([]Person) {
    if m == nil {
        return nil
    } else {
        return m.people
    }
}
// GetPhoto gets the photo property value. The user's profile photo. Read-only.
func (m *User) GetPhoto()(*ProfilePhoto) {
    if m == nil {
        return nil
    } else {
        return m.photo
    }
}
// GetPhotos gets the photos property value. Read-only. Nullable.
func (m *User) GetPhotos()([]ProfilePhoto) {
    if m == nil {
        return nil
    } else {
        return m.photos
    }
}
// GetPlanner gets the planner property value. Entry-point to the Planner resource that might exist for a user. Read-only.
func (m *User) GetPlanner()(*PlannerUser) {
    if m == nil {
        return nil
    } else {
        return m.planner
    }
}
// GetPostalCode gets the postalCode property value. The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code. Maximum length is 40 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetPostalCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postalCode
    }
}
// GetPreferredLanguage gets the preferredLanguage property value. The preferred language for the user. Should follow ISO 639-1 Code; for example en-US. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values)
func (m *User) GetPreferredLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredLanguage
    }
}
// GetPreferredName gets the preferredName property value. The preferred name for the user. Returned only on $select.
func (m *User) GetPreferredName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredName
    }
}
// GetPresence gets the presence property value. 
func (m *User) GetPresence()(*Presence) {
    if m == nil {
        return nil
    } else {
        return m.presence
    }
}
// GetProvisionedPlans gets the provisionedPlans property value. The plans that are provisioned for the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq, not, ge, le).
func (m *User) GetProvisionedPlans()([]ProvisionedPlan) {
    if m == nil {
        return nil
    } else {
        return m.provisionedPlans
    }
}
// GetProxyAddresses gets the proxyAddresses property value. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. For Azure AD B2C accounts, this property has a limit of ten unique addresses. Read-only, Not nullable. Returned only on $select. Supports $filter (eq, not, ge, le, startsWith).
func (m *User) GetProxyAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.proxyAddresses
    }
}
// GetRegisteredDevices gets the registeredDevices property value. Devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *User) GetRegisteredDevices()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.registeredDevices
    }
}
// GetResponsibilities gets the responsibilities property value. A list for the user to enumerate their responsibilities. Returned only on $select.
func (m *User) GetResponsibilities()([]string) {
    if m == nil {
        return nil
    } else {
        return m.responsibilities
    }
}
// GetSchools gets the schools property value. A list for the user to enumerate the schools they have attended. Returned only on $select.
func (m *User) GetSchools()([]string) {
    if m == nil {
        return nil
    } else {
        return m.schools
    }
}
// GetScopedRoleMemberOf gets the scopedRoleMemberOf property value. The scoped-role administrative unit memberships for this user. Read-only. Nullable.
func (m *User) GetScopedRoleMemberOf()([]ScopedRoleMembership) {
    if m == nil {
        return nil
    } else {
        return m.scopedRoleMemberOf
    }
}
// GetSettings gets the settings property value. Read-only. Nullable.
func (m *User) GetSettings()(*UserSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetShowInAddressList gets the showInAddressList property value. true if the Outlook global address list should contain this user, otherwise false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false. Returned only on $select. Supports $filter (eq, ne, not, in).
func (m *User) GetShowInAddressList()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showInAddressList
    }
}
// GetSignInSessionsValidFromDateTime gets the signInSessionsValidFromDateTime property value. Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).  If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint. Read-only. Use revokeSignInSessions to reset. Returned only on $select.
func (m *User) GetSignInSessionsValidFromDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.signInSessionsValidFromDateTime
    }
}
// GetSkills gets the skills property value. A list for the user to enumerate their skills. Returned only on $select.
func (m *User) GetSkills()([]string) {
    if m == nil {
        return nil
    } else {
        return m.skills
    }
}
// GetState gets the state property value. The state or province in the user's address. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetStreetAddress gets the streetAddress property value. The street address of the user's place of business. Maximum length is 1024 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetStreetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.streetAddress
    }
}
// GetSurname gets the surname property value. The user's surname (family name or last name). Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetSurname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.surname
    }
}
// GetTeamwork gets the teamwork property value. A container for Microsoft Teams features available for the user. Read-only. Nullable.
func (m *User) GetTeamwork()(*UserTeamwork) {
    if m == nil {
        return nil
    } else {
        return m.teamwork
    }
}
// GetTodo gets the todo property value. Represents the To Do services available to a user.
func (m *User) GetTodo()(*Todo) {
    if m == nil {
        return nil
    } else {
        return m.todo
    }
}
// GetTransitiveMemberOf gets the transitiveMemberOf property value. 
func (m *User) GetTransitiveMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.transitiveMemberOf
    }
}
// GetUsageLocation gets the usageLocation property value. A two letter country code (ISO standard 3166). Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.  Examples include: US, JP, and GB. Not nullable. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) GetUsageLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.usageLocation
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization.NOTE: This property cannot contain accent characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith) and $orderBy.
func (m *User) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetUserType gets the userType property value. A string value that can be used to classify user types in your directory, such as Member and Guest. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values). NOTE: For more information about the permissions for member and guest users, see What are the default user permissions in Azure Active Directory?
func (m *User) GetUserType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *User) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["aboutMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAboutMe(val)
        }
        return nil
    }
    res["accountEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountEnabled(val)
        }
        return nil
    }
    res["activities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserActivity() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserActivity, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserActivity))
            }
            m.SetActivities(res)
        }
        return nil
    }
    res["ageGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAgeGroup(val)
        }
        return nil
    }
    res["agreementAcceptances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreementAcceptance() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgreementAcceptance, len(val))
            for i, v := range val {
                res[i] = *(v.(*AgreementAcceptance))
            }
            m.SetAgreementAcceptances(res)
        }
        return nil
    }
    res["appRoleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppRoleAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppRoleAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*AppRoleAssignment))
            }
            m.SetAppRoleAssignments(res)
        }
        return nil
    }
    res["assignedLicenses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignedLicense() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedLicense, len(val))
            for i, v := range val {
                res[i] = *(v.(*AssignedLicense))
            }
            m.SetAssignedLicenses(res)
        }
        return nil
    }
    res["assignedPlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignedPlan() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedPlan, len(val))
            for i, v := range val {
                res[i] = *(v.(*AssignedPlan))
            }
            m.SetAssignedPlans(res)
        }
        return nil
    }
    res["authentication"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthentication() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthentication(val.(*Authentication))
        }
        return nil
    }
    res["birthday"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBirthday(val)
        }
        return nil
    }
    res["businessPhones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetBusinessPhones(res)
        }
        return nil
    }
    res["calendar"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCalendar() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalendar(val.(*Calendar))
        }
        return nil
    }
    res["calendarGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCalendarGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CalendarGroup, len(val))
            for i, v := range val {
                res[i] = *(v.(*CalendarGroup))
            }
            m.SetCalendarGroups(res)
        }
        return nil
    }
    res["calendars"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCalendar() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Calendar, len(val))
            for i, v := range val {
                res[i] = *(v.(*Calendar))
            }
            m.SetCalendars(res)
        }
        return nil
    }
    res["calendarView"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Event, len(val))
            for i, v := range val {
                res[i] = *(v.(*Event))
            }
            m.SetCalendarView(res)
        }
        return nil
    }
    res["chats"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChat() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Chat, len(val))
            for i, v := range val {
                res[i] = *(v.(*Chat))
            }
            m.SetChats(res)
        }
        return nil
    }
    res["city"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["companyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyName(val)
        }
        return nil
    }
    res["consentProvidedForMinor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConsentProvidedForMinor(val)
        }
        return nil
    }
    res["contactFolders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContactFolder() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContactFolder, len(val))
            for i, v := range val {
                res[i] = *(v.(*ContactFolder))
            }
            m.SetContactFolders(res)
        }
        return nil
    }
    res["contacts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContact() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Contact, len(val))
            for i, v := range val {
                res[i] = *(v.(*Contact))
            }
            m.SetContacts(res)
        }
        return nil
    }
    res["country"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountry(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["createdObjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetCreatedObjects(res)
        }
        return nil
    }
    res["creationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationType(val)
        }
        return nil
    }
    res["department"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepartment(val)
        }
        return nil
    }
    res["deviceEnrollmentLimit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceEnrollmentLimit(val)
        }
        return nil
    }
    res["deviceManagementTroubleshootingEvents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementTroubleshootingEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTroubleshootingEvent, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementTroubleshootingEvent))
            }
            m.SetDeviceManagementTroubleshootingEvents(res)
        }
        return nil
    }
    res["directReports"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetDirectReports(res)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["drive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDrive() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDrive(val.(*Drive))
        }
        return nil
    }
    res["drives"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDrive() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Drive, len(val))
            for i, v := range val {
                res[i] = *(v.(*Drive))
            }
            m.SetDrives(res)
        }
        return nil
    }
    res["employeeHireDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmployeeHireDate(val)
        }
        return nil
    }
    res["employeeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmployeeId(val)
        }
        return nil
    }
    res["employeeOrgData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmployeeOrgData() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmployeeOrgData(val.(*EmployeeOrgData))
        }
        return nil
    }
    res["employeeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmployeeType(val)
        }
        return nil
    }
    res["events"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Event, len(val))
            for i, v := range val {
                res[i] = *(v.(*Event))
            }
            m.SetEvents(res)
        }
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtension() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extension, len(val))
            for i, v := range val {
                res[i] = *(v.(*Extension))
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["externalUserState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalUserState(val)
        }
        return nil
    }
    res["externalUserStateChangeDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalUserStateChangeDateTime(val)
        }
        return nil
    }
    res["faxNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFaxNumber(val)
        }
        return nil
    }
    res["followedSites"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSite() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Site, len(val))
            for i, v := range val {
                res[i] = *(v.(*Site))
            }
            m.SetFollowedSites(res)
        }
        return nil
    }
    res["givenName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
        }
        return nil
    }
    res["hireDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHireDate(val)
        }
        return nil
    }
    res["identities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewObjectIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ObjectIdentity, len(val))
            for i, v := range val {
                res[i] = *(v.(*ObjectIdentity))
            }
            m.SetIdentities(res)
        }
        return nil
    }
    res["imAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetImAddresses(res)
        }
        return nil
    }
    res["inferenceClassification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInferenceClassification() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInferenceClassification(val.(*InferenceClassification))
        }
        return nil
    }
    res["insights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfficeGraphInsights() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsights(val.(*OfficeGraphInsights))
        }
        return nil
    }
    res["interests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetInterests(res)
        }
        return nil
    }
    res["isResourceAccount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsResourceAccount(val)
        }
        return nil
    }
    res["jobTitle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobTitle(val)
        }
        return nil
    }
    res["joinedTeams"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeam() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Team, len(val))
            for i, v := range val {
                res[i] = *(v.(*Team))
            }
            m.SetJoinedTeams(res)
        }
        return nil
    }
    res["lastPasswordChangeDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastPasswordChangeDateTime(val)
        }
        return nil
    }
    res["legalAgeGroupClassification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLegalAgeGroupClassification(val)
        }
        return nil
    }
    res["licenseAssignmentStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLicenseAssignmentState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LicenseAssignmentState, len(val))
            for i, v := range val {
                res[i] = *(v.(*LicenseAssignmentState))
            }
            m.SetLicenseAssignmentStates(res)
        }
        return nil
    }
    res["licenseDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLicenseDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LicenseDetails, len(val))
            for i, v := range val {
                res[i] = *(v.(*LicenseDetails))
            }
            m.SetLicenseDetails(res)
        }
        return nil
    }
    res["mail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMail(val)
        }
        return nil
    }
    res["mailboxSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMailboxSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailboxSettings(val.(*MailboxSettings))
        }
        return nil
    }
    res["mailFolders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMailFolder() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MailFolder, len(val))
            for i, v := range val {
                res[i] = *(v.(*MailFolder))
            }
            m.SetMailFolders(res)
        }
        return nil
    }
    res["mailNickname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailNickname(val)
        }
        return nil
    }
    res["managedAppRegistrations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppRegistration() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppRegistration, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppRegistration))
            }
            m.SetManagedAppRegistrations(res)
        }
        return nil
    }
    res["managedDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDevice() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDevice, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedDevice))
            }
            m.SetManagedDevices(res)
        }
        return nil
    }
    res["manager"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManager(val.(*DirectoryObject))
        }
        return nil
    }
    res["memberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetMemberOf(res)
        }
        return nil
    }
    res["messages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMessage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Message, len(val))
            for i, v := range val {
                res[i] = *(v.(*Message))
            }
            m.SetMessages(res)
        }
        return nil
    }
    res["mobilePhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobilePhone(val)
        }
        return nil
    }
    res["mySite"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMySite(val)
        }
        return nil
    }
    res["oauth2PermissionGrants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOAuth2PermissionGrant() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OAuth2PermissionGrant, len(val))
            for i, v := range val {
                res[i] = *(v.(*OAuth2PermissionGrant))
            }
            m.SetOauth2PermissionGrants(res)
        }
        return nil
    }
    res["officeLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeLocation(val)
        }
        return nil
    }
    res["onenote"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnenote() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnenote(val.(*Onenote))
        }
        return nil
    }
    res["onlineMeetings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnlineMeeting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnlineMeeting, len(val))
            for i, v := range val {
                res[i] = *(v.(*OnlineMeeting))
            }
            m.SetOnlineMeetings(res)
        }
        return nil
    }
    res["onPremisesDistinguishedName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesDistinguishedName(val)
        }
        return nil
    }
    res["onPremisesDomainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesDomainName(val)
        }
        return nil
    }
    res["onPremisesExtensionAttributes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesExtensionAttributes() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesExtensionAttributes(val.(*OnPremisesExtensionAttributes))
        }
        return nil
    }
    res["onPremisesImmutableId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesImmutableId(val)
        }
        return nil
    }
    res["onPremisesLastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesLastSyncDateTime(val)
        }
        return nil
    }
    res["onPremisesProvisioningErrors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesProvisioningError() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesProvisioningError, len(val))
            for i, v := range val {
                res[i] = *(v.(*OnPremisesProvisioningError))
            }
            m.SetOnPremisesProvisioningErrors(res)
        }
        return nil
    }
    res["onPremisesSamAccountName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSamAccountName(val)
        }
        return nil
    }
    res["onPremisesSecurityIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSecurityIdentifier(val)
        }
        return nil
    }
    res["onPremisesSyncEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSyncEnabled(val)
        }
        return nil
    }
    res["onPremisesUserPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesUserPrincipalName(val)
        }
        return nil
    }
    res["otherMails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOtherMails(res)
        }
        return nil
    }
    res["outlook"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutlookUser() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutlook(val.(*OutlookUser))
        }
        return nil
    }
    res["ownedDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetOwnedDevices(res)
        }
        return nil
    }
    res["ownedObjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetOwnedObjects(res)
        }
        return nil
    }
    res["passwordPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordPolicies(val)
        }
        return nil
    }
    res["passwordProfile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPasswordProfile() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordProfile(val.(*PasswordProfile))
        }
        return nil
    }
    res["pastProjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPastProjects(res)
        }
        return nil
    }
    res["people"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPerson() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Person, len(val))
            for i, v := range val {
                res[i] = *(v.(*Person))
            }
            m.SetPeople(res)
        }
        return nil
    }
    res["photo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProfilePhoto() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoto(val.(*ProfilePhoto))
        }
        return nil
    }
    res["photos"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProfilePhoto() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProfilePhoto, len(val))
            for i, v := range val {
                res[i] = *(v.(*ProfilePhoto))
            }
            m.SetPhotos(res)
        }
        return nil
    }
    res["planner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerUser() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlanner(val.(*PlannerUser))
        }
        return nil
    }
    res["postalCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostalCode(val)
        }
        return nil
    }
    res["preferredLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredLanguage(val)
        }
        return nil
    }
    res["preferredName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredName(val)
        }
        return nil
    }
    res["presence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPresence() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPresence(val.(*Presence))
        }
        return nil
    }
    res["provisionedPlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisionedPlan() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisionedPlan, len(val))
            for i, v := range val {
                res[i] = *(v.(*ProvisionedPlan))
            }
            m.SetProvisionedPlans(res)
        }
        return nil
    }
    res["proxyAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetProxyAddresses(res)
        }
        return nil
    }
    res["registeredDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetRegisteredDevices(res)
        }
        return nil
    }
    res["responsibilities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetResponsibilities(res)
        }
        return nil
    }
    res["schools"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSchools(res)
        }
        return nil
    }
    res["scopedRoleMemberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewScopedRoleMembership() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScopedRoleMembership, len(val))
            for i, v := range val {
                res[i] = *(v.(*ScopedRoleMembership))
            }
            m.SetScopedRoleMemberOf(res)
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(*UserSettings))
        }
        return nil
    }
    res["showInAddressList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowInAddressList(val)
        }
        return nil
    }
    res["signInSessionsValidFromDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInSessionsValidFromDateTime(val)
        }
        return nil
    }
    res["skills"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSkills(res)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["streetAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStreetAddress(val)
        }
        return nil
    }
    res["surname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSurname(val)
        }
        return nil
    }
    res["teamwork"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserTeamwork() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamwork(val.(*UserTeamwork))
        }
        return nil
    }
    res["todo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTodo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTodo(val.(*Todo))
        }
        return nil
    }
    res["transitiveMemberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetTransitiveMemberOf(res)
        }
        return nil
    }
    res["usageLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsageLocation(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["userType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserType(val)
        }
        return nil
    }
    return res
}
func (m *User) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *User) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("aboutMe", m.GetAboutMe())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("accountEnabled", m.GetAccountEnabled())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetActivities()))
        for i, v := range m.GetActivities() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("activities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ageGroup", m.GetAgeGroup())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAgreementAcceptances()))
        for i, v := range m.GetAgreementAcceptances() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("agreementAcceptances", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppRoleAssignments()))
        for i, v := range m.GetAppRoleAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appRoleAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignedLicenses()))
        for i, v := range m.GetAssignedLicenses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignedLicenses", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignedPlans()))
        for i, v := range m.GetAssignedPlans() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignedPlans", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authentication", m.GetAuthentication())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("birthday", m.GetBirthday())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("businessPhones", m.GetBusinessPhones())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("calendar", m.GetCalendar())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCalendarGroups()))
        for i, v := range m.GetCalendarGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("calendarGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCalendars()))
        for i, v := range m.GetCalendars() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("calendars", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCalendarView()))
        for i, v := range m.GetCalendarView() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("calendarView", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetChats()))
        for i, v := range m.GetChats() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("chats", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("companyName", m.GetCompanyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("consentProvidedForMinor", m.GetConsentProvidedForMinor())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetContactFolders()))
        for i, v := range m.GetContactFolders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("contactFolders", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetContacts()))
        for i, v := range m.GetContacts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("contacts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("country", m.GetCountry())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCreatedObjects()))
        for i, v := range m.GetCreatedObjects() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("createdObjects", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("creationType", m.GetCreationType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("department", m.GetDepartment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceEnrollmentLimit", m.GetDeviceEnrollmentLimit())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceManagementTroubleshootingEvents()))
        for i, v := range m.GetDeviceManagementTroubleshootingEvents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceManagementTroubleshootingEvents", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDirectReports()))
        for i, v := range m.GetDirectReports() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("directReports", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("drive", m.GetDrive())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDrives()))
        for i, v := range m.GetDrives() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("drives", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("employeeHireDate", m.GetEmployeeHireDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("employeeId", m.GetEmployeeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("employeeOrgData", m.GetEmployeeOrgData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("employeeType", m.GetEmployeeType())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEvents()))
        for i, v := range m.GetEvents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("events", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalUserState", m.GetExternalUserState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("externalUserStateChangeDateTime", m.GetExternalUserStateChangeDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("faxNumber", m.GetFaxNumber())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFollowedSites()))
        for i, v := range m.GetFollowedSites() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("followedSites", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("givenName", m.GetGivenName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("hireDate", m.GetHireDate())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIdentities()))
        for i, v := range m.GetIdentities() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("identities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("imAddresses", m.GetImAddresses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("inferenceClassification", m.GetInferenceClassification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("insights", m.GetInsights())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("interests", m.GetInterests())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isResourceAccount", m.GetIsResourceAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("jobTitle", m.GetJobTitle())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetJoinedTeams()))
        for i, v := range m.GetJoinedTeams() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("joinedTeams", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastPasswordChangeDateTime", m.GetLastPasswordChangeDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("legalAgeGroupClassification", m.GetLegalAgeGroupClassification())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLicenseAssignmentStates()))
        for i, v := range m.GetLicenseAssignmentStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("licenseAssignmentStates", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLicenseDetails()))
        for i, v := range m.GetLicenseDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("licenseDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mail", m.GetMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mailboxSettings", m.GetMailboxSettings())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMailFolders()))
        for i, v := range m.GetMailFolders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("mailFolders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mailNickname", m.GetMailNickname())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedAppRegistrations()))
        for i, v := range m.GetManagedAppRegistrations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedAppRegistrations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedDevices()))
        for i, v := range m.GetManagedDevices() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedDevices", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("manager", m.GetManager())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMemberOf()))
        for i, v := range m.GetMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("memberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMessages()))
        for i, v := range m.GetMessages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("messages", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mobilePhone", m.GetMobilePhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mySite", m.GetMySite())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOauth2PermissionGrants()))
        for i, v := range m.GetOauth2PermissionGrants() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("oauth2PermissionGrants", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("officeLocation", m.GetOfficeLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onenote", m.GetOnenote())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOnlineMeetings()))
        for i, v := range m.GetOnlineMeetings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("onlineMeetings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesDistinguishedName", m.GetOnPremisesDistinguishedName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesDomainName", m.GetOnPremisesDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onPremisesExtensionAttributes", m.GetOnPremisesExtensionAttributes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesImmutableId", m.GetOnPremisesImmutableId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("onPremisesLastSyncDateTime", m.GetOnPremisesLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOnPremisesProvisioningErrors()))
        for i, v := range m.GetOnPremisesProvisioningErrors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("onPremisesProvisioningErrors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesSamAccountName", m.GetOnPremisesSamAccountName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesSecurityIdentifier", m.GetOnPremisesSecurityIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onPremisesSyncEnabled", m.GetOnPremisesSyncEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesUserPrincipalName", m.GetOnPremisesUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("otherMails", m.GetOtherMails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("outlook", m.GetOutlook())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOwnedDevices()))
        for i, v := range m.GetOwnedDevices() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("ownedDevices", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOwnedObjects()))
        for i, v := range m.GetOwnedObjects() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("ownedObjects", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("passwordPolicies", m.GetPasswordPolicies())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("passwordProfile", m.GetPasswordProfile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("pastProjects", m.GetPastProjects())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPeople()))
        for i, v := range m.GetPeople() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("people", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("photo", m.GetPhoto())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPhotos()))
        for i, v := range m.GetPhotos() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("photos", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("planner", m.GetPlanner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("postalCode", m.GetPostalCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredLanguage", m.GetPreferredLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredName", m.GetPreferredName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("presence", m.GetPresence())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProvisionedPlans()))
        for i, v := range m.GetProvisionedPlans() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("provisionedPlans", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("proxyAddresses", m.GetProxyAddresses())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRegisteredDevices()))
        for i, v := range m.GetRegisteredDevices() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("registeredDevices", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("responsibilities", m.GetResponsibilities())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("schools", m.GetSchools())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetScopedRoleMemberOf()))
        for i, v := range m.GetScopedRoleMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("scopedRoleMemberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showInAddressList", m.GetShowInAddressList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("signInSessionsValidFromDateTime", m.GetSignInSessionsValidFromDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("skills", m.GetSkills())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("streetAddress", m.GetStreetAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("surname", m.GetSurname())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teamwork", m.GetTeamwork())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("todo", m.GetTodo())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTransitiveMemberOf()))
        for i, v := range m.GetTransitiveMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("transitiveMemberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("usageLocation", m.GetUsageLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userType", m.GetUserType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAboutMe sets the aboutMe property value. A freeform text entry field for the user to describe themselves. Returned only on $select.
func (m *User) SetAboutMe(value *string)() {
    if m != nil {
        m.aboutMe = value
    }
}
// SetAccountEnabled sets the accountEnabled property value. true if the account is enabled; otherwise, false. This property is required when a user is created. Returned only on $select. Supports $filter (eq, ne, not, and in).
func (m *User) SetAccountEnabled(value *bool)() {
    if m != nil {
        m.accountEnabled = value
    }
}
// SetActivities sets the activities property value. The user's activities across devices. Read-only. Nullable.
func (m *User) SetActivities(value []UserActivity)() {
    if m != nil {
        m.activities = value
    }
}
// SetAgeGroup sets the ageGroup property value. Sets the age group of the user. Allowed values: null, minor, notAdult and adult. Refer to the legal age group property definitions for further information. Returned only on $select. Supports $filter (eq, ne, not, and in).
func (m *User) SetAgeGroup(value *string)() {
    if m != nil {
        m.ageGroup = value
    }
}
// SetAgreementAcceptances sets the agreementAcceptances property value. The user's terms of use acceptance statuses. Read-only. Nullable.
func (m *User) SetAgreementAcceptances(value []AgreementAcceptance)() {
    if m != nil {
        m.agreementAcceptances = value
    }
}
// SetAppRoleAssignments sets the appRoleAssignments property value. Represents the app roles a user has been granted for an application. Supports $expand.
func (m *User) SetAppRoleAssignments(value []AppRoleAssignment)() {
    if m != nil {
        m.appRoleAssignments = value
    }
}
// SetAssignedLicenses sets the assignedLicenses property value. The licenses that are assigned to the user, including inherited (group-based) licenses.  Not nullable. Returned only on $select. Supports $filter (eq and not).
func (m *User) SetAssignedLicenses(value []AssignedLicense)() {
    if m != nil {
        m.assignedLicenses = value
    }
}
// SetAssignedPlans sets the assignedPlans property value. The plans that are assigned to the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq and not).
func (m *User) SetAssignedPlans(value []AssignedPlan)() {
    if m != nil {
        m.assignedPlans = value
    }
}
// SetAuthentication sets the authentication property value. 
func (m *User) SetAuthentication(value *Authentication)() {
    if m != nil {
        m.authentication = value
    }
}
// SetBirthday sets the birthday property value. The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.
func (m *User) SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.birthday = value
    }
}
// SetBusinessPhones sets the businessPhones property value. The telephone numbers for the user. NOTE: Although this is a string collection, only one number can be set for this property. Read-only for users synced from on-premises directory. Returned by default. Supports $filter (eq, not, ge, le, startsWith).
func (m *User) SetBusinessPhones(value []string)() {
    if m != nil {
        m.businessPhones = value
    }
}
// SetCalendar sets the calendar property value. The user's primary calendar. Read-only.
func (m *User) SetCalendar(value *Calendar)() {
    if m != nil {
        m.calendar = value
    }
}
// SetCalendarGroups sets the calendarGroups property value. The user's calendar groups. Read-only. Nullable.
func (m *User) SetCalendarGroups(value []CalendarGroup)() {
    if m != nil {
        m.calendarGroups = value
    }
}
// SetCalendars sets the calendars property value. The user's calendars. Read-only. Nullable.
func (m *User) SetCalendars(value []Calendar)() {
    if m != nil {
        m.calendars = value
    }
}
// SetCalendarView sets the calendarView property value. The calendar view for the calendar. Read-only. Nullable.
func (m *User) SetCalendarView(value []Event)() {
    if m != nil {
        m.calendarView = value
    }
}
// SetChats sets the chats property value. 
func (m *User) SetChats(value []Chat)() {
    if m != nil {
        m.chats = value
    }
}
// SetCity sets the city property value. The city in which the user is located. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetCity(value *string)() {
    if m != nil {
        m.city = value
    }
}
// SetCompanyName sets the companyName property value. The company name which the user is associated. This property can be useful for describing the company that an external user comes from. The maximum length of the company name is 64 characters.Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetCompanyName(value *string)() {
    if m != nil {
        m.companyName = value
    }
}
// SetConsentProvidedForMinor sets the consentProvidedForMinor property value. Sets whether consent has been obtained for minors. Allowed values: null, granted, denied and notRequired. Refer to the legal age group property definitions for further information. Returned only on $select. Supports $filter (eq, ne, not, and in).
func (m *User) SetConsentProvidedForMinor(value *string)() {
    if m != nil {
        m.consentProvidedForMinor = value
    }
}
// SetContactFolders sets the contactFolders property value. The user's contacts folders. Read-only. Nullable.
func (m *User) SetContactFolders(value []ContactFolder)() {
    if m != nil {
        m.contactFolders = value
    }
}
// SetContacts sets the contacts property value. The user's contacts. Read-only. Nullable.
func (m *User) SetContacts(value []Contact)() {
    if m != nil {
        m.contacts = value
    }
}
// SetCountry sets the country property value. The country/region in which the user is located; for example, US or UK. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetCountry(value *string)() {
    if m != nil {
        m.country = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The created date of the user object. Read-only. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetCreatedObjects sets the createdObjects property value. Directory objects that were created by the user. Read-only. Nullable.
func (m *User) SetCreatedObjects(value []DirectoryObject)() {
    if m != nil {
        m.createdObjects = value
    }
}
// SetCreationType sets the creationType property value. Indicates whether the user account was created through one of the following methods:  As a regular school or work account (null). As an external account (Invitation). As a local account for an Azure Active Directory B2C tenant (LocalAccount). Through self-service sign-up by an internal user using email verification (EmailVerified). Through self-service sign-up by an external user signing up through a link that is part of a user flow (SelfServiceSignUp). Read-only.Returned only on $select. Supports $filter (eq, ne, not, in).
func (m *User) SetCreationType(value *string)() {
    if m != nil {
        m.creationType = value
    }
}
// SetDepartment sets the department property value. The name for the department in which the user works. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, and eq on null values).
func (m *User) SetDepartment(value *string)() {
    if m != nil {
        m.department = value
    }
}
// SetDeviceEnrollmentLimit sets the deviceEnrollmentLimit property value. The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000.
func (m *User) SetDeviceEnrollmentLimit(value *int32)() {
    if m != nil {
        m.deviceEnrollmentLimit = value
    }
}
// SetDeviceManagementTroubleshootingEvents sets the deviceManagementTroubleshootingEvents property value. The list of troubleshooting events for this user.
func (m *User) SetDeviceManagementTroubleshootingEvents(value []DeviceManagementTroubleshootingEvent)() {
    if m != nil {
        m.deviceManagementTroubleshootingEvents = value
    }
}
// SetDirectReports sets the directReports property value. The users and contacts that report to the user. (The users and contacts that have their manager property set to this user.) Read-only. Nullable. Supports $expand.
func (m *User) SetDirectReports(value []DirectoryObject)() {
    if m != nil {
        m.directReports = value
    }
}
// SetDisplayName sets the displayName property value. The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial and last name. This property is required when a user is created and it cannot be cleared during updates. Maximum length is 256 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values), $orderBy, and $search.
func (m *User) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDrive sets the drive property value. The user's OneDrive. Read-only.
func (m *User) SetDrive(value *Drive)() {
    if m != nil {
        m.drive = value
    }
}
// SetDrives sets the drives property value. A collection of drives available for this user. Read-only.
func (m *User) SetDrives(value []Drive)() {
    if m != nil {
        m.drives = value
    }
}
// SetEmployeeHireDate sets the employeeHireDate property value. The date and time when the user was hired or will start work in case of a future hire. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) SetEmployeeHireDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.employeeHireDate = value
    }
}
// SetEmployeeId sets the employeeId property value. The employee identifier assigned to the user by the organization. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) SetEmployeeId(value *string)() {
    if m != nil {
        m.employeeId = value
    }
}
// SetEmployeeOrgData sets the employeeOrgData property value. Represents organization data (e.g. division and costCenter) associated with a user. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in).
func (m *User) SetEmployeeOrgData(value *EmployeeOrgData)() {
    if m != nil {
        m.employeeOrgData = value
    }
}
// SetEmployeeType sets the employeeType property value. Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith).
func (m *User) SetEmployeeType(value *string)() {
    if m != nil {
        m.employeeType = value
    }
}
// SetEvents sets the events property value. The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *User) SetEvents(value []Event)() {
    if m != nil {
        m.events = value
    }
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the user. Read-only. Nullable.
func (m *User) SetExtensions(value []Extension)() {
    if m != nil {
        m.extensions = value
    }
}
// SetExternalUserState sets the externalUserState property value. For an external user invited to the tenant using the invitation API, this property represents the invited user's invitation status. For invited users, the state can be PendingAcceptance or Accepted, or null for all other users. Returned only on $select. Supports $filter (eq, ne, not , in).
func (m *User) SetExternalUserState(value *string)() {
    if m != nil {
        m.externalUserState = value
    }
}
// SetExternalUserStateChangeDateTime sets the externalUserStateChangeDateTime property value. Shows the timestamp for the latest change to the externalUserState property. Returned only on $select. Supports $filter (eq, ne, not , in).
func (m *User) SetExternalUserStateChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.externalUserStateChangeDateTime = value
    }
}
// SetFaxNumber sets the faxNumber property value. The fax number of the user. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) SetFaxNumber(value *string)() {
    if m != nil {
        m.faxNumber = value
    }
}
// SetFollowedSites sets the followedSites property value. 
func (m *User) SetFollowedSites(value []Site)() {
    if m != nil {
        m.followedSites = value
    }
}
// SetGivenName sets the givenName property value. The given name (first name) of the user. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) SetGivenName(value *string)() {
    if m != nil {
        m.givenName = value
    }
}
// SetHireDate sets the hireDate property value. The hire date of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.  Note: This property is specific to SharePoint Online. We recommend using the native employeeHireDate property to set and update hire date values using Microsoft Graph APIs.
func (m *User) SetHireDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.hireDate = value
    }
}
// SetIdentities sets the identities property value. Represents the identities that can be used to sign in to this user account. An identity can be provided by Microsoft (also known as a local account), by organizations, or by social identity providers such as Facebook, Google, and Microsoft, and tied to a user account. May contain multiple items with the same signInType value. Returned only on $select. Supports $filter (eq) including on null values, only where the signInType is not userPrincipalName.
func (m *User) SetIdentities(value []ObjectIdentity)() {
    if m != nil {
        m.identities = value
    }
}
// SetImAddresses sets the imAddresses property value. The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user. Read-only. Returned only on $select. Supports $filter (eq, not, ge, le, startsWith).
func (m *User) SetImAddresses(value []string)() {
    if m != nil {
        m.imAddresses = value
    }
}
// SetInferenceClassification sets the inferenceClassification property value. Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *User) SetInferenceClassification(value *InferenceClassification)() {
    if m != nil {
        m.inferenceClassification = value
    }
}
// SetInsights sets the insights property value. Read-only. Nullable.
func (m *User) SetInsights(value *OfficeGraphInsights)() {
    if m != nil {
        m.insights = value
    }
}
// SetInterests sets the interests property value. A list for the user to describe their interests. Returned only on $select.
func (m *User) SetInterests(value []string)() {
    if m != nil {
        m.interests = value
    }
}
// SetIsResourceAccount sets the isResourceAccount property value. Do not use – reserved for future use.
func (m *User) SetIsResourceAccount(value *bool)() {
    if m != nil {
        m.isResourceAccount = value
    }
}
// SetJobTitle sets the jobTitle property value. The user's job title. Maximum length is 128 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values).
func (m *User) SetJobTitle(value *string)() {
    if m != nil {
        m.jobTitle = value
    }
}
// SetJoinedTeams sets the joinedTeams property value. The Microsoft Teams teams that the user is a member of. Read-only. Nullable.
func (m *User) SetJoinedTeams(value []Team)() {
    if m != nil {
        m.joinedTeams = value
    }
}
// SetLastPasswordChangeDateTime sets the lastPasswordChangeDateTime property value. The time when this Azure AD user last changed their password or when their password was created, whichever date the latest action was performed. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.
func (m *User) SetLastPasswordChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastPasswordChangeDateTime = value
    }
}
// SetLegalAgeGroupClassification sets the legalAgeGroupClassification property value. Used by enterprise applications to determine the legal age group of the user. This property is read-only and calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, minorWithOutParentalConsent, minorWithParentalConsent, minorNoParentalConsentRequired, notAdult and adult. Refer to the legal age group property definitions for further information. Returned only on $select.
func (m *User) SetLegalAgeGroupClassification(value *string)() {
    if m != nil {
        m.legalAgeGroupClassification = value
    }
}
// SetLicenseAssignmentStates sets the licenseAssignmentStates property value. State of license assignments for this user. Read-only. Returned only on $select.
func (m *User) SetLicenseAssignmentStates(value []LicenseAssignmentState)() {
    if m != nil {
        m.licenseAssignmentStates = value
    }
}
// SetLicenseDetails sets the licenseDetails property value. A collection of this user's license details. Read-only.
func (m *User) SetLicenseDetails(value []LicenseDetails)() {
    if m != nil {
        m.licenseDetails = value
    }
}
// SetMail sets the mail property value. The SMTP address for the user, for example, jeff@contoso.onmicrosoft.com.Changes to this property will also update the user's proxyAddresses collection to include the value as an SMTP address. For Azure AD B2C accounts, this property can be updated up to only ten times with unique SMTP addresses. This property cannot contain accent characters.Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith, and eq on null values).
func (m *User) SetMail(value *string)() {
    if m != nil {
        m.mail = value
    }
}
// SetMailboxSettings sets the mailboxSettings property value. Settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale and time zone.Returned only on $select.
func (m *User) SetMailboxSettings(value *MailboxSettings)() {
    if m != nil {
        m.mailboxSettings = value
    }
}
// SetMailFolders sets the mailFolders property value. The user's mail folders. Read-only. Nullable.
func (m *User) SetMailFolders(value []MailFolder)() {
    if m != nil {
        m.mailFolders = value
    }
}
// SetMailNickname sets the mailNickname property value. The mail alias for the user. This property must be specified when a user is created. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetMailNickname(value *string)() {
    if m != nil {
        m.mailNickname = value
    }
}
// SetManagedAppRegistrations sets the managedAppRegistrations property value. Zero or more managed app registrations that belong to the user.
func (m *User) SetManagedAppRegistrations(value []ManagedAppRegistration)() {
    if m != nil {
        m.managedAppRegistrations = value
    }
}
// SetManagedDevices sets the managedDevices property value. The managed devices associated with the user.
func (m *User) SetManagedDevices(value []ManagedDevice)() {
    if m != nil {
        m.managedDevices = value
    }
}
// SetManager sets the manager property value. The user or contact that is this user's manager. Read-only. (HTTP Methods: GET, PUT, DELETE.). Supports $expand.
func (m *User) SetManager(value *DirectoryObject)() {
    if m != nil {
        m.manager = value
    }
}
// SetMemberOf sets the memberOf property value. The groups and directory roles that the user is a member of. Read-only. Nullable. Supports $expand.
func (m *User) SetMemberOf(value []DirectoryObject)() {
    if m != nil {
        m.memberOf = value
    }
}
// SetMessages sets the messages property value. The messages in a mailbox or folder. Read-only. Nullable.
func (m *User) SetMessages(value []Message)() {
    if m != nil {
        m.messages = value
    }
}
// SetMobilePhone sets the mobilePhone property value. The primary cellular telephone number for the user. Read-only for users synced from on-premises directory. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetMobilePhone(value *string)() {
    if m != nil {
        m.mobilePhone = value
    }
}
// SetMySite sets the mySite property value. The URL for the user's personal site. Returned only on $select.
func (m *User) SetMySite(value *string)() {
    if m != nil {
        m.mySite = value
    }
}
// SetOauth2PermissionGrants sets the oauth2PermissionGrants property value. 
func (m *User) SetOauth2PermissionGrants(value []OAuth2PermissionGrant)() {
    if m != nil {
        m.oauth2PermissionGrants = value
    }
}
// SetOfficeLocation sets the officeLocation property value. The office location in the user's place of business. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetOfficeLocation(value *string)() {
    if m != nil {
        m.officeLocation = value
    }
}
// SetOnenote sets the onenote property value. Read-only.
func (m *User) SetOnenote(value *Onenote)() {
    if m != nil {
        m.onenote = value
    }
}
// SetOnlineMeetings sets the onlineMeetings property value. 
func (m *User) SetOnlineMeetings(value []OnlineMeeting)() {
    if m != nil {
        m.onlineMeetings = value
    }
}
// SetOnPremisesDistinguishedName sets the onPremisesDistinguishedName property value. Contains the on-premises Active Directory distinguished name or DN. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select.
func (m *User) SetOnPremisesDistinguishedName(value *string)() {
    if m != nil {
        m.onPremisesDistinguishedName = value
    }
}
// SetOnPremisesDomainName sets the onPremisesDomainName property value. Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select.
func (m *User) SetOnPremisesDomainName(value *string)() {
    if m != nil {
        m.onPremisesDomainName = value
    }
}
// SetOnPremisesExtensionAttributes sets the onPremisesExtensionAttributes property value. Contains extensionAttributes1-15 for the user. The individual extension attributes are neither selectable nor filterable. For an onPremisesSyncEnabled user, the source of authority for this set of properties is the on-premises and is read-only. For a cloud-only user (where onPremisesSyncEnabled is false), these properties can be set during creation or update of a user object.  For a cloud-only user previously synced from on-premises Active Directory, these properties are read-only in Microsoft Graph but can be fully managed through the Exchange Admin Center or the Exchange Online V2 module in PowerShell. These extension attributes are also known as Exchange custom attributes 1-15.
func (m *User) SetOnPremisesExtensionAttributes(value *OnPremisesExtensionAttributes)() {
    if m != nil {
        m.onPremisesExtensionAttributes = value
    }
}
// SetOnPremisesImmutableId sets the onPremisesImmutableId property value. This property is used to associate an on-premises Active Directory user account to their Azure AD user object. This property must be specified when creating a new user account in the Graph if you are using a federated domain for the user's userPrincipalName (UPN) property. NOTE: The $ and _ characters cannot be used when specifying this property. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in)..
func (m *User) SetOnPremisesImmutableId(value *string)() {
    if m != nil {
        m.onPremisesImmutableId = value
    }
}
// SetOnPremisesLastSyncDateTime sets the onPremisesLastSyncDateTime property value. Indicates the last time at which the object was synced with the on-premises directory; for example: 2013-02-16T03:04:54Z. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in).
func (m *User) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.onPremisesLastSyncDateTime = value
    }
}
// SetOnPremisesProvisioningErrors sets the onPremisesProvisioningErrors property value. Errors when using Microsoft synchronization product during provisioning. Returned only on $select. Supports $filter (eq, not, ge, le).
func (m *User) SetOnPremisesProvisioningErrors(value []OnPremisesProvisioningError)() {
    if m != nil {
        m.onPremisesProvisioningErrors = value
    }
}
// SetOnPremisesSamAccountName sets the onPremisesSamAccountName property value. Contains the on-premises samAccountName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith).
func (m *User) SetOnPremisesSamAccountName(value *string)() {
    if m != nil {
        m.onPremisesSamAccountName = value
    }
}
// SetOnPremisesSecurityIdentifier sets the onPremisesSecurityIdentifier property value. Contains the on-premises security identifier (SID) for the user that was synchronized from on-premises to the cloud. Read-only. Returned only on $select. Supports $filter (eq) on null values only.
func (m *User) SetOnPremisesSecurityIdentifier(value *string)() {
    if m != nil {
        m.onPremisesSecurityIdentifier = value
    }
}
// SetOnPremisesSyncEnabled sets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *User) SetOnPremisesSyncEnabled(value *bool)() {
    if m != nil {
        m.onPremisesSyncEnabled = value
    }
}
// SetOnPremisesUserPrincipalName sets the onPremisesUserPrincipalName property value. Contains the on-premises userPrincipalName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith).
func (m *User) SetOnPremisesUserPrincipalName(value *string)() {
    if m != nil {
        m.onPremisesUserPrincipalName = value
    }
}
// SetOtherMails sets the otherMails property value. A list of additional email addresses for the user; for example: ['bob@contoso.com', 'Robert@fabrikam.com']. NOTE: This property cannot contain accent characters. Returned only on $select. Supports $filter (eq, not, ge, le, in, startsWith).
func (m *User) SetOtherMails(value []string)() {
    if m != nil {
        m.otherMails = value
    }
}
// SetOutlook sets the outlook property value. Read-only.
func (m *User) SetOutlook(value *OutlookUser)() {
    if m != nil {
        m.outlook = value
    }
}
// SetOwnedDevices sets the ownedDevices property value. Devices that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *User) SetOwnedDevices(value []DirectoryObject)() {
    if m != nil {
        m.ownedDevices = value
    }
}
// SetOwnedObjects sets the ownedObjects property value. Directory objects that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *User) SetOwnedObjects(value []DirectoryObject)() {
    if m != nil {
        m.ownedObjects = value
    }
}
// SetPasswordPolicies sets the passwordPolicies property value. Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two may be specified together; for example: DisablePasswordExpiration, DisableStrongPassword. Returned only on $select. For more information on the default password policies, see Azure AD pasword policies. Supports $filter (ne, not, and eq on null values).
func (m *User) SetPasswordPolicies(value *string)() {
    if m != nil {
        m.passwordPolicies = value
    }
}
// SetPasswordProfile sets the passwordProfile property value. Specifies the password profile for the user. The profile contains the user’s password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required. NOTE: For Azure B2C tenants, the forceChangePasswordNextSignIn property should be set to false and instead use custom policies and user flows to force password reset at first logon. See Force password reset at first logon.Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *User) SetPasswordProfile(value *PasswordProfile)() {
    if m != nil {
        m.passwordProfile = value
    }
}
// SetPastProjects sets the pastProjects property value. A list for the user to enumerate their past projects. Returned only on $select.
func (m *User) SetPastProjects(value []string)() {
    if m != nil {
        m.pastProjects = value
    }
}
// SetPeople sets the people property value. People that are relevant to the user. Read-only. Nullable.
func (m *User) SetPeople(value []Person)() {
    if m != nil {
        m.people = value
    }
}
// SetPhoto sets the photo property value. The user's profile photo. Read-only.
func (m *User) SetPhoto(value *ProfilePhoto)() {
    if m != nil {
        m.photo = value
    }
}
// SetPhotos sets the photos property value. Read-only. Nullable.
func (m *User) SetPhotos(value []ProfilePhoto)() {
    if m != nil {
        m.photos = value
    }
}
// SetPlanner sets the planner property value. Entry-point to the Planner resource that might exist for a user. Read-only.
func (m *User) SetPlanner(value *PlannerUser)() {
    if m != nil {
        m.planner = value
    }
}
// SetPostalCode sets the postalCode property value. The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code. Maximum length is 40 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetPostalCode(value *string)() {
    if m != nil {
        m.postalCode = value
    }
}
// SetPreferredLanguage sets the preferredLanguage property value. The preferred language for the user. Should follow ISO 639-1 Code; for example en-US. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values)
func (m *User) SetPreferredLanguage(value *string)() {
    if m != nil {
        m.preferredLanguage = value
    }
}
// SetPreferredName sets the preferredName property value. The preferred name for the user. Returned only on $select.
func (m *User) SetPreferredName(value *string)() {
    if m != nil {
        m.preferredName = value
    }
}
// SetPresence sets the presence property value. 
func (m *User) SetPresence(value *Presence)() {
    if m != nil {
        m.presence = value
    }
}
// SetProvisionedPlans sets the provisionedPlans property value. The plans that are provisioned for the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq, not, ge, le).
func (m *User) SetProvisionedPlans(value []ProvisionedPlan)() {
    if m != nil {
        m.provisionedPlans = value
    }
}
// SetProxyAddresses sets the proxyAddresses property value. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. For Azure AD B2C accounts, this property has a limit of ten unique addresses. Read-only, Not nullable. Returned only on $select. Supports $filter (eq, not, ge, le, startsWith).
func (m *User) SetProxyAddresses(value []string)() {
    if m != nil {
        m.proxyAddresses = value
    }
}
// SetRegisteredDevices sets the registeredDevices property value. Devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *User) SetRegisteredDevices(value []DirectoryObject)() {
    if m != nil {
        m.registeredDevices = value
    }
}
// SetResponsibilities sets the responsibilities property value. A list for the user to enumerate their responsibilities. Returned only on $select.
func (m *User) SetResponsibilities(value []string)() {
    if m != nil {
        m.responsibilities = value
    }
}
// SetSchools sets the schools property value. A list for the user to enumerate the schools they have attended. Returned only on $select.
func (m *User) SetSchools(value []string)() {
    if m != nil {
        m.schools = value
    }
}
// SetScopedRoleMemberOf sets the scopedRoleMemberOf property value. The scoped-role administrative unit memberships for this user. Read-only. Nullable.
func (m *User) SetScopedRoleMemberOf(value []ScopedRoleMembership)() {
    if m != nil {
        m.scopedRoleMemberOf = value
    }
}
// SetSettings sets the settings property value. Read-only. Nullable.
func (m *User) SetSettings(value *UserSettings)() {
    if m != nil {
        m.settings = value
    }
}
// SetShowInAddressList sets the showInAddressList property value. true if the Outlook global address list should contain this user, otherwise false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false. Returned only on $select. Supports $filter (eq, ne, not, in).
func (m *User) SetShowInAddressList(value *bool)() {
    if m != nil {
        m.showInAddressList = value
    }
}
// SetSignInSessionsValidFromDateTime sets the signInSessionsValidFromDateTime property value. Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).  If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint. Read-only. Use revokeSignInSessions to reset. Returned only on $select.
func (m *User) SetSignInSessionsValidFromDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.signInSessionsValidFromDateTime = value
    }
}
// SetSkills sets the skills property value. A list for the user to enumerate their skills. Returned only on $select.
func (m *User) SetSkills(value []string)() {
    if m != nil {
        m.skills = value
    }
}
// SetState sets the state property value. The state or province in the user's address. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetState(value *string)() {
    if m != nil {
        m.state = value
    }
}
// SetStreetAddress sets the streetAddress property value. The street address of the user's place of business. Maximum length is 1024 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetStreetAddress(value *string)() {
    if m != nil {
        m.streetAddress = value
    }
}
// SetSurname sets the surname property value. The user's surname (family name or last name). Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetSurname(value *string)() {
    if m != nil {
        m.surname = value
    }
}
// SetTeamwork sets the teamwork property value. A container for Microsoft Teams features available for the user. Read-only. Nullable.
func (m *User) SetTeamwork(value *UserTeamwork)() {
    if m != nil {
        m.teamwork = value
    }
}
// SetTodo sets the todo property value. Represents the To Do services available to a user.
func (m *User) SetTodo(value *Todo)() {
    if m != nil {
        m.todo = value
    }
}
// SetTransitiveMemberOf sets the transitiveMemberOf property value. 
func (m *User) SetTransitiveMemberOf(value []DirectoryObject)() {
    if m != nil {
        m.transitiveMemberOf = value
    }
}
// SetUsageLocation sets the usageLocation property value. A two letter country code (ISO standard 3166). Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.  Examples include: US, JP, and GB. Not nullable. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
func (m *User) SetUsageLocation(value *string)() {
    if m != nil {
        m.usageLocation = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization.NOTE: This property cannot contain accent characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith) and $orderBy.
func (m *User) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
// SetUserType sets the userType property value. A string value that can be used to classify user types in your directory, such as Member and Guest. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values). NOTE: For more information about the permissions for member and guest users, see What are the default user permissions in Azure Active Directory?
func (m *User) SetUserType(value *string)() {
    if m != nil {
        m.userType = value
    }
}
