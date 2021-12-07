package getmailtips

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// GetMailTipsRequestBody 
type GetMailTipsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    emailAddresses []string;
    // 
    mailTipsOptions *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailTipsType;
}
// NewGetMailTipsRequestBody instantiates a new getMailTipsRequestBody and sets the default values.
func NewGetMailTipsRequestBody()(*GetMailTipsRequestBody) {
    m := &GetMailTipsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetMailTipsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEmailAddresses gets the emailAddresses property value. 
func (m *GetMailTipsRequestBody) GetEmailAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.emailAddresses
    }
}
// GetMailTipsOptions gets the mailTipsOptions property value. 
func (m *GetMailTipsRequestBody) GetMailTipsOptions()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailTipsType) {
    if m == nil {
        return nil
    } else {
        return m.mailTipsOptions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetMailTipsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["emailAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEmailAddresses(res)
        }
        return nil
    }
    res["mailTipsOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ParseMailTipsType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailTipsType)
            m.SetMailTipsOptions(&cast)
        }
        return nil
    }
    return res
}
func (m *GetMailTipsRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetMailTipsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("emailAddresses", m.GetEmailAddresses())
        if err != nil {
            return err
        }
    }
    if m.GetMailTipsOptions() != nil {
        cast := m.GetMailTipsOptions().String()
        err := writer.WriteStringValue("mailTipsOptions", &cast)
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
func (m *GetMailTipsRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEmailAddresses sets the emailAddresses property value. 
func (m *GetMailTipsRequestBody) SetEmailAddresses(value []string)() {
    if m != nil {
        m.emailAddresses = value
    }
}
// SetMailTipsOptions sets the mailTipsOptions property value. 
func (m *GetMailTipsRequestBody) SetMailTipsOptions(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailTipsType)() {
    if m != nil {
        m.mailTipsOptions = value
    }
}
