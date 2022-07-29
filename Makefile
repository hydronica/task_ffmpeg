PREFIX=/usr/local
DESTDIR=
version=$(shell git describe --tags 2> /dev/null || echo "no-version")
GOFLAGS=-ldflags "-X github.com/hydronica/task_ffmpeg.Version=${version} -X github.com/hydronica/task_ffmpeg.BuildTimeUTC=`date -u '+%Y-%m-%d_%I:%M:%S%p'`"

BLDDIR = deploy/bin
EXT=
ifeq (${GOOS},windows)
    EXT=.exe
endif

APPS = info encoder

all: $(APPS)

$(BLDDIR)/info:     $(wildcard apps/workers/info/*.go)
$(BLDDIR)/encoder:  $(wildcard apps/workers/encoder/*.go)

$(BLDDIR)/%: clean
	@mkdir -p $(dir $@)
	go build ${GOFLAGS} -o $@ ./apps/*/$*

$(APPS): %: $(BLDDIR)/%

clean:
	rm -rf $(BLDDIR)

.PHONY: install clean all
.PHONY: $(APPS)

install: $(APPS)
	install -m 755 -d ${DESTDIR}${BINDIR}
	for APP in $^ ; do install -m 755 ${BLDDIR}/$$APP ${DESTDIR}${BINDIR}/$$APP${EXT} ; done
	rm -rf build