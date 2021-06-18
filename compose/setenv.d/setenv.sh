export VERSION="$(grep '^VERSION=' common.mk | awk -F= '{print $2}')"
export VERSION_SUBSTRATE="$(grep '^SUBSTRATE_VERSION=' common.mk | awk -F= '{print $2}')"
PHYLUM_VERSION_FILE="$(grep '^PHYLUM_VERSION_FILE=' common.fabric.mk | awk -F= '{print $2}')"
mkdir -p $(dirname ./fabric/"$PHYLUM_VERSION_FILE")
touch ./fabric/"$PHYLUM_VERSION_FILE"
export PHYLUM_VERSION="$(cat ./fabric/"$PHYLUM_VERSION_FILE")"
export SUBSTRATE_PLUGIN_LINUX=$(make -f ./common.godynamic.mk echo:SUBSTRATE_PLUGIN_LINUX)
