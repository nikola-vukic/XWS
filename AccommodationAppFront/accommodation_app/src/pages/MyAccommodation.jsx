import React from "react";
import classes from "./Flights.module.css";
import utils from "./Utils.module.css";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Modal from "@mui/material/Modal";
import { useState, useRef, useEffect, useContext } from "react";
import dayjs, { Dayjs } from "dayjs";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";

const acc = {
  Id: 1,
  Host: 2,
  Name: "penthaus neki",
  HasWifi: true,
  HasFreeParking: true,
  HasWashingMachine: true,
  MinNumberOfGuests: 2,
  MaxNumberOfGuests: 7,
  Availability: [
    {
      Price: 20,
      IsPricePerGuest: true,
    },
    {
      Price: 60,
      IsPricePerGuest: true,
    },
    {
      Price: 60,
      IsPricePerGuest: true,
    },
    {
      Price: 40,
      IsPricePerGuest: true,
    },
  ],
  IsReservationAcceptenceManual: true,
};
const style = {
  position: "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  bgcolor: "background.paper",
  boxShadow: 24,
  borderRadius: 3,
};
const MyAccommodation = () => {
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const today = new Date();
  const tomorrow = new Date(today);
  tomorrow.setDate(tomorrow.getDate() + 1);
  const [value, setValue] = useState(dayjs(tomorrow));
  return (
    <div className={classes.body}>
      <div className={classes.home}>
        <br></br>
        <h1>Accommodation name</h1>

        <div className={classes.buttonContainerRight}>
          <button className={utils.greenButton} onClick={handleOpen}>
            Add
          </button>
        </div>
        <table className={classes.styledTable}>
          <thead>
            <tr>
              <th>Starting date</th>
              <th>Ending date</th>
              <th>Price</th>
              <th>Pricing</th>
            </tr>
          </thead>
          <tbody>
            {acc.Availability?.map((app) => (
              <tr key={app.id}>
                <td>{app.Price}</td>
                <td>{app.Price}</td>
                <td>{app.Price}</td>
                <td>{app.Price}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <div>
            <div className={classes.modalTitle}>Add pricing</div>
            <form className={classes.register}>
              <div className={classes.spanReserve}>
                <label>Starting date:</label>
                <LocalizationProvider dateAdapter={AdapterDayjs}>
                  <DatePicker
                    value={value}
                    onChange={(newValue) => {
                      setValue(newValue);
                    }}
                    className={classes.DatePicker}
                  />
                </LocalizationProvider>
              </div>
              <div className={classes.spanReserve}>
                <label>Ending date:</label>
                <LocalizationProvider dateAdapter={AdapterDayjs}>
                  <DatePicker
                    value={value}
                    onChange={(newValue) => {
                      setValue(newValue);
                    }}
                    className={classes.DatePicker}
                  />
                </LocalizationProvider>
              </div>
              <div className={classes.spanReserve}>
                <label>Price:</label>
                <input className={utils.inputPrice}></input>
              </div>
              <div className={classes.spanReserve}>
                <label>Pricing:</label>

                <input
                  type="radio"
                  id="unit"
                  name="pricing"
                  value="unit"
                ></input>
                <label for="html">Per unit</label>
                <input
                  type="radio"
                  id="person"
                  name="pricing"
                  value="person"
                ></input>
                <label for="html">Per person</label>
              </div>
              <button className={classes.reserveButton}>Add</button>
            </form>
          </div>
        </Box>
      </Modal>
    </div>
  );
};

export default MyAccommodation;