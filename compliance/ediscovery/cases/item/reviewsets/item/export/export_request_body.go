package export

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/ediscovery"
)

// 
type ExportRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    azureBlobContainer *string;
    // 
    azureBlobToken *string;
    // 
    description *string;
    // 
    exportOptions *i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.ExportOptions;
    // 
    exportStructure *i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.ExportFileStructure;
    // 
    outputName *string;
}
// Instantiates a new exportRequestBody and sets the default values.
func NewExportRequestBody()(*ExportRequestBody) {
    m := &ExportRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExportRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the azureBlobContainer property value. 
func (m *ExportRequestBody) GetAzureBlobContainer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureBlobContainer
    }
}
// Gets the azureBlobToken property value. 
func (m *ExportRequestBody) GetAzureBlobToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureBlobToken
    }
}
// Gets the description property value. 
func (m *ExportRequestBody) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the exportOptions property value. 
func (m *ExportRequestBody) GetExportOptions()(*i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.ExportOptions) {
    if m == nil {
        return nil
    } else {
        return m.exportOptions
    }
}
// Gets the exportStructure property value. 
func (m *ExportRequestBody) GetExportStructure()(*i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.ExportFileStructure) {
    if m == nil {
        return nil
    } else {
        return m.exportStructure
    }
}
// Gets the outputName property value. 
func (m *ExportRequestBody) GetOutputName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.outputName
    }
}
// The deserialization information for the current model
func (m *ExportRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["azureBlobContainer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureBlobContainer(val)
        }
        return nil
    }
    res["azureBlobToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureBlobToken(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["exportOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.ParseExportOptions)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.ExportOptions)
            m.SetExportOptions(&cast)
        }
        return nil
    }
    res["exportStructure"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.ParseExportFileStructure)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.ExportFileStructure)
            m.SetExportStructure(&cast)
        }
        return nil
    }
    res["outputName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutputName(val)
        }
        return nil
    }
    return res
}
func (m *ExportRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ExportRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("azureBlobContainer", m.GetAzureBlobContainer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("azureBlobToken", m.GetAzureBlobToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetExportOptions() != nil {
        cast := m.GetExportOptions().String()
        err := writer.WriteStringValue("exportOptions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExportStructure() != nil {
        cast := m.GetExportStructure().String()
        err := writer.WriteStringValue("exportStructure", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("outputName", m.GetOutputName())
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
func (m *ExportRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the azureBlobContainer property value. 
// Parameters:
//  - value : Value to set for the azureBlobContainer property.
func (m *ExportRequestBody) SetAzureBlobContainer(value *string)() {
    m.azureBlobContainer = value
}
// Sets the azureBlobToken property value. 
// Parameters:
//  - value : Value to set for the azureBlobToken property.
func (m *ExportRequestBody) SetAzureBlobToken(value *string)() {
    m.azureBlobToken = value
}
// Sets the description property value. 
// Parameters:
//  - value : Value to set for the description property.
func (m *ExportRequestBody) SetDescription(value *string)() {
    m.description = value
}
// Sets the exportOptions property value. 
// Parameters:
//  - value : Value to set for the exportOptions property.
func (m *ExportRequestBody) SetExportOptions(value *i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.ExportOptions)() {
    m.exportOptions = value
}
// Sets the exportStructure property value. 
// Parameters:
//  - value : Value to set for the exportStructure property.
func (m *ExportRequestBody) SetExportStructure(value *i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.ExportFileStructure)() {
    m.exportStructure = value
}
// Sets the outputName property value. 
// Parameters:
//  - value : Value to set for the outputName property.
func (m *ExportRequestBody) SetOutputName(value *string)() {
    m.outputName = value
}
