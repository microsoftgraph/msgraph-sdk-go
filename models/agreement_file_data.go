package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AgreementFileData 
type AgreementFileData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Data that represents the terms of use PDF document. Read-only. Note: You can use the .NET Convert.ToBase64String method to convert your file to binary data for uploading using the Create agreements API. A sample syntax using this method in PowerShell is [convert]::ToBase64String((Get-Content -path 'your_file_path' -Encoding byte)).
    data []byte
    // The OdataType property
    odataType *string
}
// NewAgreementFileData instantiates a new agreementFileData and sets the default values.
func NewAgreementFileData()(*AgreementFileData) {
    m := &AgreementFileData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.agreementFileData";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAgreementFileDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAgreementFileDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAgreementFileData(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AgreementFileData) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetData gets the data property value. Data that represents the terms of use PDF document. Read-only. Note: You can use the .NET Convert.ToBase64String method to convert your file to binary data for uploading using the Create agreements API. A sample syntax using this method in PowerShell is [convert]::ToBase64String((Get-Content -path 'your_file_path' -Encoding byte)).
func (m *AgreementFileData) GetData()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.data
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AgreementFileData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetData(val)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AgreementFileData) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// Serialize serializes information the current object
func (m *AgreementFileData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("data", m.GetData())
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
func (m *AgreementFileData) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetData sets the data property value. Data that represents the terms of use PDF document. Read-only. Note: You can use the .NET Convert.ToBase64String method to convert your file to binary data for uploading using the Create agreements API. A sample syntax using this method in PowerShell is [convert]::ToBase64String((Get-Content -path 'your_file_path' -Encoding byte)).
func (m *AgreementFileData) SetData(value []byte)() {
    if m != nil {
        m.data = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AgreementFileData) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
