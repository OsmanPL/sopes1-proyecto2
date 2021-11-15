import React from "react";
import TableScrollbar from "react-table-scrollbar";
import { Table } from "react-bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";

const Tabla = (props) => {
  const headers = props.headers;
  const data = props.data;

  return (
    <TableScrollbar rows={props.row || 10}>
      <Table striped bordered hover variant="light">
        <thead>
          <tr>
            {headers?.map((e, i) => (
              <th key={i}>{e}</th>
            ))}
          </tr>
        </thead>
        <tbody>
          {data?.map((e, i) => (
            <tr key={i}>
              {Object.keys(e).map((k, v) => (
                <td key={v}>{e[k]}</td>
              ))}
            </tr>
          ))}
        </tbody>
      </Table>
    </TableScrollbar>
  );
};

export default Tabla;
