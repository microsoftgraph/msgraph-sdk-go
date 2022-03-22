package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RiskyUserHistoryItem 
type RiskyUserHistoryItem struct {
    RiskyUser
    // The activity related to user risk level change.
    activity RiskUserActivityable;
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
// CreateRiskyUserHistoryItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRiskyUserHistoryItemFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRiskyUserHistoryItem(), nil
}
// GetActivity gets the activity property value. The activity related to user risk level change.
func (m *RiskyUserHistoryItem) GetActivity()(RiskUserActivityable) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RiskyUserHistoryItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.RiskyUser.GetFieldDeserializers()
    res["activity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateRiskUserActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val.(RiskUserActivityable))
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
func (m *RiskyUserHistoryItem) SetActivity(value RiskUserActivityable)() {
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
