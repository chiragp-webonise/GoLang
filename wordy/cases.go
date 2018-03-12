package wordy

type wordyTest struct {
	description string
	question    string
	ok          bool
	answer      int
}

var tests = []wordyTest{
	{
		"addition",
		"What is 1 plus 1?",
		true,
		2,	
	},
}
