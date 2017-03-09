# Absolute path to this script, e.g. /home/user/bin/foo.sh
SCRIPT=$(readlink -f "$0")
# Absolute path this script is in, thus /home/user/bin
ONDIRHOME=$(dirname "$SCRIPT")/../bin

cd() {
  builtin cd "$@" && eval "`$ONDIRHOME/ondir \"$OLDPWD\" \"$PWD\"`"
}

pushd() {
	builtin pushd "$@" && eval "`$ONDIRHOME/ondir \"$OLDPWD\" \"$PWD\"`"
}

popd() {
	builtin popd "$@" && eval "`$ONDIRHOME/ondir \"$OLDPWD\" \"$PWD\"`"
}              
