package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// MedicalRecord represents a medical record on the blockchain
type MedicalRecord struct {
    ID          string `json:"id"`
    PatientID   string `json:"patientId"`
    DoctorID    string `json:"doctorId"` 
    Diagnosis   string `json:"diagnosis"`
    Treatment   string `json:"treatment"`
    Timestamp   string `json:"timestamp"`
}

// SmartContract provides functions for managing medical records
type SmartContract struct {
    contractapi.Contract
}

// CreateMedicalRecord adds a new medical record to the blockchain
func (s *SmartContract) CreateMedicalRecord(ctx contractapi.TransactionContextInterface, id string, patientId string, doctorId string, diagnosis string, treatment string, timestamp string) error {
    record := MedicalRecord{
        ID: id,
        PatientID: patientId,
        DoctorID: doctorId,
        Diagnosis: diagnosis,
        Treatment: treatment,
        Timestamp: timestamp,
    }

    recordJSON, err := json.Marshal(record)
    if err != nil {
        return err
    }

    return ctx.GetStub().PutState(id, recordJSON)
}

// GetMedicalRecord retrieves a medical record by ID
func (s *SmartContract) GetMedicalRecord(ctx contractapi.TransactionContextInterface, id string) (*MedicalRecord, error) {
    recordJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return nil, fmt.Errorf("failed to read record: %v", err)
    }
    if recordJSON == nil {
        return nil, fmt.Errorf("record does not exist")
    }

    var record MedicalRecord
    err = json.Unmarshal(recordJSON, &record)
    if err != nil {
        return nil, err
    }

    return &record, nil
}

func main() {
    chaincode, err := contractapi.NewChaincode(&SmartContract{})
    if err != nil {
        fmt.Printf("Error creating medical records chaincode: %v\n", err)
        return
    }

    if err := chaincode.Start(); err != nil {
        fmt.Printf("Error starting medical records chaincode: %v\n", err)
    }
} 