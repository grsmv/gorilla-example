dir = spec/controllers
pkg = github.com/grsmv/gorilla-example/app/controllers
data = cover.out

run: test fmt clear

test:
	cd $(dir) && go test -v -coverpkg $(pkg) -cover -coverprofile=$(data)

fmt:
	go tool cover -html=$(dir)/$(data)

clear:
	rm $(dir)/$(data)
