import React, { useEffect, useState } from "react";
import { Container, FormControl } from "react-bootstrap";
import { socket } from "../Socket";
import Tabla from "./Tabla";

const Redis = () => {
  const [lastGames, setLastGames] = useState([]);
  const [bestPlayers, setBestPlayers] = useState([]);
  const [player, setPlayer] = useState("");
  const [stats, setStats] = useState({
    id: 0,
    gameName: "",
    state: false,
    winner: 0,
  });

  useEffect(() => {
    socket.on("getLastGames", (value) => {
      setLastGames(value);
    });

    socket.on("getBestPlayers", (value) => {
      setBestPlayers(value);
    });

    socket.on("getPlayer", (value) => {
      setStats(value);
    });

    return () => {
      // This is the cleanup function
    };
  }, []);

  useEffect(() => {
    socket.emit("statplayer", player);
    return () => {
      // This is the cleanup function
    };
  });

  return (
    <div>
      <Container style={{ width: "35%", margin: "20px auto auto auto" }}>
        <FormControl
          type="number"
          value={player}
          onChange={(e) => setPlayer(e.target.value)}
        ></FormControl>
        <center>
          <h1>Estadisticas en tiempo real del jugador: {player}</h1>
        </center>
        <br />
        <Tabla
          row={2}
          headers={["id", "Game Name", "State"]}
          data={[
            {
              id: stats.id,
              name: stats.gameName,
              state: stats.state ? "win" : "Lost",
            },
          ]}
        />
      </Container>
      <Container>
        <center>
          <h1>Reportes Redis</h1>
        </center>
      </Container>
      <Container style={{ display: "flex" }}>
        <div style={{ width: "35%", margin: "20px auto auto auto" }}>
          <center>
            <h2>Ultimos juegos</h2>
          </center>
          <br />
          <Tabla
            headers={["id", "Game Name", "Winner", "Players", "Worker"]}
            data={lastGames}
          />
        </div>
        <div style={{ width: "35%", margin: "20px auto auto auto" }}>
          <center>
            <h2>Mejores Jugadores</h2>
          </center>
          <br />
          <Tabla headers={["Player", "Wins"]} data={bestPlayers} />
        </div>
      </Container>
    </div>
  );
};

export default Redis;
