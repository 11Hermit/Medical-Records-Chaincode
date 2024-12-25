#!/bin/bash

# Set Fabric version
export FABRIC_VERSION=2.5.4
export FABRIC_CA_VERSION=1.5.7

# Create directory for binaries if it doesn't exist
mkdir -p fabric-bins
cd fabric-bins

# Download platform specific binary
OS_ARCH=$(echo "$(uname -s)-$(uname -m)" | tr '[:upper:]' '[:lower:]')

# Download and extract binaries
echo "Downloading Fabric binaries..."
curl -sSL https://github.com/hyperledger/fabric/releases/download/v${FABRIC_VERSION}/hyperledger-fabric-${OS_ARCH}-${FABRIC_VERSION}.tar.gz | tar xz

# Download and extract CA binaries
echo "Downloading Fabric CA binaries..."
curl -sSL https://github.com/hyperledger/fabric-ca/releases/download/v${FABRIC_CA_VERSION}/hyperledger-fabric-ca-${OS_ARCH}-${FABRIC_CA_VERSION}.tar.gz | tar xz

# Download Docker images
echo "Downloading Docker images..."
curl -sSL https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh | bash -s -- docker samples binary

# Add binaries to PATH
export PATH=$PATH:$(pwd)/bin
export FABRIC_CFG_PATH=$(pwd)/config

# Create symbolic links to the test network
cd ..
ln -s fabric-bins/fabric-samples/test-network test-network

echo "Setup complete! Fabric binaries and Docker images have been installed." 