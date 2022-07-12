package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationUser provides operations to manage the admin singleton.
type EducationUser struct {
    Entity
    // True if the account is enabled; otherwise, false. This property is required when a user is created. Supports /$filter.
    accountEnabled *bool
    // The licenses that are assigned to the user. Not nullable.
    assignedLicenses []AssignedLicenseable
    // The plans that are assigned to the user. Read-only. Not nullable.
    assignedPlans []AssignedPlanable
    // List of assignments for the user. Nullable.
    assignments []EducationAssignmentable
    // The telephone numbers for the user. Note: Although this is a string collection, only one number can be set for this property.
    businessPhones []string
    // Classes to which the user belongs. Nullable.
    classes []EducationClassable
    // Entity who created the user.
    createdBy IdentitySetable
    // The name for the department in which the user works. Supports /$filter.
    department *string
    // The name displayed in the address book for the user. Supports $filter and $orderby.
    displayName *string
    // The type of external source this resource was generated from (automatically determined from externalSourceDetail). Possible values are: sis, lms, or manual.
    externalSource *EducationExternalSource
    // The name of the external source this resources was generated from.
    externalSourceDetail *string
    // The given name (first name) of the user. Supports /$filter.
    givenName *string
    // The SMTP address for the user; for example, 'jeff@contoso.onmicrosoft.com'. Read-Only. Supports /$filter.
    mail *string
    // Mail address of user. Note: type and postOfficeBox are not supported for educationUser resources.
    mailingAddress PhysicalAddressable
    // The mail alias for the user. This property must be specified when a user is created. Supports /$filter.
    mailNickname *string
    // The middle name of user.
    middleName *string
    // The primary cellular telephone number for the user.
    mobilePhone *string
    // The officeLocation property
    officeLocation *string
    // Additional information used to associate the AAD user with it's Active Directory counterpart.
    onPremisesInfo EducationOnPremisesInfoable
    // Specifies password policies for the user. See standard [user] resource for additional details.
    passwordPolicies *string
    // Specifies the password profile for the user. The profile contains the user's password. This property is required when a user is created. See standard [user] resource for additional details.
    passwordProfile PasswordProfileable
    // The preferred language for the user. Should follow ISO 639-1 Code; for example, 'en-US'.
    preferredLanguage *string
    // The primaryRole property
    primaryRole *EducationUserRole
    // The plans that are provisioned for the user. Read-only. Not nullable.
    provisionedPlans []ProvisionedPlanable
    // The refreshTokensValidFromDateTime property
    refreshTokensValidFromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Related records related to the user. Possible relationships are parent, relative, aide, doctor, guardian, child, other, unknownFutureValue
    relatedContacts []RelatedContactable
    // Address where user lives. Note: type and postOfficeBox are not supported for educationUser resources.
    residenceAddress PhysicalAddressable
    // When set, the grading rubric attached to the assignment.
    rubrics []EducationRubricable
    // Schools to which the user belongs. Nullable.
    schools []EducationSchoolable
    // True if the Outlook Global Address List should contain this user; otherwise, false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false.
    showInAddressList *bool
    // If the primary role is student, this block will contain student specific data.
    student EducationStudentable
    // The user's surname (family name or last name). Supports /$filter.
    surname *string
    // Classes for which the user is a teacher.
    taughtClasses []EducationClassable
    // If the primary role is teacher, this block will contain teacher specific data.
    teacher EducationTeacherable
    // A two-letter country code ([ISO 3166 Alpha-2]). Required for users who will be assigned licenses. Not nullable. Supports /$filter.
    usageLocation *string
    // The directory user that corresponds to this user.
    user Userable
    // The user principal name (UPN) for the user. Supports $filter and $orderby. See standard [user] resource for additional details.
    userPrincipalName *string
    // A string value that can be used to classify user types in your directory, such as 'Member' and 'Guest'. Supports /$filter.
    userType *string
}
// NewEducationUser instantiates a new educationUser and sets the default values.
func NewEducationUser()(*EducationUser) {
    m := &EducationUser{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationUser(), nil
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
func (m *EducationUser) GetAssignedLicenses()([]AssignedLicenseable) {
    if m == nil {
        return nil
    } else {
        return m.assignedLicenses
    }
}
// GetAssignedPlans gets the assignedPlans property value. The plans that are assigned to the user. Read-only. Not nullable.
func (m *EducationUser) GetAssignedPlans()([]AssignedPlanable) {
    if m == nil {
        return nil
    } else {
        return m.assignedPlans
    }
}
// GetAssignments gets the assignments property value. List of assignments for the user. Nullable.
func (m *EducationUser) GetAssignments()([]EducationAssignmentable) {
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
func (m *EducationUser) GetClasses()([]EducationClassable) {
    if m == nil {
        return nil
    } else {
        return m.classes
    }
}
// GetCreatedBy gets the createdBy property value. Entity who created the user.
func (m *EducationUser) GetCreatedBy()(IdentitySetable) {
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
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accountEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountEnabled(val)
        }
        return nil
    }
    res["assignedLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignedLicenseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedLicenseable, len(val))
            for i, v := range val {
                res[i] = v.(AssignedLicenseable)
            }
            m.SetAssignedLicenses(res)
        }
        return nil
    }
    res["assignedPlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignedPlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedPlanable, len(val))
            for i, v := range val {
                res[i] = v.(AssignedPlanable)
            }
            m.SetAssignedPlans(res)
        }
        return nil
    }
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(EducationAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["businessPhones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["classes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationClassFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationClassable, len(val))
            for i, v := range val {
                res[i] = v.(EducationClassable)
            }
            m.SetClasses(res)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["department"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepartment(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["externalSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationExternalSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalSource(val.(*EducationExternalSource))
        }
        return nil
    }
    res["externalSourceDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalSourceDetail(val)
        }
        return nil
    }
    res["givenName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
        }
        return nil
    }
    res["mail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMail(val)
        }
        return nil
    }
    res["mailingAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailingAddress(val.(PhysicalAddressable))
        }
        return nil
    }
    res["mailNickname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailNickname(val)
        }
        return nil
    }
    res["middleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddleName(val)
        }
        return nil
    }
    res["mobilePhone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobilePhone(val)
        }
        return nil
    }
    res["officeLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeLocation(val)
        }
        return nil
    }
    res["onPremisesInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationOnPremisesInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesInfo(val.(EducationOnPremisesInfoable))
        }
        return nil
    }
    res["passwordPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordPolicies(val)
        }
        return nil
    }
    res["passwordProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePasswordProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordProfile(val.(PasswordProfileable))
        }
        return nil
    }
    res["preferredLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredLanguage(val)
        }
        return nil
    }
    res["primaryRole"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationUserRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryRole(val.(*EducationUserRole))
        }
        return nil
    }
    res["provisionedPlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProvisionedPlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisionedPlanable, len(val))
            for i, v := range val {
                res[i] = v.(ProvisionedPlanable)
            }
            m.SetProvisionedPlans(res)
        }
        return nil
    }
    res["refreshTokensValidFromDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRefreshTokensValidFromDateTime(val)
        }
        return nil
    }
    res["relatedContacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRelatedContactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RelatedContactable, len(val))
            for i, v := range val {
                res[i] = v.(RelatedContactable)
            }
            m.SetRelatedContacts(res)
        }
        return nil
    }
    res["residenceAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResidenceAddress(val.(PhysicalAddressable))
        }
        return nil
    }
    res["rubrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationRubricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationRubricable, len(val))
            for i, v := range val {
                res[i] = v.(EducationRubricable)
            }
            m.SetRubrics(res)
        }
        return nil
    }
    res["schools"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationSchoolFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSchoolable, len(val))
            for i, v := range val {
                res[i] = v.(EducationSchoolable)
            }
            m.SetSchools(res)
        }
        return nil
    }
    res["showInAddressList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowInAddressList(val)
        }
        return nil
    }
    res["student"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationStudentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStudent(val.(EducationStudentable))
        }
        return nil
    }
    res["surname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSurname(val)
        }
        return nil
    }
    res["taughtClasses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationClassFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationClassable, len(val))
            for i, v := range val {
                res[i] = v.(EducationClassable)
            }
            m.SetTaughtClasses(res)
        }
        return nil
    }
    res["teacher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationTeacherFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeacher(val.(EducationTeacherable))
        }
        return nil
    }
    res["usageLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsageLocation(val)
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(Userable))
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["userType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *EducationUser) GetMailingAddress()(PhysicalAddressable) {
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
// GetOfficeLocation gets the officeLocation property value. The officeLocation property
func (m *EducationUser) GetOfficeLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.officeLocation
    }
}
// GetOnPremisesInfo gets the onPremisesInfo property value. Additional information used to associate the AAD user with it's Active Directory counterpart.
func (m *EducationUser) GetOnPremisesInfo()(EducationOnPremisesInfoable) {
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
func (m *EducationUser) GetPasswordProfile()(PasswordProfileable) {
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
// GetPrimaryRole gets the primaryRole property value. The primaryRole property
func (m *EducationUser) GetPrimaryRole()(*EducationUserRole) {
    if m == nil {
        return nil
    } else {
        return m.primaryRole
    }
}
// GetProvisionedPlans gets the provisionedPlans property value. The plans that are provisioned for the user. Read-only. Not nullable.
func (m *EducationUser) GetProvisionedPlans()([]ProvisionedPlanable) {
    if m == nil {
        return nil
    } else {
        return m.provisionedPlans
    }
}
// GetRefreshTokensValidFromDateTime gets the refreshTokensValidFromDateTime property value. The refreshTokensValidFromDateTime property
func (m *EducationUser) GetRefreshTokensValidFromDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.refreshTokensValidFromDateTime
    }
}
// GetRelatedContacts gets the relatedContacts property value. Related records related to the user. Possible relationships are parent, relative, aide, doctor, guardian, child, other, unknownFutureValue
func (m *EducationUser) GetRelatedContacts()([]RelatedContactable) {
    if m == nil {
        return nil
    } else {
        return m.relatedContacts
    }
}
// GetResidenceAddress gets the residenceAddress property value. Address where user lives. Note: type and postOfficeBox are not supported for educationUser resources.
func (m *EducationUser) GetResidenceAddress()(PhysicalAddressable) {
    if m == nil {
        return nil
    } else {
        return m.residenceAddress
    }
}
// GetRubrics gets the rubrics property value. When set, the grading rubric attached to the assignment.
func (m *EducationUser) GetRubrics()([]EducationRubricable) {
    if m == nil {
        return nil
    } else {
        return m.rubrics
    }
}
// GetSchools gets the schools property value. Schools to which the user belongs. Nullable.
func (m *EducationUser) GetSchools()([]EducationSchoolable) {
    if m == nil {
        return nil
    } else {
        return m.schools
    }
}
// GetShowInAddressList gets the showInAddressList property value. True if the Outlook Global Address List should contain this user; otherwise, false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false.
func (m *EducationUser) GetShowInAddressList()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showInAddressList
    }
}
// GetStudent gets the student property value. If the primary role is student, this block will contain student specific data.
func (m *EducationUser) GetStudent()(EducationStudentable) {
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
func (m *EducationUser) GetTaughtClasses()([]EducationClassable) {
    if m == nil {
        return nil
    } else {
        return m.taughtClasses
    }
}
// GetTeacher gets the teacher property value. If the primary role is teacher, this block will contain teacher specific data.
func (m *EducationUser) GetTeacher()(EducationTeacherable) {
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
// GetUser gets the user property value. The directory user that corresponds to this user.
func (m *EducationUser) GetUser()(Userable) {
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
// Serialize serializes information the current object
func (m *EducationUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedLicenses()))
        for i, v := range m.GetAssignedLicenses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignedLicenses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignedPlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedPlans()))
        for i, v := range m.GetAssignedPlans() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignedPlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClasses()))
        for i, v := range m.GetClasses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := (*m.GetExternalSource()).String()
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
        cast := (*m.GetPrimaryRole()).String()
        err = writer.WriteStringValue("primaryRole", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetProvisionedPlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProvisionedPlans()))
        for i, v := range m.GetProvisionedPlans() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRelatedContacts()))
        for i, v := range m.GetRelatedContacts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRubrics()))
        for i, v := range m.GetRubrics() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("rubrics", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSchools() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSchools()))
        for i, v := range m.GetSchools() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaughtClasses()))
        for i, v := range m.GetTaughtClasses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
func (m *EducationUser) SetAssignedLicenses(value []AssignedLicenseable)() {
    if m != nil {
        m.assignedLicenses = value
    }
}
// SetAssignedPlans sets the assignedPlans property value. The plans that are assigned to the user. Read-only. Not nullable.
func (m *EducationUser) SetAssignedPlans(value []AssignedPlanable)() {
    if m != nil {
        m.assignedPlans = value
    }
}
// SetAssignments sets the assignments property value. List of assignments for the user. Nullable.
func (m *EducationUser) SetAssignments(value []EducationAssignmentable)() {
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
func (m *EducationUser) SetClasses(value []EducationClassable)() {
    if m != nil {
        m.classes = value
    }
}
// SetCreatedBy sets the createdBy property value. Entity who created the user.
func (m *EducationUser) SetCreatedBy(value IdentitySetable)() {
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
func (m *EducationUser) SetMailingAddress(value PhysicalAddressable)() {
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
// SetOfficeLocation sets the officeLocation property value. The officeLocation property
func (m *EducationUser) SetOfficeLocation(value *string)() {
    if m != nil {
        m.officeLocation = value
    }
}
// SetOnPremisesInfo sets the onPremisesInfo property value. Additional information used to associate the AAD user with it's Active Directory counterpart.
func (m *EducationUser) SetOnPremisesInfo(value EducationOnPremisesInfoable)() {
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
func (m *EducationUser) SetPasswordProfile(value PasswordProfileable)() {
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
// SetPrimaryRole sets the primaryRole property value. The primaryRole property
func (m *EducationUser) SetPrimaryRole(value *EducationUserRole)() {
    if m != nil {
        m.primaryRole = value
    }
}
// SetProvisionedPlans sets the provisionedPlans property value. The plans that are provisioned for the user. Read-only. Not nullable.
func (m *EducationUser) SetProvisionedPlans(value []ProvisionedPlanable)() {
    if m != nil {
        m.provisionedPlans = value
    }
}
// SetRefreshTokensValidFromDateTime sets the refreshTokensValidFromDateTime property value. The refreshTokensValidFromDateTime property
func (m *EducationUser) SetRefreshTokensValidFromDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.refreshTokensValidFromDateTime = value
    }
}
// SetRelatedContacts sets the relatedContacts property value. Related records related to the user. Possible relationships are parent, relative, aide, doctor, guardian, child, other, unknownFutureValue
func (m *EducationUser) SetRelatedContacts(value []RelatedContactable)() {
    if m != nil {
        m.relatedContacts = value
    }
}
// SetResidenceAddress sets the residenceAddress property value. Address where user lives. Note: type and postOfficeBox are not supported for educationUser resources.
func (m *EducationUser) SetResidenceAddress(value PhysicalAddressable)() {
    if m != nil {
        m.residenceAddress = value
    }
}
// SetRubrics sets the rubrics property value. When set, the grading rubric attached to the assignment.
func (m *EducationUser) SetRubrics(value []EducationRubricable)() {
    if m != nil {
        m.rubrics = value
    }
}
// SetSchools sets the schools property value. Schools to which the user belongs. Nullable.
func (m *EducationUser) SetSchools(value []EducationSchoolable)() {
    if m != nil {
        m.schools = value
    }
}
// SetShowInAddressList sets the showInAddressList property value. True if the Outlook Global Address List should contain this user; otherwise, false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false.
func (m *EducationUser) SetShowInAddressList(value *bool)() {
    if m != nil {
        m.showInAddressList = value
    }
}
// SetStudent sets the student property value. If the primary role is student, this block will contain student specific data.
func (m *EducationUser) SetStudent(value EducationStudentable)() {
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
func (m *EducationUser) SetTaughtClasses(value []EducationClassable)() {
    if m != nil {
        m.taughtClasses = value
    }
}
// SetTeacher sets the teacher property value. If the primary role is teacher, this block will contain teacher specific data.
func (m *EducationUser) SetTeacher(value EducationTeacherable)() {
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
// SetUser sets the user property value. The directory user that corresponds to this user.
func (m *EducationUser) SetUser(value Userable)() {
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
