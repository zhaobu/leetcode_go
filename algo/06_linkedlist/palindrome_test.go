package _6_linkedlist

import "testing"

// func init() {
// 	zaplog.InitLog(&zaplog.Config{
// 		Logdir:   "./",
// 		LogName:  "main.log",
// 		Loglevel: "debug",
// 		Debug:    true},
// 	)
// }

func TestPalindrome1(t *testing.T) {
	strs := []string{"heooeh", "hello", "heoeh", "a", ""}
	for _, str1 := range strs {
		l := NewLinkedList()
		for _, c := range str1 {
			l.InsertToTail(string(c))
		}
		l.Print()
		t.Log(isPalindrome1(l))
	}
}

func TestPalindrome2(t *testing.T) {
	strs := []string{"heooeh", "hello", "heoeh", "a", ""}
	for _, str1 := range strs {
		l := NewLinkedList()
		for _, c := range str1 {
			l.InsertToTail(string(c))
		}
		l.Print()
		t.Log(isPalindrome2(l))
		l.Print()
	}
}

func TestPalindrome(t *testing.T) {

	strs := []string{"12345"}
	for _, str1 := range strs {
		l := NewLinkedList()
		for _, c := range str1 {
			l.InsertToTail(string(c))
		}
		l.Print()
		t.Log(isPalindrome(l.head.next))
		l.Print()
	}
}
