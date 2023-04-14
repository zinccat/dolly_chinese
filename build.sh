#!/bin/bash
go build -tags=jsoniter -v -o "./out/dolly" -ldflags "-s -w" ./gosrc/main/*.go