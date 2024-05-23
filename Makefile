
all: 
	@go work sync
	$(MAKE) -s -C metadatamake
	$(MAKE) -s -C rating
	$(MAKE) -s -C movie	

run:
	$(MAKE) -s -C movie run
	$(MAKE) -s -C rating run
	$(MAKE) -s -C metadata run

