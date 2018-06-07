CWD="$(pwd)"
export SOURCES_FILE="$CWD/src/data/sources.txt"
export TARGETS_FILE="$CWD/src/data/targets.txt"
export TMPL_DIR="$CWD/tmpl/"
export SRC_DIR="$CWD/src"

cd $SRC_DIR
go test -coverprofile="$CWD/cover.out"
cd $CWD

