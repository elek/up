VERSION 0.6
FROM golang:1.18
WORKDIR /go/storj-up

lint:
    FROM storjlabs/ci
    COPY . /go/storj-up
    WORKDIR /go/storj-up
    RUN check-copyright
    RUN check-large-files
    RUN check-imports -race ./...
    RUN check-peer-constraints -race
    RUN check-atomic-align ./...
    RUN check-errs ./...
    RUN check-monkit ./...
    RUN staticcheck ./...
    RUN golangci-lint --build-tags mage -j=2 run
    RUN check-mod-tidy -mod .build/go.mod.orig

build:
    RUN ls -lah
    RUN pwd
    COPY . .
    RUN --mount=type=cache,target=/root/.cache/go-build \
        --mount=type=cache,target=/go/pkg/mod \
        go build -o build/ ./...
    SAVE ARTIFACT build/storj-up AS LOCAL build/storj-up

test:
   COPY . .
   RUN go install github.com/mfridman/tparse@36f80740879e24ba6695649290a240c5908ffcbb
   RUN mkdir build
   RUN --mount=type=cache,target=/root/.cache/go-build \
       --mount=type=cache,target=/go/pkg/mod \
       go test -json ./... | tee build/tests.json
   SAVE ARTIFACT build/tests.json AS LOCAL build/tests.json

integration:
   COPY +build/storj-up /usr/local/bin/storj-up
   COPY test/test.sh .
   WITH DOCKER
      RUN ./test.sh
   END

check-format:
   COPY . .
   RUN mkdir build
   RUN bash -c '[[ $(git status --short) == "" ]] || (echo "Before formatting, please commit all your work!!! (Formatter will format only last commit)" && exit -1)'
   RUN git show --name-only --pretty=format: | grep ".go" | xargs -n1 gofmt -s -w
   RUN git diff > build/format.patch
   SAVE ARTIFACT build/format.patch

check-format-all:
   RUN go install mvdan.cc/gofumpt@v0.3.1
   COPY . /go/storj-up
   WORKDIR /go/storj-up
   RUN bash -c 'find -name "*.go" | xargs -n1 gofmt -s -w'
   RUN bash -c 'find -name "*.go" | xargs -n1 gofumpt -s -w'
   RUN mkdir -p build
   RUN git diff > build/format.patch
   SAVE ARTIFACT build/format.patch

format:
   LOCALLY
   COPY +check-format/format.patch build/format.patch
   RUN git apply --allow-empty build/format.patch
   RUN git status

format-all:
   LOCALLY
   COPY +check-format-all/format.patch build/format.patch
   RUN git apply --allow-empty build/format.patch
   RUN git status
