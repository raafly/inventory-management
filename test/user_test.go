package main_test

import "testing"

type ItemDumpy struct {
	Name	int
}

func (r ItemDumpy) CreateWithoutPointr(updateName int) {
	r.Name = updateName
}

func (r *ItemDumpy) CreateWithPointer(updateName int) {
	r.Name = updateName
}

func BenchmarkPoin(b *testing.B) {
	b.Run("tanpa_pointer", func(b *testing.B) {
		person := ItemDumpy{}
		for i := 0; i < b.N; i++ {
			person.CreateWithoutPointr(i)
		}
	})

	b.Run("dengan_pointer", func(b *testing.B) {
		person := ItemDumpy{}
		for i := 0; i < b.N; i++ {
			person.CreateWithPointer(i)
		}
	})
}