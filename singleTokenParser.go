package parser

type SingleToken struct {
	Name string
}

func CreateSingleTokenParser(text string) func(t string) *Result {
	tokenId := text
	return func(t string) *Result {
		if IsNext(t, tokenId) {
			return &Result{
				Token: &SingleToken{
					Name: tokenId,
				},
				Count: len(tokenId),
			}
		}
		return nil
	}
}
