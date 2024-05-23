current_dir = $(shell pwd)
all:
	@echo ---- $(notdir $(current_dir))	
	$(MAKE) -s -C controller
	$(MAKE) -s -C handler
	$(MAKE) -s -C repository
