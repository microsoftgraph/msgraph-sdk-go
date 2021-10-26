package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ItemPreviewInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    getUrl *string;
    // 
    postParameters *string;
    // 
    postUrl *string;
}
// Instantiates a new itemPreviewInfo and sets the default values.
func NewItemPreviewInfo()(*ItemPreviewInfo) {
    m := &ItemPreviewInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemPreviewInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the getUrl property value. 
func (m *ItemPreviewInfo) GetGetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.getUrl
    }
}
// Gets the postParameters property value. 
func (m *ItemPreviewInfo) GetPostParameters()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postParameters
    }
}
// Gets the postUrl property value. 
func (m *ItemPreviewInfo) GetPostUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postUrl
    }
}
// The deserialization information for the current model
func (m *ItemPreviewInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["getUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGetUrl(val)
        return nil
    }
    res["postParameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPostParameters(val)
        return nil
    }
    res["postUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPostUrl(val)
        return nil
    }
    return res
}
func (m *ItemPreviewInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ItemPreviewInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("getUrl", m.GetGetUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("postParameters", m.GetPostParameters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("postUrl", m.GetPostUrl())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ItemPreviewInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the getUrl property value. 
// Parameters:
//  - value : Value to set for the getUrl property.
func (m *ItemPreviewInfo) SetGetUrl(value *string)() {
    m.getUrl = value
}
// Sets the postParameters property value. 
// Parameters:
//  - value : Value to set for the postParameters property.
func (m *ItemPreviewInfo) SetPostParameters(value *string)() {
    m.postParameters = value
}
// Sets the postUrl property value. 
// Parameters:
//  - value : Value to set for the postUrl property.
func (m *ItemPreviewInfo) SetPostUrl(value *string)() {
    m.postUrl = value
}
