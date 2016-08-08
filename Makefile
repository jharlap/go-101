.PHONY: tools server

slides.md: slides.embedded.md
	embedmd $< > $@

server:
	python -m SimpleHTTPServer &
	fswatch slides.embedded.md | xargs -I {} make

remark-latest.min.js:
	curl -O https://gnab.github.io/remark/downloads/remark-latest.min.js

TOOLS=github.com/campoy/embedmd
tools:
	$(foreach tool,$(TOOLS),$(call goget, $(tool)))

define goget
        go get -u $(1)
endef

