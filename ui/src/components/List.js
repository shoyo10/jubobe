import React, { useState, useCallback, useEffect } from "react";
import Box from '@mui/material/Box';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemText from '@mui/material/ListItemText';
import DialogOrder from './Dialog';

export default function BasicList() {
    const [patients, setPatients] = useState([]);
    const [open, setOpen] = useState(false);
    const [selectedPatient, setSelectedPatient] = useState({});
    const [order, setOrder] = useState({});

    const fetchPatients = useCallback(async () => {
        const resp = await fetch("/api/patients");
        const body = await resp.json();
        const { Data } = body;

        setPatients(Data);
    }, [setPatients]);

    useEffect(() => {
        fetchPatients();
    }, [fetchPatients]);

    const handleListItemClick = async (patient) => {
        console.log(patient.Id);
        if (patient.OrderId > 0) {
            await fetchOrder(patient.OrderId);
        } else {
            setOrder({});
        }
        setOpen(true);
        setSelectedPatient(patient);
    };

    const handleClose = () => {
        fetchPatientsOnClose();
        setOpen(false);
    };

    const fetchPatientsOnClose = async () => {
        const resp = await fetch("/api/patients");
        const body = await resp.json();
        const { Data } = body;
        setPatients(Data);
    }

    const fetchOrder = async (id) => {
        const resp = await fetch("/api/orders/" + id);
        const body = await resp.json();
        const { Data } = body;

        setOrder(Data);
    };

    return (
        <Box sx={{ width: '100%', maxWidth: 720, bgcolor: 'background.paper' }}>
            <nav aria-label="secondary mailbox folders">
                <List>
                    {patients.map((patient) => (
                        <ListItem disablePadding key={patient.Name}>
                            <ListItemButton onClick={() => handleListItemClick(patient)}>
                                <ListItemText primary={patient.Name} />
                            </ListItemButton>
                        </ListItem>
                    ))}
                </List>
            </nav>
            <DialogOrder open={open} onClose={handleClose} patient={selectedPatient} order={order} />
        </Box>
    );
}
