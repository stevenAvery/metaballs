package state

import (
	"fmt"
	"math"
	"math/rand"
	"metaballs-demo/src/core"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TAU = math.Pi * 2

	LOOP_DUR_MS         = 10_000
	MAX_NUM_METABALLS   = 64
	MAX_METABALL_R      = 50.0
	MAX_CYCLES_PER_LOOP = 5
	MIN_DIST_PERCENT    = 0.75
	MAX_DIST_PERCENT    = 0.95
	BLOBINESS           = 0.5
)

type metaball struct {
	CyclesPerLoop  int32
	LoopOffset     float32
	Angle          float64
	MaxDistPercent float32
}

type metaballState struct {
	start        time.Time
	windowWidth  int
	windowHeight int
	debug        bool
	metaballs    []metaball

	shader       rl.Shader
	shaderTarget rl.RenderTexture2D
}

func NewMetaballs(width int, height int) metaballState {
	state := metaballState{}
	state.start = time.Now()
	state.windowWidth = width
	state.windowHeight = height
	state.debug = false

	state.metaballs = make([]metaball, MAX_NUM_METABALLS)
	for i := range state.metaballs {
		state.metaballs[i] = metaball{
			CyclesPerLoop:  rand.Int31n(MAX_CYCLES_PER_LOOP) + 1,
			LoopOffset:     rand.Float32(),
			Angle:          rand.Float64() * TAU,
			MaxDistPercent: MIN_DIST_PERCENT + rand.Float32()*(MAX_DIST_PERCENT-MIN_DIST_PERCENT),
		}
	}

	state.shader = rl.LoadShader("state/metaballs.vs", "state/metaballs.fs")
	state.shaderTarget = rl.LoadRenderTexture(int32(width), int32(height))

	return state
}

func (state *metaballState) Update() {
	// Set uniforms for shader
	windowSize := rl.NewVector2(float32(state.windowWidth), float32(state.windowHeight))
	centre := rl.NewVector2(windowSize.X/2.0, windowSize.Y/2.0)

	durMs := time.Now().Sub(state.start).Milliseconds() % LOOP_DUR_MS
	loopPercent := float32(durMs) / LOOP_DUR_MS // From 0.0 to 1.0
	minWindowDimension := float32(math.Min(float64(windowSize.X), float64(windowSize.Y)))
	maxR := minWindowDimension / 2.0

	metaballs := core.Map(state.metaballs, func(m metaball, _ int) rl.Vector3 {
		cyclePercent := core.Fract((loopPercent + m.LoopOffset) * float32(m.CyclesPerLoop))
		return rl.Vector3{
			X: centre.X + float32(math.Cos(m.Angle))*cyclePercent*maxR*m.MaxDistPercent,
			Y: centre.Y + float32(math.Sin(m.Angle))*cyclePercent*maxR*m.MaxDistPercent,
			Z: MAX_METABALL_R - cyclePercent*MAX_METABALL_R,
		}
	})

	core.SetUniformVec2(state.shader, "windowSize", windowSize)
	core.SetUniformVec3Arr(state.shader, "metaballs", metaballs)
	core.SetUniformInt(state.shader, "numMetaballs", len(metaballs))
	core.SetUniformColour(state.shader, "backgroundColour", rl.NewColor(0, 0, 255, 255))
	core.SetUniformColour(state.shader, "metaballColour", rl.NewColor(0, 255, 0, 255))
	core.SetUniformFloat32(state.shader, "blobiness", BLOBINESS)
}

func (state *metaballState) Draw() {
	// Draw metaball post processing shader
	rl.BeginShaderMode(state.shader)
	rl.DrawTextureRec(
		state.shaderTarget.Texture,
		rl.NewRectangle(0, 0, float32(state.shaderTarget.Texture.Width), float32(state.shaderTarget.Texture.Height)),
		rl.Vector2Zero(), rl.White,
	)
	rl.EndShaderMode()

	// Debugging text
	if state.debug {
		debugMessage := fmt.Sprintf(
			"State:  Metaballs\nFPS:    %d\nWindow: %dx%d\n",
			rl.GetFPS(), state.windowWidth, state.windowHeight,
		)
		rl.DrawText(debugMessage, 10, 10, 10, rl.White)
	}
}

func (state *metaballState) OnResize(width int, height int) {
	state.windowWidth = width
	state.windowHeight = height

	// Resize shader texture
	state.shaderTarget = rl.LoadRenderTexture(int32(width), int32(height))
}

func (state *metaballState) Unload() {
	rl.UnloadShader(state.shader)
	rl.UnloadRenderTexture(state.shaderTarget)
}
