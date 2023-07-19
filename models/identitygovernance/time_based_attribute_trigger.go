package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeBasedAttributeTrigger 
type TimeBasedAttributeTrigger struct {
    WorkflowExecutionTrigger
}
// NewTimeBasedAttributeTrigger instantiates a new timeBasedAttributeTrigger and sets the default values.
func NewTimeBasedAttributeTrigger()(*TimeBasedAttributeTrigger) {
    m := &TimeBasedAttributeTrigger{
        WorkflowExecutionTrigger: *NewWorkflowExecutionTrigger(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.timeBasedAttributeTrigger"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTimeBasedAttributeTriggerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeBasedAttributeTriggerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeBasedAttributeTrigger(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeBasedAttributeTrigger) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WorkflowExecutionTrigger.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["offsetInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffsetInDays(val)
        }
        return nil
    }
    res["timeBasedAttribute"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkflowTriggerTimeBasedAttribute)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeBasedAttribute(val.(*WorkflowTriggerTimeBasedAttribute))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TimeBasedAttributeTrigger) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOffsetInDays gets the offsetInDays property value. How many days before or after the time-based attribute specified the workflow should trigger. For example, if the attribute is employeeHireDate and offsetInDays is -1, then the workflow should trigger one day before the employee hire date. The value can range between -180 and 180 days.
func (m *TimeBasedAttributeTrigger) GetOffsetInDays()(*int32) {
    val, err := m.GetBackingStore().Get("offsetInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTimeBasedAttribute gets the timeBasedAttribute property value. The timeBasedAttribute property
func (m *TimeBasedAttributeTrigger) GetTimeBasedAttribute()(*WorkflowTriggerTimeBasedAttribute) {
    val, err := m.GetBackingStore().Get("timeBasedAttribute")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WorkflowTriggerTimeBasedAttribute)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TimeBasedAttributeTrigger) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WorkflowExecutionTrigger.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("offsetInDays", m.GetOffsetInDays())
        if err != nil {
            return err
        }
    }
    if m.GetTimeBasedAttribute() != nil {
        cast := (*m.GetTimeBasedAttribute()).String()
        err = writer.WriteStringValue("timeBasedAttribute", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TimeBasedAttributeTrigger) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOffsetInDays sets the offsetInDays property value. How many days before or after the time-based attribute specified the workflow should trigger. For example, if the attribute is employeeHireDate and offsetInDays is -1, then the workflow should trigger one day before the employee hire date. The value can range between -180 and 180 days.
func (m *TimeBasedAttributeTrigger) SetOffsetInDays(value *int32)() {
    err := m.GetBackingStore().Set("offsetInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeBasedAttribute sets the timeBasedAttribute property value. The timeBasedAttribute property
func (m *TimeBasedAttributeTrigger) SetTimeBasedAttribute(value *WorkflowTriggerTimeBasedAttribute)() {
    err := m.GetBackingStore().Set("timeBasedAttribute", value)
    if err != nil {
        panic(err)
    }
}
// TimeBasedAttributeTriggerable 
type TimeBasedAttributeTriggerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WorkflowExecutionTriggerable
    GetOdataType()(*string)
    GetOffsetInDays()(*int32)
    GetTimeBasedAttribute()(*WorkflowTriggerTimeBasedAttribute)
    SetOdataType(value *string)()
    SetOffsetInDays(value *int32)()
    SetTimeBasedAttribute(value *WorkflowTriggerTimeBasedAttribute)()
}
