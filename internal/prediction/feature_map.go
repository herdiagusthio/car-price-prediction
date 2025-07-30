package prediction

// ModelInputSize is the number of features expected by the ONNX model.
const ModelInputSize = 64

// featureIndexMap maps feature names to their corresponding indices in the model input tensor.
// This map is used during preprocessing to ensure each feature is placed at the correct position.
var featureIndexMap = map[string]int{
	// Numerical features
	"symboling":        0,
	"wheelbase":        1,
	"carlength":        2,
	"carwidth":         3,
	"carheight":        4,
	"curbweight":       5,
	"enginesize":       6,
	"boreratio":        7,
	"stroke":           8,
	"compressionratio": 9,
	"horsepower":       10,
	"peakrpm":          11,
	"citympg":          12,
	"highwaympg":       13,

	// Categorical features (one-hot encoded)
	// Fuel type
	"fueltype_gas": 14,

	// Aspiration
	"aspiration_turbo": 15,

	// Door number
	"doornumber_two": 16,

	// Car body
	"carbody_hardtop":   17,
	"carbody_hatchback": 18,
	"carbody_sedan":     19,
	"carbody_wagon":     20,

	// Drive wheel
	"drivewheel_fwd": 21,
	"drivewheel_rwd": 22,

	// Engine location
	"enginelocation_rear": 23,

	// Engine type
	"enginetype_dohcv": 24,
	"enginetype_l":     25,
	"enginetype_ohc":   26,
	"enginetype_ohcf":  27,
	"enginetype_ohcv":  28,
	"enginetype_rotor": 29,

	// Cylinder number
	"cylindernumber_five":   30,
	"cylindernumber_four":   31,
	"cylindernumber_six":    32,
	"cylindernumber_three":  33,
	"cylindernumber_twelve": 34,
	"cylindernumber_two":    35,

	// Fuel system
	"fuelsystem_2bbl": 36,
	"fuelsystem_4bbl": 37,
	"fuelsystem_idi":  38,
	"fuelsystem_mfi":  39,
	"fuelsystem_mpfi": 40,
	"fuelsystem_spdi": 41,
	"fuelsystem_spfi": 42,

	// Brand
	"brand_audi":       43,
	"brand_bmw":        44,
	"brand_buick":      45,
	"brand_chevrolet":  46,
	"brand_dodge":      47,
	"brand_honda":      48,
	"brand_isuzu":      49,
	"brand_jaguar":     50,
	"brand_mazda":      51,
	"brand_mercury":    52,
	"brand_mitsubishi": 53,
	"brand_nissan":     54,
	"brand_peugeot":    55,
	"brand_plymouth":   56,
	"brand_porsche":    57,
	"brand_renault":    58,
	"brand_saab":       59,
	"brand_subaru":     60,
	"brand_toyota":     61,
	"brand_volkswagen": 62,
	"brand_volvo":      63,
}