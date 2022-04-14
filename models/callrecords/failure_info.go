package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FailureInfo 
type FailureInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Classification of why a call or portion of a call failed.
    reason *string
    // The stage when the failure occurred. Possible values are: unknown, callSetup, midcall, unknownFutureValue.
    stage *FailureStage
}
// NewFailureInfo instantiates a new failureInfo and sets the default values.
func NewFailureInfo()(*FailureInfo) {
    m := &FailureInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFailureInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFailureInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFailureInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FailureInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FailureInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    res["stage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFailureStage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStage(val.(*FailureStage))
        }
        return nil
    }
    return res
}
// GetReason gets the reason property value. Classification of why a call or portion of a call failed.
func (m *FailureInfo) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// GetStage gets the stage property value. The stage when the failure occurred. Possible values are: unknown, callSetup, midcall, unknownFutureValue.
func (m *FailureInfo) GetStage()(*FailureStage) {
    if m == nil {
        return nil
    } else {
        return m.stage
    }
}
// Serialize serializes information the current object
func (m *FailureInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    if m.GetStage() != nil {
        cast := (*m.GetStage()).String()
        err := writer.WriteStringValue("stage", &cast)
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
func (m *FailureInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetReason sets the reason property value. Classification of why a call or portion of a call failed.
func (m *FailureInfo) SetReason(value *string)() {
    if m != nil {
        m.reason = value
    }
}
// SetStage sets the stage property value. The stage when the failure occurred. Possible values are: unknown, callSetup, midcall, unknownFutureValue.
func (m *FailureInfo) SetStage(value *FailureStage)() {
    if m != nil {
        m.stage = value
    }
}
