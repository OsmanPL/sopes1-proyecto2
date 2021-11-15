import React, { useState } from "react";
import { Container } from "react-bootstrap";
import Mongo from "./Mongo";
import Redis from "./Redis";
import "./buton.css";

const Reporte = () => {
  const [reporte, setReporte] = useState(true);
  return (
    <div>
      <br />
      <Container>
        <center>
          <div className="center">
            <input type="checkbox" onClick={() => setReporte(!reporte)}></input>
          </div>
        </center>
      </Container>
      {reporte && <Mongo />}
      {!reporte && <Redis />}
    </div>
  );
};

export default Reporte;
