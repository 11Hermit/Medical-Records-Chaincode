package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing medical records
type SmartContract struct {
    contractapi.Contract
}

// MedicalRecord describes a medical record
type MedicalRecord struct {
    ID          string `json:"id"`
    PatientID   string `json:"patientId"`
    DoctorID    string `json:"doctorId"`
    Diagnosis   string `json:"diagnosis"`
    Treatment   string `json:"treatment"`
    Timestamp   string `json:"timestamp"`
}

// CreateMedicalRecord adds a new medical record to the world state
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

// GetMedicalRecord returns the medical record stored in the world state with given id
func (s *SmartContract) GetMedicalRecord(ctx contractapi.TransactionContextInterface, id string) (*MedicalRecord, error) {
    recordJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return nil, fmt.Errorf("failed to read from world state: %v", err)
    }
    if recordJSON == nil {
        return nil, fmt.Errorf("the medical record %s does not exist", id)
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