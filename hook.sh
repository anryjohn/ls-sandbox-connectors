#!/usr/bin/env bash

set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

echo "(load-file \"connector_postgres.lisp\")" >>./phylum/connector.lisp
echo "(load-file \"connector_email.lisp\")" >>./phylum/connector.lisp
echo "(load-file \"connector_camunda_start.lisp\")" >>./phylum/connector.lisp
echo "(load-file \"connector_camunda_inspect.lisp\")" >>./phylum/connector.lisp
echo "(load-file \"connector_tests.lisp\")" >>./phylum/main.lisp
sed -i -e 's/func processRequest(/func processRequestOld(/g' ./connectorhub/main.go
sed -i -e 's/func (s \*g) Run(/func (s \*g) RunOld(/g' ./connectorhub/main.go
sed -i -e 's!dns:///localhost:7051!dns:///peer0.org1.luther.systems:7051!g' ./connectorhub/main.go
git add connectorhub/.gitignore connectorhub/internal/ connectorhub/main.go fabric/client.sh phylum/claim.lisp phylum/connector.lisp tests/example/claim.martin_collection.yaml
git clean -d -f
make apiclean
make api
#make tidy-ch
if [ "${1-}" != "skip" ]
then
  ./scripts/make-up.sh
fi
