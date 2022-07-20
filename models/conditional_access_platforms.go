package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessPlatforms 
type ConditionalAccessPlatforms struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue, linux.
    excludePlatforms []string
    // Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue,linux``.
    includePlatforms []string
    // The OdataType property
    odataType *string
}
// NewConditionalAccessPlatforms instantiates a new conditionalAccessPlatforms and sets the default values.
func NewConditionalAccessPlatforms()(*ConditionalAccessPlatforms) {
    m := &ConditionalAccessPlatforms{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.conditionalAccessPlatforms";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateConditionalAccessPlatformsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessPlatformsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessPlatforms(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessPlatforms) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExcludePlatforms gets the excludePlatforms property value. Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue, linux.
func (m *ConditionalAccessPlatforms) GetExcludePlatforms()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludePlatforms
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessPlatforms) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["excludePlatforms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludePlatforms(res)
        }
        return nil
    }
    res["includePlatforms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludePlatforms(res)
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
// GetIncludePlatforms gets the includePlatforms property value. Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue,linux``.
func (m *ConditionalAccessPlatforms) GetIncludePlatforms()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includePlatforms
    }
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConditionalAccessPlatforms) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// Serialize serializes information the current object
func (m *ConditionalAccessPlatforms) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExcludePlatforms() != nil {
        err := writer.WriteCollectionOfStringValues("excludePlatforms", m.GetExcludePlatforms())
        if err != nil {
            return err
        }
    }
    if m.GetIncludePlatforms() != nil {
        err := writer.WriteCollectionOfStringValues("includePlatforms", m.GetIncludePlatforms())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessPlatforms) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExcludePlatforms sets the excludePlatforms property value. Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue, linux.
func (m *ConditionalAccessPlatforms) SetExcludePlatforms(value []string)() {
    if m != nil {
        m.excludePlatforms = value
    }
}
// SetIncludePlatforms sets the includePlatforms property value. Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue,linux``.
func (m *ConditionalAccessPlatforms) SetIncludePlatforms(value []string)() {
    if m != nil {
        m.includePlatforms = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessPlatforms) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
