package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmployeeExperienceUser 
type EmployeeExperienceUser struct {
    Entity
}
// NewEmployeeExperienceUser instantiates a new employeeExperienceUser and sets the default values.
func NewEmployeeExperienceUser()(*EmployeeExperienceUser) {
    m := &EmployeeExperienceUser{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEmployeeExperienceUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmployeeExperienceUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmployeeExperienceUser(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmployeeExperienceUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["learningCourseActivities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLearningCourseActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LearningCourseActivityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LearningCourseActivityable)
                }
            }
            m.SetLearningCourseActivities(res)
        }
        return nil
    }
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
    return res
}
// GetLearningCourseActivities gets the learningCourseActivities property value. The learningCourseActivities property
func (m *EmployeeExperienceUser) GetLearningCourseActivities()([]LearningCourseActivityable) {
    val, err := m.GetBackingStore().Get("learningCourseActivities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]LearningCourseActivityable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EmployeeExperienceUser) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EmployeeExperienceUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetLearningCourseActivities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLearningCourseActivities()))
        for i, v := range m.GetLearningCourseActivities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("learningCourseActivities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLearningCourseActivities sets the learningCourseActivities property value. The learningCourseActivities property
func (m *EmployeeExperienceUser) SetLearningCourseActivities(value []LearningCourseActivityable)() {
    err := m.GetBackingStore().Set("learningCourseActivities", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EmployeeExperienceUser) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// EmployeeExperienceUserable 
type EmployeeExperienceUserable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLearningCourseActivities()([]LearningCourseActivityable)
    GetOdataType()(*string)
    SetLearningCourseActivities(value []LearningCourseActivityable)()
    SetOdataType(value *string)()
}
