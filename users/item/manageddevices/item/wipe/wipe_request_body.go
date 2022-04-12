package wipe

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WipeRequestBody provides operations to call the wipe method.
type WipeRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The keepEnrollmentData property
    keepEnrollmentData *bool;
    // The keepUserData property
    keepUserData *bool;
    // The macOsUnlockCode property
    macOsUnlockCode *string;
    // The persistEsimDataPlan property
    persistEsimDataPlan *bool;
}
// NewWipeRequestBody instantiates a new wipeRequestBody and sets the default values.
func NewWipeRequestBody()(*WipeRequestBody) {
    m := &WipeRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWipeRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWipeRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWipeRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WipeRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WipeRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keepEnrollmentData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepEnrollmentData(val)
        }
        return nil
    }
    res["keepUserData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepUserData(val)
        }
        return nil
    }
    res["macOsUnlockCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacOsUnlockCode(val)
        }
        return nil
    }
    res["persistEsimDataPlan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersistEsimDataPlan(val)
        }
        return nil
    }
    return res
}
// GetKeepEnrollmentData gets the keepEnrollmentData property value. The keepEnrollmentData property
func (m *WipeRequestBody) GetKeepEnrollmentData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keepEnrollmentData
    }
}
// GetKeepUserData gets the keepUserData property value. The keepUserData property
func (m *WipeRequestBody) GetKeepUserData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keepUserData
    }
}
// GetMacOsUnlockCode gets the macOsUnlockCode property value. The macOsUnlockCode property
func (m *WipeRequestBody) GetMacOsUnlockCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.macOsUnlockCode
    }
}
// GetPersistEsimDataPlan gets the persistEsimDataPlan property value. The persistEsimDataPlan property
func (m *WipeRequestBody) GetPersistEsimDataPlan()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.persistEsimDataPlan
    }
}
// Serialize serializes information the current object
func (m *WipeRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("keepEnrollmentData", m.GetKeepEnrollmentData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("keepUserData", m.GetKeepUserData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("macOsUnlockCode", m.GetMacOsUnlockCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("persistEsimDataPlan", m.GetPersistEsimDataPlan())
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
func (m *WipeRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKeepEnrollmentData sets the keepEnrollmentData property value. The keepEnrollmentData property
func (m *WipeRequestBody) SetKeepEnrollmentData(value *bool)() {
    if m != nil {
        m.keepEnrollmentData = value
    }
}
// SetKeepUserData sets the keepUserData property value. The keepUserData property
func (m *WipeRequestBody) SetKeepUserData(value *bool)() {
    if m != nil {
        m.keepUserData = value
    }
}
// SetMacOsUnlockCode sets the macOsUnlockCode property value. The macOsUnlockCode property
func (m *WipeRequestBody) SetMacOsUnlockCode(value *string)() {
    if m != nil {
        m.macOsUnlockCode = value
    }
}
// SetPersistEsimDataPlan sets the persistEsimDataPlan property value. The persistEsimDataPlan property
func (m *WipeRequestBody) SetPersistEsimDataPlan(value *bool)() {
    if m != nil {
        m.persistEsimDataPlan = value
    }
}
