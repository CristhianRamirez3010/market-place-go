package domainsEndpoints

import "fmt"

type DomainsEndpoints struct {
	GetDocuments string
}

func BuildDomainsEndpoints() DomainsEndpoints {
	user := "/domains/"
	return DomainsEndpoints{
		GetDocuments: fmt.Sprintf("%s%s", user, "documents"),
	}
}
