```markdown
# Medical Records Blockchain System

A secure blockchain-based medical records management system built with Hyperledger Fabric and Python.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Directory Structure](#directory-structure)
- [Installation](#installation)
- [Network Setup](#network-setup)
- [Chaincode Deployment](#chaincode-deployment)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [Troubleshooting](#troubleshooting)

## Prerequisites

- Go 1.19 or later
- Python 3.8 or later
- Docker 20.10.x or later
- Docker Compose v2.0.x or later
- Node.js 14.x or later
- Git

Verify installations:
```bash
go version
python --version
docker --version
docker-compose --version
node --version
git --version
```

## Directory Structure

```
medical-records-blockchain/
├── app/
│   ├── frontend/          # Gradio web interface
│   ├── backend/           # Python backend application
│   └── config/            # Configuration files
├── chaincode/
│   └── medical_records/
│       └── go/
│           ├── medical_records.go
│           └── go.mod
├── network/               # Hyperledger Fabric network files
├── scripts/              # Utility scripts
└── README.md
```

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/medical-records-blockchain.git
cd medical-records-blockchain
```

2. Set up Go modules for chaincode:
```bash
cd chaincode/medical_records/go
go mod init medical_records
go get github.com/hyperledger/fabric-contract-api-go@v1.2.1
go get github.com/hyperledger/fabric-chaincode-go@v0.0.0-20230228194215-b84622ba6a7a
go get github.com/hyperledger/fabric-protos-go@v0.3.0
go mod tidy
go mod verify
cd ../../../
```

3. Install Python dependencies:
```bash
# Create and activate virtual environment
python3 -m venv venv
source venv/bin/activate  # On Windows: .\venv\Scripts\activate

# Install required packages
pip install -r requirements.txt
```

4. Install Hyperledger Fabric binaries and docker images:
```bash
curl -sSL https://bit.ly/2ysbOFE | bash -s -- 2.5.0 1.5.5
```

## Network Setup

1. Navigate to the network directory:
```bash
cd network
```

2. Start the Fabric network:
```bash
./network.sh up createChannel -c mychannel
```

3. Verify network status:
```bash
docker ps
```

## Chaincode Deployment

1. Package the chaincode:
```bash
./network.sh deployCC -ccn medical_records -ccp ../chaincode/medical_records/go -ccl go
```

2. Verify chaincode installation:
```bash
peer lifecycle chaincode queryinstalled
```

## Running the Application

1. Start the backend server:
```bash
cd app/backend
python app.py
```

2. Launch the frontend interface:
```bash
cd app/frontend
python ui.py
```

The application should now be accessible at: http://localhost:7860

## API Documentation

### Create Medical Record
```bash
curl -X POST http://localhost:5000/api/records \
  -H "Content-Type: application/json" \
  -d '{
    "patientId": "123",
    "doctorId": "456",
    "diagnosis": "Example diagnosis",
    "treatment": "Example treatment"
  }'
```

### Query Medical Record
```bash
curl -X GET http://localhost:5000/api/records/{recordId}
```

## Troubleshooting

### Common Issues

1. Network startup fails:
```bash
# Clean up existing containers and artifacts
./network.sh down
docker system prune -a
./network.sh up createChannel -c mychannel
```

2. Chaincode deployment issues:
```bash
# Check chaincode logs
docker logs <chaincode-container-id>

# Verify chaincode status
peer lifecycle chaincode querycommitted -C mychannel
```

3. Connection issues:
- Verify that all required ports are available (7051, 7054, 9051, etc.)
- Check connection profiles in app/config/
- Ensure Docker daemon is running

### Logging

- Network logs: `network/logs/`
- Application logs: `app/logs/`
- Chaincode logs: Available through Docker

## Security Considerations

1. TLS Configuration:
- Ensure proper CA certificates are configured
- Verify mutual TLS settings in connection profiles

2. Access Control:
- Review organization MSP configurations
- Verify user roles and permissions

## Additional Resources

- [Hyperledger Fabric Documentation](https://hyperledger-fabric.readthedocs.io/)
- [Fabric SDK for Python](https://github.com/hyperledger/fabric-sdk-py)
- [Gradio Documentation](https://gradio.app/docs/)

## License

This project is licensed under the MIT License - see the LICENSE file for details.
```

This README provides:
1. Clear prerequisites and installation steps
2. Detailed network setup and chaincode deployment instructions
3. API documentation with examples
4. Troubleshooting guidance
5. Security considerations
6. Additional resources

You may want to customize:
- Repository URLs
- Port numbers
- File paths
- API endpoints
- Security configurations

based on your specific implementation.
