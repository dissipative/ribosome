package sets

import (
	"context"
	"errors"
	"github.com/dissipative/ribosome/pkg/bioio"
	"github.com/dissipative/ribosome/pkg/sequence"
	"sync"
)

const (
	DNA = iota
	RNA
)

type Set struct {
	records []bioio.Record
	molType int
}

func NewDNASet(sequences []bioio.Record) *Set {
	return &Set{
		records: sequences,
		molType: DNA,
	}
}

func NewRNASet(sequences []bioio.Record) *Set {
	return &Set{
		records: sequences,
		molType: RNA,
	}
}

type ORFs struct {
	sync.Mutex
	mapped map[string][]sequence.ORF // record ID -> []ORF
}

func (s *Set) FindORFs(minCodons int, codonTable *sequence.CodonTable) (*ORFs, error) {
	if s.molType == DNA {
		return nil, errors.New("transcribe records to RNA first")
	}

	var orfs ORFs
	orfs.mapped = make(map[string][]sequence.ORF)

	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	var wg sync.WaitGroup
	wg.Add(len(s.records))

	for _, record := range s.records {
		go func(record bioio.Record) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return // if context is done, return immediately
			default:
				rna, err := sequence.NewRNASequence(record.Sequence)
				if err != nil {
					cancel(err)
					return
				}

				orfs.Lock()
				orfs.mapped[record.ID], err = rna.FindORFs(minCodons, codonTable)
				orfs.Unlock()

				if err != nil {
					cancel(err)
				}
			}
		}(record)
	}

	wg.Wait()

	if ctx.Err() != nil {
		return nil, context.Cause(ctx)
	}

	return &orfs, nil
}
