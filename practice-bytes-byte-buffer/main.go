s := []byte("hello world! Bitch")

_, err = f.Write(s)
if err != nil {
	log.Fatalf("error %s", err)
}

b := bytes.NewBufferString("Hello! ")
fmt.Println(b.String())
b.WriteString("hello gophers!")
fmt.Println(b.String())
b.Reset()
b.WriteString("this is written after reseting")
fmt.Println(b.String())

b.Write([]byte("Happy Happy"))
fmt.Println(b.String())