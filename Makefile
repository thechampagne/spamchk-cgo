CGO := go build
STATIC := -buildmode=c-archive
SHARED := -buildmode=c-shared
LIBS := build/spamchk.a build/spamchk.so

.PHONY: all
all: $(LIBS)

build/spamchk.a: spamchk.go
	$(CGO) $(STATIC) -o build/spamchk.a $<

build/spamchk.so: spamchk.go
	$(CGO) $(SHARED) -o build/spamchk.so $<

.PHONY: clean
clean:
	find build -type f \( -name '*.h' -o -name '*.so' -o -name '*.a' \) -delete
