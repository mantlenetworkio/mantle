#/bin/bash
BASEDIR="../.."
OPS="$BASEDIR/ops"
ARTIFACT="$BASEDIR/packages/contracts/artifacts"
CONTRACTS="$ARTIFACT/contracts"
WORKDIR="./contracts"

if [ ! -d "$ARTIFACT" ]; then
  echo "compile artifacts..."
  cd $OPS
  yarn clean
  yarn install --ignore-engines
  yarn build
  echo "artifacts compiled"
  cd -
fi

mkdir "$WORKDIR"
cp -a "$CONTRACTS/L1" "$WORKDIR"
cp -a "$CONTRACTS/L2" "$WORKDIR"

function listFiles()
{
    #1st param, the dir name
    #2nd param, the aligning space
    for file in `ls $1`;
    do
        if [[ -d "$1/$file" ]]; then
            listFiles "$1/$file" "$1/$file"
        else
            if echo "$file" | grep -q -E '\.dbg.json$' || echo "$file" | grep -q -E '\.go$'
            then
            	continue
            else
#            	echo "$file"
#            	echo "$2/$file"
            	[[ "$file" =~ (.*)\.json  ]]
            	package=${BASH_REMATCH[1]}
            	cat "$2/$file" | jq '.abi' | abigen --abi - --pkg=$package --out="$2/$package.go"
            fi
        fi
    done
}
listFiles "$WORKDIR/L1" "/"
listFiles "$WORKDIR/L2" "/"
