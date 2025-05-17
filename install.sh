#!/usr/bin/env bash

if ! command -v gh > /dev/null 2>&1
then
  echo "unable to find gh binary."
  exit 1
fi

if ! command -v sha256sum > /dev/null 2>&1
then
  echo "unable to find sha256sum binary."
  exit 1
fi

GITHUB_TOKEN="$(gh auth token 2>/dev/null)"
if [ -z "$GITHUB_TOKEN" ]
then
  gh auth login -p https -w
fi
GITHUB_TOKEN="$(gh auth token)"

TEMPDIR="$(mktemp -d)"
trap 'rm -rf -- "$TEMPDIR"' EXIT

REPO="github.com/dmcclory/priority-ranker"
LATEST_VERSION="$(gh release -R "$REPO" ls --exclude-pre-releases | grep Latest | cut -f1)"

echo "downloading the latest release ($LATEST_VERSION)..."
gh release -R "$REPO" download "$LATEST_VERSION" -D "$TEMPDIR"

echo "validate checksums.."
(cd "$TEMPDIR" && sha256sum -c "priority-ranker_${LATEST_VERSION:1}_checksums.txt")

# create the destination dir
DESTINATION_DIR="$HOME/.ranker/bin"
mkdir -p "$DESTINATION_DIR"

# find the right binary to use & copy to the destination dir
KERNEL="$(uname -s)"
ARCH="$(uname -m)"
(cd "$TEMPDIR" && tar xzvf "priority-ranker_${KERNEL}_${ARCH}.tar.gz" && cp priority-ranker "$DESTINATION_DIR/ranker")

if [ -f "$HOME/.zshrc" ] && ! grep -q '$HOME/.ranker/bin' "$HOME/.zshrc"
then
  echo "Adding ranker to the ZSH path"
  echo 'export PATH="$HOME/.ranker/bin:$PATH"' >> "$HOME/.zshrc"
fi

if [ -f "$HOME/.bashrc" ] && ! grep -q '$HOME/.ranker/bin' "$HOME/.bashrc"
then
  echo "Adding ranker to the bash path"
  echo 'export PATH="$HOME/.ranker/bin:$PATH"' >> "$HOME/.bashrc"
fi

if [ -f "$HOME/.config/fish/config.fish" ] && ! grep -q '$HOME/.ranker/bin' "$HOME/.config/fish/config.fish"
then
  echo "Adding ranker to the fish path"
  echo 'export PATH="$HOME/.ranker/bin:$PATH"' >> "$HOME/.config/fish/config.fish"
fi

echo "ranker $LATEST_VERSION has been successfully installed!"
