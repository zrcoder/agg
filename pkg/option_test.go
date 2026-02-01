package pkg

import (
	"testing"
)

func TestWithScene(t *testing.T) {
	opt := WithScene("test-scene", func() any { return nil })

	b := &Base{}
	opt(b)

	if b.sceneName != "test-scene" {
		t.Errorf("WithScene() set sceneName = %v, want test-scene", b.sceneName)
	}
	if b.wsPath != "/ws/test-scene" {
		t.Errorf("WithScene() set wsPath = %v, want /ws/test-scene", b.wsPath)
	}
	if b.sceneFn == nil {
		t.Error("WithScene() did not set sceneFn")
	}
}

func TestWithLevels(t *testing.T) {
	resetCalled := false
	opt := WithLevels([]Level{
		{Label: "Level 1", Data: map[string]int{"disks": 3}},
		{Label: "Level 2", Data: map[string]int{"disks": 4}},
		{Label: "Level 3", Data: map[string]int{"disks": 5}},
	}, func() {
		resetCalled = true
	})

	b := &Base{}
	opt(b)

	if len(b.levels) != 3 {
		t.Errorf("WithLevels() set len(levels) = %v, want 3", len(b.levels))
	}
	if b.levels[0].Label != "Level 1" {
		t.Errorf("WithLevels() set level[0].Label = %v, want Level 1", b.levels[0].Label)
	}
	if b.levels[0].Value != 0 {
		t.Errorf("WithLevels() set level[0].Value = %v, want 0", b.levels[0].Value)
	}
	if b.levels[1].Value != 1 {
		t.Errorf("WithLevels() set level[1].Value = %v, want 1", b.levels[1].Value)
	}
	if len(b.levelOptions) != 3 {
		t.Errorf("WithLevels() set len(levelOptions) = %v, want 3", len(b.levelOptions))
	}
	if b.reset == nil {
		t.Error("WithLevels() did not set reset")
	}

	b.reset()
	if !resetCalled {
		t.Error("reset function was not called")
	}
}

func TestWithChapters(t *testing.T) {
	resetCalled := false
	opt := WithChapters([]Chapter{
		{
			Label: "Chapter 1",
			Children: []Level{
				{Label: "Level 1", Data: map[string]int{"disks": 3}},
				{Label: "Level 2", Data: map[string]int{"disks": 4}},
			},
		},
		{
			Label: "Chapter 2",
			Children: []Level{
				{Label: "Level 1", Data: map[string]int{"disks": 5}},
				{Label: "Level 2", Data: map[string]int{"disks": 6}},
			},
		},
	}, func() {
		resetCalled = true
	})

	b := &Base{}
	opt(b)

	if len(b.chapters) != 2 {
		t.Errorf("WithChapters() set len(chapters) = %v, want 2", len(b.chapters))
	}
	if b.chapters[0].Label != "Chapter 1" {
		t.Errorf("WithChapters() set chapter[0].Label = %v, want Chapter 1", b.chapters[0].Label)
	}
	if len(b.chapters[0].Children) != 2 {
		t.Errorf("WithChapters() set len(chapter[0].Children) = %v, want 2", len(b.chapters[0].Children))
	}
	if b.chapters[0].Children[0].Value != (LevelMeta{Chapter: 0, Level: 0}) {
		t.Errorf("WithChapters() set chapter[0].Children[0].Value = %v, want {0, 0}", b.chapters[0].Children[0].Value)
	}
	if b.chapters[1].Children[1].Value != (LevelMeta{Chapter: 1, Level: 1}) {
		t.Errorf("WithChapters() set chapter[1].Children[1].Value = %v, want {1, 1}", b.chapters[1].Children[1].Value)
	}
	if len(b.chapterOptions) != 2 {
		t.Errorf("WithChapters() set len(chapterOptions) = %v, want 2", len(b.chapterOptions))
	}
	if b.reset == nil {
		t.Error("WithChapters() did not set reset")
	}

	b.reset()
	if !resetCalled {
		t.Error("reset function was not called")
	}
}
