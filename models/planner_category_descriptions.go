package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerCategoryDescriptions 
type PlannerCategoryDescriptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The label associated with Category 1
    category1 *string;
    // The label associated with Category 2
    category2 *string;
    // The label associated with Category 3
    category3 *string;
    // The label associated with Category 4
    category4 *string;
    // The label associated with Category 5
    category5 *string;
    // The label associated with Category 6
    category6 *string;
}
// NewPlannerCategoryDescriptions instantiates a new plannerCategoryDescriptions and sets the default values.
func NewPlannerCategoryDescriptions()(*PlannerCategoryDescriptions) {
    m := &PlannerCategoryDescriptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePlannerCategoryDescriptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerCategoryDescriptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerCategoryDescriptions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerCategoryDescriptions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCategory1 gets the category1 property value. The label associated with Category 1
func (m *PlannerCategoryDescriptions) GetCategory1()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category1
    }
}
// GetCategory2 gets the category2 property value. The label associated with Category 2
func (m *PlannerCategoryDescriptions) GetCategory2()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category2
    }
}
// GetCategory3 gets the category3 property value. The label associated with Category 3
func (m *PlannerCategoryDescriptions) GetCategory3()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category3
    }
}
// GetCategory4 gets the category4 property value. The label associated with Category 4
func (m *PlannerCategoryDescriptions) GetCategory4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category4
    }
}
// GetCategory5 gets the category5 property value. The label associated with Category 5
func (m *PlannerCategoryDescriptions) GetCategory5()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category5
    }
}
// GetCategory6 gets the category6 property value. The label associated with Category 6
func (m *PlannerCategoryDescriptions) GetCategory6()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category6
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerCategoryDescriptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["category1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory1(val)
        }
        return nil
    }
    res["category2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory2(val)
        }
        return nil
    }
    res["category3"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory3(val)
        }
        return nil
    }
    res["category4"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory4(val)
        }
        return nil
    }
    res["category5"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory5(val)
        }
        return nil
    }
    res["category6"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory6(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PlannerCategoryDescriptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("category1", m.GetCategory1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category2", m.GetCategory2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category3", m.GetCategory3())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category4", m.GetCategory4())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category5", m.GetCategory5())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category6", m.GetCategory6())
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
func (m *PlannerCategoryDescriptions) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCategory1 sets the category1 property value. The label associated with Category 1
func (m *PlannerCategoryDescriptions) SetCategory1(value *string)() {
    if m != nil {
        m.category1 = value
    }
}
// SetCategory2 sets the category2 property value. The label associated with Category 2
func (m *PlannerCategoryDescriptions) SetCategory2(value *string)() {
    if m != nil {
        m.category2 = value
    }
}
// SetCategory3 sets the category3 property value. The label associated with Category 3
func (m *PlannerCategoryDescriptions) SetCategory3(value *string)() {
    if m != nil {
        m.category3 = value
    }
}
// SetCategory4 sets the category4 property value. The label associated with Category 4
func (m *PlannerCategoryDescriptions) SetCategory4(value *string)() {
    if m != nil {
        m.category4 = value
    }
}
// SetCategory5 sets the category5 property value. The label associated with Category 5
func (m *PlannerCategoryDescriptions) SetCategory5(value *string)() {
    if m != nil {
        m.category5 = value
    }
}
// SetCategory6 sets the category6 property value. The label associated with Category 6
func (m *PlannerCategoryDescriptions) SetCategory6(value *string)() {
    if m != nil {
        m.category6 = value
    }
}
