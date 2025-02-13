import PageHeader from "../header/PageHeader";
import { useState } from "react";
import { useNavigate } from "react-router-dom"; // ✅ Import useNavigate
import axios from "axios";

function FlightCreate() {
    const [flight, setFlight] = useState({
        id: '14',
        number: '',
        airline_name: '',
        source: '',
        destination: '',
        capacity: 0,
        price: 0
    });

    const OnBoxChange = (event) => {
        const { id, value } = event.target;
        setFlight(prevFlight => ({
            ...prevFlight,
            [id]: value
        }));
    };

    const navigate = useNavigate();

    const OnCreate = async () => {
        try {
            const baseurl = 'http://localhost:8080';
            const response = await axios.post(
                `${baseurl}/flights`,
                { ...flight, capacity: parseInt(flight.capacity), price: parseFloat(flight.price) }
            );
            alert(response.data.message);
            navigate("/flights/list");
        } catch (error) {
            alert('Server error');
            console.error(error);
        }
    };

    return (
        <>
            <PageHeader PageNumber={2} />
            <h3>
                <a href="/flights/list" className="btn btn-light">Go Back</a> New Flight
            </h3>
            <div className="container">
                <div className="form-group mb-3">
                    <label htmlFor="number" className="form-label">Flight Number:</label>
                    <input type="text" className="form-control" id="number" placeholder="Please enter flight number" value={flight.number} onChange={OnBoxChange} />
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="airline_name" className="form-label">Airline Name:</label>
                    <input type="text" className="form-control" id="airline_name" placeholder="Please enter airline name" value={flight.airline_name} onChange={OnBoxChange} />
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="source" className="form-label">Source:</label>
                    <input type="text" className="form-control" id="source" placeholder="Please enter source" value={flight.source} onChange={OnBoxChange} />
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="destination" className="form-label">Destination:</label>
                    <input type="text" className="form-control" id="destination" placeholder="Please enter destination" value={flight.destination} onChange={OnBoxChange} />
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="capacity" className="form-label">Capacity (Number of Seats):</label>
                    <input type="number" className="form-control" id="capacity" placeholder="Please enter capacity" value={flight.capacity} onChange={OnBoxChange} />
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="price" className="form-label">Ticket Price:</label>
                    <input type="number" className="form-control" id="price" placeholder="Please enter ticket price" value={flight.price} onChange={OnBoxChange} />
                </div>
                <button className="btn btn-success" onClick={OnCreate}>Create Flight</button>
            </div>
        </>
    );
}

export default FlightCreate;
