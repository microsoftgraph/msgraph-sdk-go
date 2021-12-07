package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CopyNotebookModel 
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
// NewCopyNotebookModel instantiates a new CopyNotebookModel and sets the default values.
func NewCopyNotebookModel()(*CopyNotebookModel) {
    m := &CopyNotebookModel{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CopyNotebookModel) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCreatedBy gets the createdBy property value. 
func (m *CopyNotebookModel) GetCreatedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedByIdentity gets the createdByIdentity property value. 
func (m *CopyNotebookModel) GetCreatedByIdentity()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdByIdentity
    }
}
// GetCreatedTime gets the createdTime property value. 
func (m *CopyNotebookModel) GetCreatedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdTime
    }
}
// GetId gets the id property value. 
func (m *CopyNotebookModel) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetIsDefault gets the isDefault property value. 
func (m *CopyNotebookModel) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetIsShared gets the isShared property value. 
func (m *CopyNotebookModel) GetIsShared()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isShared
    }
}
// GetLastModifiedBy gets the lastModifiedBy property value. 
func (m *CopyNotebookModel) GetLastModifiedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedByIdentity gets the lastModifiedByIdentity property value. 
func (m *CopyNotebookModel) GetLastModifiedByIdentity()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedByIdentity
    }
}
// GetLastModifiedTime gets the lastModifiedTime property value. 
func (m *CopyNotebookModel) GetLastModifiedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedTime
    }
}
// GetLinks gets the links property value. 
func (m *CopyNotebookModel) GetLinks()(*NotebookLinks) {
    if m == nil {
        return nil
    } else {
        return m.links
    }
}
// GetName gets the name property value. 
func (m *CopyNotebookModel) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetSectionGroupsUrl gets the sectionGroupsUrl property value. 
func (m *CopyNotebookModel) GetSectionGroupsUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sectionGroupsUrl
    }
}
// GetSectionsUrl gets the sectionsUrl property value. 
func (m *CopyNotebookModel) GetSectionsUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sectionsUrl
    }
}
// GetSelf gets the self property value. 
func (m *CopyNotebookModel) GetSelf()(*string) {
    if m == nil {
        return nil
    } else {
        return m.self
    }
}
// GetUserRole gets the userRole property value. 
func (m *CopyNotebookModel) GetUserRole()(*OnenoteUserRole) {
    if m == nil {
        return nil
    } else {
        return m.userRole
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CopyNotebookModel) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCreatedBy sets the createdBy property value. 
func (m *CopyNotebookModel) SetCreatedBy(value *string)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedByIdentity sets the createdByIdentity property value. 
func (m *CopyNotebookModel) SetCreatedByIdentity(value *IdentitySet)() {
    if m != nil {
        m.createdByIdentity = value
    }
}
// SetCreatedTime sets the createdTime property value. 
func (m *CopyNotebookModel) SetCreatedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdTime = value
    }
}
// SetId sets the id property value. 
func (m *CopyNotebookModel) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetIsDefault sets the isDefault property value. 
func (m *CopyNotebookModel) SetIsDefault(value *bool)() {
    if m != nil {
        m.isDefault = value
    }
}
// SetIsShared sets the isShared property value. 
func (m *CopyNotebookModel) SetIsShared(value *bool)() {
    if m != nil {
        m.isShared = value
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. 
func (m *CopyNotebookModel) SetLastModifiedBy(value *string)() {
    if m != nil {
        m.lastModifiedBy = value
    }
}
// SetLastModifiedByIdentity sets the lastModifiedByIdentity property value. 
func (m *CopyNotebookModel) SetLastModifiedByIdentity(value *IdentitySet)() {
    if m != nil {
        m.lastModifiedByIdentity = value
    }
}
// SetLastModifiedTime sets the lastModifiedTime property value. 
func (m *CopyNotebookModel) SetLastModifiedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedTime = value
    }
}
// SetLinks sets the links property value. 
func (m *CopyNotebookModel) SetLinks(value *NotebookLinks)() {
    if m != nil {
        m.links = value
    }
}
// SetName sets the name property value. 
func (m *CopyNotebookModel) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetSectionGroupsUrl sets the sectionGroupsUrl property value. 
func (m *CopyNotebookModel) SetSectionGroupsUrl(value *string)() {
    if m != nil {
        m.sectionGroupsUrl = value
    }
}
// SetSectionsUrl sets the sectionsUrl property value. 
func (m *CopyNotebookModel) SetSectionsUrl(value *string)() {
    if m != nil {
        m.sectionsUrl = value
    }
}
// SetSelf sets the self property value. 
func (m *CopyNotebookModel) SetSelf(value *string)() {
    if m != nil {
        m.self = value
    }
}
// SetUserRole sets the userRole property value. 
func (m *CopyNotebookModel) SetUserRole(value *OnenoteUserRole)() {
    if m != nil {
        m.userRole = value
    }
}
