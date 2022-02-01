package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationUser 
type EducationUser struct {
    Entity
    // True if the account is enabled; otherwise, false. This property is required when a user is created. Supports /$filter.
    accountEnabled *bool;
    // The licenses that are assigned to the user. Not nullable.
    assignedLicenses []AssignedLicense;
    // The plans that are assigned to the user. Read-only. Not nullable.
    assignedPlans []AssignedPlan;
    // List of assignments for the user. Nullable.
    assignments []EducationAssignment;
    // The telephone numbers for the user. Note: Although this is a string collection, only one number can be set for this property.
    businessPhones []string;
    // Classes to which the user belongs. Nullable.
    classes []EducationClass;
    // Entity who created the user.
    createdBy *IdentitySet;
    // The name for the department in which the user works. Supports /$filter.
    department *string;
    // The name displayed in the address book for the user. Supports $filter and $orderby.
    displayName *string;
    // The type of external source this resource was generated from (automatically determined from externalSourceDetail). Possible values are: sis, lms, or manual.
    externalSource *EducationExternalSource;
    // The name of the external source this resources was generated from.
    externalSourceDetail *string;
    // The given name (first name) of the user. Supports /$filter.
    givenName *string;
    // The SMTP address for the user; for example, 'jeff@contoso.onmicrosoft.com'. Read-Only. Supports /$filter.
    mail *string;
    // Mail address of user. Note: type and postOfficeBox are not supported for educationUser resources.
    mailingAddress *PhysicalAddress;
    // The mail alias for the user. This property must be specified when a user is created. Supports /$filter.
    mailNickname *string;
    // The middle name of user.
    middleName *string;
    // The primary cellular telephone number for the user.
    mobilePhone *string;
    // 
    officeLocation *string;
    // Additional information used to associate the AAD user with it's Active Directory counterpart.
    onPremisesInfo *EducationOnPremisesInfo;
    // Specifies password policies for the user. See standard [user] resource for additional details.
    passwordPolicies *string;
    // Specifies the password profile for the user. The profile contains the user's password. This property is required when a user is created. See standard [user] resource for additional details.
    passwordProfile *PasswordProfile;
    // The preferred language for the user. Should follow ISO 639-1 Code; for example, 'en-US'.
    preferredLanguage *string;
    // Default role for a user. The user's role might be different in an individual class. Possible values are: student, teacher, faculty. Supports /$filter.
    primaryRole *EducationUserRole;
    // The plans that are provisioned for the user. Read-only. Not nullable.
    provisionedPlans []ProvisionedPlan;
    // 
    refreshTokensValidFromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Related records related to the user. Possible relationships are parent, relative, aide, doctor, guardian, child, other, unknownFutureValue
    relatedContacts []RelatedContact;
    // Address where user lives. Note: type and postOfficeBox are not supported for educationUser resources.
    residenceAddress *PhysicalAddress;
    // 
    rubrics []EducationRubric;
    // Schools to which the user belongs. Nullable.
    schools []EducationSchool;
    // true if the Outlook global address list should contain this user, otherwise false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false.
    showInAddressList *bool;
    // If the primary role is student, this block will contain student specific data.
    student *EducationStudent;
    // The user's surname (family name or last name). Supports /$filter.
    surname *string;
    // Classes for which the user is a teacher.
    taughtClasses []EducationClass;
    // If the primary role is teacher, this block will contain teacher specific data.
    teacher *EducationTeacher;
    // A two-letter country code ([ISO 3166 Alpha-2]). Required for users who will be assigned licenses. Not nullable. Supports /$filter.
    usageLocation *string;
    // The directory user corresponding to this user.
    user *User;
    // The user principal name (UPN) for the user. Supports $filter and $orderby. See standard [user] resource for additional details.
    userPrincipalName *string;
    // A string value that can be used to classify user types in your directory, such as 'Member' and 'Guest'. Supports /$filter.
    userType *string;
}
// NewEducationUser instantiates a new educationUser and sets the default values.
func NewEducationUser()(*EducationUser) {
    m := &EducationUser{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccountEnabled gets the accountEnabled property value. True if the account is enabled; otherwise, false. This property is required when a user is created. Supports /$filter.
func (m *EducationUser) GetAccountEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accountEnabled
    }
}
// GetAssignedLicenses gets the assignedLicenses property value. The licenses that are assigned to the user. Not nullable.
func (m *EducationUser) GetAssignedLicenses()([]AssignedLicense) {
    if m == nil {
        return nil
    } else {
        return m.assignedLicenses
    }
}
// GetAssignedPlans gets the assignedPlans property value. The plans that are assigned to the user. Read-only. Not nullable.
func (m *EducationUser) GetAssignedPlans()([]AssignedPlan) {
    if m == nil {
        return nil
    } else {
        return m.assignedPlans
    }
}
// GetAssignments gets the assignments property value. List of assignments for the user. Nullable.
func (m *EducationUser) GetAssignments()([]EducationAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetBusinessPhones gets the businessPhones property value. The telephone numbers for the user. Note: Although this is a string collection, only one number can be set for this property.
func (m *EducationUser) GetBusinessPhones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.businessPhones
    }
}
// GetClasses gets the classes property value. Classes to which the user belongs. Nullable.
func (m *EducationUser) GetClasses()([]EducationClass) {
    if m == nil {
        return nil
    } else {
        return m.classes
    }
}
// GetCreatedBy gets the createdBy property value. Entity who created the user.
func (m *EducationUser) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetDepartment gets the department property value. The name for the department in which the user works. Supports /$filter.
func (m *EducationUser) GetDepartment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.department
    }
}
// GetDisplayName gets the displayName property value. The name displayed in the address book for the user. Supports $filter and $orderby.
func (m *EducationUser) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExternalSource gets the externalSource property value. The type of external source this resource was generated from (automatically determined from externalSourceDetail). Possible values are: sis, lms, or manual.
func (m *EducationUser) GetExternalSource()(*EducationExternalSource) {
    if m == nil {
        return nil
    } else {
        return m.externalSource
    }
}
// GetExternalSourceDetail gets the externalSourceDetail property value. The name of the external source this resources was generated from.
func (m *EducationUser) GetExternalSourceDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalSourceDetail
    }
}
// GetGivenName gets the givenName property value. The given name (first name) of the user. Supports /$filter.
func (m *EducationUser) GetGivenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.givenName
    }
}
// GetMail gets the mail property value. The SMTP address for the user; for example, 'jeff@contoso.onmicrosoft.com'. Read-Only. Supports /$filter.
func (m *EducationUser) GetMail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mail
    }
}
// GetMailingAddress gets the mailingAddress property value. Mail address of user. Note: type and postOfficeBox are not supported for educationUser resources.
func (m *EducationUser) GetMailingAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.mailingAddress
    }
}
// GetMailNickname gets the mailNickname property value. The mail alias for the user. This property must be specified when a user is created. Supports /$filter.
func (m *EducationUser) GetMailNickname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mailNickname
    }
}
// GetMiddleName gets the middleName property value. The middle name of user.
func (m *EducationUser) GetMiddleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.middleName
    }
}
// GetMobilePhone gets the mobilePhone property value. The primary cellular telephone number for the user.
func (m *EducationUser) GetMobilePhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mobilePhone
    }
}
// GetOfficeLocation gets the officeLocation property value. 
func (m *EducationUser) GetOfficeLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.officeLocation
    }
}
// GetOnPremisesInfo gets the onPremisesInfo property value. Additional information used to associate the AAD user with it's Active Directory counterpart.
func (m *EducationUser) GetOnPremisesInfo()(*EducationOnPremisesInfo) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesInfo
    }
}
// GetPasswordPolicies gets the passwordPolicies property value. Specifies password policies for the user. See standard [user] resource for additional details.
func (m *EducationUser) GetPasswordPolicies()(*string) {
    if m == nil {
        return nil
    } else {
        return m.passwordPolicies
    }
}
// GetPasswordProfile gets the passwordProfile property value. Specifies the password profile for the user. The profile contains the user's password. This property is required when a user is created. See standard [user] resource for additional details.
func (m *EducationUser) GetPasswordProfile()(*PasswordProfile) {
    if m == nil {
        return nil
    } else {
        return m.passwordProfile
    }
}
// GetPreferredLanguage gets the preferredLanguage property value. The preferred language for the user. Should follow ISO 639-1 Code; for example, 'en-US'.
func (m *EducationUser) GetPreferredLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredLanguage
    }
}
// GetPrimaryRole gets the primaryRole property value. Default role for a user. The user's role might be different in an individual class. Possible values are: student, teacher, faculty. Supports /$filter.
func (m *EducationUser) GetPrimaryRole()(*EducationUserRole) {
    if m == nil {
        return nil
    } else {
        return m.primaryRole
    }
}
// GetProvisionedPlans gets the provisionedPlans property value. The plans that are provisioned for the user. Read-only. Not nullable.
func (m *EducationUser) GetProvisionedPlans()([]ProvisionedPlan) {
    if m == nil {
        return nil
    } else {
        return m.provisionedPlans
    }
}
// GetRefreshTokensValidFromDateTime gets the refreshTokensValidFromDateTime property value. 
func (m *EducationUser) GetRefreshTokensValidFromDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.refreshTokensValidFromDateTime
    }
}
// GetRelatedContacts gets the relatedContacts property value. Related records related to the user. Possible relationships are parent, relative, aide, doctor, guardian, child, other, unknownFutureValue
func (m *EducationUser) GetRelatedContacts()([]RelatedContact) {
    if m == nil {
        return nil
    } else {
        return m.relatedContacts
    }
}
// GetResidenceAddress gets the residenceAddress property value. Address where user lives. Note: type and postOfficeBox are not supported for educationUser resources.
func (m *EducationUser) GetResidenceAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.residenceAddress
    }
}
// GetRubrics gets the rubrics property value. 
func (m *EducationUser) GetRubrics()([]EducationRubric) {
    if m == nil {
        return nil
    } else {
        return m.rubrics
    }
}
// GetSchools gets the schools property value. Schools to which the user belongs. Nullable.
func (m *EducationUser) GetSchools()([]EducationSchool) {
    if m == nil {
        return nil
    } else {
        return m.schools
    }
}
// GetShowInAddressList gets the showInAddressList property value. true if the Outlook global address list should contain this user, otherwise false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false.
func (m *EducationUser) GetShowInAddressList()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showInAddressList
    }
}
// GetStudent gets the student property value. If the primary role is student, this block will contain student specific data.
func (m *EducationUser) GetStudent()(*EducationStudent) {
    if m == nil {
        return nil
    } else {
        return m.student
    }
}
// GetSurname gets the surname property value. The user's surname (family name or last name). Supports /$filter.
func (m *EducationUser) GetSurname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.surname
    }
}
// GetTaughtClasses gets the taughtClasses property value. Classes for which the user is a teacher.
func (m *EducationUser) GetTaughtClasses()([]EducationClass) {
    if m == nil {
        return nil
    } else {
        return m.taughtClasses
    }
}
// GetTeacher gets the teacher property value. If the primary role is teacher, this block will contain teacher specific data.
func (m *EducationUser) GetTeacher()(*EducationTeacher) {
    if m == nil {
        return nil
    } else {
        return m.teacher
    }
}
// GetUsageLocation gets the usageLocation property value. A two-letter country code ([ISO 3166 Alpha-2]). Required for users who will be assigned licenses. Not nullable. Supports /$filter.
func (m *EducationUser) GetUsageLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.usageLocation
    }
}
// GetUser gets the user property value. The directory user corresponding to this user.
func (m *EducationUser) GetUser()(*User) {
    if m == nil {
        return nil
    } else {
        return m.user
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) for the user. Supports $filter and $orderby. See standard [user] resource for additional details.
func (m *EducationUser) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetUserType gets the userType property value. A string value that can be used to classify user types in your directory, such as 'Member' and 'Guest'. Supports /$filter.
func (m *EducationUser) GetUserType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationAssignment))
            }
            m.SetAssignments(res)
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
    res["classes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationClass() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationClass, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationClass))
            }
            m.SetClasses(res)
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(*IdentitySet))
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
    res["externalSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationExternalSource)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(EducationExternalSource)
            m.SetExternalSource(&cast)
        }
        return nil
    }
    res["externalSourceDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalSourceDetail(val)
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
    res["mailingAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailingAddress(val.(*PhysicalAddress))
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
    res["middleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddleName(val)
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
    res["onPremisesInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationOnPremisesInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesInfo(val.(*EducationOnPremisesInfo))
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
    res["primaryRole"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationUserRole)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(EducationUserRole)
            m.SetPrimaryRole(&cast)
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
    res["refreshTokensValidFromDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRefreshTokensValidFromDateTime(val)
        }
        return nil
    }
    res["relatedContacts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRelatedContact() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RelatedContact, len(val))
            for i, v := range val {
                res[i] = *(v.(*RelatedContact))
            }
            m.SetRelatedContacts(res)
        }
        return nil
    }
    res["residenceAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResidenceAddress(val.(*PhysicalAddress))
        }
        return nil
    }
    res["rubrics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationRubric() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationRubric, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationRubric))
            }
            m.SetRubrics(res)
        }
        return nil
    }
    res["schools"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSchool() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSchool, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationSchool))
            }
            m.SetSchools(res)
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
    res["student"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationStudent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStudent(val.(*EducationStudent))
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
    res["taughtClasses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationClass() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationClass, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationClass))
            }
            m.SetTaughtClasses(res)
        }
        return nil
    }
    res["teacher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationTeacher() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeacher(val.(*EducationTeacher))
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
    res["user"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUser() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(*User))
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
func (m *EducationUser) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetAssignedLicenses() != nil {
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
    if m.GetAssignedPlans() != nil {
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
    if m.GetAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBusinessPhones() != nil {
        err = writer.WriteCollectionOfStringValues("businessPhones", m.GetBusinessPhones())
        if err != nil {
            return err
        }
    }
    if m.GetClasses() != nil {
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
    if m.GetProvisionedPlans() != nil {
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
    if m.GetRelatedContacts() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRelatedContacts()))
        for i, v := range m.GetRelatedContacts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("relatedContacts", cast)
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
    if m.GetRubrics() != nil {
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
    if m.GetSchools() != nil {
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
    if m.GetTaughtClasses() != nil {
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
// SetAccountEnabled sets the accountEnabled property value. True if the account is enabled; otherwise, false. This property is required when a user is created. Supports /$filter.
func (m *EducationUser) SetAccountEnabled(value *bool)() {
    if m != nil {
        m.accountEnabled = value
    }
}
// SetAssignedLicenses sets the assignedLicenses property value. The licenses that are assigned to the user. Not nullable.
func (m *EducationUser) SetAssignedLicenses(value []AssignedLicense)() {
    if m != nil {
        m.assignedLicenses = value
    }
}
// SetAssignedPlans sets the assignedPlans property value. The plans that are assigned to the user. Read-only. Not nullable.
func (m *EducationUser) SetAssignedPlans(value []AssignedPlan)() {
    if m != nil {
        m.assignedPlans = value
    }
}
// SetAssignments sets the assignments property value. List of assignments for the user. Nullable.
func (m *EducationUser) SetAssignments(value []EducationAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetBusinessPhones sets the businessPhones property value. The telephone numbers for the user. Note: Although this is a string collection, only one number can be set for this property.
func (m *EducationUser) SetBusinessPhones(value []string)() {
    if m != nil {
        m.businessPhones = value
    }
}
// SetClasses sets the classes property value. Classes to which the user belongs. Nullable.
func (m *EducationUser) SetClasses(value []EducationClass)() {
    if m != nil {
        m.classes = value
    }
}
// SetCreatedBy sets the createdBy property value. Entity who created the user.
func (m *EducationUser) SetCreatedBy(value *IdentitySet)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetDepartment sets the department property value. The name for the department in which the user works. Supports /$filter.
func (m *EducationUser) SetDepartment(value *string)() {
    if m != nil {
        m.department = value
    }
}
// SetDisplayName sets the displayName property value. The name displayed in the address book for the user. Supports $filter and $orderby.
func (m *EducationUser) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExternalSource sets the externalSource property value. The type of external source this resource was generated from (automatically determined from externalSourceDetail). Possible values are: sis, lms, or manual.
func (m *EducationUser) SetExternalSource(value *EducationExternalSource)() {
    if m != nil {
        m.externalSource = value
    }
}
// SetExternalSourceDetail sets the externalSourceDetail property value. The name of the external source this resources was generated from.
func (m *EducationUser) SetExternalSourceDetail(value *string)() {
    if m != nil {
        m.externalSourceDetail = value
    }
}
// SetGivenName sets the givenName property value. The given name (first name) of the user. Supports /$filter.
func (m *EducationUser) SetGivenName(value *string)() {
    if m != nil {
        m.givenName = value
    }
}
// SetMail sets the mail property value. The SMTP address for the user; for example, 'jeff@contoso.onmicrosoft.com'. Read-Only. Supports /$filter.
func (m *EducationUser) SetMail(value *string)() {
    if m != nil {
        m.mail = value
    }
}
// SetMailingAddress sets the mailingAddress property value. Mail address of user. Note: type and postOfficeBox are not supported for educationUser resources.
func (m *EducationUser) SetMailingAddress(value *PhysicalAddress)() {
    if m != nil {
        m.mailingAddress = value
    }
}
// SetMailNickname sets the mailNickname property value. The mail alias for the user. This property must be specified when a user is created. Supports /$filter.
func (m *EducationUser) SetMailNickname(value *string)() {
    if m != nil {
        m.mailNickname = value
    }
}
// SetMiddleName sets the middleName property value. The middle name of user.
func (m *EducationUser) SetMiddleName(value *string)() {
    if m != nil {
        m.middleName = value
    }
}
// SetMobilePhone sets the mobilePhone property value. The primary cellular telephone number for the user.
func (m *EducationUser) SetMobilePhone(value *string)() {
    if m != nil {
        m.mobilePhone = value
    }
}
// SetOfficeLocation sets the officeLocation property value. 
func (m *EducationUser) SetOfficeLocation(value *string)() {
    if m != nil {
        m.officeLocation = value
    }
}
// SetOnPremisesInfo sets the onPremisesInfo property value. Additional information used to associate the AAD user with it's Active Directory counterpart.
func (m *EducationUser) SetOnPremisesInfo(value *EducationOnPremisesInfo)() {
    if m != nil {
        m.onPremisesInfo = value
    }
}
// SetPasswordPolicies sets the passwordPolicies property value. Specifies password policies for the user. See standard [user] resource for additional details.
func (m *EducationUser) SetPasswordPolicies(value *string)() {
    if m != nil {
        m.passwordPolicies = value
    }
}
// SetPasswordProfile sets the passwordProfile property value. Specifies the password profile for the user. The profile contains the user's password. This property is required when a user is created. See standard [user] resource for additional details.
func (m *EducationUser) SetPasswordProfile(value *PasswordProfile)() {
    if m != nil {
        m.passwordProfile = value
    }
}
// SetPreferredLanguage sets the preferredLanguage property value. The preferred language for the user. Should follow ISO 639-1 Code; for example, 'en-US'.
func (m *EducationUser) SetPreferredLanguage(value *string)() {
    if m != nil {
        m.preferredLanguage = value
    }
}
// SetPrimaryRole sets the primaryRole property value. Default role for a user. The user's role might be different in an individual class. Possible values are: student, teacher, faculty. Supports /$filter.
func (m *EducationUser) SetPrimaryRole(value *EducationUserRole)() {
    if m != nil {
        m.primaryRole = value
    }
}
// SetProvisionedPlans sets the provisionedPlans property value. The plans that are provisioned for the user. Read-only. Not nullable.
func (m *EducationUser) SetProvisionedPlans(value []ProvisionedPlan)() {
    if m != nil {
        m.provisionedPlans = value
    }
}
// SetRefreshTokensValidFromDateTime sets the refreshTokensValidFromDateTime property value. 
func (m *EducationUser) SetRefreshTokensValidFromDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.refreshTokensValidFromDateTime = value
    }
}
// SetRelatedContacts sets the relatedContacts property value. Related records related to the user. Possible relationships are parent, relative, aide, doctor, guardian, child, other, unknownFutureValue
func (m *EducationUser) SetRelatedContacts(value []RelatedContact)() {
    if m != nil {
        m.relatedContacts = value
    }
}
// SetResidenceAddress sets the residenceAddress property value. Address where user lives. Note: type and postOfficeBox are not supported for educationUser resources.
func (m *EducationUser) SetResidenceAddress(value *PhysicalAddress)() {
    if m != nil {
        m.residenceAddress = value
    }
}
// SetRubrics sets the rubrics property value. 
func (m *EducationUser) SetRubrics(value []EducationRubric)() {
    if m != nil {
        m.rubrics = value
    }
}
// SetSchools sets the schools property value. Schools to which the user belongs. Nullable.
func (m *EducationUser) SetSchools(value []EducationSchool)() {
    if m != nil {
        m.schools = value
    }
}
// SetShowInAddressList sets the showInAddressList property value. true if the Outlook global address list should contain this user, otherwise false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false.
func (m *EducationUser) SetShowInAddressList(value *bool)() {
    if m != nil {
        m.showInAddressList = value
    }
}
// SetStudent sets the student property value. If the primary role is student, this block will contain student specific data.
func (m *EducationUser) SetStudent(value *EducationStudent)() {
    if m != nil {
        m.student = value
    }
}
// SetSurname sets the surname property value. The user's surname (family name or last name). Supports /$filter.
func (m *EducationUser) SetSurname(value *string)() {
    if m != nil {
        m.surname = value
    }
}
// SetTaughtClasses sets the taughtClasses property value. Classes for which the user is a teacher.
func (m *EducationUser) SetTaughtClasses(value []EducationClass)() {
    if m != nil {
        m.taughtClasses = value
    }
}
// SetTeacher sets the teacher property value. If the primary role is teacher, this block will contain teacher specific data.
func (m *EducationUser) SetTeacher(value *EducationTeacher)() {
    if m != nil {
        m.teacher = value
    }
}
// SetUsageLocation sets the usageLocation property value. A two-letter country code ([ISO 3166 Alpha-2]). Required for users who will be assigned licenses. Not nullable. Supports /$filter.
func (m *EducationUser) SetUsageLocation(value *string)() {
    if m != nil {
        m.usageLocation = value
    }
}
// SetUser sets the user property value. The directory user corresponding to this user.
func (m *EducationUser) SetUser(value *User)() {
    if m != nil {
        m.user = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) for the user. Supports $filter and $orderby. See standard [user] resource for additional details.
func (m *EducationUser) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
// SetUserType sets the userType property value. A string value that can be used to classify user types in your directory, such as 'Member' and 'Guest'. Supports /$filter.
func (m *EducationUser) SetUserType(value *string)() {
    if m != nil {
        m.userType = value
    }
}
