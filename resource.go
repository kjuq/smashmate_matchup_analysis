package main

type Player struct {
	name       string
	fighter    string
	rate       int
	isCanceled bool
}

func errorHandling(err error) {
	if err != nil {
		panic(err)
	}
}

