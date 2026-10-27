using DynamicQuantities
using DynamicQuantities.Units
import DynamicQuantities.Constants as C

MMCO2 = 44.08u"g/mol"
liquid_CO2_concentration = 0.143u"mol/L"

function henrys_constant(temperature)
	return 3.4e-2u"mol/L/Constants.atm" * exp(2400u"K" * (1/temperature - 1/(298.15u"K")))
end

function bottle_CO2_required(temperature, liquid_volume, gas_volume)
	return MMCO2*(
		liquid_CO2_concentration*liquid_volume +
		liquid_CO2_concentration*gas_volume/henrys_constant(temperature)/C.R/temperature
	)
end

function reactants_required(temperature, liquid_volume, gas_volume)
	factor = bottle_CO2_required(temperature, liquid_volume, gas_volume)/132.03u"g"
	return Dict(
		"acetic" => factor*192.12u"g",
		"bicarb" => factor*252.03u"g",
	)
end

required = reactants_required(10ua"°C", 500u"mL", 600u"mL")
println("Acetic: ", required["acetic"] |> us"g")
println("Bicarb: ", required["bicarb"] |> us"g")
