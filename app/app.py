import os
from hfc.fabric import Client
import gradio as gr

# Fabric client setup
cli = Client(net_profile="connection.json")
org1_admin = cli.get_user('org1.example.com', 'Admin')
cli.new_channel('mychannel')

def create_medical_record(patient_id, doctor_id, diagnosis, treatment):
    """Create a medical record on the blockchain"""
    try:
        # Invoke chaincode
        response = cli.chaincode_invoke(
            requestor=org1_admin,
            channel_name='mychannel',
            peer_names=['peer0.org1.example.com'],
            args=[patient_id, doctor_id, diagnosis, treatment],
            cc_name='medical_records',
            fcn='CreateMedicalRecord'
        )
        return "Medical record created successfully"
    except Exception as e:
        return f"Error creating medical record: {str(e)}"

def get_medical_record(record_id):
    """Retrieve a medical record from the blockchain"""
    try:
        # Query chaincode
        response = cli.chaincode_query(
            requestor=org1_admin,
            channel_name='mychannel',
            peers=['peer0.org1.example.com'],
            args=[record_id],
            cc_name='medical_records',
            fcn='GetMedicalRecord'
        )
        return response
    except Exception as e:
        return f"Error retrieving medical record: {str(e)}"

# Gradio interface
with gr.Blocks() as demo:
    gr.Markdown("# Medical Records Management System")
    
    with gr.Tab("Create Record"):
        patient_id = gr.Textbox(label="Patient ID")
        doctor_id = gr.Textbox(label="Doctor ID")
        diagnosis = gr.Textbox(label="Diagnosis")
        treatment = gr.Textbox(label="Treatment")
        create_btn = gr.Button("Create Record")
        create_output = gr.Textbox(label="Result")
        
        create_btn.click(
            create_medical_record,
            inputs=[patient_id, doctor_id, diagnosis, treatment],
            outputs=create_output
        )
    
    with gr.Tab("View Record"):
        record_id = gr.Textbox(label="Record ID")
        view_btn = gr.Button("View Record")
        view_output = gr.JSON(label="Record Details")
        
        view_btn.click(
            get_medical_record,
            inputs=record_id,
            outputs=view_output
        )

if __name__ == "__main__":
    demo.launch() 