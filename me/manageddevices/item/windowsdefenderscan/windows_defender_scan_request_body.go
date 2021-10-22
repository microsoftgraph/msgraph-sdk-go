package windowsdefenderscan

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsDefenderScanRequestBody struct {
    additionalData map[string]interface{};
    quickScan *bool;
}
func NewWindowsDefenderScanRequestBody()(*WindowsDefenderScanRequestBody) {
    m := &WindowsDefenderScanRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *WindowsDefenderScanRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *WindowsDefenderScanRequestBody) GetQuickScan()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.quickScan
    }
}
func (m *WindowsDefenderScanRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["quickScan"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetQuickScan(val)
        return nil
    }
    return res
}
func (m *WindowsDefenderScanRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *WindowsDefenderScanRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("quickScan", m.GetQuickScan())
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
func (m *WindowsDefenderScanRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *WindowsDefenderScanRequestBody) SetQuickScan(value *bool)() {
    m.quickScan = value
}
