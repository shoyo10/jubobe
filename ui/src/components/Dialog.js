import * as React from 'react';
import PropTypes from 'prop-types';
import DialogTitle from '@mui/material/DialogTitle';
import DialogContent from '@mui/material/DialogContent';
import Dialog from '@mui/material/Dialog';
import IconButton from '@mui/material/IconButton';
import AddIcon from '@mui/icons-material/Add';
import TextField from '@mui/material/TextField';
import DialogActions from '@mui/material/DialogActions';
import Button from '@mui/material/Button';

export default function DialogOrder(props) {
    const { onClose, patient, open, order } = props;
    const [showInput, setShowInput] = React.useState(false);

    const handleClose = async () => {
        onClose();
        setShowInput(false);
    };

    const handleAddClick = () => {
        setShowInput(true);
    };

    const submitOrder = async (msg) => {
        if (patient.OrderId > 0) {
            await fetch(`/api/orders/${patient.OrderId}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ Message: msg }),
            });
        } else {
            await fetch('/api/orders', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ PatientId: patient.Id, Message: msg }),
            });
        }
    }

    return (
        <Dialog
            onClose={handleClose}
            open={open}
            PaperProps={{
                component: 'form',
                onSubmit: async (event) => {
                    event.preventDefault();
                    const formData = new FormData(event.currentTarget);
                    const formJson = Object.fromEntries(formData.entries());
                    await submitOrder(formJson.message);
                    await handleClose();
                },
            }}
        >
            <DialogTitle>
                {patient.Name} Order
                <IconButton
                    onClick={handleAddClick}
                    style={{ position: 'absolute', right: 8, top: 8, color: 'gray' }}
                >
                    <AddIcon />
                </IconButton>
            </DialogTitle>
            <DialogContent>
                {(showInput || patient.OrderId > 0) &&
                    <TextField
                        id="outlined-multiline-static"
                        name="message"
                        label="請輸入醫囑"
                        multiline
                        rows={4}
                        defaultValue={order.Message}
                        variant="filled"
                    />}
            </DialogContent>
            <DialogActions>
                <Button onClick={handleClose}>Cancel</Button>
                <Button type="submit">Submit</Button>
            </DialogActions>
        </Dialog>
    );
}

DialogOrder.propTypes = {
    onClose: PropTypes.func.isRequired,
    patient: PropTypes.object.isRequired,
    open: PropTypes.bool.isRequired,
    order: PropTypes.object.isRequired,
};
