package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CopyNotebookModel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    createdBy *string;
    // 
    createdByIdentity *IdentitySet;
    // 
    createdTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    id *string;
    // 
    isDefault *bool;
    // 
    isShared *bool;
    // 
    lastModifiedBy *string;
    // 
    lastModifiedByIdentity *IdentitySet;
    // 
    lastModifiedTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    links *NotebookLinks;
    // 
    name *string;
    // 
    sectionGroupsUrl *string;
    // 
    sectionsUrl *string;
    // 
    self *string;
    // 
    userRole *OnenoteUserRole;
}
// Instantiates a new CopyNotebookModel and sets the default values.
func NewCopyNotebookModel()(*CopyNotebookModel) {
    m := &CopyNotebookModel{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CopyNotebookModel) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the createdBy property value. 
func (m *CopyNotebookModel) GetCreatedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the createdByIdentity property value. 
func (m *CopyNotebookModel) GetCreatedByIdentity()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdByIdentity
    }
}
// Gets the createdTime property value. 
func (m *CopyNotebookModel) GetCreatedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdTime
    }
}
// Gets the id property value. 
func (m *CopyNotebookModel) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the isDefault property value. 
func (m *CopyNotebookModel) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// Gets the isShared property value. 
func (m *CopyNotebookModel) GetIsShared()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isShared
    }
}
// Gets the lastModifiedBy property value. 
func (m *CopyNotebookModel) GetLastModifiedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// Gets the lastModifiedByIdentity property value. 
func (m *CopyNotebookModel) GetLastModifiedByIdentity()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedByIdentity
    }
}
// Gets the lastModifiedTime property value. 
func (m *CopyNotebookModel) GetLastModifiedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedTime
    }
}
// Gets the links property value. 
func (m *CopyNotebookModel) GetLinks()(*NotebookLinks) {
    if m == nil {
        return nil
    } else {
        return m.links
    }
}
// Gets the name property value. 
func (m *CopyNotebookModel) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the sectionGroupsUrl property value. 
func (m *CopyNotebookModel) GetSectionGroupsUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sectionGroupsUrl
    }
}
// Gets the sectionsUrl property value. 
func (m *CopyNotebookModel) GetSectionsUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sectionsUrl
    }
}
// Gets the self property value. 
func (m *CopyNotebookModel) GetSelf()(*string) {
    if m == nil {
        return nil
    } else {
        return m.self
    }
}
// Gets the userRole property value. 
func (m *CopyNotebookModel) GetUserRole()(*OnenoteUserRole) {
    if m == nil {
        return nil
    } else {
        return m.userRole
    }
}
// The deserialization information for the current model
func (m *CopyNotebookModel) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val)
        }
        return nil
    }
    res["createdByIdentity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedByIdentity(val.(*IdentitySet))
        }
        return nil
    }
    res["createdTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedTime(val)
        }
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["isShared"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsShared(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val)
        }
        return nil
    }
    res["lastModifiedByIdentity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedByIdentity(val.(*IdentitySet))
        }
        return nil
    }
    res["lastModifiedTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedTime(val)
        }
        return nil
    }
    res["links"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNotebookLinks() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(*NotebookLinks))
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["sectionGroupsUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSectionGroupsUrl(val)
        }
        return nil
    }
    res["sectionsUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSectionsUrl(val)
        }
        return nil
    }
    res["self"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelf(val)
        }
        return nil
    }
    res["userRole"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnenoteUserRole)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(OnenoteUserRole)
            m.SetUserRole(&cast)
        }
        return nil
    }
    return res
}
func (m *CopyNotebookModel) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CopyNotebookModel) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("createdByIdentity", m.GetCreatedByIdentity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdTime", m.GetCreatedTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isShared", m.GetIsShared())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lastModifiedByIdentity", m.GetLastModifiedByIdentity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastModifiedTime", m.GetLastModifiedTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("links", m.GetLinks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sectionGroupsUrl", m.GetSectionGroupsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sectionsUrl", m.GetSectionsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("self", m.GetSelf())
        if err != nil {
            return err
        }
    }
    if m.GetUserRole() != nil {
        cast := m.GetUserRole().String()
        err := writer.WriteStringValue("userRole", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *CopyNotebookModel) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the createdBy property value. 
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *CopyNotebookModel) SetCreatedBy(value *string)() {
    m.createdBy = value
}
// Sets the createdByIdentity property value. 
// Parameters:
//  - value : Value to set for the createdByIdentity property.
func (m *CopyNotebookModel) SetCreatedByIdentity(value *IdentitySet)() {
    m.createdByIdentity = value
}
// Sets the createdTime property value. 
// Parameters:
//  - value : Value to set for the createdTime property.
func (m *CopyNotebookModel) SetCreatedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdTime = value
}
// Sets the id property value. 
// Parameters:
//  - value : Value to set for the id property.
func (m *CopyNotebookModel) SetId(value *string)() {
    m.id = value
}
// Sets the isDefault property value. 
// Parameters:
//  - value : Value to set for the isDefault property.
func (m *CopyNotebookModel) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// Sets the isShared property value. 
// Parameters:
//  - value : Value to set for the isShared property.
func (m *CopyNotebookModel) SetIsShared(value *bool)() {
    m.isShared = value
}
// Sets the lastModifiedBy property value. 
// Parameters:
//  - value : Value to set for the lastModifiedBy property.
func (m *CopyNotebookModel) SetLastModifiedBy(value *string)() {
    m.lastModifiedBy = value
}
// Sets the lastModifiedByIdentity property value. 
// Parameters:
//  - value : Value to set for the lastModifiedByIdentity property.
func (m *CopyNotebookModel) SetLastModifiedByIdentity(value *IdentitySet)() {
    m.lastModifiedByIdentity = value
}
// Sets the lastModifiedTime property value. 
// Parameters:
//  - value : Value to set for the lastModifiedTime property.
func (m *CopyNotebookModel) SetLastModifiedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedTime = value
}
// Sets the links property value. 
// Parameters:
//  - value : Value to set for the links property.
func (m *CopyNotebookModel) SetLinks(value *NotebookLinks)() {
    m.links = value
}
// Sets the name property value. 
// Parameters:
//  - value : Value to set for the name property.
func (m *CopyNotebookModel) SetName(value *string)() {
    m.name = value
}
// Sets the sectionGroupsUrl property value. 
// Parameters:
//  - value : Value to set for the sectionGroupsUrl property.
func (m *CopyNotebookModel) SetSectionGroupsUrl(value *string)() {
    m.sectionGroupsUrl = value
}
// Sets the sectionsUrl property value. 
// Parameters:
//  - value : Value to set for the sectionsUrl property.
func (m *CopyNotebookModel) SetSectionsUrl(value *string)() {
    m.sectionsUrl = value
}
// Sets the self property value. 
// Parameters:
//  - value : Value to set for the self property.
func (m *CopyNotebookModel) SetSelf(value *string)() {
    m.self = value
}
// Sets the userRole property value. 
// Parameters:
//  - value : Value to set for the userRole property.
func (m *CopyNotebookModel) SetUserRole(value *OnenoteUserRole)() {
    m.userRole = value
}
