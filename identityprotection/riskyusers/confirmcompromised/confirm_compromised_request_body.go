package confirmcompromised

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConfirmCompromisedRequestBody provides operations to call the confirmCompromised method.
type ConfirmCompromisedRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    userIds []string;
}
// NewConfirmCompromisedRequestBody instantiates a new confirmCompromisedRequestBody and sets the default values.
func NewConfirmCompromisedRequestBody()(*ConfirmCompromisedRequestBody) {
    m := &ConfirmCompromisedRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConfirmCompromisedRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConfirmCompromisedRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewConfirmCompromisedRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfirmCompromisedRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfirmCompromisedRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["userIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUserIds(res)
        }
        return nil
    }
    return res
}
// GetUserIds gets the userIds property value. 
func (m *ConfirmCompromisedRequestBody) GetUserIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.userIds
    }
}
// Serialize serializes information the current object
func (m *ConfirmCompromisedRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetUserIds() != nil {
        err := writer.WriteCollectionOfStringValues("userIds", m.GetUserIds())
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
func (m *ConfirmCompromisedRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUserIds sets the userIds property value. 
func (m *ConfirmCompromisedRequestBody) SetUserIds(value []string)() {
    if m != nil {
        m.userIds = value
    }
}
