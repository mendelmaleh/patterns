RUN ?= go run cmd/pattern/main.go

examples:
	$(RUN) -x 10 '[-+]?[0-9]{1,16}[.][0-9]{1,6}'
	$(RUN) '[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{8}'
	$(RUN) '.{8,12}'
	$(RUN) '[^aeiouAEIOU0-9]{5}'
	$(RUN) '[a-f-]{5}'

install:
	go install ./cmd/pattern
