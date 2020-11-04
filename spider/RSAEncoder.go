package main

import (
	"fmt"
	"math/big"
)

//rsa n与e
type RSAPublic struct {
	n *big.Int
	e *big.Int
}

//RSA加密
func (key *RSAPublic) RSAEncrypt(pwd string, nStr string, eStr string) string {
	key.n = new(big.Int)
	key.e = new(big.Int)

	key.n, _ = key.n.SetString(nStr, 16)
	key.e, _ = key.e.SetString(eStr, 16)

	r := key.RSADoPublic(key.pkcs1pad2(pwd,(key.n.BitLen() + 7) >> 3))
	sp := r.Text(16)
	if (len(sp) & 1) != 0 {
		sp = "0" + sp
	}
	return sp
}

//快速幂取模
func powmod(a *big.Int, b *big.Int, k *big.Int) *big.Int {
	var tmp *big.Int = a
	var check big.Int
	var ret *big.Int = big.NewInt(1)
	for ;b.Cmp(big.NewInt(0)) != 0; {
		if check.And(b, big.NewInt(1)).Cmp(big.NewInt(0)) != 0 {
			ret.Mul(ret, tmp).Mod(ret, k)
		}
		b.Div(b, big.NewInt(2))
		tmp.Mul(tmp, tmp).Mod(tmp, k)
		//fmt.Printf("%v\n", b)
	}
	return ret
}

//生成密文
func (key *RSAPublic) RSADoPublic(x *big.Int) *big.Int {
	return powmod(x, key.e, key.n)
}

func (key *RSAPublic) pkcs1pad2(s string, n int) *big.Int {
	if n < len(s) + 11 {
		fmt.Println("error: Message too long for RSAEncoder")
		return big.NewInt(0)
	}
	ba := make([]byte, n)
	i := len(s) - 1
	for ;i >= 0 && n > 0; {
		c := int(s[i])
		i--
		if c < 128 {
			n--
			ba[n] = byte(c)
		} else if c > 127 && c < 2048 {
			n--
			ba[n] = byte((c & 63) | 128)
			n--
			ba[n] = byte((c >> 6) | 192)
		} else {
			n--
			ba[n] = byte((c & 63) | 128)
			n--
			ba[n] = byte((c >> 6) & 63 | 128)
			n--
			ba[n] = byte((c >> 12) | 224)
		}
	}
	n--
	ba[n] = 0
	temp := make([]byte, 1)
	for ;n > 2; {
		temp[0] = byte(0)
		for ;temp[0] == 0; {
			temp[0] = byte(1)
		}
		n--
		ba[n] = temp[0]
	}
	n--
	ba[n] = 2
	n--
	ba[n] = 0
	var ans *big.Int
	ans = new(big.Int)
	ans = ans.SetBytes(ba)
	return ans
}