package favorite

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// FavoriteRequestBody provides operations to call the favorite method.
type FavoriteRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    messageIds []string;
}
// NewFavoriteRequestBody instantiates a new favoriteRequestBody and sets the default values.
func NewFavoriteRequestBody()(*FavoriteRequestBody) {
    m := &FavoriteRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFavoriteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFavoriteRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewFavoriteRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FavoriteRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FavoriteRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["messageIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMessageIds(res)
        }
        return nil
    }
    return res
}
// GetMessageIds gets the messageIds property value. 
func (m *FavoriteRequestBody) GetMessageIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.messageIds
    }
}
func (m *FavoriteRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *FavoriteRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetMessageIds() != nil {
        err := writer.WriteCollectionOfStringValues("messageIds", m.GetMessageIds())
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
func (m *FavoriteRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMessageIds sets the messageIds property value. 
func (m *FavoriteRequestBody) SetMessageIds(value []string)() {
    if m != nil {
        m.messageIds = value
    }
}
