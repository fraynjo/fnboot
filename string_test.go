package fnboot

import (
	"testing"
)

func TestCamelCase(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "hello_world_foo"
		result := CamelCase(s, "_")
		t.Log(result)
		if result != "helloWorldFoo" {
			t.Fatal("fail")
		}
	})
}

func TestCapitalize(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "helloFoo"
		result := Capitalize(s)
		t.Log(result)
		if result != "Hellofoo" {
			t.Fatal("fail")
		}
	})
}

func TestEndsWith(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "helloFoo"
		target := "Foo"
		result := EndsWith(s, target)
		t.Log(result)
		if result != true {
			t.Fatal("fail")
		}
	})
	t.Run("empty string", func(t *testing.T) {
		s := "helloFoo"
		target := ""
		result := EndsWith(s, target)
		t.Log(result)
		if result != true {
			t.Fatal("fail")
		}
	})
}

func TestStartWith(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "helloFoo"
		target := "hello"
		result := StartWith(s, target)
		t.Log(result)
		if result != true {
			t.Fatal("fail")
		}
	})
	t.Run("empty string", func(t *testing.T) {
		s := "helloFoo"
		target := ""
		result := StartWith(s, target)
		t.Log(result)
		if result != true {
			t.Fatal("fail")
		}
	})
}

func TestLowerFirst(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "FOO"
		result := LowerFirst(s)
		t.Log(result)
		if result != "fOO" {
			t.Fatal("fail")
		}
	})
}

func TestUpperFirst(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "foo"
		result := UpperFirst(s)
		t.Log(result)
		if result != "Foo" {
			t.Fatal("fail")
		}
	})
}

func TestPadEnd(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "FOO"
		result := PadEnd(s, 4)
		t.Log(result)
		if result != "FOO    " {
			t.Fatal("fail")
		}
	})
}

func TestPadStart(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "FOO"
		result := PadStart(s, 4)
		t.Log(result)
		if result != "    FOO" {
			t.Fatal("fail")
		}
	})
}

func TestParseInt(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "99"
		result := ParseInt(s)
		t.Log(result)
		if result != int64(99) {
			t.Fatal("fail")
		}
	})
}

func TestRepeat(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "a"
		result := Repeat(s, 4)
		t.Log(result)
		if result != "aaaa" {
			t.Fatal("fail")
		}
	})
}

func TestReplace(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "Hello world"
		result := Replace(s, "world", "foo")
		t.Log(result)
		if result != "Hello foo" {
			t.Fatal("fail")
		}
	})
}

func TestReplaceAll(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "Hello world world"
		result := ReplaceAll(s, "world", "foo")
		t.Log(result)
		if result != "Hello foo foo" {
			t.Fatal("fail")
		}
	})
}

func TestSnakeCase(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "HelloWorldFoo"
		result := SnakeCase(s)
		t.Log(result)
		if result != "hello_world_foo" {
			t.Fatal("fail")
		}
	})
}

func TestSplit(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "Hello,World"
		result := Split(s, ",")
		t.Log(result)
		if result[0] != "Hello" || result[1] != "World" {
			t.Fatal("fail")
		}
	})
}

func TestToLower(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "Hello"
		result := ToLower(s)
		t.Log(result)
		if result != "hello" {
			t.Fatal("fail")
		}
	})
}

func TestToUpper(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "Hello"
		result := ToUpper(s)
		t.Log(result)
		if result != "HELLO" {
			t.Fatal("fail")
		}
	})
}

func TestTrim(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "  He llo  "
		result := Trim(s, " ")
		t.Log(result)
		if result != "He llo" {
			t.Fatal("fail")
		}
	})
}

func TestTrimEnd(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "  He llo  "
		result := TrimEnd(s, " ")
		t.Log(result)
		if result != "  He llo" {
			t.Fatal("fail")
		}
	})
}

func TestTrimStart(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "  He llo  "
		result := TrimStart(s, " ")
		t.Log(result)
		if result != "He llo  " {
			t.Fatal("fail")
		}
	})
}

func TestTruncate(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := "hello,world"
		result := Truncate(s, 8, "")
		t.Log(result)
		if result != "hello,wo..." {
			t.Fatal("fail")
		}
	})
}

func TestRandomNumber(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		result := RandomNumber(6)
		t.Log(result)
		if result == "" {
			t.Fatal("fail")
		}
	})
}

func TestRandString(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		result := RandString(6)
		t.Log(result)
		if result == "" {
			t.Fatal("fail")
		}
	})
}
