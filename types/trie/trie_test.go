package trie

import "testing"

func TestZoi(t *testing.T) {
	tr := New()
	tr.Add("zoi")
	tr.Add("java")
	tr.Add("ぞい")

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
