BIN=crtchkr
GO=$(shell command -v go)
GOMD2MAN=$(shell command -v go-md2man)
SED=$(shell command -v sed)

# Manpage generation yoinked fom podman
# Credit: https://github.com/containers/podman
# The 'sort' below is crucial: without it, 'make docs' behaves differently
# on the first run than on subsequent ones, because the generated .md
MANPAGES_SOURCE_DIR = man
MANPAGES_MD ?= $(sort $(wildcard $(MANPAGES_SOURCE_DIR)/*.md))
MANPAGES ?= $(MANPAGES_MD:%.md=%)

all: bin


bin:
	$(GO) build -o ./build/$(BIN) .

run:
	./$(BIN)

clean:
	@rm -rf ./build/$(BIN)

pdf:
	@Rscript -e 'rmarkdown::render("doc/doc.Rmd", output_format = "pdf_document")'

.PHONY: docs
docs: $(MANPAGES) ## Generate documentation

$(MANPAGES): %: %.md docdir # .install.md2man

# This does a bunch of filtering needed for man pages:
#  1. Strip markdown link targets like '[podman(1)](podman.1.md)'
#     to just '[podman(1)]', because man pages have no link mechanism;
#  2. Then remove the brackets: '[podman(1)]' -> 'podman(1)';
#  3. Then do the same for all other markdown links,
#     like '[cgroups(7)](https://.....)'  -> just 'cgroups(7)';
#  4. Remove HTML-ish stuff like '<sup>..</sup>' and '<a>..</a>'
#  5. Replace "\" (backslash) at EOL with two spaces (no idea why)
	@$(SED) -e 's/\((crtchkr[^)]*\.md\(#.*\)\?)\)//g'    \
	       -e 's/\[\(crtchkr[^]]*\)\]/\1/g'              \
	       -e 's/\[\([^]]*\)](http[^)]\+)/\1/g'         \
	       -e 's;<\(/\)\?\(a\|a\s\+[^>]*\|sup\)>;;g'    \
	       -e 's/\\$$/  /g' $<                         |\
	$(GOMD2MAN) -out $(subst man,build/man,$@)

.PHONY: docdir
docdir:
	@mkdir -p build/man