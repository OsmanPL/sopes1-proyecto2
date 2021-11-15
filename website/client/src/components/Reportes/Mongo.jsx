import { socket } from "../Socket";
import React, { useEffect, useState } from "react";
import Tabla from "./Tabla";
import { Container } from "react-bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";
import Graphic from "./Graphic";

const Mongo = () => {
  const [data, setData] = useState([]);
  const [topJuegos, setTopJuegos] = useState([]);
  const [topWorker, setTopWorker] = useState([]);
  useEffect(() => {
    socket.on("getDataMongo", (value) => {
      setData(value);
    });

    socket.on("getTopWorker", (value) => {
      setTopWorker(value);
    });

    socket.on("getTopGames", (value) => {
      setTopJuegos(value);
    });

    return () => {
      // This is the cleanup function
    };
  }, []);

  const getLabels = (top) => {
    return top.map((e) => {
      return e._id;
    });
  };

  const getData = (top) => {
    return top.map((e) => {
      return e.count;
    });
  };

  return (
    <div>
      <Container>
        <center>
          <h1>Reportes Mongo</h1>
        </center>
      </Container>
      <br />
      <Container style={{ display: "flex" }}>
        <div style={{ width: "35%", margin: "20px auto auto auto" }}>
          <center>
            <h2>Top worker</h2>
          </center>
          <Graphic
            data={getData(topWorker)}
            label={getLabels(topWorker)}
            title={"Top Worker"}
          />
          <br />
        </div>
        <div style={{ width: "35%", margin: "20px auto auto auto" }}>
          <center>
            <h2>Top 3 Juegos</h2>
          </center>
          <Graphic
            data={getData(topJuegos)}
            label={getLabels(topJuegos)}
            title={"Top Juegos"}
          />
        </div>
      </Container>
      <div style={{ width: "35%", margin: "auto" }}>
        <center>
          <h2>Datos almacenados</h2>
        </center>
        <br />
        <Tabla
          headers={["id", "Game Name", "Winner", "Players", "Worker"]}
          data={data}
        />
      </div>
    </div>
  );
};

export default Mongo;
