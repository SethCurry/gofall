package gofall_test

import (
	"testing"

	"github.com/SethCurry/gofall"
)

func Test_Legality_UnmarshalText(t *testing.T) {
	testCases := []struct {
		name    string
		txt     []byte
		want    gofall.Legality
		wantErr bool
	}{
		{
			name:    "legal",
			txt:     []byte("legal"),
			want:    gofall.LegalityLegal,
			wantErr: false,
		},
		{
			name:    "not legal",
			txt:     []byte("not_legal"),
			want:    gofall.LegalityNotLegal,
			wantErr: false,
		},
		{
			name:    "unknown legality",
			txt:     []byte("unknown"),
			want:    gofall.Legality(""),
			wantErr: true,
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			var l gofall.Legality

			err := l.UnmarshalText(v.txt)
			if (err != nil) != v.wantErr {
				t.Errorf("unexpected error: %v", err)
			}

			if l != v.want {
				t.Errorf("unexpected legality: got %v, want %v", l, v.want)
			}
		})
	}
}

func Test_Legality_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		name    string
		txt     []byte
		want    gofall.Legality
		wantErr bool
	}{
		{
			name:    "legal",
			txt:     []byte("\"legal\""),
			want:    gofall.LegalityLegal,
			wantErr: false,
		},
		{
			name:    "not legal",
			txt:     []byte("\"not_legal\""),
			want:    gofall.LegalityNotLegal,
			wantErr: false,
		},
		{
			name:    "unknown legality",
			txt:     []byte("\"unknown\""),
			want:    gofall.Legality(""),
			wantErr: true,
		},
		{
			name:    "invalid JSON",
			txt:     []byte("invalid"),
			want:    gofall.Legality(""),
			wantErr: true,
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			var l gofall.Legality

			err := l.UnmarshalJSON(v.txt)
			if (err != nil) != v.wantErr {
				t.Errorf("unexpected error: %v", err)
			}

			if l != v.want {
				t.Errorf("unexpected legality: got %v, want %v", l, v.want)
			}
		})
	}
}
