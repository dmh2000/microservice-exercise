current_dir = $(shell pwd)
all:
	@echo -------- $(notdir $(current_dir))
	staticcheck .
	go build .