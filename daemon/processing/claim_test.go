package processing

import (
	util "github.com/lbryio/lbry.go/lbrycrd"
	"testing"
)

type claimIDMatch struct {
	ClaimID string
	TxHash  string
	N       uint
}

var claimIDTests = []claimIDMatch{

	{"589bc4845caca70977332025990b2a1807732b44",
		"6a9dbe3084b86cec8aa519970d2245dfa15193294cab65819a0d96d455c2a5df",
		1},

	{"60d7ddcc211c381bad63b73415c2065b219258f2",
		"2850f854108d9fd1d9067cc51ef38664f320cda363741a5774f8c6f0c154a702",
		1},

	{"015a97bef520a8b121baec02f9fd36a9f7e8a17e",
		"6ec15e7a80205fe2fb08eb57a5e4866544db56d81ebfc21d16a19c01a3394779",
		0},

	{"ed15aff6c77d0ae46b542aced757b04f7ec4a507",
		"1cd52537daa096d5fa2b0d20cbcf907fb1a1dc22436f48902473d8af1f7ebe07",
		0},

	{"bef5806ebee8816bcc8d10684eeb5d0d7c906c87",
		"bde00b2ba8ca425fff0d814d12a67d0ce99950365eb21081011aaf4a8c5f3e8b",
		0},
	{"8b120c045300923062d82911868febffccb502bf",
		"de0a48529d193ae33402f9620a25d91d7f33e608022d5e32451acd1d30fe7933",
		0},
	{"84cae80fbe8e49eb45a69ea3af884016c58e8ccb",
		"caf6a81ec886ed2a930a16814f0fdd488a753b22a77f5fb67a11fd3b985edb15",
		0},
	{"0fa76228fe19362a1b0af300217990f061460312",
		"7f3be2ae728ce3a3b5deee57011db1d284276594164b712914acc2f41b3d7152",
		1},
}

func TestGetClaimIDFromOutput(t *testing.T) {

	for _, claimMatch := range claimIDTests {
		claimID, err := util.ClaimIDFromOutpoint(claimMatch.TxHash, int(claimMatch.N))
		if err != nil {
			t.Error(err)
		}
		if claimID != claimMatch.ClaimID {
			t.Error("Expected ", claimMatch.ClaimID, " got ", claimID)
		}
	}

}
