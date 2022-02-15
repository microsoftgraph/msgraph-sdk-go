package internalsponsors

import (
    id548309b7f062162722aee435eac927347015bc798cf5b757340af2dc3770b50 "github.com/microsoftgraph/msgraph-sdk-go/devices/getbyids"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InternalSponsorsResponse 
type InternalSponsorsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    nextLink *string;
    // 
    value []id548309b7f062162722aee435eac927347015bc798cf5b757340af2dc3770b50.DirectoryObject;
}
// NewInternalSponsorsResponse instantiates a new internalSponsorsResponse and sets the default values.
func NewInternalSponsorsResponse()(*InternalSponsorsResponse) {
    m := &InternalSponsorsResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InternalSponsorsResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetNextLink gets the @odata.nextLink property value. 
func (m *InternalSponsorsResponse) GetNextLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nextLink
    }
}
// GetValue gets the value property value. 
func (m *InternalSponsorsResponse) GetValue()([]id548309b7f062162722aee435eac927347015bc798cf5b757340af2dc3770b50.DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InternalSponsorsResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["@odata.nextLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextLink(val)
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return id548309b7f062162722aee435eac927347015bc798cf5b757340af2dc3770b50.NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]id548309b7f062162722aee435eac927347015bc798cf5b757340af2dc3770b50.DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*id548309b7f062162722aee435eac927347015bc798cf5b757340af2dc3770b50.DirectoryObject))
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
func (m *InternalSponsorsResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *InternalSponsorsResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.nextLink", m.GetNextLink())
        if err != nil {
            return err
        }
    }
    if m.GetValue() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("value", cast)
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
func (m *InternalSponsorsResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetNextLink sets the @odata.nextLink property value. 
func (m *InternalSponsorsResponse) SetNextLink(value *string)() {
    if m != nil {
        m.nextLink = value
    }
}
// SetValue sets the value property value. 
func (m *InternalSponsorsResponse) SetValue(value []id548309b7f062162722aee435eac927347015bc798cf5b757340af2dc3770b50.DirectoryObject)() {
    if m != nil {
        m.value = value
    }
}
