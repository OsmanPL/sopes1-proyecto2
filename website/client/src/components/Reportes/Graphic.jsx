import React from "react";
import { Bar } from "react-chartjs-2";

const Graphic = (props) => {
  const config = {
    type: "bar",
    data: props.data,
    options: {
      scales: {
        y: {
          beginAtZero: true,
        },
      },
    },
  };

  const data = {
    labels: props.label,
    datasets: [
      {
        label: props.title,
        data: props.data,
        borderColor: [
          "rgb(255, 99, 132)",
          "rgb(255, 159, 64)",
          "rgb(255, 205, 86)",
        ],
        backgroundColor: [
          "rgba(255, 99, 132, 0.2)",
          "rgba(255, 159, 64, 0.2)",
          "rgba(255, 205, 86, 0.2)",
        ],
        borderWidth: 1
      },
    ],
  };

  return (
    <div>
      <Bar options={config} data={data} />
    </div>
  );
};

export default Graphic;
