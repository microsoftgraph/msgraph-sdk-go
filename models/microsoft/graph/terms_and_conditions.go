package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TermsAndConditions 
type TermsAndConditions struct {
    Entity
    // Administrator-supplied explanation of the terms and conditions, typically describing what it means to accept the terms and conditions set out in the T&C policy. This is shown to the user on prompts to accept the T&C policy.
    acceptanceStatement *string;
    // The list of acceptance statuses for this T&C policy.
    acceptanceStatuses []TermsAndConditionsAcceptanceStatus;
    // The list of assignments for this T&C policy.
    assignments []TermsAndConditionsAssignment;
    // Administrator-supplied body text of the terms and conditions, typically the terms themselves. This is shown to the user on prompts to accept the T&C policy.
    bodyText *string;
    // DateTime the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Administrator-supplied description of the T&C policy.
    description *string;
    // Administrator-supplied name for the T&C policy.
    displayName *string;
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Administrator-supplied title of the terms and conditions. This is shown to the user on prompts to accept the T&C policy.
    title *string;
    // Integer indicating the current version of the terms. Incremented when an administrator makes a change to the terms and wishes to require users to re-accept the modified T&C policy.
    version *int32;
}
// NewTermsAndConditions instantiates a new termsAndConditions and sets the default values.
func NewTermsAndConditions()(*TermsAndConditions) {
    m := &TermsAndConditions{
        Entity: *NewEntity(),
    }
    return m
}
// GetAcceptanceStatement gets the acceptanceStatement property value. Administrator-supplied explanation of the terms and conditions, typically describing what it means to accept the terms and conditions set out in the T&C policy. This is shown to the user on prompts to accept the T&C policy.
func (m *TermsAndConditions) GetAcceptanceStatement()(*string) {
    if m == nil {
        return nil
    } else {
        return m.acceptanceStatement
    }
}
// GetAcceptanceStatuses gets the acceptanceStatuses property value. The list of acceptance statuses for this T&C policy.
func (m *TermsAndConditions) GetAcceptanceStatuses()([]TermsAndConditionsAcceptanceStatus) {
    if m == nil {
        return nil
    } else {
        return m.acceptanceStatuses
    }
}
// GetAssignments gets the assignments property value. The list of assignments for this T&C policy.
func (m *TermsAndConditions) GetAssignments()([]TermsAndConditionsAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetBodyText gets the bodyText property value. Administrator-supplied body text of the terms and conditions, typically the terms themselves. This is shown to the user on prompts to accept the T&C policy.
func (m *TermsAndConditions) GetBodyText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bodyText
    }
}
// GetCreatedDateTime gets the createdDateTime property value. DateTime the object was created.
func (m *TermsAndConditions) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. Administrator-supplied description of the T&C policy.
func (m *TermsAndConditions) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Administrator-supplied name for the T&C policy.
func (m *TermsAndConditions) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *TermsAndConditions) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetTitle gets the title property value. Administrator-supplied title of the terms and conditions. This is shown to the user on prompts to accept the T&C policy.
func (m *TermsAndConditions) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// GetVersion gets the version property value. Integer indicating the current version of the terms. Incremented when an administrator makes a change to the terms and wishes to require users to re-accept the modified T&C policy.
func (m *TermsAndConditions) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TermsAndConditions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acceptanceStatement"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptanceStatement(val)
        }
        return nil
    }
    res["acceptanceStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTermsAndConditionsAcceptanceStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TermsAndConditionsAcceptanceStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*TermsAndConditionsAcceptanceStatus))
            }
            m.SetAcceptanceStatuses(res)
        }
        return nil
    }
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTermsAndConditionsAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TermsAndConditionsAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*TermsAndConditionsAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["bodyText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyText(val)
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
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
func (m *TermsAndConditions) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TermsAndConditions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("acceptanceStatement", m.GetAcceptanceStatement())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAcceptanceStatuses()))
        for i, v := range m.GetAcceptanceStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("acceptanceStatuses", cast)
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
        err = writer.WriteStringValue("bodyText", m.GetBodyText())
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptanceStatement sets the acceptanceStatement property value. Administrator-supplied explanation of the terms and conditions, typically describing what it means to accept the terms and conditions set out in the T&C policy. This is shown to the user on prompts to accept the T&C policy.
func (m *TermsAndConditions) SetAcceptanceStatement(value *string)() {
    m.acceptanceStatement = value
}
// SetAcceptanceStatuses sets the acceptanceStatuses property value. The list of acceptance statuses for this T&C policy.
func (m *TermsAndConditions) SetAcceptanceStatuses(value []TermsAndConditionsAcceptanceStatus)() {
    m.acceptanceStatuses = value
}
// SetAssignments sets the assignments property value. The list of assignments for this T&C policy.
func (m *TermsAndConditions) SetAssignments(value []TermsAndConditionsAssignment)() {
    m.assignments = value
}
// SetBodyText sets the bodyText property value. Administrator-supplied body text of the terms and conditions, typically the terms themselves. This is shown to the user on prompts to accept the T&C policy.
func (m *TermsAndConditions) SetBodyText(value *string)() {
    m.bodyText = value
}
// SetCreatedDateTime sets the createdDateTime property value. DateTime the object was created.
func (m *TermsAndConditions) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Administrator-supplied description of the T&C policy.
func (m *TermsAndConditions) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Administrator-supplied name for the T&C policy.
func (m *TermsAndConditions) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *TermsAndConditions) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetTitle sets the title property value. Administrator-supplied title of the terms and conditions. This is shown to the user on prompts to accept the T&C policy.
func (m *TermsAndConditions) SetTitle(value *string)() {
    m.title = value
}
// SetVersion sets the version property value. Integer indicating the current version of the terms. Incremented when an administrator makes a change to the terms and wishes to require users to re-accept the modified T&C policy.
func (m *TermsAndConditions) SetVersion(value *int32)() {
    m.version = value
}
