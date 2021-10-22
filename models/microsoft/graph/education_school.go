package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EducationSchool struct {
    EducationOrganization
    address *PhysicalAddress;
    administrativeUnit *AdministrativeUnit;
    classes []EducationClass;
    createdBy *IdentitySet;
    externalId *string;
    externalPrincipalId *string;
    fax *string;
    highestGrade *string;
    lowestGrade *string;
    phone *string;
    principalEmail *string;
    principalName *string;
    schoolNumber *string;
    users []EducationUser;
}
func NewEducationSchool()(*EducationSchool) {
    m := &EducationSchool{
        EducationOrganization: *NewEducationOrganization(),
    }
    return m
}
func (m *EducationSchool) GetAddress()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
func (m *EducationSchool) GetAdministrativeUnit()(*AdministrativeUnit) {
    if m == nil {
        return nil
    } else {
        return m.administrativeUnit
    }
}
func (m *EducationSchool) GetClasses()([]EducationClass) {
    if m == nil {
        return nil
    } else {
        return m.classes
    }
}
func (m *EducationSchool) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
func (m *EducationSchool) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
func (m *EducationSchool) GetExternalPrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalPrincipalId
    }
}
func (m *EducationSchool) GetFax()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fax
    }
}
func (m *EducationSchool) GetHighestGrade()(*string) {
    if m == nil {
        return nil
    } else {
        return m.highestGrade
    }
}
func (m *EducationSchool) GetLowestGrade()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lowestGrade
    }
}
func (m *EducationSchool) GetPhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phone
    }
}
func (m *EducationSchool) GetPrincipalEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalEmail
    }
}
func (m *EducationSchool) GetPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalName
    }
}
func (m *EducationSchool) GetSchoolNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.schoolNumber
    }
}
func (m *EducationSchool) GetUsers()([]EducationUser) {
    if m == nil {
        return nil
    } else {
        return m.users
    }
}
func (m *EducationSchool) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.EducationOrganization.GetFieldDeserializers()
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        m.SetAddress(val.(*PhysicalAddress))
        return nil
    }
    res["administrativeUnit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAdministrativeUnit() })
        if err != nil {
            return err
        }
        m.SetAdministrativeUnit(val.(*AdministrativeUnit))
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
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalId(val)
        return nil
    }
    res["externalPrincipalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalPrincipalId(val)
        return nil
    }
    res["fax"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFax(val)
        return nil
    }
    res["highestGrade"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHighestGrade(val)
        return nil
    }
    res["lowestGrade"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLowestGrade(val)
        return nil
    }
    res["phone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPhone(val)
        return nil
    }
    res["principalEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPrincipalEmail(val)
        return nil
    }
    res["principalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPrincipalName(val)
        return nil
    }
    res["schoolNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSchoolNumber(val)
        return nil
    }
    res["users"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationUser() })
        if err != nil {
            return err
        }
        res := make([]EducationUser, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationUser))
        }
        m.SetUsers(res)
        return nil
    }
    return res
}
func (m *EducationSchool) IsNil()(bool) {
    return m == nil
}
func (m *EducationSchool) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.EducationOrganization.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("administrativeUnit", m.GetAdministrativeUnit())
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
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalPrincipalId", m.GetExternalPrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fax", m.GetFax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("highestGrade", m.GetHighestGrade())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lowestGrade", m.GetLowestGrade())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalEmail", m.GetPrincipalEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalName", m.GetPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("schoolNumber", m.GetSchoolNumber())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUsers()))
        for i, v := range m.GetUsers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("users", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *EducationSchool) SetAddress(value *PhysicalAddress)() {
    m.address = value
}
func (m *EducationSchool) SetAdministrativeUnit(value *AdministrativeUnit)() {
    m.administrativeUnit = value
}
func (m *EducationSchool) SetClasses(value []EducationClass)() {
    m.classes = value
}
func (m *EducationSchool) SetCreatedBy(value *IdentitySet)() {
    m.createdBy = value
}
func (m *EducationSchool) SetExternalId(value *string)() {
    m.externalId = value
}
func (m *EducationSchool) SetExternalPrincipalId(value *string)() {
    m.externalPrincipalId = value
}
func (m *EducationSchool) SetFax(value *string)() {
    m.fax = value
}
func (m *EducationSchool) SetHighestGrade(value *string)() {
    m.highestGrade = value
}
func (m *EducationSchool) SetLowestGrade(value *string)() {
    m.lowestGrade = value
}
func (m *EducationSchool) SetPhone(value *string)() {
    m.phone = value
}
func (m *EducationSchool) SetPrincipalEmail(value *string)() {
    m.principalEmail = value
}
func (m *EducationSchool) SetPrincipalName(value *string)() {
    m.principalName = value
}
func (m *EducationSchool) SetSchoolNumber(value *string)() {
    m.schoolNumber = value
}
func (m *EducationSchool) SetUsers(value []EducationUser)() {
    m.users = value
}
