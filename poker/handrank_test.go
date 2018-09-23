package poker

import "testing"

func TestRankCategory_String(t *testing.T) {
	tests := []struct {
		name string
		p    RankCategory
		want string
	}{
		{
			name: "Straight Flush Stringer",
			p:    StraightFlush,
			want: "straight-flush",
		},
		{
			name: "Four Of A Kind Stringer",
			p:    FourOfAKind,
			want: "four-of-a-kind",
		},
		{
			name: "Flush Stringer",
			p:    Flush,
			want: "flush",
		},
		{
			name: "Straight Stringer",
			p:    Straight,
			want: "straight",
		},
		{
			name: "ThreeOfAKind Stringer",
			p:    ThreeOfAKind,
			want: "three-of-a-kind",
		},
		{
			name: "TwoPair Stringer",
			p:    TwoPair,
			want: "two-pairs",
		},
		{
			name: "OnePair Stringer",
			p:    OnePair,
			want: "one-pair",
		},
		{
			name: "HighCard Stringer",
			p:    HighCard,
			want: "highest-card",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("RankCategory.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
