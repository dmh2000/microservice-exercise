current_dir = $(shell pwd)
all:
	@echo -- $(notdir $(current_dir))	
	$(MAKE) -s -C pkg
	$(MAKE) -s -C internal
	$(MAKE) -s -C cmd 


run:
	$(MAKE) -s -C cmd run	