package package_name

func hello() string {
 words := []string{"hello", "func", "in", "package", "hello"}
 wl := len(words)
 
 sentence := ""
 for key, word := range words {
  sentence += word
  if key < wl-1 {
   sentence += " "
  } else {
   sentence += "."
  }
 }
 return sentence
}

func ExampleHello1() string {
 hl := hello()
 return hl
 // Output: hello func in package hello.
}
