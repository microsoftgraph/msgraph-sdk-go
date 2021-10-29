package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EducationClass struct {
    Entity
    // 
    assignmentCategories []EducationCategory;
    // 
    assignmentDefaults *EducationAssignmentDefaults;
    // All assignments associated with this class. Nullable.
    assignments []EducationAssignment;
    // 
    assignmentSettings *EducationAssignmentSettings;
    // Class code used by the school to identify the class.
    classCode *string;
    // Course information for the class
    course *EducationCourse;
    // Entity who created the class
    createdBy *IdentitySet;
    // Description of the class.
    description *string;
    // Name of the class.
    displayName *string;
    // ID of the class from the syncing system.
    externalId *string;
    // Name of the class in the syncing system.
    externalName *string;
    // How this class was created. Possible values are: sis, manual.
    externalSource *EducationExternalSource;
    // The name of the external source this resources was generated from.
    externalSourceDetail *string;
    // Grade level of the class.
    grade *string;
    // The underlying Microsoft 365 group object.
    group *Group;
    // Mail name for sending email to all members, if this is enabled.
    mailNickname *string;
    // All users in the class. Nullable.
    members []EducationUser;
    // All schools that this class is associated with. Nullable.
    schools []EducationSchool;
    // All teachers in the class. Nullable.
    teachers []EducationUser;
    // Term for this class.
    term *EducationTerm;
}
// Instantiates a new educationClass and sets the default values.
func NewEducationClass()(*EducationClass) {
    m := &EducationClass{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignmentCategories property value. 
func (m *EducationClass) GetAssignmentCategories()([]EducationCategory) {
    if m == nil {
        return nil
    } else {
        return m.assignmentCategories
    }
}
// Gets the assignmentDefaults property value. 
func (m *EducationClass) GetAssignmentDefaults()(*EducationAssignmentDefaults) {
    if m == nil {
        return nil
    } else {
        return m.assignmentDefaults
    }
}
// Gets the assignments property value. All assignments associated with this class. Nullable.
func (m *EducationClass) GetAssignments()([]EducationAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the assignmentSettings property value. 
func (m *EducationClass) GetAssignmentSettings()(*EducationAssignmentSettings) {
    if m == nil {
        return nil
    } else {
        return m.assignmentSettings
    }
}
// Gets the classCode property value. Class code used by the school to identify the class.
func (m *EducationClass) GetClassCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.classCode
    }
}
// Gets the course property value. Course information for the class
func (m *EducationClass) GetCourse()(*EducationCourse) {
    if m == nil {
        return nil
    } else {
        return m.course
    }
}
// Gets the createdBy property value. Entity who created the class
func (m *EducationClass) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the description property value. Description of the class.
func (m *EducationClass) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Name of the class.
func (m *EducationClass) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the externalId property value. ID of the class from the syncing system.
func (m *EducationClass) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// Gets the externalName property value. Name of the class in the syncing system.
func (m *EducationClass) GetExternalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalName
    }
}
// Gets the externalSource property value. How this class was created. Possible values are: sis, manual.
func (m *EducationClass) GetExternalSource()(*EducationExternalSource) {
    if m == nil {
        return nil
    } else {
        return m.externalSource
    }
}
// Gets the externalSourceDetail property value. The name of the external source this resources was generated from.
func (m *EducationClass) GetExternalSourceDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalSourceDetail
    }
}
// Gets the grade property value. Grade level of the class.
func (m *EducationClass) GetGrade()(*string) {
    if m == nil {
        return nil
    } else {
        return m.grade
    }
}
// Gets the group property value. The underlying Microsoft 365 group object.
func (m *EducationClass) GetGroup()(*Group) {
    if m == nil {
        return nil
    } else {
        return m.group
    }
}
// Gets the mailNickname property value. Mail name for sending email to all members, if this is enabled.
func (m *EducationClass) GetMailNickname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mailNickname
    }
}
// Gets the members property value. All users in the class. Nullable.
func (m *EducationClass) GetMembers()([]EducationUser) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// Gets the schools property value. All schools that this class is associated with. Nullable.
func (m *EducationClass) GetSchools()([]EducationSchool) {
    if m == nil {
        return nil
    } else {
        return m.schools
    }
}
// Gets the teachers property value. All teachers in the class. Nullable.
func (m *EducationClass) GetTeachers()([]EducationUser) {
    if m == nil {
        return nil
    } else {
        return m.teachers
    }
}
// Gets the term property value. Term for this class.
func (m *EducationClass) GetTerm()(*EducationTerm) {
    if m == nil {
        return nil
    } else {
        return m.term
    }
}
// The deserialization information for the current model
func (m *EducationClass) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationCategory() })
        if err != nil {
            return err
        }
        res := make([]EducationCategory, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationCategory))
        }
        m.SetAssignmentCategories(res)
        return nil
    }
    res["assignmentDefaults"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationAssignmentDefaults() })
        if err != nil {
            return err
        }
        m.SetAssignmentDefaults(val.(*EducationAssignmentDefaults))
        return nil
    }
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationAssignment() })
        if err != nil {
            return err
        }
        res := make([]EducationAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["assignmentSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationAssignmentSettings() })
        if err != nil {
            return err
        }
        m.SetAssignmentSettings(val.(*EducationAssignmentSettings))
        return nil
    }
    res["classCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClassCode(val)
        return nil
    }
    res["course"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationCourse() })
        if err != nil {
            return err
        }
        m.SetCourse(val.(*EducationCourse))
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
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
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
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalId(val)
        return nil
    }
    res["externalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalName(val)
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
    res["grade"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGrade(val)
        return nil
    }
    res["group"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroup() })
        if err != nil {
            return err
        }
        m.SetGroup(val.(*Group))
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
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationUser() })
        if err != nil {
            return err
        }
        res := make([]EducationUser, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationUser))
        }
        m.SetMembers(res)
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
    res["teachers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationUser() })
        if err != nil {
            return err
        }
        res := make([]EducationUser, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationUser))
        }
        m.SetTeachers(res)
        return nil
    }
    res["term"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationTerm() })
        if err != nil {
            return err
        }
        m.SetTerm(val.(*EducationTerm))
        return nil
    }
    return res
}
func (m *EducationClass) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EducationClass) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignmentCategories()))
        for i, v := range m.GetAssignmentCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignmentCategories", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assignmentDefaults", m.GetAssignmentDefaults())
        if err != nil {
            return err
        }
    }
    {
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
    {
        err = writer.WriteObjectValue("assignmentSettings", m.GetAssignmentSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("classCode", m.GetClassCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("course", m.GetCourse())
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
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalName", m.GetExternalName())
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
        err = writer.WriteStringValue("grade", m.GetGrade())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTeachers()))
        for i, v := range m.GetTeachers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("teachers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("term", m.GetTerm())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignmentCategories property value. 
// Parameters:
//  - value : Value to set for the assignmentCategories property.
func (m *EducationClass) SetAssignmentCategories(value []EducationCategory)() {
    m.assignmentCategories = value
}
// Sets the assignmentDefaults property value. 
// Parameters:
//  - value : Value to set for the assignmentDefaults property.
func (m *EducationClass) SetAssignmentDefaults(value *EducationAssignmentDefaults)() {
    m.assignmentDefaults = value
}
// Sets the assignments property value. All assignments associated with this class. Nullable.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *EducationClass) SetAssignments(value []EducationAssignment)() {
    m.assignments = value
}
// Sets the assignmentSettings property value. 
// Parameters:
//  - value : Value to set for the assignmentSettings property.
func (m *EducationClass) SetAssignmentSettings(value *EducationAssignmentSettings)() {
    m.assignmentSettings = value
}
// Sets the classCode property value. Class code used by the school to identify the class.
// Parameters:
//  - value : Value to set for the classCode property.
func (m *EducationClass) SetClassCode(value *string)() {
    m.classCode = value
}
// Sets the course property value. Course information for the class
// Parameters:
//  - value : Value to set for the course property.
func (m *EducationClass) SetCourse(value *EducationCourse)() {
    m.course = value
}
// Sets the createdBy property value. Entity who created the class
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *EducationClass) SetCreatedBy(value *IdentitySet)() {
    m.createdBy = value
}
// Sets the description property value. Description of the class.
// Parameters:
//  - value : Value to set for the description property.
func (m *EducationClass) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Name of the class.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *EducationClass) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the externalId property value. ID of the class from the syncing system.
// Parameters:
//  - value : Value to set for the externalId property.
func (m *EducationClass) SetExternalId(value *string)() {
    m.externalId = value
}
// Sets the externalName property value. Name of the class in the syncing system.
// Parameters:
//  - value : Value to set for the externalName property.
func (m *EducationClass) SetExternalName(value *string)() {
    m.externalName = value
}
// Sets the externalSource property value. How this class was created. Possible values are: sis, manual.
// Parameters:
//  - value : Value to set for the externalSource property.
func (m *EducationClass) SetExternalSource(value *EducationExternalSource)() {
    m.externalSource = value
}
// Sets the externalSourceDetail property value. The name of the external source this resources was generated from.
// Parameters:
//  - value : Value to set for the externalSourceDetail property.
func (m *EducationClass) SetExternalSourceDetail(value *string)() {
    m.externalSourceDetail = value
}
// Sets the grade property value. Grade level of the class.
// Parameters:
//  - value : Value to set for the grade property.
func (m *EducationClass) SetGrade(value *string)() {
    m.grade = value
}
// Sets the group property value. The underlying Microsoft 365 group object.
// Parameters:
//  - value : Value to set for the group property.
func (m *EducationClass) SetGroup(value *Group)() {
    m.group = value
}
// Sets the mailNickname property value. Mail name for sending email to all members, if this is enabled.
// Parameters:
//  - value : Value to set for the mailNickname property.
func (m *EducationClass) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// Sets the members property value. All users in the class. Nullable.
// Parameters:
//  - value : Value to set for the members property.
func (m *EducationClass) SetMembers(value []EducationUser)() {
    m.members = value
}
// Sets the schools property value. All schools that this class is associated with. Nullable.
// Parameters:
//  - value : Value to set for the schools property.
func (m *EducationClass) SetSchools(value []EducationSchool)() {
    m.schools = value
}
// Sets the teachers property value. All teachers in the class. Nullable.
// Parameters:
//  - value : Value to set for the teachers property.
func (m *EducationClass) SetTeachers(value []EducationUser)() {
    m.teachers = value
}
// Sets the term property value. Term for this class.
// Parameters:
//  - value : Value to set for the term property.
func (m *EducationClass) SetTerm(value *EducationTerm)() {
    m.term = value
}
