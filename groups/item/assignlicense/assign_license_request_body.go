package assignlicense

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// AssignLicenseRequestBody 
type AssignLicenseRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    addLicenses []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AssignedLicense;
    // 
    removeLicenses []string;
}
// NewAssignLicenseRequestBody instantiates a new assignLicenseRequestBody and sets the default values.
func NewAssignLicenseRequestBody()(*AssignLicenseRequestBody) {
    m := &AssignLicenseRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignLicenseRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAddLicenses gets the addLicenses property value. 
func (m *AssignLicenseRequestBody) GetAddLicenses()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AssignedLicense) {
    if m == nil {
        return nil
    } else {
        return m.addLicenses
    }
}
// GetRemoveLicenses gets the removeLicenses property value. 
func (m *AssignLicenseRequestBody) GetRemoveLicenses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.removeLicenses
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignLicenseRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["addLicenses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewAssignedLicense() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AssignedLicense, len(val))
            for i, v := range val {
                res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AssignedLicense))
            }
            m.SetAddLicenses(res)
        }
        return nil
    }
    res["removeLicenses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRemoveLicenses(res)
        }
        return nil
    }
    return res
}
func (m *AssignLicenseRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AssignLicenseRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAddLicenses()))
        for i, v := range m.GetAddLicenses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("addLicenses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("removeLicenses", m.GetRemoveLicenses())
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
func (m *AssignLicenseRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAddLicenses sets the addLicenses property value. 
func (m *AssignLicenseRequestBody) SetAddLicenses(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AssignedLicense)() {
    if m != nil {
        m.addLicenses = value
    }
}
// SetRemoveLicenses sets the removeLicenses property value. 
func (m *AssignLicenseRequestBody) SetRemoveLicenses(value []string)() {
    if m != nil {
        m.removeLicenses = value
    }
}
