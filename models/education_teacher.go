package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationTeacher 
type EducationTeacher struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Id of the Teacher in external source system.
    externalId *string
    // The OdataType property
    odataType *string
    // Teacher number.
    teacherNumber *string
}
// NewEducationTeacher instantiates a new educationTeacher and sets the default values.
func NewEducationTeacher()(*EducationTeacher) {
    m := &EducationTeacher{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.educationTeacher";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEducationTeacherFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationTeacherFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationTeacher(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationTeacher) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExternalId gets the externalId property value. Id of the Teacher in external source system.
func (m *EducationTeacher) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationTeacher) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
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
    res["teacherNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeacherNumber(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationTeacher) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// GetTeacherNumber gets the teacherNumber property value. Teacher number.
func (m *EducationTeacher) GetTeacherNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teacherNumber
    }
}
// Serialize serializes information the current object
func (m *EducationTeacher) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("teacherNumber", m.GetTeacherNumber())
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
func (m *EducationTeacher) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExternalId sets the externalId property value. Id of the Teacher in external source system.
func (m *EducationTeacher) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationTeacher) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
// SetTeacherNumber sets the teacherNumber property value. Teacher number.
func (m *EducationTeacher) SetTeacherNumber(value *string)() {
    if m != nil {
        m.teacherNumber = value
    }
}
