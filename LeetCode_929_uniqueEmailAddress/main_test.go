package t_test

import (
	"strings"
	"testing"
)

func numUniqueEmails(emails []string) int {
	emailsMap := make(map[string]bool)
	for _, email := range emails {
		splitEmail := strings.Split(email, "@")
		localName := splitEmail[0]
		domainName := splitEmail[1]
		localName = strings.Split(localName, "+")[0]
		localName = strings.Replace(localName, ".", "", -1)
		uniqueEmail := localName + "@" + domainName
		if emailsMap[uniqueEmail] == false {
			emailsMap[uniqueEmail] = true
		}
	}

	return len(emailsMap)

}

func Test_Case1(t *testing.T) {
	var input []string

	res := numUniqueEmails(input)

	if res != 0 {
		t.Error("Not match")
	}
}

func Test_Case2(t *testing.T) {
	input := []string{"test.email@leetcode.com"}

	res := numUniqueEmails(input)

	if res != 1 {
		t.Error("Not match")
	}
}

func Test_Case3(t *testing.T) {
	input := []string{"test.email+1@leetcode.com"}

	res := numUniqueEmails(input)

	if res != 1 {
		t.Error("Not match")
	}
}

func Test_Case4(t *testing.T) {
	input := []string{"test.email+1@leetcode.com", "test.e.mail+bob.cathy@leetcode.com"}

	res := numUniqueEmails(input)

	if res != 1 {
		t.Error("Not match")
	}
}

func Test_Case5(t *testing.T) {
	input := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}

	res := numUniqueEmails(input)

	if res != 2 {
		t.Error("Not match")
	}
}
