import React, { useState } from "react";
import Requester from "../../services/apiSDK";

import "./index.css";

import { NoteDto, NoteRo } from "../../dto/apiDto";

function TaskForm(props: {
  taskList: NoteRo[];
  setTaskList: (val: NoteRo[]) => void;
  setIsAdding: (val: boolean) => void;
}): JSX.Element {
  const [task, setTask] = useState<NoteDto>({ title: "", content: "" });

  return (
    <div className="taskForm">
      <div className="formBox">
        <input
          className="titleInput"
          value={task.title}
          type="text"
          onChange={(event) => setTask({ ...task, title: event.target.value })}
          placeholder="Title"
        />
        <textarea
          value={task.content}
          className="contentInput"
          onChange={(event) =>
            setTask({ ...task, content: event.target.value })
          }
          placeholder="Tip your content here..."
        />
        <button
          className="formButton"
          onClick={async () => {
            if (task.title !== "" && task.content !== "") {
              const newTask = await Requester.sendNote(task);
              if (newTask.data)
                props.setTaskList([...props.taskList, newTask.data]);
              setTask({ title: "", content: "" });
              props.setIsAdding(false);
            }
          }}
          type="button"
        >
          Submit
        </button>
      </div>
    </div>
  );
}

export default TaskForm;
