set -x
h=$(cat cli_last_sync)
git --git-dir=../cli/.git diff ${h} HEAD -- spellshape/services/network | patch -p3
git --git-dir=../cli/.git diff ${h} HEAD -- spellshape/cmd/network* | patch -p2
git --git-dir=../cli/.git rev-parse HEAD > cli_last_sync
