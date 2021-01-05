package easy929

func numUniqueEmails(emails []string) int {
	set := make(map[string]struct{})
	for _, email := range emails {
		byteEmail := []byte(email)
		var domain []byte
		for idx := range byteEmail {
			if byteEmail[len(byteEmail)-1-idx] == "@"[0] {
				break
			}
			domain = append(domain, byteEmail[len(byteEmail)-1-idx])
		}

		var local []byte
		for idx := 0; idx < len(byteEmail)-len(domain)-1; idx++ {
			if byteEmail[idx] == "+"[0] {
				break
			}
			if byteEmail[idx] == "."[0] {
				continue
			}
			local = append(local, byteEmail[idx])
		}
		set[string(local)+string(domain)] = struct{}{}
	}
	return len(set)
}
