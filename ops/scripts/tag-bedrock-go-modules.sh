#!/usr/bin/env bash

BEDROCK_TAGS_REMOTE="$1"
VERSION="$2"

if [ -z "$VERSION" ]; then
	echo "You must specify a version."
	exit 0
fi

FIRST_CHAR=$(printf '%s' "$VERSION" | cut -c1)
if [ "$FIRST_CHAR" != "v" ]; then
	echo "Tag must start with v."
	exit 0
fi

git tag "mt-bindings/$VERSION"
git tag "mt-service/$VERSION"
git push $BEDROCK_TAGS_REMOTE "mt-bindings/$VERSION"
git push $BEDROCK_TAGS_REMOTE "mt-service/$VERSION"

cd mt-chain-ops
go get github.com/mantlenetworkio/mantle/mt-bindings@$VERSION
go get github.com/mantlenetworkio/mantle/mt-service@$VERSION
go mod tidy

git add .
git commit -am 'chore: Upgrade mt-chain-ops dependencies'

git tag "mt-chain-ops/$VERSION"
git push $BEDROCK_TAGS_REMOTE "mt-chain-ops/$VERSION"

cd ../mt-node
go get github.com/mantlenetworkio/mantle/mt-bindings@$VERSION
go get github.com/mantlenetworkio/mantle/mt-service@$VERSION
go get github.com/mantlenetworkio/mantle/mt-chain-ops@$VERSION
go mod tidy

echo Please update the version to ${VERSION} in mt-node/version/version.go
read -p "Press [Enter] key to continue"

git add .
git commit -am 'chore: Upgrade mt-node dependencies'
git push $BEDROCK_TAGS_REMOTE
git tag "mt-node/$VERSION"
git push $BEDROCK_TAGS_REMOTE "mt-node/$VERSION"

cd ../mt-proposer
go get github.com/mantlenetworkio/mantle/mt-bindings@$VERSION
go get github.com/mantlenetworkio/mantle/mt-service@$VERSION
go get github.com/mantlenetworkio/mantle/mt-node@$VERSION
go mod tidy

echo Please update the version to ${VERSION} in mt-proposer/cmd/main.go
read -p "Press [Enter] key to continue"

git add .
git commit -am 'chore: Upgrade mt-proposer dependencies'
git push $BEDROCK_TAGS_REMOTE
git tag "mt-proposer/$VERSION"
git push $BEDROCK_TAGS_REMOTE "mt-proposer/$VERSION"

cd ../mt-batcher
go get github.com/mantlenetworkio/mantle/mt-bindings@$VERSION
go get github.com/mantlenetworkio/mantle/mt-service@$VERSION
go get github.com/mantlenetworkio/mantle/mt-node@$VERSION
go get github.com/mantlenetworkio/mantle/mt-proposer@$VERSION
go mod tidy

echo Please update the version to ${VERSION} in mt-batcher/cmd/main.go
read -p "Press [Enter] key to continue"

git add .
git commit -am 'chore: Upgrade mt-batcher dependencies'
git push $BEDROCK_TAGS_REMOTE
git tag "mt-batcher/$VERSION"
git push $BEDROCK_TAGS_REMOTE "mt-batcher/$VERSION"

cd ../mt-e2e
go get github.com/mantlenetworkio/mantle/mt-bindings@$VERSION
go get github.com/mantlenetworkio/mantle/mt-service@$VERSION
go get github.com/mantlenetworkio/mantle/mt-node@$VERSION
go get github.com/mantlenetworkio/mantle/mt-proposer@$VERSION
go get github.com/mantlenetworkio/mantle/mt-batcher@$VERSION
go mod tidy

git add .
git commit -am 'chore: Upgrade mt-e2e dependencies'
git push $BEDROCK_TAGS_REMOTE
git tag "mt-e2e/$VERSION"
git push $BEDROCK_TAGS_REMOTE "mt-e2e/$VERSION"
