package app

//Change changes the character output of #(d35) to the first character of arg(3)
func Change(cnum, char string) string {

	charbyte := []byte(char)
	cnumbyte := []byte(cnum)
	for i := 0; i < len(cnumbyte); i++ {

		if cnumbyte[i] == 35 {
			cnumbyte[i] = charbyte[0]
		}

	}
	return string(cnumbyte)
}
