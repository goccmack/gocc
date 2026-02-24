package mail

import (
	"testing"

	"github.com/goccmack/gocc/example/mail/lexer"
	"github.com/goccmack/gocc/example/mail/token"
)

var testData1 = map[string]bool{
	"mymail@google.com":          true,
	"@google.com":                false,
	`"quoted string"@mymail.com`: true,
	`"unclosed quote@mymail.com`: false,
}

func TestScan1(t *testing.T) {
	for input, ok := range testData1 {
		l := lexer.NewLexer([]byte(input))
		tok := l.Scan()
		switch {
		case tok.Type == token.INVALID:
			if ok {
				t.Errorf("%s", input)
			}
		case tok.Type == token.TokMap.Type("addrspec"):
			if !ok {
				t.Errorf("%s", input)
			}
		default:
			t.Fatalf("This must not happen")
		}
	}
}

func TestNext1(t *testing.T) {
	for input, ok := range testData1 {
		l := lexer.NewLexer([]byte(input))
		_, err := l.Next()
		if ok {
			if err != nil {
				t.Errorf("want parsed %s, but got %v", input, err)
			}
		} else {
			if err == nil {
				t.Errorf("wanted fail %s", input)
			}
		}
	}
}

var checkData2 = []string{
	"addr1@gmail.com",
	"addr2@gmail.com",
	"addr3@gmail.com",
}

var testData2 = `
	addr1@gmail.com
	addr2@gmail.com
	addr3@gmail.com
`

func TestScan2(t *testing.T) {
	l := lexer.NewLexer([]byte(testData2))
	num := 0
	for tok := l.Scan(); tok.Type == token.TokMap.Type("addrspec"); tok = l.Scan() {
		if string(tok.Lit) != checkData2[num] {
			t.Errorf("%s != %s", string(tok.Lit), checkData2[num])
		}
		num++
	}
	if num != len(checkData2) {
		t.Fatalf("%d addresses parsed", num)
	}
}

func TestNext2(t *testing.T) {
	l := lexer.NewLexer([]byte(testData2))
	num := 0
	tokBytes, err := l.Next()
	for err == nil {
		if string(tokBytes) != checkData2[num] {
			t.Errorf("%s != %s", string(tokBytes), checkData2[num])
		}
		num++
		tokBytes, err = l.Next()
	}
	if num != len(checkData2) {
		t.Fatalf("%d addresses parsed", num)
	}
}

func BenchmarkScan(b *testing.B) {
	l := lexer.NewLexer([]byte{})
	inputs := [][]byte{
		[]byte("mymail@google.com"),
		[]byte(`"quoted string"@mymail.com`),
		[]byte("addr1@gmail.com"),
	}
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		for _, input := range inputs {
			l.Init(input)
			tok := l.Scan()
			if tok.Type == token.INVALID {
				b.Error("unexpected invalid token")
			}
		}
	}
}

func BenchmarkNext(b *testing.B) {
	l := lexer.NewLexer([]byte{})
	inputs := [][]byte{
		[]byte("mymail@google.com"),
		[]byte(`"quoted string"@mymail.com`),
		[]byte("addr1@gmail.com"),
	}
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		for _, input := range inputs {
			l.Init(input)
			_, err := l.Next()
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}
