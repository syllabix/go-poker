package poker

import "testing"

type args struct {
	hand Hand
}

func TestGetRank(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRank(tt.args.hand); got != tt.want {
				t.Errorf("GetRank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGetRank(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GetRank(test.args.hand)
			}
		})
	}
}

var tests = []struct {
	name string
	args args
	want RankCategory
}{
	{
		name: "Straight Flush",
		args: args{hand1},
		want: StraightFlush,
	},
	{
		name: "Four of a Kind",
		args: args{hand4},
		want: FourOfAKind,
	},
	{
		name: "Full House",
		args: args{hand3},
		want: FullHouse,
	},
	{
		name: "Flush",
		args: args{hand2},
		want: Flush,
	},
	{
		name: "Straight",
		args: args{hand5},
		want: Straight,
	},
	{
		name: "Three of a Kind",
		args: args{hand6},
		want: ThreeOfAKind,
	},
	{
		name: "Two Pair",
		args: args{hand7},
		want: TwoPair,
	},
	{
		name: "One Pair",
		args: args{hand8},
		want: OnePair,
	},
	{
		name: "High Card",
		args: args{hand9},
		want: HighCard,
	},
}

var (
	hand1 = Hand{
		createCard("5D"),
		createCard("3D"),
		createCard("4D"),
		createCard("6D"),
		createCard("7D"),
	}

	hand2 = Hand{
		createCard("AS"),
		createCard("TS"),
		createCard("5S"),
		createCard("KS"),
		createCard("2S"),
	}

	hand3 = Hand{
		createCard("7D"),
		createCard("QH"),
		createCard("QS"),
		createCard("7C"),
		createCard("QC"),
	}

	hand4 = Hand{
		createCard("5D"),
		createCard("2C"),
		createCard("5S"),
		createCard("5C"),
		createCard("5H"),
	}

	hand5 = Hand{
		createCard("QS"),
		createCard("JD"),
		createCard("TH"),
		createCard("9C"),
		createCard("KH"),
	}

	hand6 = Hand{
		createCard("2S"),
		createCard("JD"),
		createCard("2H"),
		createCard("9C"),
		createCard("2H"),
	}

	hand7 = Hand{
		createCard("2S"),
		createCard("JD"),
		createCard("2H"),
		createCard("9C"),
		createCard("JH"),
	}

	hand8 = Hand{
		createCard("4S"),
		createCard("JD"),
		createCard("2H"),
		createCard("9C"),
		createCard("JH"),
	}

	hand9 = Hand{
		createCard("4S"),
		createCard("JD"),
		createCard("2H"),
		createCard("9C"),
		createCard("AH"),
	}
)

func createCard(cardCode string) Card {
	card, err := NewCard(cardCode)
	if err != nil {
		panic(err)
	}
	return *card
}
