package player

import (
	"reflect"
	"testing"

	"github.com/syllabix/psychic-poker-player/poker"
)

func Test_assembleHand(t *testing.T) {
	type args struct {
		codes []string
	}
	tests := []struct {
		name    string
		args    args
		want    poker.Hand
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := assembleHand(tt.args.codes)
			if (err != nil) != tt.wantErr {
				t.Errorf("assembleHand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("assembleHand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRevealBestHand(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Straight Flush",
			args:    args{"TH JH QC QD QS QH KH AH 2S 6S"},
			want:    "Hand: TH JH QC QD QS Deck: QH KH AH 2S 6S Best hand: straight-flush",
			wantErr: false,
		},
		{
			name:    "Four of a Kind",
			args:    args{"2H 2S 3H 3S 3C 2D 3D 6C 9C TH"},
			want:    "Hand: 2H 2S 3H 3S 3C Deck: 2D 3D 6C 9C TH Best hand: four-of-a-kind",
			wantErr: false,
		},
		{
			name:    "Full House",
			args:    args{"2H 2S 3H 3S 3C 2D 9C 3D 6C TH"},
			want:    "Hand: 2H 2S 3H 3S 3C Deck: 2D 9C 3D 6C TH Best hand: full-house",
			wantErr: false,
		},
		{
			name:    "Flush",
			args:    args{"2H AD 5H AC 7H AH 6H 9H 4H 3C"},
			want:    "Hand: 2H AD 5H AC 7H Deck: AH 6H 9H 4H 3C Best hand: flush",
			wantErr: false,
		},
		{
			name:    "Straight",
			args:    args{"AC 2D 9C 3S KD 2S 3D 4S 5S 6C"},
			want:    "Hand: AC 2D 9C 3S KD Deck: 2S 3D 4S 5S 6C Best hand: straight",
			wantErr: false,
		},
		{
			name:    "Three of a Kind",
			args:    args{"KS AH 2H 3C 4H KC 2C TC 2D AS"},
			want:    "Hand: KS AH 2H 3C 4H Deck: KC 2C TC 2D AS Best hand: three-of-a-kind",
			wantErr: false,
		},
		{
			name:    "Two Pairs",
			args:    args{"AH 2C 9S AD 3C QH KS JS JD KD"},
			want:    "Hand: AH 2C 9S AD 3C Deck: QH KS JS JD KD Best hand: two-pairs",
			wantErr: false,
		},
		{
			name:    "Two Pairs",
			args:    args{"6C 9C 8C 2D 7C 2H TC 4C 9S AH"},
			want:    "Hand: 6C 9C 8C 2D 7C Deck: 2H TC 4C 9S AH Best hand: one-pair",
			wantErr: false,
		},
		{
			name:    "High Card",
			args:    args{"3D 5S 2H QD TD 6S KH 9H AD QH"},
			want:    "Hand: 3D 5S 2H QD TD Deck: 6S KH 9H AD QH Best hand: highest-card",
			wantErr: false,
		},
		{
			name:    "Invalid Input Length",
			args:    args{"2H AD 5H 6H 9H 4H 3C"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid Card Formats",
			args:    args{"2H A 5H 6H 9H 4H3 3C"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RevealBestHand(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("RevealBestHand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("\nRevealBestHand() = \n%v\nExpected:\n%v", got, tt.want)
			}
		})
	}
}
