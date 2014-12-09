package accessory

import (
    "github.com/brutella/hap/model"
    
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestThermostat(t *testing.T) {
    info := model.Info{
        Name: "My Thermostat",
        SerialNumber: "001",
        Manufacturer: "Google",
        Model: "Thermostaty",
    }
    
    var thermo model.Thermostat = NewThermostat(info, 10, 0, 100, 1)
    
    assert.Equal(t, thermo.GetId(), model.InvalidId)    
    assert.Equal(t, thermo.Name(), "My Thermostat")
    assert.Equal(t, thermo.SerialNumber(), "001")
    assert.Equal(t, thermo.Manufacturer(), "Google")
    assert.Equal(t, thermo.Model(), "Thermostaty")
    assert.Equal(t, thermo.Temperature(), 10)
    assert.Equal(t, thermo.TargetTemperature(), 10)
    assert.Equal(t, thermo.TargetMode(), model.ModeOff)
    assert.Equal(t, thermo.Mode(), model.ModeOff)
    
    thermo.SetTemperature(11)
    thermo.SetTargetTemperature(12)
    
    assert.Equal(t, thermo.Temperature(), 11)
    assert.Equal(t, thermo.TargetTemperature(), 12)
}