package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	ort "github.com/yalue/onnxruntime_go"
)

// getSharedLibPath returns the path to the ONNX runtime shared library based on the OS and architecture
func getSharedLibPath() string {
	// Use absolute path to ensure the library can be found
	basePath := "/Users/bytedance/Documents/GitHub/car-price-prediction/onnxruntime_go_examples/third_party/"
	if runtime.GOOS == "windows" {
		if runtime.GOARCH == "amd64" {
			return basePath + "onnxruntime.dll"
		}
	}
	if runtime.GOOS == "darwin" {
		if runtime.GOARCH == "arm64" {
			return basePath + "onnxruntime_arm64.dylib"
		}
		if runtime.GOARCH == "amd64" {
			return basePath + "onnxruntime_amd64.dylib"
		}
	}
	if runtime.GOOS == "linux" {
		if runtime.GOARCH == "arm64" {
			return basePath + "onnxruntime_arm64.so"
		}
		return basePath + "onnxruntime.so"
	}
	log.Fatalf("Unable to determine a path to the onnxruntime shared library for OS \"%s\" and architecture \"%s\".", runtime.GOOS, runtime.GOARCH)
	return ""
}

// showNetworkInputsAndOutputs prints the inputs and outputs of an ONNX model
func showNetworkInputsAndOutputs(libPath, networkPath string) error {
	ort.SetSharedLibraryPath(libPath)
	e := ort.InitializeEnvironment()
	if e != nil {
		return fmt.Errorf("Error initializing onnxruntime library: %w", e)
	}
	defer ort.DestroyEnvironment()

	inputs, outputs, e := ort.GetInputOutputInfo(networkPath)
	if e != nil {
		return fmt.Errorf("Error getting input and output info for %s: %w", networkPath, e)
	}

	fmt.Printf("%d inputs to %s:\n", len(inputs), networkPath)
	for i, v := range inputs {
		fmt.Printf("  Index %d: %s\n", i, &v)
	}

	fmt.Printf("%d outputs from %s:\n", len(outputs), networkPath)
	for i, v := range outputs {
		fmt.Printf("  Index %d: %s\n", i, &v)
	}

	return nil
}

func main() {
	// Define the model path with absolute path
	modelPath := "/Users/bytedance/Documents/GitHub/car-price-prediction/model/best_model.onnx"

	// Get the path to the ONNX runtime shared library
	libPath := getSharedLibPath()

	// Show the inputs and outputs of the model
	err := showNetworkInputsAndOutputs(libPath, modelPath)
	if err != nil {
		log.Fatalf("Error showing network inputs and outputs: %v", err)
	}

	os.Exit(0)
}