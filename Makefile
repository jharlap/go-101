.PHONY: tools server

slides.md: slides.embedded.md
	embedmd $< > $@

server:
	python -m SimpleHTTPServer &
	fswatch slides.embedded.md | xargs -I {} make

TOOLS=github.com/campoy/embedmd
tools:
	$(foreach tool,$(TOOLS),$(call goget, $(tool)))

define goget
        go get -u $(1)
endef

