package ast

import(
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	testCharRange(t, 'a', 'z')
	testCharRange(t, '0', '9')
}

func testCharRange(t *testing.T, from, to rune) {
	for ch := from; ch <= to; ch++ {
		src := []byte(fmt.Sprintf("'%c'", ch))
		if ch1 := charVal(src); ch1 != ch {
			t.Error(fmt.Sprintf("charVal(%s) returned %d, expected %d\n", src, ch1, ch))
		}
	}
}

type test2T struct {
	lit string
	ch  rune
}

var test2D = []test2T{
	test2T{`'\a'`, '\a'},
	test2T{`'\b'`, '\b'},
	test2T{`'\f'`, '\f'},
	test2T{`'\n'`, '\n'},
	test2T{`'\r'`, '\r'},
	test2T{`'\t'`, '\t'},
	test2T{`'\v'`, '\v'},
	test2T{`'\\'`, '\\'},
	test2T{`'\''`, '\''},
}

func Test2(t *testing.T) {
	for _, src := range test2D {
		if ch1 := charVal([]byte(src.lit)); ch1 != src.ch {
			t.Error(fmt.Sprintf("charVal(%s) returned %d, expected %d\n", src.lit, src.ch, ch1))
		}
	}
}

func Test3(t *testing.T) {
	for i := 1; i < 0100; i++ {
		src := fmt.Sprintf("'\\0%o'", i)
		if ch := charVal([]byte(src)); ch != rune(i) {
			t.Error(fmt.Printf("charVal(%s) returned %o, expected %o\n", src, i, ch))
		}
	}
}

func Test4(t *testing.T) {
	for i := 0; i < 0x100; i++ {
		src := fmt.Sprintf("'\\x%x'", i)
		if ch := charVal([]byte(src)); ch != rune(i) {
			t.Error(fmt.Printf("charVal(%s) returned %x, expected %x\n", src, i, ch))
		}
	}
}

func Test5(t *testing.T) {
	for i := 1; i <= 0x2000; i++ {
		testUnicode(t, rune(i), 'u')
	}
}

func Test6(t *testing.T) {
	for i := 1; i <= 0xd7ff; i++ {
		testUnicode(t, rune(i), 'U')
	}
}

func testUnicode(t *testing.T, ch rune, u rune) {
	var src string
	switch ch {
	case 'u':
		src = fmt.Sprintf("'\\u%04x'", ch)
	default:
		src = fmt.Sprintf("'\\U%08x'", ch)

	}
	if ch1 := charVal([]byte(src)); ch != ch {
		t.Error(fmt.Printf("charVal(%s) returned %x, expected %x\n", src, ch1, ch))
	}
}
