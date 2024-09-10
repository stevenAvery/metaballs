package core

import (
	"unsafe"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func SetUniformFloat32(shader rl.Shader, uniformName string, value float32) {
	uniformLocation := rl.GetShaderLocation(shader, uniformName)
	rl.SetShaderValue(shader, uniformLocation, []float32{value}, rl.ShaderUniformFloat)
}

func SetUniformInt32(shader rl.Shader, uniformName string, value int32) {
	uniformLocation := rl.GetShaderLocation(shader, uniformName)
	uniformValue := unsafe.Slice((*float32)(unsafe.Pointer(&value)), 4)
	rl.SetShaderValue(shader, uniformLocation, uniformValue, rl.ShaderUniformInt)
}

func SetUniformInt(shader rl.Shader, uniformName string, value int) {
	SetUniformInt32(shader, uniformName, int32(value))
}

func SetUniformVec2(shader rl.Shader, uniformName string, value rl.Vector2) {
	uniformLocation := rl.GetShaderLocation(shader, uniformName)
	uniformValue := []float32{value.X, value.Y}
	rl.SetShaderValue(shader, uniformLocation, uniformValue, rl.ShaderUniformVec2)
}

func SetUniformVec3(shader rl.Shader, uniformName string, value rl.Vector3) {
	uniformLocation := rl.GetShaderLocation(shader, uniformName)
	uniformValue := []float32{value.X, value.Y, value.Z}
	rl.SetShaderValue(shader, uniformLocation, uniformValue, rl.ShaderUniformVec3)
}

func SetUniformVec4(shader rl.Shader, uniformName string, value rl.Vector4) {
	uniformLocation := rl.GetShaderLocation(shader, uniformName)
	uniformValue := []float32{value.X, value.Y, value.Z, value.W}
	rl.SetShaderValue(shader, uniformLocation, uniformValue, rl.ShaderUniformVec4)
}

func SetUniformVec2Arr(shader rl.Shader, uniformName string, value []rl.Vector2) {
	uniformLocation := rl.GetShaderLocation(shader, uniformName)
	count := len(value)
	// Flatten vec2 array into a single array of float32s
	uniformValue := make([]float32, count*2)
	for i, v := range value {
		uniformValue[i*2] = v.X
		uniformValue[i*2+1] = v.Y
	}

	rl.SetShaderValueV(shader, uniformLocation, uniformValue, rl.ShaderUniformVec2, int32(count))
}

func SetUniformVec3Arr(shader rl.Shader, uniformName string, value []rl.Vector3) {
	uniformLocation := rl.GetShaderLocation(shader, uniformName)
	count := len(value)
	// Flatten vec3 array into a single array of float32s
	uniformValue := make([]float32, count*3)
	for i, v := range value {
		uniformValue[i*3] = v.X
		uniformValue[i*3+1] = v.Y
		uniformValue[i*3+2] = v.Z
	}

	rl.SetShaderValueV(shader, uniformLocation, uniformValue, rl.ShaderUniformVec3, int32(count))
}

func SetUniformVec4Arr(shader rl.Shader, uniformName string, value []rl.Vector4) {
	uniformLocation := rl.GetShaderLocation(shader, uniformName)
	count := len(value)
	// Flatten vec4 array into a single array of float32s
	uniformValue := make([]float32, count*4)
	for i, v := range value {
		uniformValue[i*4] = v.X
		uniformValue[i*4+1] = v.Y
		uniformValue[i*4+2] = v.Z
		uniformValue[i*4+3] = v.W
	}

	rl.SetShaderValueV(shader, uniformLocation, uniformValue, rl.ShaderUniformVec4, int32(count))
}

func SetUniformColour(shader rl.Shader, uniformName string, colour rl.Color) {
	SetUniformVec3(shader, uniformName, rl.Vector3{
		X: float32(colour.R) / 255,
		Y: float32(colour.G) / 255,
		Z: float32(colour.B) / 255,
	})
}
