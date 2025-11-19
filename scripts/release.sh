#!/bin/bash

# Script to create and push a new version tag for deployment
# Usage: ./scripts/release.sh <version>
# Example: ./scripts/release.sh 1.2.3

set -e

# Check if version is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <version>"
    echo "Example: $0 1.2.3"
    exit 1
fi

VERSION=$1

# Validate version format (semantic versioning)
if ! [[ $VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Error: Version must be in semantic versioning format (e.g., 1.2.3)"
    exit 1
fi

TAG="v$VERSION"

# Check if we're on main branch
CURRENT_BRANCH=$(git branch --show-current)
if [ "$CURRENT_BRANCH" != "main" ]; then
    echo "Warning: You are not on the main branch (current: $CURRENT_BRANCH)"
    read -p "Do you want to continue? (y/n) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

# Check if tag already exists
if git rev-parse "$TAG" >/dev/null 2>&1; then
    echo "Error: Tag $TAG already exists"
    exit 1
fi

# Check for uncommitted changes
if ! git diff-index --quiet HEAD --; then
    echo "Error: You have uncommitted changes. Please commit or stash them first."
    exit 1
fi

echo "Creating release tag: $TAG"
echo "Current commit: $(git log -1 --oneline)"
echo ""

# Confirm before proceeding
read -p "Do you want to create and push this tag? (y/n) " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Release cancelled"
    exit 1
fi

# Create annotated tag
git tag -a "$TAG" -m "Release $TAG"

# Push tag to origin
git push origin "$TAG"

echo ""
echo "Tag $TAG created and pushed successfully"
echo "GitHub Actions will now deploy this release to production"
echo ""
echo "Monitor deployment at: https://github.com/$(git config --get remote.origin.url | sed 's/.*github.com[:/]\(.*\)\.git/\1/')/actions"
