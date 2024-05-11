package main
import(
"strings"
"testing"
)
//ReverseWords takes astring asinput and and returns the string with each word reversedWord
func Reversewords(input string)string {
words:=
strings.Fields(input)
reversed:=make([]string,len(words))
for i,word :=
range words {
  reversedWord:=""
for _, char :=range word{

reversedWord = string(char) +
reversedWord
}
reversed[i]= reversedWord
}
return strings.Join(reversed,"")
}
//TestsReversedWords function
func
TestReversedWords(t *testing.T) {
testCases := []struct {
  input         string
  expectedOutput  string
}{
{"Hello, World!","olleH, !dlrow"},
  {"Go Programming","oG gnimmargorp"},
{"I love GO!","I evol !oG"},
}
for _, tc := range testCases {
output := ReverseWords
(tc.input)
if output != tc.expectedOutput {
t.Errorf("ReverseWords(%s) =%s,expected %s",tc.input, output,tc.expectedOutput)
}
}
}
