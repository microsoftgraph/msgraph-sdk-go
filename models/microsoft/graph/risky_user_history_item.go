package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RiskyUserHistoryItem 
type RiskyUserHistoryItem struct {
    RiskyUser
    // The activity related to user risk level change.
    activity *RiskUserActivity;
    // The id of actor that does the operation.
    initiatedBy *string;
    // The id of the user.
    userId *string;
}
// NewRiskyUserHistoryItem instantiates a new riskyUserHistoryItem and sets the default values.
func NewRiskyUserHistoryItem()(*RiskyUserHistoryItem) {
    m := &RiskyUserHistoryItem{
        RiskyUser: *NewRiskyUser(),
    }
    return m
}
// GetActivity gets the activity property value. The activity related to user risk level change.
func (m *RiskyUserHistoryItem) GetActivity()(*RiskUserActivity) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// GetInitiatedBy gets the initiatedBy property value. The id of actor that does the operation.
func (m *RiskyUserHistoryItem) GetInitiatedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.initiatedBy
    }
}
// GetUserId gets the userId property value. The id of the user.
func (m *RiskyUserHistoryItem) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RiskyUserHistoryItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.RiskyUser.GetFieldDeserializers()
    res["activity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRiskUserActivity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val.(*RiskUserActivity))
        }
        return nil
    }
    res["initiatedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatedBy(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
func (m *RiskyUserHistoryItem) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RiskyUserHistoryItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.RiskyUser.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initiatedBy", m.GetInitiatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivity sets the activity property value. The activity related to user risk level change.
func (m *RiskyUserHistoryItem) SetActivity(value *RiskUserActivity)() {
    if m != nil {
        m.activity = value
    }
}
// SetInitiatedBy sets the initiatedBy property value. The id of actor that does the operation.
func (m *RiskyUserHistoryItem) SetInitiatedBy(value *string)() {
    if m != nil {
        m.initiatedBy = value
    }
}
// SetUserId sets the userId property value. The id of the user.
func (m *RiskyUserHistoryItem) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
