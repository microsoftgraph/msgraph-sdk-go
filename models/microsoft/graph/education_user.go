package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EducationUser struct {
    Entity
    // True if the account is enabled; otherwise, false. This property is required when a user is created. Supports $filter.
    accountEnabled *bool;
    // The licenses that are assigned to the user. Not nullable.
    assignedLicenses []AssignedLicense;
    // The plans that are assigned to the user. Read-only. Not nullable.
    assignedPlans []AssignedPlan;
    // The telephone numbers for the user. Note: Although this is a string collection, only one number can be set for this property.
    businessPhones []string;
    // Classes to which the user belongs. Nullable.
    classes []EducationClass;
    // Entity who created the user.
    createdBy *IdentitySet;
    // The name for the department in which the user works. Supports $filter.
    department *string;
    // The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Supports $filter and $orderby.
    displayName *string;
    // Where this user was created from. Possible values are: sis, manual.
    externalSource *EducationExternalSource;
    // The name of the external source this resources was generated from.
    externalSourceDetail *string;
    // The given name (first name) of the user. Supports $filter.
    givenName *string;
    // The SMTP address for the user; for example, jeff@contoso.onmicrosoft.com. Read-Only. Supports $filter.
    mail *string;
    // Mail address of user.
    mailingAddress *PhysicalAddress;
    // The mail alias for the user. This property must be specified when a user is created. Supports $filter.
    mailNickname *string;
    // The middle name of user.
    middleName *string;
    // The primary cellular telephone number for the user.
    mobilePhone *string;
    // 
    officeLocation *string;
    // Additional information used to associate the Azure AD user with its Active Directory counterpart.
    onPremisesInfo *EducationOnPremisesInfo;
    // Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two can be specified together; for example: DisablePasswordExpiration, DisableStrongPassword.
    passwordPolicies *string;
    // Specifies the password profile for the user. The profile contains the user's password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required.
    passwordProfile *PasswordProfile;
    // The preferred language for the user. Should follow ISO 639-1 Code; for example, 'en-US'.
    preferredLanguage *string;
    // Default role for a user. The user's role might be different in an individual class. Possible values are: student, teacher, none, unknownFutureValue.
    primaryRole *EducationUserRole;
    // The plans that are provisioned for the user. Read-only. Not nullable.
    provisionedPlans []ProvisionedPlan;
    // 
    refreshTokensValidFromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Address where user lives.
    residenceAddress *PhysicalAddress;
    // 
    rubrics []EducationRubric;
    // Schools to which the user belongs. Nullable.
    schools []EducationSchool;
    // true if the Outlook global address list should contain this user, otherwise false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false.
    showInAddressList *bool;
    // If the primary role is student, this block will contain student specific data.
    student *EducationStudent;
    // The user's surname (family name or last name). Supports $filter.
    surname *string;
    // Classes for which the user is a teacher.
    taughtClasses []EducationClass;
    // If the primary role is teacher, this block will contain teacher specific data.
    teacher *EducationTeacher;
    // A two-letter country code (ISO standard 3166). Required for users who will be assigned licenses due to a legal requirement to check for availability of services in countries or regions. Examples include: 'US', 'JP', and 'GB'. Not nullable. Supports $filter.
    usageLocation *string;
    // The directory user corresponding to this user.
    user *User;
    // The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization. Supports $filter and $orderby.
    userPrincipalName *string;
    // A string value that can be used to classify user types in your directory, such as 'Member' and 'Guest'. Supports $filter.
    userType *string;
}
// Instantiates a new educationUser and sets the default values.
func NewEducationUser()(*EducationUser) {
    m := &EducationUser{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accountEnabled property value. True if the account is enabled; otherwise, false. This property is required when a user is created. Supports $filter.
func (m *EducationUser) GetAccountEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accountEnabled
    }
}
// Gets the assignedLicenses property value. The licenses that are assigned to the user. Not nullable.
func (m *EducationUser) GetAssignedLicenses()([]AssignedLicense) {
    if m == nil {
        return nil
    } else {
        return m.assignedLicenses
    }
}
// Gets the assignedPlans property value. The plans that are assigned to the user. Read-only. Not nullable.
func (m *EducationUser) GetAssignedPlans()([]AssignedPlan) {
    if m == nil {
        return nil
    } else {
        return m.assignedPlans
    }
}
// Gets the businessPhones property value. The telephone numbers for the user. Note: Although this is a string collection, only one number can be set for this property.
func (m *EducationUser) GetBusinessPhones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.businessPhones
    }
}
// Gets the classes property value. Classes to which the user belongs. Nullable.
func (m *EducationUser) GetClasses()([]EducationClass) {
    if m == nil {
        return nil
    } else {
        return m.classes
    }
}
// Gets the createdBy property value. Entity who created the user.
func (m *EducationUser) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the department property value. The name for the department in which the user works. Supports $filter.
func (m *EducationUser) GetDepartment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.department
    }
}
// Gets the displayName property value. The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Supports $filter and $orderby.
func (m *EducationUser) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the externalSource property value. Where this user was created from. Possible values are: sis, manual.
func (m *EducationUser) GetExternalSource()(*EducationExternalSource) {
    if m == nil {
        return nil
    } else {
        return m.externalSource
    }
}
// Gets the externalSourceDetail property value. The name of the external source this resources was generated from.
func (m *EducationUser) GetExternalSourceDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalSourceDetail
    }
}
// Gets the givenName property value. The given name (first name) of the user. Supports $filter.
func (m *EducationUser) GetGivenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.givenName
    }
}
// Gets the mail property value. The SMTP address for the user; for example, jeff@contoso.onmicrosoft.com. Read-Only. Supports $filter.
func (m *EducationUser) GetMail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mail
    }
}
// Gets the mailingAddress property value. Mail address of user.
func (m *EducationUser) GetMailingAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.mailingAddress
    }
}
// Gets the mailNickname property value. The mail alias for the user. This property must be specified when a user is created. Supports $filter.
func (m *EducationUser) GetMailNickname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mailNickname
    }
}
// Gets the middleName property value. The middle name of user.
func (m *EducationUser) GetMiddleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.middleName
    }
}
// Gets the mobilePhone property value. The primary cellular telephone number for the user.
func (m *EducationUser) GetMobilePhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mobilePhone
    }
}
// Gets the officeLocation property value. 
func (m *EducationUser) GetOfficeLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.officeLocation
    }
}
// Gets the onPremisesInfo property value. Additional information used to associate the Azure AD user with its Active Directory counterpart.
func (m *EducationUser) GetOnPremisesInfo()(*EducationOnPremisesInfo) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesInfo
    }
}
// Gets the passwordPolicies property value. Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two can be specified together; for example: DisablePasswordExpiration, DisableStrongPassword.
func (m *EducationUser) GetPasswordPolicies()(*string) {
    if m == nil {
        return nil
    } else {
        return m.passwordPolicies
    }
}
// Gets the passwordProfile property value. Specifies the password profile for the user. The profile contains the user's password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required.
func (m *EducationUser) GetPasswordProfile()(*PasswordProfile) {
    if m == nil {
        return nil
    } else {
        return m.passwordProfile
    }
}
// Gets the preferredLanguage property value. The preferred language for the user. Should follow ISO 639-1 Code; for example, 'en-US'.
func (m *EducationUser) GetPreferredLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredLanguage
    }
}
// Gets the primaryRole property value. Default role for a user. The user's role might be different in an individual class. Possible values are: student, teacher, none, unknownFutureValue.
func (m *EducationUser) GetPrimaryRole()(*EducationUserRole) {
    if m == nil {
        return nil
    } else {
        return m.primaryRole
    }
}
// Gets the provisionedPlans property value. The plans that are provisioned for the user. Read-only. Not nullable.
func (m *EducationUser) GetProvisionedPlans()([]ProvisionedPlan) {
    if m == nil {
        return nil
    } else {
        return m.provisionedPlans
    }
}
// Gets the refreshTokensValidFromDateTime property value. 
func (m *EducationUser) GetRefreshTokensValidFromDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.refreshTokensValidFromDateTime
    }
}
// Gets the residenceAddress property value. Address where user lives.
func (m *EducationUser) GetResidenceAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.residenceAddress
    }
}
// Gets the rubrics property value. 
func (m *EducationUser) GetRubrics()([]EducationRubric) {
    if m == nil {
        return nil
    } else {
        return m.rubrics
    }
}
// Gets the schools property value. Schools to which the user belongs. Nullable.
func (m *EducationUser) GetSchools()([]EducationSchool) {
    if m == nil {
        return nil
    } else {
        return m.schools
    }
}
// Gets the showInAddressList property value. true if the Outlook global address list should contain this user, otherwise false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false.
func (m *EducationUser) GetShowInAddressList()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showInAddressList
    }
}
// Gets the student property value. If the primary role is student, this block will contain student specific data.
func (m *EducationUser) GetStudent()(*EducationStudent) {
    if m == nil {
        return nil
    } else {
        return m.student
    }
}
// Gets the surname property value. The user's surname (family name or last name). Supports $filter.
func (m *EducationUser) GetSurname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.surname
    }
}
// Gets the taughtClasses property value. Classes for which the user is a teacher.
func (m *EducationUser) GetTaughtClasses()([]EducationClass) {
    if m == nil {
        return nil
    } else {
        return m.taughtClasses
    }
}
// Gets the teacher property value. If the primary role is teacher, this block will contain teacher specific data.
func (m *EducationUser) GetTeacher()(*EducationTeacher) {
    if m == nil {
        return nil
    } else {
        return m.teacher
    }
}
// Gets the usageLocation property value. A two-letter country code (ISO standard 3166). Required for users who will be assigned licenses due to a legal requirement to check for availability of services in countries or regions. Examples include: 'US', 'JP', and 'GB'. Not nullable. Supports $filter.
func (m *EducationUser) GetUsageLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.usageLocation
    }
}
// Gets the user property value. The directory user corresponding to this user.
func (m *EducationUser) GetUser()(*User) {
    if m == nil {
        return nil
    } else {
        return m.user
    }
}
// Gets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization. Supports $filter and $orderby.
func (m *EducationUser) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Gets the userType property value. A string value that can be used to classify user types in your directory, such as 'Member' and 'Guest'. Supports $filter.
func (m *EducationUser) GetUserType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userType
    }
}
// The deserialization information for the current model
func (m *EducationUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accountEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAccountEnabled(val)
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
    res["businessPhones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetBusinessPhones(res)
        return nil
    }
    res["classes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationClass() })
        if err != nil {
            return err
        }
        res := make([]EducationClass, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationClass))
        }
        m.SetClasses(res)
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetCreatedBy(val.(*IdentitySet))
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
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["externalSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationExternalSource)
        if err != nil {
            return err
        }
        cast := val.(EducationExternalSource)
        m.SetExternalSource(&cast)
        return nil
    }
    res["externalSourceDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalSourceDetail(val)
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
    res["mail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMail(val)
        return nil
    }
    res["mailingAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        m.SetMailingAddress(val.(*PhysicalAddress))
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
    res["middleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMiddleName(val)
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
    res["officeLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOfficeLocation(val)
        return nil
    }
    res["onPremisesInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationOnPremisesInfo() })
        if err != nil {
            return err
        }
        m.SetOnPremisesInfo(val.(*EducationOnPremisesInfo))
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
    res["preferredLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPreferredLanguage(val)
        return nil
    }
    res["primaryRole"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationUserRole)
        if err != nil {
            return err
        }
        cast := val.(EducationUserRole)
        m.SetPrimaryRole(&cast)
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
    res["refreshTokensValidFromDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetRefreshTokensValidFromDateTime(val)
        return nil
    }
    res["residenceAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        m.SetResidenceAddress(val.(*PhysicalAddress))
        return nil
    }
    res["rubrics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationRubric() })
        if err != nil {
            return err
        }
        res := make([]EducationRubric, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationRubric))
        }
        m.SetRubrics(res)
        return nil
    }
    res["schools"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSchool() })
        if err != nil {
            return err
        }
        res := make([]EducationSchool, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationSchool))
        }
        m.SetSchools(res)
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
    res["student"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationStudent() })
        if err != nil {
            return err
        }
        m.SetStudent(val.(*EducationStudent))
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
    res["taughtClasses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationClass() })
        if err != nil {
            return err
        }
        res := make([]EducationClass, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationClass))
        }
        m.SetTaughtClasses(res)
        return nil
    }
    res["teacher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationTeacher() })
        if err != nil {
            return err
        }
        m.SetTeacher(val.(*EducationTeacher))
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
    res["user"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUser() })
        if err != nil {
            return err
        }
        m.SetUser(val.(*User))
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
func (m *EducationUser) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EducationUser) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountEnabled", m.GetAccountEnabled())
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
        err = writer.WriteCollectionOfStringValues("businessPhones", m.GetBusinessPhones())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClasses()))
        for i, v := range m.GetClasses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("classes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetExternalSource() != nil {
        cast := m.GetExternalSource().String()
        err = writer.WriteStringValue("externalSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalSourceDetail", m.GetExternalSourceDetail())
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
        err = writer.WriteStringValue("mail", m.GetMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mailingAddress", m.GetMailingAddress())
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
        err = writer.WriteStringValue("middleName", m.GetMiddleName())
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
        err = writer.WriteStringValue("officeLocation", m.GetOfficeLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onPremisesInfo", m.GetOnPremisesInfo())
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
        err = writer.WriteStringValue("preferredLanguage", m.GetPreferredLanguage())
        if err != nil {
            return err
        }
    }
    if m.GetPrimaryRole() != nil {
        cast := m.GetPrimaryRole().String()
        err = writer.WriteStringValue("primaryRole", &cast)
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
        err = writer.WriteTimeValue("refreshTokensValidFromDateTime", m.GetRefreshTokensValidFromDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("residenceAddress", m.GetResidenceAddress())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRubrics()))
        for i, v := range m.GetRubrics() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("rubrics", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSchools()))
        for i, v := range m.GetSchools() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("schools", cast)
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
        err = writer.WriteObjectValue("student", m.GetStudent())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTaughtClasses()))
        for i, v := range m.GetTaughtClasses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("taughtClasses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teacher", m.GetTeacher())
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
        err = writer.WriteObjectValue("user", m.GetUser())
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
// Sets the accountEnabled property value. True if the account is enabled; otherwise, false. This property is required when a user is created. Supports $filter.
// Parameters:
//  - value : Value to set for the accountEnabled property.
func (m *EducationUser) SetAccountEnabled(value *bool)() {
    m.accountEnabled = value
}
// Sets the assignedLicenses property value. The licenses that are assigned to the user. Not nullable.
// Parameters:
//  - value : Value to set for the assignedLicenses property.
func (m *EducationUser) SetAssignedLicenses(value []AssignedLicense)() {
    m.assignedLicenses = value
}
// Sets the assignedPlans property value. The plans that are assigned to the user. Read-only. Not nullable.
// Parameters:
//  - value : Value to set for the assignedPlans property.
func (m *EducationUser) SetAssignedPlans(value []AssignedPlan)() {
    m.assignedPlans = value
}
// Sets the businessPhones property value. The telephone numbers for the user. Note: Although this is a string collection, only one number can be set for this property.
// Parameters:
//  - value : Value to set for the businessPhones property.
func (m *EducationUser) SetBusinessPhones(value []string)() {
    m.businessPhones = value
}
// Sets the classes property value. Classes to which the user belongs. Nullable.
// Parameters:
//  - value : Value to set for the classes property.
func (m *EducationUser) SetClasses(value []EducationClass)() {
    m.classes = value
}
// Sets the createdBy property value. Entity who created the user.
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *EducationUser) SetCreatedBy(value *IdentitySet)() {
    m.createdBy = value
}
// Sets the department property value. The name for the department in which the user works. Supports $filter.
// Parameters:
//  - value : Value to set for the department property.
func (m *EducationUser) SetDepartment(value *string)() {
    m.department = value
}
// Sets the displayName property value. The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Supports $filter and $orderby.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *EducationUser) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the externalSource property value. Where this user was created from. Possible values are: sis, manual.
// Parameters:
//  - value : Value to set for the externalSource property.
func (m *EducationUser) SetExternalSource(value *EducationExternalSource)() {
    m.externalSource = value
}
// Sets the externalSourceDetail property value. The name of the external source this resources was generated from.
// Parameters:
//  - value : Value to set for the externalSourceDetail property.
func (m *EducationUser) SetExternalSourceDetail(value *string)() {
    m.externalSourceDetail = value
}
// Sets the givenName property value. The given name (first name) of the user. Supports $filter.
// Parameters:
//  - value : Value to set for the givenName property.
func (m *EducationUser) SetGivenName(value *string)() {
    m.givenName = value
}
// Sets the mail property value. The SMTP address for the user; for example, jeff@contoso.onmicrosoft.com. Read-Only. Supports $filter.
// Parameters:
//  - value : Value to set for the mail property.
func (m *EducationUser) SetMail(value *string)() {
    m.mail = value
}
// Sets the mailingAddress property value. Mail address of user.
// Parameters:
//  - value : Value to set for the mailingAddress property.
func (m *EducationUser) SetMailingAddress(value *PhysicalAddress)() {
    m.mailingAddress = value
}
// Sets the mailNickname property value. The mail alias for the user. This property must be specified when a user is created. Supports $filter.
// Parameters:
//  - value : Value to set for the mailNickname property.
func (m *EducationUser) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// Sets the middleName property value. The middle name of user.
// Parameters:
//  - value : Value to set for the middleName property.
func (m *EducationUser) SetMiddleName(value *string)() {
    m.middleName = value
}
// Sets the mobilePhone property value. The primary cellular telephone number for the user.
// Parameters:
//  - value : Value to set for the mobilePhone property.
func (m *EducationUser) SetMobilePhone(value *string)() {
    m.mobilePhone = value
}
// Sets the officeLocation property value. 
// Parameters:
//  - value : Value to set for the officeLocation property.
func (m *EducationUser) SetOfficeLocation(value *string)() {
    m.officeLocation = value
}
// Sets the onPremisesInfo property value. Additional information used to associate the Azure AD user with its Active Directory counterpart.
// Parameters:
//  - value : Value to set for the onPremisesInfo property.
func (m *EducationUser) SetOnPremisesInfo(value *EducationOnPremisesInfo)() {
    m.onPremisesInfo = value
}
// Sets the passwordPolicies property value. Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two can be specified together; for example: DisablePasswordExpiration, DisableStrongPassword.
// Parameters:
//  - value : Value to set for the passwordPolicies property.
func (m *EducationUser) SetPasswordPolicies(value *string)() {
    m.passwordPolicies = value
}
// Sets the passwordProfile property value. Specifies the password profile for the user. The profile contains the user's password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required.
// Parameters:
//  - value : Value to set for the passwordProfile property.
func (m *EducationUser) SetPasswordProfile(value *PasswordProfile)() {
    m.passwordProfile = value
}
// Sets the preferredLanguage property value. The preferred language for the user. Should follow ISO 639-1 Code; for example, 'en-US'.
// Parameters:
//  - value : Value to set for the preferredLanguage property.
func (m *EducationUser) SetPreferredLanguage(value *string)() {
    m.preferredLanguage = value
}
// Sets the primaryRole property value. Default role for a user. The user's role might be different in an individual class. Possible values are: student, teacher, none, unknownFutureValue.
// Parameters:
//  - value : Value to set for the primaryRole property.
func (m *EducationUser) SetPrimaryRole(value *EducationUserRole)() {
    m.primaryRole = value
}
// Sets the provisionedPlans property value. The plans that are provisioned for the user. Read-only. Not nullable.
// Parameters:
//  - value : Value to set for the provisionedPlans property.
func (m *EducationUser) SetProvisionedPlans(value []ProvisionedPlan)() {
    m.provisionedPlans = value
}
// Sets the refreshTokensValidFromDateTime property value. 
// Parameters:
//  - value : Value to set for the refreshTokensValidFromDateTime property.
func (m *EducationUser) SetRefreshTokensValidFromDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.refreshTokensValidFromDateTime = value
}
// Sets the residenceAddress property value. Address where user lives.
// Parameters:
//  - value : Value to set for the residenceAddress property.
func (m *EducationUser) SetResidenceAddress(value *PhysicalAddress)() {
    m.residenceAddress = value
}
// Sets the rubrics property value. 
// Parameters:
//  - value : Value to set for the rubrics property.
func (m *EducationUser) SetRubrics(value []EducationRubric)() {
    m.rubrics = value
}
// Sets the schools property value. Schools to which the user belongs. Nullable.
// Parameters:
//  - value : Value to set for the schools property.
func (m *EducationUser) SetSchools(value []EducationSchool)() {
    m.schools = value
}
// Sets the showInAddressList property value. true if the Outlook global address list should contain this user, otherwise false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false.
// Parameters:
//  - value : Value to set for the showInAddressList property.
func (m *EducationUser) SetShowInAddressList(value *bool)() {
    m.showInAddressList = value
}
// Sets the student property value. If the primary role is student, this block will contain student specific data.
// Parameters:
//  - value : Value to set for the student property.
func (m *EducationUser) SetStudent(value *EducationStudent)() {
    m.student = value
}
// Sets the surname property value. The user's surname (family name or last name). Supports $filter.
// Parameters:
//  - value : Value to set for the surname property.
func (m *EducationUser) SetSurname(value *string)() {
    m.surname = value
}
// Sets the taughtClasses property value. Classes for which the user is a teacher.
// Parameters:
//  - value : Value to set for the taughtClasses property.
func (m *EducationUser) SetTaughtClasses(value []EducationClass)() {
    m.taughtClasses = value
}
// Sets the teacher property value. If the primary role is teacher, this block will contain teacher specific data.
// Parameters:
//  - value : Value to set for the teacher property.
func (m *EducationUser) SetTeacher(value *EducationTeacher)() {
    m.teacher = value
}
// Sets the usageLocation property value. A two-letter country code (ISO standard 3166). Required for users who will be assigned licenses due to a legal requirement to check for availability of services in countries or regions. Examples include: 'US', 'JP', and 'GB'. Not nullable. Supports $filter.
// Parameters:
//  - value : Value to set for the usageLocation property.
func (m *EducationUser) SetUsageLocation(value *string)() {
    m.usageLocation = value
}
// Sets the user property value. The directory user corresponding to this user.
// Parameters:
//  - value : Value to set for the user property.
func (m *EducationUser) SetUser(value *User)() {
    m.user = value
}
// Sets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization. Supports $filter and $orderby.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *EducationUser) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// Sets the userType property value. A string value that can be used to classify user types in your directory, such as 'Member' and 'Guest'. Supports $filter.
// Parameters:
//  - value : Value to set for the userType property.
func (m *EducationUser) SetUserType(value *string)() {
    m.userType = value
}
