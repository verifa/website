#!/bin/bash
echo "Building static site..."
hugo --minify
echo "Site built successfully in ./public/"
echo "You can deploy the contents of the 'public' directory to any static hosting service."
