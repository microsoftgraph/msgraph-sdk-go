package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeZoneBase 
type TimeZoneBase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The name of a time zone. It can be a standard time zone name such as 'Hawaii-Aleutian Standard Time', or 'Customized Time Zone' for a custom time zone.
    name *string;
}
// NewTimeZoneBase instantiates a new timeZoneBase and sets the default values.
func NewTimeZoneBase()(*TimeZoneBase) {
    m := &TimeZoneBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTimeZoneBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeZoneBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeZoneBase(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeZoneBase) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeZoneBase) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of a time zone. It can be a standard time zone name such as 'Hawaii-Aleutian Standard Time', or 'Customized Time Zone' for a custom time zone.
func (m *TimeZoneBase) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Serialize serializes information the current object
func (m *TimeZoneBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *TimeZoneBase) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetName sets the name property value. The name of a time zone. It can be a standard time zone name such as 'Hawaii-Aleutian Standard Time', or 'Customized Time Zone' for a custom time zone.
func (m *TimeZoneBase) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
