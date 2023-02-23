package bigint

import (
	"math/big"
	"math/rand"
)

type mbi = *big.Int

func bi(n int) mbi {
	return big.NewInt(int64(n))
}

func randBytes(n int) (res []byte) {
	res = make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = byte(rand.Intn(1 << 8))
	}
	return
}

type Int struct {
	i mbi
}

func FromInt(i int) Int {
	return Int{bi(i)}
}

var (
	Zero  = FromInt(0)
	One   = FromInt(1)
	Two   = FromInt(2)
	Three = FromInt(3)
	Four  = FromInt(4)
	Five  = FromInt(5)
	Six   = FromInt(6)
	Seven = FromInt(7)
	Eight = FromInt(8)
	Nine  = FromInt(9)
	Ten   = FromInt(10)
)

func FromBigInt(x *big.Int) Int {
	return Int{bi(0).Set(x)}
}

func FromBytes(b []byte) Int {
	return Int{bi(0).SetBytes(b)}
}

func Copy(x Int) Int {
	return Int{bi(0).Set(x.i)}
}

func Exp(a Int, y Int, n Int) Int {
	return Int{bi(0).Exp(a.i, y.i, n.i)}
}

func RandInt(n Int) Int {
	x := randBytes(len(n.Bytes()))
	return FromBytes(x).Mod(n)
}

func (b Int) Mul(x Int) Int {
	return Int{bi(0).Mul(b.i, x.i)}
}

func (b Int) Mod(x Int) Int {
	return Int{bi(0).Mod(b.i, x.i)}
}

func (b Int) Add(x Int) Int {
	return Int{bi(0).Add(b.i, x.i)}
}

func (b Int) Sub(x Int) Int {
	return Int{bi(0).Sub(b.i, x.i)}
}

func (b Int) Div(x Int) Int {
	return Int{bi(0).Div(b.i, x.i)}
}

func (b Int) Equal(x Int) bool {
	return b.Cmp(x) == 0
}

func (b Int) IsInt64() bool {
	return b.i.IsInt64()
}

func (b Int) Int() int {
	return int(b.i.Int64())
}

func (b Int) Cmp(x Int) int {
	return b.i.Cmp(x.i)
}

func (b Int) Bytes() []byte {
	return b.i.Bytes()
}
