package pkg

import (
	"testing"
)

func TestLevelIndex(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*Base)
		expected int
	}{
		{
			name: "default level index",
			setup: func(b *Base) {
				b.levelIndex = 0
			},
			expected: 0,
		},
		{
			name: "custom level index",
			setup: func(b *Base) {
				b.levelIndex = 5
			},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Base{}
			tt.setup(b)
			if got := b.LevelIndex(); got != tt.expected {
				t.Errorf("LevelIndex() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestChapterIndex(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*Base)
		expected int
	}{
		{
			name: "default chapter index",
			setup: func(b *Base) {
				b.chapterIndex = 0
			},
			expected: 0,
		},
		{
			name: "custom chapter index",
			setup: func(b *Base) {
				b.chapterIndex = 3
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Base{}
			tt.setup(b)
			if got := b.ChapterIndex(); got != tt.expected {
				t.Errorf("ChapterIndex() = %v, want %v", got, tt.expected)
			}
		})
	}
}
