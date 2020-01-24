package mypkg

import "testing"

func TestGetWord(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "case 1",
			want: "word",
		},
		{
			name: "case 2",
			want: "word",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWord(); got != tt.want {
				t.Errorf("GetWord() = %v, want %v", got, tt.want)
			}
		})
	}
}