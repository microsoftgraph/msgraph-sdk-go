package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationGradingCategory 
type EducationGradingCategory struct {
    Entity
}
// NewEducationGradingCategory instantiates a new educationGradingCategory and sets the default values.
func NewEducationGradingCategory()(*EducationGradingCategory) {
    m := &EducationGradingCategory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationGradingCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationGradingCategoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationGradingCategory(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *EducationGradingCategory) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationGradingCategory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["percentageWeight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentageWeight(val)
        }
        return nil
    }
    return res
}
// GetPercentageWeight gets the percentageWeight property value. The percentageWeight property
func (m *EducationGradingCategory) GetPercentageWeight()(*int32) {
    val, err := m.GetBackingStore().Get("percentageWeight")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationGradingCategory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("percentageWeight", m.GetPercentageWeight())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *EducationGradingCategory) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetPercentageWeight sets the percentageWeight property value. The percentageWeight property
func (m *EducationGradingCategory) SetPercentageWeight(value *int32)() {
    err := m.GetBackingStore().Set("percentageWeight", value)
    if err != nil {
        panic(err)
    }
}
// EducationGradingCategoryable 
type EducationGradingCategoryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetPercentageWeight()(*int32)
    SetDisplayName(value *string)()
    SetPercentageWeight(value *int32)()
}
