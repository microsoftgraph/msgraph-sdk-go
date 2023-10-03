package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomTrainingSetting 
type CustomTrainingSetting struct {
    TrainingSetting
}
// NewCustomTrainingSetting instantiates a new customTrainingSetting and sets the default values.
func NewCustomTrainingSetting()(*CustomTrainingSetting) {
    m := &CustomTrainingSetting{
        TrainingSetting: *NewTrainingSetting(),
    }
    odataTypeValue := "#microsoft.graph.customTrainingSetting"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCustomTrainingSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomTrainingSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomTrainingSetting(), nil
}
// GetAssignedTo gets the assignedTo property value. The assignedTo property
func (m *CustomTrainingSetting) GetAssignedTo()([]TrainingAssignedTo) {
    val, err := m.GetBackingStore().Get("assignedTo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TrainingAssignedTo)
    }
    return nil
}
// GetDescription gets the description property value. The description property
func (m *CustomTrainingSetting) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *CustomTrainingSetting) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDurationInMinutes gets the durationInMinutes property value. The durationInMinutes property
func (m *CustomTrainingSetting) GetDurationInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("durationInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomTrainingSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TrainingSetting.GetFieldDeserializers()
    res["assignedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseTrainingAssignedTo)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TrainingAssignedTo, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*TrainingAssignedTo))
                }
            }
            m.SetAssignedTo(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
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
    res["durationInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInMinutes(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetUrl gets the url property value. The url property
func (m *CustomTrainingSetting) GetUrl()(*string) {
    val, err := m.GetBackingStore().Get("url")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomTrainingSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TrainingSetting.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignedTo() != nil {
        err = writer.WriteCollectionOfStringValues("assignedTo", SerializeTrainingAssignedTo(m.GetAssignedTo()))
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
        err = writer.WriteInt32Value("durationInMinutes", m.GetDurationInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedTo sets the assignedTo property value. The assignedTo property
func (m *CustomTrainingSetting) SetAssignedTo(value []TrainingAssignedTo)() {
    err := m.GetBackingStore().Set("assignedTo", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *CustomTrainingSetting) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *CustomTrainingSetting) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDurationInMinutes sets the durationInMinutes property value. The durationInMinutes property
func (m *CustomTrainingSetting) SetDurationInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("durationInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetUrl sets the url property value. The url property
func (m *CustomTrainingSetting) SetUrl(value *string)() {
    err := m.GetBackingStore().Set("url", value)
    if err != nil {
        panic(err)
    }
}
// CustomTrainingSettingable 
type CustomTrainingSettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TrainingSettingable
    GetAssignedTo()([]TrainingAssignedTo)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDurationInMinutes()(*int32)
    GetUrl()(*string)
    SetAssignedTo(value []TrainingAssignedTo)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDurationInMinutes(value *int32)()
    SetUrl(value *string)()
}
