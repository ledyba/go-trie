package trie

import "testing"

func TestZoi(t *testing.T) {
	tr := New()
	tr.Add("zoi")
	tr.Add("java")
	tr.Add("ぞい")
	tr.Pack()

	if !tr.Match("zoi") {
		t.Error("ぞいしてよ")
	}

	if !tr.Match("ぞい") {
		t.Error("ぞいしてよ")
	}

	if tr.Match("z") {
		t.Error("まだまっちしちゃだめ")
	}

	if tr.Match("ganbaruzoi") {
		t.Error("最初だけだよ")
	}

	if tr.Match("nenetch") {
		t.Error("ねねっちそこじゃない")
	}

	if !tr.Contains("ganbaruzoi!") {
		t.Error("ぞいしてよ")
	}

	if !tr.Contains("今日も一日がんばるぞい！") {
		t.Error("ぞいしてよ")
	}

	if !tr.Contains("anatatojava") {
		t.Error("ド")
	}

	if !tr.Contains("anatatojava, imasugu download") {
		t.Error("ド")
	}
}

func TestContains(t *testing.T) {
	tr := New()
	tr.Add("うらにわ")
	tr.Add("おおにわとり")
	tr.Add("こけこっこ")
	tr.Add("ok")
	tr.Pack()

	if tr.Contains("にわにはにわにわとりがいる") == true {
		t.Error("わとは")
	}
	if tr.Contains("にわにはにわおおにわとりがいる") == false {
		t.Error("いるよ")
	}
	if !(tr.Contains("コケコッコー") == false) {
		t.Error("カタカナだよ")
	}
	if !(tr.Contains("POKEMON") == false) {
		t.Error("大文字小文字区別して。")
	}
}

func TestReadme(t *testing.T) {
	tr := New() // Animes.
	tr.Add("NewGame!")
	tr.Add("School Live!")
	tr.Add("Urara Meiro Chou")
	tr.Add("Yuki Yuna Is a Hero")
	tr.Add("Non Non Biyori.")
	tr.Add("Anne Happy")
	tr.Add("Kiniro Mosaic")
	tr.Pack()

	// Match method
	if tr.Match("NewGame!") == false {
		t.Error("NewGame! is a first season of the series.")
	}
	if tr.Match("NewGame!!") == false {
		t.Error("NewGame!! is a second season of the series.")
	}
	if tr.Match("NewGame") == true {
		t.Error("Not NewGame. NewGame\"!\"")
	}

	// Contains method
	if tr.Contains("I would like to eat udon with Fuu Inubozaki, a hero in \"Yuki Yuna Is a Hero\".") == false {
		t.Error("What????? Why????")
	}
	if tr.Contains("Alas, Ikaruga is going...") == true {
		t.Error("Ikaruga is a game. Not an animation.")
	}
}