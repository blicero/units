#!/bin/sh
# Time-stamp: <2022-12-22 21:03:23 krylon>

cd $GOPATH/src/github.com/blicero/units/

rm -vf bak.units units dbg.build.log \
    && du -sh . \
    && git fsck --full \
    && git reflog expire --expire=now \
    && git gc --aggressive --prune=now \
    && du -sh .


