package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemPreviewInfo 
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
// NewItemPreviewInfo instantiates a new itemPreviewInfo and sets the default values.
func NewItemPreviewInfo()(*ItemPreviewInfo) {
    m := &ItemPreviewInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemPreviewInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetGetUrl gets the getUrl property value. 
func (m *ItemPreviewInfo) GetGetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.getUrl
    }
}
// GetPostParameters gets the postParameters property value. 
func (m *ItemPreviewInfo) GetPostParameters()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postParameters
    }
}
// GetPostUrl gets the postUrl property value. 
func (m *ItemPreviewInfo) GetPostUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemPreviewInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["getUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGetUrl(val)
        }
        return nil
    }
    res["postParameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostParameters(val)
        }
        return nil
    }
    res["postUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostUrl(val)
        }
        return nil
    }
    return res
}
func (m *ItemPreviewInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemPreviewInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetGetUrl sets the getUrl property value. 
func (m *ItemPreviewInfo) SetGetUrl(value *string)() {
    m.getUrl = value
}
// SetPostParameters sets the postParameters property value. 
func (m *ItemPreviewInfo) SetPostParameters(value *string)() {
    m.postParameters = value
}
// SetPostUrl sets the postUrl property value. 
func (m *ItemPreviewInfo) SetPostUrl(value *string)() {
    m.postUrl = value
}
