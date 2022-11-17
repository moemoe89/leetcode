package main

func reverseString(s []byte) {
	// define the min and max index
	// which always from 0 and the max is len-1
	i, j := 0, len(s)-1

	// loop and stop if i more than j
	for i < j {
		// swap the value, e.g abcd i = 0, j = 3
		// first swap is a, d will replace by d, a -> dbca (i = 1, j = 2)
		// second swap is b, c will replace by c, b -> dcba (i = 2, j = 1)
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func main() {
	reverseString([]byte(`hello`))
}
