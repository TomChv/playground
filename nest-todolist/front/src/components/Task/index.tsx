import React from "react";

import { NoteDto } from "../../dto/apiDto";

import "./index.css";

function Task(props: { task: NoteDto }) {
  return (
    <div className="task" key={props.task.title}>
      <p className="taskName">{props.task.title}</p>
      <div className="taskContent">
        <p>{props.task.content}</p>
      </div>
    </div>
  );
}

export default Task;
