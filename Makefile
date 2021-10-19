make-forbids-spawn: main.go
	go build

test: make-forbids-spawn
	./make-forbids-spawn

test-without-spawn: make-forbids-spawn
	./make-forbids-spawn -spawn=false
