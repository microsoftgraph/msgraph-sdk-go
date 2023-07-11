package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InsightValueDouble the value in an user experience analytics insight.
type InsightValueDouble struct {
    UserExperienceAnalyticsInsightValue
    // The OdataType property
    OdataType *string
}
// NewInsightValueDouble instantiates a new insightValueDouble and sets the default values.
func NewInsightValueDouble()(*InsightValueDouble) {
    m := &InsightValueDouble{
        UserExperienceAnalyticsInsightValue: *NewUserExperienceAnalyticsInsightValue(),
    }
    odataTypeValue := "#microsoft.graph.insightValueDouble"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateInsightValueDoubleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInsightValueDoubleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInsightValueDouble(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InsightValueDouble) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UserExperienceAnalyticsInsightValue.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
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
// GetValue gets the value property value. The double value of the user experience analytics insight.
func (m *InsightValueDouble) GetValue()(*float64) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *InsightValueDouble) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UserExperienceAnalyticsInsightValue.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat64Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The double value of the user experience analytics insight.
func (m *InsightValueDouble) SetValue(value *float64)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// InsightValueDoubleable 
type InsightValueDoubleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UserExperienceAnalyticsInsightValueable
    GetValue()(*float64)
    SetValue(value *float64)()
}
