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
			var legality gofall.Legality

			err := legality.UnmarshalText(v.txt)
			if (err != nil) != v.wantErr {
				t.Errorf("unexpected error: %v", err)
			}

			if legality != v.want {
				t.Errorf("unexpected legality: got %v, want %v", legality, v.want)
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
			var legality gofall.Legality

			err := legality.UnmarshalJSON(v.txt)
			if (err != nil) != v.wantErr {
				t.Errorf("unexpected error: %v", err)
			}

			if legality != v.want {
				t.Errorf("unexpected legality: got %v, want %v", legality, v.want)
			}
		})
	}
}

func Test_Legality_MarshalText(t *testing.T) {
	testCases := []struct {
		name     string
		legality gofall.Legality
		want     []byte
		wantErr  bool
	}{
		{
			name:     "legal",
			legality: gofall.LegalityLegal,
			want:     []byte("legal"),
			wantErr:  false,
		},
		{
			name:     "not legal",
			legality: gofall.LegalityNotLegal,
			want:     []byte("not_legal"),
			wantErr:  false,
		},
		{
			name:     "restricted",
			legality: gofall.LegalityRestricted,
			want:     []byte("restricted"),
			wantErr:  false,
		},
		{
			name:     "banned",
			legality: gofall.LegalityBanned,
			want:     []byte("banned"),
			wantErr:  false,
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			got, err := v.legality.MarshalText()
			if (err != nil) != v.wantErr {
				t.Errorf("unexpected error: %v", err)
			}

			if string(got) != string(v.want) {
				t.Errorf("unexpected legality: got %v, want %v", got, v.want)
			}
		})
	}
}
