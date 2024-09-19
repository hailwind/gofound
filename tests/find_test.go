package tests

import (
	"log"
	"testing"

	"github.com/sea-team/gofound/searcher/model"
	"github.com/sea-team/gofound/searcher/sorts"
)

func TestFindx(t *testing.T) {

	fastSort := &sorts.FastSort{
		IsDebug: true,
		Order:   "desc",
	}

	fastSort.Add(&[]uint32{5566, 299, 220, 788, 24959, 48})
	fastSort.Process()
	var ris = make([]model.SliceItem, 0)
	fastSort.GetAll(&ris, 0, 5)
	for _, x := range ris {
		log.Printf("ID: %d SCORE: %d", x.Id, x.Score)
	}
	log.Fatalln(".")
}
