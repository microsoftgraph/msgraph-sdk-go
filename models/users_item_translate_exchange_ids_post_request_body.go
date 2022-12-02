package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemTranslateExchangeIdsPostRequestBody provides operations to call the translateExchangeIds method.
type UsersItemTranslateExchangeIdsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The InputIds property
    inputIds []string
    // The SourceIdType property
    sourceIdType *ExchangeIdFormat
    // The TargetIdType property
    targetIdType *ExchangeIdFormat
}
// NewUsersItemTranslateExchangeIdsPostRequestBody instantiates a new UsersItemTranslateExchangeIdsPostRequestBody and sets the default values.
func NewUsersItemTranslateExchangeIdsPostRequestBody()(*UsersItemTranslateExchangeIdsPostRequestBody) {
    m := &UsersItemTranslateExchangeIdsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemTranslateExchangeIdsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemTranslateExchangeIdsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemTranslateExchangeIdsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemTranslateExchangeIdsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemTranslateExchangeIdsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["inputIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetInputIds(res)
        }
        return nil
    }
    res["sourceIdType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExchangeIdFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceIdType(val.(*ExchangeIdFormat))
        }
        return nil
    }
    res["targetIdType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExchangeIdFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetIdType(val.(*ExchangeIdFormat))
        }
        return nil
    }
    return res
}
// GetInputIds gets the inputIds property value. The InputIds property
func (m *UsersItemTranslateExchangeIdsPostRequestBody) GetInputIds()([]string) {
    return m.inputIds
}
// GetSourceIdType gets the sourceIdType property value. The SourceIdType property
func (m *UsersItemTranslateExchangeIdsPostRequestBody) GetSourceIdType()(*ExchangeIdFormat) {
    return m.sourceIdType
}
// GetTargetIdType gets the targetIdType property value. The TargetIdType property
func (m *UsersItemTranslateExchangeIdsPostRequestBody) GetTargetIdType()(*ExchangeIdFormat) {
    return m.targetIdType
}
// Serialize serializes information the current object
func (m *UsersItemTranslateExchangeIdsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInputIds() != nil {
        err := writer.WriteCollectionOfStringValues("inputIds", m.GetInputIds())
        if err != nil {
            return err
        }
    }
    if m.GetSourceIdType() != nil {
        cast := (*m.GetSourceIdType()).String()
        err := writer.WriteStringValue("sourceIdType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargetIdType() != nil {
        cast := (*m.GetTargetIdType()).String()
        err := writer.WriteStringValue("targetIdType", &cast)
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
func (m *UsersItemTranslateExchangeIdsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetInputIds sets the inputIds property value. The InputIds property
func (m *UsersItemTranslateExchangeIdsPostRequestBody) SetInputIds(value []string)() {
    m.inputIds = value
}
// SetSourceIdType sets the sourceIdType property value. The SourceIdType property
func (m *UsersItemTranslateExchangeIdsPostRequestBody) SetSourceIdType(value *ExchangeIdFormat)() {
    m.sourceIdType = value
}
// SetTargetIdType sets the targetIdType property value. The TargetIdType property
func (m *UsersItemTranslateExchangeIdsPostRequestBody) SetTargetIdType(value *ExchangeIdFormat)() {
    m.targetIdType = value
}
