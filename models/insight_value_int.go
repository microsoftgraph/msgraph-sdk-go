package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InsightValueInt the value in an user experience analytics insight.
type InsightValueInt struct {
    UserExperienceAnalyticsInsightValue
}
// NewInsightValueInt instantiates a new insightValueInt and sets the default values.
func NewInsightValueInt()(*InsightValueInt) {
    m := &InsightValueInt{
        UserExperienceAnalyticsInsightValue: *NewUserExperienceAnalyticsInsightValue(),
    }
    odataTypeValue := "#microsoft.graph.insightValueInt"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateInsightValueIntFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInsightValueIntFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInsightValueInt(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InsightValueInt) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UserExperienceAnalyticsInsightValue.GetFieldDeserializers()
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
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *InsightValueInt) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetValue gets the value property value. The int value of the user experience analytics insight.
func (m *InsightValueInt) GetValue()(*int32) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *InsightValueInt) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UserExperienceAnalyticsInsightValue.Serialize(writer)
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
        err = writer.WriteInt32Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *InsightValueInt) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetValue sets the value property value. The int value of the user experience analytics insight.
func (m *InsightValueInt) SetValue(value *int32)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// InsightValueIntable 
type InsightValueIntable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UserExperienceAnalyticsInsightValueable
    GetOdataType()(*string)
    GetValue()(*int32)
    SetOdataType(value *string)()
    SetValue(value *int32)()
}
