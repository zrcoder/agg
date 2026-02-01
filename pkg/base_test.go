package pkg

import (
	"testing"

	"github.com/zrcoder/amisgo"
)

func TestNewWithOptions(t *testing.T) {
	app := amisgo.New()
	sceneCalled := false
	resetCalled := false

	opts := []Option{
		WithScene("test-scene", func() any {
			sceneCalled = true
			return "test"
		}),
		WithLevels([]Level{
			{Label: "Level 1"},
			{Label: "Level 2"},
		}, func() {
			resetCalled = true
		}),
	}

	base := New(app, opts...)

	if base.sceneName != "test-scene" {
		t.Errorf("WithScene() did not set scene name, got %v", base.sceneName)
	}
	if base.wsPath != "/ws/test-scene" {
		t.Errorf("WithScene() did not set ws path, got %v", base.wsPath)
	}
	if len(base.levels) != 2 {
		t.Errorf("WithLevels() did not set levels, got %v", len(base.levels))
	}
	if base.sceneFn == nil {
		t.Error("WithScene() did not set scene function")
	}
	if base.reset == nil {
		t.Error("WithLevels() did not set reset function")
	}

	base.sceneFn()
	if !sceneCalled {
		t.Error("scene function was not called")
	}

	base.reset()
	if !resetCalled {
		t.Error("reset function was not called")
	}
}
