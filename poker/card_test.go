package poker

import (
	"reflect"
	"testing"
)

func TestNewCard(t *testing.T) {
	type args struct {
		cardCode string
	}
	tests := []struct {
		name    string
		args    args
		want    Card
		wantErr bool
	}{
		{
			name: "Queen of Spades",
			args: args{"QS"},
			want: Card{
				rankValue: 12,
				rank:      "Q",
				suit:      "S",
				code:      "QS",
			},
			wantErr: false,
		},
		{
			name: "Ten of Hearts",
			args: args{"TH"},
			want: Card{
				rankValue: 10,
				rank:      "T",
				suit:      "H",
				code:      "TH",
			},
			wantErr: false,
		},
		{
			name: "Jack of Diamonds",
			args: args{"JD"},
			want: Card{
				rankValue: 11,
				rank:      "J",
				suit:      "D",
				code:      "JD",
			},
			wantErr: false,
		},
		{
			name: "2 of Clubs",
			args: args{"2C"},
			want: Card{
				rankValue: 2,
				rank:      "2",
				suit:      "C",
				code:      "2C",
			},
			wantErr: false,
		},
		{
			name:    "Queen of What?",
			args:    args{"Q"},
			want:    InvalidCard,
			wantErr: true,
		},
		{
			name:    "That's not a card...",
			args:    args{"ASD$!@@"},
			want:    InvalidCard,
			wantErr: true,
		},
		{
			name:    "Not a Suite",
			args:    args{"KL"},
			want:    InvalidCard,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCard(tt.args.cardCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCard() = %v, want %v", got, tt.want)
			}
		})
	}
}
