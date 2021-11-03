package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type User struct {
    DirectoryObject
    // A freeform text entry field for the user to describe themselves. Returned only on $select.
    aboutMe *string;
    // true if the account is enabled; otherwise, false. This property is required when a user is created. Returned only on $select. Supports $filter (eq, ne, NOT, and in).
    accountEnabled *bool;
    // The user's activities across devices. Read-only. Nullable.
    activities []UserActivity;
    // Sets the age group of the user. Allowed values: null, minor, notAdult and adult. Refer to the legal age group property definitions for further information. Returned only on $select. Supports $filter (eq, ne, NOT, and in).
    ageGroup *string;
    // The user's terms of use acceptance statuses. Read-only. Nullable.
    agreementAcceptances []AgreementAcceptance;
    // Represents the app roles a user has been granted for an application. Supports $expand.
    appRoleAssignments []AppRoleAssignment;
    // The licenses that are assigned to the user, including inherited (group-based) licenses.  Not nullable. Returned only on $select. Supports $filter (eq and NOT).
    assignedLicenses []AssignedLicense;
    // The plans that are assigned to the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq and NOT).
    assignedPlans []AssignedPlan;
    // 
    authentication *Authentication;
    // The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.
    birthday *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The telephone numbers for the user. NOTE: Although this is a string collection, only one number can be set for this property. Read-only for users synced from on-premises directory. Returned by default. Supports $filter (eq and NOT).
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
    // The city in which the user is located. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    city *string;
    // The company name which the user is associated. This property can be useful for describing the company that an external user comes from. The maximum length of the company name is 64 characters.Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    companyName *string;
    // Sets whether consent has been obtained for minors. Allowed values: null, granted, denied and notRequired. Refer to the legal age group property definitions for further information. Returned only on $select. Supports $filter (eq, ne, NOT, and in).
    consentProvidedForMinor *string;
    // The user's contacts folders. Read-only. Nullable.
    contactFolders []ContactFolder;
    // The user's contacts. Read-only. Nullable.
    contacts []Contact;
    // The country/region in which the user is located; for example, US or UK. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    country *string;
    // The created date of the user object. Read-only. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, and in operators).
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Directory objects that were created by the user. Read-only. Nullable.
    createdObjects []DirectoryObject;
    // Indicates whether the user account was created through one of the following methods:  As a regular school or work account (null). As an external account (Invitation). As a local account for an Azure Active Directory B2C tenant (LocalAccount). Through self-service sign-up by an internal user using email verification (EmailVerified). Through self-service sign-up by an external user signing up through a link that is part of a user flow (SelfServiceSignUp). Read-only.Returned only on $select. Supports $filter (eq, ne, NOT, and in).
    creationType *string;
    // The name for the department in which the user works. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, and in operators).
    department *string;
    // The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000.
    deviceEnrollmentLimit *int32;
    // The list of troubleshooting events for this user.
    deviceManagementTroubleshootingEvents []DeviceManagementTroubleshootingEvent;
    // The users and contacts that report to the user. (The users and contacts that have their manager property set to this user.) Read-only. Nullable. Supports $expand.
    directReports []DirectoryObject;
    // The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial and last name. This property is required when a user is created and it cannot be cleared during updates. Maximum length is 256 characters. Returned by default. Supports $filter (eq, ne, NOT , ge, le, in, startsWith), $orderBy, and $search.
    displayName *string;
    // The user's OneDrive. Read-only.
    drive *Drive;
    // A collection of drives available for this user. Read-only.
    drives []Drive;
    // The date and time when the user was hired or will start work in case of a future hire. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, in).
    employeeHireDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The employee identifier assigned to the user by the organization. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, in, startsWith).
    employeeId *string;
    // Represents organization data (e.g. division and costCenter) associated with a user. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, in).
    employeeOrgData *EmployeeOrgData;
    // Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, in, startsWith).
    employeeType *string;
    // The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
    events []Event;
    // The collection of open extensions defined for the user. Read-only. Nullable.
    extensions []Extension;
    // For an external user invited to the tenant using the invitation API, this property represents the invited user's invitation status. For invited users, the state can be PendingAcceptance or Accepted, or null for all other users. Returned only on $select. Supports $filter (eq, ne, NOT , in).
    externalUserState *string;
    // Shows the timestamp for the latest change to the externalUserState property. Returned only on $select. Supports $filter (eq, ne, NOT , in).
    externalUserStateChangeDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The fax number of the user. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, in, startsWith).
    faxNumber *string;
    // 
    followedSites []Site;
    // The given name (first name) of the user. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, NOT , ge, le, in, startsWith).
    givenName *string;
    // The hire date of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.  Note: This property is specific to SharePoint Online. We recommend using the native employeeHireDate property to set and update hire date values using Microsoft Graph APIs.
    hireDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Represents the identities that can be used to sign in to this user account. An identity can be provided by Microsoft (also known as a local account), by organizations, or by social identity providers such as Facebook, Google, and Microsoft, and tied to a user account. May contain multiple items with the same signInType value. Returned only on $select. Supports $filter (eq) only where the signInType is not userPrincipalName.
    identities []ObjectIdentity;
    // The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user. Read-only. Returned only on $select. Supports $filter (eq, NOT, ge, le, startsWith).
    imAddresses []string;
    // Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
    inferenceClassification *InferenceClassification;
    // Read-only. Nullable.
    insights *OfficeGraphInsights;
    // A list for the user to describe their interests. Returned only on $select.
    interests []string;
    // Do not use – reserved for future use.
    isResourceAccount *bool;
    // The user's job title. Maximum length is 128 characters. Returned by default. Supports $filter (eq, ne, NOT , ge, le, in, startsWith).
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
    // The SMTP address for the user, for example, jeff@contoso.onmicrosoft.com.Changes to this property will also update the user's proxyAddresses collection to include the value as an SMTP address. For Azure AD B2C accounts, this property can be updated up to only ten times with unique SMTP addresses. This property cannot contain accent characters.Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, endsWith).
    mail *string;
    // Settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale and time zone.Returned only on $select.
    mailboxSettings *MailboxSettings;
    // The user's mail folders. Read-only. Nullable.
    mailFolders []MailFolder;
    // The mail alias for the user. This property must be specified when a user is created. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
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
    // The primary cellular telephone number for the user. Read-only for users synced from on-premises directory. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    mobilePhone *string;
    // The URL for the user's personal site. Returned only on $select.
    mySite *string;
    // 
    oauth2PermissionGrants []OAuth2PermissionGrant;
    // The office location in the user's place of business. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    officeLocation *string;
    // Read-only.
    onenote *Onenote;
    // 
    onlineMeetings []OnlineMeeting;
    // Contains the on-premises Active Directory distinguished name or DN. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select.
    onPremisesDistinguishedName *string;
    // Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select.
    onPremisesDomainName *string;
    // Contains extensionAttributes 1-15 for the user. Note that the individual extension attributes are neither selectable nor filterable. For an onPremisesSyncEnabled user, the source of authority for this set of properties is the on-premises and is read-only. For a cloud-only user (where onPremisesSyncEnabled is false), these properties may be set during creation or update. These extension attributes are also known as Exchange custom attributes 1-15. Returned only on $select. Supports $filter (eq, NOT, ge, le, in).
    onPremisesExtensionAttributes *OnPremisesExtensionAttributes;
    // This property is used to associate an on-premises Active Directory user account to their Azure AD user object. This property must be specified when creating a new user account in the Graph if you are using a federated domain for the user's userPrincipalName (UPN) property. NOTE: The $ and _ characters cannot be used when specifying this property. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in)..
    onPremisesImmutableId *string;
    // Indicates the last time at which the object was synced with the on-premises directory; for example: 2013-02-16T03:04:54Z. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in).
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Errors when using Microsoft synchronization product during provisioning. Returned only on $select. Supports $filter (eq, NOT, ge, le).
    onPremisesProvisioningErrors []OnPremisesProvisioningError;
    // Contains the on-premises samAccountName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    onPremisesSamAccountName *string;
    // Contains the on-premises security identifier (SID) for the user that was synchronized from on-premises to the cloud. Read-only. Returned only on $select.
    onPremisesSecurityIdentifier *string;
    // true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Returned only on $select. Supports $filter (eq, ne, NOT, in).
    onPremisesSyncEnabled *bool;
    // Contains the on-premises userPrincipalName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    onPremisesUserPrincipalName *string;
    // A list of additional email addresses for the user; for example: ['bob@contoso.com', 'Robert@fabrikam.com']. NOTE: This property cannot contain accent characters. Returned only on $select. Supports $filter (eq, NOT, ge, le, in, startsWith).
    otherMails []string;
    // Read-only.
    outlook *OutlookUser;
    // Devices that are owned by the user. Read-only. Nullable. Supports $expand.
    ownedDevices []DirectoryObject;
    // Directory objects that are owned by the user. Read-only. Nullable. Supports $expand.
    ownedObjects []DirectoryObject;
    // Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two may be specified together; for example: DisablePasswordExpiration, DisableStrongPassword. Returned only on $select. For more information on the default password policies, see Azure AD pasword policies. Supports $filter (ne, NOT).
    passwordPolicies *string;
    // Specifies the password profile for the user. The profile contains the user’s password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required. NOTE: For Azure B2C tenants, the forceChangePasswordNextSignIn property should be set to false and instead use custom policies and user flows to force password reset at first logon. See Force password reset at first logon.Returned only on $select. Supports $filter (eq, ne, NOT, in).
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
    // The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code. Maximum length is 40 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    postalCode *string;
    // The preferred language for the user. Should follow ISO 639-1 Code; for example en-US. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith)
    preferredLanguage *string;
    // The preferred name for the user. Returned only on $select.
    preferredName *string;
    // 
    presence *Presence;
    // The plans that are provisioned for the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq, NOT, ge, le).
    provisionedPlans []ProvisionedPlan;
    // For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. For Azure AD B2C accounts, this property has a limit of ten unique addresses. Read-only, Not nullable. Returned only on $select. Supports $filter (eq, NOT, ge, le, startsWith).
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
    // true if the Outlook global address list should contain this user, otherwise false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false. Returned only on $select. Supports $filter (eq, ne, NOT, in).
    showInAddressList *bool;
    // Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).  If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint. Read-only. Use revokeSignInSessions to reset. Returned only on $select.
    signInSessionsValidFromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // A list for the user to enumerate their skills. Returned only on $select.
    skills []string;
    // The state or province in the user's address. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    state *string;
    // The street address of the user's place of business. Maximum length is 1024 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    streetAddress *string;
    // The user's surname (family name or last name). Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    surname *string;
    // A container for Microsoft Teams features available for the user. Read-only. Nullable.
    teamwork *UserTeamwork;
    // Represents the To Do services available to a user.
    todo *Todo;
    // 
    transitiveMemberOf []DirectoryObject;
    // A two letter country code (ISO standard 3166). Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.  Examples include: US, JP, and GB. Not nullable. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    usageLocation *string;
    // The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization.NOTE: This property cannot contain accent characters. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, endsWith) and $orderBy.
    userPrincipalName *string;
    // A string value that can be used to classify user types in your directory, such as Member and Guest. Returned only on $select. Supports $filter (eq, ne, NOT, in). NOTE: For more information about the permissions for member and guest users, see What are the default user permissions in Azure Active Directory?
    userType *string;
}
// Instantiates a new user and sets the default values.
func NewUser()(*User) {
    m := &User{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// Gets the aboutMe property value. A freeform text entry field for the user to describe themselves. Returned only on $select.
func (m *User) GetAboutMe()(*string) {
    if m == nil {
        return nil
    } else {
        return m.aboutMe
    }
}
// Gets the accountEnabled property value. true if the account is enabled; otherwise, false. This property is required when a user is created. Returned only on $select. Supports $filter (eq, ne, NOT, and in).
func (m *User) GetAccountEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accountEnabled
    }
}
// Gets the activities property value. The user's activities across devices. Read-only. Nullable.
func (m *User) GetActivities()([]UserActivity) {
    if m == nil {
        return nil
    } else {
        return m.activities
    }
}
// Gets the ageGroup property value. Sets the age group of the user. Allowed values: null, minor, notAdult and adult. Refer to the legal age group property definitions for further information. Returned only on $select. Supports $filter (eq, ne, NOT, and in).
func (m *User) GetAgeGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ageGroup
    }
}
// Gets the agreementAcceptances property value. The user's terms of use acceptance statuses. Read-only. Nullable.
func (m *User) GetAgreementAcceptances()([]AgreementAcceptance) {
    if m == nil {
        return nil
    } else {
        return m.agreementAcceptances
    }
}
// Gets the appRoleAssignments property value. Represents the app roles a user has been granted for an application. Supports $expand.
func (m *User) GetAppRoleAssignments()([]AppRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.appRoleAssignments
    }
}
// Gets the assignedLicenses property value. The licenses that are assigned to the user, including inherited (group-based) licenses.  Not nullable. Returned only on $select. Supports $filter (eq and NOT).
func (m *User) GetAssignedLicenses()([]AssignedLicense) {
    if m == nil {
        return nil
    } else {
        return m.assignedLicenses
    }
}
// Gets the assignedPlans property value. The plans that are assigned to the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq and NOT).
func (m *User) GetAssignedPlans()([]AssignedPlan) {
    if m == nil {
        return nil
    } else {
        return m.assignedPlans
    }
}
// Gets the authentication property value. 
func (m *User) GetAuthentication()(*Authentication) {
    if m == nil {
        return nil
    } else {
        return m.authentication
    }
}
// Gets the birthday property value. The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.
func (m *User) GetBirthday()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.birthday
    }
}
// Gets the businessPhones property value. The telephone numbers for the user. NOTE: Although this is a string collection, only one number can be set for this property. Read-only for users synced from on-premises directory. Returned by default. Supports $filter (eq and NOT).
func (m *User) GetBusinessPhones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.businessPhones
    }
}
// Gets the calendar property value. The user's primary calendar. Read-only.
func (m *User) GetCalendar()(*Calendar) {
    if m == nil {
        return nil
    } else {
        return m.calendar
    }
}
// Gets the calendarGroups property value. The user's calendar groups. Read-only. Nullable.
func (m *User) GetCalendarGroups()([]CalendarGroup) {
    if m == nil {
        return nil
    } else {
        return m.calendarGroups
    }
}
// Gets the calendars property value. The user's calendars. Read-only. Nullable.
func (m *User) GetCalendars()([]Calendar) {
    if m == nil {
        return nil
    } else {
        return m.calendars
    }
}
// Gets the calendarView property value. The calendar view for the calendar. Read-only. Nullable.
func (m *User) GetCalendarView()([]Event) {
    if m == nil {
        return nil
    } else {
        return m.calendarView
    }
}
// Gets the chats property value. 
func (m *User) GetChats()([]Chat) {
    if m == nil {
        return nil
    } else {
        return m.chats
    }
}
// Gets the city property value. The city in which the user is located. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *User) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// Gets the companyName property value. The company name which the user is associated. This property can be useful for describing the company that an external user comes from. The maximum length of the company name is 64 characters.Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *User) GetCompanyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.companyName
    }
}
// Gets the consentProvidedForMinor property value. Sets whether consent has been obtained for minors. Allowed values: null, granted, denied and notRequired. Refer to the legal age group property definitions for further information. Returned only on $select. Supports $filter (eq, ne, NOT, and in).
func (m *User) GetConsentProvidedForMinor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.consentProvidedForMinor
    }
}
// Gets the contactFolders property value. The user's contacts folders. Read-only. Nullable.
func (m *User) GetContactFolders()([]ContactFolder) {
    if m == nil {
        return nil
    } else {
        return m.contactFolders
    }
}
// Gets the contacts property value. The user's contacts. Read-only. Nullable.
func (m *User) GetContacts()([]Contact) {
    if m == nil {
        return nil
    } else {
        return m.contacts
    }
}
// Gets the country property value. The country/region in which the user is located; for example, US or UK. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *User) GetCountry()(*string) {
    if m == nil {
        return nil
    } else {
        return m.country
    }
}
// Gets the createdDateTime property value. The created date of the user object. Read-only. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, and in operators).
func (m *User) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the createdObjects property value. Directory objects that were created by the user. Read-only. Nullable.
func (m *User) GetCreatedObjects()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.createdObjects
    }
}
// Gets the creationType property value. Indicates whether the user account was created through one of the following methods:  As a regular school or work account (null). As an external account (Invitation). As a local account for an Azure Active Directory B2C tenant (LocalAccount). Through self-service sign-up by an internal user using email verification (EmailVerified). Through self-service sign-up by an external user signing up through a link that is part of a user flow (SelfServiceSignUp). Read-only.Returned only on $select. Supports $filter (eq, ne, NOT, and in).
func (m *User) GetCreationType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.creationType
    }
}
// Gets the department property value. The name for the department in which the user works. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, and in operators).
func (m *User) GetDepartment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.department
    }
}
// Gets the deviceEnrollmentLimit property value. The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000.
func (m *User) GetDeviceEnrollmentLimit()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceEnrollmentLimit
    }
}
// Gets the deviceManagementTroubleshootingEvents property value. The list of troubleshooting events for this user.
func (m *User) GetDeviceManagementTroubleshootingEvents()([]DeviceManagementTroubleshootingEvent) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementTroubleshootingEvents
    }
}
// Gets the directReports property value. The users and contacts that report to the user. (The users and contacts that have their manager property set to this user.) Read-only. Nullable. Supports $expand.
func (m *User) GetDirectReports()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.directReports
    }
}
// Gets the displayName property value. The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial and last name. This property is required when a user is created and it cannot be cleared during updates. Maximum length is 256 characters. Returned by default. Supports $filter (eq, ne, NOT , ge, le, in, startsWith), $orderBy, and $search.
func (m *User) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the drive property value. The user's OneDrive. Read-only.
func (m *User) GetDrive()(*Drive) {
    if m == nil {
        return nil
    } else {
        return m.drive
    }
}
// Gets the drives property value. A collection of drives available for this user. Read-only.
func (m *User) GetDrives()([]Drive) {
    if m == nil {
        return nil
    } else {
        return m.drives
    }
}
// Gets the employeeHireDate property value. The date and time when the user was hired or will start work in case of a future hire. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, in).
func (m *User) GetEmployeeHireDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.employeeHireDate
    }
}
// Gets the employeeId property value. The employee identifier assigned to the user by the organization. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, in, startsWith).
func (m *User) GetEmployeeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.employeeId
    }
}
// Gets the employeeOrgData property value. Represents organization data (e.g. division and costCenter) associated with a user. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, in).
func (m *User) GetEmployeeOrgData()(*EmployeeOrgData) {
    if m == nil {
        return nil
    } else {
        return m.employeeOrgData
    }
}
// Gets the employeeType property value. Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, in, startsWith).
func (m *User) GetEmployeeType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.employeeType
    }
}
// Gets the events property value. The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *User) GetEvents()([]Event) {
    if m == nil {
        return nil
    } else {
        return m.events
    }
}
// Gets the extensions property value. The collection of open extensions defined for the user. Read-only. Nullable.
func (m *User) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// Gets the externalUserState property value. For an external user invited to the tenant using the invitation API, this property represents the invited user's invitation status. For invited users, the state can be PendingAcceptance or Accepted, or null for all other users. Returned only on $select. Supports $filter (eq, ne, NOT , in).
func (m *User) GetExternalUserState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalUserState
    }
}
// Gets the externalUserStateChangeDateTime property value. Shows the timestamp for the latest change to the externalUserState property. Returned only on $select. Supports $filter (eq, ne, NOT , in).
func (m *User) GetExternalUserStateChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.externalUserStateChangeDateTime
    }
}
// Gets the faxNumber property value. The fax number of the user. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, in, startsWith).
func (m *User) GetFaxNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.faxNumber
    }
}
// Gets the followedSites property value. 
func (m *User) GetFollowedSites()([]Site) {
    if m == nil {
        return nil
    } else {
        return m.followedSites
    }
}
// Gets the givenName property value. The given name (first name) of the user. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, NOT , ge, le, in, startsWith).
func (m *User) GetGivenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.givenName
    }
}
// Gets the hireDate property value. The hire date of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.  Note: This property is specific to SharePoint Online. We recommend using the native employeeHireDate property to set and update hire date values using Microsoft Graph APIs.
func (m *User) GetHireDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.hireDate
    }
}
// Gets the identities property value. Represents the identities that can be used to sign in to this user account. An identity can be provided by Microsoft (also known as a local account), by organizations, or by social identity providers such as Facebook, Google, and Microsoft, and tied to a user account. May contain multiple items with the same signInType value. Returned only on $select. Supports $filter (eq) only where the signInType is not userPrincipalName.
func (m *User) GetIdentities()([]ObjectIdentity) {
    if m == nil {
        return nil
    } else {
        return m.identities
    }
}
// Gets the imAddresses property value. The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user. Read-only. Returned only on $select. Supports $filter (eq, NOT, ge, le, startsWith).
func (m *User) GetImAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.imAddresses
    }
}
// Gets the inferenceClassification property value. Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *User) GetInferenceClassification()(*InferenceClassification) {
    if m == nil {
        return nil
    } else {
        return m.inferenceClassification
    }
}
// Gets the insights property value. Read-only. Nullable.
func (m *User) GetInsights()(*OfficeGraphInsights) {
    if m == nil {
        return nil
    } else {
        return m.insights
    }
}
// Gets the interests property value. A list for the user to describe their interests. Returned only on $select.
func (m *User) GetInterests()([]string) {
    if m == nil {
        return nil
    } else {
        return m.interests
    }
}
// Gets the isResourceAccount property value. Do not use – reserved for future use.
func (m *User) GetIsResourceAccount()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isResourceAccount
    }
}
// Gets the jobTitle property value. The user's job title. Maximum length is 128 characters. Returned by default. Supports $filter (eq, ne, NOT , ge, le, in, startsWith).
func (m *User) GetJobTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jobTitle
    }
}
// Gets the joinedTeams property value. The Microsoft Teams teams that the user is a member of. Read-only. Nullable.
func (m *User) GetJoinedTeams()([]Team) {
    if m == nil {
        return nil
    } else {
        return m.joinedTeams
    }
}
// Gets the lastPasswordChangeDateTime property value. The time when this Azure AD user last changed their password or when their password was created, whichever date the latest action was performed. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.
func (m *User) GetLastPasswordChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastPasswordChangeDateTime
    }
}
// Gets the legalAgeGroupClassification property value. Used by enterprise applications to determine the legal age group of the user. This property is read-only and calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, minorWithOutParentalConsent, minorWithParentalConsent, minorNoParentalConsentRequired, notAdult and adult. Refer to the legal age group property definitions for further information. Returned only on $select.
func (m *User) GetLegalAgeGroupClassification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.legalAgeGroupClassification
    }
}
// Gets the licenseAssignmentStates property value. State of license assignments for this user. Read-only. Returned only on $select.
func (m *User) GetLicenseAssignmentStates()([]LicenseAssignmentState) {
    if m == nil {
        return nil
    } else {
        return m.licenseAssignmentStates
    }
}
// Gets the licenseDetails property value. A collection of this user's license details. Read-only.
func (m *User) GetLicenseDetails()([]LicenseDetails) {
    if m == nil {
        return nil
    } else {
        return m.licenseDetails
    }
}
// Gets the mail property value. The SMTP address for the user, for example, jeff@contoso.onmicrosoft.com.Changes to this property will also update the user's proxyAddresses collection to include the value as an SMTP address. For Azure AD B2C accounts, this property can be updated up to only ten times with unique SMTP addresses. This property cannot contain accent characters.Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, endsWith).
func (m *User) GetMail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mail
    }
}
// Gets the mailboxSettings property value. Settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale and time zone.Returned only on $select.
func (m *User) GetMailboxSettings()(*MailboxSettings) {
    if m == nil {
        return nil
    } else {
        return m.mailboxSettings
    }
}
// Gets the mailFolders property value. The user's mail folders. Read-only. Nullable.
func (m *User) GetMailFolders()([]MailFolder) {
    if m == nil {
        return nil
    } else {
        return m.mailFolders
    }
}
// Gets the mailNickname property value. The mail alias for the user. This property must be specified when a user is created. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *User) GetMailNickname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mailNickname
    }
}
// Gets the managedAppRegistrations property value. Zero or more managed app registrations that belong to the user.
func (m *User) GetManagedAppRegistrations()([]ManagedAppRegistration) {
    if m == nil {
        return nil
    } else {
        return m.managedAppRegistrations
    }
}
// Gets the managedDevices property value. The managed devices associated with the user.
func (m *User) GetManagedDevices()([]ManagedDevice) {
    if m == nil {
        return nil
    } else {
        return m.managedDevices
    }
}
// Gets the manager property value. The user or contact that is this user's manager. Read-only. (HTTP Methods: GET, PUT, DELETE.). Supports $expand.
func (m *User) GetManager()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.manager
    }
}
// Gets the memberOf property value. The groups and directory roles that the user is a member of. Read-only. Nullable. Supports $expand.
func (m *User) GetMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.memberOf
    }
}
// Gets the messages property value. The messages in a mailbox or folder. Read-only. Nullable.
func (m *User) GetMessages()([]Message) {
    if m == nil {
        return nil
    } else {
        return m.messages
    }
}
// Gets the mobilePhone property value. The primary cellular telephone number for the user. Read-only for users synced from on-premises directory. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *User) GetMobilePhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mobilePhone
    }
}
// Gets the mySite property value. The URL for the user's personal site. Returned only on $select.
func (m *User) GetMySite()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mySite
    }
}
// Gets the oauth2PermissionGrants property value. 
func (m *User) GetOauth2PermissionGrants()([]OAuth2PermissionGrant) {
    if m == nil {
        return nil
    } else {
        return m.oauth2PermissionGrants
    }
}
// Gets the officeLocation property value. The office location in the user's place of business. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *User) GetOfficeLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.officeLocation
    }
}
// Gets the onenote property value. Read-only.
func (m *User) GetOnenote()(*Onenote) {
    if m == nil {
        return nil
    } else {
        return m.onenote
    }
}
// Gets the onlineMeetings property value. 
func (m *User) GetOnlineMeetings()([]OnlineMeeting) {
    if m == nil {
        return nil
    } else {
        return m.onlineMeetings
    }
}
// Gets the onPremisesDistinguishedName property value. Contains the on-premises Active Directory distinguished name or DN. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select.
func (m *User) GetOnPremisesDistinguishedName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesDistinguishedName
    }
}
// Gets the onPremisesDomainName property value. Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select.
func (m *User) GetOnPremisesDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesDomainName
    }
}
// Gets the onPremisesExtensionAttributes property value. Contains extensionAttributes 1-15 for the user. Note that the individual extension attributes are neither selectable nor filterable. For an onPremisesSyncEnabled user, the source of authority for this set of properties is the on-premises and is read-only. For a cloud-only user (where onPremisesSyncEnabled is false), these properties may be set during creation or update. These extension attributes are also known as Exchange custom attributes 1-15. Returned only on $select. Supports $filter (eq, NOT, ge, le, in).
func (m *User) GetOnPremisesExtensionAttributes()(*OnPremisesExtensionAttributes) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesExtensionAttributes
    }
}
// Gets the onPremisesImmutableId property value. This property is used to associate an on-premises Active Directory user account to their Azure AD user object. This property must be specified when creating a new user account in the Graph if you are using a federated domain for the user's userPrincipalName (UPN) property. NOTE: The $ and _ characters cannot be used when specifying this property. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in)..
func (m *User) GetOnPremisesImmutableId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesImmutableId
    }
}
// Gets the onPremisesLastSyncDateTime property value. Indicates the last time at which the object was synced with the on-premises directory; for example: 2013-02-16T03:04:54Z. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in).
func (m *User) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesLastSyncDateTime
    }
}
// Gets the onPremisesProvisioningErrors property value. Errors when using Microsoft synchronization product during provisioning. Returned only on $select. Supports $filter (eq, NOT, ge, le).
func (m *User) GetOnPremisesProvisioningErrors()([]OnPremisesProvisioningError) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesProvisioningErrors
    }
}
// Gets the onPremisesSamAccountName property value. Contains the on-premises samAccountName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *User) GetOnPremisesSamAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSamAccountName
    }
}
// Gets the onPremisesSecurityIdentifier property value. Contains the on-premises security identifier (SID) for the user that was synchronized from on-premises to the cloud. Read-only. Returned only on $select.
func (m *User) GetOnPremisesSecurityIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSecurityIdentifier
    }
}
// Gets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Returned only on $select. Supports $filter (eq, ne, NOT, in).
func (m *User) GetOnPremisesSyncEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSyncEnabled
    }
}
// Gets the onPremisesUserPrincipalName property value. Contains the on-premises userPrincipalName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *User) GetOnPremisesUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesUserPrincipalName
    }
}
// Gets the otherMails property value. A list of additional email addresses for the user; for example: ['bob@contoso.com', 'Robert@fabrikam.com']. NOTE: This property cannot contain accent characters. Returned only on $select. Supports $filter (eq, NOT, ge, le, in, startsWith).
func (m *User) GetOtherMails()([]string) {
    if m == nil {
        return nil
    } else {
        return m.otherMails
    }
}
// Gets the outlook property value. Read-only.
func (m *User) GetOutlook()(*OutlookUser) {
    if m == nil {
        return nil
    } else {
        return m.outlook
    }
}
// Gets the ownedDevices property value. Devices that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *User) GetOwnedDevices()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.ownedDevices
    }
}
// Gets the ownedObjects property value. Directory objects that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *User) GetOwnedObjects()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.ownedObjects
    }
}
// Gets the passwordPolicies property value. Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two may be specified together; for example: DisablePasswordExpiration, DisableStrongPassword. Returned only on $select. For more information on the default password policies, see Azure AD pasword policies. Supports $filter (ne, NOT).
func (m *User) GetPasswordPolicies()(*string) {
    if m == nil {
        return nil
    } else {
        return m.passwordPolicies
    }
}
// Gets the passwordProfile property value. Specifies the password profile for the user. The profile contains the user’s password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required. NOTE: For Azure B2C tenants, the forceChangePasswordNextSignIn property should be set to false and instead use custom policies and user flows to force password reset at first logon. See Force password reset at first logon.Returned only on $select. Supports $filter (eq, ne, NOT, in).
func (m *User) GetPasswordProfile()(*PasswordProfile) {
    if m == nil {
        return nil
    } else {
        return m.passwordProfile
    }
}
// Gets the pastProjects property value. A list for the user to enumerate their past projects. Returned only on $select.
func (m *User) GetPastProjects()([]string) {
    if m == nil {
        return nil
    } else {
        return m.pastProjects
    }
}
// Gets the people property value. People that are relevant to the user. Read-only. Nullable.
func (m *User) GetPeople()([]Person) {
    if m == nil {
        return nil
    } else {
        return m.people
    }
}
// Gets the photo property value. The user's profile photo. Read-only.
func (m *User) GetPhoto()(*ProfilePhoto) {
    if m == nil {
        return nil
    } else {
        return m.photo
    }
}
// Gets the photos property value. Read-only. Nullable.
func (m *User) GetPhotos()([]ProfilePhoto) {
    if m == nil {
        return nil
    } else {
        return m.photos
    }
}
// Gets the planner property value. Entry-point to the Planner resource that might exist for a user. Read-only.
func (m *User) GetPlanner()(*PlannerUser) {
    if m == nil {
        return nil
    } else {
        return m.planner
    }
}
// Gets the postalCode property value. The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code. Maximum length is 40 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *User) GetPostalCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postalCode
    }
}
// Gets the preferredLanguage property value. The preferred language for the user. Should follow ISO 639-1 Code; for example en-US. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith)
func (m *User) GetPreferredLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredLanguage
    }
}
// Gets the preferredName property value. The preferred name for the user. Returned only on $select.
func (m *User) GetPreferredName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredName
    }
}
// Gets the presence property value. 
func (m *User) GetPresence()(*Presence) {
    if m == nil {
        return nil
    } else {
        return m.presence
    }
}
// Gets the provisionedPlans property value. The plans that are provisioned for the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq, NOT, ge, le).
func (m *User) GetProvisionedPlans()([]ProvisionedPlan) {
    if m == nil {
        return nil
    } else {
        return m.provisionedPlans
    }
}
// Gets the proxyAddresses property value. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. For Azure AD B2C accounts, this property has a limit of ten unique addresses. Read-only, Not nullable. Returned only on $select. Supports $filter (eq, NOT, ge, le, startsWith).
func (m *User) GetProxyAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.proxyAddresses
    }
}
// Gets the registeredDevices property value. Devices that are registered for the user. Read-only. Nullable. Supports $expand.
func (m *User) GetRegisteredDevices()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.registeredDevices
    }
}
// Gets the responsibilities property value. A list for the user to enumerate their responsibilities. Returned only on $select.
func (m *User) GetResponsibilities()([]string) {
    if m == nil {
        return nil
    } else {
        return m.responsibilities
    }
}
// Gets the schools property value. A list for the user to enumerate the schools they have attended. Returned only on $select.
func (m *User) GetSchools()([]string) {
    if m == nil {
        return nil
    } else {
        return m.schools
    }
}
// Gets the scopedRoleMemberOf property value. The scoped-role administrative unit memberships for this user. Read-only. Nullable.
func (m *User) GetScopedRoleMemberOf()([]ScopedRoleMembership) {
    if m == nil {
        return nil
    } else {
        return m.scopedRoleMemberOf
    }
}
// Gets the settings property value. Read-only. Nullable.
func (m *User) GetSettings()(*UserSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// Gets the showInAddressList property value. true if the Outlook global address list should contain this user, otherwise false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false. Returned only on $select. Supports $filter (eq, ne, NOT, in).
func (m *User) GetShowInAddressList()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showInAddressList
    }
}
// Gets the signInSessionsValidFromDateTime property value. Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).  If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint. Read-only. Use revokeSignInSessions to reset. Returned only on $select.
func (m *User) GetSignInSessionsValidFromDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.signInSessionsValidFromDateTime
    }
}
// Gets the skills property value. A list for the user to enumerate their skills. Returned only on $select.
func (m *User) GetSkills()([]string) {
    if m == nil {
        return nil
    } else {
        return m.skills
    }
}
// Gets the state property value. The state or province in the user's address. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *User) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the streetAddress property value. The street address of the user's place of business. Maximum length is 1024 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *User) GetStreetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.streetAddress
    }
}
// Gets the surname property value. The user's surname (family name or last name). Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *User) GetSurname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.surname
    }
}
// Gets the teamwork property value. A container for Microsoft Teams features available for the user. Read-only. Nullable.
func (m *User) GetTeamwork()(*UserTeamwork) {
    if m == nil {
        return nil
    } else {
        return m.teamwork
    }
}
// Gets the todo property value. Represents the To Do services available to a user.
func (m *User) GetTodo()(*Todo) {
    if m == nil {
        return nil
    } else {
        return m.todo
    }
}
// Gets the transitiveMemberOf property value. 
func (m *User) GetTransitiveMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.transitiveMemberOf
    }
}
// Gets the usageLocation property value. A two letter country code (ISO standard 3166). Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.  Examples include: US, JP, and GB. Not nullable. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *User) GetUsageLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.usageLocation
    }
}
// Gets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization.NOTE: This property cannot contain accent characters. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, endsWith) and $orderBy.
func (m *User) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Gets the userType property value. A string value that can be used to classify user types in your directory, such as Member and Guest. Returned only on $select. Supports $filter (eq, ne, NOT, in). NOTE: For more information about the permissions for member and guest users, see What are the default user permissions in Azure Active Directory?
func (m *User) GetUserType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userType
    }
}
// The deserialization information for the current model
func (m *User) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["aboutMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAboutMe(val)
        return nil
    }
    res["accountEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAccountEnabled(val)
        return nil
    }
    res["activities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserActivity() })
        if err != nil {
            return err
        }
        res := make([]UserActivity, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserActivity))
        }
        m.SetActivities(res)
        return nil
    }
    res["ageGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAgeGroup(val)
        return nil
    }
    res["agreementAcceptances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreementAcceptance() })
        if err != nil {
            return err
        }
        res := make([]AgreementAcceptance, len(val))
        for i, v := range val {
            res[i] = *(v.(*AgreementAcceptance))
        }
        m.SetAgreementAcceptances(res)
        return nil
    }
    res["appRoleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppRoleAssignment() })
        if err != nil {
            return err
        }
        res := make([]AppRoleAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*AppRoleAssignment))
        }
        m.SetAppRoleAssignments(res)
        return nil
    }
    res["assignedLicenses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignedLicense() })
        if err != nil {
            return err
        }
        res := make([]AssignedLicense, len(val))
        for i, v := range val {
            res[i] = *(v.(*AssignedLicense))
        }
        m.SetAssignedLicenses(res)
        return nil
    }
    res["assignedPlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignedPlan() })
        if err != nil {
            return err
        }
        res := make([]AssignedPlan, len(val))
        for i, v := range val {
            res[i] = *(v.(*AssignedPlan))
        }
        m.SetAssignedPlans(res)
        return nil
    }
    res["authentication"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthentication() })
        if err != nil {
            return err
        }
        m.SetAuthentication(val.(*Authentication))
        return nil
    }
    res["birthday"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetBirthday(val)
        return nil
    }
    res["businessPhones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetBusinessPhones(res)
        return nil
    }
    res["calendar"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCalendar() })
        if err != nil {
            return err
        }
        m.SetCalendar(val.(*Calendar))
        return nil
    }
    res["calendarGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCalendarGroup() })
        if err != nil {
            return err
        }
        res := make([]CalendarGroup, len(val))
        for i, v := range val {
            res[i] = *(v.(*CalendarGroup))
        }
        m.SetCalendarGroups(res)
        return nil
    }
    res["calendars"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCalendar() })
        if err != nil {
            return err
        }
        res := make([]Calendar, len(val))
        for i, v := range val {
            res[i] = *(v.(*Calendar))
        }
        m.SetCalendars(res)
        return nil
    }
    res["calendarView"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvent() })
        if err != nil {
            return err
        }
        res := make([]Event, len(val))
        for i, v := range val {
            res[i] = *(v.(*Event))
        }
        m.SetCalendarView(res)
        return nil
    }
    res["chats"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChat() })
        if err != nil {
            return err
        }
        res := make([]Chat, len(val))
        for i, v := range val {
            res[i] = *(v.(*Chat))
        }
        m.SetChats(res)
        return nil
    }
    res["city"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCity(val)
        return nil
    }
    res["companyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCompanyName(val)
        return nil
    }
    res["consentProvidedForMinor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetConsentProvidedForMinor(val)
        return nil
    }
    res["contactFolders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContactFolder() })
        if err != nil {
            return err
        }
        res := make([]ContactFolder, len(val))
        for i, v := range val {
            res[i] = *(v.(*ContactFolder))
        }
        m.SetContactFolders(res)
        return nil
    }
    res["contacts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContact() })
        if err != nil {
            return err
        }
        res := make([]Contact, len(val))
        for i, v := range val {
            res[i] = *(v.(*Contact))
        }
        m.SetContacts(res)
        return nil
    }
    res["country"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCountry(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["createdObjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetCreatedObjects(res)
        return nil
    }
    res["creationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCreationType(val)
        return nil
    }
    res["department"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDepartment(val)
        return nil
    }
    res["deviceEnrollmentLimit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDeviceEnrollmentLimit(val)
        return nil
    }
    res["deviceManagementTroubleshootingEvents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementTroubleshootingEvent() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementTroubleshootingEvent, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementTroubleshootingEvent))
        }
        m.SetDeviceManagementTroubleshootingEvents(res)
        return nil
    }
    res["directReports"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetDirectReports(res)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["drive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDrive() })
        if err != nil {
            return err
        }
        m.SetDrive(val.(*Drive))
        return nil
    }
    res["drives"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDrive() })
        if err != nil {
            return err
        }
        res := make([]Drive, len(val))
        for i, v := range val {
            res[i] = *(v.(*Drive))
        }
        m.SetDrives(res)
        return nil
    }
    res["employeeHireDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEmployeeHireDate(val)
        return nil
    }
    res["employeeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmployeeId(val)
        return nil
    }
    res["employeeOrgData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmployeeOrgData() })
        if err != nil {
            return err
        }
        m.SetEmployeeOrgData(val.(*EmployeeOrgData))
        return nil
    }
    res["employeeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmployeeType(val)
        return nil
    }
    res["events"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvent() })
        if err != nil {
            return err
        }
        res := make([]Event, len(val))
        for i, v := range val {
            res[i] = *(v.(*Event))
        }
        m.SetEvents(res)
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtension() })
        if err != nil {
            return err
        }
        res := make([]Extension, len(val))
        for i, v := range val {
            res[i] = *(v.(*Extension))
        }
        m.SetExtensions(res)
        return nil
    }
    res["externalUserState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalUserState(val)
        return nil
    }
    res["externalUserStateChangeDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExternalUserStateChangeDateTime(val)
        return nil
    }
    res["faxNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFaxNumber(val)
        return nil
    }
    res["followedSites"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSite() })
        if err != nil {
            return err
        }
        res := make([]Site, len(val))
        for i, v := range val {
            res[i] = *(v.(*Site))
        }
        m.SetFollowedSites(res)
        return nil
    }
    res["givenName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGivenName(val)
        return nil
    }
    res["hireDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetHireDate(val)
        return nil
    }
    res["identities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewObjectIdentity() })
        if err != nil {
            return err
        }
        res := make([]ObjectIdentity, len(val))
        for i, v := range val {
            res[i] = *(v.(*ObjectIdentity))
        }
        m.SetIdentities(res)
        return nil
    }
    res["imAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetImAddresses(res)
        return nil
    }
    res["inferenceClassification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInferenceClassification() })
        if err != nil {
            return err
        }
        m.SetInferenceClassification(val.(*InferenceClassification))
        return nil
    }
    res["insights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfficeGraphInsights() })
        if err != nil {
            return err
        }
        m.SetInsights(val.(*OfficeGraphInsights))
        return nil
    }
    res["interests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetInterests(res)
        return nil
    }
    res["isResourceAccount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsResourceAccount(val)
        return nil
    }
    res["jobTitle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJobTitle(val)
        return nil
    }
    res["joinedTeams"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeam() })
        if err != nil {
            return err
        }
        res := make([]Team, len(val))
        for i, v := range val {
            res[i] = *(v.(*Team))
        }
        m.SetJoinedTeams(res)
        return nil
    }
    res["lastPasswordChangeDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastPasswordChangeDateTime(val)
        return nil
    }
    res["legalAgeGroupClassification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLegalAgeGroupClassification(val)
        return nil
    }
    res["licenseAssignmentStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLicenseAssignmentState() })
        if err != nil {
            return err
        }
        res := make([]LicenseAssignmentState, len(val))
        for i, v := range val {
            res[i] = *(v.(*LicenseAssignmentState))
        }
        m.SetLicenseAssignmentStates(res)
        return nil
    }
    res["licenseDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLicenseDetails() })
        if err != nil {
            return err
        }
        res := make([]LicenseDetails, len(val))
        for i, v := range val {
            res[i] = *(v.(*LicenseDetails))
        }
        m.SetLicenseDetails(res)
        return nil
    }
    res["mail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMail(val)
        return nil
    }
    res["mailboxSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMailboxSettings() })
        if err != nil {
            return err
        }
        m.SetMailboxSettings(val.(*MailboxSettings))
        return nil
    }
    res["mailFolders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMailFolder() })
        if err != nil {
            return err
        }
        res := make([]MailFolder, len(val))
        for i, v := range val {
            res[i] = *(v.(*MailFolder))
        }
        m.SetMailFolders(res)
        return nil
    }
    res["mailNickname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMailNickname(val)
        return nil
    }
    res["managedAppRegistrations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppRegistration() })
        if err != nil {
            return err
        }
        res := make([]ManagedAppRegistration, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedAppRegistration))
        }
        m.SetManagedAppRegistrations(res)
        return nil
    }
    res["managedDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDevice() })
        if err != nil {
            return err
        }
        res := make([]ManagedDevice, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedDevice))
        }
        m.SetManagedDevices(res)
        return nil
    }
    res["manager"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        m.SetManager(val.(*DirectoryObject))
        return nil
    }
    res["memberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetMemberOf(res)
        return nil
    }
    res["messages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMessage() })
        if err != nil {
            return err
        }
        res := make([]Message, len(val))
        for i, v := range val {
            res[i] = *(v.(*Message))
        }
        m.SetMessages(res)
        return nil
    }
    res["mobilePhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMobilePhone(val)
        return nil
    }
    res["mySite"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMySite(val)
        return nil
    }
    res["oauth2PermissionGrants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOAuth2PermissionGrant() })
        if err != nil {
            return err
        }
        res := make([]OAuth2PermissionGrant, len(val))
        for i, v := range val {
            res[i] = *(v.(*OAuth2PermissionGrant))
        }
        m.SetOauth2PermissionGrants(res)
        return nil
    }
    res["officeLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOfficeLocation(val)
        return nil
    }
    res["onenote"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnenote() })
        if err != nil {
            return err
        }
        m.SetOnenote(val.(*Onenote))
        return nil
    }
    res["onlineMeetings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnlineMeeting() })
        if err != nil {
            return err
        }
        res := make([]OnlineMeeting, len(val))
        for i, v := range val {
            res[i] = *(v.(*OnlineMeeting))
        }
        m.SetOnlineMeetings(res)
        return nil
    }
    res["onPremisesDistinguishedName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOnPremisesDistinguishedName(val)
        return nil
    }
    res["onPremisesDomainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOnPremisesDomainName(val)
        return nil
    }
    res["onPremisesExtensionAttributes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesExtensionAttributes() })
        if err != nil {
            return err
        }
        m.SetOnPremisesExtensionAttributes(val.(*OnPremisesExtensionAttributes))
        return nil
    }
    res["onPremisesImmutableId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOnPremisesImmutableId(val)
        return nil
    }
    res["onPremisesLastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetOnPremisesLastSyncDateTime(val)
        return nil
    }
    res["onPremisesProvisioningErrors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesProvisioningError() })
        if err != nil {
            return err
        }
        res := make([]OnPremisesProvisioningError, len(val))
        for i, v := range val {
            res[i] = *(v.(*OnPremisesProvisioningError))
        }
        m.SetOnPremisesProvisioningErrors(res)
        return nil
    }
    res["onPremisesSamAccountName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOnPremisesSamAccountName(val)
        return nil
    }
    res["onPremisesSecurityIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOnPremisesSecurityIdentifier(val)
        return nil
    }
    res["onPremisesSyncEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOnPremisesSyncEnabled(val)
        return nil
    }
    res["onPremisesUserPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOnPremisesUserPrincipalName(val)
        return nil
    }
    res["otherMails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetOtherMails(res)
        return nil
    }
    res["outlook"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutlookUser() })
        if err != nil {
            return err
        }
        m.SetOutlook(val.(*OutlookUser))
        return nil
    }
    res["ownedDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetOwnedDevices(res)
        return nil
    }
    res["ownedObjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetOwnedObjects(res)
        return nil
    }
    res["passwordPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPasswordPolicies(val)
        return nil
    }
    res["passwordProfile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPasswordProfile() })
        if err != nil {
            return err
        }
        m.SetPasswordProfile(val.(*PasswordProfile))
        return nil
    }
    res["pastProjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetPastProjects(res)
        return nil
    }
    res["people"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPerson() })
        if err != nil {
            return err
        }
        res := make([]Person, len(val))
        for i, v := range val {
            res[i] = *(v.(*Person))
        }
        m.SetPeople(res)
        return nil
    }
    res["photo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProfilePhoto() })
        if err != nil {
            return err
        }
        m.SetPhoto(val.(*ProfilePhoto))
        return nil
    }
    res["photos"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProfilePhoto() })
        if err != nil {
            return err
        }
        res := make([]ProfilePhoto, len(val))
        for i, v := range val {
            res[i] = *(v.(*ProfilePhoto))
        }
        m.SetPhotos(res)
        return nil
    }
    res["planner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerUser() })
        if err != nil {
            return err
        }
        m.SetPlanner(val.(*PlannerUser))
        return nil
    }
    res["postalCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPostalCode(val)
        return nil
    }
    res["preferredLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPreferredLanguage(val)
        return nil
    }
    res["preferredName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPreferredName(val)
        return nil
    }
    res["presence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPresence() })
        if err != nil {
            return err
        }
        m.SetPresence(val.(*Presence))
        return nil
    }
    res["provisionedPlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisionedPlan() })
        if err != nil {
            return err
        }
        res := make([]ProvisionedPlan, len(val))
        for i, v := range val {
            res[i] = *(v.(*ProvisionedPlan))
        }
        m.SetProvisionedPlans(res)
        return nil
    }
    res["proxyAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetProxyAddresses(res)
        return nil
    }
    res["registeredDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetRegisteredDevices(res)
        return nil
    }
    res["responsibilities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetResponsibilities(res)
        return nil
    }
    res["schools"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetSchools(res)
        return nil
    }
    res["scopedRoleMemberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewScopedRoleMembership() })
        if err != nil {
            return err
        }
        res := make([]ScopedRoleMembership, len(val))
        for i, v := range val {
            res[i] = *(v.(*ScopedRoleMembership))
        }
        m.SetScopedRoleMemberOf(res)
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSettings() })
        if err != nil {
            return err
        }
        m.SetSettings(val.(*UserSettings))
        return nil
    }
    res["showInAddressList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowInAddressList(val)
        return nil
    }
    res["signInSessionsValidFromDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetSignInSessionsValidFromDateTime(val)
        return nil
    }
    res["skills"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetSkills(res)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetState(val)
        return nil
    }
    res["streetAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStreetAddress(val)
        return nil
    }
    res["surname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSurname(val)
        return nil
    }
    res["teamwork"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserTeamwork() })
        if err != nil {
            return err
        }
        m.SetTeamwork(val.(*UserTeamwork))
        return nil
    }
    res["todo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTodo() })
        if err != nil {
            return err
        }
        m.SetTodo(val.(*Todo))
        return nil
    }
    res["transitiveMemberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetTransitiveMemberOf(res)
        return nil
    }
    res["usageLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUsageLocation(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    res["userType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserType(val)
        return nil
    }
    return res
}
func (m *User) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the aboutMe property value. A freeform text entry field for the user to describe themselves. Returned only on $select.
// Parameters:
//  - value : Value to set for the aboutMe property.
func (m *User) SetAboutMe(value *string)() {
    m.aboutMe = value
}
// Sets the accountEnabled property value. true if the account is enabled; otherwise, false. This property is required when a user is created. Returned only on $select. Supports $filter (eq, ne, NOT, and in).
// Parameters:
//  - value : Value to set for the accountEnabled property.
func (m *User) SetAccountEnabled(value *bool)() {
    m.accountEnabled = value
}
// Sets the activities property value. The user's activities across devices. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the activities property.
func (m *User) SetActivities(value []UserActivity)() {
    m.activities = value
}
// Sets the ageGroup property value. Sets the age group of the user. Allowed values: null, minor, notAdult and adult. Refer to the legal age group property definitions for further information. Returned only on $select. Supports $filter (eq, ne, NOT, and in).
// Parameters:
//  - value : Value to set for the ageGroup property.
func (m *User) SetAgeGroup(value *string)() {
    m.ageGroup = value
}
// Sets the agreementAcceptances property value. The user's terms of use acceptance statuses. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the agreementAcceptances property.
func (m *User) SetAgreementAcceptances(value []AgreementAcceptance)() {
    m.agreementAcceptances = value
}
// Sets the appRoleAssignments property value. Represents the app roles a user has been granted for an application. Supports $expand.
// Parameters:
//  - value : Value to set for the appRoleAssignments property.
func (m *User) SetAppRoleAssignments(value []AppRoleAssignment)() {
    m.appRoleAssignments = value
}
// Sets the assignedLicenses property value. The licenses that are assigned to the user, including inherited (group-based) licenses.  Not nullable. Returned only on $select. Supports $filter (eq and NOT).
// Parameters:
//  - value : Value to set for the assignedLicenses property.
func (m *User) SetAssignedLicenses(value []AssignedLicense)() {
    m.assignedLicenses = value
}
// Sets the assignedPlans property value. The plans that are assigned to the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq and NOT).
// Parameters:
//  - value : Value to set for the assignedPlans property.
func (m *User) SetAssignedPlans(value []AssignedPlan)() {
    m.assignedPlans = value
}
// Sets the authentication property value. 
// Parameters:
//  - value : Value to set for the authentication property.
func (m *User) SetAuthentication(value *Authentication)() {
    m.authentication = value
}
// Sets the birthday property value. The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.
// Parameters:
//  - value : Value to set for the birthday property.
func (m *User) SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.birthday = value
}
// Sets the businessPhones property value. The telephone numbers for the user. NOTE: Although this is a string collection, only one number can be set for this property. Read-only for users synced from on-premises directory. Returned by default. Supports $filter (eq and NOT).
// Parameters:
//  - value : Value to set for the businessPhones property.
func (m *User) SetBusinessPhones(value []string)() {
    m.businessPhones = value
}
// Sets the calendar property value. The user's primary calendar. Read-only.
// Parameters:
//  - value : Value to set for the calendar property.
func (m *User) SetCalendar(value *Calendar)() {
    m.calendar = value
}
// Sets the calendarGroups property value. The user's calendar groups. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the calendarGroups property.
func (m *User) SetCalendarGroups(value []CalendarGroup)() {
    m.calendarGroups = value
}
// Sets the calendars property value. The user's calendars. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the calendars property.
func (m *User) SetCalendars(value []Calendar)() {
    m.calendars = value
}
// Sets the calendarView property value. The calendar view for the calendar. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the calendarView property.
func (m *User) SetCalendarView(value []Event)() {
    m.calendarView = value
}
// Sets the chats property value. 
// Parameters:
//  - value : Value to set for the chats property.
func (m *User) SetChats(value []Chat)() {
    m.chats = value
}
// Sets the city property value. The city in which the user is located. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the city property.
func (m *User) SetCity(value *string)() {
    m.city = value
}
// Sets the companyName property value. The company name which the user is associated. This property can be useful for describing the company that an external user comes from. The maximum length of the company name is 64 characters.Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the companyName property.
func (m *User) SetCompanyName(value *string)() {
    m.companyName = value
}
// Sets the consentProvidedForMinor property value. Sets whether consent has been obtained for minors. Allowed values: null, granted, denied and notRequired. Refer to the legal age group property definitions for further information. Returned only on $select. Supports $filter (eq, ne, NOT, and in).
// Parameters:
//  - value : Value to set for the consentProvidedForMinor property.
func (m *User) SetConsentProvidedForMinor(value *string)() {
    m.consentProvidedForMinor = value
}
// Sets the contactFolders property value. The user's contacts folders. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the contactFolders property.
func (m *User) SetContactFolders(value []ContactFolder)() {
    m.contactFolders = value
}
// Sets the contacts property value. The user's contacts. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the contacts property.
func (m *User) SetContacts(value []Contact)() {
    m.contacts = value
}
// Sets the country property value. The country/region in which the user is located; for example, US or UK. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the country property.
func (m *User) SetCountry(value *string)() {
    m.country = value
}
// Sets the createdDateTime property value. The created date of the user object. Read-only. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, and in operators).
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *User) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the createdObjects property value. Directory objects that were created by the user. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the createdObjects property.
func (m *User) SetCreatedObjects(value []DirectoryObject)() {
    m.createdObjects = value
}
// Sets the creationType property value. Indicates whether the user account was created through one of the following methods:  As a regular school or work account (null). As an external account (Invitation). As a local account for an Azure Active Directory B2C tenant (LocalAccount). Through self-service sign-up by an internal user using email verification (EmailVerified). Through self-service sign-up by an external user signing up through a link that is part of a user flow (SelfServiceSignUp). Read-only.Returned only on $select. Supports $filter (eq, ne, NOT, and in).
// Parameters:
//  - value : Value to set for the creationType property.
func (m *User) SetCreationType(value *string)() {
    m.creationType = value
}
// Sets the department property value. The name for the department in which the user works. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, and in operators).
// Parameters:
//  - value : Value to set for the department property.
func (m *User) SetDepartment(value *string)() {
    m.department = value
}
// Sets the deviceEnrollmentLimit property value. The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000.
// Parameters:
//  - value : Value to set for the deviceEnrollmentLimit property.
func (m *User) SetDeviceEnrollmentLimit(value *int32)() {
    m.deviceEnrollmentLimit = value
}
// Sets the deviceManagementTroubleshootingEvents property value. The list of troubleshooting events for this user.
// Parameters:
//  - value : Value to set for the deviceManagementTroubleshootingEvents property.
func (m *User) SetDeviceManagementTroubleshootingEvents(value []DeviceManagementTroubleshootingEvent)() {
    m.deviceManagementTroubleshootingEvents = value
}
// Sets the directReports property value. The users and contacts that report to the user. (The users and contacts that have their manager property set to this user.) Read-only. Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the directReports property.
func (m *User) SetDirectReports(value []DirectoryObject)() {
    m.directReports = value
}
// Sets the displayName property value. The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial and last name. This property is required when a user is created and it cannot be cleared during updates. Maximum length is 256 characters. Returned by default. Supports $filter (eq, ne, NOT , ge, le, in, startsWith), $orderBy, and $search.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *User) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the drive property value. The user's OneDrive. Read-only.
// Parameters:
//  - value : Value to set for the drive property.
func (m *User) SetDrive(value *Drive)() {
    m.drive = value
}
// Sets the drives property value. A collection of drives available for this user. Read-only.
// Parameters:
//  - value : Value to set for the drives property.
func (m *User) SetDrives(value []Drive)() {
    m.drives = value
}
// Sets the employeeHireDate property value. The date and time when the user was hired or will start work in case of a future hire. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, in).
// Parameters:
//  - value : Value to set for the employeeHireDate property.
func (m *User) SetEmployeeHireDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.employeeHireDate = value
}
// Sets the employeeId property value. The employee identifier assigned to the user by the organization. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the employeeId property.
func (m *User) SetEmployeeId(value *string)() {
    m.employeeId = value
}
// Sets the employeeOrgData property value. Represents organization data (e.g. division and costCenter) associated with a user. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, in).
// Parameters:
//  - value : Value to set for the employeeOrgData property.
func (m *User) SetEmployeeOrgData(value *EmployeeOrgData)() {
    m.employeeOrgData = value
}
// Sets the employeeType property value. Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the employeeType property.
func (m *User) SetEmployeeType(value *string)() {
    m.employeeType = value
}
// Sets the events property value. The user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the events property.
func (m *User) SetEvents(value []Event)() {
    m.events = value
}
// Sets the extensions property value. The collection of open extensions defined for the user. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the extensions property.
func (m *User) SetExtensions(value []Extension)() {
    m.extensions = value
}
// Sets the externalUserState property value. For an external user invited to the tenant using the invitation API, this property represents the invited user's invitation status. For invited users, the state can be PendingAcceptance or Accepted, or null for all other users. Returned only on $select. Supports $filter (eq, ne, NOT , in).
// Parameters:
//  - value : Value to set for the externalUserState property.
func (m *User) SetExternalUserState(value *string)() {
    m.externalUserState = value
}
// Sets the externalUserStateChangeDateTime property value. Shows the timestamp for the latest change to the externalUserState property. Returned only on $select. Supports $filter (eq, ne, NOT , in).
// Parameters:
//  - value : Value to set for the externalUserStateChangeDateTime property.
func (m *User) SetExternalUserStateChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.externalUserStateChangeDateTime = value
}
// Sets the faxNumber property value. The fax number of the user. Returned only on $select. Supports $filter (eq, ne, NOT , ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the faxNumber property.
func (m *User) SetFaxNumber(value *string)() {
    m.faxNumber = value
}
// Sets the followedSites property value. 
// Parameters:
//  - value : Value to set for the followedSites property.
func (m *User) SetFollowedSites(value []Site)() {
    m.followedSites = value
}
// Sets the givenName property value. The given name (first name) of the user. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, NOT , ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the givenName property.
func (m *User) SetGivenName(value *string)() {
    m.givenName = value
}
// Sets the hireDate property value. The hire date of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.  Note: This property is specific to SharePoint Online. We recommend using the native employeeHireDate property to set and update hire date values using Microsoft Graph APIs.
// Parameters:
//  - value : Value to set for the hireDate property.
func (m *User) SetHireDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.hireDate = value
}
// Sets the identities property value. Represents the identities that can be used to sign in to this user account. An identity can be provided by Microsoft (also known as a local account), by organizations, or by social identity providers such as Facebook, Google, and Microsoft, and tied to a user account. May contain multiple items with the same signInType value. Returned only on $select. Supports $filter (eq) only where the signInType is not userPrincipalName.
// Parameters:
//  - value : Value to set for the identities property.
func (m *User) SetIdentities(value []ObjectIdentity)() {
    m.identities = value
}
// Sets the imAddresses property value. The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user. Read-only. Returned only on $select. Supports $filter (eq, NOT, ge, le, startsWith).
// Parameters:
//  - value : Value to set for the imAddresses property.
func (m *User) SetImAddresses(value []string)() {
    m.imAddresses = value
}
// Sets the inferenceClassification property value. Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
// Parameters:
//  - value : Value to set for the inferenceClassification property.
func (m *User) SetInferenceClassification(value *InferenceClassification)() {
    m.inferenceClassification = value
}
// Sets the insights property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the insights property.
func (m *User) SetInsights(value *OfficeGraphInsights)() {
    m.insights = value
}
// Sets the interests property value. A list for the user to describe their interests. Returned only on $select.
// Parameters:
//  - value : Value to set for the interests property.
func (m *User) SetInterests(value []string)() {
    m.interests = value
}
// Sets the isResourceAccount property value. Do not use – reserved for future use.
// Parameters:
//  - value : Value to set for the isResourceAccount property.
func (m *User) SetIsResourceAccount(value *bool)() {
    m.isResourceAccount = value
}
// Sets the jobTitle property value. The user's job title. Maximum length is 128 characters. Returned by default. Supports $filter (eq, ne, NOT , ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the jobTitle property.
func (m *User) SetJobTitle(value *string)() {
    m.jobTitle = value
}
// Sets the joinedTeams property value. The Microsoft Teams teams that the user is a member of. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the joinedTeams property.
func (m *User) SetJoinedTeams(value []Team)() {
    m.joinedTeams = value
}
// Sets the lastPasswordChangeDateTime property value. The time when this Azure AD user last changed their password or when their password was created, whichever date the latest action was performed. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.
// Parameters:
//  - value : Value to set for the lastPasswordChangeDateTime property.
func (m *User) SetLastPasswordChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastPasswordChangeDateTime = value
}
// Sets the legalAgeGroupClassification property value. Used by enterprise applications to determine the legal age group of the user. This property is read-only and calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, minorWithOutParentalConsent, minorWithParentalConsent, minorNoParentalConsentRequired, notAdult and adult. Refer to the legal age group property definitions for further information. Returned only on $select.
// Parameters:
//  - value : Value to set for the legalAgeGroupClassification property.
func (m *User) SetLegalAgeGroupClassification(value *string)() {
    m.legalAgeGroupClassification = value
}
// Sets the licenseAssignmentStates property value. State of license assignments for this user. Read-only. Returned only on $select.
// Parameters:
//  - value : Value to set for the licenseAssignmentStates property.
func (m *User) SetLicenseAssignmentStates(value []LicenseAssignmentState)() {
    m.licenseAssignmentStates = value
}
// Sets the licenseDetails property value. A collection of this user's license details. Read-only.
// Parameters:
//  - value : Value to set for the licenseDetails property.
func (m *User) SetLicenseDetails(value []LicenseDetails)() {
    m.licenseDetails = value
}
// Sets the mail property value. The SMTP address for the user, for example, jeff@contoso.onmicrosoft.com.Changes to this property will also update the user's proxyAddresses collection to include the value as an SMTP address. For Azure AD B2C accounts, this property can be updated up to only ten times with unique SMTP addresses. This property cannot contain accent characters.Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, endsWith).
// Parameters:
//  - value : Value to set for the mail property.
func (m *User) SetMail(value *string)() {
    m.mail = value
}
// Sets the mailboxSettings property value. Settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale and time zone.Returned only on $select.
// Parameters:
//  - value : Value to set for the mailboxSettings property.
func (m *User) SetMailboxSettings(value *MailboxSettings)() {
    m.mailboxSettings = value
}
// Sets the mailFolders property value. The user's mail folders. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the mailFolders property.
func (m *User) SetMailFolders(value []MailFolder)() {
    m.mailFolders = value
}
// Sets the mailNickname property value. The mail alias for the user. This property must be specified when a user is created. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the mailNickname property.
func (m *User) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// Sets the managedAppRegistrations property value. Zero or more managed app registrations that belong to the user.
// Parameters:
//  - value : Value to set for the managedAppRegistrations property.
func (m *User) SetManagedAppRegistrations(value []ManagedAppRegistration)() {
    m.managedAppRegistrations = value
}
// Sets the managedDevices property value. The managed devices associated with the user.
// Parameters:
//  - value : Value to set for the managedDevices property.
func (m *User) SetManagedDevices(value []ManagedDevice)() {
    m.managedDevices = value
}
// Sets the manager property value. The user or contact that is this user's manager. Read-only. (HTTP Methods: GET, PUT, DELETE.). Supports $expand.
// Parameters:
//  - value : Value to set for the manager property.
func (m *User) SetManager(value *DirectoryObject)() {
    m.manager = value
}
// Sets the memberOf property value. The groups and directory roles that the user is a member of. Read-only. Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the memberOf property.
func (m *User) SetMemberOf(value []DirectoryObject)() {
    m.memberOf = value
}
// Sets the messages property value. The messages in a mailbox or folder. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the messages property.
func (m *User) SetMessages(value []Message)() {
    m.messages = value
}
// Sets the mobilePhone property value. The primary cellular telephone number for the user. Read-only for users synced from on-premises directory. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the mobilePhone property.
func (m *User) SetMobilePhone(value *string)() {
    m.mobilePhone = value
}
// Sets the mySite property value. The URL for the user's personal site. Returned only on $select.
// Parameters:
//  - value : Value to set for the mySite property.
func (m *User) SetMySite(value *string)() {
    m.mySite = value
}
// Sets the oauth2PermissionGrants property value. 
// Parameters:
//  - value : Value to set for the oauth2PermissionGrants property.
func (m *User) SetOauth2PermissionGrants(value []OAuth2PermissionGrant)() {
    m.oauth2PermissionGrants = value
}
// Sets the officeLocation property value. The office location in the user's place of business. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the officeLocation property.
func (m *User) SetOfficeLocation(value *string)() {
    m.officeLocation = value
}
// Sets the onenote property value. Read-only.
// Parameters:
//  - value : Value to set for the onenote property.
func (m *User) SetOnenote(value *Onenote)() {
    m.onenote = value
}
// Sets the onlineMeetings property value. 
// Parameters:
//  - value : Value to set for the onlineMeetings property.
func (m *User) SetOnlineMeetings(value []OnlineMeeting)() {
    m.onlineMeetings = value
}
// Sets the onPremisesDistinguishedName property value. Contains the on-premises Active Directory distinguished name or DN. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select.
// Parameters:
//  - value : Value to set for the onPremisesDistinguishedName property.
func (m *User) SetOnPremisesDistinguishedName(value *string)() {
    m.onPremisesDistinguishedName = value
}
// Sets the onPremisesDomainName property value. Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select.
// Parameters:
//  - value : Value to set for the onPremisesDomainName property.
func (m *User) SetOnPremisesDomainName(value *string)() {
    m.onPremisesDomainName = value
}
// Sets the onPremisesExtensionAttributes property value. Contains extensionAttributes 1-15 for the user. Note that the individual extension attributes are neither selectable nor filterable. For an onPremisesSyncEnabled user, the source of authority for this set of properties is the on-premises and is read-only. For a cloud-only user (where onPremisesSyncEnabled is false), these properties may be set during creation or update. These extension attributes are also known as Exchange custom attributes 1-15. Returned only on $select. Supports $filter (eq, NOT, ge, le, in).
// Parameters:
//  - value : Value to set for the onPremisesExtensionAttributes property.
func (m *User) SetOnPremisesExtensionAttributes(value *OnPremisesExtensionAttributes)() {
    m.onPremisesExtensionAttributes = value
}
// Sets the onPremisesImmutableId property value. This property is used to associate an on-premises Active Directory user account to their Azure AD user object. This property must be specified when creating a new user account in the Graph if you are using a federated domain for the user's userPrincipalName (UPN) property. NOTE: The $ and _ characters cannot be used when specifying this property. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in)..
// Parameters:
//  - value : Value to set for the onPremisesImmutableId property.
func (m *User) SetOnPremisesImmutableId(value *string)() {
    m.onPremisesImmutableId = value
}
// Sets the onPremisesLastSyncDateTime property value. Indicates the last time at which the object was synced with the on-premises directory; for example: 2013-02-16T03:04:54Z. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in).
// Parameters:
//  - value : Value to set for the onPremisesLastSyncDateTime property.
func (m *User) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onPremisesLastSyncDateTime = value
}
// Sets the onPremisesProvisioningErrors property value. Errors when using Microsoft synchronization product during provisioning. Returned only on $select. Supports $filter (eq, NOT, ge, le).
// Parameters:
//  - value : Value to set for the onPremisesProvisioningErrors property.
func (m *User) SetOnPremisesProvisioningErrors(value []OnPremisesProvisioningError)() {
    m.onPremisesProvisioningErrors = value
}
// Sets the onPremisesSamAccountName property value. Contains the on-premises samAccountName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the onPremisesSamAccountName property.
func (m *User) SetOnPremisesSamAccountName(value *string)() {
    m.onPremisesSamAccountName = value
}
// Sets the onPremisesSecurityIdentifier property value. Contains the on-premises security identifier (SID) for the user that was synchronized from on-premises to the cloud. Read-only. Returned only on $select.
// Parameters:
//  - value : Value to set for the onPremisesSecurityIdentifier property.
func (m *User) SetOnPremisesSecurityIdentifier(value *string)() {
    m.onPremisesSecurityIdentifier = value
}
// Sets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Returned only on $select. Supports $filter (eq, ne, NOT, in).
// Parameters:
//  - value : Value to set for the onPremisesSyncEnabled property.
func (m *User) SetOnPremisesSyncEnabled(value *bool)() {
    m.onPremisesSyncEnabled = value
}
// Sets the onPremisesUserPrincipalName property value. Contains the on-premises userPrincipalName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the onPremisesUserPrincipalName property.
func (m *User) SetOnPremisesUserPrincipalName(value *string)() {
    m.onPremisesUserPrincipalName = value
}
// Sets the otherMails property value. A list of additional email addresses for the user; for example: ['bob@contoso.com', 'Robert@fabrikam.com']. NOTE: This property cannot contain accent characters. Returned only on $select. Supports $filter (eq, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the otherMails property.
func (m *User) SetOtherMails(value []string)() {
    m.otherMails = value
}
// Sets the outlook property value. Read-only.
// Parameters:
//  - value : Value to set for the outlook property.
func (m *User) SetOutlook(value *OutlookUser)() {
    m.outlook = value
}
// Sets the ownedDevices property value. Devices that are owned by the user. Read-only. Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the ownedDevices property.
func (m *User) SetOwnedDevices(value []DirectoryObject)() {
    m.ownedDevices = value
}
// Sets the ownedObjects property value. Directory objects that are owned by the user. Read-only. Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the ownedObjects property.
func (m *User) SetOwnedObjects(value []DirectoryObject)() {
    m.ownedObjects = value
}
// Sets the passwordPolicies property value. Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two may be specified together; for example: DisablePasswordExpiration, DisableStrongPassword. Returned only on $select. For more information on the default password policies, see Azure AD pasword policies. Supports $filter (ne, NOT).
// Parameters:
//  - value : Value to set for the passwordPolicies property.
func (m *User) SetPasswordPolicies(value *string)() {
    m.passwordPolicies = value
}
// Sets the passwordProfile property value. Specifies the password profile for the user. The profile contains the user’s password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required. NOTE: For Azure B2C tenants, the forceChangePasswordNextSignIn property should be set to false and instead use custom policies and user flows to force password reset at first logon. See Force password reset at first logon.Returned only on $select. Supports $filter (eq, ne, NOT, in).
// Parameters:
//  - value : Value to set for the passwordProfile property.
func (m *User) SetPasswordProfile(value *PasswordProfile)() {
    m.passwordProfile = value
}
// Sets the pastProjects property value. A list for the user to enumerate their past projects. Returned only on $select.
// Parameters:
//  - value : Value to set for the pastProjects property.
func (m *User) SetPastProjects(value []string)() {
    m.pastProjects = value
}
// Sets the people property value. People that are relevant to the user. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the people property.
func (m *User) SetPeople(value []Person)() {
    m.people = value
}
// Sets the photo property value. The user's profile photo. Read-only.
// Parameters:
//  - value : Value to set for the photo property.
func (m *User) SetPhoto(value *ProfilePhoto)() {
    m.photo = value
}
// Sets the photos property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the photos property.
func (m *User) SetPhotos(value []ProfilePhoto)() {
    m.photos = value
}
// Sets the planner property value. Entry-point to the Planner resource that might exist for a user. Read-only.
// Parameters:
//  - value : Value to set for the planner property.
func (m *User) SetPlanner(value *PlannerUser)() {
    m.planner = value
}
// Sets the postalCode property value. The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code. Maximum length is 40 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the postalCode property.
func (m *User) SetPostalCode(value *string)() {
    m.postalCode = value
}
// Sets the preferredLanguage property value. The preferred language for the user. Should follow ISO 639-1 Code; for example en-US. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith)
// Parameters:
//  - value : Value to set for the preferredLanguage property.
func (m *User) SetPreferredLanguage(value *string)() {
    m.preferredLanguage = value
}
// Sets the preferredName property value. The preferred name for the user. Returned only on $select.
// Parameters:
//  - value : Value to set for the preferredName property.
func (m *User) SetPreferredName(value *string)() {
    m.preferredName = value
}
// Sets the presence property value. 
// Parameters:
//  - value : Value to set for the presence property.
func (m *User) SetPresence(value *Presence)() {
    m.presence = value
}
// Sets the provisionedPlans property value. The plans that are provisioned for the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq, NOT, ge, le).
// Parameters:
//  - value : Value to set for the provisionedPlans property.
func (m *User) SetProvisionedPlans(value []ProvisionedPlan)() {
    m.provisionedPlans = value
}
// Sets the proxyAddresses property value. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. For Azure AD B2C accounts, this property has a limit of ten unique addresses. Read-only, Not nullable. Returned only on $select. Supports $filter (eq, NOT, ge, le, startsWith).
// Parameters:
//  - value : Value to set for the proxyAddresses property.
func (m *User) SetProxyAddresses(value []string)() {
    m.proxyAddresses = value
}
// Sets the registeredDevices property value. Devices that are registered for the user. Read-only. Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the registeredDevices property.
func (m *User) SetRegisteredDevices(value []DirectoryObject)() {
    m.registeredDevices = value
}
// Sets the responsibilities property value. A list for the user to enumerate their responsibilities. Returned only on $select.
// Parameters:
//  - value : Value to set for the responsibilities property.
func (m *User) SetResponsibilities(value []string)() {
    m.responsibilities = value
}
// Sets the schools property value. A list for the user to enumerate the schools they have attended. Returned only on $select.
// Parameters:
//  - value : Value to set for the schools property.
func (m *User) SetSchools(value []string)() {
    m.schools = value
}
// Sets the scopedRoleMemberOf property value. The scoped-role administrative unit memberships for this user. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the scopedRoleMemberOf property.
func (m *User) SetScopedRoleMemberOf(value []ScopedRoleMembership)() {
    m.scopedRoleMemberOf = value
}
// Sets the settings property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the settings property.
func (m *User) SetSettings(value *UserSettings)() {
    m.settings = value
}
// Sets the showInAddressList property value. true if the Outlook global address list should contain this user, otherwise false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false. Returned only on $select. Supports $filter (eq, ne, NOT, in).
// Parameters:
//  - value : Value to set for the showInAddressList property.
func (m *User) SetShowInAddressList(value *bool)() {
    m.showInAddressList = value
}
// Sets the signInSessionsValidFromDateTime property value. Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).  If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint. Read-only. Use revokeSignInSessions to reset. Returned only on $select.
// Parameters:
//  - value : Value to set for the signInSessionsValidFromDateTime property.
func (m *User) SetSignInSessionsValidFromDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.signInSessionsValidFromDateTime = value
}
// Sets the skills property value. A list for the user to enumerate their skills. Returned only on $select.
// Parameters:
//  - value : Value to set for the skills property.
func (m *User) SetSkills(value []string)() {
    m.skills = value
}
// Sets the state property value. The state or province in the user's address. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the state property.
func (m *User) SetState(value *string)() {
    m.state = value
}
// Sets the streetAddress property value. The street address of the user's place of business. Maximum length is 1024 characters. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the streetAddress property.
func (m *User) SetStreetAddress(value *string)() {
    m.streetAddress = value
}
// Sets the surname property value. The user's surname (family name or last name). Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the surname property.
func (m *User) SetSurname(value *string)() {
    m.surname = value
}
// Sets the teamwork property value. A container for Microsoft Teams features available for the user. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the teamwork property.
func (m *User) SetTeamwork(value *UserTeamwork)() {
    m.teamwork = value
}
// Sets the todo property value. Represents the To Do services available to a user.
// Parameters:
//  - value : Value to set for the todo property.
func (m *User) SetTodo(value *Todo)() {
    m.todo = value
}
// Sets the transitiveMemberOf property value. 
// Parameters:
//  - value : Value to set for the transitiveMemberOf property.
func (m *User) SetTransitiveMemberOf(value []DirectoryObject)() {
    m.transitiveMemberOf = value
}
// Sets the usageLocation property value. A two letter country code (ISO standard 3166). Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.  Examples include: US, JP, and GB. Not nullable. Returned only on $select. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the usageLocation property.
func (m *User) SetUsageLocation(value *string)() {
    m.usageLocation = value
}
// Sets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization.NOTE: This property cannot contain accent characters. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, endsWith) and $orderBy.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *User) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// Sets the userType property value. A string value that can be used to classify user types in your directory, such as Member and Guest. Returned only on $select. Supports $filter (eq, ne, NOT, in). NOTE: For more information about the permissions for member and guest users, see What are the default user permissions in Azure Active Directory?
// Parameters:
//  - value : Value to set for the userType property.
func (m *User) SetUserType(value *string)() {
    m.userType = value
}
