// +build !skip_msgp_testing

package compactcert

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"testing"

	"github.com/algorand/go-algorand/protocol"
	"github.com/algorand/go-algorand/testpartitioning"
	"github.com/algorand/msgp/msgp"
)

func TestMarshalUnmarshalsigFromAddr(t *testing.T) {
	testpartitioning.PartitionTest(t)

	v := sigFromAddr{}
	bts := v.MarshalMsg(nil)
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestRandomizedEncodingsigFromAddr(t *testing.T) {
	testpartitioning.PartitionTest(t)

	protocol.RunEncodingTest(t, &sigFromAddr{})
}

func BenchmarkMarshalMsgsigFromAddr(b *testing.B) {
	v := sigFromAddr{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgsigFromAddr(b *testing.B) {
	v := sigFromAddr{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalsigFromAddr(b *testing.B) {
	v := sigFromAddr{}
	bts := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}
